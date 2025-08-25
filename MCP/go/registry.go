package main

import (
	"github.com/aws-amplify-ui-builder/mcp-server/config"
	"github.com/aws-amplify-ui-builder/mcp-server/models"
	tools_app "github.com/aws-amplify-ui-builder/mcp-server/tools/app"
	tools_export "github.com/aws-amplify-ui-builder/mcp-server/tools/export"
	tools_tokens "github.com/aws-amplify-ui-builder/mcp-server/tools/tokens"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_app.CreateListcodegenjobsTool(cfg),
		tools_app.CreateStartcodegenjobTool(cfg),
		tools_export.CreateExportcomponentsTool(cfg),
		tools_app.CreateGetcodegenjobTool(cfg),
		tools_app.CreatePutmetadataflagTool(cfg),
		tools_export.CreateExportformsTool(cfg),
		tools_app.CreateGetmetadataTool(cfg),
		tools_tokens.CreateExchangecodefortokenTool(cfg),
		tools_tokens.CreateRefreshtokenTool(cfg),
		tools_app.CreateDeletethemeTool(cfg),
		tools_app.CreateGetthemeTool(cfg),
		tools_app.CreateUpdatethemeTool(cfg),
		tools_app.CreateListthemesTool(cfg),
		tools_app.CreateCreatethemeTool(cfg),
		tools_export.CreateExportthemesTool(cfg),
		tools_app.CreateDeletecomponentTool(cfg),
		tools_app.CreateGetcomponentTool(cfg),
		tools_app.CreateUpdatecomponentTool(cfg),
		tools_app.CreateDeleteformTool(cfg),
		tools_app.CreateGetformTool(cfg),
		tools_app.CreateUpdateformTool(cfg),
		tools_app.CreateListformsTool(cfg),
		tools_app.CreateCreateformTool(cfg),
		tools_app.CreateListcomponentsTool(cfg),
		tools_app.CreateCreatecomponentTool(cfg),
	}
}
