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
	// Initilize the PocketBase instance
	app := pocketbase.New()

	// Before the PocketBase server is started
	app.OnBeforeBootstrap().Add(func(e *core.BootstrapEvent) error {
		// Initialize connection to the Google Docs API.
		log.Println("Connecting to Google Docs...")
		doc, err := extractor.NewGoogleDoc()
		if err != nil {
			log.Fatalf("Initialization failed: %v", err)
		}

		// Extract text from the Google Doc
		log.Println("Extracting text...")
		text, err := doc.GetDocumentText()
		if err != nil {
			log.Fatalf("Failed to extract document text: %v", err)
		}

		// Pass it to be parsed in the parser package.
		log.Println("Parsing Civ data...")
		parser.ParseCivilizations(text)

		return err
	})

	// Before the internal router is served:
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// Generate a new API endpoint, where the DB files will be accessible.
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	// If any errors show up during the init process, kill the app and display the error!
	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start PocketBase: %v", err)
	}

}
