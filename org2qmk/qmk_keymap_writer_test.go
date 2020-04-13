package org2qmk

import (
	"testing"

	"github.com/niklasfasching/go-org/org"
	"github.com/stretchr/testify/assert"
)

func TestQmkKeymapWriter(t *testing.T) {
	assert.Implements(t, (*org.Writer)(nil), NewQmkKeymapWriter())
}
