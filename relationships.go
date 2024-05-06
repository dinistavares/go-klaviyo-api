package klaviyo

type Relationships struct {
	Lists        *RelationshipLists        `json:"lists,omitempty"`
	Segments     *RelationshipSegments     `json:"segments,omitempty"`
	Profile      *RelationshipProfile      `json:"profile,omitempty"`
	Profiles     *RelationshipProfiles     `json:"profiles,omitempty"`
	Tags         *RelationshipTags         `json:"tags,omitempty"`
	Metric       *RelarionshipMetric       `json:"metric,omitempty"`
	Attributions *RelationShipAttributions `json:"attributions,omitempty"`
}

type GenericRelationshipData struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

type RelationshipLists struct {
	Links *GenericLinks `json:"links,omitempty"`
}

type RelationshipSegments struct {
	Links *GenericLinks `json:"links,omitempty"`
}

type RelationshipProfile struct {
	Profile *Profile      `json:"data,omitempty"`
	Links   *GenericLinks `json:"links,omitempty"`
}

type RelationshipProfiles struct {
	Links *GenericLinks `json:"links,omitempty"`
}

type RelationshipTags struct {
	Links *GenericLinks `json:"links,omitempty"`
}

type RelarionshipMetric struct {
	Data  *Event        `json:"data,omitempty"`
	Links *GenericLinks `json:"links,omitempty"`
}

type RelationShipAttributions struct {
	Data  *[]GenericRelationshipData `json:"data,omitempty"`
	Links *GenericLinks              `json:"links,omitempty"`
}
