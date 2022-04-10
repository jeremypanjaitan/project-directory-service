package cloudengine

import (
	"context"
	"path/filepath"
	"pds-backend/apperror"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type FirebaseCloudEngineEntity interface {
	GetFirebaseAuth() *auth.Client
	GetFirebaseFirestore() *firestore.Client
	GetFirebaseFcm() *messaging.Client
}

type FirebaseCloudEngine struct {
	firebaseAuth      *auth.Client
	firebaseFirestore *firestore.Client
	firebaseFcm       *messaging.Client
}

func NewFirebaseCloudEngine() FirebaseCloudEngineEntity {
	serviceAccountKeyFilePath, err := filepath.Abs("serviceAccountKey.json")
	if err != nil {
		panic(apperror.ErrUnableToLoadServiceAccount)
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}

	firebaseAuth, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}
	firestoreClient, err := app.Firestore(context.Background())
	if err != nil {
		panic(err)
	}
	firebaseFcm, err := app.Messaging(context.Background())
	if err != nil {
		panic(err)
	}
	return &FirebaseCloudEngine{firebaseAuth: firebaseAuth, firebaseFirestore: firestoreClient, firebaseFcm: firebaseFcm}
}

func (f *FirebaseCloudEngine) GetFirebaseAuth() *auth.Client {
	return f.firebaseAuth
}
func (f *FirebaseCloudEngine) GetFirebaseFirestore() *firestore.Client {
	return f.firebaseFirestore
}
func (f *FirebaseCloudEngine) GetFirebaseFcm() *messaging.Client {
	return f.firebaseFcm
}
