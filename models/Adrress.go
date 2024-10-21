package main

type Address struct {
	Name    string
	City    string
	Country string
}

func GetName(a *Address) string {
	return a.Name
}
