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

func Get_dag_runsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["execution_date_gte"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("execution_date_gte=%v", val))
		}
		if val, ok := args["execution_date_lte"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("execution_date_lte=%v", val))
		}
		if val, ok := args["start_date_gte"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start_date_gte=%v", val))
		}
		if val, ok := args["start_date_lte"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start_date_lte=%v", val))
		}
		if val, ok := args["end_date_gte"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("end_date_gte=%v", val))
		}
		if val, ok := args["end_date_lte"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("end_date_lte=%v", val))
		}
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["order_by"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("order_by=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/dags/%s/dagRuns%s", cfg.BaseURL, queryString)
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
		var result models.DAGRunCollection
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

func CreateGet_dag_runsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_dags_dag_id_dagRuns",
		mcp.WithDescription("List DAG runs"),
		mcp.WithNumber("limit", mcp.Description("The numbers of items to return.")),
		mcp.WithNumber("offset", mcp.Description("The number of items to skip before starting to collect the result set.")),
		mcp.WithString("execution_date_gte", mcp.Description("Returns objects greater or equal to the specified date.\n\nThis can be combined with execution_date_lte parameter to receive only the selected period.\n")),
		mcp.WithString("execution_date_lte", mcp.Description("Returns objects less than or equal to the specified date.\n\nThis can be combined with execution_date_gte parameter to receive only the selected period.\n")),
		mcp.WithString("start_date_gte", mcp.Description("Returns objects greater or equal the specified date.\n\nThis can be combined with start_date_lte parameter to receive only the selected period.\n")),
		mcp.WithString("start_date_lte", mcp.Description("Returns objects less or equal the specified date.\n\nThis can be combined with start_date_gte parameter to receive only the selected period.\n")),
		mcp.WithString("end_date_gte", mcp.Description("Returns objects greater or equal the specified date.\n\nThis can be combined with start_date_lte parameter to receive only the selected period.\n")),
		mcp.WithString("end_date_lte", mcp.Description("Returns objects less than or equal to the specified date.\n\nThis can be combined with start_date_gte parameter to receive only the selected period.\n")),
		mcp.WithArray("state", mcp.Description("The value can be repeated to retrieve multiple matching values (OR condition).")),
		mcp.WithString("order_by", mcp.Description("The name of the field to order the results by.\nPrefix a field name with `-` to reverse the sort order.\n\n*New in version 2.1.0*\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_dag_runsHandler(cfg),
	}
}
