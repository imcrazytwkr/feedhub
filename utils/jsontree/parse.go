package jsontree

func Parse(data []byte) (*Value, error) {
	var v Value
	err := v.UnmarshalJSON(data)
	if err != nil {
		return nil, err
	}

	return &v, nil
}
