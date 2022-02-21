package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP(t *testing.T) {
	path := Path("foo", "bar", "baz")
	str := path.String()

	t.Log(str)
	assert.Equal(t, "foo.bar.baz", str)
}
