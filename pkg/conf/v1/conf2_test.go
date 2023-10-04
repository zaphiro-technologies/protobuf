package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func generateConf2Header() *Conf2Header {
	return &Conf2Header{
		SYNC:      uint32(0x1),
		FRAMESIZE: uint32(0x2),
		IDCODE:    uint32(0x3),
		SOC:       uint32(0x4),
		FRACSEC:   uint32(0x5),
		TIME_BASE: uint32(0x6),
		NUM_PMU:   uint32(0x7),
	}
}

func TestConf2Header(t *testing.T) {
	test := generateConf2Header()
	buf, err := proto.Marshal(test)
	assert.NoError(t, err)
	conf := &Conf2Header{}
	err = proto.Unmarshal(buf, conf)
	assert.NoError(t, err)
	assert.Equal(t, conf.SYNC, uint32(0x1))
	assert.Equal(t, conf.FRAMESIZE, uint32(0x2))
	assert.Equal(t, conf.IDCODE, uint32(0x3))
	assert.Equal(t, conf.SOC, uint32(0x4))
	assert.Equal(t, conf.FRACSEC, uint32(0x5))
	assert.Equal(t, conf.TIME_BASE, uint32(0x6))
	assert.Equal(t, conf.NUM_PMU, uint32(0x7))
}

func generateConfig() *Config {
	return &Config{
		STN:     "str",
		IDCODE:  uint32(0x1),
		FORMAT:  uint32(0x2),
		PHNMR:   uint32(0x3),
		ANNMR:   uint32(0x4),
		DGNMR:   uint32(0x5),
		CHNAM:   "chnam",
		PHUNIT:  []uint32{uint32(0x7)},
		ANUNIT:  []uint32{uint32(0x8)},
		DIGUNIT: []uint32{uint32(0x9)},
		FNOM:    uint32(0x10),
		CFGCNT:  uint32(0x11),
	}
}

func TestConfig(t *testing.T) {
	test := generateConfig()
	buf, err := proto.Marshal(test)
	assert.NoError(t, err)
	conf := &Config{}
	err = proto.Unmarshal(buf, conf)
	assert.NoError(t, err)
	assert.Equal(t, conf.STN, "str")
	assert.Equal(t, conf.IDCODE, uint32(0x1))
	assert.Equal(t, conf.FORMAT, uint32(0x2))
	assert.Equal(t, conf.PHNMR, uint32(0x3))
	assert.Equal(t, conf.ANNMR, uint32(0x4))
	assert.Equal(t, conf.DGNMR, uint32(0x5))
	assert.Equal(t, conf.CHNAM, "chnam")
	assert.Equal(t, conf.PHUNIT, []uint32{uint32(0x7)})
}

func generateConf2Frame() *Conf2Frame {
	return &Conf2Frame{
		Header:    generateConf2Header(),
		Configs:   []*Config{generateConfig()},
		DATA_RATE: uint32(0x3),
	}
}

func TestConf2Frame(t *testing.T) {
	test := generateConf2Frame()
	buf, err := proto.Marshal(test)
	assert.NoError(t, err)
	conf := &Conf2Frame{}
	err = proto.Unmarshal(buf, conf)
	assert.NoError(t, err)
	assert.Equal(t, conf.DATA_RATE, uint32(0x3))
	assert.Equal(t, conf.Header, generateConf2Header())
}

func BenchmarkConf2HeaderSerialization(b *testing.B) {
	test := generateConf2Header()
	for i := 0; i < b.N; i++ {
		buf, _ := proto.Marshal(test)
		conf := &Conf2Header{}
		_ = proto.Unmarshal(buf, conf)
	}
}
