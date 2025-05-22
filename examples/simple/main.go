package main

import (
	"context"
	"fmt"
	axongo "github.com/manuelarte/axon-go"
	"log"
	"log/slog"
	"net/http"
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
		ctx := context.Background()
		hc := http.Client{}

		{
			c, err := axongo.NewClientWithResponses("http://localhost:8024", axongo.WithHTTPClient(&hc))
			if err != nil {
				log.Fatal(err)
			}
			{
				resp, err := c.EndpointsWithResponse(ctx, func(ctx context.Context, req *http.Request) error {
					req.Header.Add("accept", "*/*")
					return nil
				})
				if err != nil {
					log.Fatal(err)
				}
				if resp.StatusCode() != http.StatusOK {
					log.Fatalf("[Endpoints] Expected HTTP 200 but received %d. Body: %s", resp.StatusCode(), resp.Body)
				}
			}
			{
				params := &axongo.RegisterEndpointParams{Context: "default"}
				body := axongo.RegisterEndpointJSONRequestBody{
					BaseUrl:      Ptr("http://host.docker.internal:8081"),
					HealthUrl:    Ptr("/actuators/info"),
					Name:         Ptr("go-app"),
					Type:         Ptr("HTTP(S)"), //HTTP(S), PRocket
					WrappingType: Ptr("Raw"),
					ContentType:  Ptr("application/json"),
				}
				resp, err := c.RegisterEndpointWithResponse(ctx, params, body, func(ctx context.Context, req *http.Request) error {
					req.Header.Add("accept", "*/*")
					req.Header.Add("content-type", "application/json")
					return nil
				})
				if err != nil {
					log.Fatal(err)
				}
				if resp.StatusCode() != http.StatusCreated {
					log.Fatalf("Expected HTTP 201 but received %d. Body: %s", resp.StatusCode(), resp.Body)
				}
			}

			/*{
				params := &axongo.RegisterQueryHandlerParams{Context: "default"}
				body := axongo.RegisterQueryHandlerJSONRequestBody{
					Name:     Ptr("GetUserByID"),
					QueryUrl: Ptr("http://host.docker.internal:8081"),
				}
				resp, err := c.RegisterQueryHandlerWithResponse(ctx, uuid.New(), params, body)
				if err != nil {
					log.Fatal(err)
				}
				if resp.StatusCode() != http.StatusCreated {
					log.Fatalf("Expected HTTP 201 but received %d. Body: %s", resp.StatusCode(), resp.Body)
				}
				fmt.Printf("resp.JSON201: %v\n", resp.JSON201)
			}*/

		}
	}

	router := gin.Default()
	actuatorControllers := controllers.ActuatorControllers{}
	userController := controllers.NewUserController(userReadProjection)

	router.GET("/actuators/info", actuatorControllers.Info)
	router.GET("/users/:id", userController.GetByID)
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

func Ptr[T any](v T) *T {
	return &v
}
