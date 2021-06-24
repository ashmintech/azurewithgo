package data

import (
	"log"
)

const (
	ContainerName      = "devicefiles"
	DeviceFileCollName = "devicefiles"
)

type File struct {
	FileName   string `json:"filename"`
	FileDesc   string `json:"filedesc"`
	FileType   string `json:"filetype"`
	FileDate   string `json:"filedate"`
	FilePath   string `json:"filepath"`
	FileHeader string `json:"fileHeader"`
}

type Files []File

var FileList = []File{}

func GetFiles() Files {

	mcoll := GetCollection(DeviceFileCollName)

	err = mcoll.Find(nil).Iter().All(&FileList)
	if err != nil {
		log.Println("Error while querying the Collection:\n", err)
		return nil
	}
	return FileList
}
