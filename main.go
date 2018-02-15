package main

import (
	"log"
	"sync"

	nsq "github.com/nsqio/go-nsq"
)

func main() {
	wg, config := sync.WaitGroup{}, nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}
	err = producer.Publish("write-test2", []byte("bitchhhhh"))
	if err != nil {
		log.Fatal(err)
	}
	consumer, err := nsq.NewConsumer("write-test2", "default", config)
	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println("test message received: ", string(message.Body))
		wg.Done()
		return nil
	}))
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
