package genai

import (
	"context"
	"errors"
	"os"

	"google.golang.org/genai"
)

func GeminiFetch(bookTitle, author string) (string, error) {
	const prompt = `You're a book synopsis/summary writer, you will check on your available sources \n 
	You will write a synopsis based on content from training data and every source you can fetch from \n 
	Get the content and write a synopsis considering something that would go on a back cover or an e-commerce storefront \n


	This is a full successful example:

	The story centres on Alic \n

	
	Take this as your example and make something shorter for the book you're investigating, around 3-4 lines, 50 words at maximum, the book in question is: 
	`

	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEN_AI_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return "AI Client not initialized, check your environment variables", err
	}

	fullPrompt := prompt + bookTitle
	if author != "" {
		fullPrompt += " by author: " + author
	}

	response, err := client.Models.GenerateContent(ctx, "gemini-2.5-flash", genai.Text(fullPrompt), nil)
	if err != nil {
		return "No response from ai: %e", err
	}

	if len(response.Candidates) > 0 && len(response.Candidates[0].Content.Parts) > 0 {
		part := response.Candidates[0].Content.Parts[0]
		if part.Text != "" {
			return part.Text, nil
		}
		return "", errors.New("unexpected response from ai")
	}

	return "", errors.New("No response from ai")
}
