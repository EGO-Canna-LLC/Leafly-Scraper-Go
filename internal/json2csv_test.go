package internal

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestJson2Csv(t *testing.T) {
	// clean up test files when done
	t.Cleanup(func() {
		CleanUp("./ego.html")
		CleanUp("./ego.json")
		CleanUp("./ego.csv")
	})
	err := ConvertJSONToCSV("./ego.json", "./ego.csv")
	if err != nil {
		t.Error(err)
	}
	// does the html csv exist?
	_, err = os.Stat("./ego.csv")
	if os.IsNotExist(err) {
		t.Error("file does not exist")
	}
	csvData, err := readCsvFile("./ego.csv")
	if err != nil {
		t.Error(err)
	}
	if csvData[0][0] != "comment_ID" || csvData[0][1] != "comment_post_ID" || csvData[0][2] != "product_SKU" || csvData[0][3] != "comment_author" || csvData[0][4] != "comment_author_email" || csvData[0][5] != "comment_date" || csvData[0][6] != "comment_date_gmt" || csvData[0][7] != "comment_content" || csvData[0][8] != "comment_approved" || csvData[0][9] != "comment_parent" || csvData[0][10] != "user_id" || csvData[0][11] != "rating" || csvData[0][12] != "verified" {
		t.Error("column header is not correct")
	}
	if csvData[1][3] != "Test" {
		t.Error("author is not correct")
	}
	if csvData[1][7] != "Test" {
		t.Error("review is not correct")
	}
	if csvData[1][11] != "5" {
		t.Error("rating is not correct")
	}
}

func readCsvFile(filename string) ([][]string, error) {
	// Open CSV file
	fileContent, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer fileContent.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(fileContent).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}
