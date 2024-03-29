package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

const (
	TrailStatusUrl string = "https://fredtrails.org/trail-status/"
	UserAgent      string = "ftsb-client"
)

func GetTrailStatusSummary() (string, error) {
	req, err := http.NewRequest("GET", TrailStatusUrl, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", UserAgent)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Printf("failed to properly close body reader: %s", err)
		}
	}()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return "", err
	}

	summary := fmt.Sprintf("# Fredericksburg Trail Status\n[View Website](<%s>)\n```ansi\n", TrailStatusUrl)

	headers := doc.Find("h1, h3, h5")

	headerMatches := []string{
		"trail code",
		"green",
		"amber",
		"red",
		"caution",
	}

	headerMismatches := []string{
		"e-mail",
	}

	headers.Each(func(i int, selection *goquery.Selection) {
		matches := false

		for _, match := range headerMatches {
			if strings.Contains(strings.ToLower(selection.Text()), match) {
				matches = true
				break
			}
		}

		for _, match := range headerMismatches {
			if strings.Contains(strings.ToLower(selection.Text()), match) {
				matches = false
				break
			}
		}

		if !matches {
			return
		}

		summary += selection.Text() + "\n"
	})

	summary = strings.ReplaceAll(summary, "Red", "\u001B[2;31mRed\u001B[0;31m\u001B[0m\u001B[2;31m\u001B[0m")
	summary = strings.ReplaceAll(summary, "Amber", "\u001B[2;33mAmber\u001B[0;33m\u001B[0m\u001B[2;33m\u001B[0m")
	summary = strings.ReplaceAll(summary, "Green", "\u001B[2;32mGreen\u001B[0;32m\u001B[0m\u001B[2;32m\u001B[0m")

	return summary + "\n```", nil
}
