package nationbuilder

import (
	"fmt"
	"net/http"
)

type Membership struct {
	Name         string `json:"name"`
	PersonID     int    `json:"person_id"`
	Status       string `json:"status,omitempty"`
	StatusReason string `json:"status_reason,omitempty"`
	ExpiresOn    *Date  `json:"expires_on,omitempty"`
	StartedAt    *Date  `json:"started_at,omitempty"`
	CreatedAt    *Date  `json:"created_at,omitempty"`
	UpdatedAt    *Date  `json:"updated_at,omitempty"`
}

type Memberships struct {
	Results []*Membership `json:"results"`
	Pagination
}

type membershipWrap struct {
	Membership *Membership `json:"membership"`
}

func (n *Client) GetMemberships(personID int, options *Options) (*Memberships, *Result) {
	url := fmt.Sprintf("/people/%d/memberships", personID)
	req := n.getRequest("GET", url, options)
	memberships := &Memberships{}
	result := n.retrieve(req, memberships)

	return memberships, result
}

func (n *Client) CreateMembership(personID int, membership *Membership, options *Options) (*Membership, *Result) {
	url := fmt.Sprintf("/people/%d/memberships", personID)
	req := n.getRequest("POST", url, options)
	mr := &membershipWrap{}

	result := n.create(&membershipWrap{membership}, req, mr, http.StatusOK)

	return mr.Membership, result
}

func (n *Client) UpdateMembership(personID int, membership *Membership, options *Options) (*Membership, *Result) {
	url := fmt.Sprintf("/people/%d/memberships", personID)
	req := n.getRequest("PUT", url, options)
	mr := &membershipWrap{}

	result := n.create(&membershipWrap{membership}, req, mr, http.StatusOK)

	return mr.Membership, result
}
