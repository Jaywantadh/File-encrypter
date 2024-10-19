package main

import (
	"bytes"
	"fmt"
	"os"
	"github.com/jaywantadh/file-encrypter/filecrypt"
	"golang.org/x/term"
)

func main(){
	if len(os.Args) < 2{
		PrintHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		PrintHelp()
	case "encrypt":
		HandleEncrypt()
	case "decrypt":
		HandleDecrypt()
	default:
		fmt.Println("Run 'encrypt' to encrypt a file, and run 'decrypt' to decrypt a file.")
		os.Exit(1)
	}
}

func PrintHelp(){
	fmt.Println("_______J INFORSYSTEMS_______ ")
	fmt.Println("###### FILE ENCRYPTER ######")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tgo run. encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypt a file given a password")
	fmt.Println("\t decrypt\tTries to decrypt a file using a pasword")
	fmt.Println("\t help\t\tDisplays help text")
}

func HandleEncrypt(){
	if len(os.Args) < 2{
		fmt.Println("missing the path to the file. for more info, run go run .help")
		os.Exit(0)
	}

	file := os.Args[2]

	if !ValidateFile(file){
		panic("File not found")
	}

	password := GetPassword()
	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\n file Successfully encrypted and protected.")
}

func HandleDecrypt(){
	if len(os.Args) < 2{
		fmt.Println("missing the path to the file. for more info, run go run .help")
		os.Exit(0)
	}

	file := os.Args[2]

	if !ValidateFile(file){
		panic("File not found")
	}

	fmt.Println("Enter the password: ")
	password, _ := term.ReadPassword(0)
	fmt.Println("\nDecrypting...")
	filecrypt.Decrypt(file, password)
	fmt.Println("\n file Successfully Decrypted.")
}

func GetPassword() []byte{
	fmt.Println("Enter the Password: ")
	password, _ := term.ReadPassword(0)
	fmt.Println("\nConfirm Password: ")
	password2, _ := term.ReadPassword(0)
	if !ValidatePassword(password, password2){
		fmt.Println("\nThe passwords do not match. Try again.")
		return GetPassword()
	}
	return password
}

func ValidatePassword(password1 []byte, password2 []byte) bool{
	if !bytes.Equal(password1, password2){
		return false
	}
	return true
}

func  ValidateFile(file string) bool{
	if _,err := os.Stat(file); os.IsNotExist(err){
		return false
	}
	return true
}
