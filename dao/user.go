package dao

import (
	"just.for.test/rbacdemo/model"
)

func FindUserByName(name string) (*model.User, error) {
	result := new(model.User)
	if err := model.TestDb.Model(&model.User{}).
		Where("user_name=?", name).
		Find(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
