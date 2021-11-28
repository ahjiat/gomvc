package Status

const (
	Deleted StatusType = iota + 0
	Actived
	Inactived
)

type StatusType int
func (self *StatusType) String() string {
	switch *self {
		case Deleted: return "Deleted"
		case Actived: return "Actived"
		case Inactived: return "Inactived"
		default: return ""
	}
}
func (self *StatusType) Int() int {
	return int(*self)
}
