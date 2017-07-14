package rest

import (
	"encoding/json"
	"github.com/geodan/gost/sensorthings/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestHandleRoot(t *testing.T) {
	// arrange
	count := 0
	eps := NewMockAPI().GetEndpoints()
	for _, e := range *eps {
		if e.ShowOutputInfo() {
			count++
		}
	}

	// act
	r := request("GET", "/v1.0", nil)
	arrayResponse := models.ArrayResponseEndpoint{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &arrayResponse)

	// assert
	assert.NotNil(t, arrayResponse)
	assert.Len(t, arrayResponse.Data, count)
}