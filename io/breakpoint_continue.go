package io

import (
	"github.com/icarus-go/utils/encrypt"
	"github.com/icarus-go/utils/io/constant"
	"io/ioutil"
	"os"
	"strconv"
)

// 前端传来文件片与当前片为什么文件的第几片
// 后端拿到以后比较次分片是否上传 或者是否为不完全片
// 前端发送每片多大
// 前端告知是否为最后一片且是否完成
type _breakPointContinue struct {
	file     *os.File
	fileInfo []os.FileInfo
}

//BreakPointContinue 断点续传
var BreakPointContinue = new(_breakPointContinue)

//BreakPointContinue 断点续传
//author: [SliverHorn](https://github.com/SliverHorn)
func (f *_breakPointContinue) BreakPointContinue(content []byte, fileName, fileMd5 string, contentNumber int) (path string, err error) {
	path = constant.BreakPointDir + fileMd5 + "/"
	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		return path, err
	}
	return f.makeFileContent(content, fileName, path, contentNumber)
}

//CheckMd5 检查Md5
//author: [SliverHorn](https://github.com/SliverHorn)
// return
//  param bool 是否可以继续
func (f *_breakPointContinue) CheckMd5(content []byte, chunkMd5 string) bool {
	fileMd5 := encrypt.MD5V(content)
	return fileMd5 == chunkMd5
}

//author: [SliverHorn](https://github.com/SliverHorn)
//makeFileContent 创建切片内容
func (f *_breakPointContinue) makeFileContent(content []byte, fileName string, fileDir string, contentNumber int) (path string, err error) {
	path = fileDir + fileName + "_" + strconv.Itoa(contentNumber)
	if f.file, err = os.Create(path); err != nil {
		return path, err
	} else {
		if _, err = f.file.Write(content); err != nil {
			return path, err
		}
	}
	defer func() {
		_ = f.file.Close()
	}()
	return path, nil
}

//MakeFile 创建切片文件
//author: [SliverHorn](https://github.com/SliverHorn)
func (f *_breakPointContinue) MakeFile(fileName string, FileMd5 string) (path string, err error) {
	path = constant.FinishDir + fileName
	if f.fileInfo, err = ioutil.ReadDir(constant.BreakPointDir + FileMd5); err != nil {
		return path, err
	}
	_ = os.MkdirAll(constant.FinishDir, os.ModePerm)
	if f.file, err = os.OpenFile(constant.FinishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644); err != nil {
		return path, err
	}
	defer func() {
		_ = f.file.Close()
	}()
	for k := range f.fileInfo {
		content, _ := ioutil.ReadFile(constant.BreakPointDir + FileMd5 + "/" + fileName + "_" + strconv.Itoa(k))
		if _, err = f.file.Write(content); err != nil {
			_ = os.Remove(constant.FinishDir + fileName)
			return path, err
		}
	}
	return path, err
}

//RemoveChunk 移除切片
//author: [SliverHorn](https://github.com/SliverHorn)
func (f *_breakPointContinue) RemoveChunk(fileMD5 string) error {
	return os.RemoveAll(constant.BreakPointDir + fileMD5)
}
