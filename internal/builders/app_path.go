package builders

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/appinventories/generator-packages/util"
)

func AppPath(inputDir string, className string, outputFile string) error {
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	appPathTemps := []AppPathTemp{}
	for _, file := range files {
		constVariable := strings.Split(file.Name(), ".")[0]
		value := filepath.Join(inputDir, file.Name())
		appPathTemps = append(appPathTemps, AppPathTemp{
			ConstVariable: constVariable,
			Value:         value,
		})
	}

	resultStr, err := util.ExecuteTemplate(appPathTemps, fmt.Sprintf(appPathTemplateStr, className))
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputFile, []byte(resultStr), 0755)
	if err != nil {
		return err
	}
	return nil
}

type AppPathTemp struct {
	ConstVariable string
	Value         string
}

const appPathTemplateStr = `class %s {
{{ range . }}	static const {{.ConstVariable}} = "{{.Value}}";
{{ end }}}
`
