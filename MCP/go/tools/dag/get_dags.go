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

func Get_dagsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["order_by"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("order_by=%v", val))
		}
		if val, ok := args["tags"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tags=%v", val))
		}
		if val, ok := args["only_active"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("only_active=%v", val))
		}
		if val, ok := args["dag_id_pattern"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dag_id_pattern=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/dags%s", cfg.BaseURL, queryString)
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
		var result models.DAGCollection
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

func CreateGet_dagsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_dags",
		mcp.WithDescription("List DAGs"),
		mcp.WithNumber("limit", mcp.Description("The numbers of items to return.")),
		mcp.WithNumber("offset", mcp.Description("The number of items to skip before starting to collect the result set.")),
		mcp.WithString("order_by", mcp.Description("The name of the field to order the results by.\nPrefix a field name with `-` to reverse the sort order.\n\n*New in version 2.1.0*\n")),
		mcp.WithArray("tags", mcp.Description("List of tags to filter results.\n\n*New in version 2.2.0*\n")),
		mcp.WithBoolean("only_active", mcp.Description("Only filter active DAGs.\n\n*New in version 2.1.1*\n")),
		mcp.WithString("dag_id_pattern", mcp.Description("If set, only return DAGs with dag_ids matching this pattern.\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_dagsHandler(cfg),
	}
}
