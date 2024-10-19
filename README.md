# File Encrypter and Decrypter
A simple command-line tool written in Golang to encrypt and decrypt any type of file using the AES encryption algorithm. The tool allows you to secure your files with a password and decrypt them later using the same password.

## Features
-Encrypt any file with a password.

-Decrypt the encrypted file with the same password.

-Uses AES-GCM encryption for strong security.

-Password-based key derivation using PBKDF2 and SHA-1.

-Protects files by replacing them with their encrypted/decrypted versions.
## Prerequisites
Ensure you have the following installed on your machine:

-Go 1.16+

-Install Go from the official site: https://golang.org/dl/
## Installation
Clone the repository:

bash
```
git clone https://github.com/yourusername/file-encrypter.git
cd file-encrypter
```
## Install dependencies:

The tool uses the golang.org/x/crypto/pbkdf2 package for password-based key derivation. Make sure to install it by running:

bash
```
go get golang.org/x/crypto/pbkdf2
```
## Build the tool:

Once you have Go and the dependencies installed, you can build the program:

bash
```
go build -o filecrypt main.go filecrypt.go
This will generate an executable file named filecrypt.
```
## Usage
### Encrypt a file
To encrypt a file, run:

bash
```
./filecrypt encrypt /path/to/your/file
```
You will be prompted to enter a password twice for confirmation.
The file will be encrypted, and the original file will be replaced with the encrypted version.
### Decrypt a file
To decrypt a file, run:

bash
```
./filecrypt decrypt /path/to/your/file
```
You will be prompted to enter the password used during encryption.
The file will be decrypted and replaced with its original content.
## Help
For help on how to use the tool, run:

bash
```
./filecrypt help
```
Example
Here’s an example to encrypt and decrypt a file:

Encrypting a file:

bash
```
./filecrypt encrypt /path/to/file.txt
```
You'll be prompted to enter a password. After confirming, the file will be encrypted and replaced with its encrypted version.

Decrypting a file:

bash
```
./filecrypt decrypt /path/to/file.txt
```
Enter the same password used during encryption, and the file will be decrypted.

## Important Notes
Password Security: Make sure to remember the password you used to encrypt your files. The same password is required for decryption, and there is no way to recover files without it.

Overwrites: The tool overwrites the original file with its encrypted/decrypted version, so keep a backup if needed.

## How It Works
### Encryption

-The program reads the file into memory.

-It generates a random 12-byte nonce for AES-GCM.

-It uses PBKDF2 (Password-Based Key Derivation Function 2) with SHA-1 and the provided password to derive a 32-byte encryption key.

-The file is encrypted using AES-GCM, appending the nonce to the ciphertext.

-The encrypted file is written back to the same file path.
### Decryption

-The program reads the encrypted file, separating the nonce and the ciphertext.

-It uses PBKDF2 with the same password to derive the decryption key.

-The file is decrypted using AES-GCM and written back to the same file path.


## Dependencies
The tool uses the following Go packages:

crypto/aes: Provides AES block encryption.

crypto/cipher: Provides cipher modes like GCM.

crypto/rand: Secure random number generation.

crypto/sha1: Hash function for PBKDF2.

golang.org/x/crypto/pbkdf2: Implements PBKDF2 for password-based key derivation.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

Feel free to replace "https://github.com/yourusername/file-encrypter.git" with your actual GitHub repository link, and you can further adjust or expand the instructions based on your preferences.
