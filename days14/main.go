package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"project1/connect"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id                                         int
	Path, Title, Articel, Duration             string
	CheckBox1, CheckBox2, CheckBox3, CheckBox4 bool
	Start, End                                 time.Time
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

	e.GET("/projectAdd", projectAdd)

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
	tme, err := template.ParseFiles("html/project.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	DataQuery, errQuery := connect.Conn.Query(context.Background(), "SELECT * FROM tb_project")

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, errQuery.Error())
	}

	var Result []Project
	for DataQuery.Next() {
		var Loop = Project{}

		err := DataQuery.Scan(&Loop.Id, &Loop.Path, &Loop.Title, &Loop.Articel, &Loop.CheckBox1, &Loop.CheckBox2, &Loop.CheckBox3, &Loop.CheckBox4, &Loop.Start, &Loop.End, &Loop.Duration)

		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusBadRequest, "Error scan dataquery")
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
	var projectDetail = Project{}

	connect.Conn.QueryRow(context.Background(), "SELECT * FROM tb_project WHERE id=$1", id).Scan(&projectDetail.Id, &projectDetail.Path, &projectDetail.Title, &projectDetail.Articel, &projectDetail.CheckBox1, &projectDetail.CheckBox2, &projectDetail.CheckBox3, &projectDetail.CheckBox4, &projectDetail.Start, &projectDetail.End, &projectDetail.Duration)

	dt := "02-01-2006"
	d := time.Now().Format(dt)

	data := map[string]interface{}{
		"id":            id,
		"ProjectDetail": projectDetail,
		"time":          d,
	}
	return tme.Execute(c.Response(), data)
}

func projectAdd(c echo.Context) error {
	tme, err := template.ParseFiles("html/projectAdd.html")
	if err != nil {
		return c.String(http.StatusBadRequest, "Error projectadd")
	}
	return tme.Execute(c.Response(), nil)
}

// // creat

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

	layout := "2006-01-02T15:04" // Format dari input datetime

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
	// // Hitung durasi (selisih) antara endDate dan startDate
	duration := End.Sub(Start)
	// Mengambil komponen tanggal, bulan, tahun, jam, menit, dan detik dari durasi
	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	// Ubah durasi menjadi string dengan format "Jam:Menit:Detik"
	durationString := fmt.Sprintf("Durasi: %d hari, %02d jam, %02d menit\n", days, hours, minutes)

	// /*---------------   input file section   ----------------- */

	// //ambil value dari name input file
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
		return c.String(http.StatusInternalServerError, "Error file")
	}

	//INSERT database
	dataQuery, errQuery := connect.Conn.Exec(context.Background(), "INSERT INTO tb_project(image,title,articel,check_box1,check_box2,check_box3,check_box4,start_date,end_date,duration)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", fileName, nameProject, TxtMsg, checkbox1, checkbox2, checkbox3, checkbox4, Start, End, durationString)
	if errQuery != nil {
		return c.String(http.StatusInternalServerError, "Error add blog")
	}
	fmt.Println("row affected:", dataQuery.RowsAffected())

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

// Delete

func ProjectDelete(c echo.Context) error {
	id := c.Param("id")
	Iton, _ := strconv.Atoi(id)

	connect.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", Iton)

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

// // Render Tampilan
func UpdateDataForm(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error get")
	}
	tme, err := template.ParseFiles("html/updateProject.html")
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	data := map[string]interface{}{
		"Id": id,
	}

	return tme.Execute(c.Response(), data)
}

// prosesing
func UpdateData(c echo.Context) error {
	id := c.FormValue("id")
	nameProject := c.FormValue("PName")
	TxtMsg := c.FormValue("TxtMsg")
	startDate := c.FormValue("FDate")
	LastDate := c.FormValue("LDate")
	checkbox1, _ := strconv.ParseBool(c.FormValue("Cbx1"))
	checkbox2, _ := strconv.ParseBool(c.FormValue("Cbx2"))
	checkbox3, _ := strconv.ParseBool(c.FormValue("Cbx3"))
	checkbox4, _ := strconv.ParseBool(c.FormValue("Cbx4"))

	/*-----------------------input date section-------------------*/

	// Parse tanggal dari string ke format time.Time

	layout := "2006-01-02T15:04" // Format dari input datetime

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
	// // Hitung durasi (selisih) antara endDate dan startDate
	duration := End.Sub(Start)
	// Mengambil komponen tanggal, bulan, tahun, jam, menit, dan detik dari durasi
	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	// Ubah durasi menjadi string dengan format "Jam:Menit:Detik"
	durationString := fmt.Sprintf("Durasi: %d hari, %02d jam, %02d menit\n", days, hours, minutes)

	// /*---------------   input file section   ----------------- */

	// //ambil value dari name input file
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
		return c.String(http.StatusInternalServerError, "Error file")
	}
	// update blog berdasarkan Id
	upp, err := connect.Conn.Exec(context.Background(), "UPDATE tb_project SET image=$1,title=$2,articel=$3,check_box1=$4,check_box2=$5,check_box3=$6,check_box4=$7,start_date=$8,end_date=$9,duration=$10 WHERE id=$11", fileName, nameProject, TxtMsg, checkbox1, checkbox2, checkbox3, checkbox4, Start, End, durationString, id)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error update")
	}
	fmt.Println("Update:", upp.RowsAffected())

	return c.Redirect(http.StatusMovedPermanently, "/project")
}
