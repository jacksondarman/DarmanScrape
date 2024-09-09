package main

import (
	"log"
)

func main() {
	log.Println("Connecting to Google Docs...")
	doc, err := NewGoogleDoc()
	if err != nil {
		log.Fatalf("Initialization failed: %v", err)
	}

	log.Println("Extracting text...")
	text, err := doc.GetDocumentText()
	if err != nil {
		log.Fatalf("Failed to extract document text: %v", err)
	}

	log.Println("Parsing Civ data...")
	parseCivilizations(text)
}
