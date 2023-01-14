# hotSearch
weibo baidu zhihu 热搜列表

SLQ
```mysql
create table hotSearchInfo (
    id int primary key auto_increment comment 'ID',
    web_source varchar(20) comment '热搜来源[sina/baidu/zhihu]',
    web_id varchar(50) comment '热搜来源编码',
    hot_note varchar(300) comment '热搜词',
    hot_rank int comment '当日热搜排名',
    hot_desc varchar(500) comment '热搜详情',
    hot_score varchar(30) comment '热搜指数',
    hot_tag varchar(10) comment '热搜种类',
    hot_url varchar(500) comment '访问地址',
    hot_date date comment '热搜日期',
    mnt_time datetime comment '维护时间',
    hot_opt1 varchar(500) comment 'hot_opt1',
    hot_opt2 varchar(500) comment 'hot_opt2'

);

select * from hotSearchInfo;
```
