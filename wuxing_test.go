package wuxing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWuXing_String(t *testing.T) {
	assertions := assert.New(t)
	names := []string{"金", "木", "土", "水", "火", "空"}
	for i := int8(0); i < 1; i++ {
		assertions.Equal(names[i], fmt.Sprintf("%v", WuXing(1<<i)))
	}
}

func TestWuXing_Restrain(t *testing.T) {
	assertions := assert.New(t)
	restrains := []string{"金克木", "木克土", "土克水", "水克火", "火克金"}
	for i := 0; i < 5; i++ {
		a := WuXing(1 << i)
		for j := 0; j < 5; j++ {
			b := WuXing(1 << j)
			if a.Restrain(b) {
				assertions.Equal(restrains[i], fmt.Sprintf("%v克%v", a, b))
			}
		}
	}
}
func TestWuXing_Promote(t *testing.T) {
	assertions := assert.New(t)
	promotes := []string{"金生水", "木生火", "土生金", "水生木", "火生土"}
	for i := 0; i < 5; i++ {
		a := WuXing(1 << i)
		for j := 0; j < 5; j++ {
			b := WuXing(1 << j)
			if a.Promote(b) {
				assertions.Equal(promotes[i], fmt.Sprintf("%v生%v", a, b))
			}
		}
	}
}
