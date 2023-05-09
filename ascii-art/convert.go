package asciiart

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

var (
	hashstandard   = "ac85e83127e49ec42487f272d9b9db8b"
	hashshadow     = "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	hashthinkertoy = "db448376863a4b9a6639546de113fa6f"
)

func Converter(text string, banner string) string {
	data, err := ReadBannerFile(banner)
	if err == true {
		return "Error: Some issues during reading a banner file"
	}
	arg, err2 := SomethingChecker(text)
	if err2 == true {
		return "Error : input text contains non-printable characters"
	}
	res := WriteFileAsciiArt(arg, data)
	return res
}

func ReadBannerFile(banner string) ([]string, bool) {
	var error bool
	text, err := os.ReadFile("ascii-art/" + banner + ".txt")
	if err != nil {
		error = true
	}
	text_str := string(text)
	switch banner {
	case "standard":
		if GetMDHash(text_str) != hashstandard {
			error = true
		}
	case "shadow":
		if GetMDHash(text_str) != hashshadow {
			error = true
		}
	case "thinkertoy":
		if GetMDHash(text_str) != hashthinkertoy {
			error = true
		}
	default:
		error = true
	}
	arr := strings.Split(text_str, "\n")
	return arr, error
}

func GetMDHash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SomethingChecker(text string) ([]string, bool) {
	err := false
	text1 := strings.ReplaceAll(text, "\r\n", "\n")
	text2 := strings.Split(text1, "\n")
	for i, j := range text2 {
		if len(text2) == len(text1)+1 && j == "" {
			array := text2[i+1:]
			text2 = text2[:i]
			text2 = append(text2, array...)
		}
	}
	if AsciiChecker(text) == false {
		fmt.Println("Error : Argument is out of the range by ascii table: argument should only contain numbers, letters, spaces, special characters and /n")
		err = true
	}
	return text2, err
}

func AsciiChecker(text string) bool {
	arr := []rune(text)
	for i := 0; i < len(text); i++ {
		if (arr[i] < ' ' || arr[i] > '~') && arr[i] != '\n' && arr[i] != '\r' {
			return false
		}
	}
	return true
}

func WriteFileAsciiArt(arg []string, data []string) string {
	text := ""
	for _, k := range arg {
		if k == "" {
			text = text + "\n"
			continue
		}
		for i := 0; i <= 7; i++ {
			for _, l := range k {
				text = text + string(data[1+(9*(int(l)-32))+i])
			}
			text = text + "\n"
		}
	}
	return text
}
