package main

type SqlColumn struct {
	table SqlObject
	name  string
	alias string
}

func (self SqlColumn) As(newName string) SqlColumn {
	self.alias = newName
	return self
}

func (self SqlColumn) Name() string {
	return self.table.Alias() + "." + self.name
}

func (self SqlColumn) Alias() string {
	if len(self.alias) > 0 {
		return self.alias
	}
	return self.name
}
