package dataaccess

import (
	"go.mongodb.org/mongo-driver/mongo"

	"context"
)

// global vars
var myClient *mongo.Client
var myContext *context.Context
