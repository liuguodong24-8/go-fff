package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Pledge struct {
	ID        int32
	TokenID   int32
	Price     int64
	Status    int32
	Times     int32
	Address   string
	CreatedAt time.Time
	UpdateAt  time.Time
}

// PledgeRepo 定义student的操作接口
type PledgeRepo interface {
	Save(ctx context.Context, pledge *Pledge) (*Pledge, error)
	Get(ctx context.Context, pledge *Pledge) (*Pledge, error)
	GetPledge(context.Context, string) (*Pledge, error)                  // 根据 id 获取学生信息
	GetPledges(context.Context, string, int32, int32) ([]*Pledge, error) // 获取学生信息
}

type PledgeCase struct {
	repo PledgeRepo
	log  *log.Helper
}

// NewPledgeCase 初始化StudentUserCase
func NewPledgeCase(repo PledgeRepo, logger log.Logger) *PledgeCase {
	return &PledgeCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PledgeCase) CreatePledge(ctx context.Context, ple *Pledge) (*Pledge, error) {
	uc.log.WithContext(ctx).Info("CreateStudent: %v", ple.ID)
	return uc.repo.Save(ctx, ple)
}

func (uc *PledgeCase) GetPledge(ctx context.Context, address string) (*Pledge, error) {
	uc.log.WithContext(ctx).Infof("biz.Get: %s", address)
	return uc.repo.GetPledge(ctx, address)
}

func (uc *PledgeCase) ListPledge(ctx context.Context, name string, page int32, pageSize int32) ([]*Pledge, error) {
	uc.log.WithContext(ctx).Info("List: %v", name)
	return uc.repo.GetPledges(ctx, name, page, pageSize)
}
