package optimization

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
