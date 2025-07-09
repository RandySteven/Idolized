package enums

type AccountIDPrefix string

const (
	UserAccountPrefix   AccountIDPrefix = `00`
	TalentAccountPrefix AccountIDPrefix = `01`
	GroupAccountPrefix  AccountIDPrefix = `03`
)
