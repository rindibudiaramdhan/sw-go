package firebase

import (
	"context"
	"encoding/base64"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type Config struct {
	ParentCollection    string
	ParentDocument      string
	SubParentCollection string
	Collection          string
	DocName             string
}

func GetDecodedFireBaseKey(key string) ([]byte, error) {
	fireBaseAuthKey := key
	decodedKey, err := base64.StdEncoding.DecodeString(fireBaseAuthKey)
	if err != nil {
		return nil, err
	}

	return decodedKey, nil
}

func Initialize(key string) (*firebase.App, error) {
	decodedKey, err := GetDecodedFireBaseKey(key)
	if err != nil {
		return nil, err
	}

	opts := []option.ClientOption{option.WithCredentialsJSON(decodedKey)}

	// Initialize firebase app
	app, err := firebase.NewApp(context.Background(), nil, opts...)
	if err != nil {
		return nil, err
	}

	return app, nil
}

// Send Push notifications to a single device
func SendSingleNotification(key string, messaging *messaging.Message) error {
	app, err := Initialize(key)
	if err != nil {
		return err
	}
	fcmClient, err := app.Messaging(context.Background())
	if err != nil {
		return err
	}

	_, err = fcmClient.Send(context.Background(), messaging)
	if err != nil {
		return err
	}

	return nil
}

func AddData(key string, data interface{}, cf Config) (*firestore.DocumentRef, error) {
	app, err := Initialize(key)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Firestore client: %v", err)
	}
	defer client.Close()

	// Create a reference to the parent collection
	parentDocRef := client.Collection(cf.ParentCollection).Doc(cf.ParentDocument)

	// Create a reference to the new subcollection
	subParentDocRef := parentDocRef.Collection(cf.SubParentCollection).Doc(cf.DocName)

	// Create a new document in the subcollection
	doc, _, err := subParentDocRef.Collection(cf.Collection).Add(context.Background(), data)
	if err != nil {
		return nil, fmt.Errorf("failed to create subdocument: %v", err)
	}

	return doc, err
}
