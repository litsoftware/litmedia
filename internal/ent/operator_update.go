// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/litsoftware/litmedia/internal/ent/operator"
	"github.com/litsoftware/litmedia/internal/ent/predicate"
)

// OperatorUpdate is the builder for updating Operator entities.
type OperatorUpdate struct {
	config
	hooks    []Hook
	mutation *OperatorMutation
}

// Where appends a list predicates to the OperatorUpdate builder.
func (ou *OperatorUpdate) Where(ps ...predicate.Operator) *OperatorUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *OperatorUpdate) SetUpdatedAt(t time.Time) *OperatorUpdate {
	ou.mutation.SetUpdatedAt(t)
	return ou
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ou *OperatorUpdate) ClearUpdatedAt() *OperatorUpdate {
	ou.mutation.ClearUpdatedAt()
	return ou
}

// SetDeleteAt sets the "delete_at" field.
func (ou *OperatorUpdate) SetDeleteAt(t time.Time) *OperatorUpdate {
	ou.mutation.SetDeleteAt(t)
	return ou
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ou *OperatorUpdate) SetNillableDeleteAt(t *time.Time) *OperatorUpdate {
	if t != nil {
		ou.SetDeleteAt(*t)
	}
	return ou
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (ou *OperatorUpdate) ClearDeleteAt() *OperatorUpdate {
	ou.mutation.ClearDeleteAt()
	return ou
}

// SetName sets the "name" field.
func (ou *OperatorUpdate) SetName(s string) *OperatorUpdate {
	ou.mutation.SetName(s)
	return ou
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ou *OperatorUpdate) SetNillableName(s *string) *OperatorUpdate {
	if s != nil {
		ou.SetName(*s)
	}
	return ou
}

// ClearName clears the value of the "name" field.
func (ou *OperatorUpdate) ClearName() *OperatorUpdate {
	ou.mutation.ClearName()
	return ou
}

// SetEmail sets the "email" field.
func (ou *OperatorUpdate) SetEmail(s string) *OperatorUpdate {
	ou.mutation.SetEmail(s)
	return ou
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (ou *OperatorUpdate) SetNillableEmail(s *string) *OperatorUpdate {
	if s != nil {
		ou.SetEmail(*s)
	}
	return ou
}

// ClearEmail clears the value of the "email" field.
func (ou *OperatorUpdate) ClearEmail() *OperatorUpdate {
	ou.mutation.ClearEmail()
	return ou
}

// SetPassword sets the "password" field.
func (ou *OperatorUpdate) SetPassword(s string) *OperatorUpdate {
	ou.mutation.SetPassword(s)
	return ou
}

// SetNickname sets the "nickname" field.
func (ou *OperatorUpdate) SetNickname(s string) *OperatorUpdate {
	ou.mutation.SetNickname(s)
	return ou
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (ou *OperatorUpdate) SetNillableNickname(s *string) *OperatorUpdate {
	if s != nil {
		ou.SetNickname(*s)
	}
	return ou
}

// ClearNickname clears the value of the "nickname" field.
func (ou *OperatorUpdate) ClearNickname() *OperatorUpdate {
	ou.mutation.ClearNickname()
	return ou
}

// SetPhone sets the "phone" field.
func (ou *OperatorUpdate) SetPhone(s string) *OperatorUpdate {
	ou.mutation.SetPhone(s)
	return ou
}

// SetAvatar sets the "avatar" field.
func (ou *OperatorUpdate) SetAvatar(s string) *OperatorUpdate {
	ou.mutation.SetAvatar(s)
	return ou
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (ou *OperatorUpdate) SetNillableAvatar(s *string) *OperatorUpdate {
	if s != nil {
		ou.SetAvatar(*s)
	}
	return ou
}

// ClearAvatar clears the value of the "avatar" field.
func (ou *OperatorUpdate) ClearAvatar() *OperatorUpdate {
	ou.mutation.ClearAvatar()
	return ou
}

// SetRememberToken sets the "remember_token" field.
func (ou *OperatorUpdate) SetRememberToken(s string) *OperatorUpdate {
	ou.mutation.SetRememberToken(s)
	return ou
}

// SetNillableRememberToken sets the "remember_token" field if the given value is not nil.
func (ou *OperatorUpdate) SetNillableRememberToken(s *string) *OperatorUpdate {
	if s != nil {
		ou.SetRememberToken(*s)
	}
	return ou
}

// ClearRememberToken clears the value of the "remember_token" field.
func (ou *OperatorUpdate) ClearRememberToken() *OperatorUpdate {
	ou.mutation.ClearRememberToken()
	return ou
}

// Mutation returns the OperatorMutation object of the builder.
func (ou *OperatorUpdate) Mutation() *OperatorMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OperatorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ou.defaults()
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OperatorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OperatorUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OperatorUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OperatorUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OperatorUpdate) defaults() {
	if _, ok := ou.mutation.UpdatedAt(); !ok && !ou.mutation.UpdatedAtCleared() {
		v := operator.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
}

func (ou *OperatorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   operator.Table,
			Columns: operator.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: operator.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ou.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: operator.FieldCreatedAt,
		})
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operator.FieldUpdatedAt,
		})
	}
	if ou.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: operator.FieldUpdatedAt,
		})
	}
	if value, ok := ou.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operator.FieldDeleteAt,
		})
	}
	if ou.mutation.DeleteAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: operator.FieldDeleteAt,
		})
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldName,
		})
	}
	if ou.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: operator.FieldName,
		})
	}
	if value, ok := ou.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldEmail,
		})
	}
	if ou.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: operator.FieldEmail,
		})
	}
	if value, ok := ou.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldPassword,
		})
	}
	if value, ok := ou.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldNickname,
		})
	}
	if ou.mutation.NicknameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: operator.FieldNickname,
		})
	}
	if value, ok := ou.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldPhone,
		})
	}
	if value, ok := ou.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldAvatar,
		})
	}
	if ou.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: operator.FieldAvatar,
		})
	}
	if value, ok := ou.mutation.RememberToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldRememberToken,
		})
	}
	if ou.mutation.RememberTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: operator.FieldRememberToken,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{operator.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OperatorUpdateOne is the builder for updating a single Operator entity.
type OperatorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OperatorMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *OperatorUpdateOne) SetUpdatedAt(t time.Time) *OperatorUpdateOne {
	ouo.mutation.SetUpdatedAt(t)
	return ouo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ouo *OperatorUpdateOne) ClearUpdatedAt() *OperatorUpdateOne {
	ouo.mutation.ClearUpdatedAt()
	return ouo
}

// SetDeleteAt sets the "delete_at" field.
func (ouo *OperatorUpdateOne) SetDeleteAt(t time.Time) *OperatorUpdateOne {
	ouo.mutation.SetDeleteAt(t)
	return ouo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ouo *OperatorUpdateOne) SetNillableDeleteAt(t *time.Time) *OperatorUpdateOne {
	if t != nil {
		ouo.SetDeleteAt(*t)
	}
	return ouo
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (ouo *OperatorUpdateOne) ClearDeleteAt() *OperatorUpdateOne {
	ouo.mutation.ClearDeleteAt()
	return ouo
}

// SetName sets the "name" field.
func (ouo *OperatorUpdateOne) SetName(s string) *OperatorUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ouo *OperatorUpdateOne) SetNillableName(s *string) *OperatorUpdateOne {
	if s != nil {
		ouo.SetName(*s)
	}
	return ouo
}

// ClearName clears the value of the "name" field.
func (ouo *OperatorUpdateOne) ClearName() *OperatorUpdateOne {
	ouo.mutation.ClearName()
	return ouo
}

// SetEmail sets the "email" field.
func (ouo *OperatorUpdateOne) SetEmail(s string) *OperatorUpdateOne {
	ouo.mutation.SetEmail(s)
	return ouo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (ouo *OperatorUpdateOne) SetNillableEmail(s *string) *OperatorUpdateOne {
	if s != nil {
		ouo.SetEmail(*s)
	}
	return ouo
}

// ClearEmail clears the value of the "email" field.
func (ouo *OperatorUpdateOne) ClearEmail() *OperatorUpdateOne {
	ouo.mutation.ClearEmail()
	return ouo
}

// SetPassword sets the "password" field.
func (ouo *OperatorUpdateOne) SetPassword(s string) *OperatorUpdateOne {
	ouo.mutation.SetPassword(s)
	return ouo
}

// SetNickname sets the "nickname" field.
func (ouo *OperatorUpdateOne) SetNickname(s string) *OperatorUpdateOne {
	ouo.mutation.SetNickname(s)
	return ouo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (ouo *OperatorUpdateOne) SetNillableNickname(s *string) *OperatorUpdateOne {
	if s != nil {
		ouo.SetNickname(*s)
	}
	return ouo
}

// ClearNickname clears the value of the "nickname" field.
func (ouo *OperatorUpdateOne) ClearNickname() *OperatorUpdateOne {
	ouo.mutation.ClearNickname()
	return ouo
}

// SetPhone sets the "phone" field.
func (ouo *OperatorUpdateOne) SetPhone(s string) *OperatorUpdateOne {
	ouo.mutation.SetPhone(s)
	return ouo
}

// SetAvatar sets the "avatar" field.
func (ouo *OperatorUpdateOne) SetAvatar(s string) *OperatorUpdateOne {
	ouo.mutation.SetAvatar(s)
	return ouo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (ouo *OperatorUpdateOne) SetNillableAvatar(s *string) *OperatorUpdateOne {
	if s != nil {
		ouo.SetAvatar(*s)
	}
	return ouo
}

// ClearAvatar clears the value of the "avatar" field.
func (ouo *OperatorUpdateOne) ClearAvatar() *OperatorUpdateOne {
	ouo.mutation.ClearAvatar()
	return ouo
}

// SetRememberToken sets the "remember_token" field.
func (ouo *OperatorUpdateOne) SetRememberToken(s string) *OperatorUpdateOne {
	ouo.mutation.SetRememberToken(s)
	return ouo
}

// SetNillableRememberToken sets the "remember_token" field if the given value is not nil.
func (ouo *OperatorUpdateOne) SetNillableRememberToken(s *string) *OperatorUpdateOne {
	if s != nil {
		ouo.SetRememberToken(*s)
	}
	return ouo
}

// ClearRememberToken clears the value of the "remember_token" field.
func (ouo *OperatorUpdateOne) ClearRememberToken() *OperatorUpdateOne {
	ouo.mutation.ClearRememberToken()
	return ouo
}

// Mutation returns the OperatorMutation object of the builder.
func (ouo *OperatorUpdateOne) Mutation() *OperatorMutation {
	return ouo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OperatorUpdateOne) Select(field string, fields ...string) *OperatorUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Operator entity.
func (ouo *OperatorUpdateOne) Save(ctx context.Context) (*Operator, error) {
	var (
		err  error
		node *Operator
	)
	ouo.defaults()
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OperatorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OperatorUpdateOne) SaveX(ctx context.Context) *Operator {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OperatorUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OperatorUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OperatorUpdateOne) defaults() {
	if _, ok := ouo.mutation.UpdatedAt(); !ok && !ouo.mutation.UpdatedAtCleared() {
		v := operator.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
}

func (ouo *OperatorUpdateOne) sqlSave(ctx context.Context) (_node *Operator, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   operator.Table,
			Columns: operator.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: operator.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Operator.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, operator.FieldID)
		for _, f := range fields {
			if !operator.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != operator.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ouo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: operator.FieldCreatedAt,
		})
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operator.FieldUpdatedAt,
		})
	}
	if ouo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: operator.FieldUpdatedAt,
		})
	}
	if value, ok := ouo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operator.FieldDeleteAt,
		})
	}
	if ouo.mutation.DeleteAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: operator.FieldDeleteAt,
		})
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldName,
		})
	}
	if ouo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: operator.FieldName,
		})
	}
	if value, ok := ouo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldEmail,
		})
	}
	if ouo.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: operator.FieldEmail,
		})
	}
	if value, ok := ouo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldPassword,
		})
	}
	if value, ok := ouo.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldNickname,
		})
	}
	if ouo.mutation.NicknameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: operator.FieldNickname,
		})
	}
	if value, ok := ouo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldPhone,
		})
	}
	if value, ok := ouo.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldAvatar,
		})
	}
	if ouo.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: operator.FieldAvatar,
		})
	}
	if value, ok := ouo.mutation.RememberToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldRememberToken,
		})
	}
	if ouo.mutation.RememberTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: operator.FieldRememberToken,
		})
	}
	_node = &Operator{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{operator.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
