package main

import "letsgo/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
