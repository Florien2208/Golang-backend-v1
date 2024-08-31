// user.go
package main

import "go.mongodb.org/mongo-driver/bson/primitive"
// hello guy
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"` // Added the "username" field
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}
