package apiai

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIAIQueries(t *testing.T) {

	assert := assert.New(t)

	accessToken := os.Getenv("API_AI_ACCESS_TOKEN")
	if len(accessToken) == 0 {
		t.Fatalf("please set env var API_AI_ACCESS_TOKEN to run tests")
	}
	c := NewClient(accessToken)

	msg := "bla bla bla"
	answer, err := c.Query(randomString(), msg)
	assert.NoError(err)
	assert.Equal(200, answer.Status.Code)
	assert.NotEmpty(answer.Result.Fulfillment.TextMessages)
	assert.NotEmpty(answer.Result.Fulfillment.TextMessages[0].Speech)
}
