package producers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/models"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	gormPg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
}

func CreatePostgresContainer(config PostgresConfig, ctx context.Context) (*postgres.PostgresContainer, error) {

	// Disable ryuk to avoid race condition with the container on Podman
	os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			ExposedPorts: []string{config.DBPort},
			SkipReaper:   true,
			WaitingFor: wait.ForAll(
				wait.ForSQL(nat.Port(config.DBPort), "postgres", func(host string, port nat.Port) string {
					return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
						config.DBUser, config.DBPassword, host, port.Port(), config.DBName)
				}),
				wait.ForLog("database system is ready to accept connections"),
			),
		},
		Started: true,
	}

	postgresContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(config.DBName),
		postgres.WithUsername(config.DBUser),
		postgres.WithPassword(config.DBPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
		testcontainers.CustomizeRequest(req),
	)

	if err != nil {
		log.Printf("failed to start container: %s", err)
		return nil, err
	}

	return postgresContainer, nil
}

func SetupGenericSuite(ctx context.Context) (*postgres.PostgresContainer, *gorm.DB, error) {

	postgresConfig := PostgresConfig{
		DBName:     "mydget",
		DBUser:     "mydget",
		DBPassword: "mydget",
		DBHost:     "localhost",
		DBPort:     "5432",
	}

	pgContainer, err := CreatePostgresContainer(postgresConfig, ctx)
	if err != nil {
		return nil, nil, err
	}

	host, err := pgContainer.Host(ctx)
	if err != nil {
		return nil, nil, err
	}
	port, err := pgContainer.MappedPort(ctx, "5432")
	if err != nil {
		return nil, nil, err
	}

	db, err := gorm.Open(gormPg.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, postgresConfig.DBUser, postgresConfig.DBPassword, postgresConfig.DBName, port.Port())), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	db.AutoMigrate(&models.User{}, &models.Profile{})

	return pgContainer, db, nil
}

func TearDownGenericSuite(ctx context.Context, db *gorm.DB, pgContainer *postgres.PostgresContainer) {
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.Close()
	}

	if pgContainer != nil {
		if err := pgContainer.Terminate(ctx); err != nil {
			log.Printf("Failed to terminate container: %v", err)
		}
	}
}
