package command

import "github.com/afengliz/zett/framework/cobra"

func AddKernelCommands(rootCmd *cobra.Command){
	rootCmd.AddCommand(initAppCommand())
	rootCmd.AddCommand(initCronCommand())
}
