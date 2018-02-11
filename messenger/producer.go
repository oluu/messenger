package messenger

import "github.com/Shopify/sarama"

// NewProducer returns a new instance of a Kafka Producer. connectionString is the connection information. topics is an
// array of kafka topics. successCallback is a function that is performed when a message is successfully published.
// errorCallback is a function that is performed when a message fails to publish.
func NewProducer(connectionString string, topics []string, successCallback func(), errorCallback func()) (sarama.AsyncProducer, error) {
	// initialize configurations for kafka
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
