/*
General helper functions
*/
package helper

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)


func StringifyList(l []string, sep string) string {
	if len(l) == 0 { return "" }

	return strings.Join(l, sep)
}

func FindToken(token string, list []string, resultList *[]string) error {
	if len(list) == 0 {
		return errors.New("List is empty")
	}
	for _, item := range list {
		item = strings.Replace(item, token, fmt.Sprintf("\033[30;43m%s\033[0m", token), -1)
		if strings.Contains(item, token) {
			*resultList = append(*resultList, item)
		}
	}
	slices.Sort(*resultList)

	return nil
}
