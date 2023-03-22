package commands

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/golangci/golangci-lint/pkg/lint/linter"
	"github.com/golangci/golangci-lint/pkg/logutils"
)

func (e *Executor) initHelp() {
	helpCmd := &cobra.Command{
		Use:   "help",
		Short: "Help",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}
	e.rootCmd.SetHelpCommand(helpCmd)

	lintersHelpCmd := &cobra.Command{
		Use:               "linters",
		Short:             "Help about linters",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		Run:               e.executeLintersHelp,
	}
	helpCmd.AddCommand(lintersHelpCmd)
}

func printLinterConfigs(lcs []*linter.Config, format ...string) {
	sort.Slice(lcs, func(i, j int) bool {
		return lcs[i].Name() < lcs[j].Name()
	})

	var f string
	if len(format) > 0 {
		f = format[0]

		type l struct {
			Name           string `json:"name" yaml:"name"`
			Desc           string `json:"desc" yaml:"desc"`
			*linter.Config `json:",inline" yaml:",inline"`
		}
		converted := make([]l, 0, len(lcs))
		for _, lc := range lcs {
			converted = append(converted, l{
				Name:   lc.Name(),
				Desc:   lc.Linter.Desc(),
				Config: lc,
			})
		}

		switch f {
		case config.OutFormatJSON:
			_ = json.NewEncoder(logutils.StdOut).Encode(converted)
			return
		case config.OutFormatYAML:
			_ = yaml.NewEncoder(logutils.StdOut).Encode(converted)
			return
		}
	}

	for _, lc := range lcs {
		altNamesStr := ""
		if len(lc.AlternativeNames) != 0 {
			altNamesStr = fmt.Sprintf(" (%s)", strings.Join(lc.AlternativeNames, ", "))
		}

		// If the linter description spans multiple lines, truncate everything following the first newline
		linterDescription := lc.Linter.Desc()
		firstNewline := strings.IndexRune(linterDescription, '\n')
		if firstNewline > 0 {
			linterDescription = linterDescription[:firstNewline]
		}

		deprecatedMark := ""
		if lc.IsDeprecated() {
			deprecatedMark = " [" + color.RedString("deprecated") + "]"
		}

		_, _ = fmt.Fprintf(logutils.StdOut, "%s%s%s: %s [fast: %t, auto-fix: %t]\n",
			color.YellowString(lc.Name()),
			altNamesStr, deprecatedMark, linterDescription, !lc.IsSlowLinter(), lc.CanAutoFix)
	}
}

func (e *Executor) executeLintersHelp(_ *cobra.Command, _ []string) {
	var enabledLCs, disabledLCs []*linter.Config
	for _, lc := range e.DBManager.GetAllSupportedLinterConfigs() {
		if lc.Internal {
			continue
		}

		if lc.EnabledByDefault {
			enabledLCs = append(enabledLCs, lc)
		} else {
			disabledLCs = append(disabledLCs, lc)
		}
	}

	green("Enabled by default linters:\n")
	printLinterConfigs(enabledLCs)

	red("\nDisabled by default linters:\n")
	printLinterConfigs(disabledLCs)

	green("\nLinters presets:")
	for _, p := range e.DBManager.AllPresets() {
		linters := e.DBManager.GetAllLinterConfigsForPreset(p)
		var linterNames []string
		for _, lc := range linters {
			if lc.Internal {
				continue
			}

			linterNames = append(linterNames, lc.Name())
		}
		sort.Strings(linterNames)
		_, _ = fmt.Fprintf(logutils.StdOut, "%s: %s\n",
			color.YellowString(p), strings.Join(linterNames, ", "))
	}
}
