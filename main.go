package main

import (
	"fmt"
)

type SqlTable struct {
	name  string
	alias string
}

func (self SqlTable) Name() string {
	return self.name
}
func (self SqlTable) Alias() string {
	if len(self.alias) > 0 {
		return self.alias
	}
	return self.name
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
	fmt.Println("business", business.name, business.alias)
	fmt.Println("b", b.name, b.alias)
	query := SqlQuery{}.From(b).Select(b.id.As("Identifier"), b.businessName)
	fmt.Println("sql", query.GenSql())
}
