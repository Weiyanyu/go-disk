package vo

import "time"

type GetFileMetaReq struct {
	FileHash string `uri:"file_hash" bind:"required"`
	Username string `form:"username" bind:"required"`
}

type UpdateFileMetaReq struct {
	Username string `form:"username" bind:"required"`
	FileHash string `form:"file_hash" bind:"required"`
	Filename string `form:"filename" bind:"required"`
	Status string `form:"status"`
}

type DeleteFileReq struct {
	FileHash string `form:"file_hash" bind:"required"`
	Username string `form:"username" bind:"required"`
	Filename string `form:"filename" bind:"filename"`
}

type UserFileResp struct {
	Username string `json:"username"`
	Filename string `json:"filename"`
	FileHash string `json:"file_hash"`
	FileSize int64 `json:"file_size"`
	UploadAt time.Time `json:"upload_at"`
	LastUpdate time.Time `json:"last_update"`
}

type UserFileReq struct {
	Username string `form:"username" bind:"required"`
	Limit int `form:"limit" bind:"required"`
}

