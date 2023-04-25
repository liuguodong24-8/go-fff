package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Student struct {
	ID        int32
	Name      string
	Info      string
	Status    int32
	UpdatedAt time.Time
	CreatedAt time.Time
}

// StudentRepo 定义student的操作接口
type StudentRepo interface {
	Save(ctx context.Context, student *Student) (*Student, error)
	Get(ctx context.Context, student *Student) (*Student, error)
	GetStudent(context.Context, int32) (*Student, error)                   // 根据 id 获取学生信息
	GetStudents(context.Context, string, int32, int32) ([]*Student, error) // 获取学生信息
}

type StudentUserCase struct {
	repo StudentRepo
	log  *log.Helper
}

// NewStudentUserCase 初始化StudentUserCase
func NewStudentUserCase(repo StudentRepo, logger log.Logger) *StudentUserCase {
	return &StudentUserCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *StudentUserCase) CreateStudent(ctx context.Context, stu *Student) (*Student, error) {
	uc.log.WithContext(ctx).Info("CreateStudent: %v", stu.ID)
	return uc.repo.Save(ctx, stu)
}

func (uc *StudentUserCase) Get(ctx context.Context, id int32) (*Student, error) {
	uc.log.WithContext(ctx).Infof("biz.Get: %d", id)
	return uc.repo.GetStudent(ctx, id)
}

func (uc *StudentUserCase) List(ctx context.Context, name string, page int32, pageSize int32) ([]*Student, error) {
	uc.log.WithContext(ctx).Info("List: %v", name)
	return uc.repo.GetStudents(ctx, name, page, pageSize)
}
