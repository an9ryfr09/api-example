package helper

import (
	"a6-api/utils/loader"

	"github.com/gin-gonic/gin"
)

//Paginator get pagin params totalPage and offset
//@param totalNum total item numbers
//@param perPageNum perpage numbers
//@param page current page number
//@return totalPage total page numbers
//@return offset perpage begin offset number
func Paginator(totalNum uint32, perPageNum uint16, page uint16) (totalPage uint16, offset uint16) {
	if totalNum < 1 {
		return 0, 0
	}

	if perPageNum < 1 {
		perPageNum = loader.Load().Core.PerPageNum
	}

	if page < 1 {
		page = 1
	}

	totalPage = uint16(totalNum / uint32(perPageNum))
	offset = perPageNum * (page - 1)
	return totalPage, offset
}

//GeneratePaginInfo returns pagin info
//@param totalNum total item numbers
//@param totalPage total page numbers
//@param page current page number
//@param offset perpage begin offset number
func GeneratePaginInfo(totalNum uint32, totalPage uint16, page uint16, perPageNum uint16, offset uint16) map[string]interface{} {
	return gin.H{
		"totalNum":   totalNum,
		"totalPage":  totalPage,
		"page":       page,
		"perPageNum": perPageNum,
		"offset":     offset,
	}
}
