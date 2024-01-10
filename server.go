package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type productType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageURL    string `json:"img_url"`
}

//	добавить в корзину товар изменять товар
//
// авторизация
// вынести в json
// разграничение прав доступа
// корзина есть
// нет удаления товара от имени админисратора
// в
var products = []productType{
	{
		Name:        "Квартира в новостройке",
		Description: "Жилой комплекс на границе города и природы.",
		Price:       10500000,
		ImageURL:    "/assets/img/a1.jpeg",
	},
	{
		Name:        "Квартира в новостройке",
		Description: "Это остров спокойствия для любой семьи.",
		Price:       10000000,
		ImageURL:    "/assets/img/a2.jpeg",
	},
	{
		Name:        "Квартира в новостройке",
		Description: "Место, где забываешь о суете.",
		Price:       15000000,
		ImageURL:    "/assets/img/a3.jpeg",
	},
	{
		Name:        "Частный дом",
		Description: "Удачное месторасположение",
		Price:       9555000,
		ImageURL:    "/assets/img/d4.jpg",
	},
	{
		Name:        "Частный дом",
		Description: "Место, где забываешь о суете.",
		Price:       8455000,
		ImageURL:    "/assets/img/d5.jpg",
	},
	{
		Name:        "Частный дом",
		Description: "Вся инфраструктура рядом",
		Price:       5966000,
		ImageURL:    "/assets/img/d6.jpg",
	},
	{
		Name:        "Квартира в новостройке",
		Description: "Удачное месторасположение",
		Price:       4595000,
		ImageURL:    "/assets/img/d7.jpg",
	},
	{
		Name:        "Квартира в новостройке",
		Description: "Остров спокойствия",
		Price:       6958000,
		ImageURL:    "/assets/img/d8.jpeg",
	},
}

const productPerPage = 3

func main() {
	e := echo.New()

	e.Static("/public", "public")
	e.Static("/assets", "public/assets")

	e.POST("/getProducts", func(c echo.Context) error {
		type requestData struct {
			Page int `json:"page"`
		}
		var request requestData
		err := c.Bind(&request)

		if err != nil {
			return err
		}

		startIndex := (request.Page - 1) * productPerPage
		endIndex := startIndex + productPerPage

		if endIndex > len(products) {
			endIndex = len(products)
		}

		return c.JSON(http.StatusOK, products[startIndex:endIndex])
	})

	e.POST("/getProductCount", func(c echo.Context) error {
		type responseResult = struct {
			Count           int `json:"count"`
			ProductsPerPage int `json:"products_per_page"`
		}

		result := responseResult{
			Count:           len(products),
			ProductsPerPage: productPerPage,
		}

		return c.JSON(http.StatusOK, result)
	})

	e.GET("*", func(c echo.Context) error {
		return c.File("index.html")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
