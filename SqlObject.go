package main

type SqlObject interface {
	Name() string
	Alias() string
}

func Complete(self SqlObject) string {
	return self.Name() + " as " + self.Alias()
}
