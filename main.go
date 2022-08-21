package main

import (
	"fmt"
	"github.com/EGO-Canna-LLC/Leafly-Scrapper-Go/internal"
	"os"
)

func main() {
	url := "https://www.leafly.com/brands/ego-canna/products/ego-canna-thco-d9-designer-gummies-10ct-gummies"
	err := internal.SaveHTML(url, "./reviews.html")
	if err != nil {
		return
	}
	// extract <script type="application/ld+json"> from example.html
	// and save to example.json
	err = internal.ExtractJSON("./reviews.html", "reviews.json")
	if err != nil {
		return
	}
	_, err = os.Stat("./reviews.json")
	if os.IsNotExist(err) {
		fmt.Errorf("file does not exist")
	}
	_, err = os.Stat("./reviews.json.formatted")
	if os.IsNotExist(err) {
		fmt.Errorf("file does not exist")
	} else {
		internal.CleanUp("./reviews.json")
		os.Rename("./reviews.json.formatted", "./reviews.json")
	}
	err = internal.ConvertJSONToCSV("./reviews.json", "./reviews.csv")
	if err != nil {
		return
	}
	err = internal.CleanUp("./reviews.html")
	if err != nil {
		return
	}
}
