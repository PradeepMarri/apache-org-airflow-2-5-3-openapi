package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/airflow-api-stable/mcp-server/config"
	"github.com/airflow-api-stable/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_xcom_entryHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["deserialize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("deserialize=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/dags/%s/dagRuns/%s/taskInstances/%s/xcomEntries/%s%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
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
		var result models.XCom
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

func CreateGet_xcom_entryTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_dags_dag_id_dagRuns_dag_run_id_taskInstances_task_id_xcomEntries_xcom_key",
		mcp.WithDescription("Get an XCom entry"),
		mcp.WithBoolean("deserialize", mcp.Description("Whether to deserialize an XCom value when using a custom XCom backend.\n\nThe XCom API endpoint calls `orm_deserialize_value` by default since an XCom may contain value\nthat is potentially expensive to deserialize in the web server. Setting this to true overrides\nthe consideration, and calls `deserialize_value` instead.\n\nThis parameter is not meaningful when using the default XCom backend.\n\n*New in version 2.4.0*\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_xcom_entryHandler(cfg),
	}
}
