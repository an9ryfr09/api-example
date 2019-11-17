package helper

func Paginator(totalNum uint32, perPageNum uint16, page uint16) (totalPage uint16, offset uint16) {
	totalPage = uint16(totalNum / uint32(perPageNum))
	offset = perPageNum * (page - 1)
	return totalPage, offset
}

func GeneratePaginInfo(totalNum uint32, totalPage uint16, page uint16, perPageNum uint16, offset uint16) map[string]interface{} {
	return map[string]interface{}{
		"totalNum":   totalNum,
		"totalPage":  totalPage,
		"page":       page,
		"perPageNum": perPageNum,
		"offset":     offset,
	}
}
