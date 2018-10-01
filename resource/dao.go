package resource

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Dao interacts with mongodb
type Dao struct{}

// Server uri
const Server = "localhost:27017"

// DbName is the database name
const DbName = "golang-rest-service-db"

// CollectionName is collection name
const CollectionName = "resources"

// GetResources returns the list of Resources
func (d Dao) GetResources() Resources {
	session, err := mgo.Dial(Server)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DbName).C(CollectionName)
	results := Resources{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return results
}

// CreateResource creates a resource
func (d Dao) CreateResource(r Resource) bool {
	session, err := mgo.Dial(Server)
	defer session.Close()
	r.ID = bson.NewObjectId()
	session.DB(DbName).C(CollectionName).Insert(r)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdateResource ...
func (d Dao) UpdateResource(r Resource) bool {
	session, err := mgo.Dial(Server)
	defer session.Close()
	session.DB(DbName).C(CollectionName).UpdateId(r.ID, r)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteResource ...
func (d Dao) DeleteResource(id string) string {
	session, err := mgo.Dial(Server)
	defer session.Close()
	// Verify id is ObjectId, otherwise fail
	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}
	// Grab id
	oid := bson.ObjectIdHex(id)
	// Remove user
	if err = session.DB(DbName).C(CollectionName).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	// Write status
	return "OK"
}
