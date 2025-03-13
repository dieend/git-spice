package main

import (
	"github.com/charmbracelet/log"
	"github.com/dieend/git-spice/internal/forge"
	"github.com/dieend/git-spice/internal/secret"
	"github.com/dieend/git-spice/internal/text"
)

type authLogoutCmd struct{}

func (*authLogoutCmd) Help() string {
	return text.Dedent(`
		The stored authentication information is deleted from secure storage.
		Use 'gs auth login' to log in again.

		No-op if not logged in.
	`)
}

func (cmd *authLogoutCmd) Run(
	stash secret.Stash,
	log *log.Logger,
	f forge.Forge,
) error {
	if err := f.ClearAuthenticationToken(stash); err != nil {
		return err
	}

	// TOOD: Forges should present friendly names in addition to IDs.
	log.Infof("%s: logged out", f.ID())
	return nil
}
