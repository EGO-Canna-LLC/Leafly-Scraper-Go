package internal

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Reviews is a struct that holds the data for each review
type Reviews struct {
	Author string `json:"Author"`
	Review string `json:"Review"`
	Rating int    `json:"Rating"`
}

// ConvertJSONToCSV converts JSON data to CSV data
func ConvertJSONToCSV(jsonFile, csvFile string) error {
	fmt.Println("Extracting JSON from ", jsonFile, " to ", csvFile)
	sourceFile, err := os.ReadFile(jsonFile)
	if err != nil {
		return err
	}

	var reviews []Reviews
	err = json.Unmarshal(sourceFile, &reviews)
	if err != nil {
		fmt.Println(err)
	}
	outputFile, err := os.Create(csvFile)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Write the header of the CSV file and the successive rows by iterating through the JSON struct array
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"comment_ID", "comment_post_ID", "product_SKU", "comment_author", "comment_author_email", "comment_date", "comment_date_gmt", "comment_content", "comment_approved", "comment_parent", "user_id", "rating", "verified"}
	if err := writer.Write(header); err != nil {
		return err
	}

	currentTime := time.Now()
	currUTC := currentTime.UTC().String()
	currUTC = strings.Replace(currUTC, " +0000 UTC", "", -1)

	// Write the JSON data to the CSV file
	for _, review := range reviews {
		var record []string
		record = append(record, "")
		record = append(record, "2665")
		record = append(record, "681565140761")
		record = append(record, review.Author)
		record = append(record, "review@leafly.com")
		record = append(record, currentTime.Format("2006-01-02"))
		record = append(record, currUTC)
		record = append(record, review.Review)
		record = append(record, "1")
		record = append(record, "")
		record = append(record, "28")
		record = append(record, strconv.Itoa(review.Rating))
		record = append(record, "1")
		writer.Write(record)
	}
	return nil
}
