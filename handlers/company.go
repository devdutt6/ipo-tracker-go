package handlers

import (
	"context"
	"fmt"

	"github.com/devdutt6/ipo-tracker-go/api"
	"github.com/devdutt6/ipo-tracker-go/mongoutil"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCompanies(c *fiber.Ctx) error {
	ctx := context.TODO()

	var (
		err       error
		companies []api.CompanyDocument
		cursor    *mongo.Cursor
	)

	if cursor, err = mongoutil.CompanyCollection.Find(ctx, bson.M{}); err != nil {
		fmt.Println("failed to get companies", err)
		return c.Status(500).JSON(api.InternalErrorResponse)
	}

	if err = cursor.All(ctx, &companies); err != nil {
		fmt.Println("failed to get parse", err)
		return c.Status(500).JSON(api.InternalErrorResponse)
	}
	cursor.Close(ctx)

	if companies == nil {
		companies = []api.CompanyDocument{}
	}
	return c.Status(200).JSON(api.GetCompaniesResponse{Companies: companies})
}
