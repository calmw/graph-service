package services

import (
	"fmt"
	"gorm.io/gorm"
	"graph-service/api/common/statecode"
	"graph-service/api/models/request"
	"graph-service/db"
	models2 "graph-service/models"
	"graph-service/utils"
	"strconv"
	"strings"
)

type UserLocationService struct{}

func NewUserLocation() *UserLocationService {
	return &UserLocationService{}
}

func (c *UserLocationService) GetLocationByUserAddress(userReq *request.GetLocationByUserAddress) (int, []models2.UserLocation) {
	var userLocation []models2.UserLocation
	whereCondition := fmt.Sprintf("user='%s'", strings.ToLower(userReq.User))
	db.Mysql.Table("user_location").Where(whereCondition).Find(&userLocation)

	return statecode.CommonSuccess, userLocation
}

func (c *UserLocationService) UserLocation(req *request.Location) (int, int64, []models2.UserLocation) {
	var userLocations []models2.UserLocation

	tx := db.Mysql.Model(&models2.UserLocation{})
	if req.User != "" {
		tx = tx.Where("user=?", strings.ToLower(req.User))
	}

	var total int64
	tx.Count(&total)
	page := 1 // 第5页
	if req.Page > 0 {
		page = req.Page
	}
	pageSize := 10
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}
	offset := (page - 1) * pageSize // 计算偏移量
	err := tx.Debug().Limit(pageSize).Offset(offset).Find(&userLocations).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return statecode.CommonErrServerErr, 0, userLocations
	}
	return statecode.CommonSuccess, total, userLocations
}

func (c *UserLocationService) CityId(req *request.CityName) (int, string) {

	if req.Name == "" {
		return statecode.CommonErrServerErr, ""
	}

	var areaCode models2.AreaCode

	err := db.Mysql.Model(&models2.AreaCode{}).Where("city_name=?", req.Name).First(&areaCode).Debug().Error
	var codeStr string
	if err == nil {
		CountryCodeStr := strconv.FormatInt(int64(areaCode.CountryCode), 10)
		CityCodeStr := strconv.FormatInt(int64(areaCode.CityCode), 10)
		codeStr = fmt.Sprintf("%s,%s", CountryCodeStr, CityCodeStr)
	} else {
		return statecode.CommonErrServerErr, ""
	}
	return statecode.CommonSuccess, utils.ThreeDesEncrypt(codeStr)
}
