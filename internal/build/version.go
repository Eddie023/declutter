package build

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
	BuiltBy = "unknown"
)

func IsDev() bool {
	return Version == "dev"
}
