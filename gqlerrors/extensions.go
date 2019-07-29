package gqlerrors

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/gavv/httpexpect.v1"
)

type FormattedError struct {
	Message    string                 `json:"message"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

type graphQLError struct {
	Data   map[string]interface{} `json:"data"`
	Errors []*FormattedError      `json:"errors"`
}

// decode an objects
func (g *graphQLError) decode(input interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  g,
	})

	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func prepareFromJSON(mutateOrQuery string, actual interface{}) (*graphQLError, error) {
	data, ok := actual.(*httpexpect.Object)
	if !ok {
		return nil, fmt.Errorf("`actual` is not an json object")
	}

	// Decoding GraphQL error
	var graphQLError graphQLError
	err := graphQLError.decode(data.Raw())
	if err != nil {
		return nil, err
	}

	if len(graphQLError.Errors) == 0 {
		return nil, fmt.Errorf("expected an error is not `%s`", actual)
	}

	for key := range graphQLError.Data {
		if key != mutateOrQuery {
			return nil, fmt.Errorf("expected mutate or query name [%s] not is equal [%s]", key, mutateOrQuery)
		}
	}

	return &graphQLError, nil
}
