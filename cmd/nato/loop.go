package nato

import (
	"github.com/niradler/nato/pkg/nato"
	"github.com/spf13/cobra"
)

var separator string
var pattern string
var command string
var dryRun bool

var loopCmd = &cobra.Command{
	Use:     "loop",
	Aliases: []string{"lp"},
	Short:   "loop over string",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var data string
		data = nato.GetStdin()
		if data == "" {
			data = args[0]
		}
		nato.Loop(nato.LoopArgs{data, pattern, separator, command, dryRun})
	},
}

func init() {
	loopCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "Dry run print to console instead of execute")
	loopCmd.Flags().StringVarP(&separator, "separator", "s", "\n", "String separator")
	loopCmd.Flags().StringVarP(&pattern, "pattern", "p", "fields", "String concat pattern")
	loopCmd.Flags().StringVarP(&command, "command", "c", "echo Value: {{.Value}}, Index: {{.Index}}", "Command to run on each slice")
	rootCmd.MarkFlagRequired("command")
	rootCmd.AddCommand(loopCmd)
}
