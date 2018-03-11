// Copyright © 2018 Levi Gross <levi@levigross.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/levigross/sshproxy/pkg/ssh"
	"github.com/spf13/cobra"
)

var sshConfig ssh.Config

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH protocol proxy",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return sshConfig.Run()
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)

	sshCmd.PersistentFlags().IntVar(&sshConfig.HostPort, "hostp", 2002, "The port you want to listen on")
	sshCmd.PersistentFlags().IntVar(&sshConfig.DstPort, "dstp", 22, "The dest port")
	sshCmd.PersistentFlags().StringVar(&sshConfig.DstHostname, "dst", "", "The dest hostname")
}
