# Go Faults Example

This example demonstrates how to produce and consume fault event messages using Protocol Buffers and RabbitMQ AMQP exchanges.

## Overview

This example shows how to:
- Connect to RabbitMQ using AMQP 0.9.1
- Create an exchange for fault events
- Produce fault event messages (Fault and LineFault)
- Consume and decode fault messages based on message type
- Handle different fault event types (STARTED, ENDED, LOCATED)

## Requirements

- Go 1.21 or higher
- Docker (for running RabbitMQ)
- RabbitMQ with AMQP support

## Running the Example

1. Start the RabbitMQ infrastructure:
   ```bash
   make docker-start
   ```

2. Run the example:
   ```bash
   make example-faults-go
   ```

3. The example will:
   - Connect to RabbitMQ on port 5672
   - Create a headers exchange called "fault"
   - Create a queue called "fault-storer"
   - Send fault event messages
   - Consume and display fault events
   - Wait for user input to stop

4. Press any key to stop the example

## Expected Output

```
Getting started with AMQP client for RabbitMQ
Connecting to RabbitMQ ...
Received a fault event message: 1948cd36-f835-4dc5-83c2-ba5a7612362a, event type: FAULT_EVENT_TYPE_STARTED
Press any key to stop

Received a fault event message: 1948cd36-f835-4dc5-83c2-ba5a7612362a, event type: FAULT_EVENT_TYPE_ENDED
Received a line fault event message: 1948cd36-f835-4dc5-83c2-ba5a7612362a, event type: FAULT_EVENT_TYPE_LOCATED, faulty line: EQ-1, probability: 0.330000, length from t1: 0.526424
Received a line fault event message: 1948cd36-f835-4dc5-83c2-ba5a7612362a, event type: FAULT_EVENT_TYPE_LOCATED, faulty line: EQ-2, probability: 0.330000, length from t1: 0.484818
Received a line fault event message: 1948cd36-f835-4dc5-83c2-ba5a7612362a, event type: FAULT_EVENT_TYPE_LOCATED, faulty line: EQ-3, probability: 0.330000, length from t1: 0.315550
```

## Code Structure

### Producer

The producer:
1. Connects to RabbitMQ using AMQP 0.9.1
2. Declares a headers exchange named "fault"
3. Generates fault events with different types:
   - **FAULT_EVENT_TYPE_STARTED**: Initial fault detection
   - **FAULT_EVENT_TYPE_ENDED**: Fault cleared
   - **FAULT_EVENT_TYPE_LOCATED**: Fault location identified (with probability)
4. Marshals protobuf messages and publishes with headers
5. Demonstrates both generic `Fault` and specific `LineFault` messages

### Consumer

The consumer:
1. Creates a durable queue "fault-storer"
2. Consumes messages with auto-acknowledgment
3. Unmarshals messages based on the "type" header:
   - **"Fault"**: Generic fault event
   - **"LineFault"**: Line-specific fault with location details
4. Prints relevant fault information

## Key Protocol Buffer Messages

### Fault

Represents a generic fault event with:
- `Id`: Unique fault identifier (UUID)
- `Description`: Human-readable description
- `Kind`: Fault kind (e.g., LINE_TO_GROUND)
- `Phases`: Affected phases (e.g., ABC)
- `FaultEventType`: Event type (STARTED, ENDED, LOCATED)
- `FaultCurrent`: Magnitude of fault current
- `LocationProbability`: Probability of fault location (for LOCATED events)
- `ImpactedEquipmentIds`: List of affected equipment
- `UsedMeasurementIds`: Measurements used for fault detection
- `MeasurementTimestamp`: When the fault was measured
- `UpdatedAt`: Last update timestamp

### LineFault

Extends `Fault` for line-specific faults with:
- `Fault`: Embedded fault information
- `LengthFromTerminal1`: Distance from terminal 1 (normalized)
- `LengthUncertainty`: Uncertainty in meters

## Fault Event Lifecycle

The example demonstrates a typical fault event sequence:

1. **STARTED**: Fault is detected in the system
2. **ENDED**: Fault condition has cleared
3. **LOCATED** (multiple): Possible fault locations with probabilities
   - Three potential locations on different equipment
   - Each with a 33% probability
   - Each with calculated distance from terminal

## RabbitMQ Configuration

- **Host**: localhost
- **Port**: 5672 (RabbitMQ AMQP port)
- **User**: guest
- **Password**: guest
- **Exchange**: fault (headers type)
- **Queue**: fault-storer (durable)

## Message Headers

Each message includes headers:
- **id**: Fault UUID
- **producerId**: Producer identifier (e.g., "FL" for Fault Locator)
- **type**: Message type ("Fault" or "LineFault")

## Notes

- The example uses a headers exchange for flexible routing
- Messages are persistent (delivery mode: persistent)
- The consumer runs in a goroutine to allow concurrent publishing
- All three fault locations share the same fault ID but different equipment
