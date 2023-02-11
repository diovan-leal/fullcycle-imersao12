package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"log"
	"fmt"
)
	
type kafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

//like an factory
func NewKafkaConsumer(msgChan chan *ckafka.Message) *kafkaConsumer {
	return &kafkaConsumer{
		MsgChan: msgChan,
	}
}

func (k *kafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id": os.Getenv("KafkaConsumerGroupId"),
	}

	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatal("error consuming  kafka message" + err.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopic")}
	c.SubscribeTopics(topics, nil)

	fmt.Println("kafka consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}