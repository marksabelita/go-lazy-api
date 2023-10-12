package asset

type TemplateInterface interface{
	GenerateConfigFile() bool
}

func Generate(t TemplateInterface) {
	t.GenerateConfigFile()
}