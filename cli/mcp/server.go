// Copyright Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package mcp

import (
	"context"

	"github.com/daytona/clients/cli/mcp/tools"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const intentDescription = "What problem you are trying to solve and the context of the task. Always provide this so operators and other agents can understand your actions."

const missingIntentWarning = "WARNING: no 'intent' provided. Include an 'intent' argument describing the problem you are solving and the context of your task on every tool call."

type DaytonaMCPServer struct {
	server.MCPServer
}

func NewDaytonaMCPServer() *DaytonaMCPServer {
	s := &DaytonaMCPServer{}

	s.MCPServer = *server.NewMCPServer(
		"Daytona MCP Server",
		"0.0.0-dev",
		server.WithRecovery(),
		server.WithPromptCapabilities(false),
		server.WithResourceCapabilities(false, false),
		server.WithToolCapabilities(true),
		server.WithLogging(),
	)

	s.addTools()

	return s
}

func (s *DaytonaMCPServer) Start() error {
	return server.ServeStdio(&s.MCPServer)
}

// withIntent adds the shared 'intent' parameter to a tool's input schema
func withIntent(tool mcp.Tool) mcp.Tool {
	if tool.InputSchema.Properties == nil {
		tool.InputSchema.Properties = map[string]any{}
	}
	tool.InputSchema.Properties["intent"] = map[string]any{
		"type":        "string",
		"description": intentDescription,
	}
	return tool
}

// withIntentWarning appends a warning to successful results of calls made without an 'intent'
func withIntentWarning(handler server.ToolHandlerFunc) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		result, err := handler(ctx, request)
		if result != nil && request.GetString("intent", "") == "" {
			result.Content = append(result.Content, mcp.NewTextContent(missingIntentWarning))
		}
		return result, err
	}
}

func (s *DaytonaMCPServer) addTool(tool mcp.Tool, handler server.ToolHandlerFunc) {
	s.AddTool(withIntent(tool), withIntentWarning(handler))
}

func (s *DaytonaMCPServer) addTools() {
	s.addTool(tools.GetCreateSandboxTool(), mcp.NewTypedToolHandler(tools.CreateSandbox))
	s.addTool(tools.GetDestroySandboxTool(), mcp.NewTypedToolHandler(tools.DestroySandbox))

	s.addTool(tools.GetFileUploadTool(), mcp.NewTypedToolHandler(tools.FileUpload))
	s.addTool(tools.GetFileDownloadTool(), mcp.NewTypedToolHandler(tools.FileDownload))
	s.addTool(tools.GetFileInfoTool(), mcp.NewTypedToolHandler(tools.FileInfo))
	s.addTool(tools.GetListFilesTool(), mcp.NewTypedToolHandler(tools.ListFiles))
	s.addTool(tools.GetMoveFileTool(), mcp.NewTypedToolHandler(tools.MoveFile))
	s.addTool(tools.GetDeleteFileTool(), mcp.NewTypedToolHandler(tools.DeleteFile))
	s.addTool(tools.GetCreateFolderTool(), mcp.NewTypedToolHandler(tools.CreateFolder))

	s.addTool(tools.GetExecuteCommandTool(), mcp.NewTypedToolHandler(tools.ExecuteCommand))
	s.addTool(tools.GetPreviewLinkTool(), mcp.NewTypedToolHandler(tools.PreviewLink))
	s.addTool(tools.GetGitCloneTool(), mcp.NewTypedToolHandler(tools.GitClone))

	s.addTool(tools.GetComputerUseStartTool(), mcp.NewTypedToolHandler(tools.ComputerUseStart))
	s.addTool(tools.GetComputerUseStopTool(), mcp.NewTypedToolHandler(tools.ComputerUseStop))
	s.addTool(tools.GetComputerUseStatusTool(), mcp.NewTypedToolHandler(tools.ComputerUseStatus))
	s.addTool(tools.GetComputerUseScreenshotTool(), mcp.NewTypedToolHandler(tools.ComputerUseScreenshot))
	s.addTool(tools.GetComputerUseScreenshotRegionTool(), mcp.NewTypedToolHandler(tools.ComputerUseScreenshotRegion))
	s.addTool(tools.GetComputerUseScreenshotCompressedTool(), mcp.NewTypedToolHandler(tools.ComputerUseScreenshotCompressed))
	s.addTool(tools.GetComputerUseScreenshotCompressedRegionTool(), mcp.NewTypedToolHandler(tools.ComputerUseScreenshotCompressedRegion))
	s.addTool(tools.GetComputerUseMousePositionTool(), mcp.NewTypedToolHandler(tools.ComputerUseMousePosition))
	s.addTool(tools.GetComputerUseMouseMoveTool(), mcp.NewTypedToolHandler(tools.ComputerUseMouseMove))
	s.addTool(tools.GetComputerUseMouseClickTool(), mcp.NewTypedToolHandler(tools.ComputerUseMouseClick))
	s.addTool(tools.GetComputerUseMouseDragTool(), mcp.NewTypedToolHandler(tools.ComputerUseMouseDrag))
	s.addTool(tools.GetComputerUseMouseScrollTool(), mcp.NewTypedToolHandler(tools.ComputerUseMouseScroll))
	s.addTool(tools.GetComputerUseKeyboardTypeTool(), mcp.NewTypedToolHandler(tools.ComputerUseKeyboardType))
	s.addTool(tools.GetComputerUseKeyboardPressTool(), mcp.NewTypedToolHandler(tools.ComputerUseKeyboardPress))
	s.addTool(tools.GetComputerUseKeyboardHotkeyTool(), mcp.NewTypedToolHandler(tools.ComputerUseKeyboardHotkey))
	s.addTool(tools.GetComputerUseDisplayInfoTool(), mcp.NewTypedToolHandler(tools.ComputerUseDisplayInfo))
	s.addTool(tools.GetComputerUseWindowsTool(), mcp.NewTypedToolHandler(tools.ComputerUseWindows))
	s.addTool(tools.GetComputerUseRecordingStartTool(), mcp.NewTypedToolHandler(tools.ComputerUseRecordingStart))
	s.addTool(tools.GetComputerUseRecordingStopTool(), mcp.NewTypedToolHandler(tools.ComputerUseRecordingStop))
	s.addTool(tools.GetComputerUseRecordingListTool(), mcp.NewTypedToolHandler(tools.ComputerUseRecordingList))
	s.addTool(tools.GetComputerUseRecordingGetTool(), mcp.NewTypedToolHandler(tools.ComputerUseRecordingGet))
	s.addTool(tools.GetComputerUseRecordingDeleteTool(), mcp.NewTypedToolHandler(tools.ComputerUseRecordingDelete))
	s.addTool(tools.GetComputerUseAccessibilityTreeTool(), mcp.NewTypedToolHandler(tools.ComputerUseAccessibilityTree))
	s.addTool(tools.GetComputerUseAccessibilityFindTool(), mcp.NewTypedToolHandler(tools.ComputerUseAccessibilityFind))
	s.addTool(tools.GetComputerUseAccessibilityFocusTool(), mcp.NewTypedToolHandler(tools.ComputerUseAccessibilityFocus))
	s.addTool(tools.GetComputerUseAccessibilityInvokeTool(), mcp.NewTypedToolHandler(tools.ComputerUseAccessibilityInvoke))
	s.addTool(tools.GetComputerUseAccessibilitySetValueTool(), mcp.NewTypedToolHandler(tools.ComputerUseAccessibilitySetValue))
}
