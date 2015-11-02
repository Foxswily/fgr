package template

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	DATA_TEMPLATEFILE = "template"
	DATA_OUTPUTFILE   = "out"
)

func Write(mapData map[string]interface{}) error {
	var tList, oList []string
	temp := mapData[DATA_TEMPLATEFILE]
	switch tt := temp.(type) {
	case string:
		tList = append(tList, temp.(string))
	case []string:
		tList = temp.([]string)
	case []interface{}:
		for _, v := range temp.([]interface{}) {
			if strData, ok := v.(string); ok {
				tList = append(tList, strData)
			} else {
				return errors.New("template set error,only string can be recognized")
			}
		}
	default:
		return errors.New(fmt.Sprintf("template set error,unknown data type %T", tt))
	}

	out := mapData[DATA_OUTPUTFILE]
	switch ot := out.(type) {
	case string:
		oList = append(oList, out.(string))
	case []string:
		oList = out.([]string)
	case []interface{}:
		for _, v := range out.([]interface{}) {
			if strData, ok := v.(string); ok {
				oList = append(oList, strData)
			} else {
				return errors.New("outfile set error,only string can be recognized")
			}
		}
	default:
		return errors.New(fmt.Sprintf("outfile set error,unknown data type %T", ot))
	}
	if len(tList) != len(oList) {
		return errors.New("number of template and outfile are not equal")
	}

	for i, tf := range tList {
		t, err := template.ParseFiles(tf)
		if err != nil {
			return err
		}
		path, err := filepath.Abs(oList[i])
		if err != nil {
			return err
		}
		fw, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_APPEND, os.ModeType)
		if err != nil {
			return err
		}
		err = t.Execute(fw, mapData)
		if err != nil {
			return err
		}
	}
	return nil
}

func Read(datafile string) (map[string]interface{}, error) {
	fullpath, err := filepath.Abs(datafile)
	if err != nil {
		return nil, err
	}

	dFile, err := os.Open(fullpath)
	defer dFile.Close()
	if err != nil {
		return nil, err
	}

	fileExt := filepath.Ext(strings.ToLower(datafile))
	switch fileExt {
	case ".yaml":
		return processYAML(dFile)
	default:
		return nil, errors.New("unknown file type " + fileExt)
	}
}

func processYAML(file io.Reader) (map[string]interface{}, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	dataMap := make(map[string]interface{})
	err = yaml.Unmarshal(data, &dataMap)
	if err != nil {
		return nil, err
	}
	return dataMap, nil
}
