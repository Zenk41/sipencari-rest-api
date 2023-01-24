package constant

type CategoryEnum string

const (
	CategoryPet   CategoryEnum = "Superadmin"
	CategoryHuman CategoryEnum = "Admin"
	CategoryGoods CategoryEnum = "User"
)

func (ce CategoryEnum) String() string {
	return string(ce)
}
