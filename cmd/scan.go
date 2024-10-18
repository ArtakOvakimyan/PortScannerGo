package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	ip          string
	protocol    string
	targetPorts string
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scans ports in defined ip-address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scan called")
		ip = args[0]
		err := scanPorts(ip, protocol, targetPorts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&protocol, "protocol", "n", "TCP", "Protocol (TCP/UDP)")
	scanCmd.Flags().StringVarP(&targetPorts, "target-ports", "p", "", "Port (e.g. 5000) or port range (e.g. 80:443)")
}
