package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/thereisnotime/sdrangel-mcp/internal/sdrangel"
)

func registerFeatureTools(srv *mcp.Server, c *sdrangel.Client) {
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "list_features",
		Description: "List all feature instances currently active in the SDRAngel feature set (e.g. Map, ADS-B Demodulator, RDS, VOR Localizer).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.FeatureSetInfo, error) {
		result, err := c.GetFeatureSet(ctx)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "add_feature",
		Description: "Add a new feature instance by featureType (e.g. Map, VORLocalizer, Demodulator, AISDemod). Use list_feature_plugins to see available types.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.FeatureAdd) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.AddFeature(ctx, args)
		return nil, result, err
	})

	type FeatureIndexArgs struct {
		FeatureIndex int `json:"featureIndex"`
	}

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "delete_feature",
		Description: "Delete a feature instance by its index.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args FeatureIndexArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.DeleteFeature(ctx, args.FeatureIndex)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_feature_settings",
		Description: "Get the current settings for a feature instance. Returns featureType, settingsKey (the plugin-specific wire key, e.g. MapSettings) and settings (that plugin's settings JSON).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args FeatureIndexArgs) (*mcp.CallToolResult, sdrangel.FeatureSettings, error) {
		result, err := c.GetFeatureSettings(ctx, args.FeatureIndex)
		return nil, result, err
	})

	type SetFeatureSettingsArgs struct {
		FeatureIndex int                      `json:"featureIndex"`
		Settings     sdrangel.FeatureSettings `json:"settings"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "set_feature_settings",
		Description: "Replace all settings for a feature instance. Call get_feature_settings first to learn settings.settingsKey (SDRAngel wraps the plugin's settings object under a plugin-specific wire key, e.g. MapSettings — not a generic \"settings\" key) and echo it back along with settings.settings (the plugin-specific JSON).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args SetFeatureSettingsArgs) (*mcp.CallToolResult, sdrangel.FeatureSettings, error) {
		result, err := c.SetFeatureSettings(ctx, args.FeatureIndex, args.Settings)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "patch_feature_settings",
		Description: "Partially update settings for a feature instance. Only provided fields are changed. Call get_feature_settings first to learn settings.settingsKey and echo it back along with the changed fields in settings.settings.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args SetFeatureSettingsArgs) (*mcp.CallToolResult, sdrangel.FeatureSettings, error) {
		result, err := c.PatchFeatureSettings(ctx, args.FeatureIndex, args.Settings)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "start_feature",
		Description: "Start a feature instance (begin its processing loop).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args FeatureIndexArgs) (*mcp.CallToolResult, sdrangel.FeatureState, error) {
		result, err := c.StartFeature(ctx, args.FeatureIndex)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "stop_feature",
		Description: "Stop a feature instance.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args FeatureIndexArgs) (*mcp.CallToolResult, sdrangel.FeatureState, error) {
		result, err := c.StopFeature(ctx, args.FeatureIndex)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_feature_run_status",
		Description: "Get the current run state (idle, running, error) of a feature instance.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args FeatureIndexArgs) (*mcp.CallToolResult, sdrangel.FeatureState, error) {
		result, err := c.GetFeatureRunStatus(ctx, args.FeatureIndex)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_feature_report",
		Description: "Get the runtime report for a feature instance (feature-specific runtime data). Returns reportKey (the plugin-specific wire key, e.g. MapReport) and report (that plugin's report JSON).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args FeatureIndexArgs) (*mcp.CallToolResult, sdrangel.FeatureReport, error) {
		result, err := c.GetFeatureReport(ctx, args.FeatureIndex)
		return nil, result, err
	})

	type ExecuteFeatureActionsArgs struct {
		FeatureIndex int                     `json:"featureIndex"`
		Actions      sdrangel.FeatureActions `json:"actions"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "execute_feature_actions",
		Description: "Execute actions on a feature instance (plugin-specific operations). actions.actionsKey is the plugin-specific wire key (e.g. MapActions) wrapping actions.actions (the plugin-specific JSON) — not a generic \"actions\" key.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args ExecuteFeatureActionsArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.ExecuteFeatureActions(ctx, args.FeatureIndex, args.Actions)
		return nil, result, err
	})
}
