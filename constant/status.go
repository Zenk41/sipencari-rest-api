package constant

type StatusEnum string

const (
	StatusFound    StatusEnum = "Found"
	StatusNotFound StatusEnum = "NotFound"
)

func (se StatusEnum) String() string {
	return string(se)
}
