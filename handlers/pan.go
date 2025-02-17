package handlers

import (
	"context"
	"errors"
	"fmt"

	"github.com/devdutt6/ipo-tracker-go/api"
	"github.com/devdutt6/ipo-tracker-go/mongoutil"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPanHandler(c *fiber.Ctx) error {
	ctx := context.TODO()

	var (
		userId = c.Locals("userId").(primitive.ObjectID)
		err    error
		pans   []api.PanDocument
		cursor *mongo.Cursor
	)

	if cursor, err = mongoutil.PanCollection.Find(ctx, bson.M{"userId": userId}); err != nil {
		fmt.Println("failed to get pans", err)
		return c.Status(500).JSON(api.InternalErrorResponse)
	}

	if err = cursor.All(ctx, &pans); err != nil {
		fmt.Println("failed to get parse", err)
		return c.Status(500).JSON(api.InternalErrorResponse)
	}
	cursor.Close(ctx)

	if pans == nil {
		pans = []api.PanDocument{}
	}
	return c.Status(200).JSON(api.GetPanResponse{Pans: pans})
}

func AddPanHandler(c *fiber.Ctx) error {
	var newPan = new(api.PanRequestBody)

	if err := c.BodyParser(newPan); err != nil {
		fmt.Println("failed parsing body for create pan api")
		return c.Status(400).JSON(api.NewErrorResponse("failed parsing body"))
	}

	var (
		existingPan api.PanDocument
		err         error
		ctx         = context.TODO()
	)

	if err = mongoutil.PanCollection.FindOne(ctx, bson.M{
		"panNumber": newPan.PanNumber,
	}).Decode(&existingPan); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			panDocument := api.PanDocument{
				UserId:    c.Locals("userId").(primitive.ObjectID),
				PanNumber: newPan.PanNumber,
			}
			if _, err = mongoutil.PanCollection.InsertOne(ctx, panDocument); err != nil {
				fmt.Println("failed adding pan to db")
				return c.Status(500).JSON(api.InternalErrorResponse)
			}

			return c.Status(201).JSON(api.PanResponse{Message: "pan added successfully"})
		} else {
			fmt.Println("Failed adding pan to db")
			return c.Status(500).JSON(api.InternalErrorResponse)
		}
	}

	return c.Status(409).JSON(fiber.Map{"message": "pan already in use"})
}

func DeletePanHandler(c *fiber.Ctx) error {
	var newPan = new(api.PanRequestBody)

	if err := c.BodyParser(newPan); err != nil {
		fmt.Println("failed parsing body for delete pan api")
		return c.Status(400).JSON(api.NewErrorResponse("failed parsing body"))
	}

	var (
		userId      = c.Locals("userId").(primitive.ObjectID)
		existingPan api.PanDocument
		err         error
	)

	ctx := context.TODO()

	if err = mongoutil.PanCollection.FindOne(ctx, bson.M{
		"panNumber": newPan.PanNumber,
		"userId":    c.Locals("userId"),
	}).Decode(&existingPan); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(404).JSON(api.NewErrorResponse("no such pan registered"))
		} else {
			fmt.Println("failed adding pan to db")
			return c.Status(500).JSON(api.InternalErrorResponse)
		}
	}

	if _, err = mongoutil.PanCollection.DeleteOne(ctx, bson.M{
		"panNumber": newPan.PanNumber,
		"userId":    userId,
	}); err != nil {
		fmt.Println("failed deleting pan from db")
		return c.Status(500).JSON(api.InternalErrorResponse)
	}

	return c.Status(200).JSON(api.PanResponse{Message: "pan deleted successfully"})
}
