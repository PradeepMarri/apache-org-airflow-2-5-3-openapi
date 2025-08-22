package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// UserCollection represents the UserCollection schema from the OpenAPI specification
type UserCollection struct {
	Users []UserCollectionItem `json:"users,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// ClearDagRun represents the ClearDagRun schema from the OpenAPI specification
type ClearDagRun struct {
	Dry_run bool `json:"dry_run,omitempty"` // If set, don't actually run this operation. The response will contain a list of task instances planned to be cleaned, but not modified in any way.
}

// CronExpression represents the CronExpression schema from the OpenAPI specification
type CronExpression struct {
	TypeField string `json:"__type"`
	Value string `json:"value"`
}

// Role represents the Role schema from the OpenAPI specification
type Role struct {
	Actions []ActionResource `json:"actions,omitempty"`
	Name string `json:"name,omitempty"` // The name of the role *Changed in version 2.3.0*&#58; A minimum character length requirement ('minLength') is added.
}

// HealthInfo represents the HealthInfo schema from the OpenAPI specification
type HealthInfo struct {
	Metadatabase MetadatabaseStatus `json:"metadatabase,omitempty"` // The status of the metadatabase.
	Scheduler SchedulerStatus `json:"scheduler,omitempty"` // The status and the latest scheduler heartbeat.
}

// TimeDelta represents the TimeDelta schema from the OpenAPI specification
type TimeDelta struct {
	TypeField string `json:"__type"`
	Days int `json:"days"`
	Microseconds int `json:"microseconds"`
	Seconds int `json:"seconds"`
}

// Job represents the Job schema from the OpenAPI specification
type Job struct {
	Dag_id string `json:"dag_id,omitempty"`
	End_date string `json:"end_date,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Job_type string `json:"job_type,omitempty"`
	Unixname string `json:"unixname,omitempty"`
	Id int `json:"id,omitempty"`
	Latest_heartbeat string `json:"latest_heartbeat,omitempty"`
	Start_date string `json:"start_date,omitempty"`
	State string `json:"state,omitempty"`
	Executor_class string `json:"executor_class,omitempty"`
}

// ProviderCollection represents the ProviderCollection schema from the OpenAPI specification
type ProviderCollection struct {
	Providers []Provider `json:"providers,omitempty"`
}

// ListTaskInstanceForm represents the ListTaskInstanceForm schema from the OpenAPI specification
type ListTaskInstanceForm struct {
	Start_date_lte string `json:"start_date_lte,omitempty"` // Returns objects less or equal the specified date. This can be combined with start_date_gte parameter to receive only the selected period.
	State []string `json:"state,omitempty"` // The value can be repeated to retrieve multiple matching values (OR condition).
	Duration_gte float64 `json:"duration_gte,omitempty"` // Returns objects greater than or equal to the specified values. This can be combined with duration_lte parameter to receive only the selected period.
	End_date_lte string `json:"end_date_lte,omitempty"` // Returns objects less than or equal to the specified date. This can be combined with start_date_gte parameter to receive only the selected period.
	Execution_date_lte string `json:"execution_date_lte,omitempty"` // Returns objects less than or equal to the specified date. This can be combined with execution_date_gte parameter to receive only the selected period.
	Execution_date_gte string `json:"execution_date_gte,omitempty"` // Returns objects greater or equal to the specified date. This can be combined with execution_date_lte parameter to receive only the selected period.
	Pool []string `json:"pool,omitempty"` // The value can be repeated to retrieve multiple matching values (OR condition).
	Dag_ids []string `json:"dag_ids,omitempty"` // Return objects with specific DAG IDs. The value can be repeated to retrieve multiple matching values (OR condition).
	Duration_lte float64 `json:"duration_lte,omitempty"` // Returns objects less than or equal to the specified values. This can be combined with duration_gte parameter to receive only the selected range.
	Start_date_gte string `json:"start_date_gte,omitempty"` // Returns objects greater or equal the specified date. This can be combined with start_date_lte parameter to receive only the selected period.
	End_date_gte string `json:"end_date_gte,omitempty"` // Returns objects greater or equal the specified date. This can be combined with start_date_lte parameter to receive only the selected period.
	Queue []string `json:"queue,omitempty"` // The value can be repeated to retrieve multiple matching values (OR condition).
}

// ConnectionCollectionItem represents the ConnectionCollectionItem schema from the OpenAPI specification
type ConnectionCollectionItem struct {
	Conn_type string `json:"conn_type,omitempty"` // The connection type.
	Connection_id string `json:"connection_id,omitempty"` // The connection ID.
	Description string `json:"description,omitempty"` // The description of the connection.
	Host string `json:"host,omitempty"` // Host of the connection.
	Login string `json:"login,omitempty"` // Login of the connection.
	Port int `json:"port,omitempty"` // Port of the connection.
	Schema string `json:"schema,omitempty"` // Schema of the connection.
}

// DatasetCollection represents the DatasetCollection schema from the OpenAPI specification
type DatasetCollection struct {
	Datasets []Dataset `json:"datasets,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// Dataset represents the Dataset schema from the OpenAPI specification
type Dataset struct {
	Id int `json:"id,omitempty"` // The dataset id
	Producing_tasks []TaskOutletDatasetReference `json:"producing_tasks,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // The dataset update time
	Uri string `json:"uri,omitempty"` // The dataset uri
	Consuming_dags []DagScheduleDatasetReference `json:"consuming_dags,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The dataset creation time
	Extra map[string]interface{} `json:"extra,omitempty"` // The dataset extra
}

// ListDagRunsForm represents the ListDagRunsForm schema from the OpenAPI specification
type ListDagRunsForm struct {
	Dag_ids []string `json:"dag_ids,omitempty"` // Return objects with specific DAG IDs. The value can be repeated to retrieve multiple matching values (OR condition).
	Page_offset int `json:"page_offset,omitempty"` // The number of items to skip before starting to collect the result set.
	Execution_date_lte string `json:"execution_date_lte,omitempty"` // Returns objects less than or equal to the specified date. This can be combined with execution_date_gte key to receive only the selected period.
	Page_limit int `json:"page_limit,omitempty"` // The numbers of items to return.
	States []string `json:"states,omitempty"` // Return objects with specific states. The value can be repeated to retrieve multiple matching values (OR condition).
	End_date_gte string `json:"end_date_gte,omitempty"` // Returns objects greater or equal the specified date. This can be combined with end_date_lte parameter to receive only the selected period.
	End_date_lte string `json:"end_date_lte,omitempty"` // Returns objects less than or equal to the specified date. This can be combined with end_date_gte parameter to receive only the selected period.
	Order_by string `json:"order_by,omitempty"` // The name of the field to order the results by. Prefix a field name with `-` to reverse the sort order. *New in version 2.1.0*
	Start_date_gte string `json:"start_date_gte,omitempty"` // Returns objects greater or equal the specified date. This can be combined with start_date_lte key to receive only the selected period.
	Execution_date_gte string `json:"execution_date_gte,omitempty"` // Returns objects greater or equal to the specified date. This can be combined with execution_date_lte key to receive only the selected period.
	Start_date_lte string `json:"start_date_lte,omitempty"` // Returns objects less or equal the specified date. This can be combined with start_date_gte parameter to receive only the selected period
}

// DAG represents the DAG schema from the OpenAPI specification
type DAG struct {
	Fileloc string `json:"fileloc,omitempty"` // The absolute path to the file.
	Tags []Tag `json:"tags,omitempty"` // List of tags.
	Is_active bool `json:"is_active,omitempty"` // Whether the DAG is currently seen by the scheduler(s). *New in version 2.1.1* *Changed in version 2.2.0*&#58; Field is read-only.
	Max_active_tasks int `json:"max_active_tasks,omitempty"` // Maximum number of active tasks that can be run on the DAG *New in version 2.3.0*
	Next_dagrun_create_after string `json:"next_dagrun_create_after,omitempty"` // Earliest time at which this ``next_dagrun`` can be created. *New in version 2.3.0*
	File_token string `json:"file_token,omitempty"` // The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change.
	Has_task_concurrency_limits bool `json:"has_task_concurrency_limits,omitempty"` // Whether the DAG has task concurrency limits *New in version 2.3.0*
	Pickle_id string `json:"pickle_id,omitempty"` // Foreign key to the latest pickle_id *New in version 2.3.0*
	Is_subdag bool `json:"is_subdag,omitempty"` // Whether the DAG is SubDAG.
	Max_active_runs int `json:"max_active_runs,omitempty"` // Maximum number of active DAG runs for the DAG *New in version 2.3.0*
	Owners []string `json:"owners,omitempty"`
	Scheduler_lock bool `json:"scheduler_lock,omitempty"` // Whether (one of) the scheduler is scheduling this DAG at the moment *New in version 2.3.0*
	Next_dagrun string `json:"next_dagrun,omitempty"` // The logical date of the next dag run. *New in version 2.3.0*
	Next_dagrun_data_interval_start string `json:"next_dagrun_data_interval_start,omitempty"` // The start of the interval of the next dag run. *New in version 2.3.0*
	Root_dag_id string `json:"root_dag_id,omitempty"` // If the DAG is SubDAG then it is the top level DAG identifier. Otherwise, null.
	Default_view string `json:"default_view,omitempty"` // Default view of the DAG inside the webserver *New in version 2.3.0*
	Timetable_description string `json:"timetable_description,omitempty"` // Timetable/Schedule Interval description. *New in version 2.3.0*
	Last_parsed_time string `json:"last_parsed_time,omitempty"` // The last time the DAG was parsed. *New in version 2.3.0*
	Next_dagrun_data_interval_end string `json:"next_dagrun_data_interval_end,omitempty"` // The end of the interval of the next dag run. *New in version 2.3.0*
	Has_import_errors bool `json:"has_import_errors,omitempty"` // Whether the DAG has import errors *New in version 2.3.0*
	Schedule_interval interface{} `json:"schedule_interval,omitempty"` // Schedule interval. Defines how often DAG runs, this object gets added to your latest task instance's execution_date to figure out the next schedule.
	Last_pickled string `json:"last_pickled,omitempty"` // The last time the DAG was pickled. *New in version 2.3.0*
	Description string `json:"description,omitempty"` // User-provided DAG description, which can consist of several sentences or paragraphs that describe DAG contents.
	Is_paused bool `json:"is_paused,omitempty"` // Whether the DAG is paused.
	Last_expired string `json:"last_expired,omitempty"` // Time when the DAG last received a refresh signal (e.g. the DAG's "refresh" button was clicked in the web UI) *New in version 2.3.0*
	Dag_id string `json:"dag_id,omitempty"` // The ID of the DAG.
}

// VersionInfo represents the VersionInfo schema from the OpenAPI specification
type VersionInfo struct {
	Git_version string `json:"git_version,omitempty"` // The git version (including git commit hash)
	Version string `json:"version,omitempty"` // The version of Airflow
}

// ConfigSection represents the ConfigSection schema from the OpenAPI specification
type ConfigSection struct {
	Name string `json:"name,omitempty"`
	Options []ConfigOption `json:"options,omitempty"`
}

// DagScheduleDatasetReference represents the DagScheduleDatasetReference schema from the OpenAPI specification
type DagScheduleDatasetReference struct {
	Dag_id string `json:"dag_id,omitempty"` // The DAG ID that depends on the dataset.
	Updated_at string `json:"updated_at,omitempty"` // The dataset reference update time
	Created_at string `json:"created_at,omitempty"` // The dataset reference creation time
}

// DAGCollection represents the DAGCollection schema from the OpenAPI specification
type DAGCollection struct {
	Dags []DAG `json:"dags,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// SetDagRunNote represents the SetDagRunNote schema from the OpenAPI specification
type SetDagRunNote struct {
	Note string `json:"note,omitempty"` // Custom notes left by users for this Dag Run.
}

// Task represents the Task schema from the OpenAPI specification
type Task struct {
	Execution_timeout TimeDelta `json:"execution_timeout,omitempty"` // Time delta
	Retry_exponential_backoff bool `json:"retry_exponential_backoff,omitempty"`
	Owner string `json:"owner,omitempty"`
	Priority_weight float64 `json:"priority_weight,omitempty"`
	Start_date string `json:"start_date,omitempty"`
	Pool string `json:"pool,omitempty"`
	Retries float64 `json:"retries,omitempty"`
	Ui_color string `json:"ui_color,omitempty"` // Color in hexadecimal notation.
	Weight_rule string `json:"weight_rule,omitempty"` // Weight rule.
	Extra_links []map[string]interface{} `json:"extra_links,omitempty"`
	Is_mapped bool `json:"is_mapped,omitempty"`
	Queue string `json:"queue,omitempty"`
	Retry_delay TimeDelta `json:"retry_delay,omitempty"` // Time delta
	Template_fields []string `json:"template_fields,omitempty"`
	Ui_fgcolor string `json:"ui_fgcolor,omitempty"` // Color in hexadecimal notation.
	Class_ref ClassReference `json:"class_ref,omitempty"` // Class reference
	Depends_on_past bool `json:"depends_on_past,omitempty"`
	End_date string `json:"end_date,omitempty"`
	Wait_for_downstream bool `json:"wait_for_downstream,omitempty"`
	Sub_dag DAG `json:"sub_dag,omitempty"` // DAG
	Trigger_rule string `json:"trigger_rule,omitempty"` // Trigger rule. *Changed in version 2.2.0*&#58; 'none_failed_min_one_success' is added as a possible value.
	Pool_slots float64 `json:"pool_slots,omitempty"`
	Task_id string `json:"task_id,omitempty"`
	Downstream_task_ids []string `json:"downstream_task_ids,omitempty"`
}

// RoleCollection represents the RoleCollection schema from the OpenAPI specification
type RoleCollection struct {
	Roles []Role `json:"roles,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// ConnectionTest represents the ConnectionTest schema from the OpenAPI specification
type ConnectionTest struct {
	Message string `json:"message,omitempty"` // The success or failure message of the request.
	Status bool `json:"status,omitempty"` // The status of the request.
}

// UpdateDagRunState represents the UpdateDagRunState schema from the OpenAPI specification
type UpdateDagRunState struct {
	State string `json:"state,omitempty"` // The state to set this DagRun
}

// ConfigOption represents the ConfigOption schema from the OpenAPI specification
type ConfigOption struct {
	Value string `json:"value,omitempty"`
	Key string `json:"key,omitempty"`
}

// ConnectionCollection represents the ConnectionCollection schema from the OpenAPI specification
type ConnectionCollection struct {
	Connections []ConnectionCollectionItem `json:"connections,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// User represents the User schema from the OpenAPI specification
type User struct {
	Roles []map[string]interface{} `json:"roles,omitempty"` // User roles. *Changed in version 2.2.0*&#58; Field is no longer read-only.
	First_name string `json:"first_name,omitempty"` // The user's first name. *Changed in version 2.4.0*&#58; The requirement for this to be non-empty was removed.
	Last_login string `json:"last_login,omitempty"` // The last user login
	Login_count int `json:"login_count,omitempty"` // The login count
	Username string `json:"username,omitempty"` // The username. *Changed in version 2.2.0*&#58; A minimum character length requirement ('minLength') is added.
	Changed_on string `json:"changed_on,omitempty"` // The date user was changed
	Created_on string `json:"created_on,omitempty"` // The date user was created
	Email string `json:"email,omitempty"` // The user's email. *Changed in version 2.2.0*&#58; A minimum character length requirement ('minLength') is added.
	Failed_login_count int `json:"failed_login_count,omitempty"` // The number of times the login failed
	Last_name string `json:"last_name,omitempty"` // The user's last name. *Changed in version 2.4.0*&#58; The requirement for this to be non-empty was removed.
	Active bool `json:"active,omitempty"` // Whether the user is active
	Password string `json:"password,omitempty"`
}

// DagWarning represents the DagWarning schema from the OpenAPI specification
type DagWarning struct {
	Dag_id string `json:"dag_id,omitempty"` // The dag_id.
	Message string `json:"message,omitempty"` // The message for the dag warning.
	Timestamp string `json:"timestamp,omitempty"` // The time when this warning was logged.
	Warning_type string `json:"warning_type,omitempty"` // The warning type for the dag warning.
}

// DatasetEvent represents the DatasetEvent schema from the OpenAPI specification
type DatasetEvent struct {
	Source_run_id string `json:"source_run_id,omitempty"` // The DAG run ID that updated the dataset.
	Timestamp string `json:"timestamp,omitempty"` // The dataset event creation time
	Created_dagruns []BasicDAGRun `json:"created_dagruns,omitempty"`
	Source_dag_id string `json:"source_dag_id,omitempty"` // The DAG ID that updated the dataset.
	Source_task_id string `json:"source_task_id,omitempty"` // The task ID that updated the dataset.
	Dataset_id int `json:"dataset_id,omitempty"` // The dataset id
	Dataset_uri string `json:"dataset_uri,omitempty"` // The URI of the dataset
	Extra map[string]interface{} `json:"extra,omitempty"` // The dataset event extra
	Source_map_index int `json:"source_map_index,omitempty"` // The task map index that updated the dataset.
}

// Variable represents the Variable schema from the OpenAPI specification
type Variable struct {
	Description string `json:"description,omitempty"` // The description of the variable. *New in version 2.4.0*
	Key string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// DatasetEventCollection represents the DatasetEventCollection schema from the OpenAPI specification
type DatasetEventCollection struct {
	Dataset_events []DatasetEvent `json:"dataset_events,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// XComCollection represents the XComCollection schema from the OpenAPI specification
type XComCollection struct {
	Xcom_entries []XComCollectionItem `json:"xcom_entries,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// ClearTaskInstances represents the ClearTaskInstances schema from the OpenAPI specification
type ClearTaskInstances struct {
	Include_past bool `json:"include_past,omitempty"` // If set to True, also tasks from past DAG Runs are affected.
	Include_subdags bool `json:"include_subdags,omitempty"` // Clear tasks in subdags and clear external tasks indicated by ExternalTaskMarker.
	Dry_run bool `json:"dry_run,omitempty"` // If set, don't actually run this operation. The response will contain a list of task instances planned to be cleaned, but not modified in any way.
	End_date string `json:"end_date,omitempty"` // The maximum execution date to clear.
	Include_parentdag bool `json:"include_parentdag,omitempty"` // Clear tasks in the parent dag of the subdag.
	Include_upstream bool `json:"include_upstream,omitempty"` // If set to true, upstream tasks are also affected.
	Only_running bool `json:"only_running,omitempty"` // Only clear running tasks.
	Start_date string `json:"start_date,omitempty"` // The minimum execution date to clear.
	Task_ids []string `json:"task_ids,omitempty"` // A list of task ids to clear. *New in version 2.1.0*
	Include_future bool `json:"include_future,omitempty"` // If set to True, also tasks from future DAG Runs are affected.
	Only_failed bool `json:"only_failed,omitempty"` // Only clear failed tasks.
	Reset_dag_runs bool `json:"reset_dag_runs,omitempty"` // Set state of DAG runs to RUNNING.
	Dag_run_id string `json:"dag_run_id,omitempty"` // The DagRun ID for this task instance
	Include_downstream bool `json:"include_downstream,omitempty"` // If set to true, downstream tasks are also affected.
}

// PoolCollection represents the PoolCollection schema from the OpenAPI specification
type PoolCollection struct {
	Pools []Pool `json:"pools,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// TaskInstance represents the TaskInstance schema from the OpenAPI specification
type TaskInstance struct {
	Queue string `json:"queue,omitempty"`
	Sla_miss SLAMiss `json:"sla_miss,omitempty"`
	Duration float64 `json:"duration,omitempty"`
	Pool_slots int `json:"pool_slots,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Start_date string `json:"start_date,omitempty"`
	Unixname string `json:"unixname,omitempty"`
	Dag_run_id string `json:"dag_run_id,omitempty"` // The DagRun ID for this task instance *New in version 2.3.0*
	Triggerer_job Job `json:"triggerer_job,omitempty"`
	Pid int `json:"pid,omitempty"`
	End_date string `json:"end_date,omitempty"`
	Note string `json:"note,omitempty"` // Contains manually entered notes by the user about the TaskInstance. *New in version 2.5.0*
	Priority_weight int `json:"priority_weight,omitempty"`
	Execution_date string `json:"execution_date,omitempty"`
	Max_tries int `json:"max_tries,omitempty"`
	Operator string `json:"operator,omitempty"` // *Changed in version 2.1.1*&#58; Field becomes nullable.
	Queued_when string `json:"queued_when,omitempty"`
	Task_id string `json:"task_id,omitempty"`
	Dag_id string `json:"dag_id,omitempty"`
	Map_index int `json:"map_index,omitempty"`
	Try_number int `json:"try_number,omitempty"`
	State string `json:"state,omitempty"` // Task state. *Changed in version 2.0.2*&#58; 'removed' is added as a possible value. *Changed in version 2.2.0*&#58; 'deferred' is added as a possible value. *Changed in version 2.4.0*&#58; 'sensing' state has been removed. *Changed in version 2.4.2*&#58; 'restarting' is added as a possible value
	Trigger Trigger `json:"trigger,omitempty"`
	Pool string `json:"pool,omitempty"`
	Rendered_fields map[string]interface{} `json:"rendered_fields,omitempty"` // JSON object describing rendered fields. *New in version 2.3.0*
	Executor_config string `json:"executor_config,omitempty"`
}

// PluginCollection represents the PluginCollection schema from the OpenAPI specification
type PluginCollection struct {
	Plugins []PluginCollectionItem `json:"plugins,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// UserCollectionItem represents the UserCollectionItem schema from the OpenAPI specification
type UserCollectionItem struct {
	Last_login string `json:"last_login,omitempty"` // The last user login
	Login_count int `json:"login_count,omitempty"` // The login count
	Username string `json:"username,omitempty"` // The username. *Changed in version 2.2.0*&#58; A minimum character length requirement ('minLength') is added.
	Changed_on string `json:"changed_on,omitempty"` // The date user was changed
	Created_on string `json:"created_on,omitempty"` // The date user was created
	Email string `json:"email,omitempty"` // The user's email. *Changed in version 2.2.0*&#58; A minimum character length requirement ('minLength') is added.
	Failed_login_count int `json:"failed_login_count,omitempty"` // The number of times the login failed
	Last_name string `json:"last_name,omitempty"` // The user's last name. *Changed in version 2.4.0*&#58; The requirement for this to be non-empty was removed.
	Active bool `json:"active,omitempty"` // Whether the user is active
	Roles []map[string]interface{} `json:"roles,omitempty"` // User roles. *Changed in version 2.2.0*&#58; Field is no longer read-only.
	First_name string `json:"first_name,omitempty"` // The user's first name. *Changed in version 2.4.0*&#58; The requirement for this to be non-empty was removed.
}

// UpdateTaskInstance represents the UpdateTaskInstance schema from the OpenAPI specification
type UpdateTaskInstance struct {
	Dry_run bool `json:"dry_run,omitempty"` // If set, don't actually run this operation. The response will contain the task instance planned to be affected, but won't be modified in any way.
	New_state string `json:"new_state,omitempty"` // Expected new state.
}

// Connection represents the Connection schema from the OpenAPI specification
type Connection struct {
	Login string `json:"login,omitempty"` // Login of the connection.
	Port int `json:"port,omitempty"` // Port of the connection.
	Schema string `json:"schema,omitempty"` // Schema of the connection.
	Conn_type string `json:"conn_type,omitempty"` // The connection type.
	Connection_id string `json:"connection_id,omitempty"` // The connection ID.
	Description string `json:"description,omitempty"` // The description of the connection.
	Host string `json:"host,omitempty"` // Host of the connection.
	Password string `json:"password,omitempty"` // Password of the connection.
	Extra string `json:"extra,omitempty"` // Other values that cannot be put into another field, e.g. RSA keys.
}

// BasicDAGRun represents the BasicDAGRun schema from the OpenAPI specification
type BasicDAGRun struct {
	Data_interval_start string `json:"data_interval_start,omitempty"`
	End_date string `json:"end_date,omitempty"`
	Logical_date string `json:"logical_date,omitempty"` // The logical date (previously called execution date). This is the time or interval covered by this DAG run, according to the DAG definition. The value of this field can be set only when creating the object. If you try to modify the field of an existing object, the request fails with an BAD_REQUEST error. This together with DAG_ID are a unique key. *New in version 2.2.0*
	Run_id string `json:"run_id,omitempty"` // Run ID.
	Start_date string `json:"start_date,omitempty"` // The start time. The time when DAG run was actually created. *Changed in version 2.1.3*&#58; Field becomes nullable.
	State string `json:"state,omitempty"` // DAG State. *Changed in version 2.1.3*&#58; 'queued' is added as a possible value.
	Dag_id string `json:"dag_id,omitempty"`
	Data_interval_end string `json:"data_interval_end,omitempty"`
}

// RelativeDelta represents the RelativeDelta schema from the OpenAPI specification
type RelativeDelta struct {
	Seconds int `json:"seconds"`
	Hour int `json:"hour"`
	Months int `json:"months"`
	Day int `json:"day"`
	Hours int `json:"hours"`
	Microseconds int `json:"microseconds"`
	Minutes int `json:"minutes"`
	Month int `json:"month"`
	TypeField string `json:"__type"`
	Leapdays int `json:"leapdays"`
	Microsecond int `json:"microsecond"`
	Minute int `json:"minute"`
	Year int `json:"year"`
	Years int `json:"years"`
	Days int `json:"days"`
	Second int `json:"second"`
}

// PluginCollectionItem represents the PluginCollectionItem schema from the OpenAPI specification
type PluginCollectionItem struct {
	Source string `json:"source,omitempty"` // The plugin source
	Executors []string `json:"executors,omitempty"` // The plugin executors
	Name string `json:"name,omitempty"` // The name of the plugin
	Appbuilder_menu_items []map[string]interface{} `json:"appbuilder_menu_items,omitempty"` // The Flask Appbuilder menu items
	Global_operator_extra_links []map[string]interface{} `json:"global_operator_extra_links,omitempty"` // The global operator extra links
	Operator_extra_links []map[string]interface{} `json:"operator_extra_links,omitempty"` // Operator extra links
	Appbuilder_views []map[string]interface{} `json:"appbuilder_views,omitempty"` // The appuilder views
	Macros []map[string]interface{} `json:"macros,omitempty"` // The plugin macros
	Flask_blueprints []map[string]interface{} `json:"flask_blueprints,omitempty"` // The flask blueprints
	Hooks []string `json:"hooks,omitempty"` // The plugin hooks
}

// XCom represents the XCom schema from the OpenAPI specification
type XCom struct {
	Dag_id string `json:"dag_id,omitempty"`
	Execution_date string `json:"execution_date,omitempty"`
	Key string `json:"key,omitempty"`
	Task_id string `json:"task_id,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Value string `json:"value,omitempty"` // The value
}

// MetadatabaseStatus represents the MetadatabaseStatus schema from the OpenAPI specification
type MetadatabaseStatus struct {
	Status string `json:"status,omitempty"` // Health status
}

// TaskCollection represents the TaskCollection schema from the OpenAPI specification
type TaskCollection struct {
	Tasks []Task `json:"tasks,omitempty"`
}

// Action represents the Action schema from the OpenAPI specification
type Action struct {
	Name string `json:"name,omitempty"` // The name of the permission "action"
}

// ImportError represents the ImportError schema from the OpenAPI specification
type ImportError struct {
	Timestamp string `json:"timestamp,omitempty"` // The time when this error was created.
	Filename string `json:"filename,omitempty"` // The filename
	Import_error_id int `json:"import_error_id,omitempty"` // The import error ID.
	Stack_trace string `json:"stack_trace,omitempty"` // The full stackstrace..
}

// Provider represents the Provider schema from the OpenAPI specification
type Provider struct {
	Package_name string `json:"package_name,omitempty"` // The package name of the provider.
	Version string `json:"version,omitempty"` // The version of the provider.
	Description string `json:"description,omitempty"` // The description of the provider.
}

// UpdateTaskInstancesState represents the UpdateTaskInstancesState schema from the OpenAPI specification
type UpdateTaskInstancesState struct {
	Include_past bool `json:"include_past,omitempty"` // If set to True, also tasks from past DAG Runs are affected.
	Include_downstream bool `json:"include_downstream,omitempty"` // If set to true, downstream tasks are also affected.
	Dag_run_id string `json:"dag_run_id,omitempty"` // The task instance's DAG run ID. Either set this or execution_date but not both. *New in version 2.3.0*
	Dry_run bool `json:"dry_run,omitempty"` // If set, don't actually run this operation. The response will contain a list of task instances planned to be affected, but won't be modified in any way.
	Include_future bool `json:"include_future,omitempty"` // If set to True, also tasks from future DAG Runs are affected.
	Include_upstream bool `json:"include_upstream,omitempty"` // If set to true, upstream tasks are also affected.
	New_state string `json:"new_state,omitempty"` // Expected new state.
	Task_id string `json:"task_id,omitempty"` // The task ID.
	Execution_date string `json:"execution_date,omitempty"` // The execution date. Either set this or dag_run_id but not both.
}

// Trigger represents the Trigger schema from the OpenAPI specification
type Trigger struct {
	Kwargs string `json:"kwargs,omitempty"`
	Triggerer_id int `json:"triggerer_id,omitempty"`
	Classpath string `json:"classpath,omitempty"`
	Created_date string `json:"created_date,omitempty"`
	Id int `json:"id,omitempty"`
}

// ActionResource represents the ActionResource schema from the OpenAPI specification
type ActionResource struct {
	Action Action `json:"action,omitempty"` // An action Item. *New in version 2.1.0*
	Resource Resource `json:"resource,omitempty"` // A resource on which permissions are granted. *New in version 2.1.0*
}

// VariableCollectionItem represents the VariableCollectionItem schema from the OpenAPI specification
type VariableCollectionItem struct {
	Description string `json:"description,omitempty"` // The description of the variable. *New in version 2.4.0*
	Key string `json:"key,omitempty"`
}

// DAGRunCollection represents the DAGRunCollection schema from the OpenAPI specification
type DAGRunCollection struct {
	Dag_runs []DAGRun `json:"dag_runs,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// XComCollectionItem represents the XComCollectionItem schema from the OpenAPI specification
type XComCollectionItem struct {
	Key string `json:"key,omitempty"`
	Task_id string `json:"task_id,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Dag_id string `json:"dag_id,omitempty"`
	Execution_date string `json:"execution_date,omitempty"`
}

// TaskOutletDatasetReference represents the TaskOutletDatasetReference schema from the OpenAPI specification
type TaskOutletDatasetReference struct {
	Created_at string `json:"created_at,omitempty"` // The dataset creation time
	Dag_id string `json:"dag_id,omitempty"` // The DAG ID that updates the dataset.
	Task_id string `json:"task_id,omitempty"` // The task ID that updates the dataset.
	Updated_at string `json:"updated_at,omitempty"` // The dataset update time
}

// DAGDetail represents the DAGDetail schema from the OpenAPI specification
type DAGDetail struct {
	Fileloc string `json:"fileloc,omitempty"` // The absolute path to the file.
	Tags []Tag `json:"tags,omitempty"` // List of tags.
	Is_active bool `json:"is_active,omitempty"` // Whether the DAG is currently seen by the scheduler(s). *New in version 2.1.1* *Changed in version 2.2.0*&#58; Field is read-only.
	Max_active_tasks int `json:"max_active_tasks,omitempty"` // Maximum number of active tasks that can be run on the DAG *New in version 2.3.0*
	Next_dagrun_create_after string `json:"next_dagrun_create_after,omitempty"` // Earliest time at which this ``next_dagrun`` can be created. *New in version 2.3.0*
	File_token string `json:"file_token,omitempty"` // The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change.
	Has_task_concurrency_limits bool `json:"has_task_concurrency_limits,omitempty"` // Whether the DAG has task concurrency limits *New in version 2.3.0*
	Pickle_id string `json:"pickle_id,omitempty"` // Foreign key to the latest pickle_id *New in version 2.3.0*
	Is_subdag bool `json:"is_subdag,omitempty"` // Whether the DAG is SubDAG.
	Max_active_runs int `json:"max_active_runs,omitempty"` // Maximum number of active DAG runs for the DAG *New in version 2.3.0*
	Owners []string `json:"owners,omitempty"`
	Scheduler_lock bool `json:"scheduler_lock,omitempty"` // Whether (one of) the scheduler is scheduling this DAG at the moment *New in version 2.3.0*
	Next_dagrun string `json:"next_dagrun,omitempty"` // The logical date of the next dag run. *New in version 2.3.0*
	Next_dagrun_data_interval_start string `json:"next_dagrun_data_interval_start,omitempty"` // The start of the interval of the next dag run. *New in version 2.3.0*
	Root_dag_id string `json:"root_dag_id,omitempty"` // If the DAG is SubDAG then it is the top level DAG identifier. Otherwise, null.
	Default_view string `json:"default_view,omitempty"` // Default view of the DAG inside the webserver *New in version 2.3.0*
	Timetable_description string `json:"timetable_description,omitempty"` // Timetable/Schedule Interval description. *New in version 2.3.0*
	Last_parsed_time string `json:"last_parsed_time,omitempty"` // The last time the DAG was parsed. *New in version 2.3.0*
	Next_dagrun_data_interval_end string `json:"next_dagrun_data_interval_end,omitempty"` // The end of the interval of the next dag run. *New in version 2.3.0*
	Has_import_errors bool `json:"has_import_errors,omitempty"` // Whether the DAG has import errors *New in version 2.3.0*
	Schedule_interval interface{} `json:"schedule_interval,omitempty"` // Schedule interval. Defines how often DAG runs, this object gets added to your latest task instance's execution_date to figure out the next schedule.
	Last_pickled string `json:"last_pickled,omitempty"` // The last time the DAG was pickled. *New in version 2.3.0*
	Description string `json:"description,omitempty"` // User-provided DAG description, which can consist of several sentences or paragraphs that describe DAG contents.
	Is_paused bool `json:"is_paused,omitempty"` // Whether the DAG is paused.
	Last_expired string `json:"last_expired,omitempty"` // Time when the DAG last received a refresh signal (e.g. the DAG's "refresh" button was clicked in the web UI) *New in version 2.3.0*
	Dag_id string `json:"dag_id,omitempty"` // The ID of the DAG.
	End_date string `json:"end_date,omitempty"` // The DAG's end date. *New in version 2.3.0*.
	Orientation string `json:"orientation,omitempty"`
	Is_paused_upon_creation bool `json:"is_paused_upon_creation,omitempty"` // Whether the DAG is paused upon creation. *New in version 2.3.0*
	Render_template_as_native_obj bool `json:"render_template_as_native_obj,omitempty"` // Whether to render templates as native Python objects. *New in version 2.3.0*
	Start_date string `json:"start_date,omitempty"` // The DAG's start date. *Changed in version 2.0.1*&#58; Field becomes nullable.
	Timezone string `json:"timezone,omitempty"`
	Concurrency float64 `json:"concurrency,omitempty"`
	Dag_run_timeout TimeDelta `json:"dag_run_timeout,omitempty"` // Time delta
	Last_parsed string `json:"last_parsed,omitempty"` // The last time the DAG was parsed. *New in version 2.3.0*
	Params map[string]interface{} `json:"params,omitempty"` // User-specified DAG params. *New in version 2.0.1*
	Doc_md string `json:"doc_md,omitempty"`
	Template_search_path []string `json:"template_search_path,omitempty"` // The template search path. *New in version 2.3.0*
	Catchup bool `json:"catchup,omitempty"`
	Default_view string `json:"default_view,omitempty"`
}

// Config represents the Config schema from the OpenAPI specification
type Config struct {
	Sections []ConfigSection `json:"sections,omitempty"`
}

// ActionCollection represents the ActionCollection schema from the OpenAPI specification
type ActionCollection struct {
	Actions []Action `json:"actions,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// TaskInstanceReferenceCollection represents the TaskInstanceReferenceCollection schema from the OpenAPI specification
type TaskInstanceReferenceCollection struct {
	Task_instances []TaskInstanceReference `json:"task_instances,omitempty"`
}

// ClassReference represents the ClassReference schema from the OpenAPI specification
type ClassReference struct {
	Class_name string `json:"class_name,omitempty"`
	Module_path string `json:"module_path,omitempty"`
}

// EventLog represents the EventLog schema from the OpenAPI specification
type EventLog struct {
	Execution_date string `json:"execution_date,omitempty"` // When the event was dispatched for an object having execution_date, the value of this field.
	Extra string `json:"extra,omitempty"` // Other information that was not included in the other fields, e.g. the complete CLI command.
	Owner string `json:"owner,omitempty"` // Name of the user who triggered these events a.
	Task_id string `json:"task_id,omitempty"` // The DAG ID
	When string `json:"when,omitempty"` // The time when these events happened.
	Dag_id string `json:"dag_id,omitempty"` // The DAG ID
	Event string `json:"event,omitempty"` // A key describing the type of event.
	Event_log_id int `json:"event_log_id,omitempty"` // The event log ID
}

// Resource represents the Resource schema from the OpenAPI specification
type Resource struct {
	Name string `json:"name,omitempty"` // The name of the resource
}

// SLAMiss represents the SLAMiss schema from the OpenAPI specification
type SLAMiss struct {
	Description string `json:"description,omitempty"`
	Email_sent bool `json:"email_sent,omitempty"`
	Execution_date string `json:"execution_date,omitempty"`
	Notification_sent bool `json:"notification_sent,omitempty"`
	Task_id string `json:"task_id,omitempty"` // The task ID.
	Timestamp string `json:"timestamp,omitempty"`
	Dag_id string `json:"dag_id,omitempty"` // The DAG ID.
}

// ExtraLink represents the ExtraLink schema from the OpenAPI specification
type ExtraLink struct {
	Class_ref ClassReference `json:"class_ref,omitempty"` // Class reference
	Href string `json:"href,omitempty"`
	Name string `json:"name,omitempty"`
}

// Pool represents the Pool schema from the OpenAPI specification
type Pool struct {
	Description string `json:"description,omitempty"` // The description of the pool. *New in version 2.3.0*
	Name string `json:"name,omitempty"` // The name of pool.
	Occupied_slots int `json:"occupied_slots,omitempty"` // The number of slots used by running/queued tasks at the moment.
	Open_slots int `json:"open_slots,omitempty"` // The number of free slots at the moment.
	Queued_slots int `json:"queued_slots,omitempty"` // The number of slots used by queued tasks at the moment.
	Slots int `json:"slots,omitempty"` // The maximum number of slots that can be assigned to tasks. One job may occupy one or more slots.
	Used_slots int `json:"used_slots,omitempty"` // The number of slots used by running tasks at the moment.
}

// ExtraLinkCollection represents the ExtraLinkCollection schema from the OpenAPI specification
type ExtraLinkCollection struct {
	Extra_links []ExtraLink `json:"extra_links,omitempty"`
}

// CollectionInfo represents the CollectionInfo schema from the OpenAPI specification
type CollectionInfo struct {
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// DagWarningCollection represents the DagWarningCollection schema from the OpenAPI specification
type DagWarningCollection struct {
	Import_errors []DagWarning `json:"import_errors,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// SetTaskInstanceNote represents the SetTaskInstanceNote schema from the OpenAPI specification
type SetTaskInstanceNote struct {
	Note string `json:"note"` // The custom note to set for this Task Instance.
}

// ImportErrorCollection represents the ImportErrorCollection schema from the OpenAPI specification
type ImportErrorCollection struct {
	Import_errors []ImportError `json:"import_errors,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// TaskInstanceCollection represents the TaskInstanceCollection schema from the OpenAPI specification
type TaskInstanceCollection struct {
	Task_instances []TaskInstance `json:"task_instances,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// SchedulerStatus represents the SchedulerStatus schema from the OpenAPI specification
type SchedulerStatus struct {
	Latest_scheduler_heartbeat string `json:"latest_scheduler_heartbeat,omitempty"` // The time the scheduler last do a heartbeat.
	Status string `json:"status,omitempty"` // Health status
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Status float64 `json:"status"` // The HTTP status code generated by the API server for this occurrence of the problem.
	Title string `json:"title"` // A short, human-readable summary of the problem type.
	TypeField string `json:"type"` // A URI reference [RFC3986] that identifies the problem type. This specification encourages that, when dereferenced, it provide human-readable documentation for the problem type.
	Detail string `json:"detail,omitempty"` // A human-readable explanation specific to this occurrence of the problem.
	Instance string `json:"instance,omitempty"` // A URI reference that identifies the specific occurrence of the problem. It may or may not yield further information if dereferenced.
}

// VariableCollection represents the VariableCollection schema from the OpenAPI specification
type VariableCollection struct {
	Variables []VariableCollectionItem `json:"variables,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Name string `json:"name,omitempty"`
}

// EventLogCollection represents the EventLogCollection schema from the OpenAPI specification
type EventLogCollection struct {
	Event_logs []EventLog `json:"event_logs,omitempty"`
	Total_entries int `json:"total_entries,omitempty"` // Count of total objects in the current result set before pagination parameters (limit, offset) are applied.
}

// TaskInstanceReference represents the TaskInstanceReference schema from the OpenAPI specification
type TaskInstanceReference struct {
	Dag_run_id string `json:"dag_run_id,omitempty"` // The DAG run ID.
	Execution_date string `json:"execution_date,omitempty"`
	Task_id string `json:"task_id,omitempty"` // The task ID.
	Dag_id string `json:"dag_id,omitempty"` // The DAG ID.
}

// DAGRun represents the DAGRun schema from the OpenAPI specification
type DAGRun struct {
	Start_date string `json:"start_date,omitempty"` // The start time. The time when DAG run was actually created. *Changed in version 2.1.3*&#58; Field becomes nullable.
	Data_interval_end string `json:"data_interval_end,omitempty"`
	End_date string `json:"end_date,omitempty"`
	Dag_id string `json:"dag_id,omitempty"`
	Last_scheduling_decision string `json:"last_scheduling_decision,omitempty"`
	Conf map[string]interface{} `json:"conf,omitempty"` // JSON object describing additional configuration parameters. The value of this field can be set only when creating the object. If you try to modify the field of an existing object, the request fails with an BAD_REQUEST error.
	External_trigger bool `json:"external_trigger,omitempty"`
	Logical_date string `json:"logical_date,omitempty"` // The logical date (previously called execution date). This is the time or interval covered by this DAG run, according to the DAG definition. The value of this field can be set only when creating the object. If you try to modify the field of an existing object, the request fails with an BAD_REQUEST error. This together with DAG_ID are a unique key. *New in version 2.2.0*
	State string `json:"state,omitempty"` // DAG State. *Changed in version 2.1.3*&#58; 'queued' is added as a possible value.
	Dag_run_id string `json:"dag_run_id,omitempty"` // Run ID. The value of this field can be set only when creating the object. If you try to modify the field of an existing object, the request fails with an BAD_REQUEST error. If not provided, a value will be generated based on execution_date. If the specified dag_run_id is in use, the creation request fails with an ALREADY_EXISTS error. This together with DAG_ID are a unique key.
	Data_interval_start string `json:"data_interval_start,omitempty"`
	Execution_date string `json:"execution_date,omitempty"` // The execution date. This is the same as logical_date, kept for backwards compatibility. If both this field and logical_date are provided but with different values, the request will fail with an BAD_REQUEST error. *Changed in version 2.2.0*&#58; Field becomes nullable. *Deprecated since version 2.2.0*&#58; Use 'logical_date' instead.
	Note string `json:"note,omitempty"` // Contains manually entered notes by the user about the DagRun. *New in version 2.5.0*
	Run_type string `json:"run_type,omitempty"`
}
