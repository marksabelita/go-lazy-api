package asset

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const MainFileContent = `
package main 

import (
	"log"
	"{projectName}/src/common/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"

	_ "{projectName}/docs/{projectName}"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3001
// @BasePath /
// @schemes http

func main() {
	DEFAULT_PORT := config.GetEnv("PORT")
	app := fiber.New()

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())

	if DEFAULT_PORT == "" {
		DEFAULT_PORT = config.DEFAULT_PORT
	}

	config.ConnectDB()
	// user.UserRoutes(app)
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	if err := app.Listen(":" + DEFAULT_PORT); err != nil {
		log.Fatal(err)
	}
}
`

type MainTemplate struct {
	Template string
	Directory string
	FileName string
	Dependencies []string
	ProjectName string
}

func (m MainTemplate) GenerateConfigFile() bool {
	template := strings.Replace(m.Template, "{projectName}", m.ProjectName, -1)
	contents := []byte(template)
	path := m.ProjectName + "/" + m.FileName

	writeError := os.WriteFile(path, contents, os.ModePerm)

	if writeError != nil {
		panic(writeError)
	}

	for _, dep := range m.Dependencies {
		fmt.Println(`Installing dependencies ` + dep)
		cmdExec := exec.Command("go", "get", dep)
		cmdExec.Dir = m.ProjectName
		errExecErr := cmdExec.Run()

		if errExecErr != nil {
			panic(errExecErr)
		}
	}

	return true
}