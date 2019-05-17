package gqlerrors

import (
	"fmt"

	"github.com/lab259/graphql/gqlerrors"
	"gopkg.in/gavv/httpexpect.v1"
)

type graphQLError struct {
	Data   map[string]interface{}      `json:"data"`
	Errors []*gqlerrors.FormattedError `json:"errors"`
}

func prepare(mutateOrQuery string, actual interface{}) (*graphQLError, error) {
	data, ok := actual.(*httpexpect.Object)
	if !ok {
		return nil, fmt.Errorf("`actual` is not an json object")
	}

	// Decoding GraphQL error
	var graphQLError graphQLError
	err := Decode(data.Raw(), &graphQLError)
	if err != nil {
		return nil, err
	}

	if len(graphQLError.Data) == 0 || len(graphQLError.Errors) == 0 {
		return nil, fmt.Errorf("expected an error is not `%s`", actual)
	}

	for key := range graphQLError.Data {
		if key != mutateOrQuery {
			return nil, fmt.Errorf("expected mutate or query name [%s] not is equal [%s]", key, mutateOrQuery)
		}
	}

	return &graphQLError, nil
}
