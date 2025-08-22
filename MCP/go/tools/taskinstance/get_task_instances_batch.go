package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/airflow-api-stable/mcp-server/config"
	"github.com/airflow-api-stable/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_task_instances_batchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.ListTaskInstanceForm
		
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
		url := fmt.Sprintf("%s/dags/~/dagRuns/~/taskInstances/list%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.TaskInstanceCollection
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

func CreateGet_task_instances_batchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_dags_~_dagRuns_~_taskInstances_list",
		mcp.WithDescription("List task instances (batch)"),
		mcp.WithString("duration_gte", mcp.Description("Input parameter: Returns objects greater than or equal to the specified values.\n\nThis can be combined with duration_lte parameter to receive only the selected period.\n")),
		mcp.WithString("end_date_lte", mcp.Description("Input parameter: Returns objects less than or equal to the specified date.\n\nThis can be combined with start_date_gte parameter to receive only the selected period.\n")),
		mcp.WithString("execution_date_lte", mcp.Description("Input parameter: Returns objects less than or equal to the specified date.\n\nThis can be combined with execution_date_gte parameter to receive only the selected period.\n")),
		mcp.WithString("execution_date_gte", mcp.Description("Input parameter: Returns objects greater or equal to the specified date.\n\nThis can be combined with execution_date_lte parameter to receive only the selected period.\n")),
		mcp.WithArray("pool", mcp.Description("Input parameter: The value can be repeated to retrieve multiple matching values (OR condition).")),
		mcp.WithArray("dag_ids", mcp.Description("Input parameter: Return objects with specific DAG IDs.\nThe value can be repeated to retrieve multiple matching values (OR condition).")),
		mcp.WithString("duration_lte", mcp.Description("Input parameter: Returns objects less than or equal to the specified values.\n\nThis can be combined with duration_gte parameter to receive only the selected range.\n")),
		mcp.WithString("start_date_gte", mcp.Description("Input parameter: Returns objects greater or equal the specified date.\n\nThis can be combined with start_date_lte parameter to receive only the selected period.\n")),
		mcp.WithString("end_date_gte", mcp.Description("Input parameter: Returns objects greater or equal the specified date.\n\nThis can be combined with start_date_lte parameter to receive only the selected period.\n")),
		mcp.WithArray("queue", mcp.Description("Input parameter: The value can be repeated to retrieve multiple matching values (OR condition).")),
		mcp.WithString("start_date_lte", mcp.Description("Input parameter: Returns objects less or equal the specified date.\n\nThis can be combined with start_date_gte parameter to receive only the selected period.\n")),
		mcp.WithArray("state", mcp.Description("Input parameter: The value can be repeated to retrieve multiple matching values (OR condition).")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_task_instances_batchHandler(cfg),
	}
}
