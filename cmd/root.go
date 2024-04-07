package cmd

import (
	"fmt"
	"os"

	"github.com/Math2121/calendar-go-cli/cmd/agenda"
	"github.com/Math2121/calendar-go-cli/cmd/events"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command{
	rootCmd:= &cobra.Command{
		Use: "calendar-cli",
		Short: "YOUR CALENDAR CLI",
	}

	rootCmd.AddCommand(events.EventsCmd)
	rootCmd.AddCommand(agenda.AgendaCmd)
	return rootCmd
}

func Execute(){
	if err := NewRootCmd().Execute(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}