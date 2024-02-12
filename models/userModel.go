package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `bson:"firstName" validate:"required,min=2,max=100"`
	LastName     *string            `bson:"lastName" validate:"required,min=2,max=100"`
	Password     *string            `bson:"password" validate:"required,min=6"`
	Email        *string            `bson:"email" validate:"email,required"`
	Url          *string            `bson:"url"`
	Phone        *string            `bson:"phone" validate:"required"`
	Token        *string            `bson:"token"`
	RefreshToken *string            `bson:"refreshToken"`
	Created_at   time.Time          `bson:"createdAt"`
	Updated_at   time.Time          `bson:"updatedAt"`
	UserId       string             `bson:"userId"`
	UserType     *string            `bson:"userType" validate:"required,eq=ADMIN|eq=USER" `
}
