package gqlerrors

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/99designs/gqlgen/client"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/vektah/gqlparser/gqlerror"
	"gopkg.in/gavv/httpexpect.v1"
)

type errOutput struct {
	Gqlerror       *gqlerror.Error
	FormattedError *gqlerrors.FormattedError
}

func prepare(name string, actual interface{}) (*errOutput, error) {
	switch t := actual.(type) {
	case json.RawMessage:
		var e gqlerror.List
		if err := json.Unmarshal(t, &e); err != nil {
			return nil, fmt.Errorf("Failed to unmarshal: %s", err.Error())
		}
		return &errOutput{
			Gqlerror: e[0],
		}, nil
	case client.RawJsonError:
		return prepare(name, t.RawMessage)
	case gqlerror.Error:
		return &errOutput{
			Gqlerror: &t,
		}, nil
	case *gqlerror.Error:
		return &errOutput{
			Gqlerror: t,
		}, nil
	case *httpexpect.Object:
		var graphQLError graphQLError
		err := graphQLError.decode(t.Raw())
		if err != nil {
			return nil, err
		}

		if len(graphQLError.Errors) > 0 {
			return &errOutput{
				FormattedError: graphQLError.Errors[0],
			}, nil
		}

		return nil, fmt.Errorf("%s httpexpect.Object not is error %s", name, reflect.TypeOf(actual))
	default:
		return nil, fmt.Errorf("%s matcher does not know how to handle %s", name, reflect.TypeOf(actual))
	}
}
