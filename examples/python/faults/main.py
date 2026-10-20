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
Example demonstrating how to produce and consume fault event messages using
Protocol Buffers and RabbitMQ AMQP exchanges.
"""

import math
import os
import random
import sys
import time
import uuid
from typing import List, Optional

# Add the python directory to the path for imports
sys.path.insert(0, os.path.join(os.path.dirname(__file__), "../../../python"))

import pika

# Import the generated protobuf messages
from zaphiro.grid.v1 import fault_pb2


def generate_fault(
    fault_id: str,
    event_type: int,
    measurement_timestamp: int,
    fault_current: float,
    location_probability: Optional[float],
    used_measurements: List[fault_pb2.FaultMeasurement],
    impacted_equipments: List[str],
    faulty_equipment: Optional[str]
) -> fault_pb2.Fault:
    """Generate a fault event."""
    description = "Example fault"
    timestamp_ms = round(time.time() * 1000 / 20) * 20
    
    fault = fault_pb2.Fault(
        description=description,
        Id=fault_id,
        kind=fault_pb2.PHASE_CONNECTED_FAULT_KIND_LINE_TO_GROUND,
        phases=fault_pb2.PHASE_CODE_ABC,
        updatedAt=timestamp_ms,
        faultEventType=event_type,
        measurementTimestamp=measurement_timestamp,
        faultCurrent=fault_current,
        impactedEquipmentIds=impacted_equipments
    )
    
    if faulty_equipment:
        fault.faultyEquipmentId = faulty_equipment
    
    if location_probability is not None:
        fault.locationProbability = location_probability
    
    fault.usedMeasurementIds.extend(used_measurements)
    
    return fault


def generate_line_fault(
    fault_id: str,
    event_type: int,
    measurement_timestamp: int,
    fault_current: float,
    location_probability: Optional[float],
    used_measurements: List[fault_pb2.FaultMeasurement],
    impacted_equipments: List[str],
    faulty_equipment: Optional[str]
) -> fault_pb2.LineFault:
    """Generate a line fault event."""
    fault = generate_fault(
        fault_id,
        event_type,
        measurement_timestamp,
        fault_current,
        location_probability,
        used_measurements,
        impacted_equipments,
        faulty_equipment
    )
    
    length_from_terminal1 = random.random()
    length_uncertainty = 100.0
    
    line_fault = fault_pb2.LineFault(
        fault=fault,
        lengthFromTerminal1=length_from_terminal1,
        lengthUncertainty=length_uncertainty
    )
    
    return line_fault


def publish_message(channel, headers: dict, body: bytes):
    """Publish a message to the fault exchange."""
    properties = pika.BasicProperties(
        headers=headers,
        delivery_mode=2  # Persistent
    )
    
    channel.basic_publish(
        exchange="fault",
        routing_key="",
        body=body,
        properties=properties
    )


def consume_callback(ch, method, properties, body):
    """Callback for consuming fault messages."""
    msg_type = properties.headers.get("type")
    
    if msg_type == "Fault":
        fault = fault_pb2.Fault()
        fault.ParseFromString(body)
        print(
            f"Received a fault event message: {fault.Id}, "
            f"event type: {fault_pb2.FaultEventType.Name(fault.faultEventType)}"
        )
    
    elif msg_type == "LineFault":
        line_fault = fault_pb2.LineFault()
        line_fault.ParseFromString(body)
        print(
            f"Received a line fault event message: {line_fault.fault.Id}, "
            f"event type: {fault_pb2.FaultEventType.Name(line_fault.fault.faultEventType)}, "
            f"faulty line: {line_fault.fault.faultyEquipmentId}, "
            f"probability: {line_fault.fault.locationProbability:.6f}, "
            f"length from t1: {line_fault.lengthFromTerminal1:.6f}"
        )


def main():
    """Main entry point."""
    print("Getting started with AMQP client for RabbitMQ")
    print("Connecting to RabbitMQ ...")
    
    # Connect to RabbitMQ
    connection = pika.BlockingConnection(
        pika.ConnectionParameters(
            host="localhost",
            port=5672,
            credentials=pika.PlainCredentials("guest", "guest")
        )
    )
    channel = connection.channel()
    
    # Declare exchange
    channel.exchange_declare(
        exchange="fault",
        exchange_type="headers",
        durable=True
    )
    
    # Declare queue
    queue = channel.queue_declare(
        queue="fault-storer",
        durable=True
    )
    
    # Bind queue to exchange (headers exchange, so we bind with empty arguments)
    channel.queue_bind(
        queue="fault-storer",
        exchange="fault"
    )
    
    # Set up consumer
    channel.basic_consume(
        queue="fault-storer",
        on_message_callback=consume_callback,
        auto_ack=True
    )
    
    # Prepare fault data
    fault_id = str(uuid.uuid4())
    timestamp_ms = round(time.time() * 1000 / 20) * 20
    fault_current = 100000.0
    impacted_equipments = ["EQ-1", "EQ-2", "EQ-3"]
    used_measurements = [
        fault_pb2.FaultMeasurement(positiveSign=True, measurementID="M-1"),
        fault_pb2.FaultMeasurement(positiveSign=True, measurementID="M-2"),
        fault_pb2.FaultMeasurement(positiveSign=True, measurementID="M-3"),
    ]
    
    # Send STARTED event
    start_fault_event = generate_fault(
        fault_id,
        fault_pb2.FAULT_EVENT_TYPE_STARTED,
        timestamp_ms,
        fault_current,
        None,
        used_measurements,
        impacted_equipments,
        None
    )
    
    headers = {
        "id": fault_id,
        "producerId": "FL",
        "type": "Fault"
    }
    publish_message(channel, headers, start_fault_event.SerializeToString())
    
    time.sleep(0.02)  # 20ms delay
    
    # Send ENDED event
    end_fault_event = generate_fault(
        fault_id,
        fault_pb2.FAULT_EVENT_TYPE_ENDED,
        timestamp_ms,
        fault_current,
        None,
        used_measurements,
        impacted_equipments,
        None
    )
    publish_message(channel, headers, end_fault_event.SerializeToString())
    
    # Send three LOCATED events for different equipment
    headers["type"] = "LineFault"
    probability = 0.33
    
    for equipment_id in ["EQ-1", "EQ-2", "EQ-3"]:
        location = generate_line_fault(
            fault_id,
            fault_pb2.FAULT_EVENT_TYPE_LOCATED,
            timestamp_ms,
            fault_current,
            probability,
            used_measurements,
            impacted_equipments,
            equipment_id
        )
        publish_message(channel, headers, location.SerializeToString())
    
    print("Press Ctrl+C to stop\n")
    
    # Start consuming
    try:
        channel.start_consuming()
    except KeyboardInterrupt:
        print("\nStopping consumer...")
        channel.stop_consuming()
    finally:
        connection.close()


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print("\nExiting...")
        sys.exit(0)
