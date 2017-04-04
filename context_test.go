package apiai

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIAIContexts(t *testing.T) {

	assert := assert.New(t)

	accessToken := os.Getenv("API_AI_ACCESS_TOKEN")
	if len(accessToken) == 0 {
		t.Fatalf("please set env var API_AI_ACCESS_TOKEN to run tests")
	}
	c := NewClient(accessToken)

	sid := randomString()

	contextName := randomString()
	answer, err := c.AddContext(sid, contextName, 5)
	assert.NoError(err)
	assert.Equal(200, answer.Status.Code)

	_, err = c.ClearContext(sid)
	assert.NoError(err)
}
