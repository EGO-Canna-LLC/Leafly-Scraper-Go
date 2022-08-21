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

### Coverage
[![Coverage Status](https://img.shields.io/badge/coverage-100%25-brightgreen.svg)](https://coveralls.io/r/github/james-d-james/leafly-scrapper-go)