package main

import (
	"github.com/aws-sso-oidc/mcp-server/config"
	"github.com/aws-sso-oidc/mcp-server/models"
	tools_client "github.com/aws-sso-oidc/mcp-server/tools/client"
	tools_device_authorization "github.com/aws-sso-oidc/mcp-server/tools/device_authorization"
	tools_token "github.com/aws-sso-oidc/mcp-server/tools/token"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_client.CreateRegisterclientTool(cfg),
		tools_device_authorization.CreateStartdeviceauthorizationTool(cfg),
		tools_token.CreateCreatetokenTool(cfg),
	}
}
