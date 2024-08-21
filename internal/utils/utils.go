package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func getMessageByTag(tagName string, fieldName string, paramValue string) string {
	var msg string
	switch tagName {
	case "required":
		msg = fmt.Sprintf("The field %s is required", fieldName)
	case "min":
		if paramValue != "" {
			msg = fmt.Sprintf("The field %s must have at least %s characters", fieldName, paramValue)
		} else {
			msg = fmt.Sprintf("The field %s must have at least the minimum length", fieldName)
		}
	case "max":
		if paramValue != "" {
			msg = fmt.Sprintf("The field %s must not exceed %s characters", fieldName, paramValue)
		} else {
			msg = fmt.Sprintf("The field %s must not exceed the maximum length", fieldName)
		}
	case "email":
		msg = fmt.Sprintf("The field %s must be a valid email address", fieldName)
	default:
		msg = fmt.Sprintf("The field %s is invalid", fieldName)
	}
	return msg
}

func GetCustomErrorMessages(err error) []string {
	var errorMessages []string

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			errorMessages = append(errorMessages, getMessageByTag(e.ActualTag(), e.Field(), e.Param()))
		}
	} else {
		errorMessages = append(errorMessages, "Ocorreu um erro desconhecido")
	}

	return errorMessages
}
