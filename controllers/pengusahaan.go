package controllers

import (
	"fmt"
	"go-api-spreadsheet/utils"
	"go-api-spreadsheet/db"
	"go-api-spreadsheet/data"
	"go-api-spreadsheet/models"
	"github.com/gin-gonic/gin"
)

var con = db.DB

func PostData(c *gin.Context){

	spreadsheetId := getSpreadSheetId()

	tanggal := utils.YesterdayDate()

	units := [][]string{data.Unit1, data.Unit3, data.Unit4, data.Unit5, data.Unit6, data.Unit7, data.Unit8, data.Unit9}
	unit := []string{"1","3","4","5","6","7","8","9"}

	for i, val := range units {
		postPengusahaan(c, spreadsheetId, unit[i], tanggal, val)
	}

	c.JSON(200, gin.H{"message": fmt.Sprintf("Berhasil menginput data %s", tanggal)})
}



func postPengusahaan(c *gin.Context, spreadsheetId string, unit string, tanggal string, sheetUnit []string){

	var result []float64

	for _, val := range sheetUnit {
		result = append(result, utils.StringToDecimal(utils.GetData(spreadsheetId, val)))
	}

	var sh float64
	jamKerja := result[1] - result[0]
	if jamKerja > 24 {
		sh = 24
	}
	produksi := (result[3] - result[2]) * 500
	bbm := result[4]
	lumas := result[5] + result[6]
	ps := result[7]

	query := `INSERT INTO pengusahaan (tanggal, unit, sh, produksi, bbm, lumas, ps) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	var id int
	err := con.QueryRow(query, tanggal, unit, sh, produksi, bbm, lumas, ps).Scan(&id)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}
}


func getSpreadSheetId() string{
	
	yearMonth := utils.YearMonth()

	spreadsheet := []models.Spreadsheet{}

	query := `SELECT * FROM spreadsheet WHERE about = $1 && periode = $2`

	err := con.Select(&spreadsheet, query, "pengusahaan", yearMonth)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return spreadsheet[0].SpreadsheetId
}


	// jamKerjaAwal := utils.StringToDecimal(utils.GetData(spreadsheetId, fmt.Sprintf("'%s'!F11", sheet[2])))
	// jamKerjaAkhir  := utils.StringToDecimal(utils.GetData(spreadsheetId, fmt.Sprintf("'%s'!G104", sheet[2])))
	// prodAwal := utils.StringToDecimal(utils.GetData(spreadsheetId, fmt.Sprintf("'%s'!B11", sheet[2])))
	// prodAkhir := utils.StringToDecimal(utils.GetData(spreadsheetId, fmt.Sprintf("'%s'!C104", sheet[2])))
	// pemakaianBBM := utils.StringToDecimal(utils.GetData(spreadsheetId, fmt.Sprintf("'%s'!O104", sheet[2])))
	// lumasB1 := utils.StringToDecimal(utils.GetData(spreadsheetId, fmt.Sprintf("'%s'!M42", sheet[2])))
	// lumasB2 := utils.StringToDecimal(utils.GetData(spreadsheetId, fmt.Sprintf("'%s'!M73", sheet[2])))
	// ps := utils.StringToDecimal(utils.GetData(spreadsheetId, fmt.Sprintf("'%s'!E115", sheet[2])))
