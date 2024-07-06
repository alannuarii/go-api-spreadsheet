package models

type Spreadsheet struct {
    ID          		int    `db:"id"`
    SpreadsheetId     	string `db:"spreadsheet_id"`
    About       		string `db:"about"`
    Periode   			string `db:"periode"`
}