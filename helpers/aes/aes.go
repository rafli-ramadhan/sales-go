package aes

import (
	"log"
	"strconv"

	hashids "github.com/speps/go-hashids"
)

var hd *hashids.HashIDData
var salt string
var minLength int

func initialize() {
	if hd != nil {
		return
	}

	hd = hashids.NewData()
	salt = "UnpPHAAddqRdEDaTZOu4BkZHZqbJmcAWMEeRvTSV86t4DZixSnjb5P7JOOfGPA0afqhOjcUVcgdLZHR8fxhoHYACiRapwUCvDHNT0etqqWD6qZeQP9R3kbCGW0hhaKXO"
	minLengthStr := "32"

	if salt == "" || minLengthStr == "" {
		log.Println("aes: env not found: AES_KEY or AES_MIN_LENGTH")
	}

	minLength, _ = strconv.Atoi(minLengthStr)
}

func Encrypt(id int) string {
	initialize()
	hd.Salt = salt
	hd.MinLength = minLength
	h, _ := hashids.NewWithData(hd)
	encoded, _ := h.Encode([]int{id})
	return encoded
}

func Decrypt(data string) int {
	initialize()
	hd.Salt = salt
	hd.MinLength = minLength
	h, _ := hashids.NewWithData(hd)
	d, err := h.DecodeWithError(data)
	if err != nil || len(d) < 1 {
		return -1
	}
	return d[0]
}