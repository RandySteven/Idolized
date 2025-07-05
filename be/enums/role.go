package enums

type RoleType string

const (
	Guest  RoleType = `user_guest`
	Member RoleType = `user_member`
	Talent RoleType = `user_talent`
	Group  RoleType = `talent_group`
)
