package constant

type TypeEnum string

const (
	TypeFound StatusEnum = "Found"
	TypeLost  StatusEnum = "Lost"
)

func (te TypeEnum) String() string {
	return string(te)
}

