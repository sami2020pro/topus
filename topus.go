package main

import (
	"os"
	"topus/src/topus/repl"
)

func main() {
	// flags := struct {
	// 	Interactive struct {
	// 		Settings bool `settings:"true" allow-unknown-arg="true"`
	// 	} `command:"i" description:"Interactive Topus"`
	// }

	// flag.HandleFlag("i", func(cmd *flag.Cmd, args []string) error {
	// 	fmt.Printf("%s\n", strings.Join(cmd.FlagArgs("i")[1:], " "))
	// 	return nil
	// })

	// flag.New(flag.Options{
	// 	Name: 	 	 "Topus",
	// 	Version: 	 "",
	// 	Description: "Topus programming language",
	// 	Flags: *flags,
	// 	ConfigType: flag.ConfigTypeAuto,
	// })

	repl.Start(os.Stdin, os.Stdout)
}
