package helper

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	totalNum   uint32
	perPageNum uint16
	page       uint16
)

func TestPaginator(t *testing.T) {

	totalNum = 100
	perPageNum = 10
	page = 1

	t.Run("function Paginator", func(t *testing.T) {
		totalPage, offset := Paginator(totalNum, perPageNum, page)

		if totalPage != 10 || offset != 0 {
			t.Fatalf("want totalPage is 10, got totalPage is %d", totalPage)
			t.Fatalf("want offset is 0, got offset is %d", offset)
		}
	})

	t.Run("function GeneratePaginInfo", func(t *testing.T) {
		totalPage, offset := Paginator(totalNum, perPageNum, page)
		got := GeneratePaginInfo(totalNum, totalPage, page, perPageNum, offset)

		want := gin.H{
			"totalNum":   totalNum,
			"totalPage":  uint16(10),
			"page":       page,
			"perPageNum": perPageNum,
			"offset":     uint16(0),
		}

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want is %s, got is %s", want, got)
		}
	})
}
