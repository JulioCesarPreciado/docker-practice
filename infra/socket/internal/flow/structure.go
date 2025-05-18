package flow

// StepConfig defines the metadata for each step in the form flow.
type StepConfig struct {
	Message   string `json:"message"`
	Type      string `json:"type"`
	Reference string `json:"reference"`
	Required  bool   `json:"required"`
	MinLength int    `json:"minLength"`
	MaxLength int    `json:"maxLength"`
}

// StepConfigs contains the configuration for each form step.
var StepConfigs = map[string]StepConfig{
	"name": {
		Message:   "¿Cuál es tu nombre(s)?",
		Type:      "text",
		Reference: "",
		Required:  true,
		MinLength: 2,
		MaxLength: 100,
	},
	"lastname": {
		Message:   "¿Cuál es tu apellido(s)?",
		Type:      "text",
		Reference: "",
		Required:  true,
		MinLength: 2,
		MaxLength: 100,
	},
	"type": {
		Message:   "¡Gracias! 😁 Comenzemos con los datos del vehículo ¿Qué tipo de vehículo es? (camioneta, SUV, sedan, motocicleta etc.)",
		Type:      "text",
		Reference: "",
		Required:  true,
		MinLength: 3,
		MaxLength: 100,
	},
	"brand": {
		Message:   "¿Cuál es la marca de tu vehículo?",
		Type:      "select",
		Reference: "brands",
		Required:  true,
		MinLength: 2,
		MaxLength: 100,
	},
	"model": {
		Message:   "¿Qué modelo es?",
		Type:      "select",
		Reference: "models",
		Required:  true,
		MinLength: 1,
		MaxLength: 100,
	},
	"year": {
		Message:   "¿Año del vehículo?",
		Type:      "select",
		Reference: "years",
		Required:  true,
		MinLength: 4,
		MaxLength: 4,
	},
	"version": {
		Message:   "¿Cuál es la versión?",
		Type:      "select",
		Reference: "versions",
		Required:  true,
		MinLength: 1,
		MaxLength: 100,
	},
	"birthdate": {
		Message:   "¿Cuál es tu fecha de nacimiento?",
		Type:      "date",
		Required:  true,
		MinLength: 10,
		MaxLength: 10,
	},
}

// GetStepConfig returns the full StepConfig for a given step name.
func GetStepConfig(step string) (StepConfig, bool) {
	cfg, ok := StepConfigs[step]
	return cfg, ok
}
