package routers

import (
	"github.com/gin-gonic/gin"
	"nodejs_project/controllers"
)

func InitRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 添加 CORS 中间件
	router.Use(CORSMiddleware())

	router.POST("/controllers/login", controllers.Login)
	router.POST("/controllers/register", controllers.Register)
	router.POST("/controllers/testManageMsg", controllers.TestmanageMsg)
	router.POST("/controllers/delete", controllers.Delete)
	router.POST("/controllers/addTest", controllers.AddTest)
	router.POST("/controllers/batchDelete", controllers.BatchDelete)
	router.POST("/controllers/modification", controllers.Modification)
	router.POST("/controllers/Search", controllers.Search)
	router.POST("/controllers/LogInfo", controllers.LogInfo)
	router.POST("/controllers/OpRecord", controllers.OpRecord)
	router.POST("/controllers/PointsAcquistiom", controllers.PointsAcquistiom)

	return router
}
