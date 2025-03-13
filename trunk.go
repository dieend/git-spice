package main

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/dieend/git-spice/internal/git"
	"github.com/dieend/git-spice/internal/spice"
	"github.com/dieend/git-spice/internal/spice/state"
	"github.com/dieend/git-spice/internal/ui"
)

type trunkCmd struct {
	checkoutOptions
}

func (cmd *trunkCmd) Run(
	ctx context.Context,
	log *log.Logger,
	view ui.View,
	repo *git.Repository,
	store *state.Store,
	svc *spice.Service,
) error {
	trunk := store.Trunk()
	return (&branchCheckoutCmd{
		checkoutOptions: cmd.checkoutOptions,
		Branch:          trunk,
	}).Run(ctx, log, view, repo, store, svc)
}
