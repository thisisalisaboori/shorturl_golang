package applicationservice

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)
type IContext interface{
	Init() error
	Set(link Link ) bool
	Get(link string ) string
}
type Link struct{
	Link string
	ShortLink string
}

type LinkContext struct{
	 LDB  *leveldb.DB
}


func (self *LinkContext) Init() error{
	var err error
	self.LDB, err = leveldb.OpenFile("/tmp/links", nil)
    return err
}



func (self *LinkContext) Set(link Link ) bool{
	fmt.Println(self.LDB)
	err:= self.LDB.Put([]byte(link.ShortLink) , []byte(link.Link),nil  )
	defer self.LDB.Close()
	return err == nil
}

func (self *LinkContext) Get(link string ) string{
	data,_:= self.LDB.Get([]byte(link)  ,nil )
	defer self.LDB.Close()

 	return string(data)
}