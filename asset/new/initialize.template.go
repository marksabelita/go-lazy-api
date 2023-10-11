package asset

import (
	"os"
	"os/exec"
)

type IntializeTemplate struct {
	ProjectName string
}

func (i IntializeTemplate) GenerateConfigFile() bool {
		// GENERATE MAIN FILE 
    // create directory 
		if err := os.Mkdir(i.ProjectName, os.ModePerm); err != nil {
			panic(err)
		}

    cmdExec := exec.Command("go", "mod", "init", i.ProjectName)
    cmdExec.Dir = i.ProjectName
		errExecErr := cmdExec.Run()

		if errExecErr != nil {
			panic(errExecErr)
    }	

		return true;
}