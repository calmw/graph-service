package main

import (
	"graph-service/db"
	"graph-service/models"
)

func main() {
	UserLocationNorepeat()
}

func UserLocationNorepeat() {
	db.InitMysql()
	var userLocation []models.UserLocation
	db.Mysql.Raw("select * FROM user_location WHERE id NOT IN (SELECT MIN(id) FROM user_location GROUP BY user ) and user <> '0xd5f92fd92f8c7f9391513e3019d9441aaf5b2d9e'").Scan(&userLocation)

	for _, u := range userLocation {
		db.Mysql.Model(&models.UserLocation{}).Where("id=?", u.Id).Delete(models.UserLocation{})
	}
}
