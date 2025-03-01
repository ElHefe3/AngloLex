package openapi

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

const OPENAI_TOKEN = "OPENAI_TOKEN"

func FindWordFor(sentence string) (string, error) {
    openaiToken := os.Getenv(OPENAI_TOKEN)
	if openaiToken == "" {
		log.Fatalf("[FindWordFor] Openai token not found in environment variables")
	}

	client := openai.NewClient(
		option.WithAPIKey(openaiToken),
	)

	ctx := context.Background()

	prompt := fmt.Sprintf("Provide a single word that best describes the given prompt: %s\n\n# Output Format\n- The response should be a single descriptive word related to the prompt.", sentence)

	completion, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
		Seed:       openai.Int(1),
		Model:      openai.F(openai.ChatModelGPT4o),
		Temperature: openai.F(0.3),
	})
	if err != nil {
		return "", err
	}

	return completion.Choices[0].Message.Content, nil
}
