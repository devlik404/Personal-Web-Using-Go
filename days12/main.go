package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id                                         int
	Title, Duration, Articel                   string
	CheckBox1, CheckBox2, CheckBox3, CheckBox4 bool
}

var MyBlogs = []Project{
	{
		Id:        1,
		Title:     "Software Enginering",
		Duration:  "20-07-2023",
		Articel:   "program sekarang harus memiliki paradigma peprogrmana ",
		CheckBox1: false,
		CheckBox2: true,
		CheckBox3: false,
		CheckBox4: true,
	},
	{
		Id:        2,
		Title:     "CEO Enginer",
		Duration:  "20-07-2023",
		Articel:   "saya memliki pandangan khusus sebuah paradigma pemprograman ",
		CheckBox1: true,
		CheckBox2: false,
		CheckBox3: false,
		CheckBox4: true,
	},
}

func main() {
	e := echo.New()
	e.Static("/assets", "assets")
	e.GET("/", Home)
	e.GET("/index2", formEmail)
	e.GET("/profile/:id", BlogDetail)
	e.GET("/project", MyProject)
	e.GET("/projectAdd", projectAdd)
	e.GET("/testimonial", Testimonial)
	e.POST("/project-Blog", addProject)
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
	tme, error := template.ParseFiles("html/project.html")
	if error != nil {
		return c.JSON(http.StatusInternalServerError, error.Error())
	}
	Data := map[string]interface{}{
		"projects": MyBlogs,
	}
	return tme.Execute(c.Response(), Data)
}

func projectAdd(c echo.Context) error {
	tme, error := template.ParseFiles("html/projectAdd.html")
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
	tme, err := template.ParseFiles("html/profile.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	Iton, _ := strconv.Atoi(id)

	DetailProject := Project{}

	for indexs, data := range MyBlogs {

		if indexs == Iton {
			DetailProject = Project{
				Title:     data.Title,
				Duration:  data.Duration,
				Articel:   data.Articel,
				CheckBox1: data.CheckBox1,
				CheckBox2: data.CheckBox2,
				CheckBox3: data.CheckBox3,
				CheckBox4: data.CheckBox4,
			}
		}
	}
	data := map[string]interface{}{
		"id":            id,
		"ProjectDetail": DetailProject,
	}
	return tme.Execute(c.Response(), data)
}

// Creat & Read
func addProject(c echo.Context) error {

	nameProject := c.FormValue("PName")
	// startDate := c.FormValue("FDate")
	// LastDate := c.FormValue("LDate")
	TxtMsg := c.FormValue("TxtMsg")
	checkbox1 := c.FormValue("Cbx1")
	checkbox2 := c.FormValue("Cbx2")
	checkbox3 := c.FormValue("Cbx3")
	checkbox4 := c.FormValue("Cbx4")

	boolValue, _ := strconv.ParseBool(checkbox1)
	boolValue1, _ := strconv.ParseBool(checkbox2)
	boolValue2, _ := strconv.ParseBool(checkbox3)
	boolValue3, _ := strconv.ParseBool(checkbox4)

	ProjectNew := Project{
		Title:     nameProject,
		Duration:  "20-07-2023",
		Articel:   TxtMsg,
		CheckBox1: boolValue,
		CheckBox2: boolValue1,
		CheckBox3: boolValue2,
		CheckBox4: boolValue3,
	}

	MyBlogs = append(MyBlogs, ProjectNew)

	return c.Redirect(http.StatusMovedPermanently, "/project")
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
