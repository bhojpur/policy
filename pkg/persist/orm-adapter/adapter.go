package ormadapter

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"errors"
	"log"
	"runtime"
	"strings"

	orm "github.com/bhojpur/dbm/pkg/orm"
	"github.com/bhojpur/policy/pkg/model"
	"github.com/bhojpur/policy/pkg/persist"
	"github.com/lib/pq"
)

// TableName  if tableName=="" , adapter will use default tablename "bhojpur_rule".
func (the *BhojpurRule) TableName() string {
	if len(the.tableName) == 0 {
		return "bhojpur_rule"
	}
	return the.tableName
}

// BhojpurRule  .
type BhojpurRule struct {
	PType string `orm:"varchar(100) index not null default ''"`
	V0    string `orm:"varchar(100) index not null default ''"`
	V1    string `orm:"varchar(100) index not null default ''"`
	V2    string `orm:"varchar(100) index not null default ''"`
	V3    string `orm:"varchar(100) index not null default ''"`
	V4    string `orm:"varchar(100) index not null default ''"`
	V5    string `orm:"varchar(100) index not null default ''"`

	tableName string `orm:"-"`
}

// Adapter represents the ORM adapter for policy storage.
type Adapter struct {
	driverName     string
	dataSourceName string
	dbSpecified    bool
	isFiltered     bool
	engine         *orm.Engine
	tablePrefix    string
	tableName      string
}

// Filter  .
type Filter struct {
	PType []string
	V0    []string
	V1    []string
	V2    []string
	V3    []string
	V4    []string
	V5    []string
}

// finalizer is the destructor for Adapter.
func finalizer(a *Adapter) {
	if a.engine == nil {
		return
	}

	err := a.engine.Close()
	if err != nil {
		log.Printf("close orm adapter engine failed, err: %v", err)
	}
}

// NewAdapter is the constructor for Adapter.
// dbSpecified is an optional bool parameter. The default value is false.
// It's up to whether you have specified an existing DB in dataSourceName.
// If dbSpecified == true, you need to make sure the DB in dataSourceName exists.
// If dbSpecified == false, the adapter will automatically create a DB named "bhojpur".
func NewAdapter(driverName string, dataSourceName string, dbSpecified ...bool) (*Adapter, error) {
	a := &Adapter{
		driverName:     driverName,
		dataSourceName: dataSourceName,
	}

	if len(dbSpecified) == 0 {
		a.dbSpecified = false
	} else if len(dbSpecified) == 1 {
		a.dbSpecified = dbSpecified[0]
	} else {
		return nil, errors.New("invalid parameter: dbSpecified")
	}

	// Open the DB, create it if not existed.
	err := a.open()
	if err != nil {
		return nil, err
	}

	// Call the destructor when the object is released.
	runtime.SetFinalizer(a, finalizer)

	return a, nil
}

// NewAdapterWithTableName  .
func NewAdapterWithTableName(driverName string, dataSourceName string, tableName string, tablePrefix string, dbSpecified ...bool) (*Adapter, error) {
	a := &Adapter{
		driverName:     driverName,
		dataSourceName: dataSourceName,
		tableName:      tableName,
		tablePrefix:    tablePrefix,
	}

	if len(dbSpecified) == 0 {
		a.dbSpecified = false
	} else if len(dbSpecified) == 1 {
		a.dbSpecified = dbSpecified[0]
	} else {
		return nil, errors.New("invalid parameter: dbSpecified")
	}

	// Open the DB, create it if not existed.
	err := a.open()
	if err != nil {
		return nil, err
	}

	// Call the destructor when the object is released.
	runtime.SetFinalizer(a, finalizer)

	return a, nil
}

// NewAdapterByEngine  .
func NewAdapterByEngine(engine *orm.Engine) (*Adapter, error) {
	a := &Adapter{
		engine: engine,
	}

	err := a.createTable()
	if err != nil {
		return nil, err
	}

	return a, nil
}

// NewAdapterByEngineWithTableName  .
func NewAdapterByEngineWithTableName(engine *orm.Engine, tableName string, tablePrefix string) (*Adapter, error) {
	a := &Adapter{
		engine:      engine,
		tableName:   tableName,
		tablePrefix: tablePrefix,
	}

	err := a.createTable()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Adapter) getFullTableName() string {
	if a.tablePrefix != "" {
		return a.tablePrefix + a.tableName
	}
	return a.tableName
}

func (a *Adapter) createDatabase() error {
	var err error
	var engine *orm.Engine
	if a.driverName == "postgres" {
		engine, err = orm.NewEngine(a.driverName, a.dataSourceName+" dbname=postgres")
	} else {
		engine, err = orm.NewEngine(a.driverName, a.dataSourceName)
	}
	if err != nil {
		return err
	}

	if a.driverName == "postgres" {
		if _, err = engine.Exec("CREATE DATABASE bhojpur"); err != nil {
			// 42P04 is	duplicate_database
			if pqerr, ok := err.(*pq.Error); ok && pqerr.Code == "42P04" {
				_ = engine.Close()
				return nil
			}
		}
	} else if a.driverName != "sqlite3" {
		_, err = engine.Exec("CREATE DATABASE IF NOT EXISTS bhojpur")
	}
	if err != nil {
		_ = engine.Close()
		return err
	}

	return engine.Close()
}

func (a *Adapter) open() error {
	var err error
	var engine *orm.Engine

	if a.dbSpecified {
		engine, err = orm.NewEngine(a.driverName, a.dataSourceName)
		if err != nil {
			return err
		}
	} else {
		if err = a.createDatabase(); err != nil {
			return err
		}

		if a.driverName == "postgres" {
			engine, err = orm.NewEngine(a.driverName, a.dataSourceName+" dbname=bhojpur")
		} else if a.driverName == "sqlite3" {
			engine, err = orm.NewEngine(a.driverName, a.dataSourceName)
		} else {
			engine, err = orm.NewEngine(a.driverName, a.dataSourceName+"bhojpur")
		}
		if err != nil {
			return err
		}
	}

	a.engine = engine

	return a.createTable()
}

func (a *Adapter) createTable() error {
	return a.engine.Sync2(&BhojpurRule{tableName: a.getFullTableName()})
}

func (a *Adapter) dropTable() error {
	return a.engine.DropTables(&BhojpurRule{tableName: a.getFullTableName()})
}

func loadPolicyLine(line *BhojpurRule, model model.Model) {
	var p = []string{line.PType,
		line.V0, line.V1, line.V2, line.V3, line.V4, line.V5}
	var lineText string
	if line.V5 != "" {
		lineText = strings.Join(p, ", ")
	} else if line.V4 != "" {
		lineText = strings.Join(p[:6], ", ")
	} else if line.V3 != "" {
		lineText = strings.Join(p[:5], ", ")
	} else if line.V2 != "" {
		lineText = strings.Join(p[:4], ", ")
	} else if line.V1 != "" {
		lineText = strings.Join(p[:3], ", ")
	} else if line.V0 != "" {
		lineText = strings.Join(p[:2], ", ")
	}

	persist.LoadPolicyLine(lineText, model)
}

// LoadPolicy loads policy from database.
func (a *Adapter) LoadPolicy(model model.Model) error {
	lines := make([]*BhojpurRule, 0, 64)

	if err := a.engine.Table(&BhojpurRule{tableName: a.getFullTableName()}).Find(&lines); err != nil {
		return err
	}

	for _, line := range lines {
		loadPolicyLine(line, model)
	}

	return nil
}

func (a *Adapter) genPolicyLine(ptype string, rule []string) *BhojpurRule {
	line := BhojpurRule{PType: ptype, tableName: a.getFullTableName()}

	l := len(rule)
	if l > 0 {
		line.V0 = rule[0]
	}
	if l > 1 {
		line.V1 = rule[1]
	}
	if l > 2 {
		line.V2 = rule[2]
	}
	if l > 3 {
		line.V3 = rule[3]
	}
	if l > 4 {
		line.V4 = rule[4]
	}
	if l > 5 {
		line.V5 = rule[5]
	}

	return &line
}

// SavePolicy saves policy to database.
func (a *Adapter) SavePolicy(model model.Model) error {
	err := a.dropTable()
	if err != nil {
		return err
	}
	err = a.createTable()
	if err != nil {
		return err
	}

	lines := make([]*BhojpurRule, 0, 64)

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := a.genPolicyLine(ptype, rule)
			lines = append(lines, line)
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := a.genPolicyLine(ptype, rule)
			lines = append(lines, line)
		}
	}

	// check whether the policy is empty
	if len(lines) == 0 {
		return nil
	}

	_, err = a.engine.Insert(&lines)

	return err
}

// AddPolicy adds a policy rule to the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	line := a.genPolicyLine(ptype, rule)
	_, err := a.engine.InsertOne(line)
	return err
}

// AddPolicies adds multiple policy rule to the storage.
func (a *Adapter) AddPolicies(sec string, ptype string, rules [][]string) error {
	_, err := a.engine.Transaction(func(tx *orm.Session) (interface{}, error) {
		for _, rule := range rules {
			line := a.genPolicyLine(ptype, rule)
			_, err := tx.InsertOne(line)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	})
	return err
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	line := a.genPolicyLine(ptype, rule)
	_, err := a.engine.Delete(line)
	return err
}

// RemovePolicies removes multiple policy rule from the storage.
func (a *Adapter) RemovePolicies(sec string, ptype string, rules [][]string) error {
	_, err := a.engine.Transaction(func(tx *orm.Session) (interface{}, error) {
		for _, rule := range rules {
			line := a.genPolicyLine(ptype, rule)
			_, err := tx.Delete(line)
			if err != nil {
				return nil, nil
			}
		}
		return nil, nil
	})
	return err
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	line := BhojpurRule{PType: ptype, tableName: a.getFullTableName()}

	idx := fieldIndex + len(fieldValues)
	if fieldIndex <= 0 && idx > 0 {
		line.V0 = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && idx > 1 {
		line.V1 = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && idx > 2 {
		line.V2 = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && idx > 3 {
		line.V3 = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && idx > 4 {
		line.V4 = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && idx > 5 {
		line.V5 = fieldValues[5-fieldIndex]
	}

	_, err := a.engine.Delete(&line)
	return err
}

// LoadFilteredPolicy loads only policy rules that match the filter.
func (a *Adapter) LoadFilteredPolicy(model model.Model, filter interface{}) error {
	filterValue, ok := filter.(Filter)
	if !ok {
		return errors.New("invalid filter type")
	}

	lines := make([]*BhojpurRule, 0, 64)
	if err := a.filterQuery(a.engine.NewSession(), filterValue).Table(&BhojpurRule{tableName: a.getFullTableName()}).Find(&lines); err != nil {
		return err
	}

	for _, line := range lines {
		loadPolicyLine(line, model)
	}
	a.isFiltered = true
	return nil
}

// IsFiltered returns true if the loaded policy has been filtered.
func (a *Adapter) IsFiltered() bool {
	return a.isFiltered
}

func (a *Adapter) filterQuery(session *orm.Session, filter Filter) *orm.Session {
	filterValue := [7]struct {
		col string
		val []string
	}{
		{"p_type", filter.PType},
		{"v0", filter.V0},
		{"v1", filter.V1},
		{"v2", filter.V2},
		{"v3", filter.V3},
		{"v4", filter.V4},
		{"v5", filter.V5},
	}

	for idx := range filterValue {
		switch len(filterValue[idx].val) {
		case 0:
			continue
		case 1:
			session.And(filterValue[idx].col+" = ?", filterValue[idx].val[0])
		default:
			session.In(filterValue[idx].col, filterValue[idx].val)
		}
	}

	return session
}

// UpdatePolicy update oldRule to newPolicy permanently
func (a *Adapter) UpdatePolicy(sec string, ptype string, oldRule, newPolicy []string) error {
	oRule := a.genPolicyLine(ptype, oldRule)
	_, err := a.engine.Update(a.genPolicyLine(ptype, newPolicy), oRule)
	return err
}

// UpdatePolicies updates some policy rules to storage, like db, redis.
func (a *Adapter) UpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) error {
	session := a.engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}

	for i, oldRule := range oldRules {
		nRule, oRule := a.genPolicyLine(ptype, newRules[i]), a.genPolicyLine(ptype, oldRule)
		if _, err := session.Update(nRule, oRule); err != nil {
			return err
		}
	}

	return session.Commit()
}

func (a *Adapter) UpdateFilteredPolicies(sec string, ptype string, newPolicies [][]string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	// UpdateFilteredPolicies deletes old rules and adds new rules.
	line := &BhojpurRule{}

	line.PType = ptype
	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		line.V0 = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		line.V1 = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		line.V2 = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		line.V3 = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		line.V4 = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		line.V5 = fieldValues[5-fieldIndex]
	}

	newP := make([]BhojpurRule, 0, len(newPolicies))
	oldP := make([]BhojpurRule, 0)
	for _, newRule := range newPolicies {
		newP = append(newP, *a.genPolicyLine(ptype, newRule))
	}
	tx := a.engine.NewSession().Table(&BhojpurRule{tableName: a.getFullTableName()})
	defer tx.Close()

	if err := tx.Begin(); err != nil {
		return nil, err
	}

	for i := range newP {
		str, args := line.queryString()
		if err := tx.Where(str, args...).Find(&oldP); err != nil {
			return nil, tx.Rollback()
		}
		if _, err := tx.Where(str.(string), args...).Delete(&BhojpurRule{tableName: a.getFullTableName()}); err != nil {
			return nil, tx.Rollback()
		}
		if _, err := tx.Insert(&newP[i]); err != nil {
			return nil, tx.Rollback()
		}
	}

	// return deleted rulues
	oldPolicies := make([][]string, 0)
	for _, v := range oldP {
		oldPolicy := v.toStringPolicy()
		oldPolicies = append(oldPolicies, oldPolicy)
	}
	return oldPolicies, tx.Commit()
}

func (c *BhojpurRule) toStringPolicy() []string {
	policy := make([]string, 0)
	if c.PType != "" {
		policy = append(policy, c.PType)
	}
	if c.V0 != "" {
		policy = append(policy, c.V0)
	}
	if c.V1 != "" {
		policy = append(policy, c.V1)
	}
	if c.V2 != "" {
		policy = append(policy, c.V2)
	}
	if c.V3 != "" {
		policy = append(policy, c.V3)
	}
	if c.V4 != "" {
		policy = append(policy, c.V4)
	}
	if c.V5 != "" {
		policy = append(policy, c.V5)
	}
	return policy
}

func (c *BhojpurRule) queryString() (interface{}, []interface{}) {
	queryArgs := []interface{}{c.PType}

	queryStr := "p_type = ?"
	if c.V0 != "" {
		queryStr += " and v0 = ?"
		queryArgs = append(queryArgs, c.V0)
	}
	if c.V1 != "" {
		queryStr += " and v1 = ?"
		queryArgs = append(queryArgs, c.V1)
	}
	if c.V2 != "" {
		queryStr += " and v2 = ?"
		queryArgs = append(queryArgs, c.V2)
	}
	if c.V3 != "" {
		queryStr += " and v3 = ?"
		queryArgs = append(queryArgs, c.V3)
	}
	if c.V4 != "" {
		queryStr += " and v4 = ?"
		queryArgs = append(queryArgs, c.V4)
	}
	if c.V5 != "" {
		queryStr += " and v5 = ?"
		queryArgs = append(queryArgs, c.V5)
	}

	return queryStr, queryArgs
}
