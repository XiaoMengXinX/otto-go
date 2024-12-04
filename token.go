package otto

import (
	"fmt"
	"github.com/go-ego/gpy"
	"path/filepath"
	"regexp"
	"strings"
)

func processTokens(input []string) ([]string, error) {
	var inputFiles []string
	for _, fileName := range input {
		if isYSDD(fileName) {
			filePath := filepath.Join("token/multi", getYSDD(fileName)+".wav")
			inputFiles = append(inputFiles, filePath)
			continue
		} else if isSingle(fileName) {
			filePath := filepath.Join("token/single", getSingle(fileName)+".wav")
			inputFiles = append(inputFiles, filePath)
			continue
		} else if replacements, exists := dict[fileName]; exists {
			for _, replacement := range replacements {
				filePath := filepath.Join("token/single", replacement+".wav")
				inputFiles = append(inputFiles, filePath)
			}
		} else {
			filePath := filepath.Join("token/single", fileName+".wav")
			inputFiles = append(inputFiles, filePath)
		}
	}
	fmt.Println(inputFiles)
	for i := 0; i < 2; i++ {
		inputFiles = append(inputFiles, "token/empty.wav")
	}
	return inputFiles, nil
}

func replaceYSDD(s string) string {
	var m []string
	for k, v := range ysddDict {
		for _, value := range v {
			m = append(m, value, fmt.Sprintf("[%s]", k))
		}
	}
	return strings.NewReplacer(m...).Replace(s)
}

func isYSDD(s string) bool {
	for k, _ := range ysddDict {
		if fmt.Sprintf("[%s]", k) == s {
			return true
		}
	}
	return false
}

func isSingle(s string) bool {
	re := regexp.MustCompile(`\((.*?)\)`)
	return re.MatchString(s)
}

func getYSDD(s string) string {
	re := regexp.MustCompile(`\[(.*?)\]`)
	matches := re.FindStringSubmatch(s)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func getSingle(s string) string {
	re := regexp.MustCompile(`\((.*?)\)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func pinyin(s string) []string {
	var py []string
	pys := gpy.Pinyin(s, gpy.Args{})
	for _, p := range pys {
		py = append(py, p...)
	}
	return py
}

func splitStr(s string) []string {
	re := regexp.MustCompile(`(\p{Han})|([a-zA-Z0-9]+)|(\[.*?\])|(\(.*?\))|(\p{P})|(\W)`)
	return re.FindAllString(s, -1)
}

func isChinese(s string) bool {
	re := regexp.MustCompile(`^\p{Han}+$`)
	return re.MatchString(s)
}
