package channel

import (
	"context"
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

func write(service *sheets.Service, spreadsheetId, rangeToUpdate string) error {
	value := &sheets.ValueRange{
		Values: [][]interface{}{
			{"0"},
		},
	}
	_, err := service.Spreadsheets.Values.Update(spreadsheetId, rangeToUpdate, value).ValueInputOption("RAW").Do()
	return err
}

func rewriteToSheet() {
	srv := getClient()
	err := write(srv, "1qkMMv2iwW7XSbJiJK2MGHkpZvCavJQ4RXz-Sx2iq7wE", "Ответы на форму (1)!G2")
	if err != nil {
		panic(err)
	}
}
