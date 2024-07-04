package models

import (
	"store/db"
	_ "store/db"
)

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func GetAll() []Product {
	db := db.ConnectDatabase()

	query, err := db.Query("select * from products")
	if err != nil {
		panic(err)
	}

	p := Product{}
	products := []Product{}

	for query.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = query.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)

	}
	defer db.Close()

	return products
}

func Create(name, description string, price float64, quantity int) {
	db := db.ConnectDatabase()

	insert, err := db.Prepare("insert into products (name, description, price, quantity) values (?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	insert.Exec(name, description, price, quantity)

	defer db.Close()

}
