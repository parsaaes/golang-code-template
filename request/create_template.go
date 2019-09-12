package request

import validation "github.com/go-ozzo/ozzo-validation"

type CreateTemplateRequest struct {
	Name string `json:"name"`
}

func (c CreateTemplateRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
	)
}
