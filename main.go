package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const URL = "https://raw.githubusercontent.com/mymiscreant/birthday/refs/heads/main/users.json"

var current = time.Now()

type User []struct {
	Name        string `json:"Name"`
	DateOfBirth string `json:"DateOfBirth"`
}

func ReadJsonFromURL(string) User {
	rs, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
	}
	defer rs.Body.Close()
	var myUser User
	myjson, err := io.ReadAll(rs.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(myjson, &myUser)
	if err != nil {
		fmt.Printf("Error Json\n", err)
	}

	return myUser
}

func BirthInMonth(myUser User) {
	fmt.Println("В этом месяце день рождения празднуют:")
	for _, v := range myUser {
		userDate, _ := time.Parse("02.01.2006", v.DateOfBirth)
		if current.Month() == userDate.Month() {
			fmt.Printf("%v, %v\n", v.Name, v.DateOfBirth)

		}

	}

}

func NextDateOfBirth(myUser User) {
	fmt.Println("В ближайшее время ДР празднуют")
	var NextDateOfBirthUsers User
	for _, v := range myUser {
		userDate, _ := time.Parse("02.01.2006", v.DateOfBirth)
		userDate = time.Date(current.Year(), userDate.Month(), userDate.Day(), 0, 0, 0, 0, time.UTC)
		duration := userDate.Sub(current)
		days := int(duration.Hours() / 24)
		minDays := 365
		if days > 0 {

			if days < minDays {
				minDays = days

			}
			fmt.Println(minDays)
		}
	}

	fmt.Println(NextDateOfBirthUsers)
}

func DateOfBirthToday(myUser User) {
	for _, v := range myUser {
		userDate, _ := time.Parse("02.01.2006", v.DateOfBirth)
		if current.Day() == userDate.Day() && current.Month() == userDate.Month() {
			fmt.Printf("Сегодня ДР у %v (%v)\n", v.Name, v.DateOfBirth)

		}
	}

}

func main() {
	result := ReadJsonFromURL(URL)
	// DateOfBirthToday(result)
	// BirthInMonth(result)
	NextDateOfBirth(result)

}
