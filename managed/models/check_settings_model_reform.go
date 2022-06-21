// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type checkSettingsTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *checkSettingsTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("check_settings").
func (v *checkSettingsTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *checkSettingsTableType) Columns() []string {
	return []string{
		"name",
		"interval",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *checkSettingsTableType) NewStruct() reform.Struct {
	return new(CheckSettings)
}

// NewRecord makes a new record for that table.
func (v *checkSettingsTableType) NewRecord() reform.Record {
	return new(CheckSettings)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *checkSettingsTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// CheckSettingsTable represents check_settings view or table in SQL database.
var CheckSettingsTable = &checkSettingsTableType{
	s: parse.StructInfo{
		Type:    "CheckSettings",
		SQLName: "check_settings",
		Fields: []parse.FieldInfo{
			{Name: "Name", Type: "string", Column: "name"},
			{Name: "Interval", Type: "Interval", Column: "interval"},
		},
		PKFieldIndex: 0,
	},
	z: new(CheckSettings).Values(),
}

// String returns a string representation of this struct or record.
func (s CheckSettings) String() string {
	res := make([]string, 2)
	res[0] = "Name: " + reform.Inspect(s.Name, true)
	res[1] = "Interval: " + reform.Inspect(s.Interval, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *CheckSettings) Values() []interface{} {
	return []interface{}{
		s.Name,
		s.Interval,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *CheckSettings) Pointers() []interface{} {
	return []interface{}{
		&s.Name,
		&s.Interval,
	}
}

// View returns View object for that struct.
func (s *CheckSettings) View() reform.View {
	return CheckSettingsTable
}

// Table returns Table object for that record.
func (s *CheckSettings) Table() reform.Table {
	return CheckSettingsTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *CheckSettings) PKValue() interface{} {
	return s.Name
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *CheckSettings) PKPointer() interface{} {
	return &s.Name
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *CheckSettings) HasPK() bool {
	return s.Name != CheckSettingsTable.z[CheckSettingsTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.Name = pk.
func (s *CheckSettings) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = CheckSettingsTable
	_ reform.Struct = (*CheckSettings)(nil)
	_ reform.Table  = CheckSettingsTable
	_ reform.Record = (*CheckSettings)(nil)
	_ fmt.Stringer  = (*CheckSettings)(nil)
)

func init() {
	parse.AssertUpToDate(&CheckSettingsTable.s, new(CheckSettings))
}
