package main

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/dieend/git-spice/internal/forge"
	"github.com/dieend/git-spice/internal/git"
	"github.com/dieend/git-spice/internal/secret"
	"github.com/dieend/git-spice/internal/spice"
	"github.com/dieend/git-spice/internal/spice/state"
	"github.com/dieend/git-spice/internal/text"
	"github.com/dieend/git-spice/internal/ui"
)

type stackSubmitCmd struct {
	submitOptions
}

func (*stackSubmitCmd) Help() string {
	return text.Dedent(`
		Change Requests are created or updated
		for all branches in the current stack.
	`) + "\n" + _submitHelp
}

func (cmd *stackSubmitCmd) Run(
	ctx context.Context,
	secretStash secret.Stash,
	log *log.Logger,
	view ui.View,
	repo *git.Repository,
	store *state.Store,
	svc *spice.Service,
	forges *forge.Registry,
) error {
	currentBranch, err := repo.CurrentBranch(ctx)
	if err != nil {
		return fmt.Errorf("get current branch: %w", err)
	}

	stack, err := svc.ListStack(ctx, currentBranch)
	if err != nil {
		return fmt.Errorf("list stack: %w", err)
	}

	// TODO: generalize into a service-level method
	// TODO: separate preparation of the stack from submission

	session := newSubmitSession(repo, store, secretStash, forges, view, log)
	for _, branch := range stack {
		if branch == store.Trunk() {
			continue
		}

		err := (&branchSubmitCmd{
			submitOptions: cmd.submitOptions,
			Branch:        branch,
		}).run(ctx, session, log, view, repo, store, svc)
		if err != nil {
			return fmt.Errorf("submit %v: %w", branch, err)
		}
	}

	if cmd.DryRun {
		return nil
	}

	return updateNavigationComments(
		ctx,
		store,
		svc,
		log,
		cmd.NavigationComment,
		session,
	)
}
