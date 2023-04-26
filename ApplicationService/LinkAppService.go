package applicationservice

import (
	"github.com/labstack/gommon/random"
)

type ILinkApplicationService interface {
	SetLink(newLink CreateLink) LinkDto
	GetLink(shortLink string) string
}

func GetAppLicationService(config uint8) ILinkApplicationService {
	da := LinkApplicationService{}
	return &da
}

type LinkApplicationService struct {
	cn LinkContext
}

func makeHash() string {
	const ALPHABAT = "0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	return random.String(5, ALPHABAT)
}

func (T *LinkApplicationService) SetLink(newLink CreateLink) LinkDto {
	T.cn.Init()
	shortLink := makeHash()
	for {
		check := T.cn.Get(shortLink)
		if check == "" {
			break
		}
		shortLink = makeHash()
	}

	_newLink := Link{Link: newLink.Link, ShortLink: shortLink}
	T.cn.Set(_newLink)
	return LinkDto{Link: _newLink.Link, ShortLink: shortLink}
}

func (T *LinkApplicationService) GetLink(shortLink string) string {

	T.cn.Init()
	return T.cn.Get(shortLink)
}
