// Code generated by go-swagger; DO NOT EDIT.

package chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new chat API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for chat API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ChatGet gets chat messages
*/
func (a *Client) ChatGet(params *ChatGetParams) (*ChatGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChatGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Chat.get",
		Method:             "GET",
		PathPattern:        "/chat",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChatGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChatGetOK), nil

}

/*
ChatGetChannels gets available channels
*/
func (a *Client) ChatGetChannels(params *ChatGetChannelsParams) (*ChatGetChannelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChatGetChannelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Chat.getChannels",
		Method:             "GET",
		PathPattern:        "/chat/channels",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChatGetChannelsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChatGetChannelsOK), nil

}

/*
ChatGetConnected gets connected users

Returns an array with browser users in the first position and API users (bots) in the second position.
*/
func (a *Client) ChatGetConnected(params *ChatGetConnectedParams) (*ChatGetConnectedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChatGetConnectedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Chat.getConnected",
		Method:             "GET",
		PathPattern:        "/chat/connected",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChatGetConnectedReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChatGetConnectedOK), nil

}

/*
ChatNew sends a chat message
*/
func (a *Client) ChatNew(params *ChatNewParams, authInfo runtime.ClientAuthInfoWriter) (*ChatNewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChatNewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Chat.new",
		Method:             "POST",
		PathPattern:        "/chat",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChatNewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChatNewOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
