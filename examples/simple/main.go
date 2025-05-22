package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goapp/api"
	"goapp/controllers"
	"goapp/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	logger := setLogger()

	appCfg, err := loadAppCfg()
	if err != nil {
		log.Fatal(err)
	}

	logger.Info(fmt.Sprintf("Starting application %s...", appCfg.AppName))

	db, err := gorm.Open(mysql.Open(appCfg.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	userReadProjection := api.UserReadProjection{
		Repository: repositories.NewRepository(db),
	}

	if slices.Contains(appCfg.Profiles, "query") {
		// register query handlers
	}

	router := gin.Default()
	userController := controllers.NewUserController(userReadProjection)

	if slices.Contains(appCfg.Profiles, "api") {
		router.GET("/users/:id", userController.GetByID)
	}

	if err = router.Run(appCfg.HttpServeAddress); err != nil {
		log.Fatal(err)
	}
}

func setLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)
	return logger
}

func loadAppCfg() (Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	var appCfg Config
	return appCfg, viper.Unmarshal(&appCfg)
}
