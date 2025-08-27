package handler

import (
	"fmt"
	"net/url"
	"strings"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func errInvalidURL() error {
	return fmt.Errorf("invalid URL format - URL must be a valid http or https URL")
}

type CreateURLRequest struct {
	URL string `json:"url"`
}

func (r *CreateURLRequest) Validate() error {
	if strings.TrimSpace(r.URL) == "" {
		return errParamIsRequired("url", "string")
	}

	parsedURL, err := url.Parse(r.URL)
	if err != nil {
		return errInvalidURL()
	}

	if parsedURL.Scheme == "" {
		return errInvalidURL()
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return errInvalidURL()
	}

	if parsedURL.Host == "" {
		return errInvalidURL()
	}

	return nil
}
