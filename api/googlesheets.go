package api

import (
	"InfoBot/app/admin"
	"InfoBot/fmtogram/formatter"
	"context"
	"fmt"
	"log"
	"os"

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

func sendRequest(id, rowrange string, srv *sheets.Service) {
	data, err := readSheet(srv, id, rowrange)
	log.Print(data)
	if err == nil && len(data.Values) > 0 {
		for i, row := range data.Values {
			if len(row) > 6 {
				flag := fmt.Sprintf("%v", row[6])
				if flag == "0" {
					admin.FromGoogle(row)
					err = updateFlag(srv, id, "Хургада ответы!G"+fmt.Sprintf("%d", i+2))
					if err != nil {
						log.Printf("Unable to update flag: %v", err)
					}
				} else {
					log.Print("Row already processed, skipping")
				}
			}
		}
	} else {
		log.Printf("Error while reading the Sheet: %v", err)
	}
	log.Print("err == nil && len(data.Values) > 0", err == nil && len(data.Values) > 0)
}

// Start of reading Google Sheets
func Start() {
	service := getClient()
	id := "1SzQnlwTT6KGKp9oj8JcQEHA0ZvPwh1I_KbvVB2iaY14"
	rowrange := "Хургада ответы!A2:G"
	sendRequest(id, rowrange, service)

}

func StartTest() (fm *formatter.Formatter) {
	id := "1SzQnlwTT6KGKp9oj8JcQEHA0ZvPwh1I_KbvVB2iaY14"
	rRange := "Хургада ответы!A2:G"
	service := getClient()
	data, err := readSheet(service, id, rRange)
	if err == nil {
		for i, row := range data.Values {
			flag := fmt.Sprintf("%v", row[6])
			if flag == "0" || flag == "" {
				fm = admin.FromGoogle(row)
				err := updateFlag(service, id, "Хургада ответы!G"+fmt.Sprintf("%d", i+2))
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
