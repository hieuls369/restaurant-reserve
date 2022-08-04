package user_reserve

import (
	"context"
	"example/restaurant-reserved/internal/providers"
	reserve_model "example/restaurant-reserved/internal/share-domain/reserve/reserve-model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var reserveCollection *mongo.Collection = providers.GetCollection(providers.DB, "reservation")

func GetReserves(ctx context.Context) (*mongo.Cursor, error) {

	cursor, err := reserveCollection.Find(ctx, bson.M{})

	// find all collection data
	return cursor, err
}

func CreateReserve(ctx context.Context, reserve reserve_model.ReserveModel) (*mongo.InsertOneResult, error) {

	result, err := reserveCollection.InsertOne(ctx, reserve)
	fmt.Printf("Reserve added: %v", result)
	return result, err
}
func UpdateReserve(ctx context.Context, id string, reserve reserve_model.ReserveModel) (*mongo.UpdateResult, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{"amountpeople": reserve.AmountPeople, "phonenumber": reserve.PhoneNumber, "date": reserve.Date, "tableid": reserve.TableID}
	result, err := reserveCollection.UpdateOne(
		ctx,
		bson.M{"_id": objId}, bson.M{"$set": update},
	)
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return result, err
}

func DeleteReserve(ctx context.Context, id string, phoneNumber string) (*mongo.DeleteResult, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	result, err := reserveCollection.DeleteOne(ctx, bson.M{"_id": objId})
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	return result, err
}
