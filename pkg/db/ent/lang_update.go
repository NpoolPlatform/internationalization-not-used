// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/lang"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/predicate"
)

// LangUpdate is the builder for updating Lang entities.
type LangUpdate struct {
	config
	hooks    []Hook
	mutation *LangMutation
}

// Where appends a list predicates to the LangUpdate builder.
func (lu *LangUpdate) Where(ps ...predicate.Lang) *LangUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetLang sets the "lang" field.
func (lu *LangUpdate) SetLang(s string) *LangUpdate {
	lu.mutation.SetLang(s)
	return lu
}

// SetLogo sets the "logo" field.
func (lu *LangUpdate) SetLogo(s string) *LangUpdate {
	lu.mutation.SetLogo(s)
	return lu
}

// SetName sets the "name" field.
func (lu *LangUpdate) SetName(s string) *LangUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetShort sets the "short" field.
func (lu *LangUpdate) SetShort(s string) *LangUpdate {
	lu.mutation.SetShort(s)
	return lu
}

// SetCreateAt sets the "create_at" field.
func (lu *LangUpdate) SetCreateAt(u uint32) *LangUpdate {
	lu.mutation.ResetCreateAt()
	lu.mutation.SetCreateAt(u)
	return lu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (lu *LangUpdate) SetNillableCreateAt(u *uint32) *LangUpdate {
	if u != nil {
		lu.SetCreateAt(*u)
	}
	return lu
}

// AddCreateAt adds u to the "create_at" field.
func (lu *LangUpdate) AddCreateAt(u int32) *LangUpdate {
	lu.mutation.AddCreateAt(u)
	return lu
}

// SetUpdateAt sets the "update_at" field.
func (lu *LangUpdate) SetUpdateAt(u uint32) *LangUpdate {
	lu.mutation.ResetUpdateAt()
	lu.mutation.SetUpdateAt(u)
	return lu
}

// AddUpdateAt adds u to the "update_at" field.
func (lu *LangUpdate) AddUpdateAt(u int32) *LangUpdate {
	lu.mutation.AddUpdateAt(u)
	return lu
}

// SetDeleteAt sets the "delete_at" field.
func (lu *LangUpdate) SetDeleteAt(u uint32) *LangUpdate {
	lu.mutation.ResetDeleteAt()
	lu.mutation.SetDeleteAt(u)
	return lu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (lu *LangUpdate) SetNillableDeleteAt(u *uint32) *LangUpdate {
	if u != nil {
		lu.SetDeleteAt(*u)
	}
	return lu
}

// AddDeleteAt adds u to the "delete_at" field.
func (lu *LangUpdate) AddDeleteAt(u int32) *LangUpdate {
	lu.mutation.AddDeleteAt(u)
	return lu
}

// Mutation returns the LangMutation object of the builder.
func (lu *LangUpdate) Mutation() *LangMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LangUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	lu.defaults()
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LangMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LangUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LangUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LangUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LangUpdate) defaults() {
	if _, ok := lu.mutation.UpdateAt(); !ok {
		v := lang.UpdateDefaultUpdateAt()
		lu.mutation.SetUpdateAt(v)
	}
}

func (lu *LangUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lang.Table,
			Columns: lang.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: lang.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Lang(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lang.FieldLang,
		})
	}
	if value, ok := lu.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lang.FieldLogo,
		})
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lang.FieldName,
		})
	}
	if value, ok := lu.mutation.Short(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lang.FieldShort,
		})
	}
	if value, ok := lu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldCreateAt,
		})
	}
	if value, ok := lu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldCreateAt,
		})
	}
	if value, ok := lu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldUpdateAt,
		})
	}
	if value, ok := lu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldUpdateAt,
		})
	}
	if value, ok := lu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldDeleteAt,
		})
	}
	if value, ok := lu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lang.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// LangUpdateOne is the builder for updating a single Lang entity.
type LangUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LangMutation
}

// SetLang sets the "lang" field.
func (luo *LangUpdateOne) SetLang(s string) *LangUpdateOne {
	luo.mutation.SetLang(s)
	return luo
}

// SetLogo sets the "logo" field.
func (luo *LangUpdateOne) SetLogo(s string) *LangUpdateOne {
	luo.mutation.SetLogo(s)
	return luo
}

// SetName sets the "name" field.
func (luo *LangUpdateOne) SetName(s string) *LangUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetShort sets the "short" field.
func (luo *LangUpdateOne) SetShort(s string) *LangUpdateOne {
	luo.mutation.SetShort(s)
	return luo
}

// SetCreateAt sets the "create_at" field.
func (luo *LangUpdateOne) SetCreateAt(u uint32) *LangUpdateOne {
	luo.mutation.ResetCreateAt()
	luo.mutation.SetCreateAt(u)
	return luo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (luo *LangUpdateOne) SetNillableCreateAt(u *uint32) *LangUpdateOne {
	if u != nil {
		luo.SetCreateAt(*u)
	}
	return luo
}

// AddCreateAt adds u to the "create_at" field.
func (luo *LangUpdateOne) AddCreateAt(u int32) *LangUpdateOne {
	luo.mutation.AddCreateAt(u)
	return luo
}

// SetUpdateAt sets the "update_at" field.
func (luo *LangUpdateOne) SetUpdateAt(u uint32) *LangUpdateOne {
	luo.mutation.ResetUpdateAt()
	luo.mutation.SetUpdateAt(u)
	return luo
}

// AddUpdateAt adds u to the "update_at" field.
func (luo *LangUpdateOne) AddUpdateAt(u int32) *LangUpdateOne {
	luo.mutation.AddUpdateAt(u)
	return luo
}

// SetDeleteAt sets the "delete_at" field.
func (luo *LangUpdateOne) SetDeleteAt(u uint32) *LangUpdateOne {
	luo.mutation.ResetDeleteAt()
	luo.mutation.SetDeleteAt(u)
	return luo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (luo *LangUpdateOne) SetNillableDeleteAt(u *uint32) *LangUpdateOne {
	if u != nil {
		luo.SetDeleteAt(*u)
	}
	return luo
}

// AddDeleteAt adds u to the "delete_at" field.
func (luo *LangUpdateOne) AddDeleteAt(u int32) *LangUpdateOne {
	luo.mutation.AddDeleteAt(u)
	return luo
}

// Mutation returns the LangMutation object of the builder.
func (luo *LangUpdateOne) Mutation() *LangMutation {
	return luo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LangUpdateOne) Select(field string, fields ...string) *LangUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Lang entity.
func (luo *LangUpdateOne) Save(ctx context.Context) (*Lang, error) {
	var (
		err  error
		node *Lang
	)
	luo.defaults()
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LangMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LangUpdateOne) SaveX(ctx context.Context) *Lang {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LangUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LangUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LangUpdateOne) defaults() {
	if _, ok := luo.mutation.UpdateAt(); !ok {
		v := lang.UpdateDefaultUpdateAt()
		luo.mutation.SetUpdateAt(v)
	}
}

func (luo *LangUpdateOne) sqlSave(ctx context.Context) (_node *Lang, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lang.Table,
			Columns: lang.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: lang.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Lang.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lang.FieldID)
		for _, f := range fields {
			if !lang.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lang.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Lang(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lang.FieldLang,
		})
	}
	if value, ok := luo.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lang.FieldLogo,
		})
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lang.FieldName,
		})
	}
	if value, ok := luo.mutation.Short(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lang.FieldShort,
		})
	}
	if value, ok := luo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldCreateAt,
		})
	}
	if value, ok := luo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldCreateAt,
		})
	}
	if value, ok := luo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldUpdateAt,
		})
	}
	if value, ok := luo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldUpdateAt,
		})
	}
	if value, ok := luo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldDeleteAt,
		})
	}
	if value, ok := luo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldDeleteAt,
		})
	}
	_node = &Lang{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lang.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
