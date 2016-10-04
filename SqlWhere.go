package main

// SQLWhere - where clause for the query
type SQLWhere struct {
	column      SQLColumn
	op          string
	value       interface{}
	conjunction string
	wheres      []SQLWhere
	not         bool
}

// Not - do a NOT of the containing Where statement
func (mySelf SQLWhere) Not(wheres ...SQLWhere) SQLWhere {
	if wheres != nil && len(wheres) > 0 {
		newWhere := SQLWhere{conjunction: wheres[0].conjunction, not: true}
		newWhere.wheres = append(newWhere.wheres, mySelf)
		for _, where := range wheres {
			newWhere.wheres = append(newWhere.wheres, where)
		}
		return newWhere
	}
	mySelf.not = true
	return mySelf
}

// Or - build an OR'ed list of conditionals
func (mySelf SQLWhere) Or(where SQLWhere) SQLWhere {
	newWhere := SQLWhere{conjunction: "OR"}
	newWhere.wheres = append(newWhere.wheres, mySelf)
	newWhere.wheres = append(newWhere.wheres, where)
	return newWhere
}
