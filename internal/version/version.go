package version

var (
	// RELEASE stores the release version of the service.
	RELEASE = "UNKNOWN"

	// REPO stores the address of the repository where the source code for this service was taken from.
	REPO = "UNKNOWN"

	// COMMIT stores the hash of the commit.
	COMMIT = "UNKNOWN"

	// BUILD stores the build version (Pipeline number when compiled via CI/CD).
	BUILD = "UNKNOWN"
)
