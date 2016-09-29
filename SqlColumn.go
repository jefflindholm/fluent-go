package main

// SQLColumn represents a column in a database, or a litteral value
type SQLColumn struct {
	table SQLObject
	name  string
	alias string
}

// As - is used to determine the name that would appear in the SQL statement results
func (column SQLColumn) As(newName string) SQLColumn {
	column.alias = newName
	return column
}

// Name - is <raw table name>.<column name in table>
func (column SQLColumn) Name() string {
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
	return SQLWhere{column: column, op: " = ", value: value}
}
