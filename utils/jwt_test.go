package utils

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func TestJWT(t *testing.T) {
	token, err := CreateJWT(1, "person")
	if err != nil {
		t.Error(fmt.Sprintf("Error creating JWT: %s", err.Error()))
	}
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)
	c.Request().Header.Set("Authorization", token)

	id, userType, err := VerifyJWT(c)
	if err != nil {
		t.Error(fmt.Sprintf("Error verifying JWT: %s", err.Error()))
	}
	if id == -1 {
		t.Error(fmt.Sprintf("JWT is not valid"))
	}
	if userType != "person" {
		t.Error(fmt.Sprintf("JWT is not valid"))
	}
}