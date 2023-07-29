package constant

type TypeEnum string

const (
	TypeFound StatusEnum = "Found"
	TypeLost  StatusEnum = "Lost"
)

func (se TypeEnum) String() string {
	return string(se)
}

