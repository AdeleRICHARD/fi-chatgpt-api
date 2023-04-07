package chat

import (
	"context"

	gogpt "github.com/sashabaranov/go-openai"
)

func AskToChatGPT(request string) (string, error) {
	canalAsk := make(chan string)
	go chatGPTAsync(canalAsk, request)
	resp := <-canalAsk
	return resp, nil // récupérer une valeur d'un channel
}

func chatGPTAsync(c chan string, request string) {
	key := gogpt.NewClient("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")

	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci003,
		MaxTokens:   1000,
		Prompt:      request,
		Temperature: 0.7,
	}
	resp, err := key.CreateCompletion(ctx, req)
	if err != nil {
		c <- err.Error()
	} else {
		c <- resp.Choices[0].Text
	}

}
