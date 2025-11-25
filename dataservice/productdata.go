package dataservice

import (
	"database/sql"
	"goinventory/model"
)

// Add Product
func AddProduct(db *sql.DB, product model.Product) (int, error) {
	query := `INSERT INTO products (name, price, stock) VALUES (?, ?, ?)`

	result, err := db.Exec(query, product.Name, product.Price, product.Stock)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertID), nil
}

// Get Product
func GetProduct(db *sql.DB, id int) (model.Product, error) {
	var product model.Product

	query := `SELECT id, name, price, stock, created_at FROM products WHERE id = ?`
	row := db.QueryRow(query, id)

	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CreatedAt)
	return product, err
}

// Update Stock
func UpdateStock(db *sql.DB, id int, change int) (int, int, error) {
	var oldStock int

	selectQuery := `SELECT stock FROM products WHERE id = ?`
	row := db.QueryRow(selectQuery, id)

	err := row.Scan(&oldStock)
	if err != nil {
		return 0, 0, err
	}

	newStock := oldStock + change
	if newStock < 0 {
		newStock = 0
	}

	updateQuery := `UPDATE products SET stock = ? WHERE id = ?`
	_, err = db.Exec(updateQuery, newStock, id)
	if err != nil {
		return 0, 0, err
	}

	return oldStock, newStock, nil
}

// List Products
func ListProducts(db *sql.DB, search string) ([]model.Product, error) {
	var rows *sql.Rows
	var err error

	baseQuery := `SELECT id, name, price, stock, created_at FROM products`

	if search != "" {
		query := baseQuery + ` WHERE name LIKE ?`
		rows, err = db.Query(query, "%"+search+"%")
	} else {
		rows, err = db.Query(baseQuery)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var product model.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
