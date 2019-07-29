package user

import (
	"fmt"
)

type firestoreError struct {
	err string //error description
}

func (e *firestoreError) Error() string {
	return fmt.Sprintf("Error %s", e.err)
}

// func GetUserDocumentRef(
// 	ctx context.Context,
// 	fclient firestore.Client,
// 	methodName string,
// 	requestUser User) (User, error) {

// 	userCopy := User{}
// 	err = nil

// 	switch methodName {
// 	case "POST":
// 		_, err := fclient.Collection("users").Doc(requestUser.Id).Create(ctx, requestUser)

// 		if err != nil {
// 			return userCopy, err
// 		}

// 	case "GET":
// 		snapshot, err := fclient.Collection("users").Doc(requestUser.Id).Get(ctx)

// 		if err != nil {
// 			return userCopy, err
// 		}
// 		err = snapshot.DataTo(&userCopy)

// 		if err != nil {
// 			return userCopy, err
// 		}

// 	case "DELETE":
// 		_, err := fclient.Collection("users").Doc(requestUser.Id).Delete(ctx)

// 		if err != nil {
// 			return userCopy, err
// 		}

// 	default:
// 		return User{}, &firestoreError{err: "negative error"}
// 	}

// 	return userCopy, nil
// }
