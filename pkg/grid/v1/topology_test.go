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
	tp, _ := os.ReadFile("FullGrid_TP.xml")   // read the content of file
	ssh, _ := os.ReadFile("FullGrid_SSH.xml") // read the content of file

	return &Topology{
		CreatedAt: uint64(timestamp),
		Tp:        tp,
		Ssh:       ssh,
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
	var ssh RDF
	err = xml.Unmarshal(topology.Ssh, &ssh)
	assert.NoError(t, err)
	assert.Equal(t, "http://iec.ch/TC57/ns/CIM/SteadyStateHypothesis-EU/3.0", ssh.FullModel.Profile)
}

func BenchmarkTopologySerialization(b *testing.B) {
	test := generateTopology(time.Now().UnixMilli())
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		topology := &Topology{}
		_ = proto.Unmarshal(buf, topology)
	}
}
