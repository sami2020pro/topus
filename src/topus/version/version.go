package version

import "fmt"

const topusVersion = "0.1.0"
const description = `
	Topus programming language
`

func VersionWithDetail() string {
	result := fmt.Sprintf("%s%s", topusVersion, description)

	return result
}
