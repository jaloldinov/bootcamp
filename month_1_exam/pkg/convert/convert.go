package convert

import "encoding/json"

func ConvertStructToStruct(p interface{}, s interface{}) error {

	body, err := json.Marshal(p)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &s)
	if err != nil {
		return err
	}

	return nil
}
