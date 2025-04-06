package claude

import (
	"context"

	cl "github.com/potproject/claude-sdk-go"
)

type ClaudeAI struct {
	ApiKey string
	client *cl.Client
}

func NewClaudeAI(apiKey string) *ClaudeAI {
	c := cl.NewClient(apiKey)
	return &ClaudeAI{
		ApiKey: apiKey,
		client: c,
	}
}

// Query sends a prompt to the Claude AI and returns the response.
//
// It uses the "claude-3-7-sonnet-20250219" model and sets a maximum token limit of 1024.
// The function creates a context for the request and handles any errors that may occur.
func (c *ClaudeAI) Query(prompt string) (string, error) {
	m := cl.RequestBodyMessages{
		Model:     "claude-3-7-sonnet-20250219",
		MaxTokens: 1024,
		Messages: []cl.RequestBodyMessagesMessages{
			{
				Role:    cl.MessagesRoleUser,
				Content: prompt,
			},
		},
	}
	ctx := context.Background()
	res, err := c.client.CreateMessages(ctx, m)
	if err != nil {
		return "", err
	}

	return res.Content[0].Text, nil
}
