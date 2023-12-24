package handler

import (
	"github.com/D-building-anonymaizer/backend-service/pkg/handler/file_workers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"html/template"
	//"github.com/D-building-anonymaizer/backend-service/pkg/mail"
	"log"
	"net/http"
	"os"
	"os/exec"
	// "encoding/json"
	// "io"
)

func (h *Handler) Index(c *gin.Context) {
	t, err := template.ParseFiles("../../build/index.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(c.Writer, "")
}

func (h *Handler) FileReciever(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://127.0.0.1:1337")                                                              // разрешить любой источник
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")                                           // разрешить определенные методы
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max") // разрешить определенные заголовки
	c.Header("Access-Control-Allow-Credentials", "true")                                                                          // разрешить отправку куки
	file, err := c.FormFile("file")

	log.Print(file.Filename)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s1, s2 := files.SplitFileName(file.Filename)
	str := viper.GetString("InputFolder") + s1 + s2
	err = c.SaveUploadedFile(file, str)
	if err != nil {

		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cmd := exec.Command("C:/Users/1/go/src/backend-service/configs/run")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Успешно сохранено!")
	c.JSON(http.StatusOK, gin.H{"message": str})

}

func (h *Handler) Exit(c *gin.Context) {
	//files.RemoveContents("../../output/")
	h.server.Shutdown(c)
	os.Exit(0)
}

//Данная функция и пакет mail реализуют готовый к развертыванию и использованию сервис эл. почты.
//Для использования необходимо наличие домена для налаживания работы с SMTP
//Инициализация переменных окружения (в т.ч. адресов эл. почты и атрибутов почтового сервиса) находится в файле configs/config.yml
// func (h *Handler) MailSender(c *gin.Context) {
// 	var data mail.UData
// 	body, err := io.ReadAll(c.Request.Body)
// 	log.Print(body)
// 	log.Print(c.Request.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = json.Unmarshal(body, &data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	mail.EmailSender(data)
// }
