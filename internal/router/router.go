package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ilmovilDev/go-auth/config/env"
	"github.com/ilmovilDev/go-auth/pkg/logger"
)

// Initialize sets up the router and starts the server
func Initialize(envs *env.EnvConfig) {

	log := logger.NewLogger("Router")
	router := gin.New()

	// Configure trusted proxies (modify according to your network infrastructure)
	trustedProxies := []string{"192.168.1.1", "192.168.1.2"} // replace with your real proxy IPs
	if err := router.SetTrustedProxies(trustedProxies); err != nil {
		log.Errorf("Error setting trusted proxies: %v", err)
		return
	}

	initializeRoutes(router)

	appPort := fmt.Sprintf(":%d", envs.AppPort)
	if err := router.Run(appPort); err != nil {
		log.Errorf("Error initializing router: %v", err)
	}

}
