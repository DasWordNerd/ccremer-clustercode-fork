package main

import (
	"github.com/ccremer/clustercode/pkg/cleanupcmd"
	"github.com/urfave/cli/v2"
)

func newCleanupCommand() *cli.Command {
	command := cleanupcmd.Command{}
	return &cli.Command{
		Name:   "cleanup",
		Usage:  "Remove intermediary files and finish the task",
		Before: LogMetadata,
		Action: func(ctx *cli.Context) error {
			command.Log = AppLogger(ctx).WithName(ctx.Command.Name)
			return command.Execute(ctx.Context)
		},
		Flags: []cli.Flag{
			newTaskNameFlag(&command.TaskName),
			newNamespaceFlag(&command.Namespace),
			newSourceRootDirFlag(&command.SourceRootDir),
		},
	}
}
