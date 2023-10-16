package routes

import (
	"github.com/gin-gonic/gin"
	"graph-service/api/controllers"
)

func InitRoute(e *gin.Engine) *gin.Engine {

	// version group
	v1Group := e.Group("/api/v1")

	// pledge-defi backend
	userLocationController := controllers.UserLocationController{}
	v1Group.GET("/getLocationByUserAddress", userLocationController.GetLocationByUserAddress) // 根据用户地址获取用户地区信息 , http://127.0.0.1:8000/api/v1/getLocationByUserAddress?User=0xeea7fcbFdD1EAAc2e790EAc9d375B2ef105a4262
	v1Group.GET("/userLocation", userLocationController.UserLocation)
	v1Group.GET("/getCityIdByCityName", userLocationController.GetCityIdByCityName) // 根据城市获取城市ID

	pioneerController := controllers.PioneerController{}
	v1Group.GET("/getRechargeWeightByPioneerAddress", pioneerController.GetRechargeWeightByPioneerAddress) // 根据城市先锋账户地址查询对应城市每天的充值权重
	v1Group.GET("/rechargeWeight", pioneerController.RechargeWeight)                                       // 充值权重

	v1Group.GET("/reward", pioneerController.Reward) // 奖励查询

	return e
}
