package sqlite

import (
	_ "github.com/mattn/go-sqlite3"
	. "../../data"
	"database/sql"
	. "../error"
)

// Initialize database
func InitDB() {
	db, err := sql.Open(DRIVER_NAME, DB_NAME)
	CheckErr(err)
	// Create table.
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS " + TB_VIDEO_DATA + " (" +
		ID + " INTEGER PRIMARY KEY AUTOINCREMENT," +
		AID + " INTEGER NOT NULL," +
		VIEW + " INTEGER DEFAULT -1," +
		DANMAKU + " INTEGER DEFAULT -1," +
		REPLY + " INTEGER DEFAULT -1," +
		FAVORITE + " INTEGER DEFAULT -1," +
		COIN + " INTEGER DEFAULT -1," +
		SHARE + " INTEGER DEFAULT -1," +
		NOW_RANK + " INTEGER DEFAULT -1," +
		HIS_RANK + " INTEGER DEFAULT -1," +
		LIKE + " INTEGER DEFAULT -1," +
		NO_REPRINT + " INTEGER DEFAULT -1," +
		COPYRIGHT + " INTEGER DEFAULT -1" +
		");")
	CheckErr(err)
	_, err = stmt.Exec()
	CheckErr(err)
}

// Insert a data to sqlite3 database
func InsertData(data *Data) int64 {
	db, err := sql.Open(DRIVER_NAME, DB_NAME)
	CheckErr(err)
	stmt, err := db.Prepare("INSERT INTO " + TB_VIDEO_DATA + " (" +
		AID + ", " +
		VIEW + ", " +
		DANMAKU + ", " +
		REPLY + ", " +
		FAVORITE + ", " +
		COIN + ", " +
		SHARE + ", " +
		NOW_RANK + ", " +
		HIS_RANK + ", " +
		LIKE + ", " +
		NO_REPRINT + ", " +
		COPYRIGHT +
		") VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	CheckErr(err)

	res, err := stmt.Exec(data.Aid, data.View, data.Danmaku, data.Reply, data.Favorite,
		data.Coin, data.Share, data.NowRank, data.HisRank, data.Like, data.NoReprint, data.Copyright)
	CheckErr(err)

	id, err := res.LastInsertId()
	CheckErr(err)
	return id
}

// Get all data
func GetAllData() []Data {
	db, err := sql.Open(DRIVER_NAME, DB_NAME)
	// 查询数据
	rows, err := db.Query("SELECT * FROM " + TB_VIDEO_DATA)
	CheckErr(err)

	var allData = make([]Data, 0)
	for rows.Next() {
		var data = Data{}
		err = rows.Scan(&data.Aid, &data.View, &data.Danmaku, &data.Reply, &data.Favorite,
			&data.Coin, &data.Share, &data.NowRank, &data.HisRank, &data.Like, &data.NoReprint, &data.Copyright)
		CheckErr(err)
		allData = append(allData, data)
	}
	return allData
}