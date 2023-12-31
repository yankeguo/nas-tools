// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/yankeguo/nas-tools/model"
)

func newArchivedFileIgnore(db *gorm.DB, opts ...gen.DOOption) archivedFileIgnore {
	_archivedFileIgnore := archivedFileIgnore{}

	_archivedFileIgnore.archivedFileIgnoreDo.UseDB(db, opts...)
	_archivedFileIgnore.archivedFileIgnoreDo.UseModel(&model.ArchivedFileIgnore{})

	tableName := _archivedFileIgnore.archivedFileIgnoreDo.TableName()
	_archivedFileIgnore.ALL = field.NewAsterisk(tableName)
	_archivedFileIgnore.ID = field.NewUint(tableName, "id")
	_archivedFileIgnore.Year = field.NewString(tableName, "year")
	_archivedFileIgnore.Bundle = field.NewString(tableName, "bundle")
	_archivedFileIgnore.Dir = field.NewString(tableName, "dir")
	_archivedFileIgnore.CreatedAt = field.NewTime(tableName, "created_at")

	_archivedFileIgnore.fillFieldMap()

	return _archivedFileIgnore
}

type archivedFileIgnore struct {
	archivedFileIgnoreDo

	ALL       field.Asterisk
	ID        field.Uint
	Year      field.String
	Bundle    field.String
	Dir       field.String
	CreatedAt field.Time

	fieldMap map[string]field.Expr
}

func (a archivedFileIgnore) Table(newTableName string) *archivedFileIgnore {
	a.archivedFileIgnoreDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a archivedFileIgnore) As(alias string) *archivedFileIgnore {
	a.archivedFileIgnoreDo.DO = *(a.archivedFileIgnoreDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *archivedFileIgnore) updateTableName(table string) *archivedFileIgnore {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.Year = field.NewString(table, "year")
	a.Bundle = field.NewString(table, "bundle")
	a.Dir = field.NewString(table, "dir")
	a.CreatedAt = field.NewTime(table, "created_at")

	a.fillFieldMap()

	return a
}

func (a *archivedFileIgnore) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *archivedFileIgnore) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 5)
	a.fieldMap["id"] = a.ID
	a.fieldMap["year"] = a.Year
	a.fieldMap["bundle"] = a.Bundle
	a.fieldMap["dir"] = a.Dir
	a.fieldMap["created_at"] = a.CreatedAt
}

func (a archivedFileIgnore) clone(db *gorm.DB) archivedFileIgnore {
	a.archivedFileIgnoreDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a archivedFileIgnore) replaceDB(db *gorm.DB) archivedFileIgnore {
	a.archivedFileIgnoreDo.ReplaceDB(db)
	return a
}

type archivedFileIgnoreDo struct{ gen.DO }

func (a archivedFileIgnoreDo) Debug() *archivedFileIgnoreDo {
	return a.withDO(a.DO.Debug())
}

func (a archivedFileIgnoreDo) WithContext(ctx context.Context) *archivedFileIgnoreDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a archivedFileIgnoreDo) ReadDB() *archivedFileIgnoreDo {
	return a.Clauses(dbresolver.Read)
}

func (a archivedFileIgnoreDo) WriteDB() *archivedFileIgnoreDo {
	return a.Clauses(dbresolver.Write)
}

func (a archivedFileIgnoreDo) Session(config *gorm.Session) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Session(config))
}

func (a archivedFileIgnoreDo) Clauses(conds ...clause.Expression) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a archivedFileIgnoreDo) Returning(value interface{}, columns ...string) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a archivedFileIgnoreDo) Not(conds ...gen.Condition) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a archivedFileIgnoreDo) Or(conds ...gen.Condition) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a archivedFileIgnoreDo) Select(conds ...field.Expr) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a archivedFileIgnoreDo) Where(conds ...gen.Condition) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a archivedFileIgnoreDo) Order(conds ...field.Expr) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a archivedFileIgnoreDo) Distinct(cols ...field.Expr) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a archivedFileIgnoreDo) Omit(cols ...field.Expr) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a archivedFileIgnoreDo) Join(table schema.Tabler, on ...field.Expr) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a archivedFileIgnoreDo) LeftJoin(table schema.Tabler, on ...field.Expr) *archivedFileIgnoreDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a archivedFileIgnoreDo) RightJoin(table schema.Tabler, on ...field.Expr) *archivedFileIgnoreDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a archivedFileIgnoreDo) Group(cols ...field.Expr) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a archivedFileIgnoreDo) Having(conds ...gen.Condition) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a archivedFileIgnoreDo) Limit(limit int) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a archivedFileIgnoreDo) Offset(offset int) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a archivedFileIgnoreDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a archivedFileIgnoreDo) Unscoped() *archivedFileIgnoreDo {
	return a.withDO(a.DO.Unscoped())
}

func (a archivedFileIgnoreDo) Create(values ...*model.ArchivedFileIgnore) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a archivedFileIgnoreDo) CreateInBatches(values []*model.ArchivedFileIgnore, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a archivedFileIgnoreDo) Save(values ...*model.ArchivedFileIgnore) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a archivedFileIgnoreDo) First() (*model.ArchivedFileIgnore, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedFileIgnore), nil
	}
}

func (a archivedFileIgnoreDo) Take() (*model.ArchivedFileIgnore, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedFileIgnore), nil
	}
}

func (a archivedFileIgnoreDo) Last() (*model.ArchivedFileIgnore, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedFileIgnore), nil
	}
}

func (a archivedFileIgnoreDo) Find() ([]*model.ArchivedFileIgnore, error) {
	result, err := a.DO.Find()
	return result.([]*model.ArchivedFileIgnore), err
}

func (a archivedFileIgnoreDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ArchivedFileIgnore, err error) {
	buf := make([]*model.ArchivedFileIgnore, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a archivedFileIgnoreDo) FindInBatches(result *[]*model.ArchivedFileIgnore, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a archivedFileIgnoreDo) Attrs(attrs ...field.AssignExpr) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a archivedFileIgnoreDo) Assign(attrs ...field.AssignExpr) *archivedFileIgnoreDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a archivedFileIgnoreDo) Joins(fields ...field.RelationField) *archivedFileIgnoreDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a archivedFileIgnoreDo) Preload(fields ...field.RelationField) *archivedFileIgnoreDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a archivedFileIgnoreDo) FirstOrInit() (*model.ArchivedFileIgnore, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedFileIgnore), nil
	}
}

func (a archivedFileIgnoreDo) FirstOrCreate() (*model.ArchivedFileIgnore, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedFileIgnore), nil
	}
}

func (a archivedFileIgnoreDo) FindByPage(offset int, limit int) (result []*model.ArchivedFileIgnore, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a archivedFileIgnoreDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a archivedFileIgnoreDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a archivedFileIgnoreDo) Delete(models ...*model.ArchivedFileIgnore) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *archivedFileIgnoreDo) withDO(do gen.Dao) *archivedFileIgnoreDo {
	a.DO = *do.(*gen.DO)
	return a
}
