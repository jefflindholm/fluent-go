package main

import (
	"fmt"
)

type SqlTable struct {
	_name  string
	_alias string
}

func (self SqlTable) Name() string {
	return self._name
}
func (self SqlTable) Alias() string {
	if self._alias != "" {
		return self._alias
	}
	return self._name
}

type SqlQuery struct {
	columns []SqlColumn
	from    SqlObject
}

func (self SqlQuery) From(table SqlObject) SqlQuery {
	self.from = table
	return self
}

func (self SqlQuery) Select(columns ...SqlColumn) SqlQuery {
	for _, c := range columns {
		self.columns = append(self.columns, c)
	}
	return self
}

func (self SqlQuery) GenSql() string {
	sql := "SELECT\n"
	for i, c := range self.columns {
		if i != 0 {
			sql += "\n, "
		}
		sql += Complete(c)
	}
	sql += "\n"
	sql += "FROM\n" + Complete(self.from)

	return sql
}

func main() {
	business := MakeBusiness()
	b := business.As("b")

	fmt.Println("business = ", business.Name(), business.Alias())
	fmt.Println("b = ", b.Name(), b.Alias())

	query := SqlQuery{}.From(b).Select(b.id.As("Identifier"), b.businessName)
	fmt.Println("sql", query.GenSql())
	query = SqlQuery{}.From(b).Select(b.Star()...)
	fmt.Println("sql", query.GenSql())
	query = SqlQuery{}.From(business).Select(business.Star()...)
	fmt.Println("sql", query.GenSql())
}
