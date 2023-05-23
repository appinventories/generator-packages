package builders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"github.com/appinventories/generator-packages/util"
	"github.com/stoewer/go-strcase"
)

func GenerateLocaleKey(inputPath string, outputDir string, nameSpace string) error {

	byteValue, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return err
	}
	var result map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		return err
	}
	resultKeys := []string{}
	getKeyString(result, "", &resultKeys)

	sort.Strings(resultKeys)
	localeKeyTemps := []*LocaleKeyTemp{}
	for _, resultKey := range resultKeys {
		localeKeyTemps = append(localeKeyTemps, &LocaleKeyTemp{
			Key:   strcase.SnakeCase(strings.ReplaceAll(resultKey, ".", "_")),
			Value: resultKey,
		})
	}

	resultStr, err := util.ExecuteTemplate(localeKeyTemps, fmt.Sprintf(localeKeyTemplateStr, nameSpace, nameSpace))
	if err != nil {
		return err
	}

	outputFile := filepath.Join(outputDir, strcase.SnakeCase(nameSpace)+"_gen.ts")

	err = ioutil.WriteFile(outputFile, []byte(resultStr), 0755)
	if err != nil {
		return err
	}

	return nil

}

func getKeyString(jsonKey interface{}, currentKey string, resultKey *[]string) {
	switch jsonKey.(type) {
	case map[string]interface{}:
		for key, value := range jsonKey.(map[string]interface{}) {
			if currentKey != "" {
				key = currentKey + "." + key
			}
			switch value.(type) {
			case map[string]interface{}:
				getKeyString(value, key, resultKey)
			default:
				*resultKey = append(*resultKey, key)
			}
		}
	default:
		log.Panicf("err json")
	}

}

type LocaleKeyTemp struct {
	Key   string
	Value string
}

const localeKeyTemplateStr = `
enum %s {
{{ range . }}	{{.Key}} = "{{.Value}}",
{{ end }}}

export default %s
`
