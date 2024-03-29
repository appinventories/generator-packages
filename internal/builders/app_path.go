package builders

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/appinventories/generator-packages/util"
	"github.com/stoewer/go-strcase"
	"golang.org/x/exp/slices"
)

func AppPath(inputDir string, className string, outputFile string, ignoreFiles []string) error {
	inputDir = strings.Trim(inputDir, "./")
	inputDir = strings.TrimRight(inputDir, "/")
	inputDir += "/"

	appPathTemps := []AppPathTemp{}

	err := filepath.Walk(inputDir,
		func(path string, info os.FileInfo, err error) error {
			if slices.Contains(ignoreFiles, info.Name()) {
				return nil
			}
			if !info.IsDir() {
				tempPath := strings.TrimPrefix(path, inputDir)
				constVariableList := strings.Split(tempPath, ".")
				constVariableList = constVariableList[:len(constVariableList)-1]
				constVariable := strings.Join(constVariableList, ".")
				constVariable = strings.ReplaceAll(constVariable, ".", "_")
				constVariable = strings.ReplaceAll(constVariable, "-", "_")
				constVariable = strings.ReplaceAll(constVariable, " ", "_")
				constVariable = strcase.LowerCamelCase(constVariable)
				constVariable = strings.ReplaceAll(constVariable, "/", "_")
				appPathTemps = append(appPathTemps, AppPathTemp{
					ConstVariable: constVariable,
					Value:         path,
				})
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
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
