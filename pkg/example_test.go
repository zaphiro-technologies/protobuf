package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	test := Hello()
	assert.Equal(t, test, "Hello, World!")
}

func TestHelloWithLog(t *testing.T) {
	test := HelloWithLog()
	assert.Equal(t, test, "Hello, World!")
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello()
	}
}
