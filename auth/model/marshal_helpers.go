package model

import (
	"database/sql"
	"encoding/json"
)

// NullString is a helper type for
// marshalling and unmarshalling sql.NullString values
// this allows us to avoid an extra step of creating
// application and data later models
type NullString struct {
	sql.NullString
}

// MarshalJSON handles possible null sql.NullString
func (s NullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}

	return json.Marshal(nil)

}

// UnmarshalJSON handles possible null sql.NullString
func (s *NullString) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		s.Valid = true
		s.String = *x
	} else {
		s.Valid = false
	}
	return nil
}
