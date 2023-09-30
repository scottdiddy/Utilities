package utilities

import "encoding/json"

func UnmarshalToType[T any](original *T, incoming any) error {
	var (
		incoming_byte []byte
		err         error
	)

	switch data := incoming.(type) {
	case string:
		incoming_byte = []byte(data)
	case []byte:
		incoming_byte = data
	default:
		incoming_byte, err = json.Marshal(incoming)
		if err != nil {
			return err
		}
	}
	err = json.Unmarshal(incoming_byte, original)
	if err != nil {
		return err
	}
	return nil
}