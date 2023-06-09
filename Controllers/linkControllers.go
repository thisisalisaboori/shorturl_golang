package Controllers

import (
	"encoding/json"
	//"fmt"
	applicationservice "shorturl/ApplicationService"

	"github.com/labstack/echo/v4"
)

const DEFULTCONTEXT = 0

func GetLink(c echo.Context) error {
	da := applicationservice.GetAppLicationService(DEFULTCONTEXT)
	shortLink := c.Param("link")
	link := da.GetLink(shortLink)
	obj := applicationservice.LinkDto{Link: link, ShortLink: shortLink}
	return c.JSON(200, obj)
}

func SetLink(c echo.Context) error {
	da := applicationservice.GetAppLicationService(DEFULTCONTEXT)
	var newobj applicationservice.CreateLink
	_ = json.NewDecoder(c.Request().Body).Decode(&newobj)
	//fmt.Println(obj)
	res := da.SetLink(newobj)
	obj := applicationservice.LinkDto{Link: res.Link, ShortLink: res.ShortLink}
	return c.JSON(200, obj)
}
