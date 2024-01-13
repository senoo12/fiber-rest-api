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
	rows, err := infra.DB.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

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

	response := infra.ListResponseAPI("Products retrieved successfully", "success", fiber.StatusOK, products, len(products))
	return c.JSON(response)
}

func GetProductByID(c *fiber.Ctx) error  {
	idProductStr := c.Params("id")

	idProduct, err := strconv.Atoi(idProductStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := infra.DB.Query("SELECT * FROM `products` WHERE id = ?;", idProduct)
	
	fmt.Println(idProduct)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next(){
		var product domainProducts.Product
		if 	err := rows.Scan(&product.ID, &product.ProductName, &product.Detail, &product.Quantity); err != nil {
			log.Fatal(err)
			return err
		}
		
		products = append(product)
	}

	if !rows.Next(){
		responseFailed := infra.ResponseAPI("Products retrived failed", "failed", fiber.StatusBadRequest, products)
		return c.JSON(responseFailed)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	response := infra.ResponseAPI("Products retrieved successfully", "success", fiber.StatusOK, products)
	return c.JSON(response)

	// fmt.Println(products)
}

func CreateProducts(c *fiber.Ctx) error  {
	rawBodyProduct := c.Body()

	var product domainProducts.Product
	if err := json.Unmarshal(rawBodyProduct, &product); err != nil {
		log.Fatal(err)
		return err
	}

	_, err := infra.DB.Exec("INSERT INTO products(product_name, detail, quantity) VALUES (?, ?, ?)",
		product.ProductName, product.Detail, product.Quantity)
	if err != nil {
		log.Fatal(err)
		return err
	}

	response := infra.ResponseAPI("Product created successfully", "success", fiber.StatusOK, product)
	return c.JSON(response)
}


func UpdateProduct()  {
	
}

func DeleteProductByID(c *fiber.Ctx) error  {
	idProduct := c.Params("id")

	rows, err := infra.DB.Query("DELETE FROM products WHERE id = ?;", idProduct)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

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

	response := infra.ResponseAPI("Products retrieved successfully", "success", fiber.StatusOK, products)
	return c.JSON(response)
}