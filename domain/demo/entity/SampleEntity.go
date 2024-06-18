package entity

import (
	BaseEntity "app/base/entity"
	UserEntity "app/domain/user/entity"
	"gorm.io/gorm"
	"log"
	"time"
)

/**
notes: 领域层-模型类
说明: 负责基础层的工作,字段过滤(模型黑白名单),用户权限(模型策略),触发器(模型事件),等一系列传统DBA负责的工作.
*/

// 当前实体模型

type Sample struct {
	ID uint `json:"id"  gorm:"size:20;primaryKey"`

	Name   string `json:"name" gorm:"size:50"`
	Mobile string `json:"mobile" gorm:"size:11;unique"`
	Photo  string `json:"photo" gorm:"size:200"`
	Sex    uint8  `json:"sex" gorm:"size:3;default:0"`
	Type   uint8  `json:"type" gorm:"size:3;default:0"`
	Status uint8  `json:"status" gorm:"size:3;default:0"`

	CreatedAt time.Time      `json:"created_at" gorm:"<_;autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"<_;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"<_;index"`

	//关联模型
	User  UserEntity.User   `json:"user" gorm:"foreignKey:ID"`
	Users []UserEntity.User `json:"users" gorm:"foreignKey:ID"`
}

// 当前类原型

type SampleEntityStruct struct {
	BaseEntity.BaseEntityInterface
}

// 初始化

func SampleEntity() *gorm.DB {
	instance := &SampleEntityStruct{&BaseEntity.BaseEntityStruct{}}
	query := instance.Connector("Connector").Model(&Sample{}).Table("samples")
	return query
}

//新增&更新事件 - 按执行顺序

func (entity *Sample) BeforeSave(tx *gorm.DB) (err error) {
	log.Println("BeforeSave sample entity")
	return
}
func (entity *Sample) AfterSave(tx *gorm.DB) (err error) {
	log.Println("AfterSave sample entity")
	return
}

//新增事件 - 按执行顺序

func (entity *Sample) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("BeforeCreate sample entity")
	return
}
func (entity *Sample) AfterCreate(tx *gorm.DB) (err error) {
	log.Println("AfterCreate sample entity")
	return
}

//更新事件 - 按执行顺序

func (entity *Sample) BeforeUpdate(tx *gorm.DB) (err error) {
	log.Println("BeforeUpdate sample entity")
	return
}
func (entity *Sample) AfterUpdate(tx *gorm.DB) (err error) {
	log.Println("AfterUpdate sample entity")
	return
}

//删除事件 - 按执行顺序

func (entity *Sample) BeforeDelete(tx *gorm.DB) (err error) {
	log.Println("BeforeDelete sample entity")
	return
}
func (entity *Sample) AfterDelete(tx *gorm.DB) (err error) {
	log.Println("AfterDelete sample entity")
	return
}
