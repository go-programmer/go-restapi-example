package users

import (
	"context"

	"cloud.google.com/go/firestore"
)

type usersRepository struct {
	firestoreClient *firestore.Client
}

func NewUsersRepository(ctx context.Context) (IUserRepository, error) {
	projectID := "simprints-cloud-hiring"

	firestore, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		return nil, err
	}
	repo := &usersRepository{
		firestoreClient: firestore,
	}
	return repo, err
}

func (r *usersRepository) CreateUser(ctx context.Context, user User) error {
	_, err := r.firestoreClient.Collection("users").Doc(user.Id).Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *usersRepository) GetUser(ctx context.Context, userId string) (User, error) {
	user := User{}
	snapshot, err := r.firestoreClient.Collection("users").Doc(userId).Get(ctx)
	if err != nil {
		return user, err
	}
	err = snapshot.DataTo(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *usersRepository) DeleteUser(ctx context.Context, userId string) error {
	_, err := r.firestoreClient.Collection("users").Doc(userId).Delete(ctx, firestore.Exists)

	if err != nil {
		return err
	}
	return nil
}
