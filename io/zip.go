package io

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

type _zip struct{}

var ZIP = new(_zip)

//Compression 压缩文件
//@author: [SliverHorn](https://github.com/SliverHorn)
func (z _zip) Compression(fileName, oldForm, newForm string, files ...string) error {
	newZipFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() {
		_ = newZipFile.Close()
	}()

	zipWriter := zip.NewWriter(newZipFile)
	defer func() {
		_ = zipWriter.Close()
	}()

	// 把files添加到zip中
	for _, file := range files {
		if err = z.add(file, oldForm, newForm, zipWriter); err != nil {
			return err
		}
	}
	return nil
}

//add 创建压缩文件
func (z _zip) add(file, oldForm, newForm string, writer *zip.Writer) error {
	zipFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer zipFile.Close()
	// 获取file的基础信息
	info, err := zipFile.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// 使用上面的FileInforHeader() 就可以把文件保存的路径替换成我们自己想要的了，如下面
	header.Name = strings.Replace(file, oldForm, newForm, -1)

	// 优化压缩
	// 更多参考see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	headerWrite, err := writer.CreateHeader(header)
	if err != nil {
		return err
	}

	if _, err = io.Copy(headerWrite, zipFile); err != nil {
		return err
	}

	return nil

}
