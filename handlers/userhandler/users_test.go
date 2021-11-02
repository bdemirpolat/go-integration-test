package userhandler

import (
	"encoding/json"
	"github.com/bdemirpolat/integration-test/db"
	"github.com/bdemirpolat/integration-test/models"
	"github.com/bdemirpolat/integration-test/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"strings"
	"testing"
)

type UsersTestSuite struct {
	suite.Suite
	App          *fiber.App
}

func TestUsersTestSuite(t *testing.T) {
	suite.Run(t, &UsersTestSuite{})
}

func (s *UsersTestSuite) SetupSuite() {
	err := db.CreateDatabase()
	s.Nil(err)
	database := db.ConnectDB()
	_, err = db.CreateTable(database)
	s.Nil(err)
	userRepo := &repository.UserRepo{DB: database}
	userHandler := &UserHandler{Repo: userRepo}
	app := fiber.New()
	app.Post("/users", userHandler.Create)
	go func() {
		app.Listen(":3000")
	}()
	s.App = app
}

func (s *UsersTestSuite) TestCreateUser() {
	user := &models.User{Username: "burak"}
	userJson, err := json.Marshal(user)
	s.Nil(err)
	req := httptest.NewRequest("POST", "/users", strings.NewReader(string(userJson)))
	req.Header.Add("Content-Type", "application/json")
	res, err := s.App.Test(req)
	s.Nil(err)
	defer res.Body.Close()
	s.Nil(err)
	s.Equal(200, res.StatusCode)
}

func (s *UsersTestSuite) TearDownSuite() {
	err := db.DeleteDatabase()
	s.Nil(err)
}
