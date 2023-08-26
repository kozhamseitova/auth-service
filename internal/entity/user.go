package entity

// import "github.com/google/uuid"

type User struct {
	ID string `bson:"_id,omitempty"`
	RefreshToken string `bson:"refresh_token"`
}