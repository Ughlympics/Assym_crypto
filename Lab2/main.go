package main

import (
	//"bufio"
	"fmt"
	"math/big"

	//"os"
	features "rsa/rsa"
	//"strings"
)

func main() {
	alice, _ := features.NewUser("Alice")
	bob, _ := features.NewUser("Bob")
	k1, S1, S := features.SendKey(alice, bob, big.NewInt(123456789))

	fmt.Println("=== Sent Values ===")
	fmt.Printf("k1 (encrypted key for Bob): 0x%X\n", k1)
	fmt.Printf("S1 (encrypted signature for Bob): 0x%X\n", S1)
	fmt.Printf("S (signature by Alice): 0x%X\n", S)
	fmt.Println()

	k, s := features.ReceiveKey(bob, k1, S1)
	fmt.Println("=== Received Values ===")
	fmt.Printf("k (decrypted key): 0x%X\n", k)
	fmt.Printf("s (decrypted signature): 0x%X\n", s)

}

// func main() {
// 	reader := bufio.NewReader(os.Stdin)

// 	alice, _ := features.NewUser("Alice")
// 	bob, _ := features.NewUser("Bob")

// 	fmt.Println("=== Generated Users ===")
// 	fmt.Printf("Alice's public key (N): 0x%X\n", alice.N)
// 	fmt.Printf("Bob's public key   (N): 0x%X\n", bob.N)
// 	fmt.Println()

// 	// ===== Выводим все открытые данные Алисы =====
// 	fmt.Println("=== Alice's Public Information ===")
// 	fmt.Printf("Name: %s\n", alice.Name)
// 	fmt.Printf("Public exponent (e): %d (0x%X)\n", alice.E, alice.E)
// 	fmt.Printf("Modulus (N): 0x%X\n", alice.N)
// 	fmt.Printf("Public key pair: (e = %d, N = 0x%X)\n", alice.E, alice.N)
// 	fmt.Println("==============================\n")

// 	fmt.Println("==============================\n")

// 	fmt.Println("Commands:")
// 	fmt.Println(" 1 - Encrypt message to Bob")
// 	fmt.Println(" 2 - Decrypt message by Bob")
// 	fmt.Println(" 3 - Sign message by Alice (method)")
// 	fmt.Println(" 4 - Verify Alice's signature (method)")
// 	fmt.Println(" 5 - Sign message (universal function)")
// 	fmt.Println(" 6 - Verify signature (universal function)")
// 	fmt.Println(" q - Exit")

// 	for {
// 		fmt.Print("\nEnter command: ")
// 		cmd, _ := reader.ReadString('\n')
// 		cmd = strings.TrimSpace(cmd)

// 		switch cmd {
// 		// ===== Encryption =====
// 		case "1":
// 			fmt.Print("Enter hex message: ")
// 			msgHex, _ := reader.ReadString('\n')
// 			msgHex = strings.TrimSpace(msgHex)
// 			message, ok := new(big.Int).SetString(msgHex, 16)
// 			if !ok {
// 				fmt.Println("Invalid hex message.")
// 				continue
// 			}
// 			cipher := bob.EncryptUser(message)
// 			fmt.Printf("Encrypted message: 0x%X\n", cipher)

// 		// ===== Decryption =====
// 		case "2":
// 			fmt.Print("Enter cipher (hex): ")
// 			cipherHex, _ := reader.ReadString('\n')
// 			cipherHex = strings.TrimSpace(cipherHex)
// 			cipher, ok := new(big.Int).SetString(cipherHex, 16)
// 			if !ok {
// 				fmt.Println("Invalid hex input.")
// 				continue
// 			}
// 			plain := bob.DecryptUser(cipher)
// 			fmt.Printf("Decrypted message: 0x%X\n", plain)

// 		// ===== Method: Sign =====
// 		case "3":
// 			fmt.Print("Enter message to sign (hex): ")
// 			msgHex, _ := reader.ReadString('\n')
// 			msgHex = strings.TrimSpace(msgHex)
// 			message, ok := new(big.Int).SetString(msgHex, 16)
// 			if !ok {
// 				fmt.Println("Invalid hex input.")
// 				continue
// 			}
// 			signature := alice.UserDigitalSign(message)
// 			fmt.Printf("[Method] Digital signature: 0x%X\n", signature)

// 		// ===== Method: Verify =====
// 		case "4":
// 			fmt.Print("Enter message (hex): ")
// 			msgHex, _ := reader.ReadString('\n')
// 			msgHex = strings.TrimSpace(msgHex)
// 			message, ok := new(big.Int).SetString(msgHex, 16)
// 			if !ok {
// 				fmt.Println("Invalid hex input.")
// 				continue
// 			}

// 			fmt.Print("Enter signature (hex): ")
// 			sigHex, _ := reader.ReadString('\n')
// 			sigHex = strings.TrimSpace(sigHex)
// 			signature, ok := new(big.Int).SetString(sigHex, 16)
// 			if !ok {
// 				fmt.Println("Invalid hex signature.")
// 				continue
// 			}

// 			isValid := alice.UserVerifySign(message, signature)
// 			if isValid {
// 				fmt.Println("[Method] Signature is VALID ✅")
// 			} else {
// 				fmt.Println("[Method] Signature is INVALID ❌")
// 			}

// 		// ===== Function: Sign =====
// 		case "5":
// 			fmt.Println("=== Digital Sign (universal) ===")
// 			fmt.Print("Enter message (hex): ")
// 			msgHex, _ := reader.ReadString('\n')
// 			msgHex = strings.TrimSpace(msgHex)
// 			message, ok := new(big.Int).SetString(msgHex, 16)
// 			if !ok {
// 				fmt.Println("Invalid message input.")
// 				continue
// 			}

// 			fmt.Print("Enter private exponent d (hex): ")
// 			dHex, _ := reader.ReadString('\n')
// 			dHex = strings.TrimSpace(dHex)
// 			d, ok := new(big.Int).SetString(dHex, 16)
// 			if !ok {
// 				fmt.Println("Invalid private key (d).")
// 				continue
// 			}

// 			fmt.Print("Enter modulus n (hex): ")
// 			nHex, _ := reader.ReadString('\n')
// 			nHex = strings.TrimSpace(nHex)
// 			n, ok := new(big.Int).SetString(nHex, 16)
// 			if !ok {
// 				fmt.Println("Invalid modulus (n).")
// 				continue
// 			}

// 			signature := features.DigitalSign(message, d, n)
// 			fmt.Printf("Signature: 0x%X\n", signature)

// 		case "6":
// 			fmt.Println("=== Verify Signature (universal) ===")
// 			fmt.Print("Enter message (hex): ")
// 			msgHex, _ := reader.ReadString('\n')
// 			msgHex = strings.TrimSpace(msgHex)
// 			message, ok := new(big.Int).SetString(msgHex, 16)
// 			if !ok {
// 				fmt.Println("Invalid message.")
// 				continue
// 			}

// 			fmt.Print("Enter signature (hex): ")
// 			sigHex, _ := reader.ReadString('\n')
// 			sigHex = strings.TrimSpace(sigHex)
// 			signature, ok := new(big.Int).SetString(sigHex, 16)
// 			if !ok {
// 				fmt.Println("Invalid signature.")
// 				continue
// 			}

// 			fmt.Print("Enter modulus n (hex): ")
// 			nHex, _ := reader.ReadString('\n')
// 			nHex = strings.TrimSpace(nHex)
// 			n, ok := new(big.Int).SetString(nHex, 16)
// 			if !ok {
// 				fmt.Println("Invalid modulus (n).")
// 				continue
// 			}

// 			isValid := features.VerifySign(message, signature, n)
// 			if isValid {
// 				fmt.Println("Signature is VALID ✅")
// 			} else {
// 				fmt.Println("Signature is INVALID ❌")
// 			}

// 		case "q":
// 			fmt.Println("Exiting...")
// 			return

// 		default:
// 			fmt.Println("Unknown command.")
// 		}
// 	}
// }

// var aStr3 = "ABC"
// var message, _ = new(big.Int).SetString(aStr3, 16)
// var aStr2 = "9764510DBF1C5B9FB344BCD245D8ED7080943ADEDB24F3168C10ADD2955E14D7"
// var n, _ = new(big.Int).SetString(aStr2, 16)
// cipher := features.Encrypt(n, message)
// fmt.Printf("Cipher:   0x%X\n", cipher)
