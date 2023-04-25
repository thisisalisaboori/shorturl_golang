package Controllers

import (
	"encoding/json"
	//"fmt"
	applicationservice "shorturl/ApplicationService"

	"github.com/labstack/echo/v4"
)

func GetLink(c echo.Context) error {
	obj := applicationservice.Link{}
	obj.ShortLink = c.Param("link")
	obj.GetLink()
	return c.JSON(200, obj)
}

func SetLink(c echo.Context) error {
	var obj applicationservice.Link
	_ = json.NewDecoder(c.Request().Body).Decode(&obj)
	//fmt.Println(obj)
	obj.SetLink()

	return c.JSON(200, obj)
}
