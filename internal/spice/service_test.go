package spice

import (
	"testing"

	"github.com/charmbracelet/log"
	"github.com/dieend/git-spice/internal/forge"
	"github.com/dieend/git-spice/internal/logutil"
	"github.com/dieend/git-spice/internal/spice/state"
	"github.com/dieend/git-spice/internal/spice/state/storage"
	"github.com/stretchr/testify/require"
)

// NewTestService creates a new Service for testing.
// If forge is nil, it uses the ShamHub forge.
func NewTestService(
	repo GitRepository,
	store Store,
	forgeReg *forge.Registry,
	log *log.Logger,
) *Service {
	return newService(repo, store, forgeReg, log)
}

// NewMemoryStore builds gs state storage
// that stores everything in memory.
// The store is initialized with the trunk "main".
func NewMemoryStore(t *testing.T) *state.Store {
	t.Helper()

	ctx := t.Context()
	db := storage.NewDB(make(storage.MapBackend))
	store, err := state.InitStore(ctx, state.InitStoreRequest{
		DB:    db,
		Trunk: "main",
		Log:   logutil.TestLogger(t),
	})
	require.NoError(t, err)

	return store
}
