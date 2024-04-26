package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-server/model/request"
	"go-server/model/response"
	"net/http"
	"time"
)

func Init(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRouterGroup := engine.Group("/api")
	apiRouterGroup.GET("/list", func(ctx *gin.Context) {
		testData := response.TestDataFormat{
			Key:   "123",
			Title: "TestTitle",
			Date:  time.Now(),
		}
		ctx.JSON(http.StatusOK, testData)
	})
	apiRouterGroup.POST("/create", func(ctx *gin.Context) {
		var reqBody request.TestDataFormat
		if err := ctx.Bind(&reqBody); err != nil {
			ctx.JSON(http.StatusBadRequest, "parameters error")
			return
		}

		respData := response.TestDataFormat{
			Key:   reqBody.Id,
			Title: reqBody.Title,
			Date:  reqBody.Date,
		}

		ctx.JSON(http.StatusOK, respData)
	})
}
