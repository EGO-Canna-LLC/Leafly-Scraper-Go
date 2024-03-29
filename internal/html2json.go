package internal

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

// ExtractJSON extracts the JSON from a given HTML file. Then, it saves the JSON to a file and passed the filename to formatJSON.
// for prettifying
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
	// Find reviews in JSONLD data
	scriptOpen := "<script type=\"application/ld+json\">"
	scriptClose := "</script>"
	rx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(scriptOpen) + `(.*?)` + regexp.QuoteMeta(scriptClose))
	matches := rx.FindAllStringSubmatch(htmlData, 1)
	if len(matches) <= 0 {
		return fmt.Errorf("no JSON found in %s", filename)
	}
	jsonldData := matches[0][1]

	// find review author
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
	// find review comments
	reviewBodyOpen := "\"reviewBody\":"
	reviewBodyClose := ","
	reviewBodyRx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(reviewBodyOpen) + `(.*?)` + regexp.QuoteMeta(reviewBodyClose))
	reviewBodyMatches := reviewBodyRx.FindAllStringSubmatch(jsonldData, -1)

	var reviewBodies []string
	for _, review := range reviewBodyMatches {
		reviewBodies = append(reviewBodies, strings.Trim(review[1], "\""))
	}
	// find review ratings
	reviewRatingOpen := "\"reviewRating\":{\"@type\":\"Rating\",\"ratingValue\":"
	reviewRatingClose := "}"
	reviewRatingRx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(reviewRatingOpen) + `(.*?)` + regexp.QuoteMeta(reviewRatingClose))
	reviewRatingMatches := reviewRatingRx.FindAllStringSubmatch(jsonldData, -1)

	var reviewRatings []string
	for _, review := range reviewRatingMatches {
		reviewRatings = append(reviewRatings, strings.Trim(review[1], "\""))
	}

	// open json file and write json data to it
	jsonFile, _ := os.Create(jsonFilename)
	defer jsonFile.Close()

	jfStart := []byte(`[{`)
	jsonFile.Write(jfStart)
	var reviewsLeft = len(reviewAuthors)
	// Write reviews to file
	if len(reviewAuthors) == len(reviewBodies) && len(reviewAuthors) == len(reviewRatings) {
		for i, review := range reviewAuthors {
			for j, reviewBody := range reviewBodies {
				// reviewBody = strings.Replace(reviewBody, " ", "\\u0020", -1) // escape spaces in json with \u0020
				for k, reviewRating := range reviewRatings {
					if i == j && i == k {
						if reviewsLeft <= 1 {
							// last review, don't add comma
							jsonFile.Write([]byte(`"Author":"` + review + `",` + `"Review":"` + reviewBody + `","Rating":` + reviewRating))
						} else {
							// not last review, add comma
							jsonFile.Write([]byte(`"Author":"` + review + `",` + `"Review":"` + reviewBody + `","Rating":` + reviewRating + `,`))
						}
						reviewsLeft--
					}
				}
			}
		}
	}
	jsonFile.Write([]byte("}]"))
	return formatJSON(jsonFilename)
}

func formatJSON(filename string) (err error) {
	fmt.Println("Formatting JSON")
	incFile, _ := os.ReadFile(filename)
	var prettyJSON bytes.Buffer
	// prettify the json before writing it to the file
	error := json.Indent(&prettyJSON, incFile, "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}
	jsonFile, _ := os.Create(filename + ".formatted")
	defer jsonFile.Close()
	jsonFile.Write(prettyJSON.Bytes())
	return nil
}
