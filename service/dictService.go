// Copyright © 2018 Martin Lebeda <martin.lebeda@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package service

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
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
	dbFileName := viper.GetString("dbfile")
	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		panic(err)
	}
	return db
}
