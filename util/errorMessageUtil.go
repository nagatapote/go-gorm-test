package util

import "errors"

var (
	// ErrorEmailOrPassIsWrong error
	ErrorEmailOrPassIsWrong = errors.New("メールアドレスかパスワードに誤りがあります")
	// ErrorIDOrPassIsWrong error
	ErrorIDOrPassIsWrong = errors.New("IDかパスワードに誤りがあります")
	// ErrorNoParameter error
	ErrorNoParameter = errors.New("値が不足しています")
	// ErrorServerError error
	ErrorServerError = errors.New("サーバー内エラーです")
	// ErrorDupricationError error
	ErrorDupricationError = errors.New("値が重複しています")
	// ErrorNotAllowed error
	ErrorNotAllowed = errors.New("権限が不足しています")
	// ErrorEmailOrPassIsWrong error
	ErrorTotpIsWrong = errors.New("ワンタイムパスワードが間違っています")
)
