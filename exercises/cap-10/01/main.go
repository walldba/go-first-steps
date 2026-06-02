package main

import "fmt"

type Business struct {
	name           string
	hasDefaultLogo bool
	location       BusinessLocation
}

type BusinessLocation struct {
	address string
	city    string
	state   string
	zip     string
}

func main() {

	business := Business{
		name:           "Wallace Tech",
		hasDefaultLogo: true,
		location: BusinessLocation{
			address: "Rua das Flores, 123",
			city:    "São Paulo",
			state:   "SP",
			zip:     "01000-000",
		},
	}

	fmt.Println("nome:", business.name)
	fmt.Println("logo padrão:", business.hasDefaultLogo)

	fmt.Println("endereco:", business.location.address)
	fmt.Println("cidade:", business.location.city)
	fmt.Println("estado:", business.location.state)
	fmt.Println("cep:", business.location.zip)

}
