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
	"encoding/xml"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func generateTopology(timestamp int64) *Topology {
	tp, _ := os.ReadFile("FullGrid_TP.xml") // read the content of file

	return &Topology{
		CreatedAt: timestamp,
		Tp:        tp,
	}
}

type RDF struct {
	XMLName   xml.Name `xml:"RDF"`
	Rdf       string   `xml:"rdf,attr"`
	FullModel struct {
		Profile string `xml:"Model.profile"`
	} `xml:"FullModel"`
}

func TestTopology(t *testing.T) {
	test := generateTopology(time.Now().UnixMilli())
	buf, err := proto.Marshal(test)
	assert.NoError(t, err)
	topology := &Topology{}
	err = proto.Unmarshal(buf, topology)
	assert.NoError(t, err)
	var tp RDF
	err = xml.Unmarshal(topology.Tp, &tp)
	assert.NoError(t, err)
	assert.Equal(t, "http://iec.ch/TC57/ns/CIM/Topology-EU/3.0", tp.FullModel.Profile)
}

func BenchmarkTopologySerialization(b *testing.B) {
	test := generateTopology(time.Now().UnixMilli())
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		topology := &Topology{}
		_ = proto.Unmarshal(buf, topology)
	}
}
