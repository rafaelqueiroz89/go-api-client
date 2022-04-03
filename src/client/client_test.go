package client

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestPrepareBody_Returns_Data(t *testing.T) {
	type DummyRequest struct {
		Data string `json:"data"`
	}

	data, err := PrepareBody(DummyRequest{})

	assert.Equal(t, "{\"data\":\"\"}", bytes.NewBuffer(data).String())
	assert.Nil(t, err)

	data, err = PrepareBody(math.Inf(1))
	assert.Nil(t, data)
	assert.NotNil(t, err)
}
