# PBKDF2 File Protection: Encrypt & Decrypt

## Introduction
File Encryption is a simple command-line tool for encrypting and decrypting files using password-based encryption. This tool employs the Advanced Encryption Standard (AES) in Galois/Counter Mode (GCM) for secure encryption and decryption.

## Usage
To use the tool, run the main.go file with appropriate commands. The supported commands are:

- `encrypt`: Encrypts a file given a password.
- `decrypt`: Tries to decrypt a file using a password.
- `help`: Displays help text.

  ### Examples
  ```bash
  # Encrypt a file
  go run . encrypt /path/to/your/file
  
  # Decrypt a file
  go run . decrypt /path/to/your/file
  
  # Display help text
  go run . help
  ```

## File Encryption Algorithm

The encryption algorithm used is AES in GCM mode with a key derived from the provided password using the Password-Based Key Derivation Function 2 (PBKDF2) with SHA-1 as the hash function. The tool generates a random nonce for each encryption operation.

## Password Handling

The tool prompts the user to enter a password for encryption or decryption.
For added security, the tool uses the `golang.org/x/crypto/ssh/terminal` package to hide the password input.
The tool confirms that the entered passwords match before proceeding.

## Error Handling

Errors are handled gracefully throughout the code using the `log.Fatalf` function. If an error occurs, the tool logs an informative error message and exits.

## Dependencies

The tool uses the following external packages:

- `golang.org/x/crypto/aes`

- `golang.org/x/crypto/cipher`

- `golang.org/x/crypto/pbkdf2`

- `crypto/rand`

- `crypto/sha1`

- `encoding/hex`

- `io`

- `log`

- `os`

- `bytes`

## Disclaimer

This tool is designed for educational and personal use. Use it responsibly and ensure that you have the legal right to encrypt and decrypt files. The tool may have limitations and should not be used for sensitive or critical applications without thorough testing and validation.

*Please refer to the source code for detailed implementation and error handling.*
