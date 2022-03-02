package member

type Nickname struct {
	Value string
}

func NewNickname(value string) *Nickname {
	return &Nickname{
		Value: value,
	}
}
