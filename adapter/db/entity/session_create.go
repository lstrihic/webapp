// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lstrihic/webapp/adapter/db/entity/session"
	"github.com/lstrihic/webapp/adapter/db/entity/user"
)

// SessionCreate is the builder for creating a Session entity.
type SessionCreate struct {
	config
	mutation *SessionMutation
	hooks    []Hook
}

// SetToken sets the "token" field.
func (sc *SessionCreate) SetToken(s string) *SessionCreate {
	sc.mutation.SetToken(s)
	return sc
}

// SetIP sets the "ip" field.
func (sc *SessionCreate) SetIP(s string) *SessionCreate {
	sc.mutation.SetIP(s)
	return sc
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (sc *SessionCreate) SetNillableIP(s *string) *SessionCreate {
	if s != nil {
		sc.SetIP(*s)
	}
	return sc
}

// SetIsValid sets the "is_valid" field.
func (sc *SessionCreate) SetIsValid(b bool) *SessionCreate {
	sc.mutation.SetIsValid(b)
	return sc
}

// SetNillableIsValid sets the "is_valid" field if the given value is not nil.
func (sc *SessionCreate) SetNillableIsValid(b *bool) *SessionCreate {
	if b != nil {
		sc.SetIsValid(*b)
	}
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SessionCreate) SetCreatedAt(t time.Time) *SessionCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SessionCreate) SetNillableCreatedAt(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUserID sets the "user_id" field.
func (sc *SessionCreate) SetUserID(i int) *SessionCreate {
	sc.mutation.SetUserID(i)
	return sc
}

// SetUser sets the "user" edge to the User entity.
func (sc *SessionCreate) SetUser(u *User) *SessionCreate {
	return sc.SetUserID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (sc *SessionCreate) Mutation() *SessionMutation {
	return sc.mutation
}

// Save creates the Session in the database.
func (sc *SessionCreate) Save(ctx context.Context) (*Session, error) {
	var (
		err  error
		node *Session
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("entity: uninitialized hook (forgotten import entity/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Session)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SessionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SessionCreate) SaveX(ctx context.Context) *Session {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SessionCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SessionCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SessionCreate) defaults() {
	if _, ok := sc.mutation.IsValid(); !ok {
		v := session.DefaultIsValid
		sc.mutation.SetIsValid(v)
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := session.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SessionCreate) check() error {
	if _, ok := sc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`entity: missing required field "Session.token"`)}
	}
	if v, ok := sc.mutation.Token(); ok {
		if err := session.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`entity: validator failed for field "Session.token": %w`, err)}
		}
	}
	if _, ok := sc.mutation.IsValid(); !ok {
		return &ValidationError{Name: "is_valid", err: errors.New(`entity: missing required field "Session.is_valid"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`entity: missing required field "Session.created_at"`)}
	}
	if _, ok := sc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`entity: missing required field "Session.user_id"`)}
	}
	if _, ok := sc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`entity: missing required edge "Session.user"`)}
	}
	return nil
}

func (sc *SessionCreate) sqlSave(ctx context.Context) (*Session, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *SessionCreate) createSpec() (*Session, *sqlgraph.CreateSpec) {
	var (
		_node = &Session{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: session.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: session.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Token(); ok {
		_spec.SetField(session.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := sc.mutation.IP(); ok {
		_spec.SetField(session.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if value, ok := sc.mutation.IsValid(); ok {
		_spec.SetField(session.FieldIsValid, field.TypeBool, value)
		_node.IsValid = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(session.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SessionCreateBulk is the builder for creating many Session entities in bulk.
type SessionCreateBulk struct {
	config
	builders []*SessionCreate
}

// Save creates the Session entities in the database.
func (scb *SessionCreateBulk) Save(ctx context.Context) ([]*Session, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Session, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SessionMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SessionCreateBulk) SaveX(ctx context.Context) []*Session {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SessionCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SessionCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
