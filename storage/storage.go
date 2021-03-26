package storage

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/xonoxitron/gorilla/utils"
)

func Path(name string) string {
	ex, err := os.Getwd()
	utils.ErrorCheck(err)
	return filepath.Dir(ex) + "/gorilla/data/" + name + ".dat"
}

func Create(name string) {
	filePath := Path(name)
	if !utils.Exists(filePath) {
		err := ioutil.WriteFile(filePath, []byte(""), 0644)
		utils.ErrorCheck(err)
	}
}

func Setup() {
	Create("subscribers")
	Create("assets")
	Create("tickers")
}

func Get(name string) string {
	content, err := ioutil.ReadFile(Path(name))
	utils.ErrorCheck(err)
	return string(content)
}

func Update(name string, content string, overwrite bool) {
	if overwrite {
		file, err := os.Create(Path(name))
		utils.ErrorCheck(err)
		defer file.Close()
		file.WriteString(content)
	} else {
		file, err := os.OpenFile(Path(name), os.O_APPEND|os.O_WRONLY, 0644)
		utils.ErrorCheck(err)
		defer file.Close()
		file.WriteString(content)
	}
}
