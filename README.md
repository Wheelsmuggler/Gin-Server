# Golang爬虫监控网页

使用golang 的 goquery 对于知乎大v的"想法"（ping）进行监控

* 爬虫在后台运行，每隔50～70s对网页进行一次爬取监测有无新内容的更新

* 如果有新的内容，则会在后台发送一封邮件提醒用户有新的内容，并且发送一个post请求给后台服务器让其更新数据  ([服务器代码](https://github.com/Wheelsmuggler/golang-Gin))