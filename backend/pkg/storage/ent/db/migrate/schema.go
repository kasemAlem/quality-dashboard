// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CodeCovsColumns holds the columns for the "code_covs" table.
	CodeCovsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "repository_name", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "git_organization", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "coverage_percentage", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "numeric"}},
		{Name: "repository_codecov", Type: field.TypeUUID, Nullable: true},
	}
	// CodeCovsTable holds the schema information for the "code_covs" table.
	CodeCovsTable = &schema.Table{
		Name:       "code_covs",
		Columns:    CodeCovsColumns,
		PrimaryKey: []*schema.Column{CodeCovsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "code_covs_repositories_codecov",
				Columns:    []*schema.Column{CodeCovsColumns[4]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProwJobsColumns holds the columns for the "prow_jobs" table.
	ProwJobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "job_id", Type: field.TypeString, Unique: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamptz"}},
		{Name: "duration", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "tests_count", Type: field.TypeInt64, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "failed_count", Type: field.TypeInt64, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "skipped_count", Type: field.TypeInt64, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "job_name", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "job_type", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "state", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "job_url", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "ci_failed", Type: field.TypeInt16, SchemaType: map[string]string{"postgres": "numeric"}},
		{Name: "repository_prow_jobs", Type: field.TypeUUID, Nullable: true},
	}
	// ProwJobsTable holds the schema information for the "prow_jobs" table.
	ProwJobsTable = &schema.Table{
		Name:       "prow_jobs",
		Columns:    ProwJobsColumns,
		PrimaryKey: []*schema.Column{ProwJobsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "prow_jobs_repositories_prow_jobs",
				Columns:    []*schema.Column{ProwJobsColumns[12]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProwSuitesColumns holds the columns for the "prow_suites" table.
	ProwSuitesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "job_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "name", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "status", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "time", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "repository_prow_suites", Type: field.TypeUUID, Nullable: true},
	}
	// ProwSuitesTable holds the schema information for the "prow_suites" table.
	ProwSuitesTable = &schema.Table{
		Name:       "prow_suites",
		Columns:    ProwSuitesColumns,
		PrimaryKey: []*schema.Column{ProwSuitesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "prow_suites_repositories_prow_suites",
				Columns:    []*schema.Column{ProwSuitesColumns[5]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RepositoriesColumns holds the columns for the "repositories" table.
	RepositoriesColumns = []*schema.Column{
		{Name: "repo_id", Type: field.TypeUUID, Unique: true},
		{Name: "repository_name", Type: field.TypeString, Unique: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "git_organization", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "description", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "git_url", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "teams_repositories", Type: field.TypeUUID, Nullable: true},
	}
	// RepositoriesTable holds the schema information for the "repositories" table.
	RepositoriesTable = &schema.Table{
		Name:       "repositories",
		Columns:    RepositoriesColumns,
		PrimaryKey: []*schema.Column{RepositoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "repositories_teams_repositories",
				Columns:    []*schema.Column{RepositoriesColumns[5]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TeamsColumns holds the columns for the "teams" table.
	TeamsColumns = []*schema.Column{
		{Name: "team_id", Type: field.TypeUUID, Unique: true},
		{Name: "team_name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Unique: true},
	}
	// TeamsTable holds the schema information for the "teams" table.
	TeamsTable = &schema.Table{
		Name:       "teams",
		Columns:    TeamsColumns,
		PrimaryKey: []*schema.Column{TeamsColumns[0]},
	}
	// WorkflowsColumns holds the columns for the "workflows" table.
	WorkflowsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "workflow_id", Type: field.TypeUUID, Unique: true},
		{Name: "workflow_name", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "badge_url", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "html_url", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "job_url", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "state", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "repository_workflows", Type: field.TypeUUID, Nullable: true},
	}
	// WorkflowsTable holds the schema information for the "workflows" table.
	WorkflowsTable = &schema.Table{
		Name:       "workflows",
		Columns:    WorkflowsColumns,
		PrimaryKey: []*schema.Column{WorkflowsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workflows_repositories_workflows",
				Columns:    []*schema.Column{WorkflowsColumns[7]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CodeCovsTable,
		ProwJobsTable,
		ProwSuitesTable,
		RepositoriesTable,
		TeamsTable,
		WorkflowsTable,
	}
)

func init() {
	CodeCovsTable.ForeignKeys[0].RefTable = RepositoriesTable
	ProwJobsTable.ForeignKeys[0].RefTable = RepositoriesTable
	ProwSuitesTable.ForeignKeys[0].RefTable = RepositoriesTable
	RepositoriesTable.ForeignKeys[0].RefTable = TeamsTable
	WorkflowsTable.ForeignKeys[0].RefTable = RepositoriesTable
}
