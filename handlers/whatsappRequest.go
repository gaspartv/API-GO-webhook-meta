package handlers

import "fmt"

type profile struct {
	Name string `json:"name"`
}

type metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type valueContact struct {
	Profile profile `json:"profile"`
	WaID    string  `json:"wa_id"`
}

type errorDataReceive struct {
	Details string `json:"details"`
}

type unsupportedReceive struct {
	Code      int64            `json:"code"`
	Title     string           `json:"title"`
	Message   string           `json:"message"`
	ErrorData errorDataReceive `json:"error_data"`
}

type stickerReceive struct {
	ID       string `json:"id"`
	MimeType string `json:"mime_type"`
	SHA256   string `json:"sha256"`
	Animated bool   `json:"animated"`
}

type urlReceive struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

type phoneReceive struct {
	Phone string `json:"phone"`
	Type  string `json:"type"`
	WaID  string `json:"wa_id"`
}

type orgReceive struct {
	Company    string `json:"company"`
	Department string `json:"department"`
	Title      string `json:"title"`
}

type nameReceive struct {
	FormattedName string `json:"formatted_name"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	MiddleName    string `json:"middle_name"`
	Suffix        string `json:"suffix"`
	Prefix        string `json:"prefix"`
}

type emailReceive struct {
	Email string `json:"email"`
	Type  string `json:"type"`
}

type addressReceive struct {
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	Type        string `json:"type"`
}

type contactReceive struct {
	Addresses addressReceive `json:"addresses"`
	Birthday  string         `json:"birthday"`
	Emails    []emailReceive `json:"emails"`
	Name      nameReceive    `json:"name"`
	Org       orgReceive     `json:"org"`
	Phones    []phoneReceive `json:"phones"`
	URLs      []urlReceive   `json:"urls"`
}

type textReceive struct {
	Body string `json:"body"`
}

type replyListId string

const (
	text     replyListId = "text"
	image    replyListId = "image"
	video    replyListId = "video"
	audio    replyListId = "audio"
	document replyListId = "document"
)

type listReplyReceive struct {
	ID          replyListId `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
}

type replyButtonId string

type buttonReplyReceive struct {
	ID    replyButtonId `json:"id"`
	Title string        `json:"title"`
}

const (
	Yes      replyButtonId = "yes"
	No       replyButtonId = "no"
	Cancel   replyButtonId = "cancel"
	Ok       replyButtonId = "ok"
	Next     replyButtonId = "next"
	Previous replyButtonId = "previous"
	Done     replyButtonId = "done"
)

type typeReceive string

const (
	Button      typeReceive = "button"
	Catalog     typeReceive = "catalog_message"
	List        typeReceive = "list"
	Product     typeReceive = "product"
	ProductList typeReceive = "product_list"
	ButtonReply typeReceive = "button_reply"
	ListReply   typeReceive = "list_reply"
)

type interactiveReceive struct {
	Type        typeReceive        `json:"type"`
	ButtonReply buttonReplyReceive `json:"button_reply"`
	ListReply   listReplyReceive   `json:"list_reply"`
}

type reactionReceive struct {
	MessageID string `json:"message_id"`
	Emoji     string `json:"emoji"`
}

type locationReceive struct {
	Longitude int64  `json:"longitude"`
	Latitude  int64  `json:"latitude"`
	Name      string `json:"name"`
	Address   string `json:"address"`
}

type media struct {
	ID       string `json:"id"`
	Link     string `json:"link"`
	Filename string `json:"filename"`
	Provider string `json:"provider"`
}

type messageTypeEnum string

const (
	File        messageTypeEnum = "file"
	ListButton  messageTypeEnum = "list-button"
	Text        messageTypeEnum = "text"
	Audio       messageTypeEnum = "audio"
	Document    messageTypeEnum = "document"
	Image       messageTypeEnum = "image"
	Interactive messageTypeEnum = "interactive"
	Order       messageTypeEnum = "order"
	Sticker     messageTypeEnum = "sticker"
	System      messageTypeEnum = "system"
	Video       messageTypeEnum = "video"
	Location    messageTypeEnum = "location"
	Contacts    messageTypeEnum = "contacts"
	Reaction    messageTypeEnum = "reaction"
)

type contextReceive struct {
}

type messageWhatsapp struct {
	From        string               `json:"from"`
	ID          string               `json:"id"`
	Timestamp   string               `json:"timestamp"`
	Type        messageTypeEnum      `json:"type"`
	To          string               `json:"to"`
	TTL         string               `json:"ttl"`
	Image       media                `json:"image"`
	Video       media                `json:"video"`
	Audio       media                `json:"audio"`
	Document    media                `json:"document"`
	Location    locationReceive      `json:"location"`
	Reaction    reactionReceive      `json:"reaction"`
	Interactive interactiveReceive   `json:"interactive"`
	Text        textReceive          `json:"text"`
	Contacts    []contactReceive     `json:"contacts"`
	Sticker     stickerReceive       `json:"sticker"`
	Errors      []unsupportedReceive `json:"errors"`
	Context     contextReceive       `json:"context"`
}

type statusReceive string

const (
	Delivered statusReceive = "delivered"
	Sent      statusReceive = "sent"
	Read      statusReceive = "read"
)

type originReceive struct {
	Type string `json:"type"`
}

type conversationReceive struct {
	ID                  string        `json:"id"`
	ExpirationTimestamp string        `json:"expiration_timestamp"`
	Origin              originReceive `json:"origin"`
}

type pricingReceive struct {
	Billable     bool   `json:"billable"`
	PricingModel string `json:"pricing_model"`
	Category     string `json:"category"`
}

type status struct {
	ID           string              `json:"id"`
	Status       statusReceive       `json:"status"`
	Timestamp    string              `json:"timestamp"`
	RecipientID  string              `json:"recipient_id"`
	Conversation conversationReceive `json:"conversation"`
	Pricing      pricingReceive      `json:"pricing"`
}

type value struct {
	MessagingProduct string            `json:"messaging_product"`
	Metadata         metadata          `json:"metadata"`
	Contacts         []valueContact    `json:"contacts"`
	Messages         []messageWhatsapp `json:"messages"`
	Statuses         []status          `json:"statuses"`
}

type changes struct {
	Field string `json:"field"`
	Value value  `json:"value"`
}

type entryWhatsapp struct {
	ID      string    `json:"id"`
	Changes []changes `json:"changes"`
}

type WhatsappReceiveRequest struct {
	Object string          `json:"object"`
	Entry  []entryWhatsapp `json:"entry"`
}

func (r *WhatsappReceiveRequest) Validate() error {
	if r.Object == "" {
		return fmt.Errorf("object field is empty or not initialized")
	}
	if len(r.Entry) == 0 {
		return fmt.Errorf("entry field is empty")
	}
	return nil
}
