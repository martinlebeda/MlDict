package service

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

type Data []Dict

type Dict struct {
	Dict  string `json:"dict"`
	Terms []Term `json:"terms"`
}

type Term struct {
	Term        string `json:"term"`
	Explanation string `json:"explanation"`
}

func ListDict() []string {
	db := _OpenDB()

	var result []string
	rows := _RunSelect(db, "SELECT distinct id_dict from dictionary order by id_dict")
	for rows.Next() {
		var dict string
		err := rows.Scan(&dict)
		if err == nil {
			result = append(result, dict)
		}
	}

	rows.Close()
	db.Close()

	return result
}

func QueryDict(query string) Data {
	db := _OpenDB()

	queryVar := query + "*"
	dictQuery := "SELECT dict, key, group_concat(expl, '; ') expl from dictft WHERE key MATCH '" + queryVar + "' group by dict, key order by dict, key, expl"
	rows := _RunSelect(db, dictQuery)
	var key string
	var expl string
	var dict string

	var cDict = ""

	result := Data{}
	rDict := Dict{}

	for rows.Next() {
		err := rows.Scan(&dict, &key, &expl)
		if err == nil {

			if dict != cDict {
				if cDict != "" {
					result = append(result, rDict)
				}
				rDict = Dict{Dict: dict, Terms: []Term{}}
				cDict = dict
			}

			term := Term{key, expl}
			rDict.Terms = append(rDict.Terms, term)
		}
	}
	if cDict != "" {
		result = append(result, rDict)
	}

	rows.Close() //good habit to close
	db.Close()

	return result
}

func _RunSelect(db *sql.DB, dictQuery string) *sql.Rows {
	rows, err := db.Query(dictQuery)
	if err != nil {
		panic(err)
	}
	return rows
}

func _OpenDB() *sql.DB {
	homeDir := os.Getenv("HOME")
	dbpath := filepath.Join(homeDir, ".dictionary.db")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		panic(err)
	}
	return db
}
