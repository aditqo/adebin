package endpoints

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Fetch(c echo.Context) error {

	p := c.QueryParams()

	id := p.Get("id")
	if len(id) != 0 {
		result, err := h.Store.Fetch("id:" + p.Get("id"))
		if err != nil {
			log.Println(err)
		}

		return c.String(http.StatusOK, result)
	}

	return c.String(http.StatusNotFound, "Couldn't find note with this ID")

	
}