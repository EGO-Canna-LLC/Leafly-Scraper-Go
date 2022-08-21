package internal

import (
	"encoding/json"
	"os"
	"testing"
)

func TestHtml2Json(t *testing.T) {
	url := "https://www.egocanna.com/test/reviews.html"
	err := SaveHTML(url, "./ego.html")
	if err != nil {
		t.Error(err)
	}
	// does the html file exist?
	_, err = os.Stat("./ego.html")
	if os.IsNotExist(err) {
		t.Error("file does not exist")
	}
	err = ExtractJSON("./ego.html", "ego.json")
	if err != nil {
		return
	}
	// does the json file exist?
	_, err = os.Stat("./ego.json")
	if os.IsNotExist(err) {
		t.Error("file does not exist")
	}
	_, err = os.Stat("./ego.json.formatted")
	if os.IsNotExist(err) {
		t.Error("file does not exist")
	} else {
		CleanUp("./ego.json")
		os.Rename("./ego.json.formatted", "./ego.json")
	}

	// read the json file
	jsonData, err := readFromFile("./ego.json")
	valid := json.Valid([]byte(jsonData))
	if !valid {
		t.Error("json is not valid")
	}
}
