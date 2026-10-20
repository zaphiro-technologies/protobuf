# Go Measurements Example

This example demonstrates how to produce and consume measurement data using
Protocol Buffers and RabbitMQ Streams.

## Overview

This example shows how to:

- Connect to RabbitMQ Streams
- Create a stream for measurements
- Produce `DataSet` messages containing measurement data
- Consume and decode measurement messages
- Handle consumer lifecycle

## Requirements

- Go 1.21 or higher
- Docker (for running RabbitMQ)
- RabbitMQ with streams plugin enabled

## Running the Example

1. Start the RabbitMQ infrastructure:

   ```bash
   make docker-start
   ```

2. Run the example:

   ```bash
   make example-measurements-go
   ```

3. The example will:
   - Connect to RabbitMQ Streams on port 5552
   - Create a stream called "measurements"
   - Send 10 measurement messages
   - Consume and display the measurements
   - Wait for user input to stop

4. Press any key to stop the example

## Expected Output

```
Getting started with Streaming client for RabbitMQ
Connecting to RabbitMQ streaming ...
Sent 10 messages
Press any key to stop

consumer name: my_consumer, measurement_id: Dev0000-0005, measurement_time 1720698360980, measurement_type 20, measurement_value 4592455024224327647
consumer name: my_consumer, measurement_id: Dev0000-0006, measurement_time 1720698360980, measurement_type 20, measurement_value 4604241342922663796
...
```

## Code Structure

### Producer

The producer:

1. Connects to RabbitMQ Streams using the `rabbitmq-stream-go-client`
2. Declares a stream with a 2GB max length
3. Generates synthetic measurement data using `generatePmuDataID()`
4. Marshals `DataSet` protobuf messages
5. Adds application properties (ID, type, timestamp, producer ID, etc.)
6. Sends messages to the stream

### Consumer

The consumer:

1. Creates a consumer for the stream
2. Starts consuming from the beginning (offset: first)
3. Unmarshals received protobuf messages
4. Prints measurement details (ID, timestamp, type, value)
5. Handles close events

## Key Protocol Buffer Messages

- **DataSet**: A collection of measurement data from a single producer
  - `ProducerId`: Identifier of the data producer
  - `Data`: Map of measurement ID to `Data` objects

- **Data**: A single measurement
  - `DataType`: Type of measurement (voltage, current, frequency, etc.)
  - `Value`: The measurement value (stored as uint64)
  - `MeasuredAt`: Timestamp in milliseconds

## RabbitMQ Stream Properties

- **Host**: localhost
- **Port**: 5552 (RabbitMQ Stream port)
- **User**: guest
- **Password**: guest
- **Stream**: measurements
- **Max Length**: 2GB

## Notes

- The example deletes the stream when stopping (for demonstration purposes)
- In production, streams should not be deleted automatically
- Values are stored as uint64 representation of float64 (using
  `math.Float64bits`)
- Measurements are timestamped and rounded to 20ms intervals
