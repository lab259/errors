package gqerrors

import (
	"reflect"

	"github.com/gofrs/uuid"
	"github.com/mitchellh/mapstructure"
)

// decode an objects
func decode(input interface{}, output interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  output,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			toUUIDHookFunc,
		),
	})

	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func toUUIDHookFunc(f, t reflect.Type, data interface{}) (interface{}, error) {
	if t != reflect.TypeOf(uuid.UUID{}) {
		return data, nil
	}

	switch f.Kind() {
	case reflect.String:
		return uuid.FromString(data.(string))
	default:
		return data, nil
	}
}
