package constants

const PackagePrefix = "org.github.manuelarte.axongo.example"

func Ptr[T any](v T) *T {
	return &v
}
