package constant

type PrivacyEnum string

const (
	PrivacyPublic  StatusEnum = "Public"
	PrivacyPrivate StatusEnum = "Private"
)

func (pe PrivacyEnum) String() string {
	return string(pe)
}
