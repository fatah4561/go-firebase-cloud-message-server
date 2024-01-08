package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

func main() {
	opt := option.WithCredentialsFile("google-service-key.json")
	app, err := firebase.NewApp(context.Background(), &firebase.Config{}, opt)
	if err != nil {
		log.Fatalf("error creating new firebase app: %v\n", err)
	}

	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := "fsMLKzDRQRSX7ZCupOrMM5:APA91bHumlUCdRzW4RPVMP_bOn0-xmFS1aAEyrjsWHNcFkWk6V3bat6yaSomwbE04HGj7AjikU_O1WPUXf9k8XWzRaoZyECoETqAJSt77GmS0i6ddDl7aENRSl40NtIZBm0cSlZX7yy4"

	// These registration tokens come from the client FCM SDKs.
	registrationTokens := []string{
		registrationToken,
		// ...
		// "YOUR_REGISTRATION_TOKEN_n",
	}

	topic := "test"

	// Subscribe the devices corresponding to the registration tokens to the
	// topic.
	response, err := client.SubscribeToTopic(ctx, registrationTokens, topic)
	if err != nil {
		log.Fatalln(err)
	}
	// See the TopicManagementResponse reference documentation
	// for the contents of response.
	fmt.Println(response.SuccessCount, "tokens were subscribed successfully")

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"score":   "850",
			"time":    "2:45",
			"message": "hello world",
		},
		Token: registrationToken,
		// Topic: "test",
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	resSend, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", resSend)
}
