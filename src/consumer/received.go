package main

import (
	"bufio"
	"fmt"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
	"github.com/thebigyovadiaz/rabbitmq-stream/src/util"
	"os"
)

func main() {
	env, err := stream.NewEnvironment(
		stream.NewEnvironmentOptions().
			SetHost("localhost").
			SetPort(5552).
			SetUser("guest").
			SetPassword("guest"))

	util.LogFailOnError(err)
	util.LogSuccessful("Stream opened successfully")

	streamName := "hello-go-stream"
	err = env.DeclareStream(streamName,
		&stream.StreamOptions{
			MaxLengthBytes: stream.ByteCapacity{}.GB(2),
		},
	)

	util.LogFailOnError(err)
	util.LogSuccessful("Stream created successfully")

	messagesHandler := func(consumerContext stream.ConsumerContext, message *amqp.Message) {
		util.LogSuccessful(fmt.Sprintf("Stream: %s - Received message: %s\n", consumerContext.Consumer.GetStreamName(),
			message.Data))
	}

	consumer, err := env.NewConsumer(streamName, messagesHandler,
		stream.NewConsumerOptions().SetOffset(stream.OffsetSpecification{}.First()))

	util.LogFailOnError(err)
	util.LogSuccessful("Consumer created successfully")

	reader := bufio.NewReader(os.Stdin)
	util.LogSuccessful(" [x] Waiting for messages. enter to close the consumer")

	_, _ = reader.ReadString('\n')
	err = consumer.Close()
	util.LogFailOnError(err)
}
