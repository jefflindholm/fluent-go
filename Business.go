package fluentSQL

// Business - Business table
type Business struct {
	SQLTable
	star           []*SQLColumn
	id             SQLColumn
	businessName   SQLColumn
	businessNumber SQLColumn
	parentId       SQLColumn
}

// On - builds a Join statement for this table to be joined into the SQL statement
func (mySelf Business) On(column SQLColumn) SQLJoin {
	return SQLJoin{from: mySelf, on: column}
}

// As - allows the table to change the alias used
func (mySelf Business) As(newName string) Business {
	mySelf._alias = newName
	mySelf.setStar()
	for _, c := range mySelf.star {
		c.table = mySelf
	}
	return mySelf
}

// Star - return a list of all the columns in the table
func (mySelf Business) Star() []SQLColumn {
	var results []SQLColumn
	for _, c := range mySelf.star {
		results = append(results, *c)
	}
	return results
}

// MakeBusiness - the constructor
func MakeBusiness() Business {
	result := Business{}
	result._name = "businesses"
	result.id = SQLColumn{name: "id", table: result}
	result.businessName = SQLColumn{name: "business_name", alias: "businessName", table: result}
	result.businessNumber = SQLColumn{name: "business_number", alias: "businessNumber", table: result}
	result.parentId = SQLColumn{name: "parent_id", alias: "parentId", table: result}
	result.setStar()
	return result
}
func (mySelf *Business) setStar() {
	mySelf.star = []*SQLColumn{&mySelf.id, &mySelf.businessName, &mySelf.businessNumber, &mySelf.parentId}
}
