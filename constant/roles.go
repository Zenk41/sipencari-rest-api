package constant

type RoleEnum string

const (
	RoleSuperadmin RoleEnum = "Superadmin"
	RoleAdmin      RoleEnum = "Admin"
	RoleUser       RoleEnum = "User"
)

func (re RoleEnum) String() string {
	return string(re)
}
