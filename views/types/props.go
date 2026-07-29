package types

type PageProps struct {
	PageId PageId

	Title              string
	Description        string
	Url                string
	ImageUrl           string
	AdditionalKeywords []string

	CssFilePaths []string
}
