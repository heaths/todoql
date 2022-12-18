package model

import (
	"encoding/json"
)

func (t *Task) UnmarshalJSON(data []byte) error {
	// TODO
	return json.Unmarshal(data, t)
}

func (t *CreateTaskInput) MarshalJSON() ([]byte, error) {
	// TODO
	return json.Marshal(t)
}

func (t *UpdateTaskInput) MarshalJSON() ([]byte, error) {
	// TODO
	return json.Marshal(t)
}
