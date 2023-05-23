package util

import (
	"bytes"
	"html/template"
)

func ExecuteTemplate(model interface{}, templateStr string) (string, error) {
	buf := []byte{}
	result := bytes.NewBuffer(buf)

	temp := template.New("new-template")

	t, err := temp.Parse(templateStr)
	if err != nil {
		return "", err
	}

	err = t.Execute(result, model)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}
