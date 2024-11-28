// Copyright 2024 Zaphiro Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"math/rand"
	"testing"
	"time"

	"github.com/ccoveille/go-safecast"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

func generateEvent(
	eventId string,
	eventSourceType int32,
	eventSource string,
	occurredAt int64,
	message string,
) *Event {
	detectedAt := occurredAt + rand.Int63n(40)
	return &Event{
		Id:         eventId,
		SourceId:   eventSource,
		SourceType: EventSourceType(eventSourceType),
		OccurredAt: occurredAt,
		DetectedAt: &detectedAt,
		Message:    message,
	}
}

func TestEvent(t *testing.T) {
	for k := int32(0); k < 4; k++ {
		test := generateEvent(
			uuid.NewString(),
			k,
			uuid.NewString(),
			time.Now().UnixNano(),
			"my message test event",
		)
		buf, err := proto.Marshal(test)
		assert.NoError(t, err)
		data := &Event{}
		err = proto.Unmarshal(buf, data)
		assert.NoError(t, err)
		assert.Equal(t, data.OccurredAt, test.OccurredAt)
		assert.GreaterOrEqual(t, *data.DetectedAt, test.OccurredAt)
		assert.Equal(t, data.Message, test.Message)
		assert.Equal(t, data.SourceType, test.SourceType)
		assert.Equal(t, data.SourceType.Number(), protoreflect.EnumNumber(k))
	}
}

func BenchmarkEventSerialization(b *testing.B) {
	randEventSourceType, err := safecast.ToInt32(rand.Intn(4))
	if err != nil {
		b.Fatalf("Failed to convert int to int32: %v", err)
	}
	test := generateEvent(
		uuid.NewString(),
		randEventSourceType,
		uuid.NewString(),
		time.Now().UnixNano(),
		"my message benchmark event",
	)
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &Event{}
		_ = proto.Unmarshal(buf, conf)
	}
}

func generateGridEvent(
	event *Event,
	componentId string,
	nominalValue float64,
	referenceLimit float64,
	probability float64,
	substationId string,
) *GridEvent {
	return &GridEvent{
		Event:          event,
		ComponentID:    componentId,
		SubstationID:   &substationId,
		ReferenceLimit: referenceLimit,
		Probability:    &probability,
		Value:          nominalValue,
	}
}

func TestGridEvent(t *testing.T) {
	for k := int32(0); k < 5; k++ {
		event := generateEvent(
			uuid.NewString(),
			k,
			uuid.NewString(),
			time.Now().UnixNano(),
			"my message test grid event",
		)
		test := generateGridEvent(
			event,
			uuid.NewString(),
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
			uuid.NewString(),
		)
		buf, err := proto.Marshal(test)
		assert.NoError(t, err)
		data := &GridEvent{}
		err = proto.Unmarshal(buf, data)
		assert.NoError(t, err)
		assert.Equal(t, data.Event.OccurredAt, test.Event.OccurredAt)
		assert.Equal(t, data.Event.Message, test.Event.Message)
		assert.Equal(t, data.Value, test.Value)
	}
}

func BenchmarkGridEventSerialization(b *testing.B) {
	randEventSourceType, err := safecast.ToInt32(rand.Intn(4))
	if err != nil {
		b.Fatalf("Failed to convert int to int32: %v", err)
	}
	event := generateEvent(
		uuid.NewString(),
		randEventSourceType,
		uuid.NewString(),
		time.Now().UnixNano(),
		"my message benchmark grid event",
	)
	test := generateGridEvent(
		event,
		uuid.NewString(),
		rand.Float64(),
		rand.Float64(),
		rand.Float64(),
		uuid.NewString(),
	)
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &GridEvent{}
		_ = proto.Unmarshal(buf, conf)
	}
}
