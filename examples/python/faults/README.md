# Python Faults Example

This example demonstrates how to produce and consume fault event messages using Protocol Buffers and RabbitMQ AMQP exchanges in Python.

## Overview

This example shows how to:
- Connect to RabbitMQ using AMQP 0.9.1 via the `pika` library
- Create an exchange for fault events
- Produce fault event messages (Fault and LineFault)
- Consume and decode fault messages based on message headers
- Handle different fault event types (STARTED, ENDED, LOCATED)

## Requirements

- Python 3.8 or higher
- Docker (for running RabbitMQ)
- RabbitMQ with AMQP support

## Installation

Install the required Python dependencies:

```bash
pip install pika protobuf-zaphiro
```

Or if using poetry:

```bash
poetry add pika protobuf-zaphiro
```

## Running the Example

1. Start the RabbitMQ infrastructure:
   ```bash
   make docker-start
   ```

2. Run the example:
   ```bash
   python examples/python/faults/main.py
   ```
   
   Or using the Makefile:
   ```bash
   make example-faults-python
   ```

3. The example will:
   - Connect to RabbitMQ on port 5672
   - Create a headers exchange called "fault"
   - Create a queue called "fault-storer"
   - Send fault event messages
   - Consume and display fault events
   - Wait for Ctrl+C to stop

4. Press Ctrl+C to stop the example

## Expected Output

```
Getting started with AMQP client for RabbitMQ
Connecting to RabbitMQ ...
Press Ctrl+C to stop

Received a fault event message: 1948cd36-f835-4dc5-83c2-ba5a7612362a, event type: FAULT_EVENT_TYPE_STARTED
Received a fault event message: 1948cd36-f835-4dc5-83c2-ba5a7612362a, event type: FAULT_EVENT_TYPE_ENDED
Received a line fault event message: 1948cd36-f835-4dc5-83c2-ba5a7612362a, event type: FAULT_EVENT_TYPE_LOCATED, faulty line: EQ-1, probability: 0.330000, length from t1: 0.526424
Received a line fault event message: 1948cd36-f835-4dc5-83c2-ba5a7612362a, event type: FAULT_EVENT_TYPE_LOCATED, faulty line: EQ-2, probability: 0.330000, length from t1: 0.484818
Received a line fault event message: 1948cd36-f835-4dc5-83c2-ba5a7612362a, event type: FAULT_EVENT_TYPE_LOCATED, faulty line: EQ-3, probability: 0.330000, length from t1: 0.315550
```

## Code Structure

### Producer

The producer (in the `main()` function):
1. Connects to RabbitMQ using AMQP 0.9.1 via `pika.BlockingConnection`
2. Declares a headers exchange named "fault"
3. Generates fault events with different types:
   - **FAULT_EVENT_TYPE_STARTED**: Initial fault detection
   - **FAULT_EVENT_TYPE_ENDED**: Fault cleared
   - **FAULT_EVENT_TYPE_LOCATED**: Fault location identified (with probability)
4. Serializes protobuf messages and publishes with headers
5. Demonstrates both generic `Fault` and specific `LineFault` messages

### Consumer

The consumer (callback function):
1. Creates a durable queue "fault-storer"
2. Binds the queue to the "fault" exchange
3. Consumes messages with auto-acknowledgment
4. Parses messages based on the "type" header:
   - **"Fault"**: Generic fault event
   - **"LineFault"**: Line-specific fault with location details
5. Prints relevant fault information using protobuf enum names

## Key Protocol Buffer Messages

### Fault

Represents a generic fault event with:
- `id`: Unique fault identifier (UUID)
- `description`: Human-readable description
- `kind`: Fault kind (e.g., LINE_TO_GROUND)
- `phases`: Affected phases (e.g., ABC)
- `fault_event_type`: Event type (STARTED, ENDED, LOCATED)
- `fault_current`: Magnitude of fault current
- `location_probability`: Probability of fault location (for LOCATED events)
- `impacted_equipment_ids`: List of affected equipment
- `used_measurement_ids`: Measurements used for fault detection
- `measurement_timestamp`: When the fault was measured
- `updated_at`: Last update timestamp

### LineFault

Extends `Fault` for line-specific faults with:
- `fault`: Embedded fault information
- `length_from_terminal1`: Distance from terminal 1 (normalized, 0-1)
- `length_uncertainty`: Uncertainty in meters

## Fault Event Lifecycle

The example demonstrates a typical fault event sequence:

1. **STARTED**: Fault is detected in the system
2. **ENDED**: Fault condition has cleared
3. **LOCATED** (multiple): Possible fault locations with probabilities
   - Three potential locations on different equipment (EQ-1, EQ-2, EQ-3)
   - Each with a 33% probability
   - Each with a randomly calculated distance from terminal

## RabbitMQ Configuration

- **Host**: localhost
- **Port**: 5672 (RabbitMQ AMQP port)
- **Username**: guest
- **Password**: guest
- **Exchange**: fault (headers type, durable)
- **Queue**: fault-storer (durable)

## Message Headers

Each message includes headers:
- **id**: Fault UUID
- **producerId**: Producer identifier (e.g., "FL" for Fault Locator)
- **type**: Message type ("Fault" or "LineFault")

## Dependencies

- **pika**: Python AMQP client library for RabbitMQ
- **protobuf**: Google's Protocol Buffer library
- **protobuf-zaphiro**: Zaphiro's generated protobuf messages

## Notes

- The example uses blocking I/O for simplicity
- Messages are persistent (delivery_mode=2)
- The consumer and producer run in the same process
- Messages are published before the consumer starts, so they're queued
- All three fault locations share the same fault ID but represent different equipment
- The example uses `FaultEventType.Name()` to convert enum values to readable strings
