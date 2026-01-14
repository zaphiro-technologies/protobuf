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

func generateFault(
	faultID string,
	faultKind, phaseCode int32,
	updatedAt int64,
	faultyEquipmentID string,
) *Fault {
	return &Fault{
		Id:                faultID,
		Kind:              PhaseConnectedFaultKind(faultKind),
		UpdatedAt:         updatedAt,
		Phases:            PhaseCode(phaseCode),
		FaultyEquipmentId: &faultyEquipmentID,
	}
}

func TestFault(t *testing.T) {
	for k := int32(0); k < 5; k++ {
		randPhaseCode, err := safecast.Convert[int32](rand.Intn(26))
		if err != nil {
			t.Fatalf("Failed to convert int to int32: %v", err)
		}
		test := generateFault(
			uuid.NewString(),
			k,
			randPhaseCode,
			time.Now().UnixNano(),
			uuid.NewString(),
		)
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
	randFaultKind, err := safecast.Convert[int32](rand.Intn(4))
	if err != nil {
		b.Fatalf("Failed to convert int to int32: %v", err)
	}
	randPhaseCode, err := safecast.Convert[int32](rand.Intn(26))
	if err != nil {
		b.Fatalf("Failed to convert int to int32: %v", err)
	}
	test := generateFault(
		uuid.NewString(),
		randFaultKind,
		randPhaseCode,
		time.Now().UnixNano(),
		uuid.NewString(),
	)
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
	for k := int32(0); k < 5; k++ {
		randPhaseCode, err := safecast.Convert[int32](rand.Intn(26))
		if err != nil {
			t.Fatalf("Failed to convert int to int32: %v", err)
		}
		fault := generateFault(uuid.NewString(), k, randPhaseCode, time.Now().UnixNano(), "line1")
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
	randFaultKind, err := safecast.Convert[int32](rand.Intn(4))
	if err != nil {
		b.Fatalf("Failed to convert int to int32: %v", err)
	}
	randPhaseCode, err := safecast.Convert[int32](rand.Intn(26))
	if err != nil {
		b.Fatalf("Failed to convert int to int32: %v", err)
	}
	fault := generateFault(
		uuid.NewString(),
		randFaultKind,
		randPhaseCode,
		time.Now().UnixNano(),
		"line1",
	)
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
	for k := int32(0); k < 5; k++ {
		randPhaseCode, err := safecast.Convert[int32](rand.Intn(26))
		if err != nil {
			t.Fatalf("Failed to convert int to int32: %v", err)
		}
		fault := generateFault(
			uuid.NewString(),
			k,
			randPhaseCode,
			time.Now().UnixNano(),
			"equipment1",
		)
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
	randFaultKind, err := safecast.Convert[int32](rand.Intn(4))
	if err != nil {
		b.Fatalf("Failed to convert int to int32: %v", err)
	}
	randPhaseCode, err := safecast.Convert[int32](rand.Intn(26))
	if err != nil {
		b.Fatalf("Failed to convert int to int32: %v", err)
	}
	fault := generateFault(
		uuid.NewString(),
		randFaultKind,
		randPhaseCode,
		time.Now().UnixNano(),
		"equipment1",
	)
	test := generateEquipmentFault(fault, uuid.NewString())
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &EquipmentFault{}
		_ = proto.Unmarshal(buf, conf)
	}
}
