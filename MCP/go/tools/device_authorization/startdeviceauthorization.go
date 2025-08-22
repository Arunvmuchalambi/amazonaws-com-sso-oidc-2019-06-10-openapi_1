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

func StartdeviceauthorizationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/device_authorization", cfg.BaseURL)
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
		var result models.StartDeviceAuthorizationResponse
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

func CreateStartdeviceauthorizationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_device_authorization",
		mcp.WithDescription("Initiates device authorization by requesting a pair of verification codes from the authorization service."),
		mcp.WithString("clientId", mcp.Required(), mcp.Description("Input parameter: The unique identifier string for the client that is registered with IAM Identity Center. This value should come from the persisted result of the <a>RegisterClient</a> API operation.")),
		mcp.WithString("clientSecret", mcp.Required(), mcp.Description("Input parameter: A secret string that is generated for the client. This value should come from the persisted result of the <a>RegisterClient</a> API operation.")),
		mcp.WithString("startUrl", mcp.Required(), mcp.Description("Input parameter: The URL for the AWS access portal. For more information, see <a href=\"https://docs.aws.amazon.com/singlesignon/latest/userguide/using-the-portal.html\">Using the AWS access portal</a> in the <i>IAM Identity Center User Guide</i>.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    StartdeviceauthorizationHandler(cfg),
	}
}
