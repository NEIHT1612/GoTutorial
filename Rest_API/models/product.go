package models

import "example.com/main/db"

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

func GetAllProducts() ([]Product, error){
	query := "SELECT * FROM products"
	rows, err := db.DB.Query(query)
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next(){
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.Quantity, &product.Description)
		if err != nil{
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *Product) CreateProduct() error{
	query := `INSERT INTO products(name, category, price, quantity, description)
			  VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(p.Name, p.Category, p.Price, p.Quantity, p.Description)
	if err != nil{
		return err
	}
	id, err := result.LastInsertId()
	if err != nil{
		return err
	}
	p.ID = id
	return err
}

func GetProductById(productId int64) (*Product, error){
	query := `SELECT * FROM products WHERE id = ?`
	row := db.DB.QueryRow(query, productId)
	var product Product
	err := row.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.Quantity, &product.Description)
	if err != nil{
		return nil, err
	}
	return &product, nil
}

func (p Product) UpdateProduct() error{
	query := `UPDATE products
			  SET name = ?, category = ?, price = ?, quantity = ?, description = ?
			  WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Name, p.Category, p.Price, p.Quantity, p.Description, p.ID)
	return err
}

func (p Product) DeleteProduct() error{
	query := `DELETE FROM products
			  WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.ID)
	return err
}