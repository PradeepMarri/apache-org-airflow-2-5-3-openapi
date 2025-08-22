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

func Post_set_task_instances_stateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		var requestBody models.UpdateTaskInstancesState
		
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
		url := fmt.Sprintf("%s/dags/%s/updateTaskInstancesState%s", cfg.BaseURL, queryString)
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
		var result models.TaskInstanceReferenceCollection
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

func CreatePost_set_task_instances_stateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_dags_dag_id_updateTaskInstancesState",
		mcp.WithDescription("Set a state of task instances"),
		mcp.WithBoolean("dry_run", mcp.Description("Input parameter: If set, don't actually run this operation. The response will contain a list of task instances\nplanned to be affected, but won't be modified in any way.\n")),
		mcp.WithBoolean("include_future", mcp.Description("Input parameter: If set to True, also tasks from future DAG Runs are affected.")),
		mcp.WithBoolean("include_upstream", mcp.Description("Input parameter: If set to true, upstream tasks are also affected.")),
		mcp.WithString("new_state", mcp.Description("Input parameter: Expected new state.")),
		mcp.WithString("task_id", mcp.Description("Input parameter: The task ID.")),
		mcp.WithString("execution_date", mcp.Description("Input parameter: The execution date. Either set this or dag_run_id but not both.")),
		mcp.WithBoolean("include_past", mcp.Description("Input parameter: If set to True, also tasks from past DAG Runs are affected.")),
		mcp.WithBoolean("include_downstream", mcp.Description("Input parameter: If set to true, downstream tasks are also affected.")),
		mcp.WithString("dag_run_id", mcp.Description("Input parameter: The task instance's DAG run ID. Either set this or execution_date but not both.\n\n*New in version 2.3.0*\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_set_task_instances_stateHandler(cfg),
	}
}
