package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func CsvReader(str string) {
	reader := csv.NewReader(strings.NewReader(str))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
