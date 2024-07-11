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

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

func generateTask(taskType TaskType, createdAt int64) *Task {
	return &Task{
		TaskType:  taskType,
		CreatedAt: createdAt,
	}
}

func TestTask(t *testing.T) {
	for k := 0; k < 5; k++ {
		test := generateTask(TaskType(k), time.Now().UnixNano())
		buf, err := proto.Marshal(test)
		assert.NoError(t, err)
		data := &Task{}
		err = proto.Unmarshal(buf, data)
		assert.NoError(t, err)
		assert.Equal(t, data.CreatedAt, test.CreatedAt)
		assert.Equal(t, data.TaskType, test.TaskType)
		assert.Equal(t, data.TaskType.Number(), protoreflect.EnumNumber(k))
	}
}

func generateNotification(
	notificationType NotificationType,
	createdAt int64,
	message string,
) *Notification {
	return &Notification{
		NotificationType: notificationType,
		CreatedAt:        createdAt,
		Message:          message,
	}
}

func TestNotification(t *testing.T) {
	for k := 0; k < 5; k++ {
		test := generateNotification(
			NotificationType(k),
			time.Now().UnixNano(),
			"myMessage",
		)
		buf, err := proto.Marshal(test)
		assert.NoError(t, err)
		data := &Notification{}
		err = proto.Unmarshal(buf, data)
		assert.NoError(t, err)
		assert.Equal(t, data.CreatedAt, test.CreatedAt)
		assert.Equal(t, data.NotificationType, test.NotificationType)
		assert.Equal(t, data.Message, "myMessage")
		assert.Equal(t, data.NotificationType.Number(), protoreflect.EnumNumber(k))
	}
}

func TestTriggerNotification(t *testing.T) {
	test := generateNotification(
		NotificationType_NOTIFICATION_TYPE_TRIGGER,
		time.Now().UnixNano(),
		"1", // alternatively a string code
	)
	buf, err := proto.Marshal(test)
	assert.NoError(t, err)
	data := &Notification{}
	err = proto.Unmarshal(buf, data)
	assert.NoError(t, err)
	assert.Equal(t, data.CreatedAt, test.CreatedAt)
	assert.Equal(t, data.NotificationType, test.NotificationType)
	assert.Equal(t, data.Message, "1")
	assert.Equal(t, data.NotificationType, NotificationType_NOTIFICATION_TYPE_TRIGGER)
}

func BenchmarkNotificationSerialization(b *testing.B) {
	test := generateNotification(
		NotificationType(rand.Intn(4)),
		time.Now().UnixNano(),
		"1",
	)
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &Notification{}
		_ = proto.Unmarshal(buf, conf)
	}
}

func BenchmarkTaskSerialization(b *testing.B) {
	test := generateTask(TaskType(rand.Intn(5)), time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &Task{}
		_ = proto.Unmarshal(buf, conf)
	}
}
