/// -------------------------------------------------------------------------------
/// THIS FILE IS ORIGINALLY GENERATED BY redis2go.exe.
/// PLEASE DO NOT MODIFY THIS FILE.
/// -------------------------------------------------------------------------------

package db

import (
	"errors"

	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
	"github.com/gomodule/redigo/redis"
)

// Account : 代表 1 个 redis 对象
type Account struct {
	Key    string
	passwd string

	dirtyDataInAccount               map[string]interface{}
	dirtyDataForStructFiledInAccount map[string]interface{}
	isLoadInAccount                  bool
	dbKeyInAccount                   string
	ddbNameInAccount                 string
	expireInAccount                  uint
}

// NewAccount : NewAccount 的构造函数
func NewAccount(dbName string, key string) *Account {
	return &Account{
		Key:                              key,
		ddbNameInAccount:                 dbName,
		dbKeyInAccount:                   "Account:" + key,
		dirtyDataInAccount:               make(map[string]interface{}),
		dirtyDataForStructFiledInAccount: make(map[string]interface{}),
	}
}

// HasKey : 是否存在 KEY
//          返回值，若访问数据库失败返回-1；若 key 存在返回 1 ，否则返回 0 。
func (objAccount *Account) HasKey() (int, error) {
	db := go_redis_orm.GetDB(objAccount.ddbNameInAccount)
	val, err := redis.Int(db.Do("EXISTS", objAccount.dbKeyInAccount))
	if err != nil {
		return -1, err
	}
	return val, nil
}

// Load : 从 redis 加载数据
func (objAccount *Account) Load() error {
	if objAccount.isLoadInAccount == true {
		return errors.New("alreay load")
	}
	db := go_redis_orm.GetDB(objAccount.ddbNameInAccount)
	val, err := redis.Values(db.Do("HGETALL", objAccount.dbKeyInAccount))
	if err != nil {
		return err
	}
	if len(val) == 0 {
		return go_redis_orm.ERR_ISNOT_EXIST_KEY
	}
	var data struct {
		Passwd string `redis:"passwd"`
	}
	if err := redis.ScanStruct(val, &data); err != nil {
		return err
	}
	objAccount.passwd = data.Passwd
	objAccount.isLoadInAccount = true
	return nil
}

// Save : 保存数据到 redis
func (objAccount *Account) Save() error {
	if len(objAccount.dirtyDataInAccount) == 0 && len(objAccount.dirtyDataForStructFiledInAccount) == 0 {
		return nil
	}
	for k := range objAccount.dirtyDataForStructFiledInAccount {
		_ = k

	}
	db := go_redis_orm.GetDB(objAccount.ddbNameInAccount)
	if _, err := db.Do("HMSET", redis.Args{}.Add(objAccount.dbKeyInAccount).AddFlat(objAccount.dirtyDataInAccount)...); err != nil {
		return err
	}
	if objAccount.expireInAccount != 0 {
		if _, err := db.Do("EXPIRE", objAccount.dbKeyInAccount, objAccount.expireInAccount); err != nil {
			return err
		}
	}
	objAccount.dirtyDataInAccount = make(map[string]interface{})
	objAccount.dirtyDataForStructFiledInAccount = make(map[string]interface{})
	return nil
}

// Delete : 从 redis 删除数据
func (objAccount *Account) Delete() error {
	db := go_redis_orm.GetDB(objAccount.ddbNameInAccount)
	_, err := db.Do("DEL", objAccount.dbKeyInAccount)
	if err == nil {
		objAccount.isLoadInAccount = false
		objAccount.dirtyDataInAccount = make(map[string]interface{})
		objAccount.dirtyDataForStructFiledInAccount = make(map[string]interface{})
	}
	return err
}

// IsLoad : 是否已经从 redis 导入数据
func (objAccount *Account) IsLoad() bool {
	return objAccount.isLoadInAccount
}

// Expire : 向 redis 设置该对象过期时间
func (objAccount *Account) Expire(v uint) {
	objAccount.expireInAccount = v
}

// DirtyData : 获取该对象目前已脏的数据
func (objAccount *Account) DirtyData() (map[string]interface{}, error) {
	for k := range objAccount.dirtyDataForStructFiledInAccount {
		_ = k

	}
	data := make(map[string]interface{})
	for k, v := range objAccount.dirtyDataInAccount {
		data[k] = v
	}
	objAccount.dirtyDataInAccount = make(map[string]interface{})
	objAccount.dirtyDataForStructFiledInAccount = make(map[string]interface{})
	return data, nil
}

// Save2 : 保存数据到 redis 的第 2 种方法
func (objAccount *Account) Save2(dirtyData map[string]interface{}) error {
	if len(dirtyData) == 0 {
		return nil
	}
	db := go_redis_orm.GetDB(objAccount.ddbNameInAccount)
	if _, err := db.Do("HMSET", redis.Args{}.Add(objAccount.dbKeyInAccount).AddFlat(dirtyData)...); err != nil {
		return err
	}
	if objAccount.expireInAccount != 0 {
		if _, err := db.Do("EXPIRE", objAccount.dbKeyInAccount, objAccount.expireInAccount); err != nil {
			return err
		}
	}
	return nil
}

// GetPasswd : 获取字段值
func (objAccount *Account) GetPasswd() string {
	return objAccount.passwd
}

// SetPasswd : 设置字段值
func (objAccount *Account) SetPasswd(value string) {
	objAccount.passwd = value
	objAccount.dirtyDataInAccount["passwd"] = string([]byte(value))
}
