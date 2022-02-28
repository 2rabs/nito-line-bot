package value

type NotifyGroupId struct {
	Value string
}

func NewNotifyGroupId(value string) *NotifyGroupId {
	return &NotifyGroupId{Value: value}
}
