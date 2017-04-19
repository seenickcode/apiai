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

	c1 := &Context{
		Name:     randomString(),
		Lifespan: 5,
	}
	answer, err := c.AddContext(sid, c1)
	assert.NoError(err)
	assert.Equal(200, answer.Status.Code)

	_, err = c.ClearContext(sid)
	assert.NoError(err)
}

func TestAPIAIContextParamsWithQuery(t *testing.T) {

	assert := assert.New(t)

	accessToken := os.Getenv("API_AI_ACCESS_TOKEN")
	if len(accessToken) == 0 {
		t.Fatalf("please set env var API_AI_ACCESS_TOKEN to run tests")
	}
	c := NewClient(accessToken)

	sid := randomString()

	ctxKey1 := randomString()
	ctxVal1 := randomString()
	c1 := &Context{
		Name:       randomString(),
		Lifespan:   5,
		Parameters: map[string]string{ctxKey1: ctxVal1},
	}
	_, err := c.AddContext(sid, c1)
	assert.NoError(err)

	answer, err := c.Query(sid, "bla bla bla")
	assert.Equal(ctxVal1, answer.Result.Contexts[0].Parameters[ctxKey1])
}
