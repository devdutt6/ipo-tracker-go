package handlers

import (
	"context"
	"errors"

	"github.com/devdutt6/ipo-tracker-go/api"
	"github.com/devdutt6/ipo-tracker-go/helper"
	"github.com/devdutt6/ipo-tracker-go/mongoutil"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Authenticate(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	bearer, ok := headers["Authorization"]

	if !ok || len(bearer) == 0 {
		return c.Status(401).JSON(api.NewErrorResponse("unauthorized"))
	}

	var (
		ctx         = context.TODO()
		tokenString = bearer[0]
		email       string
		err         error
	)
	tokenString = tokenString[len("Bearer "):]

	if email, err = helper.VerifyToken(tokenString); err != nil {
		return c.Status(401).JSON(api.NewErrorResponse("unauthorized"))
	}

	user := new(api.UserDocument)

	if err := mongoutil.UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(401).JSON(api.NewErrorResponse("unauthorized login again"))
		} else {
			return c.Status(500).JSON(api.InternalErrorResponse)
		}
	}

	c.Locals("userId", user.Id)
	return c.Next()
}
