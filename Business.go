package main

type Business struct {
	SqlTable
	star           []SqlColumn
	id             SqlColumn
	businessName   SqlColumn
	businessNumber SqlColumn
}

func (self Business) As(newName string) Business {
	self.alias = newName
	// since we made a new instance of the table assign all its columns tables to itself
	self.id.table = self
	self.businessName.table = self
	self.businessNumber.table = self
	return self
}
func MakeBusiness() Business {
	result := Business{}
	result.name = "businesses"
	result.id = SqlColumn{name: "id", table: result}
	result.businessName = SqlColumn{name: "business_name", alias: "businessName", table: result}
	result.businessNumber = SqlColumn{name: "business_number", alias: "businessNumber", table: result}
	result.star = []SqlColumn{result.id, result.businessName, result.businessNumber}
	return result
}
