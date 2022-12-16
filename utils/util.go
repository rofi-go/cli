package utils

import "os"

func GetProjectPath(projectName string) string {
	goPath := os.Getenv("GOPATH")
	projectPath := goPath + "/src/github.com/" + projectName
	return projectPath
}
