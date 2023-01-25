package constant

type CategoryEnum string

const (
	CategoryPet   CategoryEnum = "Pet"
	CategoryHuman CategoryEnum = "Human"
	CategoryGoods CategoryEnum = "Goods"
)

func (ce CategoryEnum) String() string {
	return string(ce)
}
