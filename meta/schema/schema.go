package schema

import "github.com/go-mysql-org/go-mysql/schema"

type Column struct {
	Name       string
	Type       int
	Collation  string
	RawType    string
	IsAuto     bool
	IsUnsigned bool
	IsVirtual  bool
	EnumValues []string
	SetValues  []string
	FixedSize  uint
	MaxSize    uint
}

func FromColumn(col schema.TableColumn) Column {
	return Column{Name: col.Name,
		Type:       col.Type,
		Collation:  col.Collation,
		RawType:    col.RawType,
		IsAuto:     col.IsAuto,
		IsUnsigned: col.IsUnsigned,
		IsVirtual:  col.IsVirtual,
		EnumValues: col.EnumValues,
		SetValues:  col.SetValues,
		FixedSize:  col.FixedSize,
		MaxSize:    col.MaxSize,
	}
}

func FromColumns(cols []schema.TableColumn) []Column {
	l := len(cols)
	var rCols []Column
	for idx := 0; idx < l; idx++ {
		rCols = append(rCols, FromColumn(cols[idx]))
	}
	return rCols
}

type Table struct {
	Name string
}
