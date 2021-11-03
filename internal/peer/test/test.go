package test

import (
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/spf13/cobra"
)

// Program name
const ProgramName = "peer"

var logger = flogging.MustGetLogger("testCommand")

// Cmd returns the Cobra Command for test
func Cmd() *cobra.Command {
	return testCommand
}

var testCommand = &cobra.Command{
	Use:   "test",
	Short: "test for compile peer",
	Long:  `test a node that interacts with the network.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("----------------Build test by Jingbo------------------")
	},
}

