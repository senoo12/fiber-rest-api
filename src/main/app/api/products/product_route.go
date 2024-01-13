package products

import (
	productsRepository "Fiber/src/main/app/api/products/repository/impl"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App)  {
	app.Route("/products", func(api fiber.Router){
		api.Get("", productsRepository.GetAllProducts)
		api.Get("/:id", productsRepository.GetProductByID)
		api.Post("/", productsRepository.CreateProducts)
		api.Put("/:id", productsRepository.UpdateProduct)
		api.Delete("/:id", productsRepository.DeleteProductByID)
	})
}