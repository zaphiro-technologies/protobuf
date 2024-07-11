package main

import (
	"bufio"
	"context"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	grid "github.com/zaphiro-technologies/protobuf/go/grid/v1"
	"google.golang.org/protobuf/proto"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Printf("%s ", err)
		os.Exit(1)
	}
}

func generateFault(
	fault_id string,
	event_type grid.FaultEventType,
	measurement_timestamp int64,
	fault_current float32,
	location_probability *float32,
	used_measurements []*grid.FaultMeasurement,
	impacted_equipments []string,
	faulty_equipment *string,
) *grid.Fault {
	description := "Example fault"

	tm := math.Round(float64(time.Now().UnixMilli())/20) * 20

	fault := grid.Fault{
		Description:          &description,
		Id:                   fault_id,
		Kind:                 grid.PhaseConnectedFaultKind_PHASE_CONNECTED_FAULT_KIND_LINE_TO_GROUND,
		Phases:               grid.PhaseCode_PHASE_CODE_ABC,
		UpdatedAt:            int64(tm),
		FaultEventType:       event_type,
		FaultyEquipmentId:    faulty_equipment,
		MeasurementTimestamp: &measurement_timestamp,
		FaultCurrent:         &fault_current,
		UsedMeasurementIds:   used_measurements,
		LocationProbability:  location_probability,
		ImpactedEquipmentIds: impacted_equipments,
	}

	return &fault
}

func generateLineFault(
	fault_id string,
	event_type grid.FaultEventType,
	measurement_timestamp int64,
	fault_current float32,
	location_probability *float32,
	used_measurements []*grid.FaultMeasurement,
	impacted_equipments []string,
	faulty_equipment *string,
) *grid.LineFault {

	fault := generateFault(
		fault_id,
		event_type,
		measurement_timestamp,
		fault_current,
		location_probability,
		used_measurements,
		impacted_equipments,
		faulty_equipment,
	)

	length_from_terminal1 := rand.Float32()
	length_uncertainty := float32(100.0)

	lineFault := grid.LineFault{
		Fault:               fault,
		LengthFromTerminal1: &length_from_terminal1,
		LengthUncertainty:   &length_uncertainty,
	}

	return &lineFault
}

func publishMessage(ch *amqp.Channel, ctx context.Context, header amqp.Table, buf []byte) {
	err := ch.PublishWithContext(ctx,
		"fault", // exchange
		"",      // routing key
		false,   // mandatory
		false,
		amqp.Publishing{
			Headers:      header,
			DeliveryMode: amqp.Persistent,
			Body:         buf,
		})
	CheckErr(err)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	// Set log level, not mandatory by default is INFO
	// you cn set DEBUG for more information
	// stream.SetLevelInfo(logs.DEBUG)

	fmt.Println("Getting started with AMPQ client for RabbitMQ")
	fmt.Println("Connecting to RabbitMQ ...")

	// Connect to the broker ( or brokers )

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	CheckErr(err)
	defer conn.Close()
	ch, err := conn.Channel()
	CheckErr(err)
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"fault",   // name
		"headers", // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)

	ctx := context.Background()

	q, err := ch.QueueDeclare(
		"fault-storer", // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	CheckErr(err)

	go func() {
		for d := range msgs {

			if d.Headers["type"] == "Fault" {
				data := &grid.Fault{}
				err := proto.Unmarshal(d.Body, data)
				CheckErr(err)
				fmt.Printf(
					"Received a fault event message: %s, event type: %s\n",
					data.Id,
					data.FaultEventType,
				)
			}

			if d.Headers["type"] == "LineFault" {
				data := &grid.LineFault{}
				err := proto.Unmarshal(d.Body, data)
				CheckErr(err)
				fmt.Printf(
					"Received a line fault event message: %s, event type: %s, faulty line: %s, probability: %f, length from t1: %f\n",
					data.Fault.Id,
					data.Fault.FaultEventType, *data.Fault.FaultyEquipmentId, *data.Fault.LocationProbability, *data.LengthFromTerminal1,
				)
			}
		}
	}()

	//send messsages

	fault_id := uuid.Must(uuid.NewRandom()).String()
	mt := int64(math.Round(float64(time.Now().UnixMilli())/20) * 20)
	fault_current := float32(100000.0)
	fault_equipment := new(string)
	location_probability := new(float32)
	impacted_equipments := []string{"EQ-1", "EQ-2", "EQ-3"}
	used_measurements := []*grid.FaultMeasurement{
		{PositiveSign: true, MeasurementID: "M-1"},
		{PositiveSign: true, MeasurementID: "M-1"},
		{PositiveSign: true, MeasurementID: "M-1"},
	}

	startFaultEvent := generateFault(
		fault_id,
		grid.FaultEventType_FAULT_EVENT_TYPE_STARTED,
		mt,
		fault_current,
		location_probability,
		used_measurements,
		impacted_equipments,
		fault_equipment,
	)

	header := amqp.Table{
		"id":         fault_id,
		"producerId": "FL",
		"type":       "Fault",
	}
	buf, _ := proto.Marshal(startFaultEvent)
	publishMessage(ch, ctx, header, buf)

	time.Sleep(20 * time.Millisecond)

	endFaultEvent := generateFault(
		fault_id,
		grid.FaultEventType_FAULT_EVENT_TYPE_ENDED,
		mt,
		fault_current,
		location_probability,
		used_measurements,
		impacted_equipments,
		fault_equipment,
	)
	buf, _ = proto.Marshal(endFaultEvent)
	publishMessage(ch, ctx, header, buf)

	// first potential location
	equipment := "EQ-1"
	probability := float32(0.33)
	header = amqp.Table{
		"id":         fault_id,
		"producerId": "FL",
		"type":       "LineFault",
	}

	location := generateLineFault(
		fault_id,
		grid.FaultEventType_FAULT_EVENT_TYPE_LOCATED,
		mt,
		fault_current,
		&probability,
		used_measurements,
		impacted_equipments,
		&equipment,
	)
	buf, _ = proto.Marshal(location)
	publishMessage(ch, ctx, header, buf)

	// second potential location
	equipment = "EQ-2"

	location = generateLineFault(
		fault_id,
		grid.FaultEventType_FAULT_EVENT_TYPE_LOCATED,
		mt,
		fault_current,
		&probability,
		used_measurements,
		impacted_equipments,
		&equipment,
	)
	buf, _ = proto.Marshal(location)
	publishMessage(ch, ctx, header, buf)

	// third potential location
	equipment = "EQ-3"

	location = generateLineFault(
		fault_id,
		grid.FaultEventType_FAULT_EVENT_TYPE_LOCATED,
		mt,
		fault_current,
		&probability,
		used_measurements,
		impacted_equipments,
		&equipment,
	)
	buf, _ = proto.Marshal(location)
	publishMessage(ch, ctx, header, buf)

	fmt.Println("Press any key to stop\n")
	_, _ = reader.ReadString('\n')
	channelCloseQueue := make(chan *amqp.Error)
	ch.NotifyClose(channelCloseQueue)
	time.Sleep(200 * time.Millisecond)
	CheckErr(err)
}
