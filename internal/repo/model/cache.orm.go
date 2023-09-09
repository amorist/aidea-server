package model

// !!! DO NOT EDIT THIS FILE

import (
	"context"
	"encoding/json"
	"github.com/iancoleman/strcase"
	"github.com/mylxsw/eloquent/query"
	"gopkg.in/guregu/null.v3"
	"time"
)

func init() {

}

// CacheN is a Cache object, all fields are nullable
type CacheN struct {
	original   *cacheOriginal
	cacheModel *CacheModel

	Id         null.Int    `json:"id"`
	Key        null.String `json:"key"`
	Value      null.String `json:"value"`
	ValidUntil null.Time   `json:"valid_until"`
	CreatedAt  null.Time
	UpdatedAt  null.Time
}

// As convert object to other type
// dst must be a pointer to struct
func (inst *CacheN) As(dst interface{}) error {
	return query.Copy(inst, dst)
}

// SetModel set model for Cache
func (inst *CacheN) SetModel(cacheModel *CacheModel) {
	inst.cacheModel = cacheModel
}

// cacheOriginal is an object which stores original Cache from database
type cacheOriginal struct {
	Id         null.Int
	Key        null.String
	Value      null.String
	ValidUntil null.Time
	CreatedAt  null.Time
	UpdatedAt  null.Time
}

// Staled identify whether the object has been modified
func (inst *CacheN) Staled(onlyFields ...string) bool {
	if inst.original == nil {
		inst.original = &cacheOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			return true
		}
		if inst.Key != inst.original.Key {
			return true
		}
		if inst.Value != inst.original.Value {
			return true
		}
		if inst.ValidUntil != inst.original.ValidUntil {
			return true
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			return true
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			return true
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					return true
				}
			case "key":
				if inst.Key != inst.original.Key {
					return true
				}
			case "value":
				if inst.Value != inst.original.Value {
					return true
				}
			case "valid_until":
				if inst.ValidUntil != inst.original.ValidUntil {
					return true
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					return true
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					return true
				}
			default:
			}
		}
	}

	return false
}

// StaledKV return all fields has been modified
func (inst *CacheN) StaledKV(onlyFields ...string) query.KV {
	kv := make(query.KV, 0)

	if inst.original == nil {
		inst.original = &cacheOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			kv["id"] = inst.Id
		}
		if inst.Key != inst.original.Key {
			kv["key"] = inst.Key
		}
		if inst.Value != inst.original.Value {
			kv["value"] = inst.Value
		}
		if inst.ValidUntil != inst.original.ValidUntil {
			kv["valid_until"] = inst.ValidUntil
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			kv["created_at"] = inst.CreatedAt
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			kv["updated_at"] = inst.UpdatedAt
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					kv["id"] = inst.Id
				}
			case "key":
				if inst.Key != inst.original.Key {
					kv["key"] = inst.Key
				}
			case "value":
				if inst.Value != inst.original.Value {
					kv["value"] = inst.Value
				}
			case "valid_until":
				if inst.ValidUntil != inst.original.ValidUntil {
					kv["valid_until"] = inst.ValidUntil
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					kv["created_at"] = inst.CreatedAt
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					kv["updated_at"] = inst.UpdatedAt
				}
			default:
			}
		}
	}

	return kv
}

// Save create a new model or update it
func (inst *CacheN) Save(ctx context.Context, onlyFields ...string) error {
	if inst.cacheModel == nil {
		return query.ErrModelNotSet
	}

	id, _, err := inst.cacheModel.SaveOrUpdate(ctx, *inst, onlyFields...)
	if err != nil {
		return err
	}

	inst.Id = null.IntFrom(id)
	return nil
}

// Delete remove a cache
func (inst *CacheN) Delete(ctx context.Context) error {
	if inst.cacheModel == nil {
		return query.ErrModelNotSet
	}

	_, err := inst.cacheModel.DeleteById(ctx, inst.Id.Int64)
	if err != nil {
		return err
	}

	return nil
}

// String convert instance to json string
func (inst *CacheN) String() string {
	rs, _ := json.Marshal(inst)
	return string(rs)
}

type cacheScope struct {
	name  string
	apply func(builder query.Condition)
}

var cacheGlobalScopes = make([]cacheScope, 0)
var cacheLocalScopes = make([]cacheScope, 0)

// AddGlobalScopeForCache assign a global scope to a model
func AddGlobalScopeForCache(name string, apply func(builder query.Condition)) {
	cacheGlobalScopes = append(cacheGlobalScopes, cacheScope{name: name, apply: apply})
}

// AddLocalScopeForCache assign a local scope to a model
func AddLocalScopeForCache(name string, apply func(builder query.Condition)) {
	cacheLocalScopes = append(cacheLocalScopes, cacheScope{name: name, apply: apply})
}

func (m *CacheModel) applyScope() query.Condition {
	scopeCond := query.ConditionBuilder()
	for _, g := range cacheGlobalScopes {
		if m.globalScopeEnabled(g.name) {
			g.apply(scopeCond)
		}
	}

	for _, s := range cacheLocalScopes {
		if m.localScopeEnabled(s.name) {
			s.apply(scopeCond)
		}
	}

	return scopeCond
}

func (m *CacheModel) localScopeEnabled(name string) bool {
	for _, n := range m.includeLocalScopes {
		if name == n {
			return true
		}
	}

	return false
}

func (m *CacheModel) globalScopeEnabled(name string) bool {
	for _, n := range m.excludeGlobalScopes {
		if name == n {
			return false
		}
	}

	return true
}

type Cache struct {
	Id         int64     `json:"id"`
	Key        string    `json:"key"`
	Value      string    `json:"value"`
	ValidUntil time.Time `json:"valid_until"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (w Cache) ToCacheN(allows ...string) CacheN {
	if len(allows) == 0 {
		return CacheN{

			Id:         null.IntFrom(int64(w.Id)),
			Key:        null.StringFrom(w.Key),
			Value:      null.StringFrom(w.Value),
			ValidUntil: null.TimeFrom(w.ValidUntil),
			CreatedAt:  null.TimeFrom(w.CreatedAt),
			UpdatedAt:  null.TimeFrom(w.UpdatedAt),
		}
	}

	res := CacheN{}
	for _, al := range allows {
		switch strcase.ToSnake(al) {

		case "id":
			res.Id = null.IntFrom(int64(w.Id))
		case "key":
			res.Key = null.StringFrom(w.Key)
		case "value":
			res.Value = null.StringFrom(w.Value)
		case "valid_until":
			res.ValidUntil = null.TimeFrom(w.ValidUntil)
		case "created_at":
			res.CreatedAt = null.TimeFrom(w.CreatedAt)
		case "updated_at":
			res.UpdatedAt = null.TimeFrom(w.UpdatedAt)
		default:
		}
	}

	return res
}

// As convert object to other type
// dst must be a pointer to struct
func (w Cache) As(dst interface{}) error {
	return query.Copy(w, dst)
}

func (w *CacheN) ToCache() Cache {
	return Cache{

		Id:         w.Id.Int64,
		Key:        w.Key.String,
		Value:      w.Value.String,
		ValidUntil: w.ValidUntil.Time,
		CreatedAt:  w.CreatedAt.Time,
		UpdatedAt:  w.UpdatedAt.Time,
	}
}

// CacheModel is a model which encapsulates the operations of the object
type CacheModel struct {
	db        *query.DatabaseWrap
	tableName string

	excludeGlobalScopes []string
	includeLocalScopes  []string

	query query.SQLBuilder
}

var cacheTableName = "cache"

// CacheTable return table name for Cache
func CacheTable() string {
	return cacheTableName
}

const (
	FieldCacheId         = "id"
	FieldCacheKey        = "key"
	FieldCacheValue      = "value"
	FieldCacheValidUntil = "valid_until"
	FieldCacheCreatedAt  = "created_at"
	FieldCacheUpdatedAt  = "updated_at"
)

// CacheFields return all fields in Cache model
func CacheFields() []string {
	return []string{
		"id",
		"key",
		"value",
		"valid_until",
		"created_at",
		"updated_at",
	}
}

func SetCacheTable(tableName string) {
	cacheTableName = tableName
}

// NewCacheModel create a CacheModel
func NewCacheModel(db query.Database) *CacheModel {
	return &CacheModel{
		db:                  query.NewDatabaseWrap(db),
		tableName:           cacheTableName,
		excludeGlobalScopes: make([]string, 0),
		includeLocalScopes:  make([]string, 0),
		query:               query.Builder(),
	}
}

// GetDB return database instance
func (m *CacheModel) GetDB() query.Database {
	return m.db.GetDB()
}

func (m *CacheModel) clone() *CacheModel {
	return &CacheModel{
		db:                  m.db,
		tableName:           m.tableName,
		excludeGlobalScopes: append([]string{}, m.excludeGlobalScopes...),
		includeLocalScopes:  append([]string{}, m.includeLocalScopes...),
		query:               m.query,
	}
}

// WithoutGlobalScopes remove a global scope for given query
func (m *CacheModel) WithoutGlobalScopes(names ...string) *CacheModel {
	mc := m.clone()
	mc.excludeGlobalScopes = append(mc.excludeGlobalScopes, names...)

	return mc
}

// WithLocalScopes add a local scope for given query
func (m *CacheModel) WithLocalScopes(names ...string) *CacheModel {
	mc := m.clone()
	mc.includeLocalScopes = append(mc.includeLocalScopes, names...)

	return mc
}

// Condition add query builder to model
func (m *CacheModel) Condition(builder query.SQLBuilder) *CacheModel {
	mm := m.clone()
	mm.query = mm.query.Merge(builder)

	return mm
}

// Find retrieve a model by its primary key
func (m *CacheModel) Find(ctx context.Context, id int64) (*CacheN, error) {
	return m.First(ctx, m.query.Where("id", "=", id))
}

// Exists return whether the records exists for a given query
func (m *CacheModel) Exists(ctx context.Context, builders ...query.SQLBuilder) (bool, error) {
	count, err := m.Count(ctx, builders...)
	return count > 0, err
}

// Count return model count for a given query
func (m *CacheModel) Count(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {
	sqlStr, params := m.query.
		Merge(builders...).
		Table(m.tableName).
		AppendCondition(m.applyScope()).
		ResolveCount()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	rows.Next()
	var res int64
	if err := rows.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (m *CacheModel) Paginate(ctx context.Context, page int64, perPage int64, builders ...query.SQLBuilder) ([]CacheN, query.PaginateMeta, error) {
	if page <= 0 {
		page = 1
	}

	if perPage <= 0 {
		perPage = 15
	}

	meta := query.PaginateMeta{
		PerPage: perPage,
		Page:    page,
	}

	count, err := m.Count(ctx, builders...)
	if err != nil {
		return nil, meta, err
	}

	meta.Total = count
	meta.LastPage = count / perPage
	if count%perPage != 0 {
		meta.LastPage += 1
	}

	res, err := m.Get(ctx, append([]query.SQLBuilder{query.Builder().Limit(perPage).Offset((page - 1) * perPage)}, builders...)...)
	if err != nil {
		return res, meta, err
	}

	return res, meta, nil
}

// Get retrieve all results for given query
func (m *CacheModel) Get(ctx context.Context, builders ...query.SQLBuilder) ([]CacheN, error) {
	b := m.query.Merge(builders...).Table(m.tableName).AppendCondition(m.applyScope())
	if len(b.GetFields()) == 0 {
		b = b.Select(
			"id",
			"key",
			"value",
			"valid_until",
			"created_at",
			"updated_at",
		)
	}

	fields := b.GetFields()
	selectFields := make([]query.Expr, 0)

	for _, f := range fields {
		switch strcase.ToSnake(f.Value) {

		case "id":
			selectFields = append(selectFields, f)
		case "key":
			selectFields = append(selectFields, f)
		case "value":
			selectFields = append(selectFields, f)
		case "valid_until":
			selectFields = append(selectFields, f)
		case "created_at":
			selectFields = append(selectFields, f)
		case "updated_at":
			selectFields = append(selectFields, f)
		}
	}

	var createScanVar = func(fields []query.Expr) (*CacheN, []interface{}) {
		var cacheVar CacheN
		scanFields := make([]interface{}, 0)

		for _, f := range fields {
			switch strcase.ToSnake(f.Value) {

			case "id":
				scanFields = append(scanFields, &cacheVar.Id)
			case "key":
				scanFields = append(scanFields, &cacheVar.Key)
			case "value":
				scanFields = append(scanFields, &cacheVar.Value)
			case "valid_until":
				scanFields = append(scanFields, &cacheVar.ValidUntil)
			case "created_at":
				scanFields = append(scanFields, &cacheVar.CreatedAt)
			case "updated_at":
				scanFields = append(scanFields, &cacheVar.UpdatedAt)
			}
		}

		return &cacheVar, scanFields
	}

	sqlStr, params := b.Fields(selectFields...).ResolveQuery()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	caches := make([]CacheN, 0)
	for rows.Next() {
		cacheReal, scanFields := createScanVar(fields)
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}

		cacheReal.original = &cacheOriginal{}
		_ = query.Copy(cacheReal, cacheReal.original)

		cacheReal.SetModel(m)
		caches = append(caches, *cacheReal)
	}

	return caches, nil
}

// First return first result for given query
func (m *CacheModel) First(ctx context.Context, builders ...query.SQLBuilder) (*CacheN, error) {
	res, err := m.Get(ctx, append(builders, query.Builder().Limit(1))...)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, query.ErrNoResult
	}

	return &res[0], nil
}

// Create save a new cache to database
func (m *CacheModel) Create(ctx context.Context, kv query.KV) (int64, error) {

	if _, ok := kv["created_at"]; !ok {
		kv["created_at"] = time.Now()
	}

	if _, ok := kv["updated_at"]; !ok {
		kv["updated_at"] = time.Now()
	}

	sqlStr, params := m.query.Table(m.tableName).ResolveInsert(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// SaveAll save all caches to database
func (m *CacheModel) SaveAll(ctx context.Context, caches []CacheN) ([]int64, error) {
	ids := make([]int64, 0)
	for _, cache := range caches {
		id, err := m.Save(ctx, cache)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Save save a cache to database
func (m *CacheModel) Save(ctx context.Context, cache CacheN, onlyFields ...string) (int64, error) {
	return m.Create(ctx, cache.StaledKV(onlyFields...))
}

// SaveOrUpdate save a new cache or update it when it has a id > 0
func (m *CacheModel) SaveOrUpdate(ctx context.Context, cache CacheN, onlyFields ...string) (id int64, updated bool, err error) {
	if cache.Id.Int64 > 0 {
		_, _err := m.UpdateById(ctx, cache.Id.Int64, cache, onlyFields...)
		return cache.Id.Int64, true, _err
	}

	_id, _err := m.Save(ctx, cache, onlyFields...)
	return _id, false, _err
}

// UpdateFields update kv for a given query
func (m *CacheModel) UpdateFields(ctx context.Context, kv query.KV, builders ...query.SQLBuilder) (int64, error) {
	if len(kv) == 0 {
		return 0, nil
	}

	kv["updated_at"] = time.Now()

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).
		Table(m.tableName).
		ResolveUpdate(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Update update a model for given query
func (m *CacheModel) Update(ctx context.Context, builder query.SQLBuilder, cache CacheN, onlyFields ...string) (int64, error) {
	return m.UpdateFields(ctx, cache.StaledKV(onlyFields...), builder)
}

// UpdateById update a model by id
func (m *CacheModel) UpdateById(ctx context.Context, id int64, cache CacheN, onlyFields ...string) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).UpdateFields(ctx, cache.StaledKV(onlyFields...))
}

// Delete remove a model
func (m *CacheModel) Delete(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).Table(m.tableName).ResolveDelete()

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}

// DeleteById remove a model by id
func (m *CacheModel) DeleteById(ctx context.Context, id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Delete(ctx)
}
