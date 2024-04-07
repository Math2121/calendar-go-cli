package agenda

import (
	"fmt"
	"log"

	"github.com/Math2121/calendar-go-cli/internal/calendar"
	"github.com/spf13/cobra"
)

var AgendaCmd = &cobra.Command{
	Use:   "agenda",
	Short: "List all your agenda",
	Long:  "check all your agenda",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		err := c.InsertAgenda(args[0])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("success")
	},
}
