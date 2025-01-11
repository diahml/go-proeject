package main

import(
	"github.com/diahml/go-project/models"
	"github.com/diahml/go-project/controllers/productcontroller"
	"github.com/gin-gonic/gin"

) 

func main(){
	r := gin.Default();
	models.ConncectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()
}