# Copyright 2024 Zaphiro Technologies
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""
Example demonstrating how to produce and consume measurement data using
Protocol Buffers and RabbitMQ Streams.
"""

import asyncio
import math
import os
import random
import struct
import sys
import time
import uuid
from typing import Dict

# Add the python directory to the path for imports
sys.path.insert(0, os.path.join(os.path.dirname(__file__), "../../../python"))

from rstream import (
    AMQPMessage,
    Consumer,
    ConsumerOffsetSpecification,
    MessageContext,
    OffsetType,
    Producer,
    StreamDoesNotExist,
    amqp_decoder,
)

# Import the generated protobuf messages
from zaphiro.grid.v1 import data_pb2


def generate_data_id(pmu_id: int, n_measures: int) -> Dict[str, data_pb2.Data]:
    """Generate synthetic measurement data for a PMU."""
    data_id = {}
    for k in range(n_measures):
        measurement_id = f"Dev{pmu_id:04d}-{k:04d}"
        
        # Determine data type based on measurement index
        if k < 1:
            data_type_int = 13
        elif k < 2:
            data_type_int = 28
        elif k < 4:
            data_type_int = 21
        elif k < 9:
            data_type_int = 20
        elif k < 11:
            data_type_int = 44
        elif k < 15:
            data_type_int = 30
        elif k < 18:
            data_type_int = 31
        elif k % 2 == 0:
            data_type_int = 43
        else:
            data_type_int = 44
        
        # Generate random value and convert to uint64 representation
        value = struct.unpack('Q', struct.pack('d', random.random()))[0]
        
        # Round timestamp to 20ms intervals
        timestamp_ms = math.floor(time.time() * 1000)
        timestamp_ms = round(timestamp_ms / 20) * 20
        
        data_id[measurement_id] = data_pb2.Data(
            dataType=data_type_int,
            value=value,
            measuredAt=timestamp_ms
        )
    
    return data_id


def generate_pmu_data_id(n_pmu: int, n_measures: int) -> list:
    """Generate measurement data for multiple PMUs."""
    pmu_data = []
    for k in range(n_pmu):
        producer_id = f"Dev{k:04d}"
        data = generate_data_id(k, n_measures)
        dataset = data_pb2.DataSet(producerId=producer_id, data=data)
        pmu_data.append(dataset)
    return pmu_data


async def publish_measurements():
    """Publish measurement messages to RabbitMQ stream."""
    print("Getting started with Streaming client for RabbitMQ")
    print("Connecting to RabbitMQ streaming ...")
    
    stream_name = "measurements"
    
    # Create producer
    async with Producer(
        host="localhost",
        port=5552,
        username="guest",
        password="guest"
    ) as producer:
        
        # Create stream (similar to Go example)
        try:
            await producer.create_stream(
                stream_name,
                arguments={"max-length-bytes": 2 * 1024 * 1024 * 1024}  # 2GB
            )
        except Exception as e:
            # Stream might already exist, which is fine
            pass
        
        # Send 10 messages
        for i in range(10):
            data = generate_pmu_data_id(1, 10)[0]
            
            # Serialize to protobuf
            buf = data.SerializeToString()
            
            # Create message with application properties
            timestamp_ms = round(time.time() * 1000 / 20) * 20
            message = AMQPMessage(
                body=buf,
                application_properties={
                    "Id": str(uuid.uuid4()),
                    "type": "DataSet",
                    "timestampId": timestamp_ms,
                    "producerId": data.producerId,
                    "aligned": False,
                    "samplingPeriod": "second"
                }
            )
            
            # Send message
            await producer.send(stream=stream_name, message=message)
        
        print("Sent 10 messages")
        
        # Allow time for messages to be confirmed
        await asyncio.sleep(1)


async def consume_measurements():
    """Consume measurement messages from RabbitMQ stream."""
    stream_name = "measurements"
    
    async def on_message(msg: AMQPMessage, message_context: MessageContext):
        """Handle received messages."""
        # Decode and parse the protobuf message
        dataset = data_pb2.DataSet()
        dataset.ParseFromString(msg.body)
        
        # Print measurement details
        for measurement_id, measurement in dataset.data.items():
            print(
                f"consumer name: my_consumer, measurement_id: {measurement_id}, "
                f"measurement_time {measurement.measuredAt}, "
                f"measurement_type {measurement.dataType}, "
                f"measurement_value {measurement.value}"
            )
    
    # Create consumer
    consumer = Consumer(
        host="localhost",
        port=5552,
        username="guest",
        password="guest",
        vhost="/"
    )
    
    try:
        await consumer.start()
        await consumer.subscribe(
            stream=stream_name,
            callback=on_message,
            decoder=amqp_decoder,
            offset_specification=ConsumerOffsetSpecification(OffsetType.FIRST, None)
        )
        
        print("Press Ctrl+C to stop\n")
        
        # Keep consuming until interrupted
        try:
            await consumer.run()
        except KeyboardInterrupt:
            print("\nStopping consumer...")
    finally:
        await consumer.close()
        
        # Clean up - delete stream (in prod, don't do this!)
        async with Producer(
            host="localhost",
            port=5552,
            username="guest",
            password="guest"
        ) as producer:
            try:
                await producer.delete_stream(stream_name)
            except StreamDoesNotExist:
                pass


async def main():
    """Main entry point - publish then consume."""
    import asyncio
    
    # First publish messages
    await publish_measurements()
    
    # Then consume them
    await consume_measurements()


if __name__ == "__main__":
    import asyncio
    
    try:
        asyncio.run(main())
    except KeyboardInterrupt:
        print("\nExiting...")
        sys.exit(0)
