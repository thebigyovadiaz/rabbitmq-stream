package main

import (
	"bufio"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
	"github.com/thebigyovadiaz/rabbitmq-stream/src/util"
	"os"
	"time"
)

func main() {
	env, err := stream.NewEnvironment(
		stream.NewEnvironmentOptions().
			SetHost("localhost").
			SetPort(5552).
			SetUser("guest").
			SetPassword("guest").
			SetRPCTimeout(5 * time.Second))

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

	producer, err := env.NewProducer(streamName, stream.NewProducerOptions())
	util.LogFailOnError(err)
	util.LogSuccessful("Producer created successfully")

	// Send a message
	err = producer.Send(amqp.NewMessage([]byte("Hello world")))
	util.LogFailOnError(err)
	util.LogSuccessful(" [x] 'Hello world' Message sent")

	reader := bufio.NewReader(os.Stdin)
	util.LogSuccessful(" [x] Press enter to close the producer")

	_, _ = reader.ReadString('\n')
	err = producer.Close()
	util.LogFailOnError(err)
}
