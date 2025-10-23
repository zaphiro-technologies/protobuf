package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSourceType(t *testing.T) {
	st, err := NewSourceType(3)
	assert.Equal(t, SourceTypeExternalService, st)
	assert.NoError(t, err)

	st, err = NewSourceType(_sourceTypeTag.ToInt8())
	assert.Equal(t, SourceTypeUnspecified, st)
	assert.Error(t, err)
}
