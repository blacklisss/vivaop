package util

import (
	"os"
	"testing"
)

func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	os.Setenv("DB_DRIVER", "mysql")
	// Set other environment variables as needed

	config, err := LoadConfig("./../../")
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}

	if config.Environment != "test" {
		t.Errorf("Expected Environment to be 'test', got %s", config.Environment)
	}

	if config.DBDriver != "mysql" {
		t.Errorf("Expected DBDriver to be 'mysql', got %s", config.DBDriver)
	}
}

func TestLoadConfigFromFile(t *testing.T) {
	configFilePath := "./../../"

	_, err := LoadConfig(configFilePath)
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}

}

func TestLoadConfigPriority(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")

	config, err := LoadConfig("./../../")
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}

	if config.Environment != "test" {
		t.Errorf("Expected Environment to be 'test', got %s", config.Environment)
	}
}

func TestMissingConfigFile(t *testing.T) {
	_, err := LoadConfig("nonexistent_directory")
	if err == nil {
		t.Fatal("Expected error when config file is missing, got nil")
	}
}
