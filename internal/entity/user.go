package entity

// import "github.com/google/uuid"

type User struct {
	ID string `bson:"_id,omitempty"`
	Name string `bson:"name" binding:"required"`
	Password string `bson:"password" binding:"required"`
}