package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var use string
	var password_lenght int
	var save_password_ask bool = ask_save_password()
	fmt.Print("How long should your password be?: ")
	fmt.Scan(&password_lenght)
	fmt.Print("\nWhat is the password for? (Only use _ for spaces): ")
	fmt.Scan(&use)
	password := create_password(password_lenght)
	if save_password_ask == true {
		save_password(use, password)
	}
	fmt.Println("Your password is: " + password)
}

func create_password(lenght int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	password := make([]rune, lenght)
	for i := range password {
		password[i] = letters[rand.Intn(len(letters))]
	}
	return string(password)
}

func ask_save_password() bool {
	var awnser bool
	for true {
		var save_password string
		fmt.Print("Do you want to save your password y or n?: ")
		fmt.Scan(&save_password)
		if save_password == "y" {
			awnser = true
			break
		} else if save_password == "n" {
			awnser = false
			break
		} else {
			fmt.Println("Please only awnser with y or n ")
		}
	}
	return awnser
}

func check_if_file_exists(file string) bool {
	var exists bool
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		// file does not exist
		exists = false
	} else {
		// file exists
		exists = true
	}
	return exists
}

func save_password(use string, password string) {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	exists := check_if_file_exists(mydir + "/Passwords/" + use + ".txt")
	replace := true
	if exists == true {
		for true {
			var save_password string
			fmt.Print("Do you want to replace your password y or n?: ")
			fmt.Scan(&save_password)
			if save_password == "y" {
				replace = true
				break
			} else if save_password == "n" {
				replace = false
				break
			} else {
				fmt.Println("Please only awnser with y or n ")
			}
		}
	}
	if replace == true {
		f, err := os.Create(mydir + "/Passwords/" + use + ".txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		d2 := []byte(password)
		n2, err := f.Write(d2)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		fmt.Println(n2, "Save password successfully at "+mydir)
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("Stopped saving the password")
	}
}
