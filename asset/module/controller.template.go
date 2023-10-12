package module_asset

import (
	"fmt"
	"os"
	"strings"
)

const ControllerFileContent = `
package {module}

import (
	"context"
	"{projectName}/src/common/config"
	"{projectName}/src/common/defaults"
	{module}_model "{projectName}/src/module/{module}/model"
	{module}_response "{projectName}/src/module/{module}/response"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var {module}Collection *mongo.Collection = config.GetCollection(config.DB, defaults.DEFAULT_{MODULE}_COLLECTION)

// @Summary Lists all {module}s details.
// @Description Lists all {module}s details.
// @Tags {Module}s
// @Accept json
// @Produce json
// @Success 200 {array} []{module}_model.{Module}
// @Param        name    query     string  false  "name"  
// @Router /{module}s [get]
func Get{Module}(c *fiber.Ctx) error {
    query := bson.M{}
    name := c.Query("name")

    if name != "" { 
        query["name"] = name  
    }
    
    {module}s, err := FindService(query)

    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON({module}_response.ErrorResponse{Message: "error"})
    }
   
    return c.Status(http.StatusOK).JSON(
        {module}s,
    )
}

// @Summary Display {module} details
// @Description Display {module} details
// @Tags {Module}s
// @Accept json
// @Produce json
// @Success 200 {object} {module}_model.{Module}
// @Param        id   path      string  true  "Account ID"
// @Router /{module}s/{id} [get]
func Get{Module}ById(c *fiber.Ctx) error {
	{module}Id := c.Params("id")
	objId, _ := primitive.ObjectIDFromHex({module}Id)
    query := bson.M{"id": objId}
    {module}, err := FindOneService(query);

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON({module}_response.{Module}Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON({module})
}

// @Summary Update {module}
// @Description Update {module} details
// @Tags {Module}s
// @Accept json
// @Produce json
// @Success 200 {object} {module}_model.{Module}
// @Param data body {module}_model.{Module} true "{Module} data"
// @Router /{module}s [patch]
func Edit{Module}(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    {module}Id := c.Params("{module}Id")
    var {module} {module}_model.{Module}
    defer cancel()

    objId, _ := primitive.ObjectIDFromHex({module}Id)

    //validate the request body
    if err := c.BodyParser(&{module}); err != nil {
        return c.Status(http.StatusBadRequest).JSON({module}_response.{Module}Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

    //use the validator library to validate required fields
		validate := validator.New()
    if validationErr := validate.Struct(&{module}); validationErr != nil {
        return c.Status(http.StatusBadRequest).JSON({module}_response.{Module}Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
    }

    update := bson.M{"name": {module}.Name}

    result, err := {module}Collection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON({module}_response.{Module}Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
    //get updated {module} details
    var updated{Module} {module}_model.{Module}
    if result.MatchedCount == 1 {
        err := {module}Collection.FindOne(ctx, bson.M{"id": objId}).Decode(&updated{Module})

        if err != nil {
            return c.Status(http.StatusInternalServerError).JSON({module}_response.{Module}Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
        }
    }

    return c.Status(http.StatusOK).JSON({module}_response.{Module}Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updated{Module}}})
}


// @Summary Create {module}
// @Description Create {module} details
// @Tags {Module}s
// @Accept json
// @Produce json
// @Success 200 {object} {module}_model.{Module}
// @Param data body {module}_model.{Module} true "{Module} data"
// @Router /{module}s [post]
func Create{Module}(c *fiber.Ctx) error {
	var {module} {module}_model.{Module}
	ctx, cancel := context.WithTimeout(context.Background(), config.DEFAULT_TIMEOUT * time.Second)
		defer cancel()

	//validate the request body
	if err := c.BodyParser(&{module}); err != nil {
		return c.Status(http.StatusBadRequest).JSON({module}_response.{Module}Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	validate := validator.New()

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&{module}); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON({module}_response.{Module}Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	new{Module} := {module}_model.{Module}{
		Id: primitive.NewObjectID(),
		Name: {module}.Name,
	}

	result, err := {module}Collection.InsertOne(ctx, new{Module})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON({module}_response.{Module}Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON({module}_response.{Module}Response{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

// @Summary Display {module} details
// @Description Display {module} details
// @Tags {Module}s
// @Accept json
// @Produce json
// @Success 200 {object} {module}_model.{Module}
// @Param        id   path      string  true  "Account ID"
// @Router /{module}s/{id} [delete]
func Delete{Module}(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    {module}Id := c.Params("{module}Id")
    defer cancel()

    objId, _ := primitive.ObjectIDFromHex({module}Id)

    result, err := {module}Collection.DeleteOne(ctx, bson.M{"id": objId})
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON({module}_response.{Module}Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

    if result.DeletedCount < 1 {
        return c.Status(http.StatusNotFound).JSON(
            {module}_response.{Module}Response{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "{Module} with specified ID not found!"}},
        )
    }

    return c.Status(http.StatusOK).JSON(
        {module}_response.{Module}Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "{Module} successfully deleted!"}},
    )
}`


type ControllerTemplate struct {
	Template string
	Directory string
	FileName string
	ModuleName string
	ProjectName string
}

func (m ControllerTemplate) GenerateConfigFile() bool {
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
