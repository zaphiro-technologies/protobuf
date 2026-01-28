# Python Measurements Example

This example demonstrates how to produce and consume measurement data using Protocol Buffers and RabbitMQ Streams in Python.

## Overview

This example shows how to:
- Connect to RabbitMQ Streams using the `rstream` library
- Create a stream for measurements
- Produce `DataSet` messages containing measurement data
- Consume and decode measurement messages using asyncio
- Handle the consumer lifecycle

## Requirements

- Python 3.8 or higher
- Docker (for running RabbitMQ)
- RabbitMQ with streams plugin enabled

## Installation

Install the required Python dependencies:

```bash
pip install rstream protobuf-zaphiro
```

Or if using poetry:

```bash
poetry add rstream protobuf-zaphiro
```

## Running the Example

1. Start the RabbitMQ infrastructure:
   ```bash
   make docker-start
   ```

2. Run the example:
   ```bash
   python examples/python/measurements/main.py
   ```
   
   Or using the Makefile:
   ```bash
   make example-measurements-python
   ```

3. The example will:
   - Connect to RabbitMQ Streams on port 5552
   - Create a stream called "measurements"
   - Send 10 measurement messages
   - Consume and display the measurements
   - Wait for Ctrl+C to stop

4. Press Ctrl+C to stop the example

## Expected Output

```
Getting started with Streaming client for RabbitMQ
Connecting to RabbitMQ streaming ...
Sent 10 messages
Press Ctrl+C to stop

consumer name: my_consumer, measurement_id: Dev0000-0005, measurement_time 1720698360980, measurement_type 20, measurement_value 4592455024224327647
consumer name: my_consumer, measurement_id: Dev0000-0006, measurement_time 1720698360980, measurement_type 20, measurement_value 4604241342922663796
...
```

## Code Structure

### Producer

The `publish_measurements()` function:
1. Connects to RabbitMQ Streams using the `rstream.Producer`
2. Creates a stream with a 2GB max length
3. Generates synthetic measurement data using `generate_pmu_data_id()`
4. Serializes `DataSet` protobuf messages
5. Adds AMQP application properties (ID, type, timestamp, producer ID, etc.)
6. Sends messages asynchronously to the stream

### Consumer

The `consume_measurements()` function:
1. Creates a consumer for the stream using `rstream.Consumer`
2. Subscribes to consume from the beginning (offset: first)
3. Parses received protobuf messages
4. Prints measurement details (ID, timestamp, type, value)
5. Handles graceful shutdown with Ctrl+C

### Data Generation

The `generate_data_id()` function creates synthetic measurement data:
- Generates measurement IDs in the format `DevXXXX-YYYY`
- Assigns data types based on measurement index
- Creates random float values converted to uint64 representation
- Timestamps measurements rounded to 20ms intervals

## Key Protocol Buffer Messages

- **DataSet**: A collection of measurement data from a single producer
  - `producer_id`: Identifier of the data producer
  - `data`: Dictionary mapping measurement ID to `Data` objects

- **Data**: A single measurement
  - `data_type`: Type of measurement (voltage, current, frequency, etc.)
  - `value`: The measurement value (stored as uint64)
  - `measured_at`: Timestamp in milliseconds

## RabbitMQ Stream Configuration

- **Host**: localhost
- **Port**: 5552 (RabbitMQ Stream port)
- **Username**: guest
- **Password**: guest
- **Stream**: measurements
- **Max Length**: 2GB

## Dependencies

- **rstream**: Python client for RabbitMQ Streams
- **protobuf**: Google's Protocol Buffer library
- **protobuf-zaphiro**: Zaphiro's generated protobuf messages

## Notes

- The example uses Python's `asyncio` for asynchronous operations
- The stream is deleted when stopping (for demonstration purposes only)
- In production, streams should not be deleted automatically
- Values are stored as uint64 representation of float64
- Measurements are timestamped and rounded to 20ms intervals
- The example first publishes all messages, then starts consuming
