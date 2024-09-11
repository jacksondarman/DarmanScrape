package extractor

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/option"
)

type GoogleDoc struct {
	Srv        *docs.Service
	DocumentId string
}

/*
*
Creates a new GoogleDoc instance by loading environment variables, setting up credentials, and initializing a Docs service.
Returns the initialized GoogleDoc instance or an error if any step fails.
*/
func NewGoogleDoc() (*GoogleDoc, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading the .env file: %w", err)
	}

	// Set environment variables.
	credentialsFilePath := os.Getenv("GOOGLE_CREDENTIALS_FILE_PATH")
	if credentialsFilePath == "" {
		return nil, fmt.Errorf("GOOGLE_CREDENTIALS_FILE_PATH environment variable isn't set")
	}

	ctx := context.Background()

	srv, err := docs.NewService(ctx, option.WithCredentialsFile(credentialsFilePath))
	if err != nil {
		return nil, fmt.Errorf("unable to create Docs service")
	}

	documentId := os.Getenv("DOCUMENT_ID")
	if documentId == "" {
		return nil, fmt.Errorf("DOCUMENT_ID environment variable isn't set")
	}

	return &GoogleDoc{Srv: srv, DocumentId: documentId}, nil
}

/**
 * GetDocumentText retrieves the text content of the Master List of Civilizations.
 *
 * Returns:
 * - string: The text content of the Master List.
 * - error: An error if unable to retrieve the Master List.
 */
func (g *GoogleDoc) GetDocumentText() (string, error) {
	doc, err := g.Srv.Documents.Get(g.DocumentId).Do()
	if err != nil {
		return "", fmt.Errorf("unable to retrieve Google Doc: %w", err)
	}

	return extractText(doc.Body.Content), nil
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
