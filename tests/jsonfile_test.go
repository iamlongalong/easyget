package tests

import (
	"testing"

	"github.com/iamlongalong/easyget"
	"github.com/stretchr/testify/assert"
)

func TestGetKVFromJSONFile(t *testing.T) {
	g := easyget.NewJSONGetterFromJSONFile("files/test.json")

	v, ok := g.Get("name")
	assert.True(t, ok)
	assert.Equal(t, "longalong", v)
}
