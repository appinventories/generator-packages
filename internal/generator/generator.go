package generator

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/appinventories/generator-packages/internal/builders"
	"github.com/appinventories/generator-packages/internal/parsers"
	"gopkg.in/yaml.v2"
)

func GenerateLocaleKey(cfgPath string) error {
	conf, err := getConf(cfgPath)
	if err != nil {
		return err
	}
	for _, genLocaleKeyConfig := range conf.GenLocaleKeyConfigs {
		jsonFilePaths, err := getJsonFilePath(genLocaleKeyConfig.InputDir)
		if err != nil {
			return err
		}

		err = builders.GenerateLocaleKey(jsonFilePaths[0], genLocaleKeyConfig.OutputDir, genLocaleKeyConfig.Namespace)
		if err != nil {
			return err
		}
	}
	return nil

}

func CsvToJson(cfgPath string) error {
	conf, err := getConf(cfgPath)
	if err != nil {
		return err
	}

	for _, csvToJsonConfig := range conf.CsvToJsonConfigs {
		err = builders.CsvToJson(csvToJsonConfig.InputPath, csvToJsonConfig.OutputDir, csvToJsonConfig.Key)
		if err != nil {
			return err
		}
	}

	return nil
}

func AppPath(cfgPath string) error {
	conf, err := getConf(cfgPath)
	if err != nil {
		return err
	}
	for _, appPathConfig := range conf.AppPathConfigs {
		err = builders.AppPath(appPathConfig.InputDir, appPathConfig.ClassName, appPathConfig.OutputFile)
	}
	if err != nil {
		return err
	}
	return nil
}

func getConf(filePath string) (*parsers.Config, error) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panicf("Unable to read config file: '%s'\n#%v ", filePath, err)
	}

	conf := &parsers.Config{}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func getJsonFilePath(jsonDir string) ([]string, error) {
	jsonFilePaths := []string{}
	files, err := ioutil.ReadDir(jsonDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		filePath := filepath.Join(jsonDir, file.Name())
		if filepath.Ext(filePath) == ".json" {
			jsonFilePaths = append(jsonFilePaths, filePath)
		}

	}
	return jsonFilePaths, nil
}
