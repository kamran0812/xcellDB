package persist

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/xuri/excelize/v2"
)

func Upload(validCols map[string]int) {

	type res struct{}
	mapstructure.Decode(validCols, res)
	fmt.Println(&res)
	//Open excel file
	f, err := excelize.OpenFile("./persist/test.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	//  defer close the excel.
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	//get the Column wise iterator
	cols, err := f.Cols("Sheet1")
	if err != nil {
		panic(err)
	}

	//Iterate col wise
	for cols.Next() {
		//get rows of current column
		row, err := cols.Rows()
		if err != nil {
			fmt.Println(err)
		}
		//		colName := row[0]

		//check if cols are present in list then proceed
		if _, ok := validCols[row[0]]; ok {
			for idx, data := range row {
				if idx == 0 || data == "" {
					continue
				}
				_, err := Client.DB()
				if err != nil {
					panic(err)
				}
				//				fmt.Println("Writing in DB: ", data, "ColName: ", colName)
			}
		}

	}

}
