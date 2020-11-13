package dao

import (
	log "code.inke.cn/BackendPlatform/golang/logging"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/conf"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/model"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

// Dao represents data access object
type Dao struct {
	c *conf.Config
}

func New(c *conf.Config) *Dao {
	return &Dao{
		c: c,
	}
}

// Ping check db resource status
func (d *Dao) Ping(ctx context.Context) error {
	return nil
}

// Close release resource
func (d *Dao) Close() error {
	return nil
}

//数据库配置
const (
	userName = "inke_db_user"
	password = "DRhBob097)+"
	ip       = "10.100.13.143"
	port     = "3306"
	dbName   = "media_lm6_record"
)

var DB *sql.DB

//注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(200)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		log.Errorf("opon database fail:%v", err)
		return
	}
	log.Info("mysql connnect success")
}


func InsertOriginalMysql(mysqlInfos *[]model.MysqOriginalInfo, tblName string) bool {
	for _, mysqlInfo := range *mysqlInfos {
		//查询是否有重复数据
		var repeatData model.MysqOriginalInfo
		querySql := fmt.Sprintf("SELECT id FROM %s WHERE  ymd = ? and group1 = ? and group2 = ? and group3 = ?", tblName)
		err := DB.QueryRow(querySql, mysqlInfo.Ymd, mysqlInfo.Country).Scan(&repeatData.Id)
		//&repeatData.Ymd, &repeatData.All, &repeatData.Group1, &repeatData.Group1_Num, &repeatData.Group2, &repeatData.Group2_Num,
		//				&repeatData.Group3, &repeatData.Group3_Num, &repeatData.Value, &repeatData.Remark, &repeatData.Event_Time
		if err != nil {
			log.Infof("Select fail:%v", err)
		}
		//log.Infof("Select Success:%s", err)
		//删除重复数据
		if repeatData != (model.MysqOriginalInfo{}) {
			//开启事务
			tx, err := DB.Begin()
			if err != nil {
				log.Errorf("Begin fail:%v", err)
			}
			//准备sql语句
			querySql = fmt.Sprintf("DELETE FROM %s WHERE id = ?", tblName)
			stmt, err := tx.Prepare(querySql)
			if err != nil {
				log.Errorf("Sql init fail:%v", err)
				return false
			}
			//设置参数以及执行sql语句
			_, err = stmt.Exec(repeatData.Id)
			if err != nil {
				log.Errorf("Sql init fail:%v", err)
				return false
			}
			//提交事务
			tx.Commit()
			//log.Infof("Delete Success:%s", res)
		}
		//插入数据
		tx, err := DB.Begin()
		if err != nil {
			log.Errorf("tx fail:%v", err)
			return false
		}
		//准备sql语句
		querySql = fmt.Sprintf("INSERT INTO %s (`ymd`,`all`, `group1`, `group1_num`, `group2`, `group2_num`, `group3`, `group3_num`, `value`, `remark`, `event_time`) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tblName)
		stmt, err := tx.Prepare(querySql)
		if err != nil {
			log.Errorf("Prepare fail:%v", err)
			return false
		}

		//将参数传递到sql语句中并且执行
		res, err := stmt.Exec(mysqlInfo.Ymd, mysqlInfo.All, mysqlInfo.Event_Time)
		if err != nil {
			log.Errorf("Exec fail:%v", err)
			return false
		}
		//将事务提交
		tx.Commit()
		//获得上一个插入自增的id
		log.Infof("Insert Success:%s", res)
	}

	return true
}

func InsertHandleMysql(mysqlInfos *[]model.MysqHandleInfo, tblName string) bool {
	for _, mysqlInfo := range *mysqlInfos {
		//查询是否有重复数据
		var repeatData model.MysqUntreatedInfo
		querySql := fmt.Sprintf("SELECT id FROM %s WHERE  ymd = ? and group1 = ? and group2 = ? and group3 = ?", tblName)
		err := DB.QueryRow(querySql, mysqlInfo.Ymd, ).Scan(&repeatData.Id)
		//&repeatData.Ymd, &repeatData.All, &repeatData.Group1, &repeatData.Group1_Num, &repeatData.Group2, &repeatData.Group2_Num,
		//				&repeatData.Group3, &repeatData.Group3_Num, &repeatData.Value, &repeatData.Remark, &repeatData.Event_Time
		if err != nil {
			log.Infof("Select fail:%v", err)
		}
		//log.Infof("Select Success:%s", err)
		//删除重复数据
		if repeatData != (model.MysqUntreatedInfo{}) {
			//开启事务
			tx, err := DB.Begin()
			if err != nil {
				log.Errorf("Begin fail:%v", err)
			}
			//准备sql语句
			querySql = fmt.Sprintf("DELETE FROM %s WHERE id = ?", tblName)
			stmt, err := tx.Prepare(querySql)
			if err != nil {
				log.Errorf("Sql init fail:%v", err)
				return false
			}
			//设置参数以及执行sql语句
			_, err = stmt.Exec(repeatData.Id)
			if err != nil {
				log.Errorf("Sql init fail:%v", err)
				return false
			}
			//提交事务
			tx.Commit()
			//log.Infof("Delete Success:%s", res)
		}
		//插入数据
		tx, err := DB.Begin()
		if err != nil {
			log.Errorf("tx fail:%v", err)
			return false
		}
		//准备sql语句
		querySql = fmt.Sprintf("INSERT INTO %s (`ymd`,`all`, `group1`, `group1_num`, `group2`, `group2_num`, `group3`, `group3_num`, `value`, `remark`, `event_time`) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tblName)
		stmt, err := tx.Prepare(querySql)
		if err != nil {
			log.Errorf("Prepare fail:%v", err)
			return false
		}

		//将参数传递到sql语句中并且执行
		res, err := stmt.Exec(mysqlInfo.Ymd, mysqlInfo.All, mysqlInfo.Event_Time)
		if err != nil {
			log.Errorf("Exec fail:%v", err)
			return false
		}
		//将事务提交
		tx.Commit()
		//获得上一个插入自增的id
		log.Infof("Insert Success:%s", res)
	}

	return true
}

func InsertMysql(MysqlInfos *[]model.MysqUntreatedInfo, tblName string) bool {
	for _, mysqlInfo := range *MysqlInfos {
		//查询是否有重复数据
		var repeatData model.MysqUntreatedInfo
		querySql := fmt.Sprintf("SELECT id FROM %s WHERE  ymd = ? and group1 = ? and group2 = ? and group3 = ?", tblName)
		err := DB.QueryRow(querySql, mysqlInfo.Ymd, mysqlInfo.Group1, mysqlInfo.Group2, mysqlInfo.Group3).Scan(&repeatData.Id)
		//&repeatData.Ymd, &repeatData.All, &repeatData.Group1, &repeatData.Group1_Num, &repeatData.Group2, &repeatData.Group2_Num,
		//				&repeatData.Group3, &repeatData.Group3_Num, &repeatData.Value, &repeatData.Remark, &repeatData.Event_Time
		if err != nil {
			log.Infof("Select fail:%v", err)
		}
		//log.Infof("Select Success:%s", err)
		//删除重复数据
		if repeatData != (model.MysqUntreatedInfo{}) {
			//开启事务
			tx, err := DB.Begin()
			if err != nil {
				log.Errorf("Begin fail:%v", err)
			}
			//准备sql语句
			querySql = fmt.Sprintf("DELETE FROM %s WHERE id = ?", tblName)
			stmt, err := tx.Prepare(querySql)
			if err != nil {
				log.Errorf("Sql init fail:%v", err)
				return false
			}
			//设置参数以及执行sql语句
			_, err = stmt.Exec(repeatData.Id)
			if err != nil {
				log.Errorf("Sql init fail:%v", err)
				return false
			}
			//提交事务
			tx.Commit()
			//log.Infof("Delete Success:%s", res)
		}
		//插入数据
		tx, err := DB.Begin()
		if err != nil {
			log.Errorf("tx fail:%v", err)
			return false
		}
		//准备sql语句
		querySql = fmt.Sprintf("INSERT INTO %s (`ymd`,`all`, `group1`, `group1_num`, `group2`, `group2_num`, `group3`, `group3_num`, `value`, `remark`, `event_time`) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tblName)
		stmt, err := tx.Prepare(querySql)
		if err != nil {
			log.Errorf("Prepare fail:%v", err)
			return false
		}

		//将参数传递到sql语句中并且执行
		res, err := stmt.Exec(mysqlInfo.Ymd, mysqlInfo.All, mysqlInfo.Group1, mysqlInfo.Group1_Num, mysqlInfo.Group2, mysqlInfo.Group2_Num, mysqlInfo.Group3, mysqlInfo.Group3_Num, mysqlInfo.Value1, mysqlInfo.Remark, mysqlInfo.Event_Time)
		if err != nil {
			log.Errorf("Exec fail:%v", err)
			return false
		}
		//将事务提交
		tx.Commit()
		//获得上一个插入自增的id
		log.Infof("Insert Success:%s", res)
	}

	return true
}
