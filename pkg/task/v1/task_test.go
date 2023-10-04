package v1

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

func generateTask(taskId string, taskType int, createdAt int64) *Task {
	return &Task{
		Id:        taskId,
		TaskType:  TaskType(taskType),
		CreatedAt: createdAt,
	}
}

func TestTask(t *testing.T) {
	for k := 0; k < 5; k++ {
		test := generateTask(uuid.NewString(), k, time.Now().UnixNano())
		buf, err := proto.Marshal(test)
		assert.NoError(t, err)
		data := &Task{}
		err = proto.Unmarshal(buf, data)
		assert.NoError(t, err)
		assert.Equal(t, data.Id, test.Id)
		assert.Equal(t, data.CreatedAt, test.CreatedAt)
		assert.Equal(t, data.TaskType, test.TaskType)
		assert.Equal(t, data.TaskType.Number(), protoreflect.EnumNumber(k))
	}
}

func generateNotification(
	notificationId string,
	notificationType int,
	createdAt int64,
	message string,
) *Notification {
	return &Notification{
		Id:               notificationId,
		NotificationType: NotificationType(notificationType),
		CreatedAt:        createdAt,
		Message:          message,
	}
}

func TestNotification(t *testing.T) {
	for k := 0; k < 4; k++ {
		test := generateNotification(uuid.NewString(), k, time.Now().UnixNano(), "myMessage")
		buf, err := proto.Marshal(test)
		assert.NoError(t, err)
		data := &Notification{}
		err = proto.Unmarshal(buf, data)
		assert.NoError(t, err)
		assert.Equal(t, data.Id, test.Id)
		assert.Equal(t, data.CreatedAt, test.CreatedAt)
		assert.Equal(t, data.NotificationType, test.NotificationType)
		assert.Equal(t, data.Message, "myMessage")
		assert.Equal(t, data.NotificationType.Number(), protoreflect.EnumNumber(k))
	}
}

func BenchmarkNotificationSerialization(b *testing.B) {
	test := generateNotification(uuid.NewString(), rand.Intn(4), time.Now().UnixNano(), "myMessage")
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &Notification{}
		_ = proto.Unmarshal(buf, conf)
	}
}

func BenchmarkTaskSerialization(b *testing.B) {
	test := generateTask(uuid.NewString(), rand.Intn(5), time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &Task{}
		_ = proto.Unmarshal(buf, conf)
	}
}
