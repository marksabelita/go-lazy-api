package module_asset

import (
	"fmt"
	"os"
	"strings"
)

const RouteFileContent = `
package {module}

import (
	default_routes "{projectName}/src/common/defaults"

	"github.com/gofiber/fiber/v2"
)

func {Module}Routes(app *fiber.App) {
	app.Get(default_routes.DEFAULT_{MODULE}S_URI, Get{Module})
	app.Get(default_routes.DEFAULT_{MODULE}S_URI + "/:id", Get{Module}ById)
	app.Patch(default_routes.DEFAULT_{MODULE}S_URI + "/:id", Edit{Module})
	app.Delete(default_routes.DEFAULT_{MODULE}S_URI + "/:id", Delete{Module})
	app.Post(default_routes.DEFAULT_{MODULE}S_URI, Create{Module})
}`

type RouteTemplate struct {
	Template string
	Directory string
	FileName string
	ModuleName string
	ProjectName string
}

func (m RouteTemplate) GenerateConfigFile() bool {
	toUpperModuleName := strings.ToUpper(m.ModuleName)
	toCapitalize := capitalize(m.ModuleName)
	toLowername := strings.ToLower(m.ModuleName)
	template := strings.Replace(strings.Replace(strings.Replace(strings.Replace(m.Template, "{module}", toLowername, -1), "{Module}", toCapitalize, -1), "{MODULE}", toUpperModuleName, -1), "{projectName}", m.ProjectName, -1)
	
	fmt.Println(template)
	paths := m.Directory + "/" + m.FileName 
	fmt.Println(paths)
	contents := []byte(template)
	writeError := os.WriteFile(paths, contents, os.ModePerm)

	if writeError != nil {
		fmt.Println("log error")
		panic(writeError)
	}

	return true
}