package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

type MessageHandle struct {
}

func (mh *MessageHandle) HandleMessage(msg *nsq.Message) error {
	log.Println("get message: ", string(msg.Body))
	time.Sleep(time.Second)
	return nil
}

func main() {
	topic := "testAlvin"
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		i := 0
		for {
			if i >= 500 {
				break
			}
			i++
			err := producer.Publish(topic, []byte(fmt.Sprintf("test-%d", i)))
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	config.MaxInFlight = 100
	//config.MsgTimeout = time.Second
	consumer, err := nsq.NewConsumer(topic, "test-01", config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(&MessageHandle{})

	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Fatal(err)
	}

	/*go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("stats",consumer.Stats().Connections," ",consumer.Stats().MessagesRequeued," ",consumer.Stats().MessagesFinished," ",consumer.Stats().MessagesReceived)
		}
	}()*/

	select {}
}
