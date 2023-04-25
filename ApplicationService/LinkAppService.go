package applicationservice

import (
	"github.com/labstack/gommon/random"
)

type ILinkApplicationService interface{
	 SetLink() bool
	 GetLink() string

}

func makeHash() string{
	const ALPHABAT="0123456789qwertyuiopasdfghjklzxcvbnm"
 	return random.String(5,ALPHABAT)
}

func (self *Link  ) SetLink() bool{
	var cn LinkContext;
	self.ShortLink =makeHash()
	cn.Init()
	cn.Set(*self)
	return true
}


func (self *Link) GetLink( ) string{
	var cn LinkContext;
	cn.Init()
	self.Link = cn.Get(self.ShortLink)
	return self.Link
}
