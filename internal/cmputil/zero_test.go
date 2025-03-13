package cmputil_test

import (
	"testing"

	"github.com/dieend/git-spice/internal/cmputil"
	"github.com/stretchr/testify/assert"
)

func TestZero(t *testing.T) {
	assert.False(t, cmputil.Zero(1))
	assert.True(t, cmputil.Zero(0))
}
