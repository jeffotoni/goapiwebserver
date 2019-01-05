// Front-end in Go server
// @jeffotoni
// 04/01/2019

package util

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

//  Exist only file
func FileExist(name string) bool {

	//if _, err := os.Stat(name); os.IsNotExist(err) {
	if stat, err := os.Stat(name); err == nil && !stat.IsDir() {

		return true
	}

	return false
}

func DirExist(path string) bool {

	//if _, err := os.Stat(path); err == nil {
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {

		return true
	}

	return false
}

func CreateDirIfNotExist(dir string) bool {

	if !DirExist(dir) {

		err := os.MkdirAll(dir, 0755)

		if err != nil {
			//log.Save.Println(err)
			return false
		}
	}

	return true
}

func CreateDirIfNotExistNotFile(dir string) {

	vet := strings.Split(dir, "/")
	tamanho_i := len(vet)
	tamanho_f := tamanho_i - 1
	var dirNew string
	//log.Println(len(vet))
	//log.Println(vet[0:tamanho_f])
	if tamanho_i > 1 {

		dirNew = strings.Join(vet[0:tamanho_f], "/")

		//criar diretorio ou nao
		if !DirExist(dirNew) {

			os.MkdirAll(dirNew, 0755)
		}
	}
}

// removendo todos os diretorios
// e arquivos de um diretorio pai
func RemoveAllDir(pathlocal string) {

	dir, _ := ioutil.ReadDir(pathlocal)
	for _, d := range dir {
		//log.Println(path.Join([]string{pathlocal, d.Name()}...))
		os.RemoveAll(path.Join([]string{pathlocal, d.Name()}...))
	}
}

func DeleteFile(path string) {
	// delete file
	os.Remove(path)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	//fmt.Println("==> done deleting file")
}
