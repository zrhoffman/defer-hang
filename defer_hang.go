package main

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import (
	"database/sql"
	_ "github.com/lib/pq"
	"runtime"
)

func main() {
	for _ = range []int{0} {
		defer func() {
			db, err := sql.Open("postgres", "postgres://user:password@dbhost/dbname")
			if err != nil {
				println("cannot open db")
				return
			}
			defer db.Close()

			tx, err := db.Begin()
			if err != nil {
				println("transaction begin failed: " + err.Error())
				return
			}

			err = tx.Commit()
			if err != nil {
				println("commit failed: " + err.Error())
				return
			}
		}()
	}
	println("About to call runtime.Goexit()")
	runtime.Goexit()
}
