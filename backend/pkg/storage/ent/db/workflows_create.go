// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/repository"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/workflows"
)

// WorkflowsCreate is the builder for creating a Workflows entity.
type WorkflowsCreate struct {
	config
	mutation *WorkflowsMutation
	hooks    []Hook
}

// SetWorkflowID sets the "workflow_id" field.
func (wc *WorkflowsCreate) SetWorkflowID(u uuid.UUID) *WorkflowsCreate {
	wc.mutation.SetWorkflowID(u)
	return wc
}

// SetNillableWorkflowID sets the "workflow_id" field if the given value is not nil.
func (wc *WorkflowsCreate) SetNillableWorkflowID(u *uuid.UUID) *WorkflowsCreate {
	if u != nil {
		wc.SetWorkflowID(*u)
	}
	return wc
}

// SetWorkflowName sets the "workflow_name" field.
func (wc *WorkflowsCreate) SetWorkflowName(s string) *WorkflowsCreate {
	wc.mutation.SetWorkflowName(s)
	return wc
}

// SetBadgeURL sets the "badge_url" field.
func (wc *WorkflowsCreate) SetBadgeURL(s string) *WorkflowsCreate {
	wc.mutation.SetBadgeURL(s)
	return wc
}

// SetHTMLURL sets the "html_url" field.
func (wc *WorkflowsCreate) SetHTMLURL(s string) *WorkflowsCreate {
	wc.mutation.SetHTMLURL(s)
	return wc
}

// SetJobURL sets the "job_url" field.
func (wc *WorkflowsCreate) SetJobURL(s string) *WorkflowsCreate {
	wc.mutation.SetJobURL(s)
	return wc
}

// SetState sets the "state" field.
func (wc *WorkflowsCreate) SetState(s string) *WorkflowsCreate {
	wc.mutation.SetState(s)
	return wc
}

// SetWorkflowsID sets the "workflows" edge to the Repository entity by ID.
func (wc *WorkflowsCreate) SetWorkflowsID(id uuid.UUID) *WorkflowsCreate {
	wc.mutation.SetWorkflowsID(id)
	return wc
}

// SetNillableWorkflowsID sets the "workflows" edge to the Repository entity by ID if the given value is not nil.
func (wc *WorkflowsCreate) SetNillableWorkflowsID(id *uuid.UUID) *WorkflowsCreate {
	if id != nil {
		wc = wc.SetWorkflowsID(*id)
	}
	return wc
}

// SetWorkflows sets the "workflows" edge to the Repository entity.
func (wc *WorkflowsCreate) SetWorkflows(r *Repository) *WorkflowsCreate {
	return wc.SetWorkflowsID(r.ID)
}

// Mutation returns the WorkflowsMutation object of the builder.
func (wc *WorkflowsCreate) Mutation() *WorkflowsMutation {
	return wc.mutation
}

// Save creates the Workflows in the database.
func (wc *WorkflowsCreate) Save(ctx context.Context) (*Workflows, error) {
	wc.defaults()
	return withHooks[*Workflows, WorkflowsMutation](ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorkflowsCreate) SaveX(ctx context.Context) *Workflows {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WorkflowsCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WorkflowsCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WorkflowsCreate) defaults() {
	if _, ok := wc.mutation.WorkflowID(); !ok {
		v := workflows.DefaultWorkflowID()
		wc.mutation.SetWorkflowID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkflowsCreate) check() error {
	if _, ok := wc.mutation.WorkflowID(); !ok {
		return &ValidationError{Name: "workflow_id", err: errors.New(`db: missing required field "Workflows.workflow_id"`)}
	}
	if _, ok := wc.mutation.WorkflowName(); !ok {
		return &ValidationError{Name: "workflow_name", err: errors.New(`db: missing required field "Workflows.workflow_name"`)}
	}
	if v, ok := wc.mutation.WorkflowName(); ok {
		if err := workflows.WorkflowNameValidator(v); err != nil {
			return &ValidationError{Name: "workflow_name", err: fmt.Errorf(`db: validator failed for field "Workflows.workflow_name": %w`, err)}
		}
	}
	if _, ok := wc.mutation.BadgeURL(); !ok {
		return &ValidationError{Name: "badge_url", err: errors.New(`db: missing required field "Workflows.badge_url"`)}
	}
	if _, ok := wc.mutation.HTMLURL(); !ok {
		return &ValidationError{Name: "html_url", err: errors.New(`db: missing required field "Workflows.html_url"`)}
	}
	if _, ok := wc.mutation.JobURL(); !ok {
		return &ValidationError{Name: "job_url", err: errors.New(`db: missing required field "Workflows.job_url"`)}
	}
	if _, ok := wc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`db: missing required field "Workflows.state"`)}
	}
	return nil
}

func (wc *WorkflowsCreate) sqlSave(ctx context.Context) (*Workflows, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WorkflowsCreate) createSpec() (*Workflows, *sqlgraph.CreateSpec) {
	var (
		_node = &Workflows{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: workflows.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workflows.FieldID,
			},
		}
	)
	if value, ok := wc.mutation.WorkflowID(); ok {
		_spec.SetField(workflows.FieldWorkflowID, field.TypeUUID, value)
		_node.WorkflowID = value
	}
	if value, ok := wc.mutation.WorkflowName(); ok {
		_spec.SetField(workflows.FieldWorkflowName, field.TypeString, value)
		_node.WorkflowName = value
	}
	if value, ok := wc.mutation.BadgeURL(); ok {
		_spec.SetField(workflows.FieldBadgeURL, field.TypeString, value)
		_node.BadgeURL = value
	}
	if value, ok := wc.mutation.HTMLURL(); ok {
		_spec.SetField(workflows.FieldHTMLURL, field.TypeString, value)
		_node.HTMLURL = value
	}
	if value, ok := wc.mutation.JobURL(); ok {
		_spec.SetField(workflows.FieldJobURL, field.TypeString, value)
		_node.JobURL = value
	}
	if value, ok := wc.mutation.State(); ok {
		_spec.SetField(workflows.FieldState, field.TypeString, value)
		_node.State = value
	}
	if nodes := wc.mutation.WorkflowsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflows.WorkflowsTable,
			Columns: []string{workflows.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.repository_workflows = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkflowsCreateBulk is the builder for creating many Workflows entities in bulk.
type WorkflowsCreateBulk struct {
	config
	builders []*WorkflowsCreate
}

// Save creates the Workflows entities in the database.
func (wcb *WorkflowsCreateBulk) Save(ctx context.Context) ([]*Workflows, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Workflows, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkflowsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WorkflowsCreateBulk) SaveX(ctx context.Context) []*Workflows {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WorkflowsCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WorkflowsCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
