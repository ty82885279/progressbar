package main

import (
	"fmt"
	"time"
)

type Bar struct {
	percent int64  //百分比
	cur     int64  //当前进度位置
	total   int64  //总进度
	rate    string //进度条
	style   string //进度条样式
}

func (b *Bar) NewBar(start, total int64) {
	b.cur = start
	b.total = total
	if b.style == "" {
		b.style = "█"
	}
	b.percent = b.getPercent()

}
func (b *Bar) NewBarWithStyle(start, total int64, style string) {
	b.style = style
	b.NewBar(start, total)
}
func (b *Bar) getPercent() int64 {

	return int64(float32(b.cur) / float32(b.total) * 100)
}
func (b *Bar) Play(cur int64) {
	b.cur = cur
	lastP := b.percent //上次完成的百分比
	b.percent = b.getPercent()
	if lastP != b.percent && b.percent%2 == 0 {

		b.rate += b.style
	}
	fmt.Printf("\r[ %-50s ] %3d%%  %d/%d", b.rate, b.percent, b.cur, b.total)

}
func (b *Bar) Finish() {

	fmt.Println()
}
func main() {

	var bar Bar
	bar.NewBar(0, 100)
	bar.NewBarWithStyle(0, 100, "#")
	for i := 0; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Play(int64(i))
	}
	bar.Finish()
}
