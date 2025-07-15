package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/silviofgantunes/find-perfect-numbers/model"
	"github.com/silviofgantunes/find-perfect-numbers/service"
	"log"
	"net/http"
	"time"
)

var validate = validator.New()

func CheckPerfectNumbers(c echo.Context) error {
	var req model.Request

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request: " + err.Error()})
	}

	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Validation failed: " + err.Error()})
	}

	log.Println("Timing Sequential Method...")
	startTime := time.Now()
	seqResults := service.FindPerfectNumbers(req.Start, req.End)
	seqDuration := time.Since(startTime)
	log.Printf("Sequential Completed in: %v\n", seqDuration)
	log.Printf("Perfect Numbers (sequential): %v\n", seqResults)

	log.Println("Timing Concurrent Method...")
	startTime = time.Now()
	conResults := service.FindPerfectNumbersParallel(req.Start, req.End)
	conDuration := time.Since(startTime)
	log.Printf("Concurrent Completed in: %v\n", conDuration)
	log.Printf("Perfect Numbers (concurrent): %v\n", conResults)

	return c.JSON(http.StatusOK, model.Response{PerfectNumbers: conResults})
}
