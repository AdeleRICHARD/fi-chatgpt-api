package tool

import (
	"encoding/csv"
	"os"
)

func Export_chatGPT_answer_csv(answer []string) error {
	outputFile, err := os.Create("assets.csv")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"Atouts"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, asset := range answer {
		var csvRow []string
		csvRow = append(csvRow, asset)
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}
