# About

Simple quickstart google cloud message server using go programming language / golang, this is
the server if you need client apps visit [firebase github](https://github.com/firebase):

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
