# bilibili-rank
爬取B站排行榜13个分区的近期和每日排行

- 采用[xrom](https://github.com/go-xorm/xorm)来实现自动创建列表和colum，并使用批量insert功能使得存储时间大幅度下降，提高了程序运行效率
- 采用[configor](github.com/jinzhu/configor)来实现配置文件的管理。通过config.yml文件，可以轻松的改变程序的内部配置。

