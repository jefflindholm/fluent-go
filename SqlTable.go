package main

// SQLTable - a table in the SQL string
type SQLTable struct {
	_name  string
	_alias string
}

// Name - the name of the table
func (mySelf SQLTable) Name() string {
	return mySelf._name
}

// Alias - the name of the table to be used when doing <table>.<col>
func (mySelf SQLTable) Alias() string {
	if mySelf._alias != "" {
		return mySelf._alias
	}
	return mySelf._name
}
