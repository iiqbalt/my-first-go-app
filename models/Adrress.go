package main

type Address struct {
	Name    string
	City    string
	Country string
}

func GetName(a *Address) string {
	return a.Name
}

func GetCity(b *Address) string {
	return b.City
}

func GetCountry(c *Address) string {
	return c.Country
}
