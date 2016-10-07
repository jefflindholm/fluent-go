package fluentSQL

// SQLObject - a simple interface to let us get the names of SQL Objects everywhere
type SQLObject interface {
	Name() string
	Alias() string
}

// Complete - function that takes a SQLObject and uses the name and alias to generate NAME AS ALIAS
func Complete(mySelf SQLObject) string {
	return mySelf.Name() + " as " + mySelf.Alias()
}
