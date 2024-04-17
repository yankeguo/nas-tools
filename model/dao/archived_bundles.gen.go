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

func newArchivedBundle(db *gorm.DB, opts ...gen.DOOption) archivedBundle {
	_archivedBundle := archivedBundle{}

	_archivedBundle.archivedBundleDo.UseDB(db, opts...)
	_archivedBundle.archivedBundleDo.UseModel(&model.ArchivedBundle{})

	tableName := _archivedBundle.archivedBundleDo.TableName()
	_archivedBundle.ALL = field.NewAsterisk(tableName)
	_archivedBundle.ID = field.NewString(tableName, "id")
	_archivedBundle.Year = field.NewString(tableName, "year")
	_archivedBundle.CreatedAt = field.NewTime(tableName, "created_at")
	_archivedBundle.Tape = field.NewString(tableName, "tape")

	_archivedBundle.fillFieldMap()

	return _archivedBundle
}

type archivedBundle struct {
	archivedBundleDo

	ALL       field.Asterisk
	ID        field.String
	Year      field.String
	CreatedAt field.Time
	Tape      field.String

	fieldMap map[string]field.Expr
}

func (a archivedBundle) Table(newTableName string) *archivedBundle {
	a.archivedBundleDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a archivedBundle) As(alias string) *archivedBundle {
	a.archivedBundleDo.DO = *(a.archivedBundleDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *archivedBundle) updateTableName(table string) *archivedBundle {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewString(table, "id")
	a.Year = field.NewString(table, "year")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.Tape = field.NewString(table, "tape")

	a.fillFieldMap()

	return a
}

func (a *archivedBundle) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *archivedBundle) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 4)
	a.fieldMap["id"] = a.ID
	a.fieldMap["year"] = a.Year
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["tape"] = a.Tape
}

func (a archivedBundle) clone(db *gorm.DB) archivedBundle {
	a.archivedBundleDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a archivedBundle) replaceDB(db *gorm.DB) archivedBundle {
	a.archivedBundleDo.ReplaceDB(db)
	return a
}

type archivedBundleDo struct{ gen.DO }

func (a archivedBundleDo) Debug() *archivedBundleDo {
	return a.withDO(a.DO.Debug())
}

func (a archivedBundleDo) WithContext(ctx context.Context) *archivedBundleDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a archivedBundleDo) ReadDB() *archivedBundleDo {
	return a.Clauses(dbresolver.Read)
}

func (a archivedBundleDo) WriteDB() *archivedBundleDo {
	return a.Clauses(dbresolver.Write)
}

func (a archivedBundleDo) Session(config *gorm.Session) *archivedBundleDo {
	return a.withDO(a.DO.Session(config))
}

func (a archivedBundleDo) Clauses(conds ...clause.Expression) *archivedBundleDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a archivedBundleDo) Returning(value interface{}, columns ...string) *archivedBundleDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a archivedBundleDo) Not(conds ...gen.Condition) *archivedBundleDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a archivedBundleDo) Or(conds ...gen.Condition) *archivedBundleDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a archivedBundleDo) Select(conds ...field.Expr) *archivedBundleDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a archivedBundleDo) Where(conds ...gen.Condition) *archivedBundleDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a archivedBundleDo) Order(conds ...field.Expr) *archivedBundleDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a archivedBundleDo) Distinct(cols ...field.Expr) *archivedBundleDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a archivedBundleDo) Omit(cols ...field.Expr) *archivedBundleDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a archivedBundleDo) Join(table schema.Tabler, on ...field.Expr) *archivedBundleDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a archivedBundleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *archivedBundleDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a archivedBundleDo) RightJoin(table schema.Tabler, on ...field.Expr) *archivedBundleDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a archivedBundleDo) Group(cols ...field.Expr) *archivedBundleDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a archivedBundleDo) Having(conds ...gen.Condition) *archivedBundleDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a archivedBundleDo) Limit(limit int) *archivedBundleDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a archivedBundleDo) Offset(offset int) *archivedBundleDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a archivedBundleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *archivedBundleDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a archivedBundleDo) Unscoped() *archivedBundleDo {
	return a.withDO(a.DO.Unscoped())
}

func (a archivedBundleDo) Create(values ...*model.ArchivedBundle) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a archivedBundleDo) CreateInBatches(values []*model.ArchivedBundle, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a archivedBundleDo) Save(values ...*model.ArchivedBundle) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a archivedBundleDo) First() (*model.ArchivedBundle, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedBundle), nil
	}
}

func (a archivedBundleDo) Take() (*model.ArchivedBundle, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedBundle), nil
	}
}

func (a archivedBundleDo) Last() (*model.ArchivedBundle, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedBundle), nil
	}
}

func (a archivedBundleDo) Find() ([]*model.ArchivedBundle, error) {
	result, err := a.DO.Find()
	return result.([]*model.ArchivedBundle), err
}

func (a archivedBundleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ArchivedBundle, err error) {
	buf := make([]*model.ArchivedBundle, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a archivedBundleDo) FindInBatches(result *[]*model.ArchivedBundle, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a archivedBundleDo) Attrs(attrs ...field.AssignExpr) *archivedBundleDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a archivedBundleDo) Assign(attrs ...field.AssignExpr) *archivedBundleDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a archivedBundleDo) Joins(fields ...field.RelationField) *archivedBundleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a archivedBundleDo) Preload(fields ...field.RelationField) *archivedBundleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a archivedBundleDo) FirstOrInit() (*model.ArchivedBundle, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedBundle), nil
	}
}

func (a archivedBundleDo) FirstOrCreate() (*model.ArchivedBundle, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedBundle), nil
	}
}

func (a archivedBundleDo) FindByPage(offset int, limit int) (result []*model.ArchivedBundle, count int64, err error) {
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

func (a archivedBundleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a archivedBundleDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a archivedBundleDo) Delete(models ...*model.ArchivedBundle) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *archivedBundleDo) withDO(do gen.Dao) *archivedBundleDo {
	a.DO = *do.(*gen.DO)
	return a
}
