package impl

import (
	domainProducts "Fiber/src/main/app/api/products/domain"
	infra "Fiber/src/main/infra"

	"github.com/gofiber/fiber/v2"
	"log"
	"fmt"
	"strconv"
	"encoding/json"
)

func GetAllProducts(c *fiber.Ctx) error {
	// Select query
	rows, err := infra.DB.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Menampung data
	var products []domainProducts.Product

	for rows.Next(){
		var product domainProducts.Product
		if 	err := rows.Scan(&product.ID, &product.ProductName, &product.Detail, &product.Quantity); err != nil {
			log.Fatal(err)
			return err
		}
		
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Response API Product
	response := infra.ListResponseAPI("Products retrieved successfully", "success", fiber.StatusOK, products, len(products))
	return c.JSON(response)
}

func GetProductByID(c *fiber.Ctx) error  {
	// Get ID
	idProductStr := c.Params("id")

	// Convert ID to int
	idProduct, err := strconv.Atoi(idProductStr)
	if err != nil {
		log.Fatal(err)
	}

	// Select query by id
	rows, err := infra.DB.Query("SELECT * FROM `products` WHERE id = ?;", idProduct)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println(rows)
	var product domainProducts.Product
	// fmt.Println(product)

	if rows.Next() {
		if err := rows.Scan(&product.ID, &product.ProductName, &product.Detail, &product.Quantity); err != nil {
			log.Fatal(err)
			return err
		}		
		} 	else {
				responseFailed := infra.ResponseAPI("Products retrived failed", "failed", fiber.StatusBadRequest, product)
				return c.JSON(responseFailed)
			}

	response := infra.ResponseAPI("Products retrieved successfully", "success", fiber.StatusOK, product)
	return c.JSON(response)
}

func CreateProducts(c *fiber.Ctx) error  {
	rawBodyProduct := c.Body()

	var product domainProducts.Product
	if err := json.Unmarshal(rawBodyProduct, &product); err != nil {
		log.Fatal(err)
		return err
	}

	if product.ProductName == "" || product.Detail == "" {
		responseFailed := infra.ResponseAPI("Products retrived failed", "failed", fiber.StatusBadRequest, product)
		return c.JSON(responseFailed)
	} 

	result, err := infra.DB.Exec("INSERT INTO products(product_name, detail, quantity) VALUES (?, ?, ?)",
		product.ProductName, product.Detail, product.Quantity)
	if err != nil {
		log.Fatal(err)
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return err
	}

	product.ID = int(lastInsertID)

	response := infra.ResponseAPI("Product created successfully", "success", fiber.StatusOK, product)
	return c.JSON(response)
}


func UpdateProduct(c *fiber.Ctx) error  {
	idProduct := c.Params("id")
	rawBodyProduct := c.Body()

	var updateProduct domainProducts.Product

	if err := json.Unmarshal(rawBodyProduct, &updateProduct); err != nil {
		log.Fatal(err)
		return err
	}

	// Execute the UPDATE statement
	result, err := infra.DB.Exec("UPDATE products SET product_name = ?, detail = ?, quantity = ? WHERE id = ?;",
		updateProduct.ProductName, updateProduct.Detail, updateProduct.Quantity, idProduct)

	if err != nil {
		log.Fatal(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if rowsAffected == 0 {
		responseNotFound := infra.ResponseAPI("Product not failed!", "failed", fiber.StatusBadRequest, nil)
		return c.JSON(responseNotFound)
	}

	err = infra.DB.QueryRow("SELECT * FROM products WHERE id = ?;", idProduct).
		Scan(&updateProduct.ID, &updateProduct.ProductName, &updateProduct.Detail, &updateProduct.Quantity)

	if err != nil {
		log.Fatal(err)
		return err
	}

	response := infra.ResponseAPI("Product updated successfully", "success", fiber.StatusOK, updateProduct)
	return c.JSON(response)


}

func DeleteProductByID(c *fiber.Ctx) error  {
	idProduct := c.Params("id")

	// Get data sebelum hapus
	var deletedProduct domainProducts.Product
	err := infra.DB.QueryRow("SELECT * FROM products WHERE id = ?;", idProduct).Scan(
		&deletedProduct.ID, &deletedProduct.ProductName, &deletedProduct.Detail, &deletedProduct.Quantity,
	)

	if err != nil {
		responseNotFound := infra.ResponseAPI("Product not found", "failed", fiber.StatusBadRequest, nil)
		return c.JSON(responseNotFound)
	}

	// Delete query
	_, err = infra.DB.Exec("DELETE FROM products WHERE id = ?;", idProduct)
	if err != nil {
		log.Fatal(err)
		return err
	}

	response := infra.ResponseAPI("Product deleted successfully", "success", fiber.StatusOK, deletedProduct)
	return c.JSON(response)
}