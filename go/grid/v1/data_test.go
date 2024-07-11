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

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func generateData(dataType int, value uint64, timestamp int64) *Data {
	return &Data{
		DataType:   DataType(dataType),
		Value:      &value,
		MeasuredAt: timestamp,
	}
}

func generateDataSet(producerId string, data map[string]*Data) *DataSet {
	return &DataSet{ProducerId: producerId, Data: data}
}

func TestData(t *testing.T) {
	for k := 0; k < 46; k++ {
		test := generateData(k, rand.Uint64(), time.Now().UnixNano())
		buf, err := proto.Marshal(test)
		assert.NoError(t, err)
		data := &Data{}
		err = proto.Unmarshal(buf, data)
		assert.NoError(t, err)
		assert.Equal(t, data.MeasuredAt, test.MeasuredAt)
		assert.Equal(t, data.DataType, test.DataType)
		assert.Equal(t, data.Value, test.Value)
	}
}

func TestDataSet(t *testing.T) {
	dataMap := map[string]*Data{}
	for k := 0; k < 46; k++ {
		dataMap[uuid.New().String()] = generateData(k, rand.Uint64(), time.Now().UnixNano())
	}
	test := generateDataSet("myDataSet", dataMap)
	buf, err := proto.Marshal(test)
	assert.NoError(t, err)
	dataSet := &DataSet{}
	err = proto.Unmarshal(buf, dataSet)
	assert.NoError(t, err)
	assert.Equal(t, dataSet.ProducerId, "myDataSet")
	assert.Equal(t, len(dataSet.Data), 46)
}

func BenchmarkDataSerialization(b *testing.B) {
	test := generateData(rand.Intn(46), rand.Uint64(), time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &Data{}
		_ = proto.Unmarshal(buf, conf)
	}
}

func BenchmarkDataSetSerialization(b *testing.B) {
	dataMap := map[string]*Data{}
	for k := 0; k < 46; k++ {
		dataMap[uuid.New().String()] = generateData(k, rand.Uint64(), time.Now().UnixNano())
	}
	test := generateDataSet("myDataSet", dataMap)
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &Data{}
		_ = proto.Unmarshal(buf, conf)
	}
}
