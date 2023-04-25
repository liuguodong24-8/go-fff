package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
	"helloworld/internal/util"
)

type studentRepo struct {
	data *Data
	log  *log.Helper
}

func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *studentRepo) Save(ctx context.Context, stu *biz.Student) (*biz.Student, error) {
	return stu, nil
}

func (repo *studentRepo) Get(ctx context.Context, stu *biz.Student) (*biz.Student, error) {
	return stu, nil
}

func (repo *studentRepo) GetStudent(ctx context.Context, id int32) (*biz.Student, error) {
	var stu biz.Student
	repo.data.gormDB.Where("id = ?", id).First(&stu) // gorm
	repo.log.WithContext(ctx).Info("gormDB: GetStudent, id: ", id)
	return &biz.Student{
		ID:        stu.ID,
		Name:      stu.Name,
		Status:    stu.Status,
		Info:      stu.Info,
		UpdatedAt: stu.UpdatedAt,
		CreatedAt: stu.CreatedAt,
	}, nil
}

func (repo *studentRepo) GetStudents(ctx context.Context, name string, page int32, pageSize int32) ([]*biz.Student, error) {
	var stus []*biz.Student
	db := repo.data.gormDB
	if name != "" {
		db = db.Scopes(util.ColumnEqualScope("name", name))
	}
	if err := db.Scopes(util.PaginationScope((page-1)*pageSize, pageSize)).Find(&stus).Error; err != nil {
		return []*biz.Student{}, nil
	}
	return stus, nil
}
