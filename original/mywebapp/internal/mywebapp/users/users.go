package users

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

const creds = `{
  "type": "service_account",
  "project_id": "simprints-cloud-hiring",
  "private_key_id": "e39ce20f1a1a41c75caedb40a75751d4c836912d",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCXjrfbst2vqrf6\nocoMItYAxLROl87oHQOz4A5tT8AELfBXysKEYldlY3zj+8nljU4itdNwrBrJvXt/\n8dpPiXHsGwEXDOhNcGabrECM3o19EF/oOSiN2d5/zH64M10dw+q3gV1VPYmN0kYe\nBUjG26iLb4X3YXHqOSHXYkwk1eVVirOsDqCIqXRb9ECMlEgalpjHZhvgyGBhunlL\nKq4kh7ofo4ILADKCLzj3D+c61N0Ja+S6mNFrC3HDfb2Bk4jTFrM+Q+XXm995z/M9\n8Va72iTRHplVZTqKUJSjUpYwNfx9iXIyvHm4uf1MIgakbw8/zt/Hk0EvPCpEVsv8\nDnUDX7stAgMBAAECggEAHigjVH+AVhKJvkDQM98ApzF80UCqxovzbyjPphee1RYR\nA18Qmof47O9BbBnvYBJgcE6uznYodGkUOG626mmmdf0fl+cVGEeb2zEIvR8Rl0eZ\nVdt9wtAN0m8t7dgmmKQhwbbni+0NQF+NXS7e/ta5cFlWFlXY9Wq93Z4tGY1IL192\nCraQZO7ssMUu9gBGX1a9Cucu+xsYUWaJGHSTbYxPsAMMlLI9lflyqko44NQyaY5i\nyrwyGR5RRNwUPqECwLfZwr8459I7UGt0j8jJURSboCMcN/WfhEW+42iRqpOwz0hi\n+rSrQpV1GUyCMJy7S8DnE2bNO1cPEiy/VkN1a9tsnwKBgQDOymNTFanbtfoOAbaq\nQZZ2ZMfbjkvubctmtwm8PXkeb0Xp2xt5Nc2dUMYKPQ+ULmF5eLaY0IZ3k+eHw1+2\nbHl2bDb8jOqksUvIuGqyVwTjtGmt6jKTWePg4NGuS2h4hfydZ/oi+4Yi06rOLok1\ngYtmWYHYe1T5Ei1wVapAgs9dVwKBgQC7n4yi9mmzQSn+p1iKHbocwhR7bQknI35D\nIu6fbf8z8Rg8Xgy/BsTAnD/ZGnNFTqcKDSUUCM8TyAuVIGv/rPNZm5CD+XxAl1M8\nXxTRm2Y3bd34novWfhIK2bBSytVmTgqGOfJyFI6Qux8EnL/bNL5UVLd+5x10VHNv\nmFHlObVVGwKBgFCRbg1q/VkLF7fpVwASucq7DVsvn8nvoTYNzo827D+9XuH0aknZ\nepj2ZZKLo8w6HeQz8gCsDhf82lJ8/oeWz+Qh2XKgfLNaGIrFv1xopxEIs7v86WOb\nvDtbnJZp7vQ2T+wMHNsQYHylN4FJhFACfwuLiBmOQ9cfC/Eej0y/KofVAoGBAKhb\n6x7wbrjjaIOSNuj48+CKZ0a5+NSeDNbBqc52LeK2HUFOQ7HGqRcNHl36ViWDBnEG\nipcHsInXoE0Y4tByqYtnBK7oWI69O3uY9l2ATKDAfcZyvpiiWB5UaNxiFkvPC7KW\nnO+xcpx8zMCex6o0EfYqbt35FX0lqtAgppxCxGyxAoGAaa0C3rBZqEV+d+H/WOQP\nOqRl8j6HiNTJyGACSkPh/CnQubw6++MGrCPPlNluASJeBbNC6WpX8qYSxEEFZ8Wl\noIvn1jWvRTTi3pAageqnO13kAz5yLPae7LTvxMqMLWioywUcob828qebB+id1AGT\nWJnOV0oL2rGPwsAGolnJ00c=\n-----END PRIVATE KEY-----\n",
  "client_email": "senior-backend-swe@simprints-cloud-hiring.iam.gserviceaccount.com",
  "client_id": "114245278424750935158",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/senior-backend-swe%40simprints-cloud-hiring.iam.gserviceaccount.com"
}`

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, userId string) (User, error)
	DeleteUser(ctx context.Context, userId string) error
}

type repository struct {
	firestoreClient *firestore.Client
}

func NewRepository(ctx context.Context) (Repository, error) {
	creds, err := google.CredentialsFromJSON(ctx, []byte(creds), "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return nil, err
	}
	project := os.Getenv("GCLOUD_PROJECT")
	if project == "" {
		return nil, fmt.Errorf("missing required environment variable GCLOUD_PROJECT")
	}
	firestoreClient, err := firestore.NewClient(ctx, project, option.WithCredentials(creds))
	if err != nil {
		return nil, err
	}
	repo := &repository{
		firestoreClient: firestoreClient,
	}
	return repo, nil
}

func (r *repository) CreateUser(ctx context.Context, user User) error {
	_, err := r.firestoreClient.Collection("users").Doc(user.Id).Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUser(ctx context.Context, userId string) (User, error) {
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

func (r *repository) DeleteUser(ctx context.Context, userId string) error {
	_, err := r.firestoreClient.Collection("users").Doc(userId).Delete(ctx, firestore.Exists)
	if err != nil {
		return err
	}
	return nil
}
