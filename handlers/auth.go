package handlers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/devdutt6/ipo-tracker-go/api"
	"github.com/devdutt6/ipo-tracker-go/helper"
	"github.com/devdutt6/ipo-tracker-go/mongoutil"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *fiber.Ctx) error {
	newUser := new(api.RegisterRequestBody)

	if err := c.BodyParser(newUser); err != nil {
		fmt.Println("failed parsing body for register api")
		return c.Status(400).JSON(api.NewErrorResponse("error parsing body"))
	}

	if newUser.Email == "" || newUser.Password == "" {
		return c.Status(400).JSON(api.NewErrorResponse("email and password are required"))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var (
		existingUser = new(api.UserDocument)
		err          error
	)

	if err = mongoutil.UserCollection.FindOne(ctx, bson.M{"email": newUser.Email}).Decode(&existingUser); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
			if err != nil {
				fmt.Println("password hashing failed")
				return c.Status(500).JSON(api.InternalErrorResponse)
			}

			newUser.Password = string(hash)
			if _, err = mongoutil.UserCollection.InsertOne(ctx, newUser); err != nil {
				fmt.Println("failed to register user")
				return c.Status(500).JSON(api.InternalErrorResponse)
			}

			return c.Status(200).JSON(api.RegisterResponse{Message: "user registered successfully"})
		} else {
			fmt.Println(err)
			return c.Status(500).JSON(api.InternalErrorResponse)
		}
	}

	return c.Status(409).JSON(api.NewErrorResponse("email already in use"))
}

func LoginHandler(c *fiber.Ctx) error {
	var userBody = new(api.LoginRequestBody)

	if err := c.BodyParser(userBody); err != nil {
		fmt.Println("failed parsing body for login api")
		return c.Status(400).JSON(api.NewErrorResponse("failed parsing request body"))
	}

	if userBody.Email == "" || userBody.Password == "" {
		return c.Status(400).JSON(api.NewErrorResponse("email and password are required"))
	}

	var (
		ctx  = context.TODO()
		user = new(api.UserDocument)
		err  error
	)

	if err = mongoutil.UserCollection.FindOne(ctx, bson.M{"email": userBody.Email}).Decode(user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(404).JSON(api.NewErrorResponse("email not registered"))
		}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userBody.Password)); err != nil {
		return c.Status(401).JSON(api.NewErrorResponse("email and password does not match"))
	}

	tokenString, err := helper.CreateToken(user.Email)
	if err != nil {
		fmt.Printf("error creating token %v", err)
		return c.Status(500).JSON(api.InternalErrorResponse)
	}

	return c.Status(200).JSON(api.LoginResponse{Token: tokenString})
}
