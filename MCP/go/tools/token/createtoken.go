package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/aws-sso-oidc/mcp-server/config"
	"github.com/aws-sso-oidc/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreatetokenHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/token", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.CreateTokenResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCreatetokenTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_token",
		mcp.WithDescription("Creates and returns an access token for the authorized client. The access token issued will be used to fetch short-term credentials for the assigned roles in the AWS account."),
		mcp.WithString("deviceCode", mcp.Description("Input parameter: Used only when calling this API for the device code grant type. This short-term code is used to identify this authentication attempt. This should come from an in-memory reference to the result of the <a>StartDeviceAuthorization</a> API.")),
		mcp.WithString("grantType", mcp.Required(), mcp.Description("Input parameter: <p>Supports grant types for the authorization code, refresh token, and device code request. For device code requests, specify the following value:</p> <p> <code>urn:ietf:params:oauth:grant-type:<i>device_code</i> </code> </p> <p>For information about how to obtain the device code, see the <a>StartDeviceAuthorization</a> topic.</p>")),
		mcp.WithString("redirectUri", mcp.Description("Input parameter: The location of the application that will receive the authorization code. Users authorize the service to send the request to this location.")),
		mcp.WithString("refreshToken", mcp.Description("Input parameter: <p>Currently, <code>refreshToken</code> is not yet implemented and is not supported. For more information about the features and limitations of the current IAM Identity Center OIDC implementation, see <i>Considerations for Using this Guide</i> in the <a href=\"https://docs.aws.amazon.com/singlesignon/latest/OIDCAPIReference/Welcome.html\">IAM Identity Center OIDC API Reference</a>.</p> <p>The token used to obtain an access token in the event that the access token is invalid or expired.</p>")),
		mcp.WithArray("scope", mcp.Description("Input parameter: The list of scopes that is defined by the client. Upon authorization, this list is used to restrict permissions when granting an access token.")),
		mcp.WithString("clientId", mcp.Required(), mcp.Description("Input parameter: The unique identifier string for each client. This value should come from the persisted result of the <a>RegisterClient</a> API.")),
		mcp.WithString("clientSecret", mcp.Required(), mcp.Description("Input parameter: A secret string generated for the client. This value should come from the persisted result of the <a>RegisterClient</a> API.")),
		mcp.WithString("code", mcp.Description("Input parameter: The authorization code received from the authorization service. This parameter is required to perform an authorization grant request to get access to a token.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatetokenHandler(cfg),
	}
}
