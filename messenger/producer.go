package messenger

import "github.com/Shopify/sarama"

// Producer implements sarama.AsyncProducer
type Producer struct {
	producer sarama.AsyncProducer
}

// NewProducer returns a new instance of a Kafka Producer. connectionString is the connection information. topics is an
// array of kafka topics. successCallback is a function that is performed when a message is successfully published.
// errorCallback is a function that is performed when a message fails to publish.
func NewProducer(connectionString string, topics []string, successCallback, errorCallback func()) (sarama.AsyncProducer, error) {
	// initialize configurations for kafka producer
	config := sarama.NewConfig()
	if successCallback != nil {
		config.Producer.Return.Successes = true
	}
	// initialize new AsyncProducer
	producer, err := sarama.NewAsyncProducer([]string{connectionString}, config)
	if err != nil {
		return nil, err
	}
	// spin up a goroutine to constantly read for errors, if there is an error, invoke errorCallback if it isn't nil.
	go func() {
		for _ = range producer.Errors() {
			if errorCallback != nil {
				errorCallback()
			}
		}
	}()
	// spin up a goroutine if we want to return successes, then invoke successCallback.
	if config.Producer.Return.Successes {
		go func() {
			for _ = range producer.Successes() {
				successCallback()
			}
		}()
	}
	return producer, nil
}

// Close handles closing and cleaning up the producer. Call this if you're done using the producer
func (p *Producer) Close() {
	if p.producer != nil {
		err := p.producer.Close()
		if err != nil {
			// TODO: log error
		}
	}
}

// Publish handles sending the message to the producer queue. message is the value sent and topic is the destination
// topic the message is sent to.
func (p *Producer) Publish(message, topic string) {
	producerMessage := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(message)}
	p.producer.Input() <- producerMessage
}
