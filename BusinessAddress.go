package main

// BusinessAddress - BusinessAddress table defininition - this file should be genereated
type BusinessAddress struct {
	SQLTable
	star       []*SQLColumn
	id         SQLColumn
	businessId SQLColumn
	line1      SQLColumn
	line2      SQLColumn
	city       SQLColumn
	state      SQLColumn
	zip        SQLColumn
}

// On - builds a Join statement for this table to be joined into the SQL statement
func (mySelf BusinessAddress) On(column SQLColumn) SQLJoin {
	return SQLJoin{from: mySelf, on: column}
}

// As - allows the table to change the alias used
func (mySelf BusinessAddress) As(newName string) BusinessAddress {
	mySelf._alias = newName
	mySelf.setStar()
	for _, c := range mySelf.star {
		c.table = mySelf
	}
	return mySelf
}

// Star - return a list of all the columns in the table
func (mySelf BusinessAddress) Star() []SQLColumn {
	var results []SQLColumn
	for _, c := range mySelf.star {
		results = append(results, *c)
	}
	return results
}

// MakeBusinessAddress - the constructor
func MakeBusinessAddress() BusinessAddress {
	result := BusinessAddress{}
	result._name = "business_addresses"
	result.id = SQLColumn{name: "id", table: result}
	result.businessId = SQLColumn{name: "business_id", alias: "businessId", table: result}
	result.line1 = SQLColumn{name: "line1", alias: "line1", table: result}
	result.line2 = SQLColumn{name: "line2", alias: "line2", table: result}
	result.city = SQLColumn{name: "city", alias: "city", table: result}
	result.state = SQLColumn{name: "state", alias: "state", table: result}
	result.zip = SQLColumn{name: "zip", alias: "zip", table: result}
	result.setStar()
	return result
}
func (mySelf *BusinessAddress) setStar() {
	mySelf.star = []*SQLColumn{&mySelf.id, &mySelf.businessId, &mySelf.line1, &mySelf.line2, &mySelf.city, &mySelf.state, &mySelf.zip}
}
