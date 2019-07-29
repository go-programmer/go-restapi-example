package utilities

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
)

var fc *firestoreClientConnection

type firestoreClientConnection struct {
	fsClientRef *firestore.Client
}

func NewFireStoreClient(ctx context.Context) {

	// project := os.Getenv("GCLOUD_PROJECT", "~/Downloads/myapp-creds.json")
	// creds, err := google.Credentials("asdasd", "~/Downloads/myapp-creds.json")
	// creds, err := google.CredentialsFromJSON(ctx, []byte(creds), "https://www.googleapis.com/auth/cloud-platform")
	// if err != nil {
	// 	return nil, err
	// }

	// if project == "" {
	// 	return nil, fmt.Errorf("missing required environment variable GCLOUD_PROJECT")
	// }

	// firestoreClient, err := firestore.NewClient(ctx, project, option.WithCredentials(creds))

	// need to set this as the environment variable automatically
	// detects and reads the file.
	// export GOOGLE_APPLICATION_CREDENTIALS="~/Downloads/myapp-creds.json"

	projectID := os.Getenv("GCLOUD_PROJECT")

	// if projectID == "" {
	// 	return nil, fmt.Errorf("missing required environment variable GCLOUD_PROJECT")
	// }

	firestoreClient, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		// return nil, err
	}

	fc = &firestoreClientConnection{
		fsClientRef: firestoreClient,
	}

	// fc = client

	//return client, nil
}

func GetFC() *firestoreClientConnection {
	return fc
}
