package nationbuilder

import (
	"fmt"
	"net/http"
	"strconv"
)

// An EventContact describes a contact for an event and is used within the
// event type
type EventContact struct {
	Name      string `json:"name"`
	Phone     string `json:"phone,omitempty"`
	ShowPhone bool   `json:"show_phone,omitempty"`
	Email     string `json:"email,omitempty"`
	ShowEmail bool   `json:"show_email,omitempty"`
}

// An RSVPForm allows for customisation of the form used on event pages
// Phone and Address fields take a string value of e.g. 'hidden'
type RSVPForm struct {
	Phone            string `json:"phone,omitempty"`
	Address          string `json:"address,omitempty"`
	AllowGuests      bool   `json:"allow_guests,omitempty"`
	AcceptRSVPs      bool   `json:"accept_rsvps,omitempty"`
	GatherVolunteers bool   `json:"gather_volunteers,omitempty"`
}

// A Venue type for use within events
type Venue struct {
	Name    string   `json:"string,omitempty"`
	Address *Address `json:"address,omitempty"`
}

// An AutoResponse represents the mail sent when a user takes some action
// on, e.g. an event
type AutoResponse struct {
	BroadCasterID int    `json:"broadcaster_id"`
	Subject       string `json:"subject"`
	Body          string `json:"body"`
}

// A Shift for e.g. a volunteer at an event
type Shift struct {
	ID        int    `json:"id"`
	StartTime *Date  `json:"start_time"`
	EndTime   *Date  `json:"end_time"`
	TimeZone  string `json:"time_zone"`
	Goal      int    `json:"goal"`
}

// Event is a representation of Nationbuilder's event type
// events are not directly created under a calendar since a calendar
// can display all a nation's events, rather an event has a reference to the
// calendar to which it belongs via calendar_id
type Event struct {
	Page
	Intro        string        `json:"intro,omitempty"`
	CalendarID   int           `json:"calendar_id,omitempty"`
	Contact      EventContact  `json:"contact,omitempty"`
	StartTime    *Date         `json:"start_time"`
	EndTime      *Date         `json:"end_time"`
	TimeZone     string        `json:"time_zone,omitempty"`
	RSVPForm     *RSVPForm     `json:"rsvp_form,omitempty"`
	Capacity     int           `json:"capacity,omitempty"`
	ShowGuests   bool          `json:"show_guests,omitempty"`
	Venue        *Venue        `json:"venue,omitempty"`
	AutoResponse *AutoResponse `json:"autoresponse,omitempty"`
	Shifts       []Shift       `json:"shifts,omitempty"`
}

// EventOptions provides for special event based query options
// To not query by calendar ID provide a negative integer
type EventOptions struct {
	Starting   *Date
	Until      *Date
	CalendarID int
	Tags       []string
}

// Events represents a page of event results
type Events struct {
	Results []*Event `json:"results"`
	Pagination
}

type eventWrap struct {
	Event *Event `json:"event"`
}

// GetEvents retrieves a page of Events for the given site
// Events are queried with special options for the start time, end time and the
// calendar ID
func (n *Client) GetEvents(siteSlug string, eventOptions *EventOptions, options *Options) (events *Events, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/events", siteSlug)

	if eventOptions.Starting != nil {
		options.SetQueryOption("starting", eventOptions.Starting.String())
	}
	if eventOptions.Until != nil {
		options.SetQueryOption("until", eventOptions.Until.String())
	}
	if eventOptions.CalendarID > 0 {
		options.SetQueryOption("calendar_id", strconv.Itoa(eventOptions.CalendarID))
	}
	req := n.getRequest("GET", u, options)
	result = n.retrieve(req, &events)

	return
}

// CreateEvent creates an event for the specified site
func (n *Client) CreateEvent(siteSlug string, event *Event, options *Options) (newEvent *Event, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/events", siteSlug)
	req := n.getRequest("POST", u, options)
	e := &eventWrap{}
	result = n.create(&eventWrap{event}, req, e, http.StatusOK)
	newEvent = e.Event

	return
}

// // Update a Calendar for the specified site and with the specified ID
// func (n *Client) UpdateCalendar(siteSlug string, calendarID int, calendar *Calendar, options *Options) (updatedCalendar *Calendar, result *Result) {
// 	u := fmt.Sprintf("/sites/%s/pages/calendars/%d", siteSlug, calendarID)
// 	req := n.getRequest("PUT", u, options)
// 	c := &calendarWrap{}
// 	result = n.create(&calendarWrap{calendar}, req, c, http.StatusOK)
// 	updatedCalendar = c.Calendar
//
// 	return
// }
//
// // Delete a Calendar
// func (n *Client) DeleteCalendar(siteSlug string, calendarID int, options *Options) (result *Result) {
// 	u := fmt.Sprintf("/sites/%s/pages/calendars/%d", siteSlug, calendarID)
// 	req := n.getRequest("DELETE", u, options)
// 	result = n.delete(req)
//
// 	return
// }
