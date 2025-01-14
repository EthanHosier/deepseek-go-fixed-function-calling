package deepseek_examples

import (
	"context"
	"fmt"
	"log"
	"os"

	deepseek "github.com/EthanHosier/deepseek-go-fixed-function-calling"
)

func Chat() {
	client := deepseek.NewClient(os.Getenv("DEEPSEEK_API"))
	request := &deepseek.ChatCompletionRequest{
		Model: deepseek.DeepSeekChat,
		Messages: []deepseek.ChatCompletionMessage{
			{Role: deepseek.ChatMessageRoleUser, Content: "Which is the tallest mountain in the world?"},
			{Role: deepseek.ChatMessageRoleSystem, Content: "Answer every question using slang"},
		},
	}
	ctx := context.Background()
	response, err := client.CreateChatCompletion(ctx, request)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("Response:", response.Choices[0].Message.Content)
}
