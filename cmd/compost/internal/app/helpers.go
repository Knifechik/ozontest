package app

import "unicode/utf8"

const maxCommentLength = 2000

func commentLengthCheck(comment string) error {
	length := utf8.RuneCountInString(comment)
	if length > maxCommentLength {
		return ErrOverCommentLength
	}
	return nil
}
