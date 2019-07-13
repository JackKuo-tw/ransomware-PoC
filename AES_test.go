package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// "1" * 32
var testkey = []byte("11111111111111111111111111111111")

// "1" * 16
var testData = []byte("1111111111111111")
var testResult = []byte{0x9a, 0x58, 0x71, 0xe7, 0x57, 0x8, 0x61, 0x21, 0x22, 0x1d, 0x92, 0x8a, 0x8d, 0x58, 0xc5, 0x19, 0x68, 0x15, 0xfb, 0x4f, 0x7b, 0xfc, 0x7a, 0x46, 0xe1, 0xc6, 0xd9, 0x60, 0xca, 0x60, 0xc0, 0x33}

func TestCFBEncrypt(t *testing.T) {

	enc, err := CFBEncrypt(testData, testkey)
	assert.Nil(t, err)
	assert.Equal(t, enc, testResult)
}

func TestCFBDecrypt(t *testing.T) {
	data, err := CFBDecrypt(testResult, testkey)
	assert.Nil(t, err)
	assert.Equal(t, testData, data)
}
