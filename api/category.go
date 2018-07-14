package api

import (
	"net/http"
	"strconv"

	"github.com/gedelumbung/go-movie/helper"
	"github.com/gedelumbung/go-movie/model"
	"github.com/gedelumbung/go-movie/repository"
	"github.com/labstack/echo"
)

func (a *API) GetAllCategories(c echo.Context) error {
	var (
		categories         []model.Category
		strPage, strLimit  string
		page, limit, count int
		err                error
		errParams          map[string]string
	)
	errParams = map[string]string{}
	page = 1
	strPage = c.QueryParam("page")
	if len(strPage) > 0 {
		page, err = strconv.Atoi(strPage)
		if err != nil {
			errParams["page"] = "invalid numeric value"
		}
	}

	limit = 10
	strLimit = c.QueryParam("limit")
	if len(strLimit) > 0 {
		limit, err = strconv.Atoi(strLimit)
		if err != nil {
			errParams["limit"] = "invalid numeric value"
		}
	}

	if len(errParams) > 0 {
		return c.JSON(http.StatusBadRequest, helper.Abort("client", "invalid request parameters", errParams))
	}

	categories, count, err = a.db.Categories().All(page, limit)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Abort("server", "unable to load data from source", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.Response(categories, map[string]interface{}{
		"pagination": helper.NewPagination(count, page, limit),
	}))
}

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
		return c.JSON(http.StatusBadRequest, helper.Abort("client", "invalid request parameters", errParams))
	}

	category, err = a.db.Categories().FindByID(id)

	if err != nil {
		if err.Error() == repository.ErrNotFound.Error() {
			return c.JSON(http.StatusNotFound, helper.Abort("server", "record not found", err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.Abort("server", "unable to load data from source", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.Response(category, nil))
}
