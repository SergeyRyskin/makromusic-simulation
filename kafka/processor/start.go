package main

import (
	"fmt"
    
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main()  {
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
    })
    if err != nil {
        fmt.Printf("Failed to create producer: %s\n", err)
    }
    fmt.Printf("%v\n", p)
}