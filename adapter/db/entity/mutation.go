// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/lstrihic/webapp/adapter/db/entity/predicate"
	"github.com/lstrihic/webapp/adapter/db/entity/session"
	"github.com/lstrihic/webapp/adapter/db/entity/user"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeSession = "Session"
	TypeUser    = "User"
)

// SessionMutation represents an operation that mutates the Session nodes in the graph.
type SessionMutation struct {
	config
	op            Op
	typ           string
	id            *int
	token         *string
	ip            *string
	is_valid      *bool
	created_at    *time.Time
	clearedFields map[string]struct{}
	user          *int
	cleareduser   bool
	done          bool
	oldValue      func(context.Context) (*Session, error)
	predicates    []predicate.Session
}

var _ ent.Mutation = (*SessionMutation)(nil)

// sessionOption allows management of the mutation configuration using functional options.
type sessionOption func(*SessionMutation)

// newSessionMutation creates new mutation for the Session entity.
func newSessionMutation(c config, op Op, opts ...sessionOption) *SessionMutation {
	m := &SessionMutation{
		config:        c,
		op:            op,
		typ:           TypeSession,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSessionID sets the ID field of the mutation.
func withSessionID(id int) sessionOption {
	return func(m *SessionMutation) {
		var (
			err   error
			once  sync.Once
			value *Session
		)
		m.oldValue = func(ctx context.Context) (*Session, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Session.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSession sets the old Session of the mutation.
func withSession(node *Session) sessionOption {
	return func(m *SessionMutation) {
		m.oldValue = func(context.Context) (*Session, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SessionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SessionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("entity: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SessionMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *SessionMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Session.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetToken sets the "token" field.
func (m *SessionMutation) SetToken(s string) {
	m.token = &s
}

// Token returns the value of the "token" field in the mutation.
func (m *SessionMutation) Token() (r string, exists bool) {
	v := m.token
	if v == nil {
		return
	}
	return *v, true
}

// OldToken returns the old "token" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldToken(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldToken is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldToken: %w", err)
	}
	return oldValue.Token, nil
}

// ResetToken resets all changes to the "token" field.
func (m *SessionMutation) ResetToken() {
	m.token = nil
}

// SetIP sets the "ip" field.
func (m *SessionMutation) SetIP(s string) {
	m.ip = &s
}

// IP returns the value of the "ip" field in the mutation.
func (m *SessionMutation) IP() (r string, exists bool) {
	v := m.ip
	if v == nil {
		return
	}
	return *v, true
}

// OldIP returns the old "ip" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldIP(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIP is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIP requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIP: %w", err)
	}
	return oldValue.IP, nil
}

// ClearIP clears the value of the "ip" field.
func (m *SessionMutation) ClearIP() {
	m.ip = nil
	m.clearedFields[session.FieldIP] = struct{}{}
}

// IPCleared returns if the "ip" field was cleared in this mutation.
func (m *SessionMutation) IPCleared() bool {
	_, ok := m.clearedFields[session.FieldIP]
	return ok
}

// ResetIP resets all changes to the "ip" field.
func (m *SessionMutation) ResetIP() {
	m.ip = nil
	delete(m.clearedFields, session.FieldIP)
}

// SetIsValid sets the "is_valid" field.
func (m *SessionMutation) SetIsValid(b bool) {
	m.is_valid = &b
}

// IsValid returns the value of the "is_valid" field in the mutation.
func (m *SessionMutation) IsValid() (r bool, exists bool) {
	v := m.is_valid
	if v == nil {
		return
	}
	return *v, true
}

// OldIsValid returns the old "is_valid" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldIsValid(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsValid is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsValid requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsValid: %w", err)
	}
	return oldValue.IsValid, nil
}

// ResetIsValid resets all changes to the "is_valid" field.
func (m *SessionMutation) ResetIsValid() {
	m.is_valid = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *SessionMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *SessionMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *SessionMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUserID sets the "user_id" field.
func (m *SessionMutation) SetUserID(i int) {
	m.user = &i
}

// UserID returns the value of the "user_id" field in the mutation.
func (m *SessionMutation) UserID() (r int, exists bool) {
	v := m.user
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "user_id" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldUserID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// ResetUserID resets all changes to the "user_id" field.
func (m *SessionMutation) ResetUserID() {
	m.user = nil
}

// ClearUser clears the "user" edge to the User entity.
func (m *SessionMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared reports if the "user" edge to the User entity was cleared.
func (m *SessionMutation) UserCleared() bool {
	return m.cleareduser
}

// UserIDs returns the "user" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// UserID instead. It exists only for internal usage by the builders.
func (m *SessionMutation) UserIDs() (ids []int) {
	if id := m.user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUser resets all changes to the "user" edge.
func (m *SessionMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
}

// Where appends a list predicates to the SessionMutation builder.
func (m *SessionMutation) Where(ps ...predicate.Session) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *SessionMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Session).
func (m *SessionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SessionMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.token != nil {
		fields = append(fields, session.FieldToken)
	}
	if m.ip != nil {
		fields = append(fields, session.FieldIP)
	}
	if m.is_valid != nil {
		fields = append(fields, session.FieldIsValid)
	}
	if m.created_at != nil {
		fields = append(fields, session.FieldCreatedAt)
	}
	if m.user != nil {
		fields = append(fields, session.FieldUserID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SessionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case session.FieldToken:
		return m.Token()
	case session.FieldIP:
		return m.IP()
	case session.FieldIsValid:
		return m.IsValid()
	case session.FieldCreatedAt:
		return m.CreatedAt()
	case session.FieldUserID:
		return m.UserID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SessionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case session.FieldToken:
		return m.OldToken(ctx)
	case session.FieldIP:
		return m.OldIP(ctx)
	case session.FieldIsValid:
		return m.OldIsValid(ctx)
	case session.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case session.FieldUserID:
		return m.OldUserID(ctx)
	}
	return nil, fmt.Errorf("unknown Session field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SessionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case session.FieldToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetToken(v)
		return nil
	case session.FieldIP:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIP(v)
		return nil
	case session.FieldIsValid:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsValid(v)
		return nil
	case session.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case session.FieldUserID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	}
	return fmt.Errorf("unknown Session field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SessionMutation) AddedFields() []string {
	var fields []string
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SessionMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SessionMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Session numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SessionMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(session.FieldIP) {
		fields = append(fields, session.FieldIP)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SessionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SessionMutation) ClearField(name string) error {
	switch name {
	case session.FieldIP:
		m.ClearIP()
		return nil
	}
	return fmt.Errorf("unknown Session nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SessionMutation) ResetField(name string) error {
	switch name {
	case session.FieldToken:
		m.ResetToken()
		return nil
	case session.FieldIP:
		m.ResetIP()
		return nil
	case session.FieldIsValid:
		m.ResetIsValid()
		return nil
	case session.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case session.FieldUserID:
		m.ResetUserID()
		return nil
	}
	return fmt.Errorf("unknown Session field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SessionMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, session.EdgeUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SessionMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case session.EdgeUser:
		if id := m.user; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SessionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SessionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SessionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, session.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SessionMutation) EdgeCleared(name string) bool {
	switch name {
	case session.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SessionMutation) ClearEdge(name string) error {
	switch name {
	case session.EdgeUser:
		m.ClearUser()
		return nil
	}
	return fmt.Errorf("unknown Session unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SessionMutation) ResetEdge(name string) error {
	switch name {
	case session.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown Session edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op              Op
	typ             string
	id              *int
	email           *string
	username        *string
	password        *string
	token_key       *string
	is_banned       *bool
	clearedFields   map[string]struct{}
	sessions        map[int]struct{}
	removedsessions map[int]struct{}
	clearedsessions bool
	done            bool
	oldValue        func(context.Context) (*User, error)
	predicates      []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("entity: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetUsername sets the "username" field.
func (m *UserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *UserMutation) ResetUsername() {
	m.username = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetTokenKey sets the "token_key" field.
func (m *UserMutation) SetTokenKey(s string) {
	m.token_key = &s
}

// TokenKey returns the value of the "token_key" field in the mutation.
func (m *UserMutation) TokenKey() (r string, exists bool) {
	v := m.token_key
	if v == nil {
		return
	}
	return *v, true
}

// OldTokenKey returns the old "token_key" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldTokenKey(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTokenKey is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTokenKey requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTokenKey: %w", err)
	}
	return oldValue.TokenKey, nil
}

// ClearTokenKey clears the value of the "token_key" field.
func (m *UserMutation) ClearTokenKey() {
	m.token_key = nil
	m.clearedFields[user.FieldTokenKey] = struct{}{}
}

// TokenKeyCleared returns if the "token_key" field was cleared in this mutation.
func (m *UserMutation) TokenKeyCleared() bool {
	_, ok := m.clearedFields[user.FieldTokenKey]
	return ok
}

// ResetTokenKey resets all changes to the "token_key" field.
func (m *UserMutation) ResetTokenKey() {
	m.token_key = nil
	delete(m.clearedFields, user.FieldTokenKey)
}

// SetIsBanned sets the "is_banned" field.
func (m *UserMutation) SetIsBanned(b bool) {
	m.is_banned = &b
}

// IsBanned returns the value of the "is_banned" field in the mutation.
func (m *UserMutation) IsBanned() (r bool, exists bool) {
	v := m.is_banned
	if v == nil {
		return
	}
	return *v, true
}

// OldIsBanned returns the old "is_banned" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldIsBanned(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsBanned is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsBanned requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsBanned: %w", err)
	}
	return oldValue.IsBanned, nil
}

// ResetIsBanned resets all changes to the "is_banned" field.
func (m *UserMutation) ResetIsBanned() {
	m.is_banned = nil
}

// AddSessionIDs adds the "sessions" edge to the Session entity by ids.
func (m *UserMutation) AddSessionIDs(ids ...int) {
	if m.sessions == nil {
		m.sessions = make(map[int]struct{})
	}
	for i := range ids {
		m.sessions[ids[i]] = struct{}{}
	}
}

// ClearSessions clears the "sessions" edge to the Session entity.
func (m *UserMutation) ClearSessions() {
	m.clearedsessions = true
}

// SessionsCleared reports if the "sessions" edge to the Session entity was cleared.
func (m *UserMutation) SessionsCleared() bool {
	return m.clearedsessions
}

// RemoveSessionIDs removes the "sessions" edge to the Session entity by IDs.
func (m *UserMutation) RemoveSessionIDs(ids ...int) {
	if m.removedsessions == nil {
		m.removedsessions = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.sessions, ids[i])
		m.removedsessions[ids[i]] = struct{}{}
	}
}

// RemovedSessions returns the removed IDs of the "sessions" edge to the Session entity.
func (m *UserMutation) RemovedSessionsIDs() (ids []int) {
	for id := range m.removedsessions {
		ids = append(ids, id)
	}
	return
}

// SessionsIDs returns the "sessions" edge IDs in the mutation.
func (m *UserMutation) SessionsIDs() (ids []int) {
	for id := range m.sessions {
		ids = append(ids, id)
	}
	return
}

// ResetSessions resets all changes to the "sessions" edge.
func (m *UserMutation) ResetSessions() {
	m.sessions = nil
	m.clearedsessions = false
	m.removedsessions = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.username != nil {
		fields = append(fields, user.FieldUsername)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.token_key != nil {
		fields = append(fields, user.FieldTokenKey)
	}
	if m.is_banned != nil {
		fields = append(fields, user.FieldIsBanned)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldEmail:
		return m.Email()
	case user.FieldUsername:
		return m.Username()
	case user.FieldPassword:
		return m.Password()
	case user.FieldTokenKey:
		return m.TokenKey()
	case user.FieldIsBanned:
		return m.IsBanned()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldUsername:
		return m.OldUsername(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldTokenKey:
		return m.OldTokenKey(ctx)
	case user.FieldIsBanned:
		return m.OldIsBanned(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldTokenKey:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTokenKey(v)
		return nil
	case user.FieldIsBanned:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsBanned(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldTokenKey) {
		fields = append(fields, user.FieldTokenKey)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldTokenKey:
		m.ClearTokenKey()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldUsername:
		m.ResetUsername()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldTokenKey:
		m.ResetTokenKey()
		return nil
	case user.FieldIsBanned:
		m.ResetIsBanned()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.sessions != nil {
		edges = append(edges, user.EdgeSessions)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeSessions:
		ids := make([]ent.Value, 0, len(m.sessions))
		for id := range m.sessions {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedsessions != nil {
		edges = append(edges, user.EdgeSessions)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeSessions:
		ids := make([]ent.Value, 0, len(m.removedsessions))
		for id := range m.removedsessions {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedsessions {
		edges = append(edges, user.EdgeSessions)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeSessions:
		return m.clearedsessions
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeSessions:
		m.ResetSessions()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
