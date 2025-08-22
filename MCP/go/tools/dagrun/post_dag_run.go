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

func Post_dag_runHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		var requestBody models.DAGRun
		
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
		url := fmt.Sprintf("%s/dags/%s/dagRuns%s", cfg.BaseURL, queryString)
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
		var result models.DAGRun
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

func CreatePost_dag_runTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_dags_dag_id_dagRuns",
		mcp.WithDescription("Trigger a new DAG run"),
		mcp.WithString("note", mcp.Description("Input parameter: Contains manually entered notes by the user about the DagRun.\n\n*New in version 2.5.0*\n")),
		mcp.WithString("run_type", mcp.Description("")),
		mcp.WithString("start_date", mcp.Description("Input parameter: The start time. The time when DAG run was actually created.\n\n*Changed in version 2.1.3*&#58; Field becomes nullable.\n")),
		mcp.WithString("data_interval_end", mcp.Description("")),
		mcp.WithString("end_date", mcp.Description("")),
		mcp.WithString("dag_id", mcp.Description("")),
		mcp.WithString("last_scheduling_decision", mcp.Description("")),
		mcp.WithObject("conf", mcp.Description("Input parameter: JSON object describing additional configuration parameters.\n\nThe value of this field can be set only when creating the object. If you try to modify the\nfield of an existing object, the request fails with an BAD_REQUEST error.\n")),
		mcp.WithBoolean("external_trigger", mcp.Description("")),
		mcp.WithString("logical_date", mcp.Description("Input parameter: The logical date (previously called execution date). This is the time or interval covered by\nthis DAG run, according to the DAG definition.\n\nThe value of this field can be set only when creating the object. If you try to modify the\nfield of an existing object, the request fails with an BAD_REQUEST error.\n\nThis together with DAG_ID are a unique key.\n\n*New in version 2.2.0*\n")),
		mcp.WithString("state", mcp.Description("Input parameter: DAG State.\n\n*Changed in version 2.1.3*&#58; 'queued' is added as a possible value.\n")),
		mcp.WithString("dag_run_id", mcp.Description("Input parameter: Run ID.\n\nThe value of this field can be set only when creating the object. If you try to modify the\nfield of an existing object, the request fails with an BAD_REQUEST error.\n\nIf not provided, a value will be generated based on execution_date.\n\nIf the specified dag_run_id is in use, the creation request fails with an ALREADY_EXISTS error.\n\nThis together with DAG_ID are a unique key.\n")),
		mcp.WithString("data_interval_start", mcp.Description("")),
		mcp.WithString("execution_date", mcp.Description("Input parameter: The execution date. This is the same as logical_date, kept for backwards compatibility.\nIf both this field and logical_date are provided but with different values, the request\nwill fail with an BAD_REQUEST error.\n\n*Changed in version 2.2.0*&#58; Field becomes nullable.\n\n*Deprecated since version 2.2.0*&#58; Use 'logical_date' instead.\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_dag_runHandler(cfg),
	}
}
