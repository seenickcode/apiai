package apiai

import (
	"os"
	"strings"
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
	a1, err := c.AddContext(sid, c1)
	assert.NoError(err)
	assert.Equal(200, a1.Status.Code)

	c2 := &Context{
		Name:     randomString(),
		Lifespan: 5,
	}
	a2, err := c.AddContext(sid, c2)
	assert.NoError(err)
	assert.Equal(200, a2.Status.Code)

	contexts, err := c.FetchContexts(sid)
	assert.NoError(err)
	assert.Equal(2, len(contexts))
	c1Found := false
	for _, c := range contexts {
		if c.Name == strings.ToLower(c1.Name) {
			c1Found = true
			break
		}
	}
	assert.True(c1Found)
	c2Found := false
	for _, c := range contexts {
		if c.Name == strings.ToLower(c2.Name) {
			c2Found = true
			break
		}
	}
	assert.True(c2Found)

	_, err = c.DeleteContext(sid, c1.Name)
	assert.NoError(err)

	contexts, err = c.FetchContexts(sid)
	assert.NoError(err)
	assert.Equal(1, len(contexts))
	assert.Equal(strings.ToLower(c2.Name), contexts[0].Name) // NOTE api.ai lowercases any set context names

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
	m1 := map[string]interface{}{
		ctxKey1: ctxVal1,
	}
	c1 := &Context{
		Name:       randomString(),
		Lifespan:   5,
		Parameters: m1,
	}
	_, err := c.AddContext(sid, c1)
	assert.NoError(err)

	answer, err := c.Query(sid, "bla bla bla")
	assert.Equal(ctxVal1, answer.Result.Contexts[0].Parameters[ctxKey1])
}
