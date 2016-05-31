package nationbuilder

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

const apiKey = "testkey"
const slug = "testslug"
const siteSlug = "testSite"
const testName = "testName"
const testID = 1

var c *Client

var (
	basicPagesURL   = fmt.Sprintf("/api/v1/sites/%s/pages/basic_pages", siteSlug)
	basicPageURL    = fmt.Sprintf("/api/v1/sites/%s/pages/basic_pages/%d", siteSlug, testID)
	blogPostsURL    = fmt.Sprintf("/api/v1/sites/%s/pages/blogs/%d/posts", siteSlug, testID)
	blogPostURL     = fmt.Sprintf("/api/v1/sites/%s/pages/blogs/%d/posts/%d", siteSlug, testID, testID)
	blogsURL        = fmt.Sprintf("/api/v1/sites/%s/pages/blogs", siteSlug)
	blogURL         = fmt.Sprintf("/api/v1/sites/%s/pages/blogs/%d", siteSlug, testID)
	calendarsURL    = fmt.Sprintf("/api/v1/sites/%s/pages/calendars", siteSlug)
	calendarURL     = fmt.Sprintf("/api/v1/sites/%s/pages/calendars/%d", siteSlug, testID)
	peopleURL       = "/api/v1/people"
	personURL       = "/api/v1/people/"
	peopleNearbyURL = "/api/v1/people/nearby"
	peopleSearchURL = "/api/v1/people/search"
	personMatchURL  = "/api/v1/people/match"
	personPushURL   = "/api/v1/people/push"
	meURL           = "/api/v1/people/me"
	membershipsURL  = fmt.Sprintf("/api/v1/people/%d/memberships", testID)
	donationsURL    = "/api/v1/donations"
	sitesURL        = "/api/v1/sites"
	attachmentsURL  = fmt.Sprintf("/api/v1/sites/%s/pages/%s/attachments", siteSlug, siteSlug)
	attachmentURL   = fmt.Sprintf("/api/v1/sites/%s/pages/%s/attachments/%d", siteSlug, siteSlug, testID)
)

var testPeople = [][]byte{[]byte(`{
	"person": {
		"active_customer_expires_at": null,
		"active_customer_started_at": null,
		"author": null,
		"author_id": null,
		"auto_import_id": null,
		"availability": "",
		"ballots": [],
		"banned_at": null,
		"billing_address": null,
		"bio": null,
		"birthdate": null,
		"call_status_id": null,
		"call_status_name": null,
		"capital_amount_in_cents": 500,
		"children_count": 0,
		"church": "",
		"city_district": null,
		"city_sub_district": null,
		"civicrm_id": null,
		"closed_invoices_amount_in_cents": null,
		"closed_invoices_count": null,
		"contact_status_id": null,
		"contact_status_name": null,
		"could_vote_status": false,
		"county_district": "London",
		"county_file_id": null,
		"created_at": "2016-04-29T17:48:00+01:00",
		"datatrust_id": null,
		"demo": "",
		"do_not_call": false,
		"do_not_contact": false,
		"donations_amount_in_cents": 0,
		"donations_amount_this_cycle_in_cents": 0,
		"donations_count": 0,
		"donations_count_this_cycle": 0,
		"donations_pledged_amount_in_cents": 0,
		"donations_raised_amount_in_cents": 0,
		"donations_raised_amount_this_cycle_in_cents": 0,
		"donations_raised_count": 0,
		"donations_raised_count_this_cycle": 0,
		"donations_to_raise_amount_in_cents": 0,
		"dw_id": null,
		"email": "sample.address@example.com",
		"email1": "sample.address2@example.com",
		"email1_is_bad": false,
		"email2": null,
		"email2_is_bad": false,
		"email3": null,
		"email3_is_bad": false,
		"email4": null,
		"email4_is_bad": false,
		"email_opt_in": true,
		"employer": "",
		"ethnicity": "",
		"external_id": null,
		"facebook_address": null,
		"facebook_profile_url": null,
		"facebook_updated_at": null,
		"facebook_username": null,
		"fax_number": "",
		"federal_district": "London",
		"federal_donotcall": false,
		"fire_district": null,
		"first_donated_at": null,
		"first_fundraised_at": null,
		"first_invoice_at": null,
		"first_name": "Sample",
		"first_prospect_at": null,
		"first_recruited_at": null,
		"first_supporter_at": "2016-04-29T17:48:00+01:00",
		"first_volunteer_at": null,
		"full_name": "Sample Person",
		"has_facebook": false,
		"home_address": null,
		"id": 1,
		"import_id": null,
		"inferred_party": "",
		"inferred_support_level": null,
		"invoice_payments_amount_in_cents": 0,
		"invoice_payments_referred_amount_in_cents": 0,
		"invoices_amount_in_cents": null,
		"invoices_count": null,
		"is_absentee_voter": null,
		"is_active_voter": null,
		"is_deceased": false,
		"is_donor": false,
		"is_dropped_from_file": null,
		"is_early_voter": null,
		"is_fundraiser": false,
		"is_ignore_donation_limits": false,
		"is_leaderboardable": true,
		"is_mobile_bad": false,
		"is_permanent_absentee_voter": null,
		"is_possible_duplicate": false,
		"is_profile_private": false,
		"is_profile_searchable": true,
		"is_prospect": false,
		"is_supporter": true,
		"is_survey_question_private": false,
		"is_twitter_follower": false,
		"is_volunteer": false,
		"judicial_district": null,
		"labour_region": null,
		"language": "",
		"last_call_id": null,
		"last_contacted_at": null,
		"last_contacted_by": null,
		"last_donated_at": null,
		"last_fundraised_at": null,
		"last_invoice_at": null,
		"last_name": "Person",
		"last_rule_violation_at": null,
		"legal_name": null,
		"linkedin_id": null,
		"locale": "",
		"mailing_address": null,
		"marital_status": "",
		"media_market_name": null,
		"meetup_address": null,
		"membership_expires_at": null,
		"membership_level_name": null,
		"membership_started_at": null,
		"middle_name": null,
		"mobile": "",
		"mobile_normalized": null,
		"mobile_opt_in": true,
		"nbec_guid": null,
		"nbec_precinct": null,
		"nbec_precinct_code": null,
		"ngp_id": null,
		"note": null,
		"note_updated_at": null,
		"occupation": "",
		"outstanding_invoices_amount_in_cents": null,
		"outstanding_invoices_count": null,
		"overdue_invoices_count": 0,
		"page_slug": null,
		"parent": null,
		"parent_id": null,
		"party": "",
		"party_member": false,
		"pf_strat_id": null,
		"phone": "",
		"phone_normalized": null,
		"phone_time": "",
		"precinct_code": null,
		"precinct_id": 1,
		"precinct_name": "Sample Precinct",
		"prefix": null,
		"previous_party": "",
		"primary_address": {
			"address1": null,
			"address2": null,
			"address3": null,
			"carrier_route": null,
			"city": null,
			"country_code": "GB",
			"county": null,
			"delivery_point": null,
			"fips": null,
			"lat": "51.4995692",
			"lng": "-0.1268362",
			"lot": null,
			"sort_sequence": null,
			"state": null,
			"street_name": null,
			"street_number": null,
			"street_prefix": null,
			"street_suffix": null,
			"street_type": null,
			"unit_number": null,
			"zip": "SW1A 0AA",
			"zip4": null,
			"zip5": "SW1A 0AA"
		},
		"primary_email_id": 1,
		"priority_level": null,
		"priority_level_changed_at": null,
		"profile_content": null,
		"profile_content_html": null,
		"profile_headline": null,
		"profile_image_url_ssl": "https://d3n8a8pro7vhmx.cloudfront.net/assets/icons/buddy.png",
		"received_capital_amount_in_cents": 500,
		"recruiter": null,
		"recruiter_id": null,
		"recruits_count": 0,
		"registered_address": null,
		"registered_at": null,
		"religion": "",
		"rnc_id": null,
		"rnc_regid": null,
		"rule_violations_count": 0,
		"salesforce_id": null,
		"school_district": "London",
		"school_sub_district": "Somewhere",
		"sex": "",
		"signup_sources": [],
		"signup_type": 0,
		"spent_capital_amount_in_cents": 0,
		"state_file_id": null,
		"state_lower_district": null,
		"state_upper_district": null,
		"submitted_address": "A Sample Address, London",
		"subnations": [],
		"suffix": null,
		"support_level": null,
		"support_level_changed_at": null,
		"support_probability_score": null,
		"supranational_district": "London Region",
		"tags": [],
		"township": null,
		"turnout_probability_score": null,
		"twitter_address": null,
		"twitter_description": null,
		"twitter_followers_count": null,
		"twitter_friends_count": null,
		"twitter_id": null,
		"twitter_location": null,
		"twitter_login": null,
		"twitter_name": null,
		"twitter_updated_at": null,
		"twitter_website": null,
		"unsubscribed_at": null,
		"updated_at": "2016-05-11T12:18:39+01:00",
		"user_submitted_address": {
			"address1": null,
			"address2": null,
			"address3": null,
			"carrier_route": null,
			"city": null,
			"country_code": "GB",
			"county": null,
			"delivery_point": null,
			"fips": null,
			"lat": "52.2205396",
			"lng": "0.1744521",
			"lot": null,
			"sort_sequence": null,
			"state": null,
			"street_name": null,
			"street_number": null,
			"street_prefix": null,
			"street_suffix": null,
			"street_type": null,
			"unit_number": null,
			"zip": "SW1A 0AA",
			"zip4": null,
			"zip5": "SW1A 0AA"
		},
		"username": "ausername",
		"van_id": null,
		"village_district": null,
		"voter_updated_at": null,
		"ward": null,
		"warnings_count": 0,
		"website": null,
		"work_address": null,
		"work_phone_number": ""
	},
	"precinct": {
		"code": null,
		"id": 1,
		"name": "Sample precinct name"
	}
}`),
	[]byte(`{
  "person": {
    "birthdate": "1973-04-04",
    "city_district": "city_district_1",
    "civicrm_id": 123489,
    "county_district": "county_district_1",
    "county_file_id": "mycountyfileid",
    "created_at": "2014-02-02T11:51:23-05:00",
    "do_not_call": false,
    "do_not_contact": false,
    "dw_id": 987,
    "email": "test@example.com",
    "email_opt_in": true,
    "employer": "Texas Instruments",
    "external_id": "8491",
    "federal_district": "federal_district_1",
    "fire_district": "fire_district_1",
    "first_name": "Bob",
    "has_facebook": true,
    "id": 672,
    "is_twitter_follower": false,
    "is_volunteer": false,
    "judicial_district": "judicial_district_1",
    "labour_region": "labour_region_1",
    "last_name": "Smith",
    "linkedin_id": "linkedinid",
    "mobile": "5738882938",
    "mobile_opt_in": true,
    "nbec_guid": "c4545",
    "ngp_id": "myngpid",
    "note": "Has 3 kids. Ask about the baby on the way",
    "occupation": "Programmer",
    "party": "P",
    "pf_strat_id": 876,
    "phone": "2156726335",
    "precinct_id": 21,
    "primary_address": null,
    "recruiter_id": 3421,
    "rnc_id": 1234567,
    "rnc_regid": "003F4ED9-3C5C-4E15-8178-B84920F41BFF",
    "salesforce_id": "salesforceid",
    "school_district": "school_district_1",
    "school_sub_district": "school_sub_district_1",
    "sex": "F",
    "signup_type": 0,
    "state_file_id": "a12345z",
    "state_lower_district": "state_lower_district_1",
    "state_upper_district": "state_upper_district_1",
    "support_level": 1,
    "supranational_district": "region_abc",
    "tags": [],
    "twitter_id": "765",
    "twitter_name": "AngryPig",
    "updated_at": "2014-02-02T11:51:23-05:00",
    "van_id": "d6767",
    "village_district": "village_district_1",
    "active_customer_expires_at": null,
    "active_customer_started_at": null,
    "author": null,
    "author_id": null,
    "auto_import_id": null,
    "availability": "anytime",
    "banned_at": null,
    "billing_address": null,
    "bio": "Specialist in management and business growth.",
    "call_status_id": 5,
    "call_status_name": "no-answer",
    "capital_amount_in_cents": 0,
    "children_count": 2,
    "church": "Liberty Church",
    "city_sub_district": "DOWNTOWN LOS ANGELES",
    "closed_invoices_amount_in_cents": null,
    "closed_invoices_count": null,
    "contact_status_id": 3,
    "contact_status_name": "Answered",
    "could_vote_status": 1,
    "demo": "W",
    "donations_amount_in_cents": 0,
    "donations_amount_this_cycle_in_cents": 0,
    "donations_count": 0,
    "donations_count_this_cycle": 0,
    "donations_pledged_amount_in_cents": 0,
    "donations_raised_amount_in_cents": 0,
    "donations_raised_amount_this_cycle_in_cents": 0,
    "donations_raised_count": 0,
    "donations_raised_count_this_cycle": 0,
    "donations_to_raise_amount_in_cents": 0,
    "email1": "test@example.com",
    "email1_is_bad": false,
    "email2": null,
    "email2_is_bad": false,
    "email3": null,
    "email3_is_bad": false,
    "email4": null,
    "email4_is_bad": false,
    "ethnicity": null,
    "facebook_address": null,
    "facebook_profile_url": "http://www.facebook.com/profile.php?id=124071",
    "facebook_updated_at": "2014-02-25T11:22:12-04:00",
    "facebook_username": "alice",
    "fax_number": null,
    "federal_donotcall": false,
    "first_donated_at": "2013-02-02T11:51:23-05:00",
    "first_fundraised_at": "2013-04-02T11:51:23-05:00",
    "first_invoice_at": "2013-01-02T11:51:23-05:00",
    "first_prospect_at": "2013-11-02T11:51:23-05:00",
    "first_recruited_at": "2013-05-02T11:51:23-05:00",
    "first_supporter_at": "2014-02-02T11:51:23-05:00",
    "first_volunteer_at": "2013-05-02T11:51:23-05:00",
    "full_name": "Bob Smith",
    "home_address": null,
    "import_id": null,
    "inferred_party": "H",
    "inferred_support_level": 1,
    "invoice_payments_amount_in_cents": 0,
    "invoice_payments_referred_amount_in_cents": 0,
    "invoices_amount_in_cents": null,
    "invoices_count": null,
    "is_deceased": false,
    "is_donor": false,
    "is_fundraiser": false,
    "is_ignore_donation_limits": false,
    "is_leaderboardable": true,
    "is_mobile_bad": false,
    "is_possible_duplicate": false,
    "is_profile_private": false,
    "is_profile_searchable": true,
    "is_prospect": false,
    "is_supporter": true,
    "is_survey_question_private": false,
    "language": "EN",
    "last_call_id": 2378,
    "last_contacted_at": null,
    "last_contacted_by": null,
    "last_donated_at": "2014-02-02T11:51:23-05:00",
    "last_fundraised_at": "2014-01-01T11:51:23-05:00",
    "last_invoice_at": null,
    "last_rule_violation_at": null,
    "legal_name": null,
    "locale": "en",
    "mailing_address": null,
    "marital_status": "S",
    "media_market_name": null,
    "meetup_address": null,
    "membership_expires_at": null,
    "membership_level_name": null,
    "membership_started_at": null,
    "middle_name": "T.",
    "mobile_normalized": null,
    "nbec_precinct_code": null,
    "note_updated_at": "2014-01-01T11:51:23-05:00",
    "outstanding_invoices_amount_in_cents": null,
    "outstanding_invoices_count": null,
    "overdue_invoices_count": 0,
    "page_slug": null,
    "parent": null,
    "parent_id": null,
    "party_member": false,
    "phone_normalized": null,
    "phone_time": "daytime",
    "precinct_code": "9876",
    "precinct_name": "9876",
    "prefix": null,
    "previous_party": "U",
    "primary_email_id": 1,
    "priority_level": 4,
    "priority_level_changed_at": null,
    "profile_content": "I support Abraham Lincoln for President.",
    "profile_content_html": "<p>I support Abraham Lincoln for President.</p>",
    "profile_headline": "Abe is back!",
    "received_capital_amount_in_cents": 0,
    "recruiter": {
      "birthdate": "1962-03-03",
      "city_district": null,
      "civicrm_id": 9,
      "county_district": "11",
      "county_file_id": "97076395",
      "created_at": "2014-02-02T11:51:23-05:00",
      "do_not_call": false,
      "do_not_contact": false,
      "dw_id": null,
      "email": null,
      "email_opt_in": true,
      "employer": "Sun Film Foundation",
      "external_id": "781986",
      "federal_district": "0",
      "fire_district": null,
      "first_name": "Mina",
      "has_facebook": false,
      "id": 3421,
      "is_twitter_follower": false,
      "is_volunteer": false,
      "judicial_district": null,
      "labour_region": null,
      "last_name": "Abbott",
      "linkedin_id": null,
      "mobile": "5738882938",
      "mobile_opt_in": true,
      "nbec_guid": "000099dc-00f0-11e2-b349-27c4eb6c1979",
      "ngp_id": null,
      "note": "Has 3 kids. Ask about the baby on the way",
      "occupation": "Community Organizer",
      "party": "P",
      "pf_strat_id": 91207702,
      "phone": "2156726335",
      "precinct_id": null,
      "primary_address": null,
      "recruiter_id": null,
      "rnc_id": 2653431,
      "rnc_regid": "003F4ED9-3C5C-4E15-8178-B84920F41BFF",
      "salesforce_id": null,
      "school_district": "Chicago Public Schools",
      "school_sub_district": null,
      "sex": "F",
      "signup_type": 0,
      "state_file_id": "100000012",
      "state_lower_district": "024",
      "state_upper_district": "1",
      "support_level": 1,
      "supranational_district": "Eastern Euro Region",
      "tags": [],
      "twitter_id": "63403772",
      "twitter_name": "AngryPig",
      "updated_at": "2014-02-02T11:51:23-05:00",
      "van_id": "2",
      "village_district": null
    },
    "recruits_count": 0,
    "registered_address": null,
    "registered_at": null,
    "religion": null,
    "rule_violations_count": 0,
    "spent_capital_amount_in_cents": 0,
    "submitted_address": "123 Main St Portland, ME 04101",
    "subnations": ["abeforprez", "operalovers"],
    "suffix": null,
    "support_level_changed_at": "2014-02-02T11:51:23-05:00",
    "support_probability_score": null,
    "turnout_probability_score": null,
    "twitter_address": null,
    "twitter_description": "Animal lover.",
    "twitter_followers_count": 564,
    "twitter_friends_count": 32,
    "twitter_location": "Dallas, TX",
    "twitter_login": "trinity",
    "twitter_updated_at": "2014-01-01T11:51:23-05:00",
    "twitter_website": "http://mytwittersite.com",
    "unsubscribed_at": null,
    "user_submitted_address": null,
    "username": null,
    "warnings_count": 0,
    "website": null,
    "work_address": null,
    "work_phone_number": null
  },
  "precinct": {
    "id": 21,
    "code": "9876",
    "name": "9876"
  }
}`)}

func attachmentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(&attachmentWrap{
			Attachment: &Attachment{
				FileName: testName,
			},
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	case "DELETE":
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func attachmentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		d, _ := NewDate(testTime)
		err := json.NewEncoder(w).Encode(&Attachments{
			Results: []*Attachment{
				&Attachment{
					FileName:    testName,
					UpdatedAt:   d,
					ID:          testID,
					ContentType: "image/jpeg",
					URL:         "/foo/file.jpg",
				},
			},
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	case "POST":
		foo64 := base64.StdEncoding.EncodeToString([]byte("foo"))
		upload := &uploadWrap{}
		err := json.NewDecoder(r.Body).Decode(upload)
		if err != nil {
			log.Fatal(err.Error())
		}
		if upload.Attachment.Content != foo64 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(&attachmentWrap{
			Attachment: &Attachment{
				FileName:    upload.Attachment.FileName,
				UpdatedAt:   upload.Attachment.UpdatedAt,
				URL:         "/foo/foo.jpg",
				ContentType: upload.Attachment.ContentType,
			},
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func sitesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(&Sites{
			Results: []*Site{
				&Site{
					Name: testName,
				},
			},
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func personMatchHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		q := r.URL.Query()
		fName := q.Get("first_name")
		lName := q.Get("last_name")
		email := q.Get("email")
		if fName == "" || lName == "" || email == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if fName == "MatchMissingF" && lName == "MatchMissingL" && email == "match-missing@example.com" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := json.NewEncoder(w).Encode(&personWrap{
			Person: &Person{
				FirstName: fName,
				LastName:  lName,
				Email:     email,
			},
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func peopleSearchHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		q := r.URL.Query()
		fName := q.Get("first_name")
		lName := q.Get("last_name")
		if fName == "" || lName == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := json.NewEncoder(w).Encode(&People{
			Results: []*Person{
				&Person{
					FirstName: "Wally",
					LastName:  "Waldo",
				},
			},
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func peopleNearbyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		q := r.URL.Query()
		loc := q.Get("location")
		if loc == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		dist := q.Get("distance")
		if dist == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := json.NewEncoder(w).Encode(&People{
			Results: []*Person{
				&Person{
					FirstName: "Wally",
					LastName:  "Waldo",
				},
			},
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		pid, err := strconv.Atoi(r.URL.Path[len(r.URL.Path)-1:])
		if err != nil {
			log.Fatal(err.Error())
		}
		_, err = w.Write(testPeople[pid])
		if err != nil {
			log.Fatal(err.Error())
		}
	case "PUT":
		pw := &personWrap{}
		err := json.NewDecoder(r.Body).Decode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = json.NewEncoder(w).Encode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "DELETE":
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		_, err := w.Write(testPeople[0])
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		people := &People{
			Results: []*Person{
				&Person{
					ID:        testID,
					FirstName: "Phileas",
					LastName:  "Fogg",
				},
			},
		}
		err := json.NewEncoder(w).Encode(people)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "POST":
		pw := &personWrap{}
		err := json.NewDecoder(r.Body).Decode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
		pw.Person.ID = testID
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func membershipsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		memberships := &Memberships{
			Results: []*Membership{
				&Membership{
					Name:         "test_membership",
					PersonID:     1,
					Status:       "active",
					StatusReason: "API Test",
					ExpiresOn:    NewDateFromTime(time.Date(2016, time.December, 25, 0, 0, 0, 0, time.UTC)),
					StartedAt:    NewDateFromTime(time.Now()),
				},
			},
		}
		err := json.NewEncoder(w).Encode(memberships)
		if err != nil {
			log.Fatalf(err.Error())
		}
	case "POST":
		mw := &membershipWrap{}
		err := json.NewDecoder(r.Body).Decode(mw)
		if err != nil {
			log.Fatal(err.Error())
		}
		mw.Membership.PersonID = testID
		mw.Membership.CreatedAt = NewDateFromTime(time.Now())
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(mw)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "PUT":
		mw := &membershipWrap{}
		err := json.NewDecoder(r.Body).Decode(mw)
		if err != nil {
			log.Fatal(err.Error())
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(mw)
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func calendarHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		cw := &calendarWrap{
			Calendar: &Calendar{
				Page: Page{
					Name: testName,
					ID:   testID,
				},
			},
		}
		err := json.NewEncoder(w).Encode(cw)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "PUT":
		cw := &calendarWrap{}
		err := json.NewDecoder(r.Body).Decode(cw)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = json.NewEncoder(w).Encode(cw)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "DELETE":
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func donationsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		dw := &donationWrap{}
		err := json.NewDecoder(r.Body).Decode(dw)
		if err != nil {
			log.Fatal(err.Error())
		}

		if dw.Donation.DonorID != testID || dw.Donation.AmountInCents != 100 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		dw.Donation.ID = testID
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(dw)
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func calendarsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c := &Calendars{
			Results: []*Calendar{
				&Calendar{
					Page: Page{
						Name: testName,
						ID:   testID,
					},
				},
			},
		}
		err := json.NewEncoder(w).Encode(c)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "POST":
		cw := &calendarWrap{}
		err := json.NewDecoder(r.Body).Decode(cw)
		if err != nil {
			log.Fatal(err.Error())
		}
		cw.Calendar.ID = testID
		err = json.NewEncoder(w).Encode(cw)
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

var basicPagesHandler = func(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		pages := &BasicPages{
			Results: []*BasicPage{
				&BasicPage{
					Page: Page{
						Name: testName,
					},
				},
			},
		}
		err := json.NewEncoder(w).Encode(pages)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if req.Method == "POST" {
		bpw := &basicPageWrap{}
		err := json.NewDecoder(req.Body).Decode(bpw)
		if err != nil {
			log.Fatal(err.Error())
		}

		if bpw.BasicPage.Name != testName {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			err := json.NewEncoder(w).Encode(bpw)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}

var basicPageHandler = func(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "PUT":
		bpw := &basicPageWrap{}
		err := json.NewDecoder(req.Body).Decode(bpw)
		if err != nil {
			log.Fatal(err.Error())
		}
		page := bpw.BasicPage
		if page.Name != testName {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			err := json.NewEncoder(w).Encode(bpw)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	case "DELETE":
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

var blogPostsHandler = func(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		posts := &BlogPosts{
			Results: []*BlogPost{
				&BlogPost{
					Page: Page{
						Name: testName,
						ID:   testID,
					},
				},
			},
		}
		err := json.NewEncoder(w).Encode(posts)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "POST":
		pw := &blogPostWrap{}
		err := json.NewDecoder(req.Body).Decode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
		pw.BlogPost.ID = testID
		err = json.NewEncoder(w).Encode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func blogPostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		p := &BlogPost{
			Page: Page{
				Name: testName,
			},
		}
		err := json.NewEncoder(w).Encode(&blogPostWrap{p})
		if err != nil {
			log.Fatal(err.Error())
		}
	case "PUT":
		pw := &blogPostWrap{}
		err := json.NewDecoder(r.Body).Decode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
		if pw.BlogPost.Name != testName {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = json.NewEncoder(w).Encode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "DELETE":
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func blogsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		blogs := &Blogs{
			Results: []*Blog{
				&Blog{
					Page: Page{
						Name: testName,
					},
				},
			},
		}

		err := json.NewEncoder(w).Encode(blogs)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "POST":
		bw := &blogWrap{}
		err := json.NewDecoder(r.Body).Decode(bw)
		if err != nil {
			log.Fatal(err.Error())
		}
		bw.Blog.ID = testID
		err = json.NewEncoder(w).Encode(bw)
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		bw := &blogWrap{}
		err := json.NewDecoder(r.Body).Decode(bw)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = json.NewEncoder(w).Encode(bw)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "DELETE":
		w.WriteHeader(http.StatusNoContent)
	}
}

func init() {
	client, err := NewClient(slug, apiKey)
	if err != nil {
		log.Fatal(err.Error())
	}
	c = client

	apiMux := http.NewServeMux()
	apiMux.HandleFunc(basicPagesURL, basicPagesHandler)
	apiMux.HandleFunc(basicPageURL, basicPageHandler)
	apiMux.HandleFunc(blogPostsURL, blogPostsHandler)
	apiMux.HandleFunc(blogPostURL, blogPostHandler)
	apiMux.HandleFunc(blogsURL, blogsHandler)
	apiMux.HandleFunc(blogURL, blogHandler)
	apiMux.HandleFunc(calendarsURL, calendarsHandler)
	apiMux.HandleFunc(calendarURL, calendarHandler)
	apiMux.HandleFunc(peopleURL, peopleHandler)
	apiMux.HandleFunc(personURL, personHandler)
	apiMux.HandleFunc(personMatchURL, personMatchHandler)
	apiMux.HandleFunc(peopleNearbyURL, peopleNearbyHandler)
	apiMux.HandleFunc(peopleSearchURL, peopleSearchHandler)
	apiMux.HandleFunc(sitesURL, sitesHandler)
	apiMux.HandleFunc(personPushURL, personHandler)
	apiMux.HandleFunc(meURL, meHandler)
	apiMux.HandleFunc(donationsURL, donationsHandler)
	apiMux.HandleFunc(membershipsURL, membershipsHandler)
	apiMux.HandleFunc(attachmentsURL, attachmentsHandler)
	apiMux.HandleFunc(attachmentURL, attachmentHandler)

	server := httptest.NewServer(apiMux)

	u, err := url.Parse(server.URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	u.Path += "/api/v1"
	c.baseURL = &nationbuilderURL{*u}
}

func TestNationbuilderURLExtendPath(t *testing.T) {
	n, err := NewClient(slug, apiKey)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedPath := "/api/v1"
	url := n.baseURL

	if url.u.Path != expectedPath {
		t.Errorf("Expected path %s but saw %s", expectedPath, url.u.Path)
	}

	url.extendPath("/foo")
	expectedPath = "/api/v1/foo"

	if url.u.Path != expectedPath {
		t.Errorf("Expected extended path to be %s but saw %s", expectedPath, url.u.Path)
	}

	url.extendPath("bar")
	expectedPath = "/api/v1/foo/bar"
	if url.u.Path != expectedPath {
		t.Errorf("Expected extended path to be %s but saw %s", expectedPath, url.u.Path)
	}
}

func TestNewClient(t *testing.T) {
	c, err := NewClient(slug, apiKey)
	if err != nil {
		t.Fatal(err.Error())
	}

	if c.Slug != slug {
		t.Errorf("expected client to have slug %s but saw %s", slug, c.Slug)
	}

	if c.ApiKey != apiKey {
		t.Errorf("expected client to have api key %s but saw %s", apiKey, c.ApiKey)
	}

}

func TestBasicPagesGet(t *testing.T) {
	pages, result := c.GetBasicPages(siteSlug, nil)
	if result.HasError() {
		t.Error(result.Error())
	}

	if result.StatusCode != http.StatusOK {
		t.Errorf("Retrieving basic pages - expect status %d but saw %d", http.StatusOK, result.StatusCode)
	}

	if pages != nil {
		if len(pages.Results) != 1 {
			t.Errorf("Expected number of results to be %d but received %d results", 1, len(pages.Results))
		}

		if pages.Results[0].Name != testName {
			t.Errorf("Expected name of %s but saw %s", testName, pages.Results[0].Name)
		}

	} else {
		t.Error("Expected page of basic pages but received no response")
	}
}

func TestBasicPagesCreate(t *testing.T) {
	bp := &BasicPage{
		Page: Page{
			Name: testName,
		},
	}
	page, result := c.CreateBasicPage(siteSlug, bp, nil)
	if result.HasError() {
		t.Errorf("Unexpected error - code %d, Err: %s", result.StatusCode, result.Error())
	}

	if page == nil {
		t.Error("Expected new page but received a nil response")
	} else {
		if page.Name != testName {
			t.Errorf("Expected page name to be %s but saw %s", testName, page.Name)
		}
	}

	if result.StatusCode != http.StatusOK {
		t.Errorf("NB BasicPages returns a 200 on page creation, not 201 - make sure test reflects that: saw code %d", result.StatusCode)
	}
}

func TestBasicPageUpdate(t *testing.T) {
	bp := &BasicPage{
		Page: Page{
			Name: testName,
		},
	}
	page, result := c.UpdateBasicPage(siteSlug, testID, bp, nil)
	if result.HasError() {
		t.Errorf("Unexpected error - code %d, Err: %s", result.StatusCode, result.Error())
	}

	if page == nil {
		t.Error("Expected new page but received a nil response")
	} else {
		if page.Name != testName {
			t.Errorf("Expected page name to be %s but saw %s", testName, page.Name)
		}
	}

	if result.StatusCode != http.StatusOK {
		t.Errorf("NB BasicPages returns a 200 on page creation, not 201 - make sure test reflects that: saw code %d", result.StatusCode)
	}
}

func TestBasicPageDelete(t *testing.T) {
	result := c.DeleteBasicPage(siteSlug, testID, nil)
	if result.HasError() {
		t.Errorf("Unexpected error - code %d, Err: %s", result.StatusCode, result.Error())
	}

	if result.StatusCode != http.StatusNoContent {
		t.Errorf("NB delete returns a 204 on resource deletion - make sure test reflects that: saw code %d", result.StatusCode)
	}
}

func TestBlogPostsGet(t *testing.T) {
	posts, result := c.GetBlogPosts(siteSlug, testID, nil)
	if result.HasError() {
		t.Errorf("Unexpected error - code %d, Err: %s", result.StatusCode, result.Error())
	}

	if posts == nil {
		t.Error("Unexpected nil posts response")
		t.SkipNow()
	}

	if len(posts.Results) != 1 {
		t.Error("Check API test - should return 1 post")
	}

	pName := posts.Results[0].Name
	if pName != testName {
		t.Errorf("Expected name of %s but saw %s", testName, pName)
	}
}

func TestBlogPostsCreate(t *testing.T) {
	p := &BlogPost{
		Page: Page{
			Name: testName,
		},
	}
	post, result := c.CreateBlogPost(siteSlug, testID, p, nil)
	if result.HasError() {
		t.Errorf("Unexpected error creating blog post: %s", result.Error())
		t.SkipNow()
	}

	if post == nil {
		t.Error("Unexpeced nil post - check test api")
		t.SkipNow()
	}

	name, expected := post.Name, testName
	if name != expected {
		t.Errorf("Expected post to have name %s but saw %s: check API test server", expected, name)
	}
}

func TestBlogPostGet(t *testing.T) {
	post, result := c.GetBlogPost(siteSlug, testID, testID, nil)
	if result.HasError() {
		t.Errorf("Unexpected error creating blog post: %s", result.Error())
		t.SkipNow()
	}

	name, expected := post.Name, testName
	if name != expected {
		t.Errorf("Expected %s but saw %s", expected, name)
	}
}

func TestBlogPostUpdate(t *testing.T) {
	p := &BlogPost{
		Page: Page{
			Name: testName,
		},
	}
	post, result := c.UpdateBlogPost(siteSlug, testID, testID, p, nil)
	if result.HasError() {
		t.Errorf("Unexpected error updating blog post: %s", result.Error())
		t.SkipNow()
	}

	if post == nil {
		t.Error("Unexpeced nil post - check test api")
		t.SkipNow()
	}

	name, expected := post.Name, testName
	if name != expected {
		t.Errorf("Expected post to have name %s but saw %s: check API test server", expected, name)
	}
}

func TestBlogPostDelete(t *testing.T) {
	result := c.DeleteBlogPost(siteSlug, testID, testID, nil)
	if result.HasError() {
		t.Errorf("Unexpected error deleting blog post: %s", result.Error())
		t.SkipNow()
	}

	if result.StatusCode != http.StatusNoContent {
		t.Errorf("NB delete returns a 204 on resource deletion - make sure test reflects that: saw code %d", result.StatusCode)
	}
}

func TestBlogsGet(t *testing.T) {
	blogs, result := c.GetBlogs(siteSlug, nil)
	if result.HasError() {
		t.Errorf("Unexpected error retrieving blog: %s", result.Error())
		t.SkipNow()
	}

	if blogs == nil {
		t.Error("Unexpected nil blogs response")
		t.SkipNow()
	}

	if len(blogs.Results) != 1 {
		t.Error("Check API test - should return 1")
	}

	pName := blogs.Results[0].Name
	if pName != testName {
		t.Errorf("Expected name of %s but saw %s", testName, pName)
	}
}

func TestBlogsCreate(t *testing.T) {
	blog := &Blog{
		Page: Page{
			Name: testName,
		},
	}

	newBlog, result := c.CreateBlog(siteSlug, blog, nil)
	if result.HasError() {
		t.Errorf("Unexpected error creating blog: %s", result.Error())
		t.SkipNow()
	}

	if newBlog == nil {
		t.Error("Unexpected nil newblog response")
		t.SkipNow()
	} else {
		name, expected := newBlog.Name, testName
		if name != expected {
			t.Errorf("Expected new blog to have name %s but saw %s - check test API", testName, name)
		}
	}
}

func TestBlogUpdate(t *testing.T) {
	blog := &Blog{
		Page: Page{
			Name: testName,
		},
	}

	updatedBlog, result := c.UpdateBlog(siteSlug, testID, blog, nil)
	if result.HasError() {
		t.Errorf("Unexpected error retrieving blog: %s", result.Error())
		t.SkipNow()
	}

	if updatedBlog == nil {
		t.Error("Unexpected nil updatedBlog response")
		t.SkipNow()
	} else {
		name, expected := updatedBlog.Name, testName
		if name != expected {
			t.Errorf("Expected updated blog to have name %s but saw %s - check test API", testName, name)
		}
	}
}

func TestBlogDelete(t *testing.T) {
	result := c.DeleteBlog(siteSlug, testID, nil)
	if result.HasError() {
		t.Errorf("Unexpected error deleting blog: %s", result.Error())
		t.SkipNow()
	}

	if result.StatusCode != http.StatusNoContent {
		t.Errorf("NB delete returns a 204 on resource deletion - make sure test reflects that: saw code %d", result.StatusCode)
	}
}

func TestCalendarsGet(t *testing.T) {
	calendars, result := c.GetCalendars(siteSlug, nil)
	if result.HasError() {
		t.Errorf("Unexpected error retrieving calendars: %s", result.Error())
		t.SkipNow()
	}

	if calendars == nil {
		t.Error("Unexpected nil calendars response")
		t.SkipNow()
	}

	if len(calendars.Results) != 1 {
		t.Error("Check API test - should return 1")
	}

	pName := calendars.Results[0].Name
	if pName != testName {
		t.Errorf("Expected name of %s but saw %s", testName, pName)
	}
}

func TestCalendarCreate(t *testing.T) {
	cal := &Calendar{
		Page: Page{
			Name: testName,
		},
	}
	newCal, result := c.CreateCalendar(siteSlug, cal, nil)
	if result.HasError() {
		t.Errorf("Unexpected error creating calendar: %s", result.Error())
		t.SkipNow()
	}

	if newCal == nil {
		t.Error("Unexpected nil calendar response")
		t.SkipNow()
	} else {
		name, expected := newCal.Name, testName
		if name != expected {
			t.Errorf("Expected new calendar to have name %s but saw %s - check test API", testName, name)
		}
	}
}

func TestCalendarUpdate(t *testing.T) {
	cal := &Calendar{
		Page: Page{
			Name: testName,
		},
	}

	updatedCal, result := c.UpdateCalendar(siteSlug, testID, cal, nil)
	if result.HasError() {
		t.Errorf("Unexpected error updating calendar: %s", result.Error())
		t.SkipNow()
	}

	if updatedCal == nil {
		t.Error("Unexpected nil updatedCal response")
		t.SkipNow()
	} else {
		name, expected := updatedCal.Name, testName
		if name != expected {
			t.Errorf("Expected updated blog to have name %s but saw %s - check test API", testName, name)
		}
	}
}

func TestCalendarDelete(t *testing.T) {
	result := c.DeleteCalendar(siteSlug, testID, nil)
	if result.HasError() {
		t.Errorf("Unexpected error deleting calendar: %s", result.Error())
		t.SkipNow()
	}

	if result.StatusCode != http.StatusNoContent {
		t.Errorf("NB delete returns a 204 on resource deletion - make sure test reflects that: saw code %d", result.StatusCode)
	}
}

func TestCalendarGet(t *testing.T) {
	cal, result := c.GetCalendar(siteSlug, testID, nil)
	if result.HasError() {
		t.Errorf("Unexpected error retrieving calendar: %s", result.Error())
		t.SkipNow()
	}

	name, expected := cal.Name, testName
	if name != expected {
		t.Errorf("Expected %s but saw %s", expected, name)
	}
}

func TestDonationsCreate(t *testing.T) {
	donation := &Donation{
		DonorID:         testID,
		AmountInCents:   100,
		SucceededAt:     NewDateFromTime(time.Now()),
		PaymentTypeName: "Cash",
	}
	newDonation, result := c.CreateDonation(donation, nil)
	if result.HasError() {
		t.Errorf("Error when creating donation: %s", result.Error())
		t.SkipNow()
	}
	if newDonation == nil {
		t.Error("Newly created donation is nil")
		t.SkipNow()
	}
	if newDonation.DonorID != testID {
		t.Errorf("Expected donor ID to be %d but saw %d", testID, newDonation.DonorID)
	}
	if newDonation.AmountInCents != 100 {
		t.Errorf("Expected amountincents to be 100 but saw %d", newDonation.AmountInCents)
	}
	if newDonation.ID != testID {
		t.Errorf("Expected donation ID to be %d but saw %d", testID, newDonation.ID)
	}
}

func TestMembershipsGet(t *testing.T) {
	memberships, result := c.GetMemberships(testID, nil)
	if result.HasError() {
		t.Error(result.Error())
		t.SkipNow()
	}

	if result.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 when retrieving memberships but saw %d", result.StatusCode)
		t.SkipNow()
	}

	if memberships == nil {
		t.Error("Received nil when running membership test")
	} else {
		if len(memberships.Results) != 1 {
			t.Errorf("Expected to see 1 membership but saw %d", len(memberships.Results))
		}

		m := memberships.Results[0]
		if m.Name != "test_membership" {
			t.Errorf("Expected membership name test_membership but saw %s", m.Name)
		}

		expiresOn := time.Date(2016, time.December, 25, 0, 0, 0, 0, time.UTC)

		if !m.ExpiresOn.Time.Equal(expiresOn) {
			t.Errorf("Expected membership expiry to equal %s but saw %s", m.ExpiresOn.Time, expiresOn)
		}
	}
}

func TestMembershipCreate(t *testing.T) {
	membership := &Membership{
		Name:   "test_membership",
		Status: "active",
	}
	newM, result := c.CreateMembership(testID, membership, nil)
	if result.HasError() {
		t.Errorf("Error creating membership: %s", result.Error())
		t.SkipNow()
	}
	if newM == nil {
		t.Error("Expected new membership to be returned")
		t.SkipNow()
	} else {
		if newM.PersonID != testID && newM.Name != "test_membership" && newM.Status != "active" {
			t.Error("Returned values for new membership incorrect")
		}
	}

}

func TestMembershipUpdate(t *testing.T) {
	membership := &Membership{
		Name:         "test_membership",
		Status:       "expired",
		StatusReason: "test_reason",
	}

	updatedM, result := c.UpdateMembership(testID, membership, nil)
	if result.HasError() {
		t.Errorf("Error updating membership: %s", result.Error())
		t.SkipNow()
	}

	if updatedM == nil {
		t.Error("Expected updated membership to be returned")
		t.SkipNow()
	} else {
		if updatedM.PersonID != testID && updatedM.Name != "test_membership" && updatedM.Status != "expired" &&
			updatedM.StatusReason != "test_reason" {

			t.Error("Returned values for updated membership incorrect")
		}
	}
}

func TestPeopleGet(t *testing.T) {
	people, result := c.GetPeople(nil)
	if result.HasError() {
		t.Error(result.Error())
		t.SkipNow()
	}

	if result.StatusCode != http.StatusOK {
		t.Errorf("Retrieving people - expect status %d but saw %d", http.StatusOK, result.StatusCode)
	}

	if people != nil {
		if len(people.Results) != 1 {
			t.Errorf("Expected number of results to be %d but received %d results", 1, len(people.Results))
		}

		p := people.Results[0]
		if p.FirstName != "Phileas" && p.LastName != "Fogg" {
			t.Errorf("Expected name of Phileas Fogg but saw %s %s", p.FirstName, p.LastName)
		}

	} else {
		t.Error("Expected page of people but received no response")
	}
}

func TestPersonCreate(t *testing.T) {
	p := &Person{
		FirstName: "Phileas",
		LastName:  "Fogg",
	}
	newP, result := c.CreatePerson(p, nil)
	if result.HasError() {
		t.Errorf("Unexpected error creating person: %s", result.Error())
		t.SkipNow()
	}

	if newP == nil {
		t.Error("Unexpected nil person response")
		t.SkipNow()
	} else {
		firstName, lastName, expectedFirstName, expectedLastName := newP.FirstName, newP.LastName, "Phileas", "Fogg"
		if firstName != expectedFirstName || lastName != expectedLastName {
			t.Errorf("Expected new person to have name %s %s but saw %s %s - check test API", expectedFirstName, expectedLastName, firstName, lastName)
		}
	}
}

func TestPersonGet(t *testing.T) {
	for id, _ := range testPeople {
		person, result := c.GetPerson(id, nil)
		if result.HasError() {
			t.Errorf("Unexpected error retrieving person %d: %s", id, result.Error())
		}

		if person == nil {
			t.Error("Unexpected nil person response")
		}
	}
}

func TestPersonUpdate(t *testing.T) {
	p := &Person{
		FirstName: "Slarti",
		LastName:  "Bartfast",
	}
	updatedPerson, result := c.UpdatePerson(testID, p, nil)
	if result.HasError() {
		t.Errorf("Unexpected error updating person: %s", result.Error())
		t.SkipNow()
	}

	if updatedPerson == nil {
		t.Error("Unexpected nil person response")
		t.SkipNow()
	} else {
		firstName, lastName, expectedFirstName, expectedLastName := updatedPerson.FirstName, updatedPerson.LastName, "Slarti", "Bartfast"
		if firstName != expectedFirstName || lastName != expectedLastName {
			t.Errorf("Expected updated person to have name %s %s but saw %s %s - check test API", expectedFirstName, expectedLastName, firstName, lastName)
		}
	}
}

func TestPersonDelete(t *testing.T) {
	result := c.DeletePerson(testID, nil)
	if result.HasError() {
		t.Errorf("Unexpected error deleting person: %s", result.Error())
		t.SkipNow()
	}

	if result.StatusCode != http.StatusNoContent {
		t.Errorf("NB delete returns a 204 on resource deletion - make sure test reflects that: saw code %d", result.StatusCode)
	}
}

func TestNearbyPeople(t *testing.T) {
	people, result := c.NearbyPeople(-78.301233, 14.129705, 10, nil)
	if result.HasError() {
		t.Errorf("Unexpected error finding nearby people: %s", result.Error())
		t.SkipNow()
	}

	if people == nil {
		t.Error("Unexpected nil people response")
		t.SkipNow()
	}
}

func TestSearchPeople(t *testing.T) {
	opts := &PeopleSearchOptions{
		FirstName: "Wally",
		LastName:  "Waldo",
	}
	people, result := c.SearchPeople(opts, nil)

	if result.HasError() {
		t.Errorf("Unexpected error searching people: %s", result.Error())
		t.SkipNow()
	}

	if people == nil {
		t.Error("Unexpected nil people response")
		t.SkipNow()
	}
}

func TestMatchPerson(t *testing.T) {
	opts := &PersonMatchOptions{
		FirstName: "MatchFirst",
		LastName:  "MatchLast",
		Email:     "match@example.com",
	}
	person, result := c.MatchPerson(opts, nil)

	if result.HasError() {
		t.Errorf("Unexpected error matching a person: %s", result.Error())
		t.SkipNow()
	}

	if person == nil {
		t.Errorf("Expected match api to return person but returned nil")
		t.SkipNow()
	}

	if person.FirstName != "MatchFirst" && person.LastName != "MatchLast" && person.Email != "match@example.com" {
		t.Errorf("Expected match call to return test person but details differ: %s", person)
		t.SkipNow()
	}
}

func TestMatchMissingPerson(t *testing.T) {
	opts := &PersonMatchOptions{
		FirstName: "MatchMissingF",
		LastName:  "MatchMissingL",
		Email:     "match-missing@example.com",
	}
	person, result := c.MatchPerson(opts, nil)

	if person != nil {
		t.Errorf("Expected person to be nil when no match found but saw: %s", person)
		t.SkipNow()
	}

	if !result.HasError() {
		t.Errorf("Expected no match of person to be as NB behaviour and return an error")
		t.SkipNow()
	}

	if result.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected error code to be 400 when no person matched but saw %d", result.StatusCode)
		t.SkipNow()
	}
}

func TestPushPerson(t *testing.T) {
	p, result := c.PushPerson(&Person{
		FirstName: "Phileas",
		LastName:  "Fogg",
	}, nil)
	if result.HasError() {
		t.Errorf("Unexpected error pushing person: %s", result.Error())
		t.SkipNow()
	}

	if p == nil {
		t.Error("Unexpected nil person response after push person - check API")
	}
}

func TestGetYourself(t *testing.T) {
	me, result := c.GetYourself(nil)
	if result.HasError() {
		t.Errorf("Unexpected error fetching yourself: %s", result.Error())
		t.SkipNow()
	}

	if me == nil {
		t.Error("Unexpected nil person response after call to 'me' endpoint - check API")
	}
}

func TestSitesGet(t *testing.T) {
	sites, result := c.GetSites(nil)
	if result.HasError() {
		t.Errorf("Unexpected error fetching sites: %s", result.Error())
		t.SkipNow()
	}

	if sites == nil {
		t.Error("Unexpected nil sites response")
	}
}

func TestAttachmentsGet(t *testing.T) {
	attachments, result := c.GetAttachments(siteSlug, siteSlug, nil)
	if result.HasError() {
		t.Errorf("Unexpected error fetching attachments: %s", result.Error())
		t.SkipNow()
	}

	if attachments == nil {
		t.Error("Unexpected nil attachments response - check API")
	}
}

func TestAttachmentsCreate(t *testing.T) {
	d, _ := NewDate(testTime)
	u := &Upload{
		FileName:    testName,
		UpdatedAt:   d,
		ContentType: "image/jpeg",
	}
	r := strings.NewReader("foo")
	a, result := c.CreateAttachment(siteSlug, siteSlug, u, r, nil)
	if result.HasError() {
		t.Errorf("Unexpected error creating attachment: %s", result.Error())
		t.SkipNow()
	}

	if a == nil {
		t.Error("Unexpected nil attachment response - check API")
	}

	if a.FileName != u.FileName {
		t.Error("Check test API - name mismatch for attachment and upload")
	}
}

func TestAttachmentGet(t *testing.T) {
	a, result := c.GetAttachment(siteSlug, siteSlug, testID, nil)
	if result.HasError() {
		t.Errorf("Unexpected error retrieving attachment: %s", result.Error())
		t.SkipNow()
	}

	if a == nil {
		t.Error("Unexpected nil attachment response - check API")
	}
}

func TestAttachmentDelete(t *testing.T) {
	result := c.DeleteAttachment(siteSlug, siteSlug, testID, nil)
	if result.HasError() {
		t.Errorf("Unexpected error deleting attachment: %s", result.Error())
		t.SkipNow()
	}
}
