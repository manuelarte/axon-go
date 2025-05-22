package api

type GetUserByIDQuery struct {
	ID int `json:"id" uri:"id" binding:"required"`
}

func (g GetUserByIDQuery) GetPackageName() string {
	return "org.github.axonserver.connector.go.example.kotlinapp.api.GetUserByIDQuery"
}
