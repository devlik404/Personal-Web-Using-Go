package main

import (
	"context"
	"html/template"
	"net/http"
	"project1/connect"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id                                         int
	Image, Title, Duration, Articel            string
	StartDate, EndDate                         time.Time
	CheckBox1, CheckBox2, CheckBox3, CheckBox4 bool
}

var MyBlogs = []Project{}

func main() {
	e := echo.New()
	connect.DbConection()
	e.Static("/assets", "assets")
	e.GET("/", Home)
	e.GET("/index2", formEmail)
	e.GET("/profile/:id", BlogDetail)
	e.GET("/project", MyProject)
	e.GET("/testimonial", Testimonial)
	e.POST("/Delete/:id", ProjectDelete)
	e.GET("/update/:id", UpdateDataForm)
	e.POST("/updateProject", UpdateData)

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
	tme, err := template.ParseFiles("html/project.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	DataQuery, errQuery := connect.Conn.Query(context.Background(), "SELECT * FROM tb_projects")

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, errQuery.Error())
	}

	var Result []Project
	for DataQuery.Next() {
		var Loop = Project{}

		err := DataQuery.Scan(&Loop.Id, &Loop.Image, &Loop.Title, &Loop.StartDate, &Loop.EndDate, &Loop.Articel, &Loop.CheckBox1, &Loop.CheckBox2, &Loop.CheckBox3, &Loop.CheckBox4)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		Result = append(Result, Loop)
	}

	Data := map[string]interface{}{
		"projects": Result,
	}
	return tme.Execute(c.Response(), Data)
}

func Testimonial(c echo.Context) error {
	tme, error := template.ParseFiles("html/testimonial.html")
	if error != nil {
		return c.JSON(http.StatusInternalServerError, error.Error())
	}
	return tme.Execute(c.Response(), nil)
}
func BlogDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tme, err := template.ParseFiles("html/profile.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	DataQuery, errQuery := connect.Conn.Query(context.Background(), "SELECT * FROM tb_projects")

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, errQuery.Error())
	}

	var projectDetail = Project{}

	for DataQuery.Next() {
		if projectDetail.Id == id {
			err := DataQuery.Scan(&projectDetail.Title, &projectDetail.Duration, &projectDetail.Articel, &projectDetail.CheckBox1, &projectDetail.CheckBox2, &projectDetail.CheckBox3, &projectDetail.CheckBox4)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

	}
	data := map[string]interface{}{
		"id":            id,
		"ProjectDetail": projectDetail,
	}
	return tme.Execute(c.Response(), data)
}

// Delete

func ProjectDelete(c echo.Context) error {
	id := c.Param("id")
	Iton, _ := strconv.Atoi(id)

	MyBlogs = append(MyBlogs[:Iton], MyBlogs[Iton+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

// // update
func UpdateDataForm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	UpdateForm := Project{}

	for i, update := range MyBlogs {
		if id == i {
			UpdateForm = Project{
				Id:        i,
				Title:     update.Title,
				Articel:   update.Articel,
				CheckBox1: update.CheckBox1,
				CheckBox2: update.CheckBox2,
				CheckBox3: update.CheckBox3,
				CheckBox4: update.CheckBox4,
			}

		}
	}
	data := map[string]interface{}{
		"projects": UpdateForm,
	}
	tme, err := template.ParseFiles("html/updateProject.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tme.Execute(c.Response(), data)
}

func UpdateData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	nameProject := c.FormValue("PName")
	TxtMsg := c.FormValue("TxtMsg")
	checkbox1, _ := strconv.ParseBool(c.FormValue("Cbx1"))
	checkbox2, _ := strconv.ParseBool(c.FormValue("Cbx2"))
	checkbox3, _ := strconv.ParseBool(c.FormValue("Cbx3"))
	checkbox4, _ := strconv.ParseBool(c.FormValue("Cbx4"))

	MyBlogs[id].Title = nameProject
	MyBlogs[id].Articel = TxtMsg
	MyBlogs[id].CheckBox1 = checkbox1
	MyBlogs[id].CheckBox2 = checkbox2
	MyBlogs[id].CheckBox3 = checkbox3
	MyBlogs[id].CheckBox4 = checkbox4

	return c.Redirect(http.StatusMovedPermanently, "/project")
}
