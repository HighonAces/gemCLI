package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <text>", os.Args[0])
	}

	text := os.Args[1]

	// Initialize the generative AI client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatalf("Failed to create generative AI client: %v", err)
	}
	defer client.Close()

	// Make the API call with the text input
	model := client.GenerativeModel("gemini-2.0-flash-exp")
	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text("command-line style with 7 examples")},
	}
	session := model.StartChat()
	session.History = []*genai.Content{}
	resp, err := session.SendMessage(ctx, genai.Text(text))
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Printf("%v\n", part)
	}
}
