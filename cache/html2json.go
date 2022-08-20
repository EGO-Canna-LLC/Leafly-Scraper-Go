package cache

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// SaveHTML saves the HTML of a given URL to a file.
func SaveHTML(url, filename string) (err error) {
	fmt.Println("Downloading ", url, " to ", filename)

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return
}

func ExtractJSON(filename, jsonFilename string) (err error) {
	fmt.Println("Extracting JSON from ", filename, " to ", jsonFilename)

	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	htmlData, err := readFromFile(filename)
	if err != nil {
		return
	}

	scriptOpen := "<script type=\"application/ld+json\">"
	scriptClose := "</script>"
	rx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(scriptOpen) + `(.*?)` + regexp.QuoteMeta(scriptClose))
	matches := rx.FindAllStringSubmatch(htmlData, 1)
	if len(matches) <= 0 {
		return fmt.Errorf("no JSON found in %s", filename)
	}
	jsonldData := matches[0][1]

	// Find reviews in JSONLD data
	reviewOpen := "\"Review\",\"author\":{\"@type\":\"Person\",\"name\":"
	reviewClose := "},"
	reviewRx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(reviewOpen) + `(.*?)` + regexp.QuoteMeta(reviewClose))
	reviewMatches := reviewRx.FindAllStringSubmatch(jsonldData, -1)
	if len(reviewMatches) <= 0 {
		return fmt.Errorf("no reviews found in %s", filename)
	}
	var reviewAuthors []string
	for _, review := range reviewMatches {
		reviewAuthors = append(reviewAuthors, strings.Trim(review[1], "\""))
	}
	reviewBodyOpen := "\"reviewBody\":"
	reviewBodyClose := ","
	reviewBodyRx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(reviewBodyOpen) + `(.*?)` + regexp.QuoteMeta(reviewBodyClose))
	reviewBodyMatches := reviewBodyRx.FindAllStringSubmatch(jsonldData, -1)
	var reviewBodies []string
	for _, review := range reviewBodyMatches {
		reviewBodies = append(reviewBodies, strings.Trim(review[1], "\""))
	}
	reviewRatingOpen := "\"reviewRating\":{\"@type\":\"Rating\",\"ratingValue\":"
	reviewRatingClose := "}"
	reviewRatingRx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(reviewRatingOpen) + `(.*?)` + regexp.QuoteMeta(reviewRatingClose))
	reviewRatingMatches := reviewRatingRx.FindAllStringSubmatch(jsonldData, -1)
	var reviewRatings []string
	for _, review := range reviewRatingMatches {
		reviewRatings = append(reviewRatings, strings.Trim(review[1], "\""))
	}
	jsonFile, _ := os.Create(jsonFilename)
	defer jsonFile.Close()
	jfStart := []byte(`{"Reviews":` + "{")
	jsonFile.Write(jfStart)
	var reviewsLeft = len(reviewAuthors)
	if len(reviewAuthors) == len(reviewBodies) && len(reviewAuthors) == len(reviewRatings) {
		for i, review := range reviewAuthors {
			for j, reviewBody := range reviewBodies {
				for k, reviewRating := range reviewRatings {
					if i == j && i == k {
						if reviewsLeft <= 1 {
							jsonFile.Write([]byte(`"` + review + `":` + `{"body":"` + reviewBody + `","rating":"` + reviewRating + `"}`))
						} else {
							jsonFile.Write([]byte(`"` + review + `":` + `{"body":"` + reviewBody + `","rating":"` + reviewRating + `"}` + ","))
						}
						reviewsLeft--
					}
				}
			}
		}
	}
	jsonFile.Write([]byte("}" + "}"))
	return formatJSON(jsonFilename)
}

func formatJSON(filename string) (err error) {
	fmt.Println("Formatting JSON")
	incJsonFile, _ := os.OpenFile(filename, os.O_RDWR, 0644)
	defer incJsonFile.Close()
	jsonData, _ := io.ReadAll(incJsonFile)
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, jsonData, "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}
	jsonFile, _ := os.Create("example.json")
	defer jsonFile.Close()
	jsonFile.Write(prettyJSON.Bytes())
	return nil
}
