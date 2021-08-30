package DictSeg

import (
	"fmt"
	"time"

	"github.com/ksclouds/PowerNLP/Config"
	"github.com/ksclouds/PowerNLP/Seg/Collections"
)

var (
	MapTrieSeg *Collections.MapTrie = nil
)

func init() {
	MapTrieSeg = Collections.NewMapTrie()
	MapTrieSeg.LoadDict("data/dict.txt")
	time.Sleep(time.Second * 2)
	fmt.Println("字典加载完毕")
	MapTrieSeg.LoadDict(Config.DictPath)
}
