package constant

type StatusEnum string

const (
	StatusHasBeenFound StatusEnum = "HasBeenFound"
	StatusNotYetFound  StatusEnum = "NotYetFound"
)

func (se StatusEnum) String() string {
	return string(se)
}
