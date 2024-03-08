package types

import "encoding/json"

type Questions []Question

func UnmarshalQuestions(data []byte) (Questions, error) {
	var r Questions
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Questions) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Question struct {
	Question string `json:"question"`
}
