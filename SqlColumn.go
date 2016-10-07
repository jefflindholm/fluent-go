package fluentSQL

// SQLColumn represents a column in a database, or a litteral value
type SQLColumn struct {
	table    SQLObject
	name     string
	alias    string
	literal  string
	sortDesc bool
}

// Desc - is used to set the sort order for column to Desc
func (column SQLColumn) Desc() SQLColumn {
	column.sortDesc = true
	return column
}

// As - is used to determine the name that would appear in the SQL statement results
func (column SQLColumn) As(newName string) SQLColumn {
	column.alias = newName
	return column
}

// Name - is <raw table name>.<column name in table>
func (column SQLColumn) Name() string {
	if len(column.literal) > 0 {
		return "(" + column.literal + ")"
	}
	return column.table.Alias() + "." + column.name
}

// Alias - is the name that appears in the SQL statement results, set via As above
func (column SQLColumn) Alias() string {
	if len(column.alias) > 0 {
		return column.alias
	}
	return column.name
}

// Eq - create a where clause that does an equality compare
func (column SQLColumn) Eq(value interface{}) SQLWhere {
	return column.Op("=", value)
}

// Op - create where clause that does some simple one item compare (=, <, >, >=, <=, !=, etc)
func (column SQLColumn) Op(op string, value interface{}) SQLWhere {
	return SQLWhere{column: column, op: " " + op + " ", value: value}
}

// Like - create where clause for LIKE compare
func (column SQLColumn) Like(value string) SQLWhere {
	return SQLWhere{column: column, op: " LIKE ", value: "%" + value + "%"}
}

// Between - create where clause for column between 2 values
func (column SQLColumn) Between(val1 interface{}, val2 interface{}) SQLWhere {
	return SQLWhere{column: column, op: " BETWEEN ", value: []interface{}{val1, val2}}
}

// In - create where clause for SQL IN statement
func (column SQLColumn) In(value ...interface{}) SQLWhere {
	var values []interface{}
	for _, v := range value {
		values = append(values, v)
	}
	return SQLWhere{column: column, op: " IN ", value: values}
}
