package gqerrors

import (
	"fmt"

	"github.com/lab259/errors"
	"github.com/lab259/graphql/gqlerrors"
	"gopkg.in/gavv/httpexpect.v1"
)

type GraphQLError struct {
	Data   map[string]interface{}      `json:"data"`
	Errors []*gqlerrors.FormattedError `json:"errors"`
}

func prepare(mutateOrQuery string, actual interface{}) (*GraphQLError, error) {
	data, ok := actual.(*httpexpect.Object)
	if !ok {
		return nil, errors.New("`actual` is not an json object")
	}

	// Decoding GraphQL error
	var graphQLError GraphQLError
	err := decode(data.Raw(), &graphQLError)
	if err != nil {
		return nil, err
	}

	if len(graphQLError.Data) == 0 || len(graphQLError.Errors) == 0 {
		return nil, errors.New(fmt.Sprintf("expected an error is not `%s`", actual))
	}

	for key := range graphQLError.Data {
		if key != mutateOrQuery {
			return nil, errors.New(fmt.Sprintf("expected mutate or query name [%s] not is equal [%s]", key, mutateOrQuery))
		}
	}

	return &graphQLError, nil
}
