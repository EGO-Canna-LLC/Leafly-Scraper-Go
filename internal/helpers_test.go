package internal

import (
	"log"
	"os"
	"testing"
)

func TestHelpers(t *testing.T) {
	file, err := os.Create("./test_reviews.html")
	if err != nil {
		t.Error(err)
	}
	_, err = file.WriteString("old falcon")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	htmlData, err := readFromFile("./test_reviews.html")
	if err != nil {
		t.Error(err)
	}
	// what does the file contain?
	if htmlData != "old falcon" {
		t.Errorf("failed to read from file")
	}
	err = CleanUp("./test_reviews.html")
	if err != nil {
		t.Error(err)
	}
	// does file exist?
	_, err = os.Stat("./test_reviews.html")
	if !os.IsNotExist(err) {
		t.Error("file still exists")
	}
}
