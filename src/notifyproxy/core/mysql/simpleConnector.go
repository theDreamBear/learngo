package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gomcpack/mcpack"
	"gomcpack/npc"
	"notifyproxy/core/define"
	"notifyproxy/core/util"
	"time"
)

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("mysql", "nichao:nichao@tcp(127.0.0.1:3306)/notify?charset=utf8")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	} else {
		if db != nil {
			fmt.Println("db has been initialed already")
		} else {
			fmt.Println("db is none")
		}
	}
	// Open doesn't open a connection. Validate DSN data:
	db.Ping()
}

func InsertToDb(obj *define.NotifyNmqOnlineInfo, rw npc.ResponseWriter) {
	dt := util.TimeStampToYearMonthDay(int64(obj.Utime))
	if err := insertToDb(obj.Uid, obj.Product, obj.Extra, obj.Cmdno, obj.Stime, obj.Utime, dt); err == nil {
		fmt.Println("insert success")
		response := &define.NmqResponse{ErrorNo: 0, ErrorMsg: "success", TransId: obj.TransId}
		rB, _ := mcpack.Marshal(response)
		rw.Write(rB)
	} else {
		fmt.Println("insert error\t", err)
	}
}

func insertToDb(uid uint64, proudct string, extra string, cmdno uint32, stime int32, utime int32, dt int32) error {
	index := uid % 20
	insertSql := fmt.Sprintf("insert into  tblOnlineStat%d (uid, product, extra, cmdno, stime, utime, dt ) values (?, ?, ?, ?, ? ,?, ?) ", index)

	now := time.Now().UnixNano()
	if smt, err := db.Prepare(insertSql); err == nil {

		if _, err := smt.Exec(uid, proudct, extra, cmdno, stime, utime, dt); err != nil {
			fmt.Println(err)
			return err
		} else {
			elapse := time.Now().UnixNano() - now
			fmt.Printf("Prepare time %dus\n", elapse/1000)
			fmt.Println(insertSql)
			return nil
		}
	} else {
		fmt.Println(insertSql)
		return err
	}
}
