package gpx

type NullableString struct {
	data string
	null bool
}

func (n *NullableString) Null() bool {
	return n.null
}

func (n *NullableString) NotNull() bool {
	return !n.null
}

func (n *NullableString) Value() string {
	return n.data
}

func (n *NullableString) SetValue(data string) {
	n.data = data
}

func NewNullableString(data string) *NullableString {
	result := new(NullableString)
	result.data = data
	return result
}
