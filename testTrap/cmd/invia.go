// Copyright © 2019 Alberto Bregliano
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
	"flag"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var nasname, community, serverip, deviceip, argomento, summary string
var port, specific, severity int

// inviaCmd represents the invia command
var inviaCmd = &cobra.Command{
	Use:   "invia",
	Short: "Invia le trap snmp di Ramses",
	Long: `Genera e invia trap versione 1 di Ramses a server snmp
	Copyright © 2019 Alberto Bregliano`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("invia called")
		flag.Parse()
		portu := uint16(port)
		_, err := CreaTrap(nasname, argomento, summary, deviceip, serverip, community, portu, specific, severity)
		if err != nil {
			log.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(inviaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inviaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inviaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	inviaCmd.Flags().StringVarP(&nasname, "nasname", "n", "r-finto", "il nome di un nas")
	inviaCmd.Flags().StringVarP(&argomento, "argomento", "a", "sessioni_ppp", "argomento della trap")
	inviaCmd.Flags().StringVarP(&summary, "summary", "m", "Forte abbassamento sessioni ppp", "summary della trap")
	inviaCmd.Flags().StringVarP(&community, "community", "c", "public", "la community da usare")
	inviaCmd.Flags().StringVarP(&serverip, "serverip", "s", "127.0.0.1", "ip del server snmp")
	inviaCmd.Flags().StringVarP(&deviceip, "deviceip", "d", "10.10.10.10", "ip del device")
	inviaCmd.Flags().IntVarP(&port, "port", "p", 162, "la porta snmp da usare")
	inviaCmd.Flags().IntVarP(&severity, "severity", "z", 1, "severity della trap")
	inviaCmd.Flags().IntVarP(&specific, "specific", "f", 5, "specific della trap")

}
