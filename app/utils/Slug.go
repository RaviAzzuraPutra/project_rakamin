package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func GenerateSlug(nama string) string {
	slug := strings.ToLower(nama)

	reg := regexp.MustCompile("[^a-z0-9]+")
	slug = reg.ReplaceAllString(slug, "-")

	slug = strings.Trim(slug, "-")

	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	randomNumber := 1000 + r.Intn(9000)

	return fmt.Sprintf("%s-%04d", slug, randomNumber)
}
