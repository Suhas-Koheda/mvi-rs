package controllers

import (
	"net/http"

	"github.com/suhas-koheda/mvi-rs/internal/oauth"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http/response"
)

func authControllers() {
	app.POST("/auth/signup", signUpHandler)
	app.POST("/auth/signin", signInHandler)
}

func signUpHandler(ctx *gofr.Context) (interface{}, error) {
	var request struct {
		FirstName string `json:"FirstName"`
		LastName  string `json:"LastName"`
		Passowrd  string `json:"Password"`
		Email     string `json:"Email"`
		Role      string `json:"Role"`
	}
	if err := ctx.Bind(&request); err != nil {
		return response.Raw{
			Data: map[string]interface{}{
				"StatusCode": http.StatusBadRequest,
				"Body":       map[string]string{"error": "Invalid request body"},
			},
		}, nil
	}

	// TODO: to add the registerUser() function to add the user to databse
	token, err := oauth.GetJWT(request.Email, request.Role)
	if err != nil {
		return response.Raw{
			Data: map[string]interface{}{
				"statuscode": http.StatusUnauthorized,
				"body":       map[string]string{"error": "invalid credentials"},
			},
		}, nil
	}
	return response.Raw{
		Data: map[string]interface{}{
			"statuscode": http.StatusAccepted,
			"body": map[string]interface{}{
				"token":   token,
				"message": "Sign Up Successful",
			},
		},
	}, nil
}

func signInHandler(ctx *gofr.Context) (interface{}, error) {
	var request struct {
		Email    string `json:"Email"`
		Password string `json:"Password"`
		Role     string `json:"Role"`
	}
	if err := ctx.Bind(&request); err != nil {
		return response.Raw{
			Data: map[string]interface{}{
				"statuscode": http.StatusBadRequest,
				"body":       map[string]string{"error": "Invalid request body"},
			},
		}, nil
	}

	if !authenticateUser(request.Email, request.Password) {
		return response.Raw{
			Data: map[string]interface{}{
				"statuscode": http.StatusUnauthorized,
				"body":       map[string]string{"error": "invalid credentials"},
			},
		}, nil
	}

	token, err := oauth.GetJWT(request.Email, request.Role)
	if err != nil {
		return response.Raw{
			Data: map[string]interface{}{
				"statuscode": http.StatusInternalServerError,
				"body":       map[string]string{"error": "failed to generate token"},
			},
		}, nil
	}

	return response.Raw{
		Data: map[string]interface{}{
			"statuscode": http.StatusOK,
			"body": map[string]interface{}{
				"token":   token,
				"message": "Sign In Successful",
			},
		},
	}, nil
}

func authenticateUser(username string, password string) bool {
	// TODO: the logic to check if the user exists in the database or not
	if username == "test" {
		return password == "password"
	}
	return false
}

