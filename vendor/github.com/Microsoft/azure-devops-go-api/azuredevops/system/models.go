// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package system

type SqlDbType string

type sqlDbTypeValuesType struct {
	BigInt           SqlDbType
	Binary           SqlDbType
	Bit              SqlDbType
	Char             SqlDbType
	DateTime         SqlDbType
	Decimal          SqlDbType
	Float            SqlDbType
	Image            SqlDbType
	Int              SqlDbType
	Money            SqlDbType
	NChar            SqlDbType
	NText            SqlDbType
	NVarChar         SqlDbType
	Real             SqlDbType
	UniqueIdentifier SqlDbType
	SmallDateTime    SqlDbType
	SmallInt         SqlDbType
	SmallMoney       SqlDbType
	Text             SqlDbType
	Timestamp        SqlDbType
	TinyInt          SqlDbType
	VarBinary        SqlDbType
	VarChar          SqlDbType
	Variant          SqlDbType
	Xml              SqlDbType
	Udt              SqlDbType
	Structured       SqlDbType
	Date             SqlDbType
	Time             SqlDbType
	DateTime2        SqlDbType
	DateTimeOffset   SqlDbType
}

var SqlDbTypeValues = sqlDbTypeValuesType{
	BigInt:           "bigInt",
	Binary:           "binary",
	Bit:              "bit",
	Char:             "char",
	DateTime:         "dateTime",
	Decimal:          "decimal",
	Float:            "float",
	Image:            "image",
	Int:              "int",
	Money:            "money",
	NChar:            "nChar",
	NText:            "nText",
	NVarChar:         "nVarChar",
	Real:             "real",
	UniqueIdentifier: "uniqueIdentifier",
	SmallDateTime:    "smallDateTime",
	SmallInt:         "smallInt",
	SmallMoney:       "smallMoney",
	Text:             "text",
	Timestamp:        "timestamp",
	TinyInt:          "tinyInt",
	VarBinary:        "varBinary",
	VarChar:          "varChar",
	Variant:          "variant",
	Xml:              "xml",
	Udt:              "udt",
	Structured:       "structured",
	Date:             "date",
	Time:             "time",
	DateTime2:        "dateTime2",
	DateTimeOffset:   "dateTimeOffset",
}

type TraceLevel string

type traceLevelValuesType struct {
	Off     TraceLevel
	Error   TraceLevel
	Warning TraceLevel
	Info    TraceLevel
	Verbose TraceLevel
}

var TraceLevelValues = traceLevelValuesType{
	Off:     "off",
	Error:   "error",
	Warning: "warning",
	Info:    "info",
	Verbose: "verbose",
}
