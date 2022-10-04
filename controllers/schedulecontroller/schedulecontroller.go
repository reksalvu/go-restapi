package schedulecontroller

import (
	"encoding/json"
	"example/restapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Get_Schedules(c *gin.Context) {
	var schedules []models.Schedule

	if title := c.DefaultQuery("title", ""); title == "" {
		models.DB.Model(&schedules).Preload("Activities").Find(&schedules)
	} else {
		title := "%" + title + "%"
		models.DB.Model(&schedules).Preload("Activities").Where("title LIKE ?", title).Find(&schedules)
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": schedules})
}

func Get_Schedule_byID(c *gin.Context) {
	var schedule models.Schedule
	id := c.Param("id")

	if err := models.DB.Model(&schedule).Preload("Activities").First(&schedule, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "schedule not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": schedule})
}

func Create_Schedule(c *gin.Context) {
	var schedule models.Schedule

	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Model(&schedule).Preload("Activities").Create(&schedule)
	c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "data": schedule})
}

func Update_Schedule_byID(c *gin.Context) {
	var schedule models.Schedule
	id := c.Param("id")

	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&schedule).Where("id = ?", id).Updates(&schedule).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "schedule not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "schedule has been updated"})
}

func Delete_Schedule(c *gin.Context) {
	var schedule models.Schedule
	var input struct {
		Id json.Number
	}
	id, _ := input.Id.Int64()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Delete(&schedule, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "schedule not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "schedule has been deleted"})
}
