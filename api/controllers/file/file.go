package file

import (
	gt "github.com/dreamlu/go-tool"
	File "github.com/dreamlu/go-tool/tool/file"
	"github.com/dreamlu/go-tool/tool/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 单文件上传
// use gin upload file
func UpoadFile(u *gin.Context) {

	name := u.PostForm("name") //指定文件名
	file, err := u.FormFile("file")
	if err != nil {
		gt.Logger().Error(err.Error())
		u.JSON(http.StatusOK, result.MapData{Status: result.CodeError, Msg: err.Error()})
	}
	upFile := File.File{
		Name: name,
	}
	path, err := upFile.GetUploadFile(file)
	if err != nil {
		gt.Logger().Error(err.Error())
	}
	u.JSON(http.StatusOK, result.GetSuccess(map[string]interface{}{"path": path}))
}
