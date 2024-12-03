package otto

import (
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
		}
		trimmedFileName := strings.TrimSpace(fileName)
		if trimmedFileName != "" {
			if replacements, exists := dict[trimmedFileName]; exists {
				for _, replacement := range replacements {
					filePath := filepath.Join("token/single", replacement+".wav")
					inputFiles = append(inputFiles, filePath)
				}
			} else {
				filePath := filepath.Join("token/single", trimmedFileName+".wav")
				inputFiles = append(inputFiles, filePath)
			}
		}
	}
	for i := 0; i < 2; i++ {
		inputFiles = append(inputFiles, "token/empty.wav")
	}
	return inputFiles, nil
}

func replaceYSDD(s string) string {
	for chinese, phrase := range ysddDict {
		if strings.Contains(s, chinese) {
			s = strings.ReplaceAll(s, chinese, phrase)
		}
	}
	return s
}

func isYSDD(s string) bool {
	for _, value := range ysddDict {
		if value == s {
			return true
		}
	}
	return false
}

func getYSDD(input string) string {
	re := regexp.MustCompile(`\[(.*?)\]`)
	matches := re.FindStringSubmatch(input)
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
	re := regexp.MustCompile(`(\p{Han})|([a-zA-Z0-9]+)|(\[.*?\])|(\p{P})|(\W)`)
	return re.FindAllString(s, -1)
}

func isChinese(s string) bool {
	re := regexp.MustCompile(`^\p{Han}+$`)
	return re.MatchString(s)
}
