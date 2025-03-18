package validators

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kalinchkma/geko/internal/utils"
)

func NormalizeJsonValidationError(err error, messages map[string]string) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, vErr := range validationErrors {
			fieldKey := vErr.Field() + "." + vErr.Tag()

			var msg string
			if customMsg, exists := messages[fieldKey]; exists {
				msg = customMsg
			} else {
				msg = fmt.Sprintf("Invalid value '%v' for field '%s' (expected: %s)", vErr.Value(), vErr.Field(), vErr.Tag())
			}

			// errors[strings.ToLower(vErr.Field())] = msg
			errors[utils.CamelToSnake(vErr.Field())] = msg
		}
	} else {

		switch e := err.(type) {
		case *json.UnmarshalTypeError:
			errors[e.Field] = fmt.Sprintf(
				"field '%s' has an invalid type, expected %s but got %s",
				e.Field, e.Type.String(), e.Value,
			)
		case *json.SyntaxError:
			errors["message"] = fmt.Sprintf("syntax error at offset %d: %v", e.Offset, err)
		case *json.InvalidUnmarshalError:
			errors["message"] = fmt.Sprintf("invalid unmarshal: %v", e)
		case *json.UnsupportedTypeError:
			errors["message"] = fmt.Sprintf("unsupported type: %v", e.Type)
		case *json.MarshalerError:
			errors["message"] = fmt.Sprintf("error marshaling JSON: %v", e.Err)
		default:
			errors["message"] = fmt.Sprintf("unexpected error: %v", err)
		}
	}
	return errors
}
