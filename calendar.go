package nationbuilder

import (
	"fmt"
	"net/http"
)

type UserEvents struct {
	StartDate        string   `json:"start_date,omitempty"`
	EndDate          string   `json:"end_date,omitempty"`
	DefaultStartTime string   `json:"default_start_time,omitempty"`
	TagList          []string `json:"tag_list,omitempty"`
}

type Calendar struct {
	Page
	Content             string     `json:"content,omitempty"`
	EventName           string     `json:"event_name,omitempty"`
	ShowMap             bool       `json:"show_map,omitempty"`
	Order               string     `json:"order"`
	UserSubmittedEvents UserEvents `json:"user_submitted_events,omitempty"`
}

func (c *Calendar) String() string {
	return fmt.Sprintf("Calendar ID %d - %s", c.ID, c.Name)
}

type Calendars struct {
	Results []*Calendar `json:"results"`
	Pagination
}

type calendarWrap struct {
	Calendar *Calendar `json:"calendar"`
}

// Retrieve a page of Calendars for the given site and blog id
func (n *Client) GetCalendars(siteSlug string, options *Options) (calendars *Calendars, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/calendars", siteSlug)
	req := n.getRequest("GET", u, options)
	result = n.retrieve(req, &calendars)

	return
}

// Retrieve an individual Calendar
func (n *Client) GetCalendar(siteSlug string, calendarID int, options *Options) (calendar *Calendar, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/calendars/%d", siteSlug, calendarID)
	req := n.getRequest("GET", u, options)
	c := &calendarWrap{}
	result = n.retrieve(req, c)
	calendar = c.Calendar

	return
}

// Create a Calendar for the specified site
func (n *Client) CreateCalendar(siteSlug string, calendar *Calendar, options *Options) (newCalendar *Calendar, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/calendars", siteSlug)
	req := n.getRequest("POST", u, options)
	c := &calendarWrap{}
	result = n.create(&calendarWrap{calendar}, req, c, http.StatusOK)
	newCalendar = c.Calendar

	return
}

// Update a Calendar for the specified site and with the specified ID
func (n *Client) UpdateCalendar(siteSlug string, calendarID int, calendar *Calendar, options *Options) (updatedCalendar *Calendar, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/calendars/%d", siteSlug, calendarID)
	req := n.getRequest("PUT", u, options)
	c := &calendarWrap{}
	result = n.create(&calendarWrap{calendar}, req, c, http.StatusOK)
	updatedCalendar = c.Calendar

	return
}

// Delete a Calendar
func (n *Client) DeleteCalendar(siteSlug string, calendarID int, options *Options) (result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/calendars/%d", siteSlug, calendarID)
	req := n.getRequest("DELETE", u, options)
	result = n.delete(req)

	return
}
