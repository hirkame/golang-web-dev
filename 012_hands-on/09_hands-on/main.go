package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

type dataRow struct {
	Date time.Time
	Open float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// Open a file.
	file, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Read the file.
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	// Create a empty data structure.
	data := make([]dataRow, 0, len(records))

	// Append data.
	for _, record := range records {
		date, _ := time.Parse("2006-01-02", record[0])
		open, _ := strconv.ParseFloat(record[1], 32)
		data = append(data, dataRow{
			Date: date,
			Open: open,
		})
	}

	err = tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
