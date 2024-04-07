package events

import (
	"log"

	"github.com/Math2121/calendar-go-cli/internal/calendar"
	"github.com/spf13/cobra"
)

var EventsWeekCmd = &cobra.Command{
	Use:   "week",
	Short: "List all events of week ",
	Long:  "check all your events of week ",
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		err := c.GetAgendaId()
		if err != nil {
			log.Fatal(err.Error())
		}
		c.ListWeekEvents()
	},
}
