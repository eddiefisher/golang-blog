package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dev/utils"
)

//PostsIndex show all posts
func PostsIndex(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo(utils.FunctionName(), r)
	if r.URL.Path != "/" {
		utils.PageNotFound(w, r, http.StatusNotFound)
		return
	}
	utils.RenderTemplate("posts", "index").ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title":   "All Posts",
		"Content": "Lorem Ipsum",
	})
}

//PostsShow show post
func PostsShow(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo(utils.FunctionName(), r)
	if !strings.HasPrefix(r.URL.Path, "/post") {
		utils.PageNotFound(w, r, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.URL.Path[len("/post/"):])
	if err != nil {
		utils.PageNotFound(w, r, http.StatusNotFound)
		return
	}
	utils.RenderTemplate("posts", "show").ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title":   "Post " + strconv.Itoa(id),
		"Content": "show: " + strconv.Itoa(id),
	})
}
