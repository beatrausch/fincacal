package csv

import (
	"encoding/csv"
	"os"
)

func WriteCSV(file string, table [][]string) error {
	fd, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	w := csv.NewWriter(fd)
	w.WriteAll(table) // calls Flush internally

	if err := w.Error(); err != nil {
		return err
	}
	return nil
}
