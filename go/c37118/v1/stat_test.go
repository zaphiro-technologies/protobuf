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
)

func generateStat(timestamp int64) *Stat {
	return &Stat{
		Error:         rand.Uint32() % 8,
		Sync:          rand.Intn(2) == 1,
		Sorting:       rand.Intn(2) == 1,
		Trigger:       rand.Intn(2) == 1,
		ConfigChange:  rand.Intn(2) == 1,
		DataModified:  rand.Intn(2) == 1,
		TimeQuality:   rand.Uint32() % 8,
		UnlockedTime:  rand.Uint32() % 8,
		TriggerReason: rand.Uint32() % 8,
		MeasuredAt:    timestamp,
	}
}

func TestData(t *testing.T) {
	for k := int32(0); k < 46; k++ {
		test := generateStat(time.Now().UnixNano())
		buf, err := proto.Marshal(test)
		assert.NoError(t, err)
		data := &Stat{}
		err = proto.Unmarshal(buf, data)
		assert.NoError(t, err)
		assert.Equal(t, data.MeasuredAt, test.MeasuredAt)
		assert.Equal(t, data.TimeQuality, test.TimeQuality)
		assert.Equal(t, data.ConfigChange, test.ConfigChange)
	}
}

func BenchmarkStatSerialization(b *testing.B) {
	test := generateStat(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &Stat{}
		_ = proto.Unmarshal(buf, conf)
	}
}
