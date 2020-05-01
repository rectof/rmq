package main

import (
	"github.com/adjust/rmq"
)

func main() {
	connection, err := rmq.OpenConnection("cleaner", "tcp", "localhost:6379", 2)
	if err != nil {
		panic(err)
	}

	queue := connection.OpenQueue("things")
	queue.PurgeReady()
}
