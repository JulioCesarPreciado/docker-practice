package flow

import (
	"strings"
	"time"
)

// LengthBetween1And255 checks if the input string length is between 1 and 255 characters.
func LengthBetween1And255(s string) bool {
	return len(s) >= 1 && len(s) <= 255
}

// IsValidAdultBirthdate validates if the input is a valid date in YYYY-MM-DD format and if the user is 18+ years old.
func IsValidAdultBirthdate(s string) bool {
	s = strings.TrimSpace(s)
	birthdate, err := time.Parse("2006-01-02", s)
	if err != nil {
		return false
	}

	today := time.Now()
	age := today.Year() - birthdate.Year()
	if today.Month() < birthdate.Month() || (today.Month() == birthdate.Month() && today.Day() < birthdate.Day()) {
		age--
	}

	return age >= 18
}

// StepRule defines the name of the step, its validation function, and an error message.
type StepRule struct {
	Name         string
	Validator    func(string) bool
	ErrorMessage string
}

// StepDefinitions defines the ordered form steps and their associated validation rules.
var StepDefinitions = []StepRule{
	{
		Name:         "name",
		Validator:    LengthBetween1And255,
		ErrorMessage: "❌ El nombre no puede estar vacío.",
	},
	{
		Name:         "lastname",
		Validator:    LengthBetween1And255,
		ErrorMessage: "❌ El apellido no puede estar vacío.",
	},
	{
		Name:         "type",
		Validator:    LengthBetween1And255,
		ErrorMessage: "❌ El tipo de vehículo es obligatorio.",
	},
	{
		Name:         "brand",
		Validator:    LengthBetween1And255,
		ErrorMessage: "❌ La marca es obligatoria.",
	},
	{
		Name:         "model",
		Validator:    LengthBetween1And255,
		ErrorMessage: "❌ El modelo es obligatorio.",
	},
	{
		Name:         "year",
		Validator:    LengthBetween1And255,
		ErrorMessage: "❌ El año debe tener 4 dígitos (ej. 2020).",
	},
	{
		Name:         "version",
		Validator:    LengthBetween1And255,
		ErrorMessage: "❌ La versión no puede estar vacía.",
	},
	{
		Name:         "birthdate",
		Validator:    IsValidAdultBirthdate,
		ErrorMessage: "❌ No eres mayor de edad.",
	},
}

// ValidateStep checks if the given value is valid for the specified step.
func ValidateStep(step string, value string) (bool, string) {
	for _, def := range StepDefinitions {
		if def.Name == step {
			if def.Validator(value) {
				return true, ""
			}
			return false, def.ErrorMessage
		}
	}
	return false, "❌ Paso no reconocido."
}
