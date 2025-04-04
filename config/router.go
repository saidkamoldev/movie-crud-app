package config

import (
	"github.com/gin-gonic/gin"
	"log"
)

// NewRouter - yangi Gin routerini qaytaradi va portni ishga tushiradi
func NewRouter() *gin.Engine {
	r := gin.Default()

	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Serverni ishga tushirishda xatolik: %v", err)
		}
	}()

	return r
}
