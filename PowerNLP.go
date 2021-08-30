package PowerNLP

import "github.com/ksclouds/PowerNLP/Seg"

//默认分词方法
func Segment(sentence string) []string {
	return Seg.DefaultSegment().Segment(sentence)
}
