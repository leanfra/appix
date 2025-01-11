package repo

type Group struct {
	User string
	Role string
}

type GroupFilter struct {
	User string
	Role string
}

type Rule struct {
	Sub        string
	ResourceId string
	Action     string
}

type RuleFilter struct {
	Sub        string
	ResourceId string
}

type AuthenRequest struct {
	Sub        string
	ResourceId string
	Action     string
}
