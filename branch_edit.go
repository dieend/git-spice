package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/dieend/git-spice/internal/git"
	"github.com/dieend/git-spice/internal/spice"
	"github.com/dieend/git-spice/internal/spice/state"
	"github.com/dieend/git-spice/internal/text"
)

type branchEditCmd struct{}

func (*branchEditCmd) Help() string {
	return text.Dedent(`
		Starts an interactive rebase with only the commits
		in this branch.
		Following the rebase, branches upstack from this branch
		will be restacked.
	`)
}

func (*branchEditCmd) Run(
	ctx context.Context,
	log *log.Logger,
	repo *git.Repository,
	store *state.Store,
	svc *spice.Service,
) error {
	currentBranch, err := repo.CurrentBranch(ctx)
	if err != nil {
		return fmt.Errorf("get current branch: %w", err)
	}

	b, err := svc.LookupBranch(ctx, currentBranch)
	if err != nil {
		if errors.Is(err, state.ErrNotExist) {
			return fmt.Errorf("branch not tracked: %s", currentBranch)
		}
		return fmt.Errorf("get branch: %w", err)
	}

	req := git.RebaseRequest{
		Interactive: true,
		Branch:      currentBranch,
		Upstream:    b.Base,
	}
	if err := repo.Rebase(ctx, req); err != nil {
		// if the rebase is interrupted,
		// recover with an 'upstack restack' later.
		return svc.RebaseRescue(ctx, spice.RebaseRescueRequest{
			Err:     err,
			Command: []string{"upstack", "restack"},
			Branch:  currentBranch,
			Message: "interrupted: edit branch " + currentBranch,
		})
	}

	return (&upstackRestackCmd{}).Run(ctx, log, repo, store, svc)
}
