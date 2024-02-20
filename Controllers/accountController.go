package controllers

import (
	"database/sql"
	"net/http"
	Models "shidqi/WebGo/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AccountAPIController struct {
	DB *sql.DB
}

func NewAccountAPIController(db *sql.DB) *AccountAPIController {
	return &AccountAPIController{DB: db}
}

func (controller *AccountAPIController) ServeLoginPage(c *fiber.Ctx) error {
	return c.SendFile("Views/index.html")
}

type LoginForm struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

func (controller *AccountAPIController) LoginUser(c *fiber.Ctx) error {
	var form LoginForm
	if err := c.BodyParser(&form); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var account Models.Account
	if err := controller.DB.QueryRow("SELECT password FROM accounts WHERE username = ?", form.Username).Scan(&account.Password); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(form.Password)); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid password"})
	}

	return c.JSON(fiber.Map{"message": "Logged in"})
}
