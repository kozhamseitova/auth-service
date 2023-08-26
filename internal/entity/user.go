package entity

// import "github.com/google/uuid"

type User struct {
	ID string `bson:"_id,omitempty" json:"id"`
	RefreshToken string `bson:"refresh_token" json:"refresh_token"`
}