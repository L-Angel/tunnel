package worker

import (
	"github.com/l-angel/tunnel/meta/schema"
	"strings"
)

type Action int

const (
	None   Action = 0
	Update Action = 1
	Delete Action = 2
	Insert Action = 3
	Select Action = 4
	Alter  Action = 5
	Create Action = 6
)

func (self Action) String() string {
	switch self {
	case Update:
		return "update"
	case Delete:
		return "delete"
	case Insert:
		return "insert"
	case Select:
		return "select"
	case Alter:
		return "alter"
	case Create:
		return "create"
	default:
		return "none"
	}
}

func Of(action string) Action {
	switch strings.ToLower(action) {
	case "update":
		return Update
	case "delete":
		return Delete
	case "insert":
		return Insert
	case "select":
		return Select
	case "alter":
		return Alter
	case "create":
		return Create
	default:
		return None
	}
}

type OEvent struct {
	Action  Action
	Table   string
	Rows    [][]interface{}
	Columns []schema.Column
}

type DataEvent struct {
	Action  string          `json:"action"`
	Table   string          `json:"table"`
	Columns []schema.Column `json:"columns"`
}

type ChangeValue struct {
	Index  int         `json:"index"`
	Column string      `json:"column"`
	Value  interface{} `json:"value"`
	Before interface{} `json:"before"`
	After  interface{} `json:"after"`
}

type DeleteEvent struct {
	DataEvent
	Rows [][]ChangeValue `json:"rows"`
}

type UpdateEvent struct {
	DataEvent
	Rows [][]ChangeValue `json:"rows"`
}

type InsertEvent struct {
	DataEvent
	Rows [][]ChangeValue `json:"rows"`
}
