package model

import (
    "gorm.io/gorm"
)

type UserEntity struct {
    gorm.Model
    Name string
}

func GetAll() (datas []UserEntity) {
    result := Db.Find(&datas)
    if result.Error != nil {
        panic(result.Error)
    }
    return
}

func GetOne(id int) (data UserEntity) {
    result := Db.First(&data, id)
    if result.Error != nil {
        panic(result.Error)
    }
    return
}

func (b *UserEntity) Create() {
    result := Db.Create(b)
    if result.Error != nil {
        panic(result.Error)
    }
}

func (b *UserEntity) Update() {
    result := Db.Save(b)
    if result.Error != nil {
        panic(result.Error)
    }
}

func (b *UserEntity) Delete() {
    result := Db.Delete(b)
    if result.Error != nil {
        panic(result.Error)
    }
}
