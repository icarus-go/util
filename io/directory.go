package io

import (
	"go.uber.org/zap"
	"os"
)

var Directory = new(_directory)

type _directory struct{}

//PathExists 文件目录是否存在
//@author: [SliverHorn](https://github.com/SliverHorn)
func (d *_directory) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//BatchCreate 批量创建文件夹
//@author: [SliverHorn](https://github.com/SliverHorn)
func (d *_directory) BatchCreate(directories ...string) error {
	for _, directory := range directories {
		if exist, err := d.PathExists(directory); err != nil {
			return err
		} else {
			if !exist {
				if err = os.MkdirAll(directory, os.ModePerm); err != nil {
					zap.L().Info("Function os.MkdirAll Failed!", zap.Error(err))
				}
			}
		}
	}
	zap.L().Info("Batch Create Succeed!", zap.Strings("directory", directories))
	return nil
}

//Create 批量创建文件夹
//@author: [SliverHorn](https://github.com/SliverHorn)
func (d *_directory) Create(dirs ...string) error {
	for _, v := range dirs {
		if exist, err := d.PathExists(v); err != nil {
			return err
		} else {
			if !exist {
				zap.L().Debug("create directory" + v)
				err = os.MkdirAll(v, os.ModePerm)
				if err != nil {
					zap.L().Error("create directory"+v, zap.Any(" error:", err))
				}
			}
		}
	}
	return nil
}
