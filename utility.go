package klaviyo

type GenericLinks struct {
	Self    string `json:"self,omitempty"`
	First   string `json:"first,omitempty"`
	Last    string `json:"last,omitempty"`
	Prev    string `json:"prev,omitempty"`
	Next    string `json:"next,omitempty"`
	Related string `json:"related,omitempty"`
}
