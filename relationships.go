package klaviyo

type Relationships struct {
	Attributions         *RelationShipAttributions         `json:"attributions,omitempty"`
	AttributedEvent      *RelationShipAttributedEvent      `json:"attributed-event,omitempty"`
	Campaign             *RelationShipCampaign             `json:"campaign,omitempty"`
	CampaignMessage      *RelationShipCampaignMessage      `json:"campaign-message,omitempty"`
	Conversation         *RelationShipConversation         `json:"conversation,omitempty"`
	Event                *RelationShipEvent                `json:"event,omitempty"`
	Flow                 *RelationShipFlow                 `json:"flow,omitempty"`
	FlowMessage          *RelationShipFlowMessage          `json:"flow-message,omitempty"`
	FlowMessageVariation *RelationShipFlowMessageVariation `json:"low-message-variation,omitempty"`
	Lists                *RelationshipLists                `json:"lists,omitempty"`
	Metric               *RelationshipMetric               `json:"metric,omitempty"`
	Profile              *RelationshipProfile              `json:"profile,omitempty"`
	Profiles             *RelationshipProfiles             `json:"profiles,omitempty"`
	Segments             *RelationshipSegments             `json:"segments,omitempty"`
	Tags                 *RelationshipTags                 `json:"tags,omitempty"`
}

type GenericRelationshipData struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

type RelationShipAttributions struct {
	Data  *[]GenericRelationshipData `json:"data,omitempty"`
	Links *GenericLinks              `json:"links,omitempty"`
}

type RelationShipAttributedEvent struct {
	Data  *GenericRelationshipData `json:"data,omitempty"`
	Links *GenericLinks            `json:"links,omitempty"`
}

type RelationShipCampaign struct {
	Data  *GenericRelationshipData `json:"data,omitempty"`
	Links *GenericLinks            `json:"links,omitempty"`
}

type RelationShipCampaignMessage struct {
	Data  *GenericRelationshipData `json:"data,omitempty"`
	Links *GenericLinks            `json:"links,omitempty"`
}

type RelationShipConversation struct {
	Links *GenericLinks `json:"links,omitempty"`
}

type RelationShipEvent struct {
	Data  *Event        `json:"data,omitempty"`
	Links *GenericLinks `json:"links,omitempty"`
}

type RelationShipFlow struct {
	Data  *Event        `json:"data,omitempty"`
	Links *GenericLinks `json:"links,omitempty"`
}

type RelationShipFlowMessage struct {
	Data  *Event        `json:"data,omitempty"`
	Links *GenericLinks `json:"links,omitempty"`
}

type RelationShipFlowMessageVariation struct {
	Data  *Event        `json:"data,omitempty"`
	Links *GenericLinks `json:"links,omitempty"`
}

type RelationshipLists struct {
	Links *GenericLinks `json:"links,omitempty"`
}

type RelationshipMetric struct {
	Data  *Event        `json:"data,omitempty"`
	Links *GenericLinks `json:"links,omitempty"`
}

type RelationshipProfile struct {
	Profile *Profile      `json:"data,omitempty"`
	Links   *GenericLinks `json:"links,omitempty"`
}

type RelationshipProfiles struct {
	Links *GenericLinks `json:"links,omitempty"`
}

type RelationshipSegments struct {
	Links *GenericLinks `json:"links,omitempty"`
}

type RelationshipTags struct {
	Links *GenericLinks `json:"links,omitempty"`
}
