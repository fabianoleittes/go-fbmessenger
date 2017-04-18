package fbmodelsend

/*
TemplateElement - Elements of Facebook Generic Template Message
*/
type TemplateElement struct {
	Title    string    `json:"title,omitempty"`
	ItemURL  string    `json:"item_url,omitempty"`
	ImageURL string    `json:"image_url,omitempty"`
	Subtitle string    `json:"subtitle,omitempty"`
	Buttons  []*Button `json:"buttons,omitempty"`
}

/*
TemplateElementShareContent - Elements of Facebook Generic Template Message for Share Content
*/
type TemplateElementShareContent struct {
	Title    string                 `json:"title,omitempty"`
	ItemURL  string                 `json:"item_url,omitempty"`
	ImageURL string                 `json:"image_url,omitempty"`
	Subtitle string                 `json:"subtitle,omitempty"`
	Buttons  []*ButtonSharedContent `json:"buttons,omitempty"`
}

/*
TemplateElementShareContentInner - Elements of Facebook Generic Template Message for Share Content
*/
type TemplateElementShareContentInner struct {
	Title         string                 `json:"title,omitempty"`
	ItemURL       string                 `json:"item_url,omitempty"`
	ImageURL      string                 `json:"image_url,omitempty"`
	Subtitle      string                 `json:"subtitle,omitempty"`
	DefaultAction DefaultAction          `json:"default_action,omitempty"`
	Buttons       []*ButtonSharedContent `json:"buttons,omitempty"`
}
