package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func ExampleGenerativeModel_GenerateContent() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.0-pro")
	resp, err := model.GenerateContent(ctx, genai.Text("What are the color of rainbow? Please explain"))
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
}

func ExampleGenerativeModel_GenerateContentStream() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.0-pro")

	iter := model.GenerateContentStream(ctx, genai.Text("Tell me a story about a lumberjack and his giant ox. Keep it very short."))
	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		printResponse(resp)
	}
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}

func main() {
	ExampleGenerativeModel_GenerateContent()
	fmt.Println("---****---")
	ExampleGenerativeModel_GenerateContentStream()
}
