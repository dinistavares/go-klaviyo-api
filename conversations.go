package klaviyo

import (
	"fmt"
)

// Conversations service
type ConversationsService service

type Conversation struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

type ConversationMessageResponse struct {
	Data *ConverstaionMessage `json:"data,omitempty"`
}

type ConverstaionMessage struct {
	Type       string      `json:"type,omitempty"`
	ID         string      `json:"id,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

type CreateConversationMessageCard struct {
	Data *CreateConversationMessage `json:"data,omitempty"`
}

type CreateConversationMessage struct {
	Type          string                               `json:"type,omitempty"`
	Attributes    *CreateConversationMessageAttributes `json:"attributes,omitempty"`
	Relationships *Relationships                       `json:"relationships,omitempty"`
}

type CreateConversationMessageAttributes struct {
	Body string `json:"body,omitempty"`
}

//  ***********************************************************************************
//  CREATE CONVERSATION MESSAGE
//  TODO: Add reference
//  ***********************************************************************************

// Sets new conversation message body
func (conversationMessage *CreateConversationMessageCard) SetConversationMessageBody(message string) {
	conversationMessage.setConversationMessageDataAttributes()

	conversationMessage.Data.Attributes.Body = message
}

// Sets new conversation message converation ID
func (conversationMessage *CreateConversationMessageCard) SetConversationMessageConversationID(id string) {
	conversationMessage.setConversationMessageDataRelationships()

	conversationMessage.Data.Relationships.Conversation = &RelationShipConversation{
		Data: &Conversation{
			Type: "conversation",
			ID: id,
		},
	}
}

// Create conversation message. Reference: (TODO: add reference)
func (service *ConversationsService) CreateConverstaionMessage(converationMessage *CreateConversationMessageCard) (*ConversationMessageResponse, *Response, error) {
	_url := fmt.Sprintf("%s/conversation-messages", ApiTypePrivate)

	// Ensure type is set to "conversation-message" if empty
	service.setCreateConversationMessageType(converationMessage)

	req, _ := service.client.NewRequest("POST", _url, nil, converationMessage)

	conversationMessage := new(ConversationMessageResponse)
	response, err := service.client.Do(req, converationMessage)

	if err != nil {
		return conversationMessage, response, err
	}

	return conversationMessage, response, nil
}

// Sets CreateCoupon.Type to 'conversation-message' if it is not set
func (service *ConversationsService) setCreateConversationMessageType(converationMessage *CreateConversationMessageCard) {
	if converationMessage != nil && converationMessage.Data != nil && converationMessage.Data.Type == "" {
		converationMessage.Data.Type = "conversation-message"
	}
}

// Ensure conversation message data and attribute pointers are created
func (conversationMessage *CreateConversationMessageCard) setConversationMessageDataAttributes() {
	if conversationMessage.Data == nil {
		conversationMessage.Data = &CreateConversationMessage{}
	}

	if conversationMessage.Data.Attributes == nil {
		conversationMessage.Data.Attributes = &CreateConversationMessageAttributes{}
	}
}

// Ensure conversation message data and relationships pointers are created
func (conversationMessage *CreateConversationMessageCard) setConversationMessageDataRelationships() {
	if conversationMessage.Data == nil {
		conversationMessage.Data = &CreateConversationMessage{}
	}

	if conversationMessage.Data.Relationships == nil {
		conversationMessage.Data.Relationships = &Relationships{}
	}
}
