package ai

import (
	"aictions/ai/llms/claude"
	"fmt"
)

// SetAiFromList sets the AI instance based on the provided name.
func (ac *AiConnector) SetAiFromList(name string) error {
	var ai AiInterface

	switch name {
	case "claude":
		ai = claude.NewClaudeAI("")
	default:
		return fmt.Errorf("unsupported AI name: %s", name)
	}

	ac.SetAi(ai)
	return nil
}
