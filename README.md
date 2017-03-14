# Overview

A simple API.ai wrapper with flexible query response object leveraging *json.RawData.

## Use

```

  c := apiai.NewClient("<YOUR API.AI ACCESS TOKEN>")

  // perform a query and access some response params
  answer, err := c.Query("<SOME UNIQUE SESSION ID>", "Hi there!")

  params := struct {
    MyParam1 string `json:"my-param-2"`
    MyParam2   string `json:"my-param-1"`
  }{}
  json.Unmarshal(*answer.Result.Parameters, &params)

  fmt.Printf("got params: %+v", params)

```

## Running Tests

``API_AI_ACCESS_TOKEN=<token> go test``
