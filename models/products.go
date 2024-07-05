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

	query, err := db.Query("select * from products order by id")
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
		p.Id = id
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

func Edit(id int) Product {
	db := db.ConnectDatabase()

	query, err := db.Query("select * from products where id = ?", id)

	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for query.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = query.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

	}

	defer db.Close()

	return product
}

func Update(id int, name, description string, price float64, quantity int) {
	db := db.ConnectDatabase()

	update, err := db.Prepare("update products SET name = ?, description = ?, price = ?, quantity = ? where id = ?")

	if err != nil {
		panic(err.Error())
	}

	update.Exec(name, description, price, quantity, id)

	defer db.Close()

}

func Delete(id int) {
	db := db.ConnectDatabase()

	delete, err := db.Prepare("delete from products where id = ?")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}
