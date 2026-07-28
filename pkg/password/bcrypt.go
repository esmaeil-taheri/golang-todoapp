// Package password isolates bcrypt behind two small functions so the rest of
// the code never imports the crypto library directly. It's a generic utility,
// hence its home under pkg/.

package password

import "golang.org/x/crypto/bcrypt"

// Hash returns the bcrypt hash of a plain-text password.
func Hash(plain string, cost int) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), cost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// Compare reports whether plain matches a previously hashed password.
func Compare(hashed, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}
