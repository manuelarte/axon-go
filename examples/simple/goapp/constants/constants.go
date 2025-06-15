package constants

const packagePrefix = "org.github.manuelarte.axongo.example"

const (
	UserReadType         = packagePrefix + ".api.UserRead"
	GetUserByIDQueryType = packagePrefix + ".api.GetUserByIDQuery"
)

func Ptr[T any](v T) *T {
	return &v
}
