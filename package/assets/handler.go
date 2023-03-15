package assets

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type DataFormat string

const (
	Svg DataFormat = "svg"
	Wav DataFormat = "wav"
	Png DataFormat = "png"
)

type Handler struct{}

func (h *Handler) InitAssetsRoutes(router *gin.Engine) {
	assets := router.Group("/")
	{
		assets.GET("/:md5ext", h.GetAsset)
		assets.POST("/:md5ext", h.StoreAsset)
		//assets.DELETE("/:md5ext", h.GetAsset)
		//assets.PUT("/:md5ext", h.GetAsset)
	}
}

func (h *Handler) GetAsset(c *gin.Context) {
	log.Print("GetAsset")
	md5ext := c.Param("md5ext")
	md5extParts := strings.Split(md5ext, ".")
	if len(md5extParts) != 2 {
		log.Print("invalid asset md5ext")
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid asset md5ext"))
		return
	}
	buf, err := ioutil.ReadFile("./static/" + md5ext)
	if err != nil {
		log.Print(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Header("Content-Type", "img/svg+xml")
	c.Header("Content-Disposition", "attachment; filename=\""+md5ext+"\"")
	c.Writer.Write(buf)
	c.Status(http.StatusOK)
	return
}

func (h *Handler) StoreAsset(c *gin.Context) {
	log.Print("StoreAsset")
	md5ext := c.Param("md5ext")
	md5extParts := strings.Split(md5ext, ".")
	if len(md5extParts) != 2 {
		log.Print("invalid asset md5ext")
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid asset md5ext"))
		return
	}
	body, readingErr := io.ReadAll(c.Request.Body)
	if readingErr != nil {
		log.Print(readingErr)
		c.AbortWithError(http.StatusBadRequest, readingErr)
		return
	}
	newAssetFile, creatingErr := os.Create("./static/" + md5ext)
	if creatingErr != nil {
		log.Print(readingErr)
		c.AbortWithError(http.StatusBadRequest, creatingErr)
		return
	}
	newAssetFile.Write(body)
	newAssetFile.Close()
	c.Status(http.StatusOK)
	return
}
