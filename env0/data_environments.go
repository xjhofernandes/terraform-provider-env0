package env0

import (
	"context"

	"github.com/env0/terraform-provider-env0/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataEnvironments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataEnvironmentsRead,

		Schema: map[string]*schema.Schema{
			"include_archived_environments": {
				Type:        schema.TypeBool,
				Description: "set to 'true' to include archived environments (defaults to 'false')",
				Optional:    true,
				Default:     false,
			},
			"environments": {
				Type:        schema.TypeList,
				Description: "list of environments",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "the environment's id",
							Optional:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "name of the environment",
							Optional:    true,
						},
						"is_archived": {
							Type:        schema.TypeBool,
							Description: "set to 'true' to exclude archived environments when getting an environment by name",
							Computed:    true,
						},
						"project_id": {
							Type:        schema.TypeString,
							Description: "project id of the environment",
							Computed:    true,
							Optional:    true,
						},
						"approve_plan_automatically": {
							Type:        schema.TypeBool,
							Description: "the default require approval of the environment",
							Computed:    true,
						},
						"run_plan_on_pull_requests": {
							Type:        schema.TypeBool,
							Description: "does pr plan enable",
							Computed:    true,
						},
						"auto_deploy_on_path_changes_only": {
							Type:        schema.TypeBool,
							Description: "does continuous deployment on file changes in path enable",
							Computed:    true,
						},
						"deploy_on_push": {
							Type:        schema.TypeBool,
							Description: "does continuous deployment is enabled",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "the status of the environment",
							Computed:    true,
						},
						"deployment_id": {
							Type:        schema.TypeString,
							Description: "the id of the latest deployment",
							Computed:    true,
						},
						"template_id": {
							Type:        schema.TypeString,
							Description: "the template id the environment is to be created from",
							Computed:    true,
						},
						"template_name": {
							Type:        schema.TypeString,
							Description: "the template id the environment is to be created from",
							Computed:    true,
						},
						"template_revision": {
							Type:        schema.TypeString,
							Description: "the template id the environment is to be created from",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataEnvironmentsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var filteredEnvironments []client.Environment
	apiClient := meta.(client.ApiClientInterface)

	includeArchivedEnvironments := d.Get("include_archived_environments").(bool)

	projectId, ok := d.GetOk("project_id")
	if ok {
		environments, err := apiClient.ProjectEnvironments(projectId.(string))
		if err != nil {
			return diag.Errorf("failed to get list of projects: %v", err)
		}

		filteredEnvironments = filterProjectEnvironments(apiClient, projectId.(string), environments, includeArchivedEnvironments)
	} else {
		projects, err := apiClient.Projects()
		if err != nil {
			return diag.Errorf("failed to get list of projects: %v", err)
		}

		filteredProjectsId := []string{}
		for _, project := range projects {
			if includeArchivedEnvironments || !project.IsArchived {
				filteredProjectsId = append(filteredProjectsId, project.Id)
			}
		}

		for _, projectId := range filteredProjectsId {
			environments, err := apiClient.ProjectEnvironments(projectId)
			if err != nil {
				return diag.Errorf("failed to get list of projects: %v", err)
			}

			filteredEnvironmentsPerProjectId := filterProjectEnvironments(apiClient, projectId, environments, includeArchivedEnvironments)
			filteredEnvironments = append(filteredEnvironments, filteredEnvironmentsPerProjectId...)
		}
	}

	if err := writeResourceDataSlice(filteredEnvironments, "environments", d); err != nil {
		return diag.Errorf("schema slice resource data serialization failed: %v", err)
	}

	d.SetId("environments")

	return nil
}

func filterProjectEnvironments(apiClient client.ApiClientInterface, projectId string, environments []client.Environment, includeArchivedEnvironments bool) []client.Environment {
	filteredEnvironments := []client.Environment{}
	for _, environment := range environments {
		if includeArchivedEnvironments || (environment.IsArchived != nil && !*environment.IsArchived) {
			if environment.LatestDeploymentLog.BlueprintId != "" {
				environment.BlueprintId = environment.LatestDeploymentLog.BlueprintId
				environment.BlueprintName = environment.LatestDeploymentLog.BlueprintName
				environment.BlueprintRevision = environment.LatestDeploymentLog.BlueprintRevision
			}

			filteredEnvironments = append(filteredEnvironments, environment)
		}
	}
	return filteredEnvironments
}
