package gqlerrors

import (
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mitchellh/mapstructure"
)

type graphQLError struct {
	Data   map[string]interface{}      `json:"data"`
	Errors []*gqlerrors.FormattedError `json:"errors"`
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
