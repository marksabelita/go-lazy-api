package new_asset

import (
	"fmt"
	"log"
	"os"
)

const ConfigDefaultFileContent = `
package config

const DEFAULT_PORT string = "3001"
const DEFAULT_TIMEOUT = 10`

const EnvDefaultFileContent = `
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(c string) string {
	c = os.Getenv(c)
	return c
}
 
func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGO_URI")
}`

const MongoDefaultFileContent = `
package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(EnvMongoURI()))

	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
			log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	DEFAULT_DATABASE := GetEnv("DATABASE")
	
	if(DEFAULT_DATABASE == "") {
		log.Fatal("Database is not configured!")
	}

	collection := client.Database(DEFAULT_DATABASE).Collection(collectionName)
	return collection
}`

type ConfigTemplate struct {
	Template     string
	Directory 	 string
	FileName string
	ProjectName string
	Tidy bool
}

func (m ConfigTemplate) GenerateConfigFile() bool {
	configPaths := m.Directory + "/" + m.FileName 
	fmt.Println(configPaths)

	// Use os.MkdirAll to create nested directories
	mkdirAllError := os.MkdirAll(m.Directory, os.ModePerm)

	if mkdirAllError != nil {
		panic(mkdirAllError)
	}

	configContents := []byte(m.Template)
	writeError := os.WriteFile(configPaths, configContents, os.ModePerm)
	
	if writeError != nil {
		log.Fatal(writeError)
	}

	return true
}

