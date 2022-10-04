package activitycontroller

import (
	"encoding/json"
	"example/restapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Act_bySchId(c *gin.Context) {
	var activities []models.Activity
	id := c.Param("id")

	if err := models.DB.Model(&activities).Where("schedule_id = ?", id).Find(&activities).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "activity not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": activities})
}

func Act_byActId(c *gin.Context) {
	var activity models.Activity
	id := c.Param("id")

	if err := models.DB.First(&activity, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "activity not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": activity})
}

func Create_Activity(c *gin.Context) {
	var activity models.Activity

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&activity)
	c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "data": activity})
}

func Update_Activity_byID(c *gin.Context) {
	var activity models.Activity
	id := c.Param("id")

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&activity).Where("id = ?", id).Updates(&activity).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "activity not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "activity has been updated"})
}

func Delete_Activity(c *gin.Context) {
	var activity models.Activity
	var input struct {
		Id json.Number
	}
	id, _ := input.Id.Int64()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Delete(&activity, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "activity not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "activity has been deleted"})
}
