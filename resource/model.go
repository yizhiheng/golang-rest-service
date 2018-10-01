package resource

import "github.com/globalsign/mgo/bson"

type Resource struct {
	ID     bson.ObjectId `bson:"_id"`
	Title  string        `json:"title"`
	Number int32         `json:"number"`
}

type Resources []Resource
