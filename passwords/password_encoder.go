package passwords

type PasswordEncoder interface {
	// 加密
	Encode(raw string) ([]byte, error)

	// 是否匹配
	Match(raw, encode string) (bool, error)
}
