package nationbuilder

import "net/http"

type Donation struct {
	Amount                string   `json:"amount,omitempty"`
	AmountInCents         int      `json:"amount_in_cents,omitempty"`
	AuthorID              int      `json:"author_id,omitempty"`
	BillingAddress        *Address `json:"billing_address,omitempty"`
	CancelledAt           *Date    `json:"canceled_at,omitempty"`
	ChequeNumber          int      `json:"check_number,omitempty"`
	CorporateContribution bool     `json:"corporate_contribution,omitempty"`
	CreatedAt             *Date    `json:"created_at,omitempty"`
	DonorID               int      `json:"donor_id,omitempty"`
	Donor                 *Person  `json:"donor,omitempty"`
	Email                 string   `json:"email,omitempty"`
	Employer              string   `json:"employer,omitempty"`
	FailedAt              *Date    `json:"failed_at,omitempty"`
	FirstName             string   `json:"first_name,omitempty"`
	ID                    int      `json:"id,omitempty"`
	ImportID              int      `json:"import_id,omitempty"`
	IsPrivate             bool     `json:"is_private,omitempty"`
	LastName              string   `json:"last_name,omitempty"`
	MailingSlug           string   `json:"mailing_slug,omitempty"`
	MerchantAccountID     string   `json:"merchant_account_id,omitempty"`
	NGPID                 int      `json:"ngp_id,omitempty"`
	Note                  string   `json:"note,omitempty"`
	Occupation            string   `json:"occupation,omitempty"`
	PageSlug              string   `json:"page_slug,omitempty"`
	PaymentTypeName       string   `json:"payment_type_name,omitempty"`
	PaymentTypeNGPCode    string   `json:"payment_type_ngp_code,omitempty"`
	PledgeID              int      `json:"pledge_id,omitempty"`
	RecruiterNameOrEmail  string   `json:"recruiter_name_or_email,omitempty"`
	RecurringDonationID   int      `json:"recurring_donation_id,omitempty"`
	SucceededAt           *Date    `json:"succeeded_at,omitempty"`
	TrackingCodeSlug      string   `json:"tracking_code_slug,omitempty"`
	UpdatedAt             *Date    `json:"updated_at,omitempty"`
	WorkAddress           *Address `json:"work_address,omitempty"`
}

type donationWrap struct {
	Donation *Donation `json:"donation"`
}


func (n *Client) CreateDonation(donation *Donation, options *Options) (*Donation, *Result) {
	req := n.getRequest("POST", "/donations", options)
	dw := &donationWrap{}
	result := n.create(&donationWrap{donation}, req, dw, http.StatusOK)

	return dw.Donation, result
}