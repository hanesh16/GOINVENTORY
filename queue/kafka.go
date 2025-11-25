package queue

import (
	"fmt"

	"github.com/IBM/sarama"
)

func ProduceKafkaMessage(topic, message string, producer sarama.SyncProducer) error {
	messageToSend := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(messageToSend)
	if err != nil {
		return err
	}

	fmt.Printf("Kafka message sent → topic=%s partition=%d offset=%d\n", topic, partition, offset)
	return nil
}
