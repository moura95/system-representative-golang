package imports

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"io"
	"strings"

	"github.com/gocarina/gocsv"
)

func unmarshalCsvRequestBody[T any](requestBody []byte, dataSlice *[]T) error {
	// Create a bytes buffer from the request body data
	buf := bytes.NewBuffer(requestBody)

	// Create a new CSV reader using the bytes buffer
	csvReader := csv.NewReader(buf)

	// Parse the CSV data into a slice of string slices
	var csvData [][]string
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		csvData = append(csvData, row)
	}

	// Convert CSV data into a string
	var csvString strings.Builder
	for _, row := range csvData {
		csvString.WriteString(strings.Join(row, ","))
		csvString.WriteRune('\n')
	}

	// Unmarshal the CSV string into the data slice
	if err := gocsv.UnmarshalString(csvString.String(), dataSlice); err != nil {
		return err
	}

	if len(*dataSlice) == 0 {
		return ErrEmptyCsvFile
	}

	return nil
}

func sqlNullString(value string, valid bool) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  valid,
	}
}
