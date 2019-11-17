package helper

func Paginator(totalNum uint32, perPageNum uint16, page uint16) (totalPage uint16, offset uint16) {
	totalPage = uint16(totalNum / uint32(perPageNum))
	offset = perPageNum * (page - 1)
	return totalPage, offset
}
