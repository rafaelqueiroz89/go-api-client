//TODO: Api tests not needed in this level

package client

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"math"
	"net/http"
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

func TestAddQueryParams(t *testing.T) {
	type args struct {
		query   map[string]string
		request *http.Request
	}

	var queryParams args
	queryParams.query = make(map[string]string)
	queryParams.query["wololo"] = "p1"
	queryParams.query["bla"] = "p2"
	queryParams.request, _ = http.NewRequest("GET", "test", nil)

	AddQueryParams(queryParams.query, queryParams.request)
	assert.Equal(t, "bla=p2&wololo=p1", queryParams.request.URL.RawQuery)
}
