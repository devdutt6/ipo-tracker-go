package handlers

import (
	"context"

	"github.com/devdutt6/ipo-tracker-go/api"
	"github.com/devdutt6/ipo-tracker-go/mongoutil"
	"github.com/devdutt6/ipo-tracker-go/static"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckAllotmentHandler(c *fiber.Ctx) error {
	companyId := c.Params("companyId")

	cid, err := primitive.ObjectIDFromHex(companyId)
	if err != nil {
		return c.Status(400).JSON(api.NewErrorResponse("please provide a valid companyId"))
	}

	var (
		ctx        = context.TODO()
		company    api.CompanyDocument
		pans       []api.PanDocument
		pansCursor *mongo.Cursor
		userId     = c.Locals("userId").(primitive.ObjectID)
	)

	if err = mongoutil.CompanyCollection.FindOne(ctx, bson.M{"_id": cid}).Decode(&company); err != nil {
		return c.Status(404).JSON(api.NewErrorResponse("no such company found"))
	}

	if pansCursor, err = mongoutil.PanCollection.Find(ctx, bson.M{
		"userId": userId,
	}); err != nil {
		return c.Status(500).JSON(api.InternalErrorResponse)
	}

	if err = pansCursor.All(ctx, &pans); err != nil {
		return c.Status(500).JSON(api.InternalErrorResponse)
	}

	if pans == nil { // no pans to fetch status for
		return c.Status(200).JSON(fiber.Map{})
	}

	var response map[string]string

	switch company.Registrar {
	case static.REGISTRAR[static.CAMEO]:
		response = CheckWithCameo(&company, &pans)
	case static.REGISTRAR[static.BIGSHARE]:
		response = CheckWithBigShare(&company, &pans)
	case static.REGISTRAR[static.MAASHITLA]:
		response = CheckWithMaashitla(&company, &pans)
	case static.REGISTRAR[static.LINKINTIME]:
		response = CheckWithLinkintime(&company, &pans)
	default:
		response = map[string]string{}
	}

	return c.Status(200).JSON(response)
}
