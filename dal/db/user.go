package db

import (
	"douyin/pkg/constant"
)

type User struct {
	ID              uint64 `json:"id"`                                                                  // 自增主键
	Username        string `gorm:"index:idx_username,unique;type:varchar(63);not null" json:"username"` //用户名
	Password        string `gorm:"type:varchar(255);not null" json:"password"`                          //用户密码
	FollowingCount  uint64 `gorm:"default:0;not null" json:"following_count"`                           //关注数
	FollowerCount   uint64 `gorm:"default:0;not null" json:"follower_count"`                            //粉丝数
	Avatar          string `gorm:"type:varchar(255);not null" json:"avatar"`                            //用户头像
	BackgroundImage string `gorm:"type:varchar(255);not null" json:"background_image"`                  //用户个人页顶部大图
	Signature       string `gorm:"type:varchar(255);not null" json:"signature"`                         //个人简介
	TotalFavorited  uint64 `gorm:"default:0;not null" json:"total_favorited"`                           //获赞数量
	WorkCount       uint64 `gorm:"default:0;not null" json:"work_count"`                                //作品数量
	FavoriteCount   uint64 `gorm:"default:0;not null" json:"favorite_count"`                            //点赞数量
}

func (u *User) TableName() string {
	return constant.UserTableName
}

func CreateUser(user *User) (userID int64, err error) {
	if err := DB.Create(user).Error; err != nil {
		return 0, err
	}
	return int64(user.ID), nil
}

func SelectUserByID(userID uint64) (*User, error) {
	res := User{ID: userID}
	if err := DB.First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func SelectUserByName(username string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.Where("username = ?", username).Limit(1).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func IncreaseUserFavoriteCount(userID uint64) (uint64, error) {
	user := &User{
		ID: userID,
	}
	err := DB.First(&user).Error
	if err != nil {
		return 0, err
	}
	if err := DB.Model(&user).Update("favorite_count", user.FavoriteCount+1).Error; err != nil {
		return 0, err
	}
	return user.FavoriteCount, nil
}

func DecreaseUserFavoriteCount(userID uint64) (uint64, error) {
	user := &User{
		ID: userID,
	}
	err := DB.First(&user).Error
	if err != nil {
		return 0, err
	}
	if err := DB.Model(&user).Update("favorite_count", user.FavoriteCount-1).Error; err != nil {
		return 0, err
	}
	return user.FavoriteCount, nil
}