// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"go-gin/database/ent/history_answer"
	"go-gin/database/ent/predicate"
	"go-gin/database/ent/questions"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// QuestionsQuery is the builder for querying Questions entities.
type QuestionsQuery struct {
	config
	ctx                *QueryContext
	order              []questions.OrderOption
	inters             []Interceptor
	predicates         []predicate.Questions
	withHistoryAnswers *HistoryAnswerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the QuestionsQuery builder.
func (qq *QuestionsQuery) Where(ps ...predicate.Questions) *QuestionsQuery {
	qq.predicates = append(qq.predicates, ps...)
	return qq
}

// Limit the number of records to be returned by this query.
func (qq *QuestionsQuery) Limit(limit int) *QuestionsQuery {
	qq.ctx.Limit = &limit
	return qq
}

// Offset to start from.
func (qq *QuestionsQuery) Offset(offset int) *QuestionsQuery {
	qq.ctx.Offset = &offset
	return qq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (qq *QuestionsQuery) Unique(unique bool) *QuestionsQuery {
	qq.ctx.Unique = &unique
	return qq
}

// Order specifies how the records should be ordered.
func (qq *QuestionsQuery) Order(o ...questions.OrderOption) *QuestionsQuery {
	qq.order = append(qq.order, o...)
	return qq
}

// QueryHistoryAnswers chains the current query on the "history_answers" edge.
func (qq *QuestionsQuery) QueryHistoryAnswers() *HistoryAnswerQuery {
	query := (&HistoryAnswerClient{config: qq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := qq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(questions.Table, questions.FieldID, selector),
			sqlgraph.To(history_answer.Table, history_answer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, questions.HistoryAnswersTable, questions.HistoryAnswersColumn),
		)
		fromU = sqlgraph.SetNeighbors(qq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Questions entity from the query.
// Returns a *NotFoundError when no Questions was found.
func (qq *QuestionsQuery) First(ctx context.Context) (*Questions, error) {
	nodes, err := qq.Limit(1).All(setContextOp(ctx, qq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{questions.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (qq *QuestionsQuery) FirstX(ctx context.Context) *Questions {
	node, err := qq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Questions ID from the query.
// Returns a *NotFoundError when no Questions ID was found.
func (qq *QuestionsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qq.Limit(1).IDs(setContextOp(ctx, qq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{questions.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (qq *QuestionsQuery) FirstIDX(ctx context.Context) int {
	id, err := qq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Questions entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Questions entity is found.
// Returns a *NotFoundError when no Questions entities are found.
func (qq *QuestionsQuery) Only(ctx context.Context) (*Questions, error) {
	nodes, err := qq.Limit(2).All(setContextOp(ctx, qq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{questions.Label}
	default:
		return nil, &NotSingularError{questions.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (qq *QuestionsQuery) OnlyX(ctx context.Context) *Questions {
	node, err := qq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Questions ID in the query.
// Returns a *NotSingularError when more than one Questions ID is found.
// Returns a *NotFoundError when no entities are found.
func (qq *QuestionsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qq.Limit(2).IDs(setContextOp(ctx, qq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{questions.Label}
	default:
		err = &NotSingularError{questions.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (qq *QuestionsQuery) OnlyIDX(ctx context.Context) int {
	id, err := qq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of QuestionsSlice.
func (qq *QuestionsQuery) All(ctx context.Context) ([]*Questions, error) {
	ctx = setContextOp(ctx, qq.ctx, "All")
	if err := qq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Questions, *QuestionsQuery]()
	return withInterceptors[[]*Questions](ctx, qq, qr, qq.inters)
}

// AllX is like All, but panics if an error occurs.
func (qq *QuestionsQuery) AllX(ctx context.Context) []*Questions {
	nodes, err := qq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Questions IDs.
func (qq *QuestionsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if qq.ctx.Unique == nil && qq.path != nil {
		qq.Unique(true)
	}
	ctx = setContextOp(ctx, qq.ctx, "IDs")
	if err = qq.Select(questions.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (qq *QuestionsQuery) IDsX(ctx context.Context) []int {
	ids, err := qq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (qq *QuestionsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, qq.ctx, "Count")
	if err := qq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, qq, querierCount[*QuestionsQuery](), qq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (qq *QuestionsQuery) CountX(ctx context.Context) int {
	count, err := qq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (qq *QuestionsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, qq.ctx, "Exist")
	switch _, err := qq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (qq *QuestionsQuery) ExistX(ctx context.Context) bool {
	exist, err := qq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the QuestionsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (qq *QuestionsQuery) Clone() *QuestionsQuery {
	if qq == nil {
		return nil
	}
	return &QuestionsQuery{
		config:             qq.config,
		ctx:                qq.ctx.Clone(),
		order:              append([]questions.OrderOption{}, qq.order...),
		inters:             append([]Interceptor{}, qq.inters...),
		predicates:         append([]predicate.Questions{}, qq.predicates...),
		withHistoryAnswers: qq.withHistoryAnswers.Clone(),
		// clone intermediate query.
		sql:  qq.sql.Clone(),
		path: qq.path,
	}
}

// WithHistoryAnswers tells the query-builder to eager-load the nodes that are connected to
// the "history_answers" edge. The optional arguments are used to configure the query builder of the edge.
func (qq *QuestionsQuery) WithHistoryAnswers(opts ...func(*HistoryAnswerQuery)) *QuestionsQuery {
	query := (&HistoryAnswerClient{config: qq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	qq.withHistoryAnswers = query
	return qq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Questions.Query().
//		GroupBy(questions.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (qq *QuestionsQuery) GroupBy(field string, fields ...string) *QuestionsGroupBy {
	qq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &QuestionsGroupBy{build: qq}
	grbuild.flds = &qq.ctx.Fields
	grbuild.label = questions.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Questions.Query().
//		Select(questions.FieldCreatedAt).
//		Scan(ctx, &v)
func (qq *QuestionsQuery) Select(fields ...string) *QuestionsSelect {
	qq.ctx.Fields = append(qq.ctx.Fields, fields...)
	sbuild := &QuestionsSelect{QuestionsQuery: qq}
	sbuild.label = questions.Label
	sbuild.flds, sbuild.scan = &qq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a QuestionsSelect configured with the given aggregations.
func (qq *QuestionsQuery) Aggregate(fns ...AggregateFunc) *QuestionsSelect {
	return qq.Select().Aggregate(fns...)
}

func (qq *QuestionsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range qq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, qq); err != nil {
				return err
			}
		}
	}
	for _, f := range qq.ctx.Fields {
		if !questions.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if qq.path != nil {
		prev, err := qq.path(ctx)
		if err != nil {
			return err
		}
		qq.sql = prev
	}
	return nil
}

func (qq *QuestionsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Questions, error) {
	var (
		nodes       = []*Questions{}
		_spec       = qq.querySpec()
		loadedTypes = [1]bool{
			qq.withHistoryAnswers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Questions).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Questions{config: qq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, qq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := qq.withHistoryAnswers; query != nil {
		if err := qq.loadHistoryAnswers(ctx, query, nodes,
			func(n *Questions) { n.Edges.HistoryAnswers = []*History_Answer{} },
			func(n *Questions, e *History_Answer) { n.Edges.HistoryAnswers = append(n.Edges.HistoryAnswers, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (qq *QuestionsQuery) loadHistoryAnswers(ctx context.Context, query *HistoryAnswerQuery, nodes []*Questions, init func(*Questions), assign func(*Questions, *History_Answer)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Questions)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(history_answer.FieldQuestionID)
	}
	query.Where(predicate.History_Answer(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(questions.HistoryAnswersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.QuestionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "question_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (qq *QuestionsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := qq.querySpec()
	_spec.Node.Columns = qq.ctx.Fields
	if len(qq.ctx.Fields) > 0 {
		_spec.Unique = qq.ctx.Unique != nil && *qq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, qq.driver, _spec)
}

func (qq *QuestionsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(questions.Table, questions.Columns, sqlgraph.NewFieldSpec(questions.FieldID, field.TypeInt))
	_spec.From = qq.sql
	if unique := qq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if qq.path != nil {
		_spec.Unique = true
	}
	if fields := qq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, questions.FieldID)
		for i := range fields {
			if fields[i] != questions.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := qq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := qq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := qq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := qq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (qq *QuestionsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(qq.driver.Dialect())
	t1 := builder.Table(questions.Table)
	columns := qq.ctx.Fields
	if len(columns) == 0 {
		columns = questions.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if qq.sql != nil {
		selector = qq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if qq.ctx.Unique != nil && *qq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range qq.predicates {
		p(selector)
	}
	for _, p := range qq.order {
		p(selector)
	}
	if offset := qq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := qq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// QuestionsGroupBy is the group-by builder for Questions entities.
type QuestionsGroupBy struct {
	selector
	build *QuestionsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (qgb *QuestionsGroupBy) Aggregate(fns ...AggregateFunc) *QuestionsGroupBy {
	qgb.fns = append(qgb.fns, fns...)
	return qgb
}

// Scan applies the selector query and scans the result into the given value.
func (qgb *QuestionsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, qgb.build.ctx, "GroupBy")
	if err := qgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*QuestionsQuery, *QuestionsGroupBy](ctx, qgb.build, qgb, qgb.build.inters, v)
}

func (qgb *QuestionsGroupBy) sqlScan(ctx context.Context, root *QuestionsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(qgb.fns))
	for _, fn := range qgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*qgb.flds)+len(qgb.fns))
		for _, f := range *qgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*qgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// QuestionsSelect is the builder for selecting fields of Questions entities.
type QuestionsSelect struct {
	*QuestionsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (qs *QuestionsSelect) Aggregate(fns ...AggregateFunc) *QuestionsSelect {
	qs.fns = append(qs.fns, fns...)
	return qs
}

// Scan applies the selector query and scans the result into the given value.
func (qs *QuestionsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, qs.ctx, "Select")
	if err := qs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*QuestionsQuery, *QuestionsSelect](ctx, qs.QuestionsQuery, qs, qs.inters, v)
}

func (qs *QuestionsSelect) sqlScan(ctx context.Context, root *QuestionsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(qs.fns))
	for _, fn := range qs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*qs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}