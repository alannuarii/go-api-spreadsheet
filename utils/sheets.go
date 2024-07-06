package utils

import (
    "context"
    "os"

    "golang.org/x/oauth2/google"
    "google.golang.org/api/option"
    "google.golang.org/api/sheets/v4"
)

func GetData(spreadsheetId string, readRange string) (string) {
    ctx := context.Background()

    serviceAccountFile := "credential/secret.json"

    b, err := os.ReadFile(serviceAccountFile)
    if err != nil {
        return ""
    }

    config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsReadonlyScope)
    if err != nil {
        return ""
    }

    client := config.Client(ctx)

    srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
    if err != nil {
        return ""
    }

    resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
    if err != nil {
        return ""
    }

    if len(resp.Values) == 0 || len(resp.Values[0]) == 0 {
        return ""
    }

    value, ok := resp.Values[0][0].(string)
    if !ok {
        return ""
    }

    return value
}
