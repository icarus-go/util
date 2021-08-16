package optimization

import "math"

var Page = new(_page)

type _page struct{}

//Each 分页循环
func (p *_page) Each(length, max int, function func(start, end int) error) error {
	remainder := length % max // 取余
	page := length / max
	if remainder > 0 {
		page++
	}
	start := 0
	end := 0
	for i := 0; i < page; i++ {
		start = i * max
		end = (i + 1) * max
		if length < end {
			end = length
		}

		if err := function(start, end); err != nil {
			return err
		}
	}
	return nil
}

//One 根据页码以及页面条数 计算当前长度是否可以返回切片的起始索引下标和结束索引下标
// page 页码 min: 1
// size 页码 min: 1 , max: length
// length 切片长度
func (p *_page) One(page, size, length int) (start, end int) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	// 总页数
	pageCount := int(math.Ceil(float64(length) / float64(size)))

	if pageCount < page {
		return
	}

	if size > length {
		end = length
	}

	if page <= pageCount {
		start = (page - 1) * size
		end = start + size

		if end > length {
			end = length
		}
	}

	return start, end

}
