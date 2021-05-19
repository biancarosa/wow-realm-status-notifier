package configuration

import (
	"context"
	"fmt"
	"log"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type ConfigurationProvider interface {
	PopulateConfig() *Config
}

type EnvironmentVariablesConfigurationProvider struct{}

func (p *EnvironmentVariablesConfigurationProvider) PopulateConfig(c *Config) *Config {
	c.TelegramToken = os.Getenv("TELEGRAM_TOKEN")
	c.Port = os.Getenv("PORT")
	c.GoogleProjectID = os.Getenv("GOOGLE_PROJECT_ID")
	c.MongoDB.Host = os.Getenv("MONGODB_HOST")
	c.MongoDB.Port = os.Getenv("MONGODB_PORT")
	c.MongoDB.Database = os.Getenv("MONGODB_DATABASE")
	c.MongoDB.Collection = os.Getenv("MONGODB_COLLECTION")
	return c
}

func getSecretFromGSM(projectID, secretName string) string {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, secretName),
	}
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Fatalf("failed to access secret version: %v", err)
	}
	return fmt.Sprintf("%s", result.Payload.Data)
}

type GoogleSecretManagerConfigurationProvider struct{}

func (p *GoogleSecretManagerConfigurationProvider) PopulateConfig(c *Config) *Config {
	if c.GoogleProjectID == "" {
		panic("Can't get env vars from Google without Google Project ID set")
	}
	c.TelegramToken = getSecretFromGSM(c.GoogleProjectID, "TELEGRAM_TOKEN")
	return c
}
