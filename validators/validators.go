package validators

import (
	"fmt"
	"net/mail"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func Validate(val interface{}) error {
	value := reflect.ValueOf(val)

	for i := 0; i < value.NumField(); i++ {
		tag := value.Type().Field(i).Tag.Get("validate")
		if tag == "" {
			continue
		}
		validator := validator{rule: "", fieldValue: value.Field(i), fieldName: value.Type().Field(i).Name}
		rules := strings.Split(tag, ",")
		for _, ruleName := range rules {
			validator.rule = ruleName
			if err := validator.checkRule(); err != nil {
				return err
			}
		}
	}
	return nil
}

type validator struct {
	rule       string
	fieldValue reflect.Value
	fieldName  string
}

func (r *validator) checkRule() error {
	switch {
	case strings.HasPrefix(r.rule, "min="):
		return r.validateMinLength()
	case strings.HasPrefix(r.rule, "max="):
		return r.validateMaxLength()
	case r.rule == "password":
		return r.validatePassword()
	case r.rule == "email":
		return r.validateEmail()
	}
	return nil
}

func (r *validator) validateMinLength() error {
	min, err := strconv.Atoi(strings.TrimPrefix(r.rule, "min="))
	if err != nil {
		return nil
	}

	if len(r.fieldValue.String()) < min {
		return fmt.Errorf("%s must be at least %d characters long", r.fieldName, min)
	}
	return nil
}

func (r *validator) validateMaxLength() error {
	max, err := strconv.Atoi(strings.TrimPrefix(r.rule, "max="))
	if err != nil {
		return nil
	}

	if len(r.fieldValue.String()) < max {
		return fmt.Errorf("%s must be at most %d characters long", r.fieldName, max)
	}
	return nil
}

func (r *validator) validatePassword() error {
	hasLetter := regexp.MustCompile(`[A-Za-z]`).MatchString(r.fieldValue.String())
	hasDigit := regexp.MustCompile(`\d`).MatchString(r.fieldValue.String())
	hasSpecialChar := regexp.MustCompile(`[@$!%*#?&]`).MatchString(r.fieldValue.String())
	hasProperLen := len(r.fieldValue.String()) >= 8
	if !(hasLetter && hasDigit && hasSpecialChar && hasProperLen) {
		return fmt.Errorf("%s must have at least eight characters, at least one letter, one number and one special character", r.fieldName)
	}

	return nil
}

func (r *validator) validateEmail() error {
	_, err := mail.ParseAddress(r.fieldValue.String())
	if err != nil {
		return fmt.Errorf("%s must be a valid email", r.fieldName)
	}
	return nil
}
