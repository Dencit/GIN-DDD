package entity

import (
	BaseEntity "app/base/entity"
	"gorm.io/gorm"
	"log"
	"time"
)

/**
notes: 领域层-模型类
说明: 负责基础层的工作,字段过滤(模型黑白名单),用户权限(模型策略),触发器(模型事件),等一系列传统DBA负责的工作.
*/

// 当前实体模型

type User struct {
	ID uint `json:"id"  gorm:"size:20;primaryKey"`

	Role         string `json:"role" gorm:"size:200"`
	Name         string `json:"name" gorm:"size:200"`
	NickName     string `json:"nick_name" gorm:"size:200"`
	Mobile       string `json:"mobile" gorm:"size:11;unique"`
	Avatar       string `json:"avatar" gorm:"size:200"`
	PassWord     string `json:"pass_word" gorm:"size:200"`
	ClientDriver string `json:"client_driver" gorm:"size:65530"`

	Sex        uint8 `json:"sex" gorm:"size:3;default:0"`
	Status     uint8 `json:"status" gorm:"size:3;default:0"`
	ClientType uint8 `json:"client_type" gorm:"size:3;default:0"`

	Lat float32 `json:"lat" gorm:"size:10,6;default:0.0"`
	Lng float32 `json:"lng" gorm:"size:10,6;default:0.0"`

	OnLineTime  time.Time `json:"on_line_time" gorm:"<_;default:null"`
	OffLineTime time.Time `json:"off_line_time" gorm:"<_;default:null"`

	CreatedAt time.Time      `json:"created_at" gorm:"<_;autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"<_;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"<_;index"`
}

// 当前类原型

type UserEntityStruct struct {
	BaseEntity.BaseEntityInterface
}

// 初始化

func UserEntity() *gorm.DB {
	instance := &UserEntityStruct{&BaseEntity.BaseEntityStruct{}}
	query := instance.Connector("Connector").Model(&User{}).Table("users")
	return query
}

//新增&更新事件 - 按执行顺序

func (entity *User) BeforeSave(tx *gorm.DB) (err error) {
	log.Println("BeforeSave sample entity")
	return
}
func (entity *User) AfterSave(tx *gorm.DB) (err error) {
	log.Println("AfterSave sample entity")
	return
}

//新增事件 - 按执行顺序

func (entity *User) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("BeforeCreate sample entity")
	return
}
func (entity *User) AfterCreate(tx *gorm.DB) (err error) {
	log.Println("AfterCreate sample entity")
	return
}

//更新事件 - 按执行顺序

func (entity *User) BeforeUpdate(tx *gorm.DB) (err error) {
	log.Println("BeforeUpdate sample entity")
	return
}
func (entity *User) AfterUpdate(tx *gorm.DB) (err error) {
	log.Println("AfterUpdate sample entity")
	return
}

//删除事件 - 按执行顺序

func (entity *User) BeforeDelete(tx *gorm.DB) (err error) {
	log.Println("BeforeDelete sample entity")
	return
}
func (entity *User) AfterDelete(tx *gorm.DB) (err error) {
	log.Println("AfterDelete sample entity")
	return
}
