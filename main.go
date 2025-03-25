package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func readPartsFromCSV(filePath string) ([]genai.Part, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var parts []genai.Part
	for _, record := range records {
		if len(record) < 2 {
			continue
		}
		parts = append(parts, genai.Text(fmt.Sprintf("input: %s", record[0])))
		parts = append(parts, genai.Text(fmt.Sprintf("output: %s", record[1])))
	}

	return parts, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <cli_input>", os.Args[0])
	}

	cliInput := os.Args[1]

	ctx := context.Background()

	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		log.Fatalln("Environment variable GEMINI_API_KEY not set")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash-exp")
	model.SetTemperature(0.1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text("list of 5 lines with examples as nested list")},
	}

	csvFilePath := "cli_training.csv" // Hardcoded CSV file path
	parts, err := readPartsFromCSV(csvFilePath)
	if err != nil {
		log.Fatalf("Error reading parts from CSV: %v", err)
	}
	// Append CLI input to parts
	parts = append(parts, genai.Text(fmt.Sprintf("input: %s", cliInput)))
	resp, err := model.GenerateContent(ctx, parts...)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Printf("%v\n", part)
	}
	fmt.Printf("token count: %d\n", resp.Candidates[0].TokenCount)
}
