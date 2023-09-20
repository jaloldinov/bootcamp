package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Write(fileName string, objectMap map[string]interface{}) error {
	var objects []interface{}
	for _, val := range objectMap {
		objects = append(objects, val)
	}
	body, err := json.MarshalIndent(objects, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
