package grpc

import (
	"context"
	"net/url"
	"path/filepath"
	"testing"
	"time"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/repo/cache"
	"github.com/Hidayathamir/gocheck/internal/repo/db"
	"github.com/Hidayathamir/gocheck/internal/repo/db/migration/migrate"
	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	postgrescontainers "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/mock/gomock"
)

func setup(t *testing.T) (cfg *config.Config, ctrl *gomock.Controller, pg *db.Postgres, redis *cache.Redis) { //nolint:unparam
	t.Helper()

	cfg, err := config.Load(filepath.Join("..", "..", "..", "config.yml"))
	require.NoError(t, err)

	pgContainer := createPGContainer(t, cfg)
	t.Cleanup(func() { require.NoError(t, pgContainer.Terminate(context.Background())) })

	updateConfigPGPort(t, cfg, pgContainer)

	pg, err = db.NewPostgres(cfg)
	require.NoError(t, err)

	err = migrate.Up(pg.DB, migrate.WithDir(filepath.Join("..", "..", "..", "internal", "repo", "db", "migration", "schema_migration")))
	require.NoError(t, err)

	miniRedis := miniredis.RunT(t)

	updateConfigRedisPort(t, cfg, miniRedis)

	redis, err = cache.NewRedis(cfg)
	require.NoError(t, err)

	ctrl = gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	return cfg, ctrl, pg, redis
}

type mute struct{}

func (n mute) Printf(string, ...interface{}) {}

func createPGContainer(t *testing.T, cfg *config.Config) *postgrescontainers.PostgresContainer {
	t.Helper()

	pgContainer, err := postgrescontainers.RunContainer(context.Background(),
		testcontainers.WithLogger(&mute{}),
		testcontainers.WithImage("postgres:16.2"),
		postgrescontainers.WithDatabase(cfg.GetPostgresDBName()),
		postgrescontainers.WithUsername(cfg.GetPostgresUsername()),
		postgrescontainers.WithPassword(cfg.GetPostgresPassword()),
		testcontainers.WithWaitStrategy(
			wait.
				ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	require.NoError(t, err)
	return pgContainer
}

func updateConfigPGPort(t *testing.T, cfg *config.Config, pgContainer *postgrescontainers.PostgresContainer) {
	t.Helper()

	dbURL, err := pgContainer.ConnectionString(context.Background())
	require.NoError(t, err)

	url, err := url.Parse(dbURL)
	require.NoError(t, err)

	cfg.SetPostgresPort(url.Port())
}

func updateConfigRedisPort(t *testing.T, cfg *config.Config, miniRedis *miniredis.Miniredis) {
	t.Helper()

	cfg.SetRedisPort(miniRedis.Port())
}
