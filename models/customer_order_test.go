// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testCustomerOrders(t *testing.T) {
	t.Parallel()

	query := CustomerOrders()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCustomerOrdersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CustomerOrder{}
	if err = randomize.Struct(seed, o, customerOrderDBTypes, true, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
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

	count, err := CustomerOrders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCustomerOrdersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CustomerOrder{}
	if err = randomize.Struct(seed, o, customerOrderDBTypes, true, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := CustomerOrders().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CustomerOrders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCustomerOrdersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CustomerOrder{}
	if err = randomize.Struct(seed, o, customerOrderDBTypes, true, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CustomerOrderSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CustomerOrders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCustomerOrdersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CustomerOrder{}
	if err = randomize.Struct(seed, o, customerOrderDBTypes, true, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CustomerOrderExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if CustomerOrder exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CustomerOrderExists to return true, but got false.")
	}
}

func testCustomerOrdersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CustomerOrder{}
	if err = randomize.Struct(seed, o, customerOrderDBTypes, true, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	customerOrderFound, err := FindCustomerOrder(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if customerOrderFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCustomerOrdersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CustomerOrder{}
	if err = randomize.Struct(seed, o, customerOrderDBTypes, true, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = CustomerOrders().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCustomerOrdersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CustomerOrder{}
	if err = randomize.Struct(seed, o, customerOrderDBTypes, true, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := CustomerOrders().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCustomerOrdersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customerOrderOne := &CustomerOrder{}
	customerOrderTwo := &CustomerOrder{}
	if err = randomize.Struct(seed, customerOrderOne, customerOrderDBTypes, false, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}
	if err = randomize.Struct(seed, customerOrderTwo, customerOrderDBTypes, false, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = customerOrderOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = customerOrderTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CustomerOrders().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCustomerOrdersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	customerOrderOne := &CustomerOrder{}
	customerOrderTwo := &CustomerOrder{}
	if err = randomize.Struct(seed, customerOrderOne, customerOrderDBTypes, false, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}
	if err = randomize.Struct(seed, customerOrderTwo, customerOrderDBTypes, false, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = customerOrderOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = customerOrderTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CustomerOrders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func customerOrderBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *CustomerOrder) error {
	*o = CustomerOrder{}
	return nil
}

func customerOrderAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *CustomerOrder) error {
	*o = CustomerOrder{}
	return nil
}

func customerOrderAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *CustomerOrder) error {
	*o = CustomerOrder{}
	return nil
}

func customerOrderBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CustomerOrder) error {
	*o = CustomerOrder{}
	return nil
}

func customerOrderAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CustomerOrder) error {
	*o = CustomerOrder{}
	return nil
}

func customerOrderBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CustomerOrder) error {
	*o = CustomerOrder{}
	return nil
}

func customerOrderAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CustomerOrder) error {
	*o = CustomerOrder{}
	return nil
}

func customerOrderBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CustomerOrder) error {
	*o = CustomerOrder{}
	return nil
}

func customerOrderAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CustomerOrder) error {
	*o = CustomerOrder{}
	return nil
}

func testCustomerOrdersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &CustomerOrder{}
	o := &CustomerOrder{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, customerOrderDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CustomerOrder object: %s", err)
	}

	AddCustomerOrderHook(boil.BeforeInsertHook, customerOrderBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	customerOrderBeforeInsertHooks = []CustomerOrderHook{}

	AddCustomerOrderHook(boil.AfterInsertHook, customerOrderAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	customerOrderAfterInsertHooks = []CustomerOrderHook{}

	AddCustomerOrderHook(boil.AfterSelectHook, customerOrderAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	customerOrderAfterSelectHooks = []CustomerOrderHook{}

	AddCustomerOrderHook(boil.BeforeUpdateHook, customerOrderBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	customerOrderBeforeUpdateHooks = []CustomerOrderHook{}

	AddCustomerOrderHook(boil.AfterUpdateHook, customerOrderAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	customerOrderAfterUpdateHooks = []CustomerOrderHook{}

	AddCustomerOrderHook(boil.BeforeDeleteHook, customerOrderBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	customerOrderBeforeDeleteHooks = []CustomerOrderHook{}

	AddCustomerOrderHook(boil.AfterDeleteHook, customerOrderAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	customerOrderAfterDeleteHooks = []CustomerOrderHook{}

	AddCustomerOrderHook(boil.BeforeUpsertHook, customerOrderBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	customerOrderBeforeUpsertHooks = []CustomerOrderHook{}

	AddCustomerOrderHook(boil.AfterUpsertHook, customerOrderAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	customerOrderAfterUpsertHooks = []CustomerOrderHook{}
}

func testCustomerOrdersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CustomerOrder{}
	if err = randomize.Struct(seed, o, customerOrderDBTypes, true, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CustomerOrders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCustomerOrdersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CustomerOrder{}
	if err = randomize.Struct(seed, o, customerOrderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(customerOrderColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := CustomerOrders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCustomerOrderToManyOrderOrderLines(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CustomerOrder
	var b, c OrderLine

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerOrderDBTypes, true, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, orderLineDBTypes, false, orderLineColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, orderLineDBTypes, false, orderLineColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.OrderID = a.ID
	c.OrderID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.OrderOrderLines().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.OrderID == b.OrderID {
			bFound = true
		}
		if v.OrderID == c.OrderID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CustomerOrderSlice{&a}
	if err = a.L.LoadOrderOrderLines(ctx, tx, false, (*[]*CustomerOrder)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OrderOrderLines); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.OrderOrderLines = nil
	if err = a.L.LoadOrderOrderLines(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OrderOrderLines); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testCustomerOrderToManyAddOpOrderOrderLines(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CustomerOrder
	var b, c, d, e OrderLine

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerOrderDBTypes, false, strmangle.SetComplement(customerOrderPrimaryKeyColumns, customerOrderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*OrderLine{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, orderLineDBTypes, false, strmangle.SetComplement(orderLinePrimaryKeyColumns, orderLineColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*OrderLine{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddOrderOrderLines(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.OrderID {
			t.Error("foreign key was wrong value", a.ID, first.OrderID)
		}
		if a.ID != second.OrderID {
			t.Error("foreign key was wrong value", a.ID, second.OrderID)
		}

		if first.R.Order != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Order != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.OrderOrderLines[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.OrderOrderLines[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.OrderOrderLines().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCustomerOrdersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CustomerOrder{}
	if err = randomize.Struct(seed, o, customerOrderDBTypes, true, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
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

func testCustomerOrdersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CustomerOrder{}
	if err = randomize.Struct(seed, o, customerOrderDBTypes, true, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CustomerOrderSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCustomerOrdersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CustomerOrder{}
	if err = randomize.Struct(seed, o, customerOrderDBTypes, true, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CustomerOrders().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	customerOrderDBTypes = map[string]string{`ID`: `bigint`, `OrderNumber`: `int`, `Name`: `varchar`, `Note`: `varchar`, `Email`: `varchar`, `FinancialStatus`: `varchar`, `FulfillmentStatus`: `varchar`, `TotalDiscounts`: `decimal`, `TotalLineItemsPrice`: `decimal`, `TotalOutstanding`: `decimal`, `TotalPrice`: `decimal`, `TotalTax`: `decimal`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`, `ProcessedAt`: `datetime`, `CancelledAt`: `datetime`, `ClosedAt`: `datetime`}
	_                    = bytes.MinRead
)

func testCustomerOrdersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(customerOrderPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(customerOrderAllColumns) == len(customerOrderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CustomerOrder{}
	if err = randomize.Struct(seed, o, customerOrderDBTypes, true, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CustomerOrders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, customerOrderDBTypes, true, customerOrderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCustomerOrdersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(customerOrderAllColumns) == len(customerOrderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CustomerOrder{}
	if err = randomize.Struct(seed, o, customerOrderDBTypes, true, customerOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CustomerOrders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, customerOrderDBTypes, true, customerOrderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(customerOrderAllColumns, customerOrderPrimaryKeyColumns) {
		fields = customerOrderAllColumns
	} else {
		fields = strmangle.SetComplement(
			customerOrderAllColumns,
			customerOrderPrimaryKeyColumns,
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

	slice := CustomerOrderSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCustomerOrdersUpsert(t *testing.T) {
	t.Parallel()

	if len(customerOrderAllColumns) == len(customerOrderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLCustomerOrderUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := CustomerOrder{}
	if err = randomize.Struct(seed, &o, customerOrderDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CustomerOrder: %s", err)
	}

	count, err := CustomerOrders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, customerOrderDBTypes, false, customerOrderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CustomerOrder struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CustomerOrder: %s", err)
	}

	count, err = CustomerOrders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
