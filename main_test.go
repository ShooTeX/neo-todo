package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	_, err := Init()

	assert.Nil(t, err)
}
