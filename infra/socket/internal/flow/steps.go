package flow

import "encoding/json"

// ProcessMessage dynamically processes form steps based on the order defined in StepDefinitions.
func ProcessMessage(msg []byte, state map[string]string) []byte {
	text := string(msg)

	for _, step := range StepDefinitions {
		if _, ok := state[step.Name]; !ok {
			valid, errorMessage := ValidateStep(step.Name, text)
			if !valid {
				return []byte(errorMessage)
			}
			state[step.Name] = text

			// Encontrar siguiente paso
			for i, def := range StepDefinitions {
				if def.Name == step.Name && i+1 < len(StepDefinitions) {
					next := StepDefinitions[i+1].Name
					cfg, ok := GetStepConfig(next)
					if ok {
						jsonCfg, _ := json.Marshal(cfg)
						return jsonCfg
					}
				}
			}

			return []byte("✅ Formulario completado.")
		}
	}

	return []byte("✅ Formulario ya completado.")
}
