package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateURLRequest struct {
	URL string `json:"url"`
}

func (r *CreateURLRequest) Validate() error {
	if r.URL == "" {
		return errParamIsRequired("url", "string")
	}
	return nil
}
