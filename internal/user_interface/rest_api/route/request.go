package route

import "encoding/json"

type requestBody interface {
	registeringUserStruct
}

func unmarshalToStruct[T requestBody](body []byte) (T, error) {
	var t T
	if err := json.Unmarshal(body, &t); err != nil {
		return t, err
	}
	return t, nil
}
