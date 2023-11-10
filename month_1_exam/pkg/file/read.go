package file

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spf13/cast"
)

func Read(fileName string) (map[string]interface{}, error) {

	var (
		objects   []interface{}
		objectMap = make(map[string]interface{})
	)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &objects)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, object := range objects {
		obj := cast.ToStringMap(object)
		objectMap[cast.ToString(obj["id"])] = object
	}

	return objectMap, nil
}
