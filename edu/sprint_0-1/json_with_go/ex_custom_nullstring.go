type NullString struct {
	String string
	Valid  bool
}

func (ns *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ns.Valid = false
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	ns.String = str
	ns.Valid = true
	return nil
}