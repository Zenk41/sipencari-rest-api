package constant

type HelpfulEnum string

const (
	HelpfulYes HelpfulEnum = "Yes"
	HelpfulNo  HelpfulEnum = "No"
)

func (he HelpfulEnum) String() string {
	return string(he)
}
