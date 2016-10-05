package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// SQLQuery - the actual SQL Query to be processed
type SQLQuery struct {
	columns []SQLColumn
	from    SQLObject
	joins   []SQLJoin
	wheres  []SQLWhere
	orderBy []SQLColumn
}

// From - add the table the values should be taken from - the main table
func (query SQLQuery) From(table SQLObject) SQLQuery {
	query.from = table
	return query
}

// Select - the list of columns to retrieve in the SQL statement
func (query SQLQuery) Select(columns ...SQLColumn) SQLQuery {
	for _, c := range columns {
		query.columns = append(query.columns, c)
	}
	return query
}

// OrderBy - the list of columns to order by
func (query SQLQuery) OrderBy(columns ...SQLColumn) SQLQuery {
	for _, c := range columns {
		query.orderBy = append(query.orderBy, c)
	}
	return query
}

// GenSQL - create the SQL string from the information built
func (query SQLQuery) GenSQL() string {
	SQL := "SELECT\n"
	for i, c := range query.columns {
		if i != 0 {
			SQL += "\n, "
		}
		SQL += Complete(c)
	}
	SQL += "\n"
	SQL += "FROM\n" + Complete(query.from)

	join := ""
	for _, j := range query.joins {
		cmd := ""
		if j.right {
			cmd += "RIGHT "
		}
		if j.outer {
			cmd += "OUTER "
		}
		cmd += "JOIN"
		join += "\n" + cmd + "\n" + Complete(j.from) + " ON " + j.on.Name() + " = " + j.using.Name()

	}
	SQL += join

	where := buildWhere(query.wheres, "AND")
	SQL += "\nWHERE\n" + where

	if len(query.orderBy) > 0 {
		SQL += "\nORDER BY\n"
		for _, o := range query.orderBy {
			dir := " ASC\n"
			if o.sortDesc {
				dir = " DESC\n"
			}
			SQL += o.Name() + dir
		}
	}

	return SQL
}
func buildWhere(wheres []SQLWhere, conjuction string) string {
	where := ""
	for _, w := range wheres {
		var subWhere string
		if len(w.wheres) > 0 {
			subWhere = "(" + buildWhere(w.wheres, w.conjunction) + ")"
		} else {
			val, err := buildValue(w.value, w.op)
			if err != nil {
				fmt.Println(err)
			} else {
				subWhere = w.column.Name() + w.op + val
			}
		}
		if w.not {
			subWhere = "NOT(" + subWhere + ")"
		}
		if where != "" {
			where += "\n" + conjuction + "\n"
		}
		where += subWhere
	}
	return where
}
func buildString(value string) string {
	return "'" + value + "'"
}
func buildValue(value interface{}, op string) (string, error) {
	switch value.(type) {
	case int:
		return strconv.Itoa(value.(int)), nil
	case string:
		return buildString(value.(string)), nil
	case []interface{}:
		if op == " BETWEEN " {
			v1, _ := buildValue(value.([]interface{})[0], "")
			v2, _ := buildValue(value.([]interface{})[1], "")
			return v1 + " AND " + v2, nil
		}
		list := "("
		for i, val := range value.([]interface{}) {
			if i > 0 {
				list += ", "
			}
			v, _ := buildValue(val, "")
			list += v
		}
		list += ")"
		return list, nil
	}
	return "", errors.New("unkown type")
}

// Join - the information used to join a second or more table into the result set
func (query SQLQuery) Join(join SQLJoin) SQLQuery {
	query.joins = append(query.joins, join)
	return query
}

// Where - create a where clause for the query
func (query SQLQuery) Where(where SQLWhere) SQLQuery {
	query.wheres = append(query.wheres, where)
	return query
}

func main() {
	business := MakeBusiness()
	businessAddress := MakeBusinessAddress()
	b := business.As("b")
	b2 := business.As("b2")

	fmt.Println("business = ", business.Name(), business.Alias())
	fmt.Println("b = ", b.Name(), b.Alias())

	query := SQLQuery{}.From(b).Select(b.id.As("Identifier"), b.businessName)
	fmt.Println("SQL", query.GenSQL())
	query = SQLQuery{}.From(b).Select(b.Star()...)
	fmt.Println("SQL", query.GenSQL())
	query = SQLQuery{}.From(business).Select(business.Star()...)
	fmt.Println("SQL", query.GenSQL())
	query = SQLQuery{}.From(businessAddress.As("ba")).Select(businessAddress.As("ba").Star()...)
	fmt.Println("SQL", query.GenSQL())

	query = SQLQuery{}.From(b).Select(b.Star()...)
	query = query.Join(businessAddress.On(businessAddress.businessId).Using(b.id).Right().Outer())
	query = query.Select(businessAddress.Star()...)
	fmt.Println("SQL - join", query.GenSQL())

	literal := SQLColumn{literal: "SELECT count(*) FROM business_addresses WHERE business_id = b.id", alias: "addressCount"}
	query = query.Where(b.businessNumber.Eq("12345")).Where(b.id.Eq(4001)).Where(b.id.Between(1, 10))
	query = query.
		Select(literal).
		Where(SQLWhere{}.Not(businessAddress.zip.In(46062, 46032))).
		Where(b.businessName.In("Bubba Car World", "Bubbas Cars").Or(b.businessName.Like("fred")).Not()).
		Join(b2.On(b2.id).Using(b.id)).
		Select(b2.businessName.As("parentName")).
		OrderBy(b.businessName.Desc(), b.businessNumber, literal)
	fmt.Println("SQL - where", query.GenSQL())
	var u interface{}
	u = "some string"
	fmt.Println(u, reflect.TypeOf(u))
	u = 123
	fmt.Println(u, reflect.TypeOf(u))
}
