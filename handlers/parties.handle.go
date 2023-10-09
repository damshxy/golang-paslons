package handlers

import (
	"go-paslons-crud/config"
	"go-paslons-crud/models"

	"github.com/gin-gonic/gin"
)

func CreateParties(c *gin.Context) {
	partiesCreate := new(models.Parties)

	err := c.ShouldBindJSON(&partiesCreate)
	if err != nil {
		c.JSON(400, gin.H{"error": "failed create votes"})
	}

	if partiesCreate.Name == "" {
		c.JSON(400, gin.H{
			"err": "parties name is required",
		})
	}

	if partiesCreate.PaslonsID == 0 {
		c.JSON(400, gin.H{
			"error": "paslons id is required",
		})
	}

	config.DB.Preload("Paslons").Create(&partiesCreate)

	c.JSON(200, gin.H{
		"message": "Create data succesfully",
		"parties": partiesCreate,
	})
}

func GetParties(c *gin.Context) {
	var parties []models.Parties

	config.DB.Preload("Paslons").Find(&parties)

	c.JSON(200, gin.H{
		"message": "Data Parties Found !!!",
		"parties": parties,
	})
}

func GetPartiesById(c *gin.Context) {
	id := c.Param("id")

	var parties models.Parties

	config.DB.Preload("Paslons").First(&parties, id)

	c.JSON(200, gin.H{
		"message": "Data found !!!",
		"parites": parties,
	})
}

func UpdateParties(c *gin.Context) {
	id := c.Param("id")

	var partiesBody struct {
		Name string
	}

	c.Bind(&partiesBody)

	var parties models.Parties
	config.DB.First(&parties, id)
	
	config.DB.Model(&parties).Preload("Paslons").Updates(models.Parties{
		Name: partiesBody.Name,
	})

	c.JSON(200, gin.H{
		"message": "Update data succesfully",
		"parties": parties,
	})
}

func DeleteParties(c *gin.Context) {
	id := c.Param("id")

	config.DB.Delete(&models.Parties{}, id)

	c.JSON(200, gin.H{
		"message": "Parties Deleted !!!",
	})
}