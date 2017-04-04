package apiai

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIAIIntents(t *testing.T) {

	assert := assert.New(t)

	accessToken := os.Getenv("API_AI_ACCESS_TOKEN")
	if len(accessToken) == 0 {
		t.Fatalf("please set env var API_AI_ACCESS_TOKEN to run tests")
	}
	c := NewClient(accessToken)

	list, err := c.ListIntents()
	assert.NoError(err)
	assert.NotEmpty(list)
	assert.NotEmpty(list[0].Name)
}
