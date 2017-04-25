# Overview

A simple API.ai SDK for Go.

# Features

- Query
  - Supports all query response fields, including:
    - Multiple Fulfillment message types
    - Params
- List Intents
- Add/Fetch/Delete/Clear Context

# Examples

Example: Perform a query and access some response params.

```golang

  c := apiai.NewClient("<YOUR API.AI ACCESS TOKEN>")

  answer, err := c.Query("<SOME UNIQUE SESSION ID>", "Hi there!")

  params := struct {
    MyParam1 string `json:"my-param-1"`
    MyParam2 string `json:"my-param-2"`
  }{}
  json.Unmarshal(*answer.Result.Parameters, &params)

  fmt.Printf("got params: %+v", params)

```

# Running Tests

NOTE: Running tests require an API.ai Access Token. The tests use this token to
perform a query triggering the Default Fallback Intent for that account to ensure
basic tests pass.

``API_AI_ACCESS_TOKEN=<token> go test``
