/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"helm-external-val/util"
	"os"

	"github.com/spf13/cobra"
)

var kubeNamespace string
var output string

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "cm <name>",
	Short: "Get the content of values from a cm and write it to a file",
	Long:  `Get the content of values from a cm and write it to a file`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmName := args[0]
		fmt.Println("cm called", cmName)
		client := util.GetK8sClient()
		cm, err := util.GetConfigMap(kubeNamespace, cmName, client)
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}
		values := util.ComposeValues(cm)
		util.WriteValuesToFile(values, output)
		fmt.Printf("%s written to %s\n", cmName, output)
	},
}

func init() {
	rootCmd.AddCommand(cmCmd)
	cmCmd.PersistentFlags().StringVar(&kubeNamespace, "kube_namespace", "default", "The namespace to get the cm from")
	cmCmd.PersistentFlags().StringVarP(&output, "out", "o", "values-cm.yaml", "The file to output the values to")
}
