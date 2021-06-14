package setup

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"os"
)

type ConfSetup struct {
	Session      *session.Session
	DynamoClient *dynamodb.DynamoDB
}

func envVarSetup() {

	log.Println("Loading Env Var")
	REGION := os.Getenv("FINDERIO_REGION")
	ID := os.Getenv("FINDERIO_ID")
	SECRET := os.Getenv("FINDERIO_SECRET")
	if REGION != "" && ID != "" && SECRET != "" {
		os.Setenv("AWS_REGION", REGION)
		os.Setenv("AWS_ACCESS_KEY_ID", ID)
		os.Setenv("AWS_SECRET_ACCESS_KEY", SECRET)
	}
}

func sessionSetup() *ConfSetup {

	log.Println("Starting sessions")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	return &ConfSetup{sess, svc}

}

func MainSetup() *ConfSetup {

	envVarSetup()
	confSetup := sessionSetup()
	log.Println("End of setup")
	return confSetup
}
