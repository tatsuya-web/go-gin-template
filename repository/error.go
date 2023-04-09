package repository

import "errors"

var (
	ErrNotFound          = errors.New("見つかりません。")
	ErrNotExistEmail     = errors.New("メールアドレスが存在しません。")
	ErrAlreadyEntry      = errors.New("登録済みのメールアドレスは登録できません。")
	ErrNotFoundSession   = errors.New("確認コードまたは、セッションキーが無効です。")
	ErrNotMatchLogInfo   = errors.New("メールアドレスまたは、パスワードが異なります。")
	ErrNotUser           = errors.New("ユーザが存在しません。")
	ErrDifferentPassword = errors.New("パスワードが異なります。")
	ErrUnauthorizedUser  = errors.New("権限の無いユーザーです。")
)

const (
	// ErrCodeMySQLDuplicateEntry はMySQL系のDUPLICATEエラーコード
	// https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html
	// Error number: 1062; Symbol: ER_DUP_ENTRY; SQLSTATE: 23000
	ErrCodeMySQLDuplicateEntry = 1062
)
