package calendar

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	gCalendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const AGENDA = "Roles"

var (
	ErrAddAgenda  = errors.New("Error to add agenda")
	ErrEventsWeek = errors.New("Error to show events week")
)

type Calendar struct {
	Service    *gCalendar.Service
	CalendarId string
}

func NewClient() *Calendar {
	ctx := context.Background()

	credentials, err := os.ReadFile("./credentials.json")
	if err != nil {
		log.Fatal("Unable to read client secret file: %v", err)
	}
	service, err := gCalendar.NewService(ctx, option.WithCredentialsJSON(credentials))
	if err != nil {
		log.Fatal("Error to create google calendar service: %s", err.Error())
	}

	return &Calendar{Service: service}

}

func (c *Calendar) GetAgendaId() error {
	list, err := c.Service.CalendarList.List().Do()
	if err != nil {
		return err
	}
	for _, v := range list.Items {
		if v.Summary == AGENDA {
			c.CalendarId = v.Id
		}
	}

	return nil
}

func (c *Calendar) InsertAgenda(id string) error {
	entry := &gCalendar.CalendarListEntry{
		Id: id,
	}
	_, err := c.Service.CalendarList.Insert(entry).Do()
	if err != nil {
		return ErrAddAgenda
	}

	return nil
}

func (c *Calendar) ListWeekEvents() error {
	now := time.Now()
	weekDay := now.Weekday()
	starDate := now.AddDate(0, 0, -int(weekDay))
	endDate := starDate.AddDate(0, 0, 7)

	events, err := c.Service.Events.List(c.CalendarId).TimeMin(starDate.Format(time.RFC3339)).TimeMax(endDate.Format(time.RFC3339)).Do()

	if err != nil {
		return ErrEventsWeek
	}

	for _, v := range events.Items {
		fmt.Printf("%s |%s  at %s\n ", v.Summary, v.Status, v.Start.DateTime)

	}
	return nil
}
