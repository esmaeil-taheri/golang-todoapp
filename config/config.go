package config

import "golang.org/x/crypto/bcrypt"

// Config gathers every runtime setting in one place instead of hardcoding
// values across the codebase. The composition root (cmd/) calls Load once and
// hands each layer only the pieces it needs.
type Config struct {
	UserFilePath string
	BcryptCost int
}

func Load() Config {
	return Config{
		UserFilePath: "user.txt",
		BcryptCost: bcrypt.DefaultCost,
	}
}
