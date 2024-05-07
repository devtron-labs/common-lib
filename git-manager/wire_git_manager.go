package git_manager

import (
	"github.com/google/wire"
)

var GitManagerWireSet = wire.NewSet(
	NewGitManagerImpl,
	NewGitCliManager,
	wire.Bind(new(GitCliManager), new(*GitCliManagerImpl)),
)
