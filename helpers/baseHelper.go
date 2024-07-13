package helpers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func Pagination(c *fiber.Ctx) (int, int) {
	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage", "10"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	startIndex := (page - 1) * recordPerPage

	return startIndex, recordPerPage
}
