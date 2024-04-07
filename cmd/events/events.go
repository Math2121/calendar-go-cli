package events

import (
	"log"

	"github.com/Math2121/calendar-go-cli/internal/calendar"
	"github.com/spf13/cobra"
)
func init() {
	EventsCmd.AddCommand(EventsWeekCmd)
}
var EventsCmd = &cobra.Command{
	Use: "events",
	Short: "List all calendar events",
	Long: "check all your events",
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		 err := c.GetAgendaId()
		if err != nil {
			log.Fatal(err.Error())
		}

	},
}