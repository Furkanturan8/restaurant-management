package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	helper "restaurant-management/helpers"
	"restaurant-management/models"
	"restaurant-management/services"
	"strconv"
	"time"
)

type UserHandler struct {
	Service   *services.UserService
	Validator *validator.Validate
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service, Validator: validator.New()}
}

func (uh *UserHandler) GetUsers(c *fiber.Ctx) error {
	startIndex, recordPerPage := helper.Pagination(c)

	users, total, err := uh.Service.GetUsers(startIndex, recordPerPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error occurred while listing users"})
	}

	return c.JSON(fiber.Map{
		"total_count": total,
		"users":       users,
	})
}

func (uh *UserHandler) GetUser(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz user ID"})
	}

	user, err := uh.Service.GetUserByID(int(uint(userID)))
	user.Password = nil
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"user": user})
}

func (uh *UserHandler) SignUp(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	validationErr := uh.Validator.Struct(user)
	if validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationErr.Error()})
	}

	count, err := uh.Service.CountByEmail(*user.Email)
	if err != nil || count > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Bu email zaten var"})
	}

	count, err = uh.Service.CountByPhone(*user.Phone)
	if err != nil || count > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Bu telefon numarası zaten var"})
	}

	password, err := HashPassword(*user.Password)
	if err != nil {
		return err
	}
	user.Password = &password
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := uh.Service.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "user oluşturulamadı!"})
	}

	return c.JSON(fiber.Map{"message": "user başarıyla oluşturuldu!"})
}

func (uh *UserHandler) Login(c *fiber.Ctx) error {
	var user models.User
	var foundUser models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := uh.Service.FindByEmail(*user.Email, &foundUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kullanıcı bulunamadı!"})
	}

	if !VerifyPassword(*user.Password, *foundUser.Password) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Geçersiz giriş bilgileri"})
	}

	token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, foundUser.UserID)
	if err := uh.Service.UpdateTokens(foundUser.UserID, token, refreshToken); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Tokenlar güncellenemedi"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24), // 24 saatlik geçerlilik süresi
		HTTPOnly: true,
		Secure:   false, // Sadece HTTPS üzerinden gönderilsin
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * 24 * 7), // 7 günlük geçerlilik süresi
		HTTPOnly: true,
		Secure:   false, // Sadece HTTPS üzerinden gönderilsin
	})

	foundUser.Password = nil

	return c.JSON(fiber.Map{
		"user":          foundUser,
		"token":         token,
		"refresh_token": refreshToken,
	})
}

func (uh *UserHandler) LogOut(c *fiber.Ctx) error {
	userID, _ := strconv.Atoi(c.Params("user_id"))
	user, err := uh.Service.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz user ID"})
	}

	fmt.Println("token:", *user.Token)
	fmt.Println("tokenRef:", *user.RefreshToken)
	fmt.Println("email:", *user.Email)

	err = uh.Service.ClearTokens(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Çıkış yapılamadı"})
	}

	// Token cookie'lerini sil
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Geçmiş bir süre sonu
		HTTPOnly: true,
		Secure:   false, // HTTPS gereksinimi yok
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Geçmiş bir süre sonu
		HTTPOnly: true,
		Secure:   false, // HTTPS gereksinimi yok
	})

	return c.JSON(fiber.Map{"message": "Başarıyla çıkış yapıldı"})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := err == nil

	return check
}
