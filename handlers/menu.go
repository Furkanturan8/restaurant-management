package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	helper "restaurant-management/helpers"
	"restaurant-management/models"
	"restaurant-management/services"
	"strconv"
)

type MenuHandler struct {
	Service   *services.MenuService
	Validator *validator.Validate
}

func NewMenuHandler(service *services.MenuService) *MenuHandler {
	return &MenuHandler{
		Service:   service,
		Validator: validator.New(),
	}
}

func (mh *MenuHandler) GetMenus(c *fiber.Ctx) error {
	startIndex, recordPerPage := helper.Pagination(c)

	menus, total, err := mh.Service.GetMenus(startIndex, recordPerPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error occurred while listing menu items"})
	}

	return c.JSON(fiber.Map{
		"total_count": total,
		"menu_items":  menus,
	})
}

func (mh *MenuHandler) GetMenu(c *fiber.Ctx) error {
	menuID, err := strconv.Atoi(c.Params("menu_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz menu ID"})
	}

	menu, err := mh.Service.GetMenuByID(int(uint(menuID)))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"menu": menu})
}

func (mh *MenuHandler) CreateMenu(c *fiber.Ctx) error {
	var menu models.Menu
	if err := c.BodyParser(&menu); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz JSON formatı"})
	}

	// Validation
	if err := mh.Validator.Struct(menu); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Kontrol: createdAt veya updatedAt alanları girilmişse hata ver
	if !menu.CreatedAt.IsZero() || !menu.UpdatedAt.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "createdAt veya updatedAt alanları manuel olarak doldurulamaz"})
	}

	err := mh.Service.CreateMenu(menu)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Menu oluşturulurken hata oluştu!"})
	}

	return c.SendString("Yeni Menu oluşturuldu!")
}

func (mh *MenuHandler) UpdateMenu(c *fiber.Ctx) error {
	menuID, err := strconv.Atoi(c.Params("menu_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz menu ID"})
	}

	var menu models.Menu
	if err := c.BodyParser(&menu); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	menu.MenuID = uint(menuID)
	err = mh.Service.UpdateMenu(menu)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Menu başarılı şekilde güncellendi!"})

}
