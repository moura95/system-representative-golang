package file

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/util"
)

func (f *File) upload(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorReadFile.Error()))
		return
	}

	loc, err := util.UploadFile(&file, &f.Config, util.GenerateFilename(representativeID, header.Filename, ctx.Param("dir")))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorUploadFile.Error()))
		return
	}
	data := map[string]string{
		"location": loc,
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse(200, data, ""))
	return
}

func (f *File) delete(ctx *gin.Context) {
	_ = util.DeleteFile(ctx.Param("filename"))

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
