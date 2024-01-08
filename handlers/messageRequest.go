package handlers

import "fmt"

type adsContextData struct {
	AdTitle  string `json:"ad_title"`
	PhotoURL string `json:"photo_url"`
	VideoURL string `json:"video_url"`
}

type product struct {
	ID string `json:"id"`
}

type referral struct {
	Product        product        `json:"product"`
	Ref            string         `json:"ref"`
	AdID           string         `json:"ad_id"`
	Source         string         `json:"source"`
	Type           string         `json:"type"`
	AdsContextData adsContextData `json:"ads_context_data"`
}

type payload struct {
	URL       string `json:"url"`
	Title     string `json:"title"`
	StickerID int64  `json:"sticker_id"`
}

type attachment struct {
	Type    string  `json:"type"`
	Payload payload `json:"payload"`
}

type story struct {
	URL string `json:"url"`
	ID  string `json:"id"`
}

type replyTo struct {
	MID   string `json:"mid"`
	Story story  `json:"story"`
}

type quickReply struct {
	Payload string `json:"payload"`
}

type postback struct {
	Title   string `json:"title"`
	Payload string `json:"payload"`
	MID     string `json:"mid"`
}

type message struct {
	MID           string       `json:"mid"`
	Text          string       `json:"text"`
	Postback      postback     `json:"postback"`
	QuickReply    quickReply   `json:"quick_reply"`
	ReplyTo       replyTo      `json:"reply_to"`
	Attachments   []attachment `json:"attachments"`
	Referral      referral     `json:"referral"`
	IsDeleted     bool         `json:"is_deleted"`
	IsEcho        bool         `json:"is_echo"`
	IsUnsupported bool         `json:"is_unsupported"`
}

type read struct {
	MID string `json:"mid"`
}

type messagingSender struct {
	ID      string `json:"id"`
	UserRef string `json:"user_ref"`
}

type messagingRecipient struct {
	ID string `json:"id"`
}

type messaging struct {
	Sender    messagingSender    `json:"sender"`
	Recipient messagingRecipient `json:"recipient"`
	Timestamp int64              `json:"timestamp"`
	Postback  postback           `json:"postback"`
	Message   message            `json:"message"`
	Read      read               `json:"read"`
}

const (
	Page      = "page"
	Instagram = "instagram"
)

type eventObject struct {
	Page      string `json:"page"`
	Instagram string `json:"instagram"`
}

type entry struct {
	ID        string      `json:"id"`
	Time      int64       `json:"time"`
	Messaging []messaging `json:"messaging"`
}

type MessageReceiveRequest struct {
	Object eventObject `json:"object"`
	Entry  []entry     `json:"entry"`
}

func (r *MessageReceiveRequest) Validate() error {
	if r.Object == (eventObject{}) {
		return fmt.Errorf("object field is empty or not initialized")
	}
	if len(r.Entry) == 0 {
		return fmt.Errorf("entry field is empty")
	}
	return nil
}
