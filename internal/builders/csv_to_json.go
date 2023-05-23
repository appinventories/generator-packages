package builders

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CsvToJson(inputPath string, outputDir string, key *string) error {

	csvFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	langs := []string{}

	// header : key,lang1,lang2,.....
	header := csvLines[0]

	for i := range header {
		// ignore first index of header
		if i == 0 {
			continue
		}
		langs = append(langs, csvLines[0][i])

	}

	langsMap := []map[string]interface{}{}
	for range langs {
		langsMap = append(langsMap, map[string]interface{}{})
	}

	for i, lines := range csvLines {
		// ignore header
		if i == 0 {
			continue
		}

		for j := range lines {
			// ignore key
			if j == 0 {
				continue
			}
			langsMap[j-1][lines[0]] = lines[j]

		}
	}

	for i, lang := range langs {
		outputJson := langsMap[i]
		if key != nil {
			outputJson = map[string]interface{}{*key: langsMap[i]}
		}

		content, err := json.MarshalIndent(outputJson, "", "	")
		if err != nil {
			return err
		}

		outputFile := filepath.Join(outputDir, lang+".json")

		err = ioutil.WriteFile(outputFile, content, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

type JsonLang struct {
	T map[string]string
}
