package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"project1/connect"
	"project1/midleware"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

type Project struct {
	Id                                         int
	Author, Path, Title, Articel, Duration     string
	CheckBox1, CheckBox2, CheckBox3, CheckBox4 bool
	Start, End                                 time.Time
}
type Users struct {
	Id                           int
	Name, Email, HashPwd, Author string
}
type LogEndSess struct {
	Name    string
	LogSess bool
}

var loginsession = LogEndSess{}

var MyBlogs = []Project{}

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 5 * time.Second,
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	// connection in database func
	connect.DbConection()

	e.Static("/assets", "assets")
	e.Static("/upload", "upload")
	e.GET("/", Home)
	e.GET("/index2", formEmail)
	e.GET("/testimonial", Testimonial)
	//creat
	e.GET("/projectAdd", projectAdd)
	e.POST("/project-Blog", midleware.MiddleFile(addProject))
	//read
	e.GET("/project", MyProject)
	e.GET("/profile/:id", BlogDetail)
	// delete
	e.POST("/Delete/:id", ProjectDelete)
	// update
	e.GET("/update/:id", UpdateDataForm)
	e.POST("/updateProject", UpdateData)
	//Auth
	e.GET("/register", FormRegister)
	e.POST("/form-register", Registrasi)

	e.GET("/validation", FormLogin)
	e.POST("/valid-form", Login)
	// logoutsession
	e.POST("/logout", LogoutSession)

	e.Logger.Fatal(e.Start("localhost:500"))
}
func Home(c echo.Context) error {
	tme, error := template.ParseFiles("html/index.html")
	if error != nil {
		return c.JSON(http.StatusInternalServerError, error.Error())
	}
	sess, _ := session.Get("cookie", c)

	// Ambil nilai Iduser dari session
	IduserInterface := sess.Values["id"]

	if IduserInterface == nil {
		// Jika Iduser belum ada di dalam session, mungkin user belum login atau sesi telah berakhir
		// ridirect ke halaman tapi fiturnya terbatas
		return tme.Execute(c.Response(), nil)
	}
	//konversi menjadi integer
	Iduser := sess.Values["id"].(int)
	if IduserInterface == nil {
		// Jika Iduser belum ada di dalam session, mungkin user belum login atau sesi telah berakhir
		// ridirect ke halaman tapi fiturnya terbatas
		return tme.Execute(c.Response(), nil)
	}
	var user = Users{}

	connect.Conn.QueryRow(context.Background(), "SELECT id,name,email FROM tb_users WHERE id=$1", Iduser).Scan(&user.Id, &user.Name, &user.Email)

	loginsession := GetLoginSession(c)

	Data := map[string]interface{}{
		"Loginsession": loginsession,
	}

	return tme.Execute(c.Response(), Data)
}

func formEmail(c echo.Context) error {
	tme, error := template.ParseFiles("html/index2.html")
	if error != nil {
		return c.JSON(http.StatusInternalServerError, error.Error())
	}
	loginsession := GetLoginSession(c)

	Data := map[string]interface{}{
		"Loginsession": loginsession,
	}

	return tme.Execute(c.Response(), Data)
}

func MyProject(c echo.Context) error {
	tme, err := template.ParseFiles("html/project.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	DataQuery, errQuery := connect.Conn.Query(context.Background(), "SELECT tb_project.id,tb_users.name,tb_project.image,tb_project.title,tb_project.articel,tb_project.check_box1,tb_project.check_box2,tb_project.check_box3,tb_project.check_box4,tb_project.start_date,tb_project.end_date,tb_project.duration FROM tb_project LEFT JOIN tb_users ON tb_project.author_id = tb_users.id")

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, errQuery.Error())
	}

	var Result []Project
	for DataQuery.Next() {
		var Loop = Project{}
		var nullname sql.NullString

		err := DataQuery.Scan(&Loop.Id, &nullname, &Loop.Path, &Loop.Title, &Loop.Articel, &Loop.CheckBox1, &Loop.CheckBox2, &Loop.CheckBox3, &Loop.CheckBox4, &Loop.Start, &Loop.End, &Loop.Duration)

		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusBadRequest, "Error scan dataquery")
		}
		// fmt.Println(nullname.String)
		Loop.Author = nullname.String
		Result = append(Result, Loop)

	}

	loginsession := GetLoginSession(c)

	Data := map[string]interface{}{
		"projects":     Result,
		"Loginsession": loginsession,
	}
	return tme.Execute(c.Response(), Data)
}

func Testimonial(c echo.Context) error {
	tme, error := template.ParseFiles("html/testimonial.html")
	if error != nil {
		return c.JSON(http.StatusInternalServerError, error.Error())
	}
	loginsession := GetLoginSession(c)
	Data := map[string]interface{}{
		"Loginsession": loginsession,
	}

	return tme.Execute(c.Response(), Data)
}
func BlogDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tme, err := template.ParseFiles("html/profile.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var projectDetail = Project{}
	var nullname sql.NullString
	connect.Conn.QueryRow(context.Background(), "SELECT tb_project.id,tb_users.name,tb_project.image,tb_project.title,tb_project.articel,tb_project.check_box1,tb_project.check_box2,tb_project.check_box3,tb_project.check_box4,tb_project.start_date,tb_project.end_date,tb_project.duration FROM tb_project LEFT JOIN tb_users ON tb_project.author_id = tb_users.id WHERE tb_project.id=$1", id).Scan(&projectDetail.Id, &nullname, &projectDetail.Path, &projectDetail.Title, &projectDetail.Articel, &projectDetail.CheckBox1, &projectDetail.CheckBox2, &projectDetail.CheckBox3, &projectDetail.CheckBox4, &projectDetail.Start, &projectDetail.End, &projectDetail.Duration)

	dt := "02-01-2006"

	d := time.Now().Format(dt)
	//variabel sql null
	projectDetail.Author = nullname.String
	//func session
	loginsession := GetLoginSession(c)

	data := map[string]interface{}{
		"id":            id,
		"ProjectDetail": projectDetail,
		"time":          d,
		"Loginsession":  loginsession,
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

	fileName := c.Get("DataFile").(string)

	sess, _ := session.Get("cookie", c)
	//INSERT database
	dataQuery, errQuery := connect.Conn.Exec(context.Background(), "INSERT INTO tb_project(image,title,articel,check_box1,check_box2,check_box3,check_box4,start_date,end_date,duration,author_id)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)", fileName, nameProject, TxtMsg, checkbox1, checkbox2, checkbox3, checkbox4, Start, End, durationString, sess.Values["id"].(int))
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
	// fileName := c.Get(c)
	// update blog berdasarkan Id
	upp, err := connect.Conn.Exec(context.Background(), "UPDATE tb_project SET image=$1,title=$2,articel=$3,check_box1=$4,check_box2=$5,check_box3=$6,check_box4=$7,start_date=$8,end_date=$9,duration=$10 WHERE id=$11", "1.jpg", nameProject, TxtMsg, checkbox1, checkbox2, checkbox3, checkbox4, Start, End, durationString, id)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error update")
	}
	fmt.Println("Update:", upp.RowsAffected())

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

//auth

func FormRegister(c echo.Context) error {
	tme, err := template.ParseFiles("html/register.html")
	if err != nil {
		return c.String(http.StatusBadRequest, "Error get register")
	}
	sess, errsess := session.Get("cookie", c)
	if errsess != nil {
		return c.JSON(http.StatusInternalServerError, errsess.Error())
	}
	Flashes := map[string]interface{}{
		"message": sess.Values["message"],
		"alert":   sess.Values["alert"],
	}
	delete(sess.Values, "message")
	delete(sess.Values, "alert")

	sess.Save(c.Request(), c.Response())

	return tme.Execute(c.Response(), Flashes)
}
func Registrasi(c echo.Context) error {
	InputName := c.FormValue("inputName")
	inputEmail := c.FormValue("inputEmail")
	inputPassword := c.FormValue("inputPassword")
	//Trim Auth
	// Lakukan validasi data input
	if inputPassword == "" || InputName == "" || inputEmail == "" {
		return FlashMessage(c, "inputkan terlebih dahulu !!", false, "/register")
	}
	// Hapus spasi putih di awal atau akhir string
	trimmedInput := strings.TrimSpace(inputPassword)
	// Pengecekan panjang data
	if len(trimmedInput) < 5 || len(trimmedInput) > 20 {
		return FlashMessage(c, "Registrasi Gagal ! Masukan Password 5 hingga 20 karakter", false, "/register")
	}

	//bcrypt hashing
	hashiteration, hasherr := bcrypt.GenerateFromPassword([]byte(inputPassword), 10)
	if hasherr != nil {
		return c.JSON(http.StatusInternalServerError, hasherr.Error())
	}

	QueryUser, QueryErr := connect.Conn.Exec(context.Background(), "INSERT INTO tb_users(name,email,password)VALUES($1,$2,$3)", InputName, inputEmail, hashiteration)

	// Cookie store

	fmt.Println("register berhasil :", QueryUser.RowsAffected())

	if QueryErr != nil {
		return c.JSON(http.StatusInternalServerError, hasherr.Error())
	}

	return FlashMessage(c, "Registerasi Berhasil Silahkan Login :)", true, "/validation")

}

//Validation

func FormLogin(c echo.Context) error {
	tme, err := template.ParseFiles("html/validation.html")
	if err != nil {
		return c.String(http.StatusBadRequest, "Error get Login")
	}
	sess, errsess := session.Get("cookie", c)
	if errsess != nil {
		return c.JSON(http.StatusInternalServerError, errsess.Error())
	}
	Flashes := map[string]interface{}{
		"message": sess.Values["message"],
		"alert":   sess.Values["alert"],
	}
	delete(sess.Values, "message")
	delete(sess.Values, "alert")

	sess.Save(c.Request(), c.Response())

	return tme.Execute(c.Response(), Flashes)
}
func Login(c echo.Context) error {

	inputEmail := c.FormValue("inputEmail")
	inputPassword := c.FormValue("inputPassword")

	users := Users{}

	QuerLogerr := connect.Conn.QueryRow(context.Background(), "SELECT id, name, email,password FROM tb_users WHERE email=$1", inputEmail).Scan(&users.Id, &users.Name, &users.Email, &users.HashPwd)

	if QuerLogerr != nil {
		return FlashMessage(c, "Masukan Email/Password terlebih dahulu!!", false, "/validation")
	}

	Comperr := bcrypt.CompareHashAndPassword([]byte(users.HashPwd), []byte(inputPassword))
	if Comperr != nil {
		return FlashMessage(c, "Email/Password Salah!!", false, "/validation")
	}

	sess, _ := session.Get("cookie", c)
	sess.Options.MaxAge = 10800 //per/detik
	sess.Values["name"] = users.Name
	sess.Values["email"] = users.Email
	sess.Values["id"] = users.Id
	sess.Values["login"] = true
	sess.Save(c.Request(), c.Response())

	// Tunggu sebentar untuk memberi waktu pada goroutine untuk menyelesaikan tindakannya
	time.Sleep(1 * time.Second)

	return c.Redirect(http.StatusMovedPermanently, "/")

}
func LogoutSession(c echo.Context) error {
	sess, _ := session.Get("cookie", c)

	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return FlashMessage(c, "logout berhasil", true, "/")
}

// function session Store
func FlashMessage(c echo.Context, message string, alert bool, redirectPath string) error {
	sess, errsess := session.Get("cookie", c)
	if errsess != nil {
		return c.JSON(http.StatusInternalServerError, errsess.Error())
	}
	sess.Values["message"] = message
	sess.Values["alert"] = alert
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, redirectPath)
}

// Fungsi untuk mengambil data dari session dan membuat data untuk template
func GetLoginSession(c echo.Context) LogEndSess {
	sess, _ := session.Get("cookie", c)
	loginsession := LogEndSess{}

	if sess.Values["login"] != true {
		loginsession.LogSess = false
	} else {
		loginsession.LogSess = true
		loginsession.Name = sess.Values["name"].(string)
	}

	return loginsession
}
