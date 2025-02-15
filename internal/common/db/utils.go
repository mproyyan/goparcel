package db

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertToObjectId(hex string) (primitive.ObjectID, error) {
	if hex == "" {
		return primitive.NilObjectID, nil
	}

	return primitive.ObjectIDFromHex(hex)
}
