package fluentSQL

import (
	"fmt"
	"reflect"
)

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
