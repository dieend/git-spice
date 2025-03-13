package main

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/dieend/git-spice/internal/git"
	"github.com/dieend/git-spice/internal/must"
	"github.com/dieend/git-spice/internal/spice"
	"github.com/dieend/git-spice/internal/spice/state"
	"github.com/dieend/git-spice/internal/text"
	"github.com/dieend/git-spice/internal/ui"
	"github.com/dieend/git-spice/internal/ui/widget"
)

type topCmd struct {
	checkoutOptions
}

func (*topCmd) Help() string {
	return text.Dedent(`
		Checks out the top-most branch in the current branch's stack.
		If there are multiple possible top-most branches,
		a prompt will ask you to pick one.
		Use the -n flag to print the branch without checking it out.
	`)
}

func (cmd *topCmd) Run(
	ctx context.Context,
	log *log.Logger,
	view ui.View,
	repo *git.Repository,
	store *state.Store,
	svc *spice.Service,
) error {
	current, err := repo.CurrentBranch(ctx)
	if err != nil {
		// TODO: handle not a branch
		return fmt.Errorf("get current branch: %w", err)
	}

	tops, err := svc.FindTop(ctx, current)
	if err != nil {
		return fmt.Errorf("find top-most branches: %w", err)
	}
	must.NotBeEmptyf(tops, "FindTopmost always returns at least one branch")

	branch := tops[0]
	if len(tops) > 1 {
		desc := "There are multiple top-level branches reachable from the current branch."
		if !ui.Interactive(view) {
			log.Error(desc)
			return errNoPrompt
		}

		items := make([]widget.BranchTreeItem, len(tops))
		for i, b := range tops {
			items[i] = widget.BranchTreeItem{
				Branch: b,
				Base:   current,
			}
		}

		// If there are multiple top-most branches,
		// prompt the user to pick one.
		prompt := widget.NewBranchTreeSelect().
			WithValue(&branch).
			WithItems(items...).
			WithTitle("Pick a branch").
			WithDescription(desc)
		if err := ui.Run(view, prompt); err != nil {
			return fmt.Errorf("a branch is required: %w", err)
		}
	}

	if branch == current && !cmd.DryRun {
		log.Info("Already on the top-most branch in this stack")
		return nil
	}

	return (&branchCheckoutCmd{
		checkoutOptions: cmd.checkoutOptions,
		Branch:          branch,
	}).Run(ctx, log, view, repo, store, svc)
}
