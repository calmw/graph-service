package validate

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"graph-service/api/common/statecode"
	"graph-service/api/models/request"
	"io"
)

type UserLocation struct{}

func NewUserLocation() *UserLocation {
	return &UserLocation{}
}

func (v *UserLocation) GetLocationByUserAddress(c *gin.Context, req *request.GetLocationByUserAddress) int {

	//err := c.ShouldBind(req)
	err := c.BindQuery(req)
	fmt.Println(err)
	if err == io.EOF {
		return statecode.ParameterEmptyErr
	} else if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "user" && e.Tag() == "required" {
				return statecode.UserEmpty
			}
		}
		return statecode.CommonErrServerErr
	}

	return statecode.CommonSuccess
}
