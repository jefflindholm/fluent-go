package main

type Business struct {
	SqlTable
	star           []*SqlColumn
	id             SqlColumn
	businessName   SqlColumn
	businessNumber SqlColumn
}

func (self Business) As(newName string) Business {
	self._alias = newName
	// since we made a new instance of the table assign all its columns tables to itself
	self.id.table = self
	self.businessName.table = self
	self.businessNumber.table = self
	self.SetStar()
	return self
}
func (self Business) Star() []SqlColumn {
	var results []SqlColumn
	for _, c := range self.star {
		results = append(results, *c)
	}
	return results
}
func MakeBusiness() Business {
	result := Business{}
	result._name = "businesses"
	result.id = SqlColumn{name: "id", table: result}
	result.businessName = SqlColumn{name: "business_name", alias: "businessName", table: result}
	result.businessNumber = SqlColumn{name: "business_number", alias: "businessNumber", table: result}
	result.SetStar()
	return result
}
func (self *Business) SetStar() {
	self.star = []*SqlColumn{&self.id, &self.businessName, &self.businessNumber}
}
