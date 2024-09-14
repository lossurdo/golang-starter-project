package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateUsername(fullname string) string {
	words := strings.Fields(fullname)
	return strings.ToLower(words[0] + "-" + words[len(words)-1])
}

func removeDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func GenerateSimplePassword() string {
	password_list := []string{"funchal", "goiania", "regina", "tanger", "caracas", "cordoba", "cardiff", "napoles",
		"catania", "harbin", "pequim", "kolkata", "ottawa", "munique", "bangcoc", "nantes", "houston", "jacarta",
		"fuzhou", "roterda", "recife", "toquio", "shiraz", "zagreb", "toronto", "jaipur", "wroclaw", "exeter",
		"xangai", "cologne", "bangkok", "kigali", "lucknow", "calcuta", "genebra", "genova", "raleigh", "viamao",
		"seattle", "zurich", "glasgow", "dublin", "chicago", "luanda", "detroit", "angeles", "calgary", "damasco",
		"halifax", "lahore", "vitoria", "sapporo", "tallinn", "vilnius", "botucatu", "niteroi", "fukuoka",
		"londres", "eugene", "dallas", "rangoon", "maputo", "teheran", "bilbau", "manila", "havana", "sydney",
		"madurai", "tijuana", "racoon", "toluca", "jeddah", "beirut", "lublin", "puebla", "zahedan", "harare",
		"colombo", "bogota", "bufalo", "nagoia", "moscou", "lisboa", "memphis", "uruguai", "jakarta", "karachi",
		"denver", "sevilha", "chennai", "nanjing", "belize", "bremen", "rosario", "tianjin", "manama", "dresden",
		"bergen", "merida", "mumbai", "atenas", "berlim", "erevan", "leipzig", "tehran", "manila", "oakland",
		"gdansk", "cuiaba", "cancun", "mexico", "glasgow", "dublin", "chicago", "luanda", "detroit", "angeles",
		"calgary", "damasco"}

	password_list = removeDuplicate(password_list)

	random_two_digits := rand.Intn(100)

	return password_list[rand.Intn(len(password_list))] + strconv.Itoa(random_two_digits)
}

func GeneratePassword(length int, strong bool) string {
	var chars string
	if strong {
		chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()"
	} else {
		chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	}

	rand.NewSource(time.Now().UnixNano())

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = chars[rand.Intn(len(chars))]
	}

	return string(password)
}
