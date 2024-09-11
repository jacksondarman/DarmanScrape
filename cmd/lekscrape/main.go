package main

import (
	"log"
	"os"

	"github.com/jacksondarman/lekscrape/internal/extractor"
	"github.com/jacksondarman/lekscrape/internal/parser"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	log.Println("Connecting to Google Docs...")
	doc, err := extractor.NewGoogleDoc()
	if err != nil {
		log.Fatalf("Initialization failed: %v", err)
	}

	log.Println("Extracting text...")
	text, err := doc.GetDocumentText()
	if err != nil {
		log.Fatalf("Failed to extract document text: %v", err)
	}

	log.Println("Parsing Civ data...")
	parser.ParseCivilizations(text)

	app := pocketbase.New()
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start PocketBase: %v", err)
	}

}
