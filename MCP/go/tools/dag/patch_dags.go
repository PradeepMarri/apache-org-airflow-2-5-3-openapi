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

func Patch_dagsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["tags"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tags=%v", val))
		}
		if val, ok := args["update_mask"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("update_mask=%v", val))
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
		// Create properly typed request body using the generated schema
		var requestBody models.DAG
		
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
		url := fmt.Sprintf("%s/dags%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
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

func CreatePatch_dagsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_dags",
		mcp.WithDescription("Update DAGs"),
		mcp.WithNumber("limit", mcp.Description("The numbers of items to return.")),
		mcp.WithNumber("offset", mcp.Description("The number of items to skip before starting to collect the result set.")),
		mcp.WithArray("tags", mcp.Description("List of tags to filter results.\n\n*New in version 2.2.0*\n")),
		mcp.WithArray("update_mask", mcp.Description("The fields to update on the resource. If absent or empty, all modifiable fields are updated.\nA comma-separated list of fully qualified names of fields.\n")),
		mcp.WithBoolean("only_active", mcp.Description("Only filter active DAGs.\n\n*New in version 2.1.1*\n")),
		mcp.WithString("dag_id_pattern", mcp.Required(), mcp.Description("If set, only update DAGs with dag_ids matching this pattern.\n")),
		mcp.WithBoolean("has_task_concurrency_limits", mcp.Description("Input parameter: Whether the DAG has task concurrency limits\n\n*New in version 2.3.0*\n")),
		mcp.WithString("pickle_id", mcp.Description("Input parameter: Foreign key to the latest pickle_id\n\n*New in version 2.3.0*\n")),
		mcp.WithBoolean("is_subdag", mcp.Description("Input parameter: Whether the DAG is SubDAG.")),
		mcp.WithNumber("max_active_runs", mcp.Description("Input parameter: Maximum number of active DAG runs for the DAG\n\n*New in version 2.3.0*\n")),
		mcp.WithArray("owners", mcp.Description("")),
		mcp.WithBoolean("scheduler_lock", mcp.Description("Input parameter: Whether (one of) the scheduler is scheduling this DAG at the moment\n\n*New in version 2.3.0*\n")),
		mcp.WithString("next_dagrun", mcp.Description("Input parameter: The logical date of the next dag run.\n\n*New in version 2.3.0*\n")),
		mcp.WithString("next_dagrun_data_interval_start", mcp.Description("Input parameter: The start of the interval of the next dag run.\n\n*New in version 2.3.0*\n")),
		mcp.WithString("root_dag_id", mcp.Description("Input parameter: If the DAG is SubDAG then it is the top level DAG identifier. Otherwise, null.")),
		mcp.WithString("default_view", mcp.Description("Input parameter: Default view of the DAG inside the webserver\n\n*New in version 2.3.0*\n")),
		mcp.WithString("timetable_description", mcp.Description("Input parameter: Timetable/Schedule Interval description.\n\n*New in version 2.3.0*\n")),
		mcp.WithString("last_parsed_time", mcp.Description("Input parameter: The last time the DAG was parsed.\n\n*New in version 2.3.0*\n")),
		mcp.WithString("next_dagrun_data_interval_end", mcp.Description("Input parameter: The end of the interval of the next dag run.\n\n*New in version 2.3.0*\n")),
		mcp.WithBoolean("has_import_errors", mcp.Description("Input parameter: Whether the DAG has import errors\n\n*New in version 2.3.0*\n")),
		mcp.WithObject("schedule_interval", mcp.Description("Input parameter: Schedule interval. Defines how often DAG runs, this object gets added to your latest task instance's\nexecution_date to figure out the next schedule.\n")),
		mcp.WithString("last_pickled", mcp.Description("Input parameter: The last time the DAG was pickled.\n\n*New in version 2.3.0*\n")),
		mcp.WithString("description", mcp.Description("Input parameter: User-provided DAG description, which can consist of several sentences or paragraphs that describe DAG contents.\n")),
		mcp.WithBoolean("is_paused", mcp.Description("Input parameter: Whether the DAG is paused.")),
		mcp.WithString("last_expired", mcp.Description("Input parameter: Time when the DAG last received a refresh signal\n(e.g. the DAG's \"refresh\" button was clicked in the web UI)\n\n*New in version 2.3.0*\n")),
		mcp.WithString("dag_id", mcp.Description("Input parameter: The ID of the DAG.")),
		mcp.WithString("fileloc", mcp.Description("Input parameter: The absolute path to the file.")),
		mcp.WithArray("tags", mcp.Description("Input parameter: List of tags.")),
		mcp.WithBoolean("is_active", mcp.Description("Input parameter: Whether the DAG is currently seen by the scheduler(s).\n\n*New in version 2.1.1*\n\n*Changed in version 2.2.0*&#58; Field is read-only.\n")),
		mcp.WithNumber("max_active_tasks", mcp.Description("Input parameter: Maximum number of active tasks that can be run on the DAG\n\n*New in version 2.3.0*\n")),
		mcp.WithString("next_dagrun_create_after", mcp.Description("Input parameter: Earliest time at which this ``next_dagrun`` can be created.\n\n*New in version 2.3.0*\n")),
		mcp.WithString("file_token", mcp.Description("Input parameter: The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change.\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Patch_dagsHandler(cfg),
	}
}
