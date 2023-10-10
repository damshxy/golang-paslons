package handlers

import (
	"go-paslons-crud/database"
	"go-paslons-crud/models"

	"github.com/gin-gonic/gin"
)

func CreateVotes(c *gin.Context) {
	votesCreate := new(models.Votes)

	err := c.ShouldBindJSON(&votesCreate)
	if err != nil {
		c.JSON(400, gin.H{"error": "failed create votes"})
	}

	if votesCreate.VoterName == "" {
		c.JSON(400, gin.H{
			"err": "voter name is required",
		})
	}

	if votesCreate.PaslonsID == 0 {
		c.JSON(400, gin.H{
			"err": "paslons id is required",
		})
	}

	database.DB.Preload("Paslons").Create(&votesCreate)

	c.JSON(200, gin.H{
		"message": "create data succesfully",
		"votes": votesCreate,
	})
}

func GetVotes(c *gin.Context) {
	var votes []models.Votes

	database.DB.Preload("Paslons").Find(&votes)

	c.JSON(200, gin.H{
		"message": "data found",
		"votes": votes,
	})
}

func GetVotesById(c *gin.Context) {
	id := c.Param("id")
	var votes models.Votes

	database.DB.Preload("Paslons").First(&votes, id)

	c.JSON(200, votes)
}

func UpdateVotes(c *gin.Context) {
	id := c.Param("id")
	var votesBody struct {
		VoterName string
		PaslonsID int
	}

	c.Bind(&votesBody)

	var votes []models.Votes

	database.DB.First(&votes, id)
	database.DB.Model(&votes).Updates(models.Votes{
		VoterName: votesBody.VoterName,
		PaslonsID: votesBody.PaslonsID,
	})

	c.JSON(200, gin.H{
		"message": "Update successfully",
		"votes": votes,
	})
}

func DeleteVotes(c *gin.Context) {
	id := c.Param("id")

	database.DB.Delete(&models.Votes{}, id)

	c.JSON(200, gin.H{
		"message": "Votes Deleted !!!",
	})
}