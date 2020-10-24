package helpers

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// Database is global database variable
	Database *mongo.Database
	// URI is mongo connection string
	URI string = "mongodb+srv://egnite:Aqbfjotld9@cluster0-wtkg5.mongodb.net/egnite?retryWrites=true&w=majority"
)

func init() {

	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	Database = client.Database("egnite")
}

// GetRecord will fetch single record
func GetRecord(db string, filter map[string]interface{}) (map[string]interface{}, error) {
	var record map[string]interface{}
	collection := Database.Collection(db)
	documentReturned := collection.FindOne(context.TODO(), filter)
	err := documentReturned.Decode(&record)
	if err != nil {
		return record, err
	}
	return record, nil
}

// GetRecords will fetch multiple records
func GetRecords(db string, filter map[string]interface{}) ([]map[string]interface{}, error) {
	var records []map[string]interface{}
	collection := Database.Collection(db)
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return records, err
	}
	for cur.Next(context.TODO()) {
		var record map[string]interface{}
		_ = cur.Decode(&record)
		records = append(records, record)
	}
	return records, nil
}

// InsertRecord will insert single record
func InsertRecord(db string, record interface{}) error {
	collection := Database.Collection(db)
	_, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		return err
	}
	return nil
}

// InsertRecords will insert multiple records
func InsertRecords(db string, records []interface{}) error {
	collection := Database.Collection(db)
	_, err := collection.InsertMany(context.TODO(), records)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRecord will update single record
func UpdateRecord(db string, filter map[string]interface{}, update map[string]interface{}) error {
	collection := Database.Collection(db)
	_, err := collection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return err
	}
	return nil
}

// UpdateRecords will update multiple records
func UpdateRecords(db string, filter map[string]interface{}, update map[string]interface{}) error {
	collection := Database.Collection(db)
	_, err := collection.UpdateMany(context.TODO(), filter, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return err
	}
	return nil
}

// DeleteRecord will delete single record
func DeleteRecord(db string, filter interface{}) error {
	collection := Database.Collection(db)
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRecords will delete multiple records
func DeleteRecords(db string, filter map[string]interface{}) error {
	collection := Database.Collection(db)
	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
