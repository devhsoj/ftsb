package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

const (
	TrailStatusUrl string = "https://fredtrails.org/trail-status/"
	UserAgent      string = "fts-client-v1.0"
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

	var summary string

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

	return summary, nil
}
