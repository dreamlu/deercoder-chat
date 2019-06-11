package file

import (
	File "github.com/dreamlu/go-tool/util/file"
	"github.com/dreamlu/go-tool/util/lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 单文件上传
// use gin upload file
func UpoadFile(u *gin.Context) {

	fname := u.PostForm("fname") //指定文件名
	file, err := u.FormFile("file")
	if err != nil {
		u.JSON(http.StatusOK, lib.MapData{Status: lib.CodeFile, Msg: err.Error()})
	}
	upFile := File.File{}
	path := upFile.GetUploadFile(file, fname)
	u.JSON(http.StatusOK, map[string]interface{}{lib.Status: lib.CodeFile, lib.Msg: lib.MsgFile, "path": path})
}
