package config

import (
	"budget4home/src/expense"
	"budget4home/src/group"
	"budget4home/src/label"
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDb() *gorm.DB {
	if len(os.Getenv("DATABASE")) == 0 {
		// set default connection string
		os.Setenv("DATABASE", "host=localhost port=5432 database=postgres user=postgres password=postgres sslmode=disable")
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE")), &gorm.Config{})
	if err != nil {
		fmt.Println("Error starting db")
		return nil
	}

	db.AutoMigrate(&label.Label{})
	db.AutoMigrate(&expense.Expense{})
	db.AutoMigrate(&group.Group{})
	db.AutoMigrate(&group.GroupUser{})

	return db
}

func SetupFirebase() (*firebase.App, *auth.Client) {
	if len(os.Getenv("FIREBASE_ADMIN")) == 0 {
		// set default connection string
		os.Setenv(
			"FIREBASE_ADMIN",
			"{\n  \"type\": \"service_account\",\n  \"project_id\": \"lfmachadodasilva-dev\",\n  \"private_key_id\": \"07b395ede330ee05c8fd597a79aa43e6721fd90e\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDYN1MH9VMtBq6W\\nLYhvX67wTtrgkgCmdPdJbN7DFsoEdwqx8Kp8QaQCSqMYsqksJsg0F8zsd3e7BWdp\\nqgvVuKr1WSchAxZDZfvnbyOgWv2cSqbsAKpQdpQ8Avuh5KqWTuWc3oAMGiQJjPhk\\nc4mV79vaki3rIKIDAf8S3QK5zD5F+vnEwk5naQzi7GvUlHMOzb/vgpUS0ntEKtL+\\nLNolAqzotEx2mo9H60QzdnSH9rwCYIp26+KVQ2FrVtIJg1gt30cSSxD5Od7cZy69\\nH96Yee6qPK7faqqmIGtN4bJsoS6DBpp+LNxkIlMfH7aVNJZxMvC7IZJ0mQ6JxRop\\nI8yQWR8VAgMBAAECggEAPLldAnEuCQlFYzwqg00WLruW0VdwA0/6W47OCXCPEC32\\nvP7ggsFUtKgNolQbGyFRPQAqN4deAxKDdrOhW0bgDMQbLzRUvl0YIGmgUdr8Ozbk\\nJogRTUPgYfJchZ1ZI5nI9wOuZre4w03qPRwN2iRhVMOVTeH+XmXjI4Mazt/D5ZcR\\nArqftI8L8XrBFihI2ZcICXJbtiLDI10/LgvoCX4Mx4d0pzPh6rlpoM6s6VhjXrSW\\nNgY1Ov32shzW94kjQrFWSZRXdIDr0HqNa8/RYNkhDgMoefkiGsqkz7iiQd1jI2Nl\\n8XHr92wLEDyZxz5WmJbn/xx6Rn/XUZi9xXYq6b46CQKBgQDtU6idOWvFdjb6kur7\\n9EYBIF4gWbMNs/woqnIoNBMvxGnsCZ3G/0d9uTrAxjfxfHLyeW0QS0iuFoD1o2ew\\nrVKIqcN4uNjnJ5FuFWwcAj95AaX1/mV0DbmaEYBV/nIbBVxtXtWroqX5+Tv5CAuq\\non8EGCV6CQ3vbZ7adD/0f+kMuwKBgQDpOnHr8MgO3YZePqqemwTJvA3vlYKpT/mq\\nllX/DuvJD2ysKY0xVxYK7p2J3TXdpWPGhQOaOJO57EcWyBP+WWcmPIrOnzrFOhLI\\n1AZ4qvHqB3EbVeKUepJF5TWiDKmRfb2hMOw+AX4YDWMyoqko6JRfKLjtAw+k8gia\\nzfQ5F1kubwKBgQDok2CdW6Va6KOuYgY4jcWA2xiDOYR4PFcz/v1KmuXmnOR0tWPS\\nnV+RPNHwExDF7gCz0P1px21ddZ5BmzZFdUV5umxeRUADH8qsh//fXvCXBF+AuCLy\\nXXTzII2VOHrYMiPG5vYccpOXaoE0ZocQXJh/Ca7IblEiv1m6mcwHLsbTewKBgQDL\\nxOeBt9ZivoNGVkQl3NTbGLWoTJ1jJl/A/iZWeQfim9pbtNYKdMbqeD4mFKKaa4T4\\nDoDuPl47Q8d8jQSC5kOr8ZtpGU90v99nnW+l/9zqluPoeSpEha7E6JTqZ1vDpPOI\\nIMFXD2DGzPONavaWqlXvqhmf4lvJxjlkX2rwcojVewKBgQClGC1I8eIgR1l+GQGN\\nud+L9dBrL2P03dnCGWqkw8zuksPYtxcIEuY7GOCC6qKRsy1Kt4xUjmS1vzzQMsdH\\ncGsOQCFGvGS2rJlAY4RpG6esQCdLyT9CEO7AHe4k2rwA9j3j/iUiIwydwuIxXMwv\\n2yH71eyEzp8Nzm4SNlOUkT6n3Q==\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"firebase-adminsdk-161wz@lfmachadodasilva-dev.iam.gserviceaccount.com\",\n  \"client_id\": \"107587945077183294620\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-161wz%40lfmachadodasilva-dev.iam.gserviceaccount.com\"\n}\n")
	}

	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_ADMIN")))
	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}
	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic("Firebase load error" + err.Error())
	}
	return app, auth
}
