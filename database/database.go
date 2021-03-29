package database

import (
	errHandling "urlShortner/error"

	mgm "github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB()  {
	err := mgm.SetDefaultConfig(nil, "URL-Shortner", options.Client().ApplyURI("mongodb://localhost:27017"))
	errHandling.HandleErr(err)
}

