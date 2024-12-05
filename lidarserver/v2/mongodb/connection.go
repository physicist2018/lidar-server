package mongodb

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kataras/golog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DB_NAME = "Lidar"

type Client struct {
	client *mongo.Client
}

var client *Client

// GetDefaultMongoClient is a function that returns a MongoDB client and an error.
// It checks if the MongoDB client is already created. If it is, it returns the existing client.
// If it is not, it creates a new client using the NewClient function with the MongoDB URI from the environment variable "LIDARDB_URI" and returns the new client.
// If there is an error while creating the client, it returns the error.
func GetDefaultMongoClient() (*Client, error) {
	if client == nil {
		client, err := NewClient(os.Getenv("LIDARDB_URI"))
		if err != nil {
			return nil, err
		}
		return client, nil
	}
	return client, nil
}

// GetDefaultMongoClientPanic is a function that returns a MongoDB client.
// It calls the GetDefaultMongoClient function and panics if there is an error.
// It is used when the program cannot continue without a valid MongoDB client.
func GetDefaultMongoClientPanic() *Client {
	client, err := GetDefaultMongoClient()
	if err != nil {
		panic(err)
	}
	return client
}

// GetMongoClient is a function that returns a MongoDB client and an error.
// It checks if the MongoDB client is already created. If it is, it returns the existing client.
// If it is not, it creates a new client using the NewClient function with the provided URI and returns the new client.
// If there is an error while creating the client, it returns the error.
func GetMongoClient(uri string) (*Client, error) {
	if client == nil {
		client, err := NewClient(uri)
		if err != nil {
			return nil, err
		}
		return client, nil
	}
	return client, nil
}

func NewClient(uri string) (*Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	return &Client{client: client}, nil
}

func (c *Client) Ping() error {
	return c.client.Ping(context.Background(), nil)
}

func (c *Client) Close() error {
	return c.client.Disconnect(context.Background())
}

// ListCollectionNames lists all collection names in Lidar database
func (c *Client) ListCollectionNames() ([]string, error) {
	return c.client.Database(DB_NAME).ListCollectionNames(context.Background(), bson.D{})
}

// CreateDatabase creates a new database in Lidar database and panics if there is an error
func (c *Client) ListCollectionNamesPanic() []string {
	names, err := c.ListCollectionNames()
	if err != nil {
		panic(err)
	}
	return names
}

// CreateCollection creates a new collection in Lidar database
// name - is the name of the collection
// returns error if collection already exists or if there is an error
func (c *Client) CreateCollection(name string) error {
	return c.client.Database(DB_NAME).CreateCollection(context.Background(), name)
}

// GetCollection returns a collection in Lidar database
// name - is the name of the collection
// returns nil if collection does not exist or if there is an error
func (c *Client) GetCollection(name string) *mongo.Collection {
	return c.client.Database(DB_NAME).Collection(name)
}

// DropCollection drops a collection in Lidar database
// name - is the name of the collection
// returns error if collection does not exist or if there is an error
func (c *Client) DropCollection(name string) error {
	return c.client.Database(DB_NAME).Collection(name).Drop(context.Background())
}

func init() {
	godotenv.Load()
	golog.Info("Connecting to MONGO")
	clt := GetDefaultMongoClientPanic()
	golog.Info("Creating collection")
	clt.CreateCollection("Experiment")
}
