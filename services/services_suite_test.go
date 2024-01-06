package services_test

import (
	"context"
	"dermsnap/database"
	"dermsnap/services"
	"os"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"
)

var DB *gorm.DB
var UserService services.UserService
var DermsnapService services.DermsnapService

func TestServices(t *testing.T) {
	ctx := context.Background()

	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:14-alpine"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		panic(err)
	}

	// Clean up the container
	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			panic(err)
		}
	}()

	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		panic(err)
	}
	os.Setenv("DATABASE_URL", connStr)

	DB = database.NewDatabase()
	UserService = services.NewUserService(DB)
	DermsnapService = services.NewDermsnapService(DB)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Services Suite Test")
}
