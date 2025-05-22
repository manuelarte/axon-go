package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"goapp/api"
	"goapp/controllers"
	"goapp/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	axongo "github.com/manuelarte/axon-go"
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
			endpoint, err := deleteAndRegisterEndpoint(ctx, c, "default", "go-app")
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Endpoint Registered: %s", endpoint.String())

			{
				/*params := &axongo.RegisterQueryHandlerParams{Context: "default"}
				  body := axongo.RegisterQueryHandlerJSONRequestBody{
				  	Name:     ptr("GetUserByIDQuery"),
				  	QueryUrl: ptr("http://host.docker.internal:8081/queries/GetUserByIDQuery"),
				  }
				  resp, err := c.RegisterQueryHandlerWithResponse(ctx, endpoint, params, body)
				    if err != nil {
				    	log.Fatal(err)
				    }
				    if resp.StatusCode() != http.StatusCreated {
				    	log.Fatalf("Expected HTTP 201 but received %d. Body: %s", resp.StatusCode(), resp.Body)
				    }
				    fmt.Printf("resp.JSON201: %v\n", resp.Body)*/
			}

		}
	}

	router := gin.Default()
	actuatorControllers := controllers.ActuatorControllers{}
	userController := controllers.NewUserController(userReadProjection)
	queryController := controllers.QueryController{}

	router.GET("/actuators/info", actuatorControllers.Info)
	router.GET("/queries", queryController.Get)
	router.GET("/queries/GetUserByIDQuery", queryController.GetUserByIDQuery)
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

func deleteAndRegisterEndpoint(
	ctx context.Context,
	c axongo.ClientWithResponsesInterface,
	endpointContext, endpointName string,
) (uuid.UUID, error) {
	var endpoint uuid.UUID
	{
		resp, err := c.EndpointsWithResponse(ctx)
		if err != nil {
			return endpoint, err
		}
		if resp.StatusCode() != http.StatusOK {
			return endpoint, fmt.Errorf("expected HTTP 200 but received %d. Body %+v", resp.StatusCode(), string(resp.Body))
		}
		var endpoints []*axongo.EndpointOverview
		if err = json.Unmarshal(resp.Body, &endpoints); err != nil {
			return endpoint, err
		}
		if i := slices.IndexFunc(endpoints, func(a *axongo.EndpointOverview) bool {
			return *a.Name == endpointName && *a.Context == endpointContext
		}); i >= 0 {
			e := uuid.MustParse(*endpoints[i].Id)
			_, err := c.DeleteEndpointWithResponse(ctx, e, &axongo.DeleteEndpointParams{Context: endpointContext})
			if err != nil {
				return endpoint, err
			}
		}
	}

	{
		params := &axongo.RegisterEndpointParams{Context: endpointContext}
		body := axongo.RegisterEndpointJSONRequestBody{
			BaseUrl:      ptr("http://host.docker.internal:8081"),
			HealthUrl:    ptr("/actuators/info"),
			QueryUrl:     ptr("/queries"),
			Name:         ptr(endpointName),
			Type:         ptr("HTTP(S)"), // HTTP(S), PRocket
			WrappingType: ptr("Raw"),
			ContentType:  ptr("application/json"),
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
		log.Print("Health endpoint Registered")
		type ResponseBody struct {
			ID uuid.UUID `json:"id"`
		}
		var respBody ResponseBody
		err = json.Unmarshal(resp.Body, &respBody)
		if err != nil {
			log.Fatal(err)
		}
		endpoint = respBody.ID
	}
	return endpoint, nil
}

func ptr[T any](v T) *T {
	return &v
}
