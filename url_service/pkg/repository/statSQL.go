package repository

import (
	"database/sql"
	"fmt"
	"url_service"
)

type StatSQL struct {
	db *sql.DB
}

func NewStatSQL(db *sql.DB) *StatSQL {
	return &StatSQL{
		db: db,
	}
}

func (r *URLSQL) AddStat(stat url_service.Stats, statDetail url_service.StatsDetail) (int, error) {
	var idResult int64

	id := r.CheckStat(stat)

	fmt.Println(id)

	if id != 0 {
		query := fmt.Sprintf("UPDATE %s SET counttransition=`counttransition` + 1 WHERE id=%d", statTable, id)

		result, err := r.db.Exec(query)

		if err != nil {
			return 0, err
		}

		idResult, err = result.RowsAffected()

		if err != nil {
			return 0, err
		}

		statDetail.Statid = int(id)

		r.AddStatDetail(statDetail)

		fmt.Println(statDetail)
	} else {
		query := fmt.Sprintf("INSERT INTO %s (url_id, counttransition, date, dateDB) VALUES (%d, %d, %d, '%s')", statTable, stat.Url_id, stat.Counttransition, stat.Date, stat.DateDB)

		result, err := r.db.Exec(query)

		if err != nil {
			return 0, err
		}

		idResult, err = result.LastInsertId()

		if err != nil {
			return 0, err
		}

		statDetail.Statid = int(idResult)

		r.AddStatDetail(statDetail)
	}

	return int(idResult), nil
}

func (r *URLSQL) AddStatDetail(statDetail url_service.StatsDetail) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (statid, city, browser, os, ismobile) VALUES (%d, %d, '%s', '%s', %v)", statDetailTable, statDetail.Statid, statDetail.City, statDetail.Browser, statDetail.Os, statDetail.IsMobile)

	_, err := r.db.Exec(query)

	fmt.Println(err)

	if err != nil {
		return 0, err
	}

	fmt.Println("ДОбавлена статистика детальная")

	return statDetail.Statid, nil
}

func (r *URLSQL) LoadIdRegion(stat url_service.RegionsCode) int {
	query := fmt.Sprintf("SELECT id FROM ru WHERE code='%s'", stat.Code)

	rows, err := r.db.Query(query)

	if err != nil {
		return 0
	}

	for rows.Next() {
		rows.Scan(&stat.Id)
	}

	return stat.Id
}

func (r *URLSQL) CheckStat(stat url_service.Stats) int {
	query := fmt.Sprintf("SELECT id FROM %s WHERE url_id=%d AND dateDB='%s'", statTable, stat.Url_id, stat.DateDB)

	rows, err := r.db.Query(query)

	if err != nil {
		return 0
	}

	for rows.Next() {
		rows.Scan(&stat.Id)
	}

	return stat.Id
}

func (r *URLSQL) GetStatURLid(urlId int) ([]url_service.Stats, error) {
	var statList []url_service.Stats

	query := fmt.Sprintf("SELECT id, counttransition, date FROM %s WHERE url_id=%d", statTable, urlId)

	rows, err := r.db.Query(query)

	for rows.Next() {
		var singleList url_service.Stats

		errScan := rows.Scan(&singleList.Id, &singleList.Counttransition, &singleList.Date)

		if errScan != nil {
			return statList, errScan
		}

		statList = append(statList, singleList)
	}

	return statList, err
}

//select sd.browser, count(sd.browser) from stat s join statdetail sd on s.id=sd.statid where s.url_id=2 group by sd.browser;
func (r *URLSQL) GetStatBrowser(urlId int) ([]url_service.StatsBrowser, error) {
	var statList []url_service.StatsBrowser

	query := fmt.Sprintf("SELECT sd.browser, COUNT(sd.browser) AS count FROM stat s JOIN statdetail sd ON s.id=sd.statid WHERE s.url_id=%d GROUP BY sd.browser", urlId)

	rows, err := r.db.Query(query)

	for rows.Next() {
		var singleList url_service.StatsBrowser

		errScan := rows.Scan(&singleList.Browser, &singleList.Count)

		if errScan != nil {
			return statList, errScan
		}

		statList = append(statList, singleList)
	}

	return statList, err
}

func (r *URLSQL) GetOsURL(urlId int) ([]url_service.StatsOs, error) {
	var statList []url_service.StatsOs

	query := fmt.Sprintf("SELECT DISTINCT sd.os FROM statdetail sd JOIN stat s ON sd.statid=s.id WHERE s.url_id=%d", urlId)

	rows, err := r.db.Query(query)

	for rows.Next() {
		var singleList url_service.StatsOs

		singleList.Count = 0

		errScan := rows.Scan(&singleList.Os)

		if errScan != nil {
			return statList, errScan
		}

		statList = append(statList, singleList)
	}

	return statList, err
}

func (r *URLSQL) GetCountType(urlId int) ([]url_service.StatCountType, error) {
	var statList []url_service.StatCountType
	ismobile := true

	query := fmt.Sprintf("(select count(os) as count from statdetail sd join stat s on sd.statid=s.id where ismobile=1 and url_id=%d) union all (select count(os) from statdetail sd join stat s on sd.statid=s.id where ismobile=0 and url_id=%d)", urlId, urlId)

	rows, err := r.db.Query(query)

	for rows.Next() {
		var singleList url_service.StatCountType

		if ismobile {
			singleList.IsMobile = true
			ismobile = false
		}

		errScan := rows.Scan(&singleList.Count)

		if errScan != nil {
			return statList, errScan
		}

		statList = append(statList, singleList)
	}

	return statList, err
}

func (r *URLSQL) GetCountIsMobile(urlId int) ([]url_service.StatCountIsMobile, error) {
	var statList []url_service.StatCountIsMobile

	query := fmt.Sprintf("select sd.os, count(sd.ismobile) as count from statdetail sd join stat s on sd.statid=s.id where sd.ismobile=1 and s.url_id=%d group by os", urlId)

	rows, err := r.db.Query(query)

	for rows.Next() {
		var singleList url_service.StatCountIsMobile

		errScan := rows.Scan(&singleList.Os, &singleList.Count)

		if errScan != nil {
			return statList, errScan
		}

		statList = append(statList, singleList)
	}

	return statList, err
}

func (r *URLSQL) GetCountIsTab(urlId int) ([]url_service.StatCountIsMobile, error) {
	var statList []url_service.StatCountIsMobile

	query := fmt.Sprintf("select sd.os, count(sd.ismobile) as count from statdetail sd join stat s on sd.statid=s.id where sd.ismobile=0 and s.url_id=%d group by os", urlId)

	rows, err := r.db.Query(query)

	for rows.Next() {
		var singleList url_service.StatCountIsMobile

		errScan := rows.Scan(&singleList.Os, &singleList.Count)

		if errScan != nil {
			return statList, errScan
		}

		statList = append(statList, singleList)
	}

	return statList, err
}

func (r *URLSQL) GetCountRegion(urlId int) ([]url_service.RegionsCode, error) {
	var statList []url_service.RegionsCode

	query := fmt.Sprintf("select r.name, r.code, count(sd.statid) as value from statdetail sd join stat s on sd.statid=s.id join ru r on sd.city=r.id where s.url_id=%d group by r.name, r.code;", urlId)

	rows, err := r.db.Query(query)

	for rows.Next() {
		var singleList url_service.RegionsCode

		errScan := rows.Scan(&singleList.Name, &singleList.Code, &singleList.Id)

		if errScan != nil {
			return statList, errScan
		}

		statList = append(statList, singleList)
	}

	return statList, err
}
