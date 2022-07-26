package diabetes

func (u *GlucoseUnit) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	*u, err = InterpretGlucoseUnit(string(data))
	return err
}
