package data

import (
	"fmt"
	"strings"
	"strconv"
	"go-api-spreadsheet/utils"
)

var tanggal = utils.YesterdayDate()
var day = strings.Split(tanggal, "-")
var sheetInt, _ = strconv.Atoi(day[2])
var sheet = strconv.Itoa(sheetInt)

var Unit1 = []string{fmt.Sprintf("'%s'!F7", sheet), fmt.Sprintf("'%s'!G100", sheet), fmt.Sprintf("'%s'!B7", sheet), fmt.Sprintf("'%s'!C100", sheet), fmt.Sprintf("'%s'!O100", sheet), fmt.Sprintf("'%s'!M38", sheet), fmt.Sprintf("'%s'!M69", sheet), fmt.Sprintf("'%s'!E111", sheet)}
var Unit3 = []string{fmt.Sprintf("'%s'!F8", sheet), fmt.Sprintf("'%s'!G101", sheet), fmt.Sprintf("'%s'!B8", sheet), fmt.Sprintf("'%s'!C101", sheet), fmt.Sprintf("'%s'!O101", sheet), fmt.Sprintf("'%s'!M39", sheet), fmt.Sprintf("'%s'!M70", sheet), fmt.Sprintf("'%s'!E112", sheet)}
var Unit4 = []string{fmt.Sprintf("'%s'!F9", sheet), fmt.Sprintf("'%s'!G102", sheet), fmt.Sprintf("'%s'!B9", sheet), fmt.Sprintf("'%s'!C102", sheet), fmt.Sprintf("'%s'!O102", sheet), fmt.Sprintf("'%s'!M40", sheet), fmt.Sprintf("'%s'!M71", sheet), fmt.Sprintf("'%s'!E113", sheet)}
var Unit5 = []string{fmt.Sprintf("'%s'!F10", sheet), fmt.Sprintf("'%s'!G103", sheet), fmt.Sprintf("'%s'!B10", sheet), fmt.Sprintf("'%s'!C103", sheet), fmt.Sprintf("'%s'!O103", sheet), fmt.Sprintf("'%s'!M41", sheet), fmt.Sprintf("'%s'!M72", sheet), fmt.Sprintf("'%s'!E114", sheet)}
var Unit6 = []string{fmt.Sprintf("'%s'!F11", sheet), fmt.Sprintf("'%s'!G104", sheet), fmt.Sprintf("'%s'!B11", sheet), fmt.Sprintf("'%s'!C104", sheet), fmt.Sprintf("'%s'!O104", sheet), fmt.Sprintf("'%s'!M42", sheet), fmt.Sprintf("'%s'!M73", sheet), fmt.Sprintf("'%s'!E115", sheet)}
var Unit7 = []string{fmt.Sprintf("'%s'!F12", sheet), fmt.Sprintf("'%s'!G105", sheet), fmt.Sprintf("'%s'!B12", sheet), fmt.Sprintf("'%s'!C105", sheet), fmt.Sprintf("'%s'!O105", sheet), fmt.Sprintf("'%s'!M43", sheet), fmt.Sprintf("'%s'!M74", sheet), fmt.Sprintf("'%s'!E116", sheet)}
var Unit8 = []string{fmt.Sprintf("'%s'!F13", sheet), fmt.Sprintf("'%s'!G106", sheet), fmt.Sprintf("'%s'!B13", sheet), fmt.Sprintf("'%s'!C106", sheet), fmt.Sprintf("'%s'!O106", sheet), fmt.Sprintf("'%s'!M44", sheet), fmt.Sprintf("'%s'!M75", sheet), fmt.Sprintf("'%s'!E117", sheet)}
var Unit9 = []string{fmt.Sprintf("'%s'!F14", sheet), fmt.Sprintf("'%s'!G107", sheet), fmt.Sprintf("'%s'!B14", sheet), fmt.Sprintf("'%s'!C107", sheet), fmt.Sprintf("'%s'!O107", sheet), fmt.Sprintf("'%s'!M45", sheet), fmt.Sprintf("'%s'!M76", sheet), fmt.Sprintf("'%s'!E118", sheet)}
