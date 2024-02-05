package dto

type Email struct {
	SMTPServer     string
	SMTPUsername   string
	SMTPPassword   string
	SenderEmail    string
	RecipientEmail string
	SubjectEmail   string
	BodyEmail      string
	SMTPPort       int
}
