package gqlerrors

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/99designs/gqlgen/client"
	"github.com/vektah/gqlparser/gqlerror"
)

func prepare(name string, actual interface{}) (*gqlerror.Error, error) {
	switch t := actual.(type) {
	case json.RawMessage:
		var e gqlerror.List
		if err := json.Unmarshal(t, &e); err != nil {
			return nil, fmt.Errorf("Failed to unmarshal: %s", err.Error())
		}
		return e[0], nil
	case client.RawJsonError:
		return prepare(name, t.RawMessage)
	case gqlerror.Error:
		return &t, nil
	case *gqlerror.Error:
		return t, nil
	default:
		return nil, fmt.Errorf("%s matcher does not know how to handle %s", name, reflect.TypeOf(actual))
	}
}
