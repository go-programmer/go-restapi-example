package user

import (
	"context"
	"mywebapp/v9/utilities"
)

var fsClientConnection firestoreClientConnection

func SetClient(fsClientConnection *firestoreClientConnection) {
	fsClientConnection = fsClientConnection.get
}

func (dbClient *IUserRepository) CreateUser(ctx context.Context, user User) error {

	__, err := utilities.GetFC().firestoreClient.Collection("users").Doc(user.Id).Create(ctx, user)

	if err != nil {
		return err
	}
	return nil
}

// func (r *repository) GetUser(ctx context.Context, userId string) (User, error) {
// 	user := User{}
// 	snapshot, err := r.firestoreClient.Collection("users").Doc(userId).Get(ctx)
// 	if err != nil {
// 		return user, err
// 	}
// 	err = snapshot.DataTo(&user)
// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

// func (r *repository) DeleteUser(ctx context.Context, userId string) error {
// 	_, err := r.firestoreClient.Collection("users").Doc(userId).Delete(ctx, firestore.Exists)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
