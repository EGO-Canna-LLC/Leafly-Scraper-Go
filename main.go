package main

import (
	"github.com/EGO-Canna-LLC/Leafly-Scrapper-Go/internal"
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
	err = internal.ConvertJSONToCSV("./reviews.json", "./reviews.csv")
	if err != nil {
		return
	}
	err = internal.CleanUp("./reviews.html")
	if err != nil {
		return
	}
}
