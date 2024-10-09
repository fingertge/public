// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/16 14:59:19
// * Proj: work
// * Pack: kafka
// * File: kafka.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

type KafkaWriter struct {
	*kafka.Writer
	timeOut time.Duration
}

func NewKafkaClient(addr string, timeOut time.Duration) *KafkaWriter {
	return &KafkaWriter{
		Writer: &kafka.Writer{
			Addr:                   kafka.TCP(addr),
			Balancer:               &kafka.LeastBytes{},
			AllowAutoTopicCreation: true,
		},
		timeOut: timeOut,
	}
}

func (k *KafkaWriter) SendMessage(topic string, msg []byte, header []kafka.Header) error {
	message := kafka.Message{
		Topic:   topic,
		Value:   msg,
		Headers: header,
	}
	return k.WriteMessages(context.Background(), message)
}

type KafkaReaderExecute func(val []byte, header []kafka.Header)

type KafkaReader struct {
	*kafka.Reader
	execute KafkaReaderExecute
}

func NewKafkaReader(addrs []string, topic string, execute KafkaReaderExecute) *KafkaReader {
	return &KafkaReader{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  addrs,
			GroupID:  "reaed-001",
			Topic:    topic,
			MinBytes: 10e3,
			MaxBytes: 10e6,
		}),
		execute: execute,
	}
}

func (k *KafkaReader) ReadMessages() {
	for {
		msg, err := k.Reader.ReadMessage(context.Background())
		if err != nil {
			continue
		}
		k.execute(msg.Value, msg.Headers)
	}
}
