/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"lazy-api/asset"
	module_asset "lazy-api/asset/module"
	new_asset "lazy-api/asset/new"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create new REST API project",
	Long: `Create new REST API with fiber go.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		projectDirectory := projectName + "/"

		initializeTemplate := new_asset.IntializeTemplate{
			ProjectName: projectName,
		}
		asset.Generate(initializeTemplate)
		
		mainTemplate := new_asset.MainTemplate{
			Template: new_asset.MainFileContent,
			Directory: projectDirectory,
			FileName: "main.go",
			Dependencies: []string{
				"github.com/gofiber/fiber/v2", 
				"github.com/gofiber/swagger", 
				"go.mongodb.org/mongo-driver",
				"github.com/joho/godotenv",
				"github.com/go-playground/validator/v10",
			},
			ProjectName: projectName,
		}
		asset.Generate(mainTemplate)

		defaultConfigTemplate := new_asset.ConfigTemplate{
			Template: new_asset.ConfigDefaultFileContent,
			Directory: projectDirectory + "src/common/config",
			FileName: "defaults.config.go",
		}
		asset.Generate(defaultConfigTemplate)

		envConfigTemplate := new_asset.ConfigTemplate{
			Template: new_asset.EnvDefaultFileContent,
			Directory: projectDirectory + "src/common/config",
			FileName: "env.config.go",
		}
		asset.Generate(envConfigTemplate)

		envMongoTemplate := new_asset.ConfigTemplate{
			Template: new_asset.MongoDefaultFileContent,
			Directory: projectDirectory + "src/common/config",
			FileName: "mongo.config.go",
			Tidy: true,
		}
		asset.Generate(envMongoTemplate)


		// tidyCmdExec := exec.Command("go", "mod", "tidy")
		// tidyCmdExec.Dir = projectName
		// tidyErrExecErr := tidyCmdExec.Run()

		// if tidyErrExecErr != nil {
		// 	panic(tidyErrExecErr)
		// }	

		mainModuleTemplate := module_asset.MainModuleTemplate{
			ModuleName: "user",
			ProjectName: projectName,
		}

		asset.Generate(mainModuleTemplate)

	

	},


	
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
