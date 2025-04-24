package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func UserRegister(err error) map[string]string {
	result := map[string]string{}

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range errs {
			field := strings.ToLower(fieldErr.Field())

			// You can customize messages based on tag
			var msg string
			switch fieldErr.Tag() {
			case "required":
				msg = fmt.Sprintf("اجباری است.")
			case "email":
				msg = fmt.Sprintf("باید یک آدرس ایمیل معتبر باشد.")
			case "min":
				msg = fmt.Sprintf("باید حداقل %s کاراکتر باشد.", fieldErr.Param())
			case "max":
				msg = fmt.Sprintf("باید حداکثر %s کاراکتر باشد.", fieldErr.Param())
			case "gte":
				msg = fmt.Sprintf("باید بزرگتر یا مساوی %s کاراکتر باشد.", fieldErr.Param())
			case "lte":
				msg = fmt.Sprintf("باید کوچکتر یا مساوی %s کاراکتر باشد.", fieldErr.Param())
			case "alphanum":
				msg = fmt.Sprintf("فقط باید شامل حروف و عدد باشد.")
			case "e164":
				msg = fmt.Sprintf("باید یک شماره تلفن معتر باشد.")
			case "username":
				msg = fmt.Sprintf("باید با یک حرف شروع شود و فقط شامل حروف، اعداد، زیرخط و نقطه باشد.")
			default:
				msg = fmt.Sprintf("نامعتبر است.")
			}

			result[field] = msg
		}
	} else {
		result["error"] = "Invalid request format"
	}

	return result
}
