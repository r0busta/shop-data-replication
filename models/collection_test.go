// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testCollections(t *testing.T) {
	t.Parallel()

	query := Collections()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCollectionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Collection{}
	if err = randomize.Struct(seed, o, collectionDBTypes, true, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
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

	count, err := Collections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCollectionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Collection{}
	if err = randomize.Struct(seed, o, collectionDBTypes, true, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Collections().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Collections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCollectionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Collection{}
	if err = randomize.Struct(seed, o, collectionDBTypes, true, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CollectionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Collections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCollectionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Collection{}
	if err = randomize.Struct(seed, o, collectionDBTypes, true, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CollectionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Collection exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CollectionExists to return true, but got false.")
	}
}

func testCollectionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Collection{}
	if err = randomize.Struct(seed, o, collectionDBTypes, true, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	collectionFound, err := FindCollection(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if collectionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCollectionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Collection{}
	if err = randomize.Struct(seed, o, collectionDBTypes, true, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Collections().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCollectionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Collection{}
	if err = randomize.Struct(seed, o, collectionDBTypes, true, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Collections().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCollectionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	collectionOne := &Collection{}
	collectionTwo := &Collection{}
	if err = randomize.Struct(seed, collectionOne, collectionDBTypes, false, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}
	if err = randomize.Struct(seed, collectionTwo, collectionDBTypes, false, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = collectionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = collectionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Collections().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCollectionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	collectionOne := &Collection{}
	collectionTwo := &Collection{}
	if err = randomize.Struct(seed, collectionOne, collectionDBTypes, false, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}
	if err = randomize.Struct(seed, collectionTwo, collectionDBTypes, false, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = collectionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = collectionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Collections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func collectionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Collection) error {
	*o = Collection{}
	return nil
}

func collectionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Collection) error {
	*o = Collection{}
	return nil
}

func collectionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Collection) error {
	*o = Collection{}
	return nil
}

func collectionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Collection) error {
	*o = Collection{}
	return nil
}

func collectionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Collection) error {
	*o = Collection{}
	return nil
}

func collectionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Collection) error {
	*o = Collection{}
	return nil
}

func collectionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Collection) error {
	*o = Collection{}
	return nil
}

func collectionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Collection) error {
	*o = Collection{}
	return nil
}

func collectionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Collection) error {
	*o = Collection{}
	return nil
}

func testCollectionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Collection{}
	o := &Collection{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, collectionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Collection object: %s", err)
	}

	AddCollectionHook(boil.BeforeInsertHook, collectionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	collectionBeforeInsertHooks = []CollectionHook{}

	AddCollectionHook(boil.AfterInsertHook, collectionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	collectionAfterInsertHooks = []CollectionHook{}

	AddCollectionHook(boil.AfterSelectHook, collectionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	collectionAfterSelectHooks = []CollectionHook{}

	AddCollectionHook(boil.BeforeUpdateHook, collectionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	collectionBeforeUpdateHooks = []CollectionHook{}

	AddCollectionHook(boil.AfterUpdateHook, collectionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	collectionAfterUpdateHooks = []CollectionHook{}

	AddCollectionHook(boil.BeforeDeleteHook, collectionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	collectionBeforeDeleteHooks = []CollectionHook{}

	AddCollectionHook(boil.AfterDeleteHook, collectionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	collectionAfterDeleteHooks = []CollectionHook{}

	AddCollectionHook(boil.BeforeUpsertHook, collectionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	collectionBeforeUpsertHooks = []CollectionHook{}

	AddCollectionHook(boil.AfterUpsertHook, collectionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	collectionAfterUpsertHooks = []CollectionHook{}
}

func testCollectionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Collection{}
	if err = randomize.Struct(seed, o, collectionDBTypes, true, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Collections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCollectionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Collection{}
	if err = randomize.Struct(seed, o, collectionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(collectionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Collections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCollectionToManyProducts(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Collection
	var b, c Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, collectionDBTypes, true, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, productDBTypes, false, productColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, productDBTypes, false, productColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into `collection_product` (`collection_id`, `product_id`) values (?, ?)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into `collection_product` (`collection_id`, `product_id`) values (?, ?)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Products().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CollectionSlice{&a}
	if err = a.L.LoadProducts(ctx, tx, false, (*[]*Collection)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Products); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Products = nil
	if err = a.L.LoadProducts(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Products); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testCollectionToManyAddOpProducts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Collection
	var b, c, d, e Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, collectionDBTypes, false, strmangle.SetComplement(collectionPrimaryKeyColumns, collectionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Product{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, productDBTypes, false, strmangle.SetComplement(productPrimaryKeyColumns, productColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Product{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddProducts(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.Collections[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Collections[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Products[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Products[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Products().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCollectionToManySetOpProducts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Collection
	var b, c, d, e Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, collectionDBTypes, false, strmangle.SetComplement(collectionPrimaryKeyColumns, collectionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Product{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, productDBTypes, false, strmangle.SetComplement(productPrimaryKeyColumns, productColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetProducts(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Products().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetProducts(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Products().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.Collections) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Collections) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Collections[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Collections[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Products[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Products[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testCollectionToManyRemoveOpProducts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Collection
	var b, c, d, e Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, collectionDBTypes, false, strmangle.SetComplement(collectionPrimaryKeyColumns, collectionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Product{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, productDBTypes, false, strmangle.SetComplement(productPrimaryKeyColumns, productColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddProducts(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Products().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveProducts(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Products().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.Collections) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Collections) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Collections[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Collections[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Products) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Products[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Products[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testCollectionToOneImageUsingImage(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Collection
	var foreign Image

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, collectionDBTypes, true, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, imageDBTypes, false, imageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Image struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ImageID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Image().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := CollectionSlice{&local}
	if err = local.L.LoadImage(ctx, tx, false, (*[]*Collection)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Image == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Image = nil
	if err = local.L.LoadImage(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Image == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCollectionToOneSetOpImageUsingImage(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Collection
	var b, c Image

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, collectionDBTypes, false, strmangle.SetComplement(collectionPrimaryKeyColumns, collectionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, imageDBTypes, false, strmangle.SetComplement(imagePrimaryKeyColumns, imageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, imageDBTypes, false, strmangle.SetComplement(imagePrimaryKeyColumns, imageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Image{&b, &c} {
		err = a.SetImage(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Image != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Collections[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ImageID, x.ID) {
			t.Error("foreign key was wrong value", a.ImageID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ImageID))
		reflect.Indirect(reflect.ValueOf(&a.ImageID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ImageID, x.ID) {
			t.Error("foreign key was wrong value", a.ImageID, x.ID)
		}
	}
}

func testCollectionToOneRemoveOpImageUsingImage(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Collection
	var b Image

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, collectionDBTypes, false, strmangle.SetComplement(collectionPrimaryKeyColumns, collectionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, imageDBTypes, false, strmangle.SetComplement(imagePrimaryKeyColumns, imageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetImage(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveImage(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Image().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Image != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ImageID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Collections) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testCollectionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Collection{}
	if err = randomize.Struct(seed, o, collectionDBTypes, true, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
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

func testCollectionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Collection{}
	if err = randomize.Struct(seed, o, collectionDBTypes, true, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CollectionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCollectionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Collection{}
	if err = randomize.Struct(seed, o, collectionDBTypes, true, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Collections().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	collectionDBTypes = map[string]string{`ID`: `bigint`, `ImageID`: `bigint`, `BodyHTML`: `text`, `Handle`: `varchar`, `PublishedAt`: `datetime`, `PublishedScope`: `varchar`, `SortOrder`: `varchar`, `Title`: `varchar`, `TemplateSuffix`: `varchar`, `UpdatedAt`: `datetime`}
	_                 = bytes.MinRead
)

func testCollectionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(collectionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(collectionAllColumns) == len(collectionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Collection{}
	if err = randomize.Struct(seed, o, collectionDBTypes, true, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Collections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, collectionDBTypes, true, collectionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCollectionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(collectionAllColumns) == len(collectionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Collection{}
	if err = randomize.Struct(seed, o, collectionDBTypes, true, collectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Collections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, collectionDBTypes, true, collectionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(collectionAllColumns, collectionPrimaryKeyColumns) {
		fields = collectionAllColumns
	} else {
		fields = strmangle.SetComplement(
			collectionAllColumns,
			collectionPrimaryKeyColumns,
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

	slice := CollectionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCollectionsUpsert(t *testing.T) {
	t.Parallel()

	if len(collectionAllColumns) == len(collectionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLCollectionUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Collection{}
	if err = randomize.Struct(seed, &o, collectionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Collection: %s", err)
	}

	count, err := Collections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, collectionDBTypes, false, collectionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Collection struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Collection: %s", err)
	}

	count, err = Collections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
