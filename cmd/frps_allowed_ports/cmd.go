package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofrp/fp-multiuser/pkg/server"

	"github.com/spf13/cobra"
)

const version = "1.0.2"

var (
	showVersion bool
	bindAddr    string
	// tokenFile   string
	portsFile string
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "version")
	rootCmd.PersistentFlags().StringVarP(&bindAddr, "bind_addr", "l", "127.0.0.1:9000", "bind address")
	rootCmd.PersistentFlags().StringVarP(&portsFile, "ports_file", "p", "./ports", "ports file")
}

var rootCmd = &cobra.Command{
	Use:   "fp-multiuser",
	Short: "fp-multiuser is the server plugin of frp to support multiple users.",
	RunE: func(cmd *cobra.Command, args []string) error {

		portslist, _ := ParseportsFromFile(portsFile)
		if showVersion {
			fmt.Println(version)
			return nil
		}
		s, err := server.New(server.Config{
			BindAddress: bindAddr,
			Ports:       portslist,
		})
		if err != nil {
			return err
		}
		s.Run()
		return nil

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func ParseportsFromFile(file string) (map[string][]string, error) {
	buf, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	result := make(map[string][]string)
	rows := strings.Split(string(buf), "\n")
	for _, row := range rows {
		parts := strings.Split(row, "=")
		key := parts[0]
		value := parts[1]
		result[key] = append(result[key], value)
	}
	return result, nil

}
