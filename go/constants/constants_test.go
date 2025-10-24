package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSourceType(t *testing.T) {
	_, err := NewSourceType("foo")
	assert.Error(t, err)

	st, err := NewSourceType(SourceTypeExternalService.String())
	assert.NoError(t, err)
	assert.Equal(t, SourceTypeExternalService, st)
}
