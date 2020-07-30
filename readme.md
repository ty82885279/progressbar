## 进度条

### 有些工作可能需要一段时间才能完成,为了提交体验,需要加入进度条功能

默认样式

```go
var bar Bar
	bar.NewBar(0, 100)
	for i := 0; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Play(int64(i))
	}
	bar.Finish()
```



![1](/Users/mrlee/Desktop/progressbar/1.GIF)

指定样式

```go
var bar Bar
	bar.NewBarWithStyle(0, 100, "#")
	for i := 0; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Play(int64(i))
	}
	bar.Finish()
```

![2](/Users/mrlee/Desktop/progressbar/2.GIF)