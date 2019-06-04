package models

import (
	"github.com/jinzhu/gorm"
)

// Post model
type Post struct {
	gorm.Model
	Post string `json:"post"`
	URI  string `gorm:"not null;unique" json:"URI"`
}

// IncrementURI to generate the URI
func IncrementURI(last string) string {
	return string(nextChar([]byte(last))[:])
}

func nextChar(ch []byte) []byte {
	// Last Character
	// aaz -> aaA
	if ch[2] == 'z' {
		ch[2] = ch[2] - 57

		// Second character
		// aaZ -> aba
	} else if ch[2] == 'Z' && ch[1] != 'z' && ch[1] != 'Z' {
		ch[1]++
		ch[2] = 'a'
		// azZ -> aAa
	} else if ch[1] == 'z' && ch[2] == 'Z' {
		ch[1] = ch[1] - 57
		ch[2] = 'a'

		// Third Character
		// aZZ -> baa
	} else if ch[1] == 'Z' && ch[2] == 'Z' && ch[0] != 'z' {
		ch[0]++
		ch[1] = 'a'
		ch[2] = 'a'

		// zZZ -> Aaa
	} else if ch[0] == 'z' {
		ch[0] = ch[0] - 57
		ch[1] = 'a'
		ch[2] = 'a'

	} else {
		ch[2] += 1
	}
	return ch
}
