package utils

import (
	"encoding/json"
	"strings"
	"time"
)

type JsonParsingDate time.Time

func (j *JsonParsingDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "" || s == "null" {
		return nil
	}
	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}
	*j = JsonParsingDate(t)
	return nil

}

func (j JsonParsingDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}
