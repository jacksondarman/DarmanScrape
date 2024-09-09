package main

import "google.golang.org/api/docs/v1"

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
