package main

import (
	kafka2 "github.com/diovan-leal/fullcycle-imersao12-simulator/application/kafka"
	"github.com/diovan-leal/fullcycle-imersao12-simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func init () {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main () {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)

	//async process gorountine parallel threads
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		//async process gorountine parallel threads
		go kafka2.Produce(msg)
	}

	/* producer := kafka.NewKafkaProducer()
	kafka.Publish("Ola", "readtest", producer) */

	for {
		_ = 1
	}

	//route := route2.Route{
	//	ID: "1",
	//	ClientID: "1",
	//}
	//route.LoadPositions()
	//stringjson, _ := route.ExportJsonPositions()

	//fmt.Println(stringjson[1])
}