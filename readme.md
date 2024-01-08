# About

Simple quickstart google firebase cloud message server using go programming language / golang, **this is
the server** (to send the messages) if you need client apps visit [firebase github](https://github.com/firebase):

- [Android firebase client apps](https://github.com/firebase/quickstart-android)
- [Web js firebase client apps](https://github.com/firebase/quickstart-js)
- [Ios firebase client apps](https://github.com/firebase/quickstart-ios)

## Setup

Before starting replace the "google-service-key.json" to correct path private key file in main.go

```go
// REPLACE "google-service-key.json" to correct path private key file
// visit https://firebase.google.com/docs/cloud-messaging/auth-server#provide-credentials-manually
opt := option.WithCredentialsFile("google-service-key.json")
```

Also replace the registration token "CLIENT_REGISTRATION_TOKEN" with yours registration token in main.go

```go
// This registration token comes from the client FCM SDKs. Check about section in readme if you need client apps
registrationToken := "CLIENT_REGISTRATION_TOKEN"
```

If you want to send message to multiple clients / topcis, add tokens here in main.go

```go
// These registration tokens come from the client FCM SDKs. Check about section in readme if you need client apps
registrationTokens := []string{
    registrationToken,
    // ...
    // "YOUR_REGISTRATION_TOKEN_n",
}
```

For more info on sending messages (individual clients, multiple clients, or topic) check [firebase docs](https://firebase.google.com/docs/cloud-messaging/send-message)

## Run / Build

Run

```bash
go run main.go
```

Build

```bash
go build main.go
```

## Docs

 [Google cloud message documentation](https://firebase.google.com/docs/cloud-messaging/).
