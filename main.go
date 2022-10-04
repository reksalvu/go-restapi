package main

import (
	"example/restapi/controllers/activitycontroller"
	"example/restapi/controllers/schedulecontroller"
	"example/restapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/schedules", schedulecontroller.Get_Schedules)
	r.GET("/schedules/:id", schedulecontroller.Get_Schedule_byID)
	r.POST("/schedules", schedulecontroller.Create_Schedule)
	r.PUT("/schedules/:id", schedulecontroller.Update_Schedule_byID)
	r.DELETE("/schedules", schedulecontroller.Delete_Schedule)
	r.GET("/activities/schedule/:id", activitycontroller.Act_bySchId)
	r.GET("/activities/:id", activitycontroller.Act_byActId)
	r.POST("/activities", activitycontroller.Create_Activity)
	r.PUT("/activities/:id", activitycontroller.Update_Activity_byID)
	r.DELETE("/activities", activitycontroller.Delete_Activity)

	r.Run()
}
