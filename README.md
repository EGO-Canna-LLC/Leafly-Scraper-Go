[![Coverage Status](https://img.shields.io/badge/coverage-88%25-brightgreen.svg)](#)
[![Go](https://github.com/EGO-Canna-LLC/Leafly-Scraper-Go/actions/workflows/go.yml/badge.svg)](https://github.com/EGO-Canna-LLC/Leafly-Scraper-Go/actions/workflows/go.yml)

# Leafly-Scrapper-Go
Leafly Review Scrapper

Will download EGO Canna Product reviews for https://www.leafly.com/brands/ego-canna/products/ego-canna-thco-d9-designer-gummies-10ct-gummie, parse them into JSON and CSV for importing into egocanna.com

## Requirements
go 1.19+

## Running
go run .
or run the binary after compiling

## Compiling
go build .

## Testing
go test ./... -v -cover
