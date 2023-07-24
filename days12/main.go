package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id int

	StartDate, EndDate                         time.Time
	Duration                                   time.Duration
	Path, Title, Articel                       string
	CheckBox1, CheckBox2, CheckBox3, CheckBox4 bool
}

var MyBlogs = []Project{
	{
		Id:        1,
		Path:      "2.jpg",
		Title:     "Software Enginering",
		Articel:   "program sekarang harus memiliki paradigma peprogrmana ",
		CheckBox1: false,
		CheckBox2: true,
		CheckBox3: false,
		CheckBox4: true,
	},
	{
		Id:        2,
		Path:      "7.jpg",
		Title:     "CEO Enginer",
		Articel:   "saya memliki pandangan khusus sebuah paradigma pemprograman ",
		CheckBox1: true,
		CheckBox2: false,
		CheckBox3: false,
		CheckBox4: true,
	},
	{
		Id:        2,
		Path:      "9.jpg",
		Title:     "CEO Enginer",
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
	tme, err := template.ParseFiles("html/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tme.Execute(c.Response(), nil)
}

func formEmail(c echo.Context) error {
	tme, err := template.ParseFiles("html/index2.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tme.Execute(c.Response(), nil)
}

func MyProject(c echo.Context) error {
	tme, err := template.ParseFiles("html/project.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	Data := map[string]interface{}{
		"projects": MyBlogs,
	}
	return tme.Execute(c.Response(), Data)
}

func projectAdd(c echo.Context) error {
	tme, err := template.ParseFiles("html/projectAdd.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
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
				Path:      data.Path,
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
	startDate := c.FormValue("FDate")
	LastDate := c.FormValue("LDate")
	TxtMsg := c.FormValue("TxtMsg")

	//ambil value dari stiap checkbox dan convert ke bool
	checkbox1, _ := strconv.ParseBool(c.FormValue("Cbx1"))
	checkbox2, _ := strconv.ParseBool(c.FormValue("Cbx2"))
	checkbox3, _ := strconv.ParseBool(c.FormValue("Cbx3"))
	checkbox4, _ := strconv.ParseBool(c.FormValue("Cbx4"))

	/*-----------------------input date section-------------------*/

	// Parse tanggal dari string ke format time.Time

	layout := "2006-01-02" // Format dari input date

	Start, err := time.Parse(layout, startDate)
	if err != nil {
		return fmt.Errorf("Error parsing Start Date: %s", err)
	}
	End, err := time.Parse(layout, LastDate)
	if err != nil {
		return fmt.Errorf("Error parsing End Date: %s", err)
	}
	if Start.After(End) {
		return c.String(http.StatusRequestURITooLong, "WARNING: Please insert a First Date > Last Date!!")
	}
	// Hitung durasi (selisih) antara endDate dan startDate
	duration := End.Sub(Start)

	/*---------------   input file section   ----------------- */

	//ambil value dari name input file
	file, err := c.FormFile("image")
	if err != nil {
		return c.String(http.StatusBadRequest, "Error uploading file")
	}
	// Dapatkan nama file dari file yang diunggah
	fileName := file.Filename
	// Buka file yang akan diunggah
	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error opening file")
	}
	defer src.Close()

	// Buat file baru di server untuk menyimpan gambar
	dstPath := filepath.Join("assets", "image", fileName)
	dst, err := os.Create(dstPath)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error creating destination file")
	}
	defer dst.Close()

	// Salin isi file dari src ke dst (unggah)
	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	ProjectNew := Project{
		Path:      fileName,
		Title:     nameProject,
		Duration:  duration,
		Articel:   TxtMsg,
		CheckBox1: checkbox1,
		CheckBox2: checkbox2,
		CheckBox3: checkbox3,
		CheckBox4: checkbox4,
	}
	fmt.Println(ProjectNew.Path)
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
				Path:      update.Path,
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

	//ambil value dari name input file
	file, err := c.FormFile("image")
	if err != nil {
		return c.String(http.StatusBadRequest, "Error uploading file")
	}
	// Dapatkan nama file dari file yang diunggah
	fileName := file.Filename
	// Buka file yang akan diunggah
	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error opening file")
	}
	defer src.Close()

	// Buat file baru di server untuk menyimpan gambar
	dstPath := filepath.Join("assets", "image", fileName)
	dst, err := os.Create(dstPath)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error creating destination file")
	}
	defer dst.Close()

	// Salin isi file dari src ke dst (unggah)
	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	MyBlogs[id].Path = fileName
	MyBlogs[id].Title = nameProject
	MyBlogs[id].Articel = TxtMsg
	MyBlogs[id].CheckBox1 = checkbox1
	MyBlogs[id].CheckBox2 = checkbox2
	MyBlogs[id].CheckBox3 = checkbox3
	MyBlogs[id].CheckBox4 = checkbox4

	return c.Redirect(http.StatusMovedPermanently, "/project")
}
