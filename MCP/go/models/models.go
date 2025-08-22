package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// StartDeviceAuthorizationResponse represents the StartDeviceAuthorizationResponse schema from the OpenAPI specification
type StartDeviceAuthorizationResponse struct {
	Expiresin interface{} `json:"expiresIn,omitempty"`
	Interval interface{} `json:"interval,omitempty"`
	Usercode interface{} `json:"userCode,omitempty"`
	Verificationuri interface{} `json:"verificationUri,omitempty"`
	Verificationuricomplete interface{} `json:"verificationUriComplete,omitempty"`
	Devicecode interface{} `json:"deviceCode,omitempty"`
}

// RegisterClientResponse represents the RegisterClientResponse schema from the OpenAPI specification
type RegisterClientResponse struct {
	Clientid interface{} `json:"clientId,omitempty"`
	Clientidissuedat interface{} `json:"clientIdIssuedAt,omitempty"`
	Clientsecret interface{} `json:"clientSecret,omitempty"`
	Clientsecretexpiresat interface{} `json:"clientSecretExpiresAt,omitempty"`
	Tokenendpoint interface{} `json:"tokenEndpoint,omitempty"`
	Authorizationendpoint interface{} `json:"authorizationEndpoint,omitempty"`
}

// RegisterClientRequest represents the RegisterClientRequest schema from the OpenAPI specification
type RegisterClientRequest struct {
	Clientname interface{} `json:"clientName"`
	Clienttype interface{} `json:"clientType"`
	Scopes interface{} `json:"scopes,omitempty"`
}

// CreateTokenResponse represents the CreateTokenResponse schema from the OpenAPI specification
type CreateTokenResponse struct {
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Expiresin interface{} `json:"expiresIn,omitempty"`
	Idtoken interface{} `json:"idToken,omitempty"`
	Refreshtoken interface{} `json:"refreshToken,omitempty"`
	Tokentype interface{} `json:"tokenType,omitempty"`
}

// StartDeviceAuthorizationRequest represents the StartDeviceAuthorizationRequest schema from the OpenAPI specification
type StartDeviceAuthorizationRequest struct {
	Clientid interface{} `json:"clientId"`
	Clientsecret interface{} `json:"clientSecret"`
	Starturl interface{} `json:"startUrl"`
}

// CreateTokenRequest represents the CreateTokenRequest schema from the OpenAPI specification
type CreateTokenRequest struct {
	Clientid interface{} `json:"clientId"`
	Clientsecret interface{} `json:"clientSecret"`
	Code interface{} `json:"code,omitempty"`
	Devicecode interface{} `json:"deviceCode,omitempty"`
	Granttype interface{} `json:"grantType"`
	Redirecturi interface{} `json:"redirectUri,omitempty"`
	Refreshtoken interface{} `json:"refreshToken,omitempty"`
	Scope interface{} `json:"scope,omitempty"`
}
