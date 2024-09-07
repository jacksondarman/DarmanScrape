package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/option"
)

func main() {
	// Load Credentials key from .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	// Set the environment variable...
	credentialsFilePath := os.Getenv("GOOGLE_CREDENTIALS_FILE_PATH")
	if credentialsFilePath == "" {
		log.Fatalf("GOOGLE_CREDENTIALS_FILE_PATH environment variable isn't set.")
	}

	// Add new context
	ctx := context.Background()

	// Log in to Google Docs API.
	srv, err := docs.NewService(ctx, option.WithCredentialsFile(credentialsFilePath))
	if err != nil {
		log.Fatalf("Unable to create Docs service: %v", err)
	}

	// Fetch Document with this Document ID
	documentId := os.Getenv("DOCUMENT_ID")
	doc, err := srv.Documents.Get(documentId).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve Google Doc: %v", err)
	}

	fmt.Println(extractText(doc.Body.Content))
}

func extractText(elements []*docs.StructuralElement) string {
	var text string
	for _, elem := range elements {
		if elem.Paragraph != nil {
			for _, element := range elem.Paragraph.Elements {
				if element.TextRun != nil {
					text += element.TextRun.Content
				}
			}
		}
	}
	return text
}
