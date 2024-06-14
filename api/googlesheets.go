package api

import (
	"InfoBot/app/admin"
	"InfoBot/fmtogram/formatter"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func getClient() *sheets.Service {
	ctx := context.Background()
	b, err := os.ReadFile("able-scope-403808-14f43a5ce016.json")
	if err != nil {
		panic(err)
	}
	config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsScope)
	if err != nil {
		panic(err)
	}
	client := config.Client(ctx)
	service, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		panic(err)
	}
	return service
}

func readSheet(service *sheets.Service, spreadsheetId, readRange string) (*sheets.ValueRange, error) {
	return service.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
}

func updateFlag(service *sheets.Service, spreadsheetId, rangeToUpdate string) error {
	value := &sheets.ValueRange{
		Values: [][]interface{}{
			{"1"},
		},
	}
	_, err := service.Spreadsheets.Values.Update(spreadsheetId, rangeToUpdate, value).ValueInputOption("RAW").Do()
	return err
}

// Start of reading Google Sheets
func Start() {
	id := "1qkMMv2iwW7XSbJiJK2MGHkpZvCavJQ4RXz-Sx2iq7wE"
	counter := 2
	service := getClient()
	for {
		rRange := fmt.Sprintf("Ответы на форму (1)!A%d:G", counter)
		data, err := readSheet(service, id, rRange)
		if err == nil {
			for i, row := range data.Values {
				flag := fmt.Sprintf("%v", row[6])
				if flag == "0" || flag == "" {
					admin.FromGoogle(row)
					err := updateFlag(service, id, "Ответы на форму (1)!G"+fmt.Sprintf("%d", i+2))
					if err != nil {
						log.Printf("Unable to update flag: %v", err)
					}
					counter++
				} else {
					fmt.Println("Row already processed, skipping.")
				}
			}
		}
		fmt.Println("Data processed and flags updated")
		time.Sleep(time.Second * 10)
	}
}

func StartTest() (fm *formatter.Formatter) {
	id := "1qkMMv2iwW7XSbJiJK2MGHkpZvCavJQ4RXz-Sx2iq7wE"
	rRange := "Ответы на форму (1)!A2:G"
	service := getClient()
	data, err := readSheet(service, id, rRange)
	if err == nil {
		for i, row := range data.Values {
			flag := fmt.Sprintf("%v", row[6])
			if flag == "0" || flag == "" {
				fm = admin.FromGoogle(row)
				err := updateFlag(service, id, "Ответы на форму (1)!G"+fmt.Sprintf("%d", i+2))
				if err != nil {
					log.Printf("Unable to update flag: %v", err)
				}
			} else {
				fmt.Println("Row already processed, skipping.")
			}
		}
	}
	fmt.Println("Data processed and flags updated")
	return fm
}
