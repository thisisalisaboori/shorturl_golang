package applicationservice

import (
	"github.com/labstack/gommon/random"
)

type ILinkApplicationService interface {
	SetLink(newLink CreateLink) LinkDto
	GetLink(shortLink string) string
}
type LinkApplicationService struct {
	cn LinkContext
}

func makeHash() string {
	const ALPHABAT = "0123456789qwertyuiopasdfghjklzxcvbnm"
	return random.String(5, ALPHABAT)
}

func (self *LinkApplicationService) SetLink(newLink CreateLink) LinkDto {
	shortLink := makeHash()
	self.cn.Init()
	_newLink := Link{Link: newLink.Link, ShortLink: shortLink}
	self.cn.Set(_newLink)
	return LinkDto{Link: _newLink.Link, ShortLink: shortLink}
}

func (self *LinkApplicationService) GetLink(shortLink string) string {

	self.cn.Init()
	return self.cn.Get(shortLink)
}
