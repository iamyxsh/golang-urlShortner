package schema

import (
	mgm "github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type UrlReq struct {
	Url string `json: "url", bson: "url"`
}

type AdminLoginReq struct {
	Username string `json: "username", bson: "username"`
	Password string `json: "password", bson: "password"`
}

type UrlDoc struct {
	mgm.DefaultModel `bson:",inline"`
	Short string `json:"short" bson:"short"`
	Full string `json:"full" bson:"full"`
	Clicks int `json:"clicks" bson:"clicks"`
	Status bool `json:"status" bson:"status"`
}

type AdminDoc struct {
	mgm.DefaultModel `bson:",inline"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func NewUrl(short string,full string, clicks int) *UrlDoc {
   return &UrlDoc{
      Short: short,
	  Full: full,
	  Clicks: clicks,
   }
}

func FindByShort(s string) (error , *UrlDoc) {
	var urlDoc = &UrlDoc{}
	var coll = mgm.Coll(urlDoc)

	err := coll.First(bson.M{"short": s }, urlDoc)
	
	if err != nil {
		return err, nil
	}

	return nil, urlDoc

}

func FindAdmin(s string) (error , *AdminDoc) {
	var adminDoc= &AdminDoc{}
	var coll = mgm.Coll(adminDoc)

	err := coll.First(bson.M{"username": s }, adminDoc)
	
	if err != nil {
		return err, nil
	}

	return nil, adminDoc

}


