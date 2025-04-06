package ai

type AiInterface interface {
	// Query the AI with a prompt and return the response
	Query(prompt string) (string, error)
}

type AiConnector struct {
	// The AI instance to use for querying
	Ai AiInterface
}

// SetAi sets the AI instance to use for querying
func (ac *AiConnector) SetAi(ai AiInterface) {
	ac.Ai = ai
}

// Query queries the AI with a prompt and returns the response
func (ac *AiConnector) Query(prompt string) (string, error) {
	if ac.Ai == nil {
		return "", nil
	}
	return ac.Ai.Query(prompt)
}
