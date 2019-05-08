package model

// Form is creatd by http post x-www-form-urlencoded
type Form struct {
	FromEmail  string `json:"from_email"`
	FromName   string `json:"from_name"`
	EnvFrom    string `json:"env_from"`
	EnvToList  string `json:"env_to_list"`
	ToList     string `json:"to_list"`
	HeaderList string `json:"header_list"`
	Subject    string `json:"subject"`
	BodyText   string `json:"body_text"`
	BodyHTML   string `json:"body_html"`
}
