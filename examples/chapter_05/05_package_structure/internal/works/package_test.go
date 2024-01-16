package works

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetState(t *testing.T) {
	assert.Equal(t, "initialized", GetState())
}
