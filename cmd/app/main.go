package main

import (
	"hiyoko-fiver/pkg/logging/file"
	"hiyoko-fiver/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gofiber/fiber/v3"
)

const (
	envRoot = "./cmd/app"
	logDir  = "./log/app"
)

func init() {
	logger.SetLogDir(logDir)
	logger.Initialize()
	util.LoadEnv(envRoot)
}

func main() {
	app := fiber.New()

	app.Get("/", home)

	app.Post("/login", login)

	err := app.Listen(":8080")
	if err != nil {
		return
	}
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Tokens struct {
	Id      string
	Access  string
	Refresh string
}

func home(c fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func login(c fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := &User{
		Email:    email,
		Password: password,
	}

	tokens := cognitoAuth(*user)

	return c.JSON(tokens)
}

func cognitoAuth(user User) Tokens {
	params := &cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow: aws.String(cognitoidentityprovider.AuthFlowTypeAdminUserPasswordAuth),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(user.Email),
			"PASSWORD": aws.String(user.Password),
		},
		ClientId:   aws.String(util.Env("AWS_COGNITO_USER_CLIENT_ID").GetString("")),
		UserPoolId: aws.String(util.Env("AWS_COGNITO_USER_POOL_ID").GetString("")),
	}

	awsConfig := &aws.Config{
		Endpoint: aws.String(util.Env("AWS_COGNITO_ENDPOINT_URL").GetString("")),
		Region:   aws.String(util.Env("AWS_REGION").GetString("")),
	}

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		logger.Fatal("Create sess error", err)
	}

	client := cognitoidentityprovider.New(sess)

	res, err := client.AdminInitiateAuth(params)
	if err != nil {
		logger.Fatal("AdminAuth Error", err)
	}
	if res == nil || res.AuthenticationResult == nil || res.AuthenticationResult.IdToken == nil {
		logger.Fatal("failed to login")
	}

	var tokens Tokens

	tokens.Id = *res.AuthenticationResult.IdToken
	tokens.Access = *res.AuthenticationResult.AccessToken
	tokens.Refresh = *res.AuthenticationResult.RefreshToken

	return tokens
}
