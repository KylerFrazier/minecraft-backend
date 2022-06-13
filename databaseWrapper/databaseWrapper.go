package databaseWrapper

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseWrapper struct {
	client *mongo.Client
}

func (*DatabaseWrapper) GetUserHash(userName string) string {
	return ""
}

func (*DatabaseWrapper) GetPasswordHash(userName string) string {
	return ""
}

// Schema

// User {
//   id: ID!
//   hash: ID!
//   name: String!
// }

// Auth {
//   id: ID!
//   passwordHash: ID!
//   userHash: ID!
// }

// Passwords {
//   id: ID!
//   hash: ID!
// }

// Roles {
//   id: ID!
//   userHash: ID!
//   list: [String!]
// }
