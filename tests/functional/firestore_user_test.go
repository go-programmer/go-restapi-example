package firestore_user_test

import (
	"context"
	pUser "mywebapp/v9/user"
	pUtil "mywebapp/v9/utilities"
	"os"

	"testing"
)

func TestHello(t *testing.T) {
	userData := pUser.User{
		Id:    "11112",
		Email: "codeanit@gmail.com",
	}

	ctx := context.Background()

	// router := util.NewRouter(usermod.UserRoutes)

	projectID := os.Getenv("GCLOUD_PROJECT")

	// if projectID == "" {
	// 	return nil, fmt.Errorf("missing required environment variable GCLOUD_PROJECT")
	// }

	// firestoreClient, err := firestore.NewClient(ctx, projectID)
	// firestoreClientConn

	// get database connection type
	pUtil.NewFireStoreClient(ctx)
	pUser.SetClient(fsClientRef)

	// and set it to the user repository
	
	expected := "Hello Go!"

	actual, err := pUser.GetUserDocumentRef(ctx, firestoreClient, "POST", userData)

	if err != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
