#!/bin/bash

# config
CONFIG_STREAM="c37-118-configuration"
MEASUREMENT_STREAM="measurement"
NOTIFICATION_EXCHANGE_TYPE="headers"
NOTIFICATION_EXCHANGE="notification"
DEAD_LETTER_EXCHANGE="notification-unprocessed"
NOTIFICATION_QUEUE="notification-data-labview"
TRIGGER_QUEUE="notification-trigger-labview"
TIMEOUT=$((10*60*1000))
NOTIFICATION_QUEUE_ARGUMENTS="{\"exclusive\":false,\"x-message-ttl\":$TIMEOUT, \"x-dead-letter-exchange\":\"$DEAD_LETTER_EXCHANGE\"}"
BYTES=$((100 * 1000000))
NOTIFICATION_QUEUE_BINDING_1_ARGUMENTS='{"x-match":"all","type":"Notification","notificationType":1}'
NOTIFICATION_QUEUE_BINDING_2_ARGUMENTS='{"x-match":"all","type":"Notification","notificationType":2}'
NOTIFICATION_QUEUE_BINDING_3_ARGUMENTS='{"x-match":"all","type":"Notification","notificationType":3}'
TRIGGER_QUEUE_BINDING_1_ARGUMENTS='{"x-match":"all","type":"Notification","notificationType":4}'
ASYNC_QUEUE_BINDING_1_ARGUMENTS='{"x-match":"all","type":"Notification","notificationType":0}'
RABBITMQ_FAULT_EXCHANGE_NAME=fault
RABBITMQ_FAULT_CONSUMER_TAG=fault-storer
RABBITMQ_ESTIMATION_EXCHANGE_NAME=estimated-measurements
RABBITMQ_ESTIMATION_CONSUMER_TAG=estimated-measurements
RABBITMQ_TOPOLOGY_EXCHANGE_NAME=topology
RABBITMQ_TOPOLOGY_CONSUMER_TAG=topology
RABBITMQ_EVENT_EXCHANGE_NAME=event
RABBITMQ_EVENT_CONSUMER_TAG=event-storer
RABBITMQ_STAT_STREAM_NAME=stat
RABBITMQ_STAT_LABVIEW_EXCHANGE_NAME=stat-labview
RABBITMQ_STAT_LABVIEW_QUEUE_NAME=stat-labview
RABBITMQ_ASYNC_QUEUE_NAME=async-measurements

# Declare the streams and queues
rabbitmqadmin declare queue name="$CONFIG_STREAM" queue_type=stream arguments="{\"x-max-length-bytes\":$BYTES}"
rabbitmqadmin declare queue name="$MEASUREMENT_STREAM" queue_type=stream arguments="{\"x-max-length-bytes\":$BYTES}"
rabbitmqadmin declare queue name="$RABBITMQ_STAT_STREAM_NAME" queue_type=stream arguments="{\"x-max-length-bytes\":$BYTES}"
rabbitmqadmin declare exchange name="$NOTIFICATION_EXCHANGE" type="$NOTIFICATION_EXCHANGE_TYPE"
rabbitmqadmin declare exchange name="$DEAD_LETTER_EXCHANGE" type="fanout"
rabbitmqctl set_policy DLX ".*" "{\"dead-letter-exchange\":\"$DEAD_LETTER_EXCHANGE\"}" --apply-to queues
rabbitmqadmin declare queue name="$NOTIFICATION_QUEUE" durable=true
rabbitmqadmin declare queue name="$TRIGGER_QUEUE" durable=true
rabbitmqadmin declare queue name="$RABBITMQ_ASYNC_QUEUE_NAME" durable=true
rabbitmqadmin declare binding source="$NOTIFICATION_EXCHANGE" destination_type=queue destination="$NOTIFICATION_QUEUE" arguments="$NOTIFICATION_QUEUE_BINDING_1_ARGUMENTS"
rabbitmqadmin declare binding source="$NOTIFICATION_EXCHANGE" destination_type=queue destination="$NOTIFICATION_QUEUE" arguments="$NOTIFICATION_QUEUE_BINDING_2_ARGUMENTS"
rabbitmqadmin declare binding source="$NOTIFICATION_EXCHANGE" destination_type=queue destination="$NOTIFICATION_QUEUE" arguments="$NOTIFICATION_QUEUE_BINDING_3_ARGUMENTS"
rabbitmqadmin declare binding source="$NOTIFICATION_EXCHANGE" destination_type=queue destination="$TRIGGER_QUEUE" arguments="$TRIGGER_QUEUE_BINDING_1_ARGUMENTS"
rabbitmqadmin declare binding source="$NOTIFICATION_EXCHANGE" destination_type=queue destination="$RABBITMQ_ASYNC_QUEUE_NAME" arguments="$ASYNC_QUEUE_BINDING_1_ARGUMENTS"
rabbitmqadmin declare exchange name="$RABBITMQ_FAULT_EXCHANGE_NAME" type=headers durable=true auto_delete=false internal=false
rabbitmqadmin declare queue name="$RABBITMQ_FAULT_CONSUMER_TAG" durable=true auto_delete=false
rabbitmqadmin declare binding source="$RABBITMQ_FAULT_EXCHANGE_NAME" destination="$RABBITMQ_FAULT_CONSUMER_TAG" destination_type=queue routing_key=""
rabbitmqadmin declare exchange name="$RABBITMQ_ESTIMATION_EXCHANGE_NAME" type=headers durable=true auto_delete=false internal=false
rabbitmqadmin declare queue name="$RABBITMQ_ESTIMATION_CONSUMER_TAG" durable=true auto_delete=false
rabbitmqadmin declare binding source="$RABBITMQ_ESTIMATION_EXCHANGE_NAME" destination="$RABBITMQ_ESTIMATION_CONSUMER_TAG" destination_type=queue routing_key=""
rabbitmqadmin declare exchange name="$RABBITMQ_TOPOLOGY_EXCHANGE_NAME" type=headers durable=true auto_delete=false internal=false
rabbitmqadmin declare queue name="$RABBITMQ_TOPOLOGY_CONSUMER_TAG" durable=true auto_delete=false
rabbitmqadmin declare binding source="$RABBITMQ_TOPOLOGY_EXCHANGE_NAME" destination="$RABBITMQ_TOPOLOGY_CONSUMER_TAG" destination_type=queue routing_key=""
rabbitmqadmin declare exchange name="$RABBITMQ_EVENT_EXCHANGE_NAME" type=headers durable=true auto_delete=false internal=false
rabbitmqadmin declare queue name="$RABBITMQ_EVENT_CONSUMER_TAG" durable=true auto_delete=false
rabbitmqadmin declare binding source="$RABBITMQ_EVENT_EXCHANGE_NAME" destination="$RABBITMQ_EVENT_CONSUMER_TAG" destination_type=queue routing_key=""
rabbitmqadmin declare exchange name="$RABBITMQ_STAT_LABVIEW_EXCHANGE_NAME" type=headers durable=true auto_delete=false internal=false
rabbitmqadmin declare queue name="$RABBITMQ_STAT_LABVIEW_QUEUE_NAME" durable=true auto_delete=false
rabbitmqadmin declare binding source="$RABBITMQ_STAT_LABVIEW_EXCHANGE_NAME" destination="$RABBITMQ_STAT_LABVIEW_QUEUE_NAME" destination_type=queue routing_key=""

