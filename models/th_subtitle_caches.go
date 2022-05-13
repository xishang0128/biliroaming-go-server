// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// THSubtitleCach is an object representing the database table.
type THSubtitleCach struct {
	EpisodeID int64      `boil:"episode_id" json:"episode_id" toml:"episode_id" yaml:"episode_id"`
	Data      types.JSON `boil:"data" json:"data" toml:"data" yaml:"data"`
	CreatedAt time.Time  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time  `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *thSubtitleCachR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L thSubtitleCachL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var THSubtitleCachColumns = struct {
	EpisodeID string
	Data      string
	CreatedAt string
	UpdatedAt string
}{
	EpisodeID: "episode_id",
	Data:      "data",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var THSubtitleCachTableColumns = struct {
	EpisodeID string
	Data      string
	CreatedAt string
	UpdatedAt string
}{
	EpisodeID: "th_subtitle_caches.episode_id",
	Data:      "th_subtitle_caches.data",
	CreatedAt: "th_subtitle_caches.created_at",
	UpdatedAt: "th_subtitle_caches.updated_at",
}

// Generated where

var THSubtitleCachWhere = struct {
	EpisodeID whereHelperint64
	Data      whereHelpertypes_JSON
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	EpisodeID: whereHelperint64{field: "\"th_subtitle_caches\".\"episode_id\""},
	Data:      whereHelpertypes_JSON{field: "\"th_subtitle_caches\".\"data\""},
	CreatedAt: whereHelpertime_Time{field: "\"th_subtitle_caches\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"th_subtitle_caches\".\"updated_at\""},
}

// THSubtitleCachRels is where relationship names are stored.
var THSubtitleCachRels = struct {
}{}

// thSubtitleCachR is where relationships are stored.
type thSubtitleCachR struct {
}

// NewStruct creates a new relationship struct
func (*thSubtitleCachR) NewStruct() *thSubtitleCachR {
	return &thSubtitleCachR{}
}

// thSubtitleCachL is where Load methods for each relationship are stored.
type thSubtitleCachL struct{}

var (
	thSubtitleCachAllColumns            = []string{"episode_id", "data", "created_at", "updated_at"}
	thSubtitleCachColumnsWithoutDefault = []string{"episode_id", "data", "created_at", "updated_at"}
	thSubtitleCachColumnsWithDefault    = []string{}
	thSubtitleCachPrimaryKeyColumns     = []string{"episode_id"}
	thSubtitleCachGeneratedColumns      = []string{}
)

type (
	// THSubtitleCachSlice is an alias for a slice of pointers to THSubtitleCach.
	// This should almost always be used instead of []THSubtitleCach.
	THSubtitleCachSlice []*THSubtitleCach
	// THSubtitleCachHook is the signature for custom THSubtitleCach hook methods
	THSubtitleCachHook func(context.Context, boil.ContextExecutor, *THSubtitleCach) error

	thSubtitleCachQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	thSubtitleCachType                 = reflect.TypeOf(&THSubtitleCach{})
	thSubtitleCachMapping              = queries.MakeStructMapping(thSubtitleCachType)
	thSubtitleCachPrimaryKeyMapping, _ = queries.BindMapping(thSubtitleCachType, thSubtitleCachMapping, thSubtitleCachPrimaryKeyColumns)
	thSubtitleCachInsertCacheMut       sync.RWMutex
	thSubtitleCachInsertCache          = make(map[string]insertCache)
	thSubtitleCachUpdateCacheMut       sync.RWMutex
	thSubtitleCachUpdateCache          = make(map[string]updateCache)
	thSubtitleCachUpsertCacheMut       sync.RWMutex
	thSubtitleCachUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var thSubtitleCachAfterSelectHooks []THSubtitleCachHook

var thSubtitleCachBeforeInsertHooks []THSubtitleCachHook
var thSubtitleCachAfterInsertHooks []THSubtitleCachHook

var thSubtitleCachBeforeUpdateHooks []THSubtitleCachHook
var thSubtitleCachAfterUpdateHooks []THSubtitleCachHook

var thSubtitleCachBeforeDeleteHooks []THSubtitleCachHook
var thSubtitleCachAfterDeleteHooks []THSubtitleCachHook

var thSubtitleCachBeforeUpsertHooks []THSubtitleCachHook
var thSubtitleCachAfterUpsertHooks []THSubtitleCachHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *THSubtitleCach) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSubtitleCachAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *THSubtitleCach) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSubtitleCachBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *THSubtitleCach) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSubtitleCachAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *THSubtitleCach) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSubtitleCachBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *THSubtitleCach) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSubtitleCachAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *THSubtitleCach) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSubtitleCachBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *THSubtitleCach) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSubtitleCachAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *THSubtitleCach) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSubtitleCachBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *THSubtitleCach) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSubtitleCachAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTHSubtitleCachHook registers your hook function for all future operations.
func AddTHSubtitleCachHook(hookPoint boil.HookPoint, thSubtitleCachHook THSubtitleCachHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		thSubtitleCachAfterSelectHooks = append(thSubtitleCachAfterSelectHooks, thSubtitleCachHook)
	case boil.BeforeInsertHook:
		thSubtitleCachBeforeInsertHooks = append(thSubtitleCachBeforeInsertHooks, thSubtitleCachHook)
	case boil.AfterInsertHook:
		thSubtitleCachAfterInsertHooks = append(thSubtitleCachAfterInsertHooks, thSubtitleCachHook)
	case boil.BeforeUpdateHook:
		thSubtitleCachBeforeUpdateHooks = append(thSubtitleCachBeforeUpdateHooks, thSubtitleCachHook)
	case boil.AfterUpdateHook:
		thSubtitleCachAfterUpdateHooks = append(thSubtitleCachAfterUpdateHooks, thSubtitleCachHook)
	case boil.BeforeDeleteHook:
		thSubtitleCachBeforeDeleteHooks = append(thSubtitleCachBeforeDeleteHooks, thSubtitleCachHook)
	case boil.AfterDeleteHook:
		thSubtitleCachAfterDeleteHooks = append(thSubtitleCachAfterDeleteHooks, thSubtitleCachHook)
	case boil.BeforeUpsertHook:
		thSubtitleCachBeforeUpsertHooks = append(thSubtitleCachBeforeUpsertHooks, thSubtitleCachHook)
	case boil.AfterUpsertHook:
		thSubtitleCachAfterUpsertHooks = append(thSubtitleCachAfterUpsertHooks, thSubtitleCachHook)
	}
}

// One returns a single thSubtitleCach record from the query.
func (q thSubtitleCachQuery) One(ctx context.Context, exec boil.ContextExecutor) (*THSubtitleCach, error) {
	o := &THSubtitleCach{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for th_subtitle_caches")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all THSubtitleCach records from the query.
func (q thSubtitleCachQuery) All(ctx context.Context, exec boil.ContextExecutor) (THSubtitleCachSlice, error) {
	var o []*THSubtitleCach

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to THSubtitleCach slice")
	}

	if len(thSubtitleCachAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all THSubtitleCach records in the query.
func (q thSubtitleCachQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count th_subtitle_caches rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q thSubtitleCachQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if th_subtitle_caches exists")
	}

	return count > 0, nil
}

// THSubtitleCaches retrieves all the records using an executor.
func THSubtitleCaches(mods ...qm.QueryMod) thSubtitleCachQuery {
	mods = append(mods, qm.From("\"th_subtitle_caches\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"th_subtitle_caches\".*"})
	}

	return thSubtitleCachQuery{q}
}

// FindTHSubtitleCach retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTHSubtitleCach(ctx context.Context, exec boil.ContextExecutor, episodeID int64, selectCols ...string) (*THSubtitleCach, error) {
	thSubtitleCachObj := &THSubtitleCach{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"th_subtitle_caches\" where \"episode_id\"=$1", sel,
	)

	q := queries.Raw(query, episodeID)

	err := q.Bind(ctx, exec, thSubtitleCachObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from th_subtitle_caches")
	}

	if err = thSubtitleCachObj.doAfterSelectHooks(ctx, exec); err != nil {
		return thSubtitleCachObj, err
	}

	return thSubtitleCachObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *THSubtitleCach) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no th_subtitle_caches provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(thSubtitleCachColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	thSubtitleCachInsertCacheMut.RLock()
	cache, cached := thSubtitleCachInsertCache[key]
	thSubtitleCachInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			thSubtitleCachAllColumns,
			thSubtitleCachColumnsWithDefault,
			thSubtitleCachColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(thSubtitleCachType, thSubtitleCachMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(thSubtitleCachType, thSubtitleCachMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"th_subtitle_caches\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"th_subtitle_caches\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into th_subtitle_caches")
	}

	if !cached {
		thSubtitleCachInsertCacheMut.Lock()
		thSubtitleCachInsertCache[key] = cache
		thSubtitleCachInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the THSubtitleCach.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *THSubtitleCach) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	thSubtitleCachUpdateCacheMut.RLock()
	cache, cached := thSubtitleCachUpdateCache[key]
	thSubtitleCachUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			thSubtitleCachAllColumns,
			thSubtitleCachPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update th_subtitle_caches, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"th_subtitle_caches\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, thSubtitleCachPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(thSubtitleCachType, thSubtitleCachMapping, append(wl, thSubtitleCachPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update th_subtitle_caches row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for th_subtitle_caches")
	}

	if !cached {
		thSubtitleCachUpdateCacheMut.Lock()
		thSubtitleCachUpdateCache[key] = cache
		thSubtitleCachUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q thSubtitleCachQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for th_subtitle_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for th_subtitle_caches")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o THSubtitleCachSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thSubtitleCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"th_subtitle_caches\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, thSubtitleCachPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in thSubtitleCach slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all thSubtitleCach")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *THSubtitleCach) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no th_subtitle_caches provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(thSubtitleCachColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	thSubtitleCachUpsertCacheMut.RLock()
	cache, cached := thSubtitleCachUpsertCache[key]
	thSubtitleCachUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			thSubtitleCachAllColumns,
			thSubtitleCachColumnsWithDefault,
			thSubtitleCachColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			thSubtitleCachAllColumns,
			thSubtitleCachPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert th_subtitle_caches, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(thSubtitleCachPrimaryKeyColumns))
			copy(conflict, thSubtitleCachPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"th_subtitle_caches\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(thSubtitleCachType, thSubtitleCachMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(thSubtitleCachType, thSubtitleCachMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert th_subtitle_caches")
	}

	if !cached {
		thSubtitleCachUpsertCacheMut.Lock()
		thSubtitleCachUpsertCache[key] = cache
		thSubtitleCachUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single THSubtitleCach record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *THSubtitleCach) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no THSubtitleCach provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), thSubtitleCachPrimaryKeyMapping)
	sql := "DELETE FROM \"th_subtitle_caches\" WHERE \"episode_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from th_subtitle_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for th_subtitle_caches")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q thSubtitleCachQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no thSubtitleCachQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from th_subtitle_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for th_subtitle_caches")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o THSubtitleCachSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(thSubtitleCachBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thSubtitleCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"th_subtitle_caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, thSubtitleCachPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from thSubtitleCach slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for th_subtitle_caches")
	}

	if len(thSubtitleCachAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *THSubtitleCach) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTHSubtitleCach(ctx, exec, o.EpisodeID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *THSubtitleCachSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := THSubtitleCachSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thSubtitleCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"th_subtitle_caches\".* FROM \"th_subtitle_caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, thSubtitleCachPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in THSubtitleCachSlice")
	}

	*o = slice

	return nil
}

// THSubtitleCachExists checks if the THSubtitleCach row exists.
func THSubtitleCachExists(ctx context.Context, exec boil.ContextExecutor, episodeID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"th_subtitle_caches\" where \"episode_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, episodeID)
	}
	row := exec.QueryRowContext(ctx, sql, episodeID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if th_subtitle_caches exists")
	}

	return exists, nil
}
