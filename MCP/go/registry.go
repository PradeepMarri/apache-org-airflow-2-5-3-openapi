package main

import (
	"github.com/airflow-api-stable/mcp-server/config"
	"github.com/airflow-api-stable/mcp-server/models"
	tools_dag "github.com/airflow-api-stable/mcp-server/tools/dag"
	tools_dataset "github.com/airflow-api-stable/mcp-server/tools/dataset"
	tools_role "github.com/airflow-api-stable/mcp-server/tools/role"
	tools_xcom "github.com/airflow-api-stable/mcp-server/tools/xcom"
	tools_variable "github.com/airflow-api-stable/mcp-server/tools/variable"
	tools_taskinstance "github.com/airflow-api-stable/mcp-server/tools/taskinstance"
	tools_eventlog "github.com/airflow-api-stable/mcp-server/tools/eventlog"
	tools_connection "github.com/airflow-api-stable/mcp-server/tools/connection"
	tools_plugin "github.com/airflow-api-stable/mcp-server/tools/plugin"
	tools_user "github.com/airflow-api-stable/mcp-server/tools/user"
	tools_dagrun "github.com/airflow-api-stable/mcp-server/tools/dagrun"
	tools_config "github.com/airflow-api-stable/mcp-server/tools/config"
	tools_importerror "github.com/airflow-api-stable/mcp-server/tools/importerror"
	tools_monitoring "github.com/airflow-api-stable/mcp-server/tools/monitoring"
	tools_dagwarning "github.com/airflow-api-stable/mcp-server/tools/dagwarning"
	tools_pool "github.com/airflow-api-stable/mcp-server/tools/pool"
	tools_provider "github.com/airflow-api-stable/mcp-server/tools/provider"
	tools_permission "github.com/airflow-api-stable/mcp-server/tools/permission"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_dag.CreateGet_dag_sourceTool(cfg),
		tools_dataset.CreateGet_datasetsTool(cfg),
		tools_role.CreateDelete_roleTool(cfg),
		tools_role.CreateGet_roleTool(cfg),
		tools_role.CreatePatch_roleTool(cfg),
		tools_dag.CreatePost_set_task_instances_stateTool(cfg),
		tools_xcom.CreateGet_xcom_entryTool(cfg),
		tools_role.CreateGet_rolesTool(cfg),
		tools_role.CreatePost_roleTool(cfg),
		tools_variable.CreateGet_variablesTool(cfg),
		tools_variable.CreatePost_variablesTool(cfg),
		tools_taskinstance.CreateGet_task_instances_batchTool(cfg),
		tools_eventlog.CreateGet_event_logTool(cfg),
		tools_connection.CreateTest_connectionTool(cfg),
		tools_dag.CreatePatch_dagTool(cfg),
		tools_dag.CreateDelete_dagTool(cfg),
		tools_dag.CreateGet_dagTool(cfg),
		tools_taskinstance.CreateGet_logTool(cfg),
		tools_plugin.CreateGet_pluginsTool(cfg),
		tools_user.CreateGet_userTool(cfg),
		tools_user.CreatePatch_userTool(cfg),
		tools_user.CreateDelete_userTool(cfg),
		tools_dagrun.CreateGet_dag_runs_batchTool(cfg),
		tools_config.CreateGet_configTool(cfg),
		tools_taskinstance.CreateSet_task_instance_noteTool(cfg),
		tools_dagrun.CreatePost_dag_runTool(cfg),
		tools_dagrun.CreateGet_dag_runsTool(cfg),
		tools_dagrun.CreateSet_dag_run_noteTool(cfg),
		tools_dagrun.CreateDelete_dag_runTool(cfg),
		tools_dagrun.CreateGet_dag_runTool(cfg),
		tools_dagrun.CreateUpdate_dag_run_stateTool(cfg),
		tools_importerror.CreateGet_import_errorsTool(cfg),
		tools_dag.CreateGet_dag_detailsTool(cfg),
		tools_taskinstance.CreateGet_task_instancesTool(cfg),
		tools_dagrun.CreateClear_dag_runTool(cfg),
		tools_taskinstance.CreateGet_mapped_task_instanceTool(cfg),
		tools_taskinstance.CreatePatch_mapped_task_instanceTool(cfg),
		tools_user.CreateGet_usersTool(cfg),
		tools_user.CreatePost_userTool(cfg),
		tools_connection.CreateDelete_connectionTool(cfg),
		tools_connection.CreateGet_connectionTool(cfg),
		tools_connection.CreatePatch_connectionTool(cfg),
		tools_monitoring.CreateGet_healthTool(cfg),
		tools_dagwarning.CreateGet_dag_warningsTool(cfg),
		tools_dataset.CreateGet_datasetTool(cfg),
		tools_pool.CreateDelete_poolTool(cfg),
		tools_pool.CreateGet_poolTool(cfg),
		tools_pool.CreatePatch_poolTool(cfg),
		tools_taskinstance.CreateGet_mapped_task_instancesTool(cfg),
		tools_dag.CreateGet_tasksTool(cfg),
		tools_taskinstance.CreateGet_extra_linksTool(cfg),
		tools_taskinstance.CreateSet_mapped_task_instance_noteTool(cfg),
		tools_taskinstance.CreateGet_task_instanceTool(cfg),
		tools_taskinstance.CreatePatch_task_instanceTool(cfg),
		tools_provider.CreateGet_providersTool(cfg),
		tools_dag.CreateGet_dagsTool(cfg),
		tools_dag.CreatePatch_dagsTool(cfg),
		tools_dag.CreateGet_taskTool(cfg),
		tools_pool.CreateGet_poolsTool(cfg),
		tools_pool.CreatePost_poolTool(cfg),
		tools_variable.CreateDelete_variableTool(cfg),
		tools_variable.CreateGet_variableTool(cfg),
		tools_variable.CreatePatch_variableTool(cfg),
		tools_xcom.CreateGet_xcom_entriesTool(cfg),
		tools_permission.CreateGet_permissionsTool(cfg),
		tools_dagrun.CreateGet_upstream_dataset_eventsTool(cfg),
		tools_dataset.CreateGet_dataset_eventsTool(cfg),
		tools_monitoring.CreateGet_versionTool(cfg),
		tools_eventlog.CreateGet_event_logsTool(cfg),
		tools_connection.CreateGet_connectionsTool(cfg),
		tools_connection.CreatePost_connectionTool(cfg),
		tools_importerror.CreateGet_import_errorTool(cfg),
		tools_dag.CreatePost_clear_task_instancesTool(cfg),
	}
}
