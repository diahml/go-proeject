package productcontroller

import (
	"net/http"
	"strconv"

	"github.com/diahml/go-project/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index (c *gin.Context){
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show (c *gin.Context){
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
			return
		default : 
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create (c *gin.Context){

	var product models.Product // make product variable with struct type Product

	//mengambil data dr body json
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	models.DB.Create(&product) //insert data to database
	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Update (c *gin.Context){
	var product models.Product
	id := c.Param("id")  //variable id mengambil dari parameter id di url

	if err := c.ShouldBindJSON(&product); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id =?", id).Updates(&product).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : "Data cannot be updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Data updated successfully"})
}

func Delete (c *gin.Context){
	var product models.Product

	input := map[string]string{"id" : "0"} //menampung variable id yg ada di body

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	id, _ := strconv.ParseInt(input["id"], 10, 64)

	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : "Data cannot be deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Data Deleted"})
	
}