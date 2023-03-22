package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/golangci/golangci-lint/pkg/lint/linter"
)

func (e *Executor) initLinters() {
	e.lintersCmd = &cobra.Command{
		Use:               "linters",
		Short:             "List current linters configuration",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE:              e.executeLinters,
	}
	e.rootCmd.AddCommand(e.lintersCmd)
	e.initRunConfiguration(e.lintersCmd)
}

// executeLinters runs the 'linters' CLI command, which displays the supported linters.
func (e *Executor) executeLinters(_ *cobra.Command, _ []string) error {
	enabledLintersMap, err := e.EnabledLintersSet.GetEnabledLintersMap()
	if err != nil {
		return fmt.Errorf("can't get enabled linters: %w", err)
	}

	green("Enabled by your configuration linters:\n")
	enabledLinters := make([]*linter.Config, 0, len(enabledLintersMap))
	for _, l := range enabledLintersMap {
		enabledLinters = append(enabledLinters, l)
	}
	printLinterConfigs(enabledLinters, e.cfg.Output.Format)

	var disabledLCs []*linter.Config
	for _, lc := range e.DBManager.GetAllSupportedLinterConfigs() {
		if enabledLintersMap[lc.Name()] == nil {
			disabledLCs = append(disabledLCs, lc)
		}
	}

	red("\nDisabled by your configuration linters:\n")
	printLinterConfigs(disabledLCs, e.cfg.Output.Format)

	return nil
}
