// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPlayURLCaches(t *testing.T) {
	t.Parallel()

	query := PlayURLCaches()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPlayURLCachesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayURLCach{}
	if err = randomize.Struct(seed, o, playURLCachDBTypes, true, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PlayURLCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPlayURLCachesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayURLCach{}
	if err = randomize.Struct(seed, o, playURLCachDBTypes, true, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PlayURLCaches().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PlayURLCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPlayURLCachesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayURLCach{}
	if err = randomize.Struct(seed, o, playURLCachDBTypes, true, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PlayURLCachSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PlayURLCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPlayURLCachesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayURLCach{}
	if err = randomize.Struct(seed, o, playURLCachDBTypes, true, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PlayURLCachExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if PlayURLCach exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PlayURLCachExists to return true, but got false.")
	}
}

func testPlayURLCachesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayURLCach{}
	if err = randomize.Struct(seed, o, playURLCachDBTypes, true, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	playURLCachFound, err := FindPlayURLCach(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if playURLCachFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPlayURLCachesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayURLCach{}
	if err = randomize.Struct(seed, o, playURLCachDBTypes, true, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PlayURLCaches().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPlayURLCachesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayURLCach{}
	if err = randomize.Struct(seed, o, playURLCachDBTypes, true, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PlayURLCaches().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPlayURLCachesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	playURLCachOne := &PlayURLCach{}
	playURLCachTwo := &PlayURLCach{}
	if err = randomize.Struct(seed, playURLCachOne, playURLCachDBTypes, false, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}
	if err = randomize.Struct(seed, playURLCachTwo, playURLCachDBTypes, false, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = playURLCachOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = playURLCachTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PlayURLCaches().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPlayURLCachesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	playURLCachOne := &PlayURLCach{}
	playURLCachTwo := &PlayURLCach{}
	if err = randomize.Struct(seed, playURLCachOne, playURLCachDBTypes, false, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}
	if err = randomize.Struct(seed, playURLCachTwo, playURLCachDBTypes, false, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = playURLCachOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = playURLCachTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PlayURLCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func playURLCachBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *PlayURLCach) error {
	*o = PlayURLCach{}
	return nil
}

func playURLCachAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *PlayURLCach) error {
	*o = PlayURLCach{}
	return nil
}

func playURLCachAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *PlayURLCach) error {
	*o = PlayURLCach{}
	return nil
}

func playURLCachBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PlayURLCach) error {
	*o = PlayURLCach{}
	return nil
}

func playURLCachAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PlayURLCach) error {
	*o = PlayURLCach{}
	return nil
}

func playURLCachBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PlayURLCach) error {
	*o = PlayURLCach{}
	return nil
}

func playURLCachAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PlayURLCach) error {
	*o = PlayURLCach{}
	return nil
}

func playURLCachBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PlayURLCach) error {
	*o = PlayURLCach{}
	return nil
}

func playURLCachAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PlayURLCach) error {
	*o = PlayURLCach{}
	return nil
}

func testPlayURLCachesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &PlayURLCach{}
	o := &PlayURLCach{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, playURLCachDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PlayURLCach object: %s", err)
	}

	AddPlayURLCachHook(boil.BeforeInsertHook, playURLCachBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	playURLCachBeforeInsertHooks = []PlayURLCachHook{}

	AddPlayURLCachHook(boil.AfterInsertHook, playURLCachAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	playURLCachAfterInsertHooks = []PlayURLCachHook{}

	AddPlayURLCachHook(boil.AfterSelectHook, playURLCachAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	playURLCachAfterSelectHooks = []PlayURLCachHook{}

	AddPlayURLCachHook(boil.BeforeUpdateHook, playURLCachBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	playURLCachBeforeUpdateHooks = []PlayURLCachHook{}

	AddPlayURLCachHook(boil.AfterUpdateHook, playURLCachAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	playURLCachAfterUpdateHooks = []PlayURLCachHook{}

	AddPlayURLCachHook(boil.BeforeDeleteHook, playURLCachBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	playURLCachBeforeDeleteHooks = []PlayURLCachHook{}

	AddPlayURLCachHook(boil.AfterDeleteHook, playURLCachAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	playURLCachAfterDeleteHooks = []PlayURLCachHook{}

	AddPlayURLCachHook(boil.BeforeUpsertHook, playURLCachBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	playURLCachBeforeUpsertHooks = []PlayURLCachHook{}

	AddPlayURLCachHook(boil.AfterUpsertHook, playURLCachAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	playURLCachAfterUpsertHooks = []PlayURLCachHook{}
}

func testPlayURLCachesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayURLCach{}
	if err = randomize.Struct(seed, o, playURLCachDBTypes, true, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PlayURLCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPlayURLCachesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayURLCach{}
	if err = randomize.Struct(seed, o, playURLCachDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(playURLCachColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PlayURLCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPlayURLCachesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayURLCach{}
	if err = randomize.Struct(seed, o, playURLCachDBTypes, true, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPlayURLCachesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayURLCach{}
	if err = randomize.Struct(seed, o, playURLCachDBTypes, true, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PlayURLCachSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPlayURLCachesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayURLCach{}
	if err = randomize.Struct(seed, o, playURLCachDBTypes, true, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PlayURLCaches().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	playURLCachDBTypes = map[string]string{`ID`: `integer`, `EpisodeID`: `bigint`, `IsVip`: `boolean`, `Area`: `smallint`, `DeviceType`: `smallint`, `FormatType`: `smallint`, `Data`: `json`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                  = bytes.MinRead
)

func testPlayURLCachesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(playURLCachPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(playURLCachAllColumns) == len(playURLCachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PlayURLCach{}
	if err = randomize.Struct(seed, o, playURLCachDBTypes, true, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PlayURLCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, playURLCachDBTypes, true, playURLCachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPlayURLCachesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(playURLCachAllColumns) == len(playURLCachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PlayURLCach{}
	if err = randomize.Struct(seed, o, playURLCachDBTypes, true, playURLCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PlayURLCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, playURLCachDBTypes, true, playURLCachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(playURLCachAllColumns, playURLCachPrimaryKeyColumns) {
		fields = playURLCachAllColumns
	} else {
		fields = strmangle.SetComplement(
			playURLCachAllColumns,
			playURLCachPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PlayURLCachSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPlayURLCachesUpsert(t *testing.T) {
	t.Parallel()

	if len(playURLCachAllColumns) == len(playURLCachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PlayURLCach{}
	if err = randomize.Struct(seed, &o, playURLCachDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PlayURLCach: %s", err)
	}

	count, err := PlayURLCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, playURLCachDBTypes, false, playURLCachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PlayURLCach struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PlayURLCach: %s", err)
	}

	count, err = PlayURLCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
