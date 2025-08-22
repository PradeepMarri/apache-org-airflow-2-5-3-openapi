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

func Post_userHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		var requestBody models.User
		
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
		url := fmt.Sprintf("%s/users%s", cfg.BaseURL, queryString)
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
		var result models.User
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

func CreatePost_userTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_users",
		mcp.WithDescription("Create a user"),
		mcp.WithArray("roles", mcp.Description("Input parameter: User roles.\n\n*Changed in version 2.2.0*&#58; Field is no longer read-only.\n")),
		mcp.WithString("first_name", mcp.Description("Input parameter: The user's first name.\n\n*Changed in version 2.4.0*&#58; The requirement for this to be non-empty was removed.\n")),
		mcp.WithString("last_login", mcp.Description("Input parameter: The last user login")),
		mcp.WithNumber("login_count", mcp.Description("Input parameter: The login count")),
		mcp.WithString("username", mcp.Description("Input parameter: The username.\n\n*Changed in version 2.2.0*&#58; A minimum character length requirement ('minLength') is added.\n")),
		mcp.WithString("changed_on", mcp.Description("Input parameter: The date user was changed")),
		mcp.WithString("created_on", mcp.Description("Input parameter: The date user was created")),
		mcp.WithString("email", mcp.Description("Input parameter: The user's email.\n\n*Changed in version 2.2.0*&#58; A minimum character length requirement ('minLength') is added.\n")),
		mcp.WithNumber("failed_login_count", mcp.Description("Input parameter: The number of times the login failed")),
		mcp.WithString("last_name", mcp.Description("Input parameter: The user's last name.\n\n*Changed in version 2.4.0*&#58; The requirement for this to be non-empty was removed.\n")),
		mcp.WithBoolean("active", mcp.Description("Input parameter: Whether the user is active")),
		mcp.WithString("password", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_userHandler(cfg),
	}
}
