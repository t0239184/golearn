package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	url := "/sys/ping"
	response := Get(&url, nil)

	/* Processing response */
	json, err := ConvertResponseToMap(response)
	if err != nil {
		t.Error(err)
	}
	value, _ := json["health"]

	/* Assert */
	expectedStatus := http.StatusOK
	assert.Equal(t, expectedStatus, response.Code)
	assert.True(t, value.(bool))
}
