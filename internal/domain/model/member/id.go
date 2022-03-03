package member

type Id struct {
	Value int64
}

func NewId(value int64) Id {
	return Id{
		Value: value,
	}
}
