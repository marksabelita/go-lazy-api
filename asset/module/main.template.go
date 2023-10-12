package module_asset

import (
	"lazy-api/asset"
	"os"
	"strings"
)

type MainModuleTemplate struct {
	ProjectName string
	ModuleName string
}

func (m MainModuleTemplate) GenerateConfigFile() bool {
	directory := m.ProjectName + "/src/module/" + m.ModuleName;
	mkdirAllError := os.MkdirAll(directory, os.ModePerm)

	if mkdirAllError != nil {
		panic(mkdirAllError)
	}

	moduleName := strings.ToLower(m.ModuleName)

	controllerModuleTemplate := ControllerTemplate{
		Template: ControllerFileContent,
		Directory: directory,
		FileName:  moduleName + ".controller.go",
		ModuleName: m.ModuleName,
		ProjectName: m.ProjectName,
	}
	asset.Generate(controllerModuleTemplate)

	routeModuleTemplate := RouteTemplate{
		Template: RouteFileContent,
		Directory: directory,
		FileName: moduleName + ".route.go",
		ModuleName: m.ModuleName,
		ProjectName: m.ProjectName,
	}
	asset.Generate(routeModuleTemplate)

	return true
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}