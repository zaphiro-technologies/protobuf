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

func generateFault(faultId string, faultKind, phaseCode int, updatedAt int64, faultyEquipmentId string) *Fault {
	return &Fault{
		Id:                faultId,
		Kind:              PhaseConnectedFaultKind(faultKind),
		UpdatedAt:         updatedAt,
		Phases:            PhaseCode(phaseCode),
		FaultyEquipmentId: &faultyEquipmentId,
	}
}

func TestFault(t *testing.T) {
	for k := 0; k < 5; k++ {
		test := generateFault(uuid.NewString(), k, rand.Intn(26), time.Now().UnixNano(), uuid.NewString())
		buf, err := proto.Marshal(test)
		assert.NoError(t, err)
		data := &Fault{}
		err = proto.Unmarshal(buf, data)
		assert.NoError(t, err)
		assert.Equal(t, data.UpdatedAt, test.UpdatedAt)
		assert.Equal(t, data.Kind, test.Kind)
		assert.Equal(t, data.Kind.Number(), protoreflect.EnumNumber(k))
		assert.Equal(t, data.Phases, test.Phases)
	}
}

func BenchmarkFaultSerialization(b *testing.B) {
	test := generateFault(uuid.NewString(), rand.Intn(4), rand.Intn(26), time.Now().UnixNano(), uuid.NewString())
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &Fault{}
		_ = proto.Unmarshal(buf, conf)
	}
}

func generateLineFault(
	fault *Fault,
	lengthFromTerminal1 float32,
) *LineFault {
	return &LineFault{
		Fault:               fault,
		LengthFromTerminal1: &lengthFromTerminal1,
	}
}

func TestLineFault(t *testing.T) {
	for k := 0; k < 5; k++ {
		fault := generateFault(uuid.NewString(), k, rand.Intn(26), time.Now().UnixNano(), "line1")
		test := generateLineFault(fault, rand.Float32())
		buf, err := proto.Marshal(test)
		assert.NoError(t, err)
		data := &LineFault{}
		err = proto.Unmarshal(buf, data)
		assert.NoError(t, err)
		assert.Equal(t, data.LengthFromTerminal1, test.LengthFromTerminal1)
	}
}

func BenchmarkLineFaultSerialization(b *testing.B) {
	fault := generateFault(uuid.NewString(), rand.Intn(4), rand.Intn(26), time.Now().UnixNano(), "line1")
	test := generateLineFault(fault, rand.Float32())
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &LineFault{}
		_ = proto.Unmarshal(buf, conf)
	}
}

func generateEquipmentFault(
	fault *Fault,
	terminalID string,
) *EquipmentFault {
	return &EquipmentFault{
		Fault:      fault,
		TerminalID: &terminalID,
	}
}

func TestEquipmentFault(t *testing.T) {
	for k := 0; k < 5; k++ {
		fault := generateFault(uuid.NewString(), k, rand.Intn(26), time.Now().UnixNano(), "equipment1")
		test := generateEquipmentFault(fault, uuid.NewString())
		buf, err := proto.Marshal(test)
		assert.NoError(t, err)
		data := &EquipmentFault{}
		err = proto.Unmarshal(buf, data)
		assert.NoError(t, err)
		assert.Equal(t, data.TerminalID, test.TerminalID)
	}
}

func BenchmarkEquipmentFaultSerialization(b *testing.B) {
	fault := generateFault(uuid.NewString(), rand.Intn(4), rand.Intn(26), time.Now().UnixNano(), "equipment1")
	test := generateEquipmentFault(fault, uuid.NewString())
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &EquipmentFault{}
		_ = proto.Unmarshal(buf, conf)
	}
}
