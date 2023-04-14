/* execute in mysql */

create table product (
	id int not null auto_increment,
	name varchar(100) not null,
	price float not null,
	created_at datetime,
	updated_at datetime,
	deleted_at datetime,
	primary key (id)
)

create table voucher (
	id int not null auto_increment,
	code varchar(100) not null,
	persen float not null,
	created_at datetime,
	updated_at datetime,
	deleted_at datetime,
	primary key (id)
)

create table transaction (
	id int not null auto_increment,
	transactionnumber int not null unique,
	name varchar(100) not null,
	quantity int not null,
	discount float not null,
	total float not null,
	pay float not null,
	primary key (id)
)

create table transaction_detail (
	id int not null auto_increment,
	transaction_id int not null,
	item varchar(100) not null,
	price float not null,
	quantity int not null,
	total float not null,
	primary key (id),
	foreign key (transaction_id) references Transaction(id)
)

alter table transaction rename column transactionnumber to transaction_number

alter table transaction modify column transaction_number int not null
