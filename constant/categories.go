package constant

type CategoryEnum string

const (
	CategoryPet   CategoryEnum = "Pet"
	CategoryHuman CategoryEnum = "Human"
	CategoryGoods CategoryEnum = "Enum"
)

func (ce CategoryEnum) String() string {
	return string(ce)
}
