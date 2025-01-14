package repo

import (
	"context"
	"errors"
	"strings"
)

const ActWrite = "write"

type Group struct {
	User string
	Role string
}

type GroupFilter struct {
	User string
	Role string
}

type IResource interface {
	ResourceStr() string
	ParseStr(raw string) error
}

type Rule struct {
	Sub      string
	Resource IResource
	Action   string
}

type RuleFilter struct {
	Sub string
}

type AuthenRequest struct {
	Sub      string
	Resource IResource
	Action   string
}

type AuthzRepo interface {
	CreateRule(ctx context.Context, tx TX, policy *Rule) error
	DeleteRule(ctx context.Context, tx TX, policy *Rule) error
	ListRule(ctx context.Context, tx TX, filter *RuleFilter) ([]*Rule, error)
	Enforce(ctx context.Context, tx TX, request *AuthenRequest) (bool, error)
	CreateGroup(ctx context.Context, tx TX, role *Group) error
	DeleteGroup(ctx context.Context, tx TX, role *Group) error
	ListGroup(ctx context.Context, tx TX, filter *GroupFilter) ([]*Group, error)
}

// Resource4Sv1 is a resource string for casbin.
// including 4 sections as resource type, team name, resource instance and user name
type Resource4Sv1 struct {
	ResType  string
	TeamName string
	ResInst  string
	UserName string
}

func NewResource4Sv1(resType, teamName, resInst, userName string) IResource {
	return &Resource4Sv1{
		ResType:  resType,
		TeamName: teamName,
		ResInst:  resInst,
		UserName: userName,
	}
}

func (r *Resource4Sv1) ResourceStr() string {
	str := &strings.Builder{}
	if r.ResType != "" {
		str.WriteString("v1/")
		str.WriteString(r.ResType)
	} else {
		str.WriteString("v1/{resource}")
	}
	if r.TeamName != "" {
		str.WriteString("/")
		str.WriteString(r.TeamName)
	} else {
		str.WriteString("/{team}")
	}
	if r.ResInst != "" {
		str.WriteString("/")
		str.WriteString(r.ResInst)
	} else {
		str.WriteString("/{resource_id}")
	}
	if r.UserName != "" {
		str.WriteString("/")
		str.WriteString(r.UserName)
	} else {
		str.WriteString("/{user}")
	}
	return str.String()
}

func (r *Resource4Sv1) ParseStr(raw string) error {
	_raw := strings.TrimLeft(raw, "v1/")
	parts := strings.Split(_raw, "/")
	if len(parts) != 4 {
		return errors.New("invalid resource string")
	}
	if parts[0] != "" && parts[0] != "{resource}" {
		r.ResType = parts[0]
	}
	if parts[1] != "" && parts[1] != "{team}" {
		r.TeamName = parts[1]
	}
	if parts[2] != "" && parts[2] != "{resource_id}" {
		r.ResInst = parts[2]
	}
	if parts[3] != "" && parts[3] != "{user}" {
		r.UserName = parts[3]
	}
	return nil
}
