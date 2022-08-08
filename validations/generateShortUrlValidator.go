package validations

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type GenerateShortUrlValidator struct {
	Url string `validate:"required,url"`
}

func GenerateShortUrlValidatorFunc(sl validator.StructLevel) {
	validSl := sl.Current().Interface().(GenerateShortUrlValidator)
	isValid := false
	jsonmarshal, _ := json.Marshal(validSl)
	var m map[string]interface{}
	_ = json.Unmarshal(jsonmarshal, &m)
	for _, v := range m {
		if len(fmt.Sprintf("%s", v)) > 0 {
			isValid = true
		}
	}
	if !isValid {
		sl.ReportError(validSl, "url", "Url", "required", "")
	}
}
