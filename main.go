package main

import (
	"github.com/kamran0812/xcelldb/persist"
	readconfig "github.com/kamran0812/xcelldb/readConfig"
)

func main() {
	//Read YAML file to get List of COLUMN names that needed to be persisted
	validCols := readconfig.Load()

	//Pass Valid COLUMNS list for Persistence
	persist.Upload(validCols)

}
