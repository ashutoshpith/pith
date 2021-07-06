package xlsx

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)
type XlsxStr struct{
	Text string
}

type Sheet struct {
	Sheet string
	Axis string
}

func openXlsxFile(text string ) *excelize.File {
	file, err := excelize.OpenFile(text)
	if err != nil {
		fmt.Println(nil)
	}
	return file
}
func (xl XlsxStr) Count() {
	f := openXlsxFile(xl.Text)
	count := f.SheetCount
	fmt.Println("SheetCount ", count)
}

func (xl XlsxStr) ViewSheet(sheet Sheet) {
     file := openXlsxFile(xl.Text)
	 cell, err := file.GetCellValue(sheet.Sheet, sheet.Axis)
	 if err != nil {
		 fmt.Println(err)
	 }
	 fmt.Println("Cell ",cell)

	 rows, err := file.GetRows(sheet.Sheet)
	 if err != nil {
		 fmt.Println(err)
	 }
	 for _, row := range rows {
		 for _, col := range row {
			 fmt.Print(col, "\t")
		 }
		 fmt.Println()
	 }

} 