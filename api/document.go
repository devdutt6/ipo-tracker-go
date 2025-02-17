package api

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserDocument struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
}

type CompanyDocument struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CompanyName string             `json:"companyName" bson:"companyName,omitempty"`
	CompanyCode string             `json:"companyCode" bson:"companyCode,omitempty"`
	Registrar   string             `json:"registrar" bson:"registrar,omitempty"`
}

type PanDocument struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId    primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	PanNumber string             `json:"panNumber" bson:"panNumber,omitempty"`
}
