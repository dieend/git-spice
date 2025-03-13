package main

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/dieend/git-spice/internal/forge"
	"github.com/dieend/git-spice/internal/git"
	"github.com/dieend/git-spice/internal/spice"
	"github.com/dieend/git-spice/internal/spice/state"
	"github.com/dieend/git-spice/internal/text"
)

type logLongCmd struct {
	branchLogCmd
}

func (*logLongCmd) Help() string {
	return text.Dedent(`
		Only branches that are upstack and downstack from the current
		branch are shown.
		Use with the -a/--all flag to show all tracked branches.
	`)
}

func (cmd *logLongCmd) Run(
	ctx context.Context,
	log *log.Logger,
	repo *git.Repository,
	store *state.Store,
	svc *spice.Service,
	forges *forge.Registry,
) (err error) {
	return cmd.run(ctx, &branchLogOptions{
		Log:     log,
		Commits: true,
	}, repo, store, svc, forges)
}
