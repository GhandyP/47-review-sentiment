package main

import (
	"context"
	"log"
	"os"

	"github.com/GhandyP/adk-go-enterprise-cookbook/internal/adkhelper"
)

func main() {
	if err := adkhelper.Run(
		context.Background(),
		os.Args[1:],
		"case_47_review_sentiment",
		"A text-only starting point for review sentiment.",
		"Analyze review text supplied by the user for sentiment and aspects. Quote evidence, preserve review IDs if provided, and avoid inferring facts not stated.",
	); err != nil {
		log.Fatal(err)
	}
}
