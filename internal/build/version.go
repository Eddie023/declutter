package build

var (
	Version = "dev"
)

func IsDev() bool {
	return Version == "dev"
}
