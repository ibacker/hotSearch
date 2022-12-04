package utils

import (
	"fmt"
	"hotSearch/model"
	"io/ioutil"
	"os"
	"regexp"
)

func CreateReadMe(data []model.Question) {
	fileName := "README.md"
	if file, err := ioutil.ReadFile(fileName); err == nil {
		reg, _ := regexp.Compile(`<!-- BEGIN -->[\W\w]*<!-- END -->`)
		allString := reg.ReplaceAll(file, []byte(CreateList(data)))
		if err := ioutil.WriteFile(fileName, allString, os.ModePerm); err != nil {
			fmt.Println("README写入错误")
			fmt.Println(err)
		}
	}
}

func CreateArchives(data []model.Question, fileName string) {
	filePath := fmt.Sprintf("./archives/%v.md", fileName)
	if file, err := os.Create(filePath); err == nil {
		defer file.Close()
		content := fmt.Sprintf("# %v\n 共%v条\n %v", fileName, len(data), CreateList(data))
		file.WriteString(content)
	}
}

func CreateList(data []model.Question) string {
	var word string
	for _, v := range data {
		word += fmt.Sprintf("1. [%v](%v)\n", v.Title, v.Url)
	}
	template := fmt.Sprintf("<!-- BEGIN -->\n\n%v\n<!-- END -->", word)
	return template
}
