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

func testCustomers(t *testing.T) {
	t.Parallel()

	query := Customers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCustomersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
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

	count, err := Customers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCustomersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Customers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Customers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCustomersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CustomerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Customers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCustomersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CustomerExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Customer exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CustomerExists to return true, but got false.")
	}
}

func testCustomersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	customerFound, err := FindCustomer(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if customerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCustomersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Customers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCustomersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Customers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCustomersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customerOne := &Customer{}
	customerTwo := &Customer{}
	if err = randomize.Struct(seed, customerOne, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}
	if err = randomize.Struct(seed, customerTwo, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = customerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = customerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Customers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCustomersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	customerOne := &Customer{}
	customerTwo := &Customer{}
	if err = randomize.Struct(seed, customerOne, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}
	if err = randomize.Struct(seed, customerTwo, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = customerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = customerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Customers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func customerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Customer) error {
	*o = Customer{}
	return nil
}

func testCustomersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Customer{}
	o := &Customer{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, customerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Customer object: %s", err)
	}

	AddCustomerHook(boil.BeforeInsertHook, customerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	customerBeforeInsertHooks = []CustomerHook{}

	AddCustomerHook(boil.AfterInsertHook, customerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	customerAfterInsertHooks = []CustomerHook{}

	AddCustomerHook(boil.AfterSelectHook, customerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	customerAfterSelectHooks = []CustomerHook{}

	AddCustomerHook(boil.BeforeUpdateHook, customerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	customerBeforeUpdateHooks = []CustomerHook{}

	AddCustomerHook(boil.AfterUpdateHook, customerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	customerAfterUpdateHooks = []CustomerHook{}

	AddCustomerHook(boil.BeforeDeleteHook, customerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	customerBeforeDeleteHooks = []CustomerHook{}

	AddCustomerHook(boil.AfterDeleteHook, customerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	customerAfterDeleteHooks = []CustomerHook{}

	AddCustomerHook(boil.BeforeUpsertHook, customerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	customerBeforeUpsertHooks = []CustomerHook{}

	AddCustomerHook(boil.AfterUpsertHook, customerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	customerAfterUpsertHooks = []CustomerHook{}
}

func testCustomersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Customers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCustomersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(customerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Customers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCustomerToManyCustomerAddresses(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Customer
	var b, c CustomerAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, customerAddressDBTypes, false, customerAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, customerAddressDBTypes, false, customerAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.CustomerID = a.ID
	c.CustomerID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.CustomerAddresses().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.CustomerID == b.CustomerID {
			bFound = true
		}
		if v.CustomerID == c.CustomerID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CustomerSlice{&a}
	if err = a.L.LoadCustomerAddresses(ctx, tx, false, (*[]*Customer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CustomerAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.CustomerAddresses = nil
	if err = a.L.LoadCustomerAddresses(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CustomerAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testCustomerToManyAddOpCustomerAddresses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Customer
	var b, c, d, e CustomerAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*CustomerAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, customerAddressDBTypes, false, strmangle.SetComplement(customerAddressPrimaryKeyColumns, customerAddressColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*CustomerAddress{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCustomerAddresses(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.CustomerID {
			t.Error("foreign key was wrong value", a.ID, first.CustomerID)
		}
		if a.ID != second.CustomerID {
			t.Error("foreign key was wrong value", a.ID, second.CustomerID)
		}

		if first.R.Customer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Customer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.CustomerAddresses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.CustomerAddresses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.CustomerAddresses().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testCustomerToOneCustomerAddressUsingDefaultAddress(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Customer
	var foreign CustomerAddress

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, customerAddressDBTypes, false, customerAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CustomerAddress struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.DefaultAddressID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.DefaultAddress().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := CustomerSlice{&local}
	if err = local.L.LoadDefaultAddress(ctx, tx, false, (*[]*Customer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.DefaultAddress == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.DefaultAddress = nil
	if err = local.L.LoadDefaultAddress(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.DefaultAddress == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCustomerToOneSetOpCustomerAddressUsingDefaultAddress(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Customer
	var b, c CustomerAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, customerAddressDBTypes, false, strmangle.SetComplement(customerAddressPrimaryKeyColumns, customerAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, customerAddressDBTypes, false, strmangle.SetComplement(customerAddressPrimaryKeyColumns, customerAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*CustomerAddress{&b, &c} {
		err = a.SetDefaultAddress(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.DefaultAddress != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.DefaultAddressCustomers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.DefaultAddressID, x.ID) {
			t.Error("foreign key was wrong value", a.DefaultAddressID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DefaultAddressID))
		reflect.Indirect(reflect.ValueOf(&a.DefaultAddressID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.DefaultAddressID, x.ID) {
			t.Error("foreign key was wrong value", a.DefaultAddressID, x.ID)
		}
	}
}

func testCustomerToOneRemoveOpCustomerAddressUsingDefaultAddress(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Customer
	var b CustomerAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, customerAddressDBTypes, false, strmangle.SetComplement(customerAddressPrimaryKeyColumns, customerAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetDefaultAddress(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveDefaultAddress(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.DefaultAddress().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.DefaultAddress != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.DefaultAddressID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.DefaultAddressCustomers) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testCustomersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
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

func testCustomersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CustomerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCustomersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Customers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	customerDBTypes = map[string]string{`ID`: `bigint`, `DefaultAddressID`: `bigint`, `AcceptsMarketing`: `tinyint`, `AcceptsMarketingUpdatedAt`: `datetime`, `CreatedAt`: `datetime`, `Email`: `varchar`, `FirstName`: `varchar`, `LastName`: `varchar`, `MarketingOptInLevel`: `varchar`, `Note`: `varchar`, `OrderCount`: `int`, `Phone`: `varchar`, `State`: `varchar`, `TaxExempt`: `tinyint`, `TotalSpent`: `decimal`, `UpdatedAt`: `datetime`, `VerifiedEmail`: `tinyint`}
	_               = bytes.MinRead
)

func testCustomersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(customerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(customerAllColumns) == len(customerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Customers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, customerDBTypes, true, customerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCustomersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(customerAllColumns) == len(customerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Customers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, customerDBTypes, true, customerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(customerAllColumns, customerPrimaryKeyColumns) {
		fields = customerAllColumns
	} else {
		fields = strmangle.SetComplement(
			customerAllColumns,
			customerPrimaryKeyColumns,
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

	slice := CustomerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCustomersUpsert(t *testing.T) {
	t.Parallel()

	if len(customerAllColumns) == len(customerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLCustomerUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Customer{}
	if err = randomize.Struct(seed, &o, customerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Customer: %s", err)
	}

	count, err := Customers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, customerDBTypes, false, customerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Customer: %s", err)
	}

	count, err = Customers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
