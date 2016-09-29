package main

// SQLJoin represents a join definition
type SQLJoin struct {
	from  SQLObject
	on    SQLColumn
	using SQLColumn
	outer bool
	right bool
}

// Outer - set the join to be an outer join
func (mySelf SQLJoin) Outer() SQLJoin {
	mySelf.outer = true
	return mySelf
}

// Right - set the join to be a right join
func (mySelf SQLJoin) Right() SQLJoin {
	mySelf.right = true
	return mySelf
}

// Using - the function to set the column from existing tables used to join the new table into the result
func (mySelf SQLJoin) Using(column SQLColumn) SQLJoin {
	mySelf.using = column
	return mySelf
}
