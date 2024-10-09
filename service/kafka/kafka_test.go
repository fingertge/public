// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/16 15:03:34
// * Proj: work
// * Pack: kafka
// * File: kafka_test.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"testing"
	"time"
)

func TestKafka(t *testing.T) {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"192.168.31.28:9002"},
		Partition: 0,
		GroupID:   "test-001",
		Topic:     "test",
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	//r.SetOffset(-2)

	for {
		m, e := r.ReadMessage(context.Background())
		if e != nil {
			t.Error(e)
			return
		}
		t.Logf("mess:%+v", m)
		off := r.Offset()
		t.Logf("got offset:%d", off)
	}

}

func TestKafkaWriter_SendMessage(t *testing.T) {
	w := NewKafkaClient("192.168.31.28:9002", 3*time.Second)
	headers := []kafka.Header{
		{
			Key:   "appid",
			Value: []byte("1022"),
		},
		{
			Key:   "gameId",
			Value: []byte("1022"),
		},
	}
	err := w.SendMessage("test", []byte("this is test message"), headers)
	if err != nil {
		t.Error(err)
	}
}

func Execute(val []byte, headers []kafka.Header) {
	fmt.Printf("val:%s\nHeaders:\n", val)
	for _, v := range headers {
		fmt.Printf("key:%s, val:%s\n", v.Key, v.Value)
	}
}

func TestKafkaReader_ReadMessages(t *testing.T) {
	r := NewKafkaReader([]string{"192.168.31.28:9002"}, "test", Execute)
	r.ReadMessages()
}
