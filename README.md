# Sniffer
嗅探器：golang实现的图片抓取程序

**注意：仅供学习、娱乐**

## 说明

- 图片来源网站：[煎蛋](http://jandan.net)
- 依赖项目：[goquery](https://github.com/PuerkitoBio/goquery)
- 参考项目：[Meizar](https://github.com/qibin0506/Meizar)

## 版本日志

- 创建日期：2017年3月29日

## 参数说明

```bash
Sniffer --help
Usage of Sniffer:
  -end int
        end holds the last page number to finish sniffer (default 1)
  -path string
        path holds the path to store the pictures (default "./images")
  -start int
        start holds the first page to start sniffer (default 1)
```

## 编译运行

0、前提：需要安装了golang运行环境

1、在命令行环境下，进入Sniffer目录

2、编译

```bash
go install
```

3、运行

```bash
Sniffer -start 1 -end 10
```


## 输出结果

程序运行成功后，命令行输出如下，并且会将抓取到的图片存储在指定文件夹下。如果未指定图片存储路径，默认会将图片存放在Sniffer/images目录下。

```bash
$ Sniffer -start 1 -end 2
create storage directory:       ../images/
download picture:       http://ww3.sinaimg.cn/large/7c8e8afbjw1dh9yimwp4xj.jpg
success
download picture:       http://ww4.sinaimg.cn/large/62ca3fd6jw1dhckht5jlwj.jpg
success
download picture:       http://ww2.sinaimg.cn/large/6f86cef5jw1dhccu6q2mpj.jpg
success
download picture:       http://ww3.sinaimg.cn/large/64e5c3e2jw1dhch767eiqj.jpg
success
download picture:       http://ww3.sinaimg.cn/large/64e5c3e2jw1dhch3jiklpj.jpg
success
download picture:       http://ww2.sinaimg.cn/large/6ed8a358jw1dhanny2mpuj.jpg
success
download picture:       http://ww2.sinaimg.cn/large/792427dbjw1dhdnkuc6jqj.jpg
success
download picture:       http://ww3.sinaimg.cn/large/79242f11jw1dhdglgsx5hj.jpg
success
download picture:       http://ww1.sinaimg.cn/large/8001cc87jw1dhdpwj2nhhj.jpg
success
download picture:       http://ww2.sinaimg.cn/large/6aa4686bjw1dhdpl7ffruj.jpg
success
download picture:       http://ww2.sinaimg.cn/large/8001cc87jw1dhdpj6uw3nj.jpg
success
download picture:       http://ww3.sinaimg.cn/large/792427dbjw1dhbsjda0f9j.jpg
success
download picture:       http://ww3.sinaimg.cn/large/7ce249cbjw1dhdhxx9n9lj.jpg
success
download picture:       http://ww2.sinaimg.cn/large/62a8fcf3jw1dhcrxl2754j.jpg
success
download picture:       http://ww4.sinaimg.cn/large/7c8e8afbjw1dhdh0zwj5oj.jpg
success
download picture:       http://ww3.sinaimg.cn/large/61e8a1fdjw1dhcvzracy1j.jpg
success
download picture:       http://ww1.sinaimg.cn/large/6c83d38djw1dhbp3kvouxj.jpg
success
```

## 自定义扩展

程序默认是从[煎蛋](http://jandan.net)上抓取图片，然而用户可以自定义抓取规则，根据想要抓取的图片网站的特性，实现Sniffer/code/rule/rule.go里的Rule接口里的3个方法，即可进行抓取。

例如，煎蛋的抓取规则如下：

```go
// JandanRule is an implementation of Rule
type JandanRule struct{}

// NewJandanRule create a new JandanRule object
func NewJandanRule() Rule {
	return &JandanRule{}
}

// UrlRule implements the method in interface of Rule
func (p *JandanRule) URLRule() (url string) {
	return "http://jandan.net/ooxx/"
}

// PageRule implements the method in interface of Rule
func (p *JandanRule) PageRule(pageNumber int) (page string) {
	return "page-" + strconv.Itoa(pageNumber)
}

// ImageRule implements the method in interface of Rule
func (p *JandanRule) ImageRule(doc *goquery.Document, f func(image string)) {
	doc.Find("a.view_img_link").Each(func(i int, s *goquery.Selection) {
		if img, exist := s.Attr("href"); exist {
			f(img)
		}
	})
}
```

说明：
1. 第一个方法返回我们要抓取的URL
2. 第二个方法根据当前页返回URL后面的页面信息
3. 第三个方法是内容匹配规则， 将匹配到的内容利用f函数返回，f函数可根据不同情况自定义

