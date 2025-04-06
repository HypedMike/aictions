package functions

// FunctionModel represents a function that can be called by the AI
type FunctionModel struct {
	// Name of the function
	Name string `json:"name"`

	// Description of the function
	Description string `json:"description"`

	// Parameters of the function
	Parameters []struct {
		// Name of the parameter
		Name string `json:"name"`

		// Type of the parameter
		Type string `json:"type"`

		// Type of the parameter
		Value any `json:"value"`
	} `json:"parameters"`
}
