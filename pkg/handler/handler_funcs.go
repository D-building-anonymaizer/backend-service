package handler

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (h *Handler) Index(c *gin.Context) {
	t, err := template.ParseFiles("../../build/index.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(c.Writer, "")
}

func (h *Handler) FileReciever(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ext := filepath.Ext(file.Filename)
	allowed := viper.GetStringSlice("allowedXt")
	valid := false
	for _, a := range allowed {
		if ext == a {
			valid = true
			break
		}
	}
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file format"})
		return
	}

	folder := "../.../root"
	err = os.MkdirAll(folder, 0755)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dst := filepath.Join(folder, file.Filename)
	out, err := os.Create(dst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer out.Close()
	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()
	_, err = io.Copy(out, f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("file %s saved to %s", file.Filename, dst)})
}

func (h *Handler) Exit(c *gin.Context) {
	h.server.Shutdown(c)
}
