package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	actual := getHello()
	expected := "Hello, world!"
	assert.Equal(t, actual, expected)
}
