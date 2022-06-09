package worker

type Mapping interface {
	mapTo(oe OEvent) interface{}
}

type DeleteMapping struct {
}

func (self *DeleteMapping) mapTo(oe OEvent) interface{} {
	dataLen := len(oe.Rows)
	colLen := len(oe.Columns)
	var data [][]ChangeValue
	for idx := 0; idx < dataLen; idx += 1 {
		var row []ChangeValue
		for cIdx := 0; cIdx < colLen; cIdx += 1 {
			row = append(row, ChangeValue{Index: cIdx, Column: oe.Columns[cIdx].Name, Value: oe.Rows[idx][cIdx]})
		}
		data = append(data, row)
	}
	return &InsertEvent{
		DataEvent: DataEvent{
			Action:  "delete",
			Table:   oe.Table,
			Columns: oe.Columns,
		},
		Rows: data,
	}

}

type InsertMapping struct {
}

func (self *InsertMapping) mapTo(oe OEvent) interface{} {
	dataLen := len(oe.Rows)
	colLen := len(oe.Columns)
	var data [][]ChangeValue
	for idx := 0; idx < dataLen; idx += 1 {
		var row []ChangeValue
		for cIdx := 0; cIdx < colLen; cIdx += 1 {
			row = append(row, ChangeValue{Index: cIdx, Column: oe.Columns[cIdx].Name, Value: oe.Rows[idx][cIdx]})
		}
		data = append(data, row)
	}
	return &InsertEvent{
		DataEvent: DataEvent{
			Action:  "insert",
			Table:   oe.Table,
			Columns: oe.Columns,
		},
		Rows: data,
	}
}

type UpdateMapping struct {
}

func (self *UpdateMapping) mapTo(oe OEvent) interface{} {
	dataLen := len(oe.Rows)
	colLen := len(oe.Columns)
	var data [][]ChangeValue
	for idx := 0; idx < dataLen; idx += 2 {
		before := oe.Rows[idx]
		after := oe.Rows[idx+1]
		for cIdx := 0; cIdx < colLen; cIdx += 1 {
			var row []ChangeValue
			if before[cIdx] != after[cIdx] {
				row = append(row, ChangeValue{Index: cIdx, Column: oe.Columns[cIdx].Name, Before: before[cIdx], After: after[cIdx]})
			}
			data = append(data, row)
		}
	}
	return &InsertEvent{
		DataEvent: DataEvent{
			Action:  "update",
			Table:   oe.Table,
			Columns: oe.Columns,
		},
		Rows: data,
	}
}
