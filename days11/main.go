package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/assets", "assets")
	e.GET("/index", Home)
	e.GET("/index2", formEmail)
	e.GET("/profile/:id", BlogDetail)
	e.GET("/project", MyProject)
	e.GET("/testimonial", Testimonial)
	e.POST("/addProject", addProject)

	e.Logger.Fatal(e.Start("localhost:500"))
}

func Home(c echo.Context) error {
	tme, error := template.ParseFiles("html/index.html")
	if error != nil {
		return c.JSON(http.StatusInternalServerError, error.Error())
	}
	return tme.Execute(c.Response(), nil)
}

func formEmail(c echo.Context) error {
	tme, error := template.ParseFiles("html/index2.html")
	if error != nil {
		return c.JSON(http.StatusInternalServerError, error.Error())
	}
	return tme.Execute(c.Response(), nil)
}

func MyProject(c echo.Context) error {
	tme, error := template.ParseFiles("html/project.html")
	if error != nil {
		return c.JSON(http.StatusInternalServerError, error.Error())
	}
	return tme.Execute(c.Response(), nil)
}

func Testimonial(c echo.Context) error {
	tme, error := template.ParseFiles("html/testimonial.html")
	if error != nil {
		return c.JSON(http.StatusInternalServerError, error.Error())
	}
	return tme.Execute(c.Response(), nil)
}
func BlogDetail(c echo.Context) error {
	id := c.Param("id")
	tme, error := template.ParseFiles("html/profile.html")
	if error != nil {
		return c.JSON(http.StatusInternalServerError, error.Error())
	}
	BlogDetail := map[string]interface{}{
		"id":           id,
		"technologies": "Technologies",
		"Duration":     "duration",
		"articel":      "hallo nama sya malim fajar",
	}
	return tme.Execute(c.Response(), BlogDetail)
}

func addProject(c echo.Context) error {
	nameProject := c.FormValue("PName")
	startDate := c.FormValue("FDate")
	LastDate := c.FormValue("LDate")
	TxtMsg := c.FormValue("TxtMsg")
	checkbox1 := c.FormValue("Cbx1")
	checkbox2 := c.FormValue("Cbx2")
	checkbox3 := c.FormValue("Cbx3")
	checkbox4 := c.FormValue("Cbx4")

	fmt.Println("nameProject", nameProject)
	fmt.Println("Start Date", startDate)
	fmt.Println("LastDate", LastDate)
	fmt.Println("TxtMsg", TxtMsg)
	fmt.Println("checkbox1", checkbox1)
	fmt.Println("checkbox2", checkbox2)
	fmt.Println("checkbox3", checkbox3)
	fmt.Println("checkbox4", checkbox4)

	return c.Redirect(http.StatusMovedPermanently, "/project")
}
