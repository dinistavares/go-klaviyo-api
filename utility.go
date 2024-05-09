package klaviyo

type GenericLinks struct {
	Self    string `json:"self,omitempty"`
	First   string `json:"first,omitempty"`
	Last    string `json:"last,omitempty"`
	Prev    string `json:"prev,omitempty"`
	Next    string `json:"next,omitempty"`
	Related string `json:"related,omitempty"`
}

type GenericIncluded struct {
	Type          string         `json:"type,omitempty"`
	ID            string         `json:"id,omitempty"`
	Attributes    *interface{}   `json:"attributes,omitempty"`
	Links         *GenericLinks  `json:"links,omitempty"`
	RelationShips *Relationships `json:"relationships,omitempty"`
}