package controller_single

import (
	"aictions/ai"
	"aictions/functions"
	"aictions/prompts"
	"fmt"
)

type Options struct {
	Ai    *ai.AiConnector
	Query string
}

func Serve(opt Options) (string, error) {
	ai := opt.Ai
	query := opt.Query

	// Check if ai and query are not nil or empty
	if ai == nil {
		panic("AI instance is nil")
	}
	if query == "" {
		panic("Query is empty")
	}

	// Get the functions
	functionsHandler := functions.NewFunctionsHandler()
	functions, err := functionsHandler.SprintFunction()
	if err != nil {
		fmt.Printf("Failed to get functions: %v", err)
	}

	// Build the prompt
	prompt := prompts.ChooseBetweenFunctionsSingle(struct {
		Prompt    string
		Functions string
	}{
		Prompt:    query,
		Functions: functions,
	})

	// Send the prompt to the AI
	response, err := ai.Query(prompt)
	if err != nil {
		return "", fmt.Errorf("failed to query AI: %v", err)
	}

	// Print the response
	return response, nil
}
