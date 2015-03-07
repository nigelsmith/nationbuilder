package nationbuilder

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Attachment represents a file 'attached' to a particular page
type Attachment struct {
	ID          int    `json:"id,omitempty"`
	FileName    string `json:"filename,omitempty"`
	UpdatedAt   *Date  `json:"updated_at,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	URL         string `json:"url,omitempty"`
}

func (a *Attachment) String() string {
	return fmt.Sprintf("Attachment: %s", a.FileName)
}

// Paginated collection of attachments
type Attachments struct {
	Results []*Attachment
	Pagination
}

type attachmentWrap struct {
	Attachment *Attachment `json:"attachment,omitempty"`
}

// An upload with a content payload of a base64 encoded string representation
// of a file.
type Upload struct {
	FileName    string `json:"filename,omitempty"`
	UpdatedAt   *Date  `json:"updated_at"`
	ContentType string `json:"content_type,omitempty"`
	Content     string `json:"content,omitempty"`
}

type uploadWrap struct {
	Attachment *Upload `json:"attachment,omitempty"`
}

// Retrieve a page of attachments for the given site and page
func (n *Client) GetAttachments(siteSlug string, pageSlug string, options *Options) (attachments *Attachments, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/%s/attachments", siteSlug, pageSlug)
	req := n.getRequest("GET", u, options)
	result = n.retrieve(req, &attachments)

	return
}

// Create an attachment for the given site and page by providing an upload object.
// If upload's content field is empty then file must be non-nil.
// I.e. you may base64 encode your own content and provide it that way or provide a reader and the resulting
// bytes will be base64 encoded for you.  If the upload's content is not empty and an io.Reader is provided
// then the upload.Content will take precedence.
func (n *Client) CreateAttachment(siteSlug string, pageSlug string, upload *Upload, file io.Reader, options *Options) (attachment *Attachment, result *Result) {
	if upload == nil {
		return nil, &Result{
			Err: errors.New("Please supply an attachment object to upload"),
		}
	}

	if upload.Content == "" && file == nil {
		return nil, &Result{
			Err: errors.New("No file provided to upload"),
		}
	}

	if upload.Content == "" {
		b, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, &Result{
				Err: err,
			}
		}
		upload.Content = base64.StdEncoding.EncodeToString(b)
	}

	u := fmt.Sprintf("/sites/%s/pages/%s/attachments", siteSlug, pageSlug)
	r := n.getRequest("POST", u, options)
	aw := &attachmentWrap{}
	result = n.create(&uploadWrap{upload}, r, aw, http.StatusOK)
	attachment = aw.Attachment

	return
}

// Retrieve a single attachment for the given site and page with the specified attachment id
func (n *Client) GetAttachment(siteSlug string, pageSlug string, id int, options *Options) (attachment *Attachment, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/%s/attachments/%d", siteSlug, pageSlug, id)
	r := n.getRequest("GET", u, options)
	a := &attachmentWrap{}
	result = n.retrieve(r, a)
	attachment = a.Attachment

	return
}

// Delete a page attachment
func (n *Client) DeleteAttachment(siteSlug string, pageSlug string, id int, options *Options) (result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/%s/attachments/%d", siteSlug, pageSlug, id)
	r := n.getRequest("DELETE", u, options)
	result = n.delete(r)

	return
}
