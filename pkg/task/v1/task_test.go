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

func generateTask(taskId string, taskType TaskType, createdAt int64) *Task {
	return &Task{
		Id:        taskId,
		TaskType:  taskType,
		CreatedAt: createdAt,
	}
}

func TestTask(t *testing.T) {
	for k := 0; k < 5; k++ {
		test := generateTask(uuid.NewString(), TaskType(k), time.Now().UnixNano())
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
	notificationType NotificationType,
	createdAt int64,
	message string,
) *Notification {
	return &Notification{
		Id:               notificationId,
		NotificationType: notificationType,
		CreatedAt:        createdAt,
		Message:          message,
	}
}

func TestNotification(t *testing.T) {
	for k := 0; k < 5; k++ {
		test := generateNotification(
			uuid.NewString(),
			NotificationType(k),
			time.Now().UnixNano(),
			"myMessage",
		)
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

func TestTriggerNotification(t *testing.T) {
	test := generateNotification(
		uuid.NewString(),
		NotificationType_NOTIFICATION_TYPE_TRIGGER,
		time.Now().UnixNano(),
		"1", //alternatively a string code
	)
	test.TimestampID = func() *int64 { i := int64(4); return &i }()
	test.ProducerID = func() *string { i := "CIM_PMU_CODE"; return &i }()
	buf, err := proto.Marshal(test)
	assert.NoError(t, err)
	data := &Notification{}
	err = proto.Unmarshal(buf, data)
	assert.NoError(t, err)
	assert.Equal(t, data.Id, test.Id)
	assert.Equal(t, data.CreatedAt, test.CreatedAt)
	assert.Equal(t, data.NotificationType, test.NotificationType)
	assert.Equal(t, data.Message, "1")
	assert.Equal(t, data.NotificationType, NotificationType_NOTIFICATION_TYPE_TRIGGER)
	assert.Equal(t, *data.TimestampID, int64(4))
	assert.Equal(t, *data.ProducerID, "CIM_PMU_CODE")
}

func BenchmarkNotificationSerialization(b *testing.B) {
	test := generateNotification(
		uuid.NewString(),
		NotificationType(rand.Intn(4)),
		time.Now().UnixNano(),
		"myMessage",
	)
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &Notification{}
		_ = proto.Unmarshal(buf, conf)
	}
}

func BenchmarkTaskSerialization(b *testing.B) {
	test := generateTask(uuid.NewString(), TaskType(rand.Intn(5)), time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &Task{}
		_ = proto.Unmarshal(buf, conf)
	}
}
