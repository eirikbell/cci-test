package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Hello World", helloWorld())
}
