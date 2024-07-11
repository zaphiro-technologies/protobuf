package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
	grid "github.com/zaphiro-technologies/protobuf/go/grid/v1"
	"google.golang.org/protobuf/proto"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Printf("%s ", err)
		os.Exit(1)
	}
}

func consumerClose(channelClose stream.ChannelClose) {
	event := <-channelClose
	fmt.Printf("Consumer: %s closed on the stream: %s, reason: %s \n", event.Name, event.StreamName, event.Reason)
}

func generateDataID(pmuId int, nMeasures int) map[string]*grid.Data {
	dataID := map[string]*grid.Data{}
	for k := 0; k < nMeasures; k++ {
		id := fmt.Sprintf("Dev%04d-%04d", pmuId, k)
		dataTypeInt := 0
		switch {
		case k < 1:
			dataTypeInt = 13
		case k < 2:
			dataTypeInt = 28
		case k < 4:
			dataTypeInt = 21
		case k < 9:
			dataTypeInt = 20
		case k < 11:
			dataTypeInt = 44
		case k < 15:
			dataTypeInt = 30
		case k < 18:
			dataTypeInt = 31
		case k%2 == 0:
			dataTypeInt = 43
		case k%2 == 1:
			dataTypeInt = 44
		}
		dataType := grid.DataType(dataTypeInt)
		dataID[id] = &grid.Data{
			DataType:   dataType,
			Value:      new(uint64),
			MeasuredAt: time.Now().UnixMilli(),
		}
	}
	return dataID
}

func generatePmuDataID(nPmu int, nMeasures int) []*grid.DataSet {
	pmuData := []*grid.DataSet{}
	for k := 0; k < nPmu; k++ {
		id := fmt.Sprintf("Dev%04d", k)
		pmuData = append(
			pmuData,
			&grid.DataSet{ProducerId: id, Data: generateDataID(k, nMeasures)},
		)
	}
	return pmuData
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	// Set log level, not mandatory by default is INFO
	// you cn set DEBUG for more information
	// stream.SetLevelInfo(logs.DEBUG)

	fmt.Println("Getting started with Streaming client for RabbitMQ")
	fmt.Println("Connecting to RabbitMQ streaming ...")

	// Connect to the broker ( or brokers )

	env, err := stream.NewEnvironment(
		stream.NewEnvironmentOptions().
			SetHost("localhost").
			SetPort(5552).
			SetUser("guest").
			SetPassword("guest"))
	CheckErr(err)

	// The default stream for measurements in a prod set-up, you don't have to
	// create it

	streamName := "measurements"
	err = env.DeclareStream(streamName,
		&stream.StreamOptions{
			MaxLengthBytes: stream.ByteCapacity{}.GB(2),
		},
	)

	CheckErr(err)

	// Get a new producer for a stream
	producer, err := env.NewProducer(streamName, nil)
	CheckErr(err)

	// the send method automatically aggregates the messages
	// based on batch size
	for i := 0; i < 10; i++ {
		data := generatePmuDataID(1, 10)[0]
		buf, _ := proto.Marshal(data)
		message := amqp.NewMessage(buf)
		messageProperties := make(map[string]interface{})
		messageProperties["Id"] = uuid.New().String()
		messageProperties["type"] = "DataSet"
		messageProperties["timestampId"] = int64(math.Round(float64(time.Now().UnixMilli())/20) * 20)
		messageProperties["producerId"] = data.GetProducerId()
		message.ApplicationProperties = messageProperties
		err := producer.Send(message)
		CheckErr(err)
	}

	// this sleep is not mandatory, just to show the confirmed messages
	time.Sleep(1 * time.Second)
	err = producer.Close()
	CheckErr(err)

	fmt.Println("Sent 10 messages")

	handleMessages := func(consumerContext stream.ConsumerContext, message *amqp.Message) {

		data := &grid.DataSet{}
		err := proto.Unmarshal(message.GetData(), data)
		for measurement_id, measurement := range data.GetData() {
			fmt.Printf("consumer name: %s, measurement_id: %s, measurement_time %d, measurement_type %d, measurement_value %d\n ",
				consumerContext.Consumer.GetName(), measurement_id, measurement.DataType, *measurement.Value, measurement.MeasuredAt)
		}
		CheckErr(err)
	}

	consumer, err := env.NewConsumer(
		streamName,
		handleMessages,
		stream.NewConsumerOptions().
			SetClientProvidedName("my_consumer").            // connection name
			SetConsumerName("my_consumer").                  // set a consumer name
			SetOffset(stream.OffsetSpecification{}.First())) // start consuming from the beginning

	CheckErr(err)
	channelClose := consumer.NotifyClose()
	// channelClose receives all the closing events, here you can handle the
	// client reconnection or just log
	defer consumerClose(channelClose)

	fmt.Println("Press any key to stop ")
	_, _ = reader.ReadString('\n')
	err = consumer.Close()
	time.Sleep(200 * time.Millisecond)
	CheckErr(err)
	err = env.DeleteStream(streamName)
	CheckErr(err)
	err = env.Close()
	CheckErr(err)
}
