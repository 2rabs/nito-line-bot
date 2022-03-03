package member

type LineUserId struct {
	Value string
}

func NewLineUserId(value string) LineUserId {
	return LineUserId{
		Value: value,
	}
}
