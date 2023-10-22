package csvreader

import (
	"encoding/csv"
	"log"
	"log-analysis/config"
	"os"
)

func ReadCSV(conf config.Config) [][]string {
	f := conf.CSVFile
	file, err := os.Open(f)
	if err != nil {
		log.Fatalf("open file error %v", err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalf("csv reading error %v", err)
	}
	return rows
}
