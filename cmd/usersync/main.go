package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var (
	apiKey        = flag.String("apikey", "", "api key from https://console.cloud.google.com/apis/credentials")
	spreadsheetId = flag.String("id", "", "spreadsheet id")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	sheetsService, err := sheets.NewService(ctx, option.WithAPIKey(*apiKey))
	if err != nil {
		log.Fatal(err)
	}
	s, err := sheetsService.Spreadsheets.Get(*spreadsheetId).Do()
	if err != nil {
		log.Fatal(err)
	}
	b, err := s.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
