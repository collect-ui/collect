// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model_gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"collect/model"
)

func newAttendance(db *gorm.DB, opts ...gen.DOOption) attendance {
	_attendance := attendance{}

	_attendance.attendanceDo.UseDB(db, opts...)
	_attendance.attendanceDo.UseModel(&model.Attendance{})

	tableName := _attendance.attendanceDo.TableName()
	_attendance.ALL = field.NewAsterisk(tableName)
	_attendance.ID = field.NewString(tableName, "id")
	_attendance.AttendanceTime = field.NewString(tableName, "attendance_time")
	_attendance.UserID = field.NewString(tableName, "user_id")

	_attendance.fillFieldMap()

	return _attendance
}

type attendance struct {
	attendanceDo

	ALL            field.Asterisk
	ID             field.String
	AttendanceTime field.String
	UserID         field.String

	fieldMap map[string]field.Expr
}

func (a attendance) Table(newTableName string) *attendance {
	a.attendanceDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a attendance) As(alias string) *attendance {
	a.attendanceDo.DO = *(a.attendanceDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *attendance) updateTableName(table string) *attendance {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewString(table, "id")
	a.AttendanceTime = field.NewString(table, "attendance_time")
	a.UserID = field.NewString(table, "user_id")

	a.fillFieldMap()

	return a
}

func (a *attendance) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *attendance) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 3)
	a.fieldMap["id"] = a.ID
	a.fieldMap["attendance_time"] = a.AttendanceTime
	a.fieldMap["user_id"] = a.UserID
}

func (a attendance) clone(db *gorm.DB) attendance {
	a.attendanceDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a attendance) replaceDB(db *gorm.DB) attendance {
	a.attendanceDo.ReplaceDB(db)
	return a
}

type attendanceDo struct{ gen.DO }

type IAttendanceDo interface {
	gen.SubQuery
	Debug() IAttendanceDo
	WithContext(ctx context.Context) IAttendanceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAttendanceDo
	WriteDB() IAttendanceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAttendanceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAttendanceDo
	Not(conds ...gen.Condition) IAttendanceDo
	Or(conds ...gen.Condition) IAttendanceDo
	Select(conds ...field.Expr) IAttendanceDo
	Where(conds ...gen.Condition) IAttendanceDo
	Order(conds ...field.Expr) IAttendanceDo
	Distinct(cols ...field.Expr) IAttendanceDo
	Omit(cols ...field.Expr) IAttendanceDo
	Join(table schema.Tabler, on ...field.Expr) IAttendanceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAttendanceDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAttendanceDo
	Group(cols ...field.Expr) IAttendanceDo
	Having(conds ...gen.Condition) IAttendanceDo
	Limit(limit int) IAttendanceDo
	Offset(offset int) IAttendanceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAttendanceDo
	Unscoped() IAttendanceDo
	Create(values ...*model.Attendance) error
	CreateInBatches(values []*model.Attendance, batchSize int) error
	Save(values ...*model.Attendance) error
	First() (*model.Attendance, error)
	Take() (*model.Attendance, error)
	Last() (*model.Attendance, error)
	Find() ([]*model.Attendance, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Attendance, err error)
	FindInBatches(result *[]*model.Attendance, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Attendance) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAttendanceDo
	Assign(attrs ...field.AssignExpr) IAttendanceDo
	Joins(fields ...field.RelationField) IAttendanceDo
	Preload(fields ...field.RelationField) IAttendanceDo
	FirstOrInit() (*model.Attendance, error)
	FirstOrCreate() (*model.Attendance, error)
	FindByPage(offset int, limit int) (result []*model.Attendance, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAttendanceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a attendanceDo) Debug() IAttendanceDo {
	return a.withDO(a.DO.Debug())
}

func (a attendanceDo) WithContext(ctx context.Context) IAttendanceDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a attendanceDo) ReadDB() IAttendanceDo {
	return a.Clauses(dbresolver.Read)
}

func (a attendanceDo) WriteDB() IAttendanceDo {
	return a.Clauses(dbresolver.Write)
}

func (a attendanceDo) Session(config *gorm.Session) IAttendanceDo {
	return a.withDO(a.DO.Session(config))
}

func (a attendanceDo) Clauses(conds ...clause.Expression) IAttendanceDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a attendanceDo) Returning(value interface{}, columns ...string) IAttendanceDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a attendanceDo) Not(conds ...gen.Condition) IAttendanceDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a attendanceDo) Or(conds ...gen.Condition) IAttendanceDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a attendanceDo) Select(conds ...field.Expr) IAttendanceDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a attendanceDo) Where(conds ...gen.Condition) IAttendanceDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a attendanceDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAttendanceDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a attendanceDo) Order(conds ...field.Expr) IAttendanceDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a attendanceDo) Distinct(cols ...field.Expr) IAttendanceDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a attendanceDo) Omit(cols ...field.Expr) IAttendanceDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a attendanceDo) Join(table schema.Tabler, on ...field.Expr) IAttendanceDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a attendanceDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAttendanceDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a attendanceDo) RightJoin(table schema.Tabler, on ...field.Expr) IAttendanceDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a attendanceDo) Group(cols ...field.Expr) IAttendanceDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a attendanceDo) Having(conds ...gen.Condition) IAttendanceDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a attendanceDo) Limit(limit int) IAttendanceDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a attendanceDo) Offset(offset int) IAttendanceDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a attendanceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAttendanceDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a attendanceDo) Unscoped() IAttendanceDo {
	return a.withDO(a.DO.Unscoped())
}

func (a attendanceDo) Create(values ...*model.Attendance) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a attendanceDo) CreateInBatches(values []*model.Attendance, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a attendanceDo) Save(values ...*model.Attendance) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a attendanceDo) First() (*model.Attendance, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Attendance), nil
	}
}

func (a attendanceDo) Take() (*model.Attendance, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Attendance), nil
	}
}

func (a attendanceDo) Last() (*model.Attendance, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Attendance), nil
	}
}

func (a attendanceDo) Find() ([]*model.Attendance, error) {
	result, err := a.DO.Find()
	return result.([]*model.Attendance), err
}

func (a attendanceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Attendance, err error) {
	buf := make([]*model.Attendance, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a attendanceDo) FindInBatches(result *[]*model.Attendance, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a attendanceDo) Attrs(attrs ...field.AssignExpr) IAttendanceDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a attendanceDo) Assign(attrs ...field.AssignExpr) IAttendanceDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a attendanceDo) Joins(fields ...field.RelationField) IAttendanceDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a attendanceDo) Preload(fields ...field.RelationField) IAttendanceDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a attendanceDo) FirstOrInit() (*model.Attendance, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Attendance), nil
	}
}

func (a attendanceDo) FirstOrCreate() (*model.Attendance, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Attendance), nil
	}
}

func (a attendanceDo) FindByPage(offset int, limit int) (result []*model.Attendance, count int64, err error) {
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

func (a attendanceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a attendanceDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a attendanceDo) Delete(models ...*model.Attendance) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *attendanceDo) withDO(do gen.Dao) *attendanceDo {
	a.DO = *do.(*gen.DO)
	return a
}
