package api

import (
	"net/http"
	"strconv"

	"github.com/gedelumbung/go-movie/model"
	"github.com/gedelumbung/go-movie/repository"
	"github.com/labstack/echo"
)

func (a *API) GetCategory(c echo.Context) error {
	var (
		category model.Category
		err      error
		id       int
	)
	errParams := map[string]string{}
	strId := c.Param("id")
	if id, err = strconv.Atoi(strId); err != nil {
		errParams["id"] = "invalid numeric value"
	}
	if len(errParams) > 0 {
		return c.JSON(http.StatusBadRequest, ErrRespond("client", "invalid request parameters", errParams))
	}

	category, err = a.db.Categories().FindByID(id)

	if err != nil {
		if err.Error() == repository.ErrNotFound.Error() {
			return c.JSON(http.StatusNotFound, ErrRespond("server", "record not found", err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, ErrRespondString("server", "unable to load data from source", err.Error()))
	}

	return c.JSON(http.StatusOK, OKRespond(category))
}
