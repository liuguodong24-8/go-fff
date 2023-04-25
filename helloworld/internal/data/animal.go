package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
	"helloworld/internal/util"
)

type pledgeRepo struct {
	data *Data
	log  *log.Helper
}

func NewPledgeRepo(data *Data, logger log.Logger) biz.PledgeRepo {
	return &pledgeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *pledgeRepo) Save(ctx context.Context, ple *biz.Pledge) (*biz.Pledge, error) {
	err := repo.data.gormDB.Create(ple)
	if err != nil {
		return &biz.Pledge{}, err.Error
	}
	return ple, nil
}

func (repo *pledgeRepo) Get(ctx context.Context, stu *biz.Pledge) (*biz.Pledge, error) {
	return stu, nil
}

func (repo *pledgeRepo) GetPledge(ctx context.Context, address string) (*biz.Pledge, error) {
	var ple biz.Pledge
	repo.data.gormDB.Where("address = ?", address).First(&ple) // gorm
	repo.log.WithContext(ctx).Info("gormDB: GetPledge, address: ", address)
	return &biz.Pledge{
		ID:        ple.ID,
		TokenID:   ple.TokenID,
		Address:   ple.Address,
		Price:     ple.Price,
		Times:     ple.Times,
		CreatedAt: ple.CreatedAt,
		UpdateAt:  ple.UpdateAt,
	}, nil
}

func (repo *pledgeRepo) GetPledges(ctx context.Context, name string, page int32, pageSize int32) ([]*biz.Pledge, error) {
	var ples []*biz.Pledge
	db := repo.data.gormDB
	if name != "" {
		db = db.Scopes(util.ColumnEqualScope("name", name))
	}
	if err := db.Scopes(util.PaginationScope((page-1)*pageSize, pageSize)).Find(&ples).Error; err != nil {
		return []*biz.Pledge{}, nil
	}
	return ples, nil
}
