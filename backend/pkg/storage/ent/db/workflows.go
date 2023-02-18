// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/repository"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/workflows"
)

// Workflows is the model entity for the Workflows schema.
type Workflows struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// WorkflowID holds the value of the "workflow_id" field.
	WorkflowID uuid.UUID `json:"workflow_id,omitempty"`
	// WorkflowName holds the value of the "workflow_name" field.
	WorkflowName string `json:"workflow_name,omitempty"`
	// BadgeURL holds the value of the "badge_url" field.
	BadgeURL string `json:"badge_url,omitempty"`
	// HTMLURL holds the value of the "html_url" field.
	HTMLURL string `json:"html_url,omitempty"`
	// JobURL holds the value of the "job_url" field.
	JobURL string `json:"job_url,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkflowsQuery when eager-loading is set.
	Edges                WorkflowsEdges `json:"edges"`
	repository_workflows *uuid.UUID
}

// WorkflowsEdges holds the relations/edges for other nodes in the graph.
type WorkflowsEdges struct {
	// Workflows holds the value of the workflows edge.
	Workflows *Repository `json:"workflows,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WorkflowsOrErr returns the Workflows value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowsEdges) WorkflowsOrErr() (*Repository, error) {
	if e.loadedTypes[0] {
		if e.Workflows == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: repository.Label}
		}
		return e.Workflows, nil
	}
	return nil, &NotLoadedError{edge: "workflows"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Workflows) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case workflows.FieldID:
			values[i] = new(sql.NullInt64)
		case workflows.FieldWorkflowName, workflows.FieldBadgeURL, workflows.FieldHTMLURL, workflows.FieldJobURL, workflows.FieldState:
			values[i] = new(sql.NullString)
		case workflows.FieldWorkflowID:
			values[i] = new(uuid.UUID)
		case workflows.ForeignKeys[0]: // repository_workflows
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Workflows", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Workflows fields.
func (w *Workflows) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workflows.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case workflows.FieldWorkflowID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_id", values[i])
			} else if value != nil {
				w.WorkflowID = *value
			}
		case workflows.FieldWorkflowName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_name", values[i])
			} else if value.Valid {
				w.WorkflowName = value.String
			}
		case workflows.FieldBadgeURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field badge_url", values[i])
			} else if value.Valid {
				w.BadgeURL = value.String
			}
		case workflows.FieldHTMLURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field html_url", values[i])
			} else if value.Valid {
				w.HTMLURL = value.String
			}
		case workflows.FieldJobURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field job_url", values[i])
			} else if value.Valid {
				w.JobURL = value.String
			}
		case workflows.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				w.State = value.String
			}
		case workflows.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field repository_workflows", values[i])
			} else if value.Valid {
				w.repository_workflows = new(uuid.UUID)
				*w.repository_workflows = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryWorkflows queries the "workflows" edge of the Workflows entity.
func (w *Workflows) QueryWorkflows() *RepositoryQuery {
	return NewWorkflowsClient(w.config).QueryWorkflows(w)
}

// Update returns a builder for updating this Workflows.
// Note that you need to call Workflows.Unwrap() before calling this method if this Workflows
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Workflows) Update() *WorkflowsUpdateOne {
	return NewWorkflowsClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Workflows entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Workflows) Unwrap() *Workflows {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("db: Workflows is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Workflows) String() string {
	var builder strings.Builder
	builder.WriteString("Workflows(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("workflow_id=")
	builder.WriteString(fmt.Sprintf("%v", w.WorkflowID))
	builder.WriteString(", ")
	builder.WriteString("workflow_name=")
	builder.WriteString(w.WorkflowName)
	builder.WriteString(", ")
	builder.WriteString("badge_url=")
	builder.WriteString(w.BadgeURL)
	builder.WriteString(", ")
	builder.WriteString("html_url=")
	builder.WriteString(w.HTMLURL)
	builder.WriteString(", ")
	builder.WriteString("job_url=")
	builder.WriteString(w.JobURL)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(w.State)
	builder.WriteByte(')')
	return builder.String()
}

// WorkflowsSlice is a parsable slice of Workflows.
type WorkflowsSlice []*Workflows

func (w WorkflowsSlice) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}
