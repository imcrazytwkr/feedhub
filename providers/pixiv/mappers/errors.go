package mappers

import (
	"errors"

	"github.com/valyala/fastjson"
)

func processErrorFields(contents *fastjson.Value) (error, bool) {
	if !contents.GetBool(errorKey) {
		return nil, false
	}

	message := contents.GetStringBytes(messageKey)
	if len(message) == 0 {
		return nil, true
	}

	return errors.New(string(message)), true
}
