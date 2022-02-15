package Controllers

import (
	"Project/Service/ScheduleCourse"
	"Project/Types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ScheCourseController struct {
}

func (s ScheCourseController) ScheduleCourse(c *gin.Context) {
	scheduleCourseRequest := &Types.ScheduleCourseRequest{}

	if err := c.ShouldBindJSON(&scheduleCourseRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	scheduleCourseResponse := ScheduleCourse.Schedule(scheduleCourseRequest)
	c.JSON(http.StatusOK, &scheduleCourseResponse)
}
