package version

import "fmt"

const (
	Version   = "0.1.0"
	GitCommit = "HEAD"
)

func CurrentVersion() string {
	return fmt.Sprintf("%s (%s)", Version, GitCommit)
}
