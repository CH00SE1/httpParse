# [作者](https://github.com/CH00SE1/)  说明 

解析一些网站

## 技术选型

### ***golang + gorm + gin***

hs|gjyb|mysql|redis|golang|
:-----:|:-----:|:-----:|:-----:|:-----:|
基础网站|药品|8.0.29|5.0.24|1.8.4

## 项目内容
1. 基础hs网站(6个)
2. 药品网站(4个)

#### mysql 表名 t_hs_info
```aidl
CREATE TABLE `t_hs_info` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `title` varchar(256) NOT NULL COMMENT '标题',
  `url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '网站url',
  `m3u8_url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '下载url',
  `class_id` bigint DEFAULT NULL COMMENT '分类id',
  `platform` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '平台名称',
  `page` bigint DEFAULT NULL COMMENT '页码',
  `location` varchar(255) DEFAULT NULL COMMENT '位置',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_title_unique` (`title`) USING BTREE,
  KEY `idx_t_hs_info_deleted_at` (`deleted_at`),
  KEY `sel_title_normal` (`title`) USING BTREE,
  KEY `idx` (`id`) USING BTREE,
  KEY `m3u8Url_Idx` (`m3u8_url`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=309181 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```
#### 示例一（mysql数据同步redis）
```
// mysql数据同步redis
func Mysql2Redis() {
	db, err := db.MysqlConfigure()
	if err != nil {
		fmt.Println(err)
	}
	var infos []HsInfo
	// 查询数据
	db.Find(&infos)
	for _, info := range infos {
		// 添加序列化后的数据到redis
		marshal, _ := json.Marshal(info)
		redis.SetKey(info.Title, marshal)
	}
}
```
#### 示例二（遍历redis数据保存mysql）
```
// 同步redis数据 遍历redis数据
func Redis2Mysql() {
	keys := redis.GetKeyList()
	mysql, err := db.MysqlConfigure()
	if err != nil {
		fmt.Println("connent datebase err:", err)
	}
	infos := []*HsInfo{}
	for _, key := range keys {
		values, _ := redis.GetKey(key)
		var hsInfo HsInfo
		json.Unmarshal(utils.String2Bytes(values), &hsInfo)
		infos = append(infos, &hsInfo)
	}
	mysql.CreateInBatches(infos, 15).Callback()
}
```
#### 示例三（生成文件小案例）
```
func CreateFile() {
	content := capePageByIdInfo.Data.Content
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(content)))
	if err != nil {
		log.Fatal(err)
	}
	var Text string
	num := 1
	dom.Find("p").Each(func(i int, selection *goquery.Selection) {
		text := utils.StringStrip(selection.Text())
		if len(text) != 0 {
			Text += strconv.Itoa(num) + "." + text + "\n"
			num++
		}
	})
	if num != 1 {
		utils.CreateFile(&Text, "C:\\Users\\Administrator\\Desktop\\海角社区\\", title, ".txt")
	}
}
```
#### 数据展示
![image](https://github.com/CH00SE1/httpParse/blob/prod/other/t_hs_info.png)
