package enums

type AccountIDPrefix string

const (
	UserAccountPrefix   AccountIDPrefix = `00`
	TalentAccountPrefix AccountIDPrefix = `01`
	GroupAccountPrefix  AccountIDPrefix = `03`
)

func (a AccountIDPrefix) String() string {
	return string(a)
}
