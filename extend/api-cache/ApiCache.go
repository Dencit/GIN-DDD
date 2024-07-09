package api_cache

import (
	"app/base/config"
	"app/extend/convert/arrays"
	"app/extend/convert/maps"
	"app/extend/convert/values"
	JsonTool "app/extend/json-tool"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"
)

var rdbList = make(map[string]*redis.Client) // 连接实例-列表
var rdbSync = sync.Map{}                     // 连接实例-原子标记

//api json 缓存结构

type jsonStruct struct {
	Data any `json:"data"`
	Meta any `json:"meta"`
}

//类结构

type ApiCacheStruct struct {
	rdb      *redis.Client   //连接实例
	ctx      context.Context //载体结构
	QueryMap map[string]any
}

//实例初始化

func Instance(requestQuery map[string]any) *ApiCacheStruct {
	instance := &ApiCacheStruct{}
	//设置载体
	instance.ctx = context.Background()
	//暂存请求
	instance.QueryMap = requestQuery
	//redis库名
	rdbName := "0"

	//获取原子标记
	rdbTag, _ := rdbSync.Load(rdbName)
	if rdbTag == nil {
		//记录原子标记
		rdbSync.Store(rdbName, "1")

		//连接设置
		rdb := redis.NewClient(&redis.Options{
			Addr:     config.Rdb.Addr,
			Password: config.Rdb.Pwd,
			DB:       config.Rdb.Select,
		})
		//按库名 记录连接
		rdbList[rdbName] = rdb
	}
	//把 已记录连接 传递给 实例
	instance.rdb = rdbList[rdbName]

	return instance
}

//生成哈希KEY - 通过类和函数

func (receiver ApiCacheStruct) HKeyByClassMethod(classAndMethod string) string {
	hKey := "api_cache:" + strings.Join(strings.Split(classAndMethod, "\\"), "_")
	return hKey
}

//生成 QUERY KEY/ HEADER KEY - 通过GET请求参数

func (receiver ApiCacheStruct) QueryKeyByRequest(requestQuery map[string]any) string {

	queryKey := ""
	//生成get参数key
	if !maps.IsEmpty(requestQuery) {
		maps.KSort(requestQuery, func(value any, keyName any) {
			val, key := values.ToString(value), values.ToString(keyName)
			//不拼接 缓存键
			if key != "_time" {
				queryKey += "&" + key + "=" + val
			}
		})
		queryKey = "/" + queryKey
	}

	//防止空值
	if maps.IsEmpty(requestQuery) {
		queryKey = "-"
	}

	log.Println("queryKey::", queryKey) //

	return queryKey
}

//缓存方法-闭包结构

type collectClosure func(result *any, meta *any)

//闭包缓存业务返回数据

func (receiver ApiCacheStruct) Collect(hKey string, queryKey string, callback collectClosure, expire time.Duration) (any, any) {
	var result, meta any

	result, meta = receiver.GetDataByMineKey(hKey, queryKey)

	//无缓存泽添加
	if result == nil || receiver.QueryMap["_time"] == "1" {
		log.Println("api-realtime:", hKey+queryKey) //

		//传递对象引用, 让闭包逻辑修改实参
		callback(&result, &meta)

		//设置db集合全局信息
		receiver.setDbInfo(hKey, expire)
		//添加缓存
		receiver.SetDataByMineKey(hKey, queryKey, result, meta, expire)
		//更新db集合全局信息
		receiver.updateDbInfo(hKey)

	} else {
		log.Println("api-cache:", hKey+queryKey) //
	}

	return result, meta
}

//获取数据

func (receiver ApiCacheStruct) GetDataByMineKey(hKey string, queryKey string) (any, any) {
	var apiJson = &jsonStruct{}

	//获取子数据
	ctx := receiver.ctx
	jsonStr, _ := receiver.rdb.HGet(ctx, hKey, queryKey).Result()
	if len(jsonStr) > 0 {

		var err error
		err = json.Unmarshal([]byte(jsonStr), apiJson)
		if err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}

	}
	return apiJson.Data, apiJson.Meta
}

//保存数据

func (receiver ApiCacheStruct) SetDataByMineKey(hKey string, queryKey string, data any, meta any, expire time.Duration) any {
	var apiJson = &jsonStruct{}
	apiJson.Data = data
	apiJson.Meta = meta

	//任意对象 转 json文本
	jsonStr := JsonTool.Encode(apiJson)

	//添加子数据
	ctx := receiver.ctx
	err := receiver.rdb.HSet(ctx, hKey, queryKey, jsonStr).Err()
	if err != nil {
		log.Println("err::", err.Error())
	}
	//不过期,则不设置
	if expire != 0 {
		receiver.rdb.Expire(ctx, hKey, expire)
	}
	return err
}

//设置db集合全局信息

func (receiver ApiCacheStruct) setDbInfo(hKey string, expire time.Duration) {
	ctx := receiver.ctx
	hKeysArr, _ := receiver.rdb.HKeys(ctx, hKey).Result()
	if arrays.IsEmpty(hKeysArr) {
		//哈希键不存在时,需要先设置过期时间, 作用于所有子键.
		receiver.rdb.HSet(ctx, hKey, "db_total", 0)
		receiver.rdb.HSet(ctx, hKey, "db_expire", expire/time.Second)
		receiver.rdb.HSet(ctx, hKey, "db_create_time", time.DateTime)
		receiver.rdb.HSet(ctx, hKey, "db_update_time", time.DateTime)
		receiver.rdb.Expire(ctx, hKey, expire)
	}
}

//更新db集合全局信息

func (receiver ApiCacheStruct) updateDbInfo(hKey string) {
	ctx := receiver.ctx
	hKeysArr, _ := receiver.rdb.HKeys(ctx, hKey).Result()
	dbTotal := len(hKeysArr) - 4
	if !arrays.IsEmpty(hKeysArr) {
		//哈希键存在时,子数据添加.
		receiver.rdb.HMSet(ctx, hKey, "db_total", dbTotal)
		receiver.rdb.HMSet(ctx, hKey, "db_update_time", time.DateTime)
	}
}

//缓存集合筛选-返回结构

type Result struct {
	Data map[string]string
	Meta struct {
		Count int
	}
}

//缓存集合筛选 : apiCache.GetCollect(hKey, "&_page=1")

func (receiver ApiCacheStruct) GetCollect(hKey string, queryKey string) *Result { //
	result := &Result{}
	//
	ctx := receiver.ctx
	allData, _ := receiver.rdb.HGetAll(ctx, hKey).Result()

	//准备获取的 部分键名
	tempKeysMap := make(map[string]string, 0)

	//获取部分键
	maps.Walk(allData, func(value any, keyName any) {
		val, key := values.ToString(value), values.ToString(keyName)
		//log.Println("val, key::", val, key) //
		//正则匹配 连接参数
		wildcard := ".*"
		regex := regexp.MustCompile(wildcard + "(" + queryKey + ")" + wildcard)
		match := regex.FindString(key)
		if len(match) > 0 {
			tempKeysMap[key] = val
		}
	})
	if !maps.IsEmpty(tempKeysMap) {
		result.Data = tempKeysMap
		result.Meta.Count = len(tempKeysMap)
	}

	return result
}

//缓存集合清理

func (receiver ApiCacheStruct) DropStore(hKey string) {
	ctx := receiver.ctx
	receiver.rdb.Expire(ctx, hKey, -1)
	return
}

//缓存集合-部分键清理 : apiCache.DropCollect(hKey, "&_page=1")

func (receiver ApiCacheStruct) DropCollect(hKey string, queryKey string) []string {
	ctx := receiver.ctx

	//准备删除的 部分键名
	tempKeysArr := make([]string, 0)

	//删除部分键
	hKeysArr, _ := receiver.rdb.HKeys(ctx, hKey).Result()
	if !arrays.IsEmpty(hKeysArr) {
		arrays.Walk(hKeysArr, func(value any, index any) {
			val := values.ToString(value)
			//正则匹配 连接参数
			wildcard := ".*"
			regex := regexp.MustCompile(wildcard + "(" + queryKey + ")" + wildcard)
			match := regex.FindString(val)
			if len(match) > 0 {
				tempKeysArr = append(tempKeysArr, val)
			}
		})

		if !arrays.IsEmpty(tempKeysArr) {
			log.Println("tempKeysArr::", tempKeysArr) //
			//批量删除
			receiver.rdb.HDel(ctx, hKey, tempKeysArr...)
			//更新db集合全局信息
			receiver.updateDbInfo(hKey)
		}

	}

	return hKeysArr
}
