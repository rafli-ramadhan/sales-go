package transaction

import (
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"sales-go/config"
	"sales-go/model"
	"sales-go/publisher"
)

type repositoryhttpgormpostgresql struct {
	db *gorm.DB
}

func NewPostgreSQLGormHTTPRepository(db *gorm.DB) *repositoryhttpgormpostgresql {
	return &repositoryhttpgormpostgresql{
		db: db,
	}
}

func (repo *repositoryhttpgormpostgresql) GetTransactionByNumber(transactionNumber int) (result []model.TransactionDetail, err error) {
	defer repo.db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// GetTransaction
	query := `SELECT id, transaction_number, name, quantity, discount, total, pay FROM transaction WHERE transaction_number = $1`
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.QueryContext(ctx, transactionNumber)
	if err != nil {
		return
	}

	var transaction model.Transaction
	for res.Next() {
		res.Scan(&transaction.Id, &transaction.TransactionNumber, &transaction.Name, &transaction.Quantity, &transaction.Discount, &transaction.Total, &transaction.Pay)
	}

	// GetTransactionDetail
	query2 := `SELECT id, item, price, quantity, total FROM transaction_detail WHERE transaction_id = $1`
	stmt2, err := repo.db.PrepareContext(ctx, query2)
	if err != nil {
		return
	}

	res2, err := stmt2.QueryContext(ctx, transaction.Id)
	if err != nil {
		return
	}

	for res2.Next() {
		var temp model.TransactionDetail
		res2.Scan(&temp.Id, &temp.Item, &temp.Price, &temp.Quantity, &temp.Total)
		// append transaction in each of transaction detail
		temp.Transaction = transaction
		result = append(result, temp)
	}

	return
}

func (repo *repositoryhttpgormpostgresql) CreateBulkTransactionDetail(voucher model.VoucherRequest, listTransactionDetail []model.TransactionDetail, req model.TransactionDetailBulkRequest) (res []model.TransactionDetail, err error) {
	sendingData := model.TransactionDataRabbitMQ{
		voucher, 
		listTransactionDetail, 
		req,
	}

	// publish data to RabbitMQ
	err = publisher.Publish(sendingData)
	if err != nil {
		err = fmt.Errorf("error publish data to RabbitMQ : %s", err.Error())
		return
	}

	time.Sleep(3 * time.Second)

	// get response after publish data
	config, err := config.LoadConfig()
	if err != nil {
		err = fmt.Errorf("failed to load config : %s", err)
		return
	}

	conn, err := amqp.Dial(config.RabbitMQURL)
	if err != nil {
		err = fmt.Errorf("failed to connect to RabbitMQ : %s", err.Error())
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		err = fmt.Errorf("failed to open a channel : %s", err.Error())
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"create_transaction_response", // name
		true,                 // durable
		false,                // auto delete queue when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		err = fmt.Errorf("failed to declare a queue : %s", err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		err = fmt.Errorf("failed to register a consumer : %s", err)
		return
	}

	// worker to receive value from variable msgs
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			err := json.Unmarshal(d.Body, &res)
			if err != nil {
				break
			}

			d.Ack(false)
		}
		if err != nil {
			return
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	return
}

