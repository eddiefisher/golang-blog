// Copyright 2017, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// main.go [created: Tue, 30 May 2017]

package main

import (
	"net/http"

	"./controller"

	"github.com/Sirupsen/logrus"
	"github.com/dev/dbcon"
)

func main() {
	db := dbcon.Connection()
	defer db.Close()

	http.HandleFunc("/", controller.PostsIndex)
	http.HandleFunc("/post/", controller.PostsShow)

	logrus.Fatal(http.ListenAndServe(":3000", nil))
}
