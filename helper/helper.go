package helper

import (
	"context"
	"time"

	"github.com/dchest/uniuri"
	jwt "github.com/form3tech-oss/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Context() (context.Context,context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	return ctx, cancel
}

func ShortUrl() string {
	short := uniuri.NewLen(5)
	return short
}

func CreateToken(userid primitive.ObjectID) (string, error) {
  var err error
  secretKey := "superSecretKey" // This should be in env but for this project, I am hardcoding it
  atClaims := jwt.MapClaims{}
  atClaims["authorized"] = true
  atClaims["user_id"] = userid
  atClaims["exp"] = time.Now().Add(10 * time.Hour).Unix()
  at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  token, err := at.SignedString([]byte(secretKey))
  if err != nil {
     return "", err
  }
  return token, nil
}