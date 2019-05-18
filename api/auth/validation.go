package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/frayeralex/go-api/models"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

type validationError struct {
	ValidationError url.Values `json:"validation_error"`
}

func validateCredentials(w http.ResponseWriter, r *http.Request) error {
	var err error
	var cred models.Credentials

	rules := govalidator.MapData{
		"username": []string{"required", "min:6"},
		"password": []string{"required", "min:6", "max:20"},
	}

	opts := govalidator.Options{
		Request: r,
		Data:    &cred,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	fmt.Println(len(e))
	if len(e) > 0 {
		_ = json.NewEncoder(w).Encode(validationError{e})
		err = errors.New("api/auth.validateCredentials: validation failed")
	}
	return err
}
