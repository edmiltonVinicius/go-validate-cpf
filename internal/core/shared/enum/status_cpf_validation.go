package enum

type StatusValidation int

const (
	APPROVED StatusValidation = iota + 1
	REPROVED
	PENDING
)

func (s StatusValidation) String() string {
	return [...]string{"APPROVED", "REPROVED", "PENDING"}[s-1]
}
