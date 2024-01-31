package dto

type Email struct {
	SMTPServer     string
	SMTPPort       int
	SMTPUsername   string
	SMTPPassword   string
	SenderEmail    string
	RecipientEmail string
	SubjectEmail   string
	BodyEmail      string
}
