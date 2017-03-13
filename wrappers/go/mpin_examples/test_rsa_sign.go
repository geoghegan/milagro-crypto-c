/**
 * @file rsa_test.go
 * @author Alessandro Budroni
 * @brief RSA Wrappers tests
 *
 * LICENSE
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

 package main

import (
	"encoding/hex"
	"fmt"

	"github.com/miracl/amcl-go-wrapper"
)

func main() {

	var rtn int;

	// Seed value for Random Number Generator (RNG)
	seedHex := "9e8b4178790cd57a5761c4a6f164ba72"
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		fmt.Println("Error decoding seed value")
		return
	}

	rng := amcl.CreateCSPRNG(seed)

	// Generating public/private key pair
	fmt.Printf("Generating public/private key pair\n")
	RSA_PrivKey, RSA_PubKey := amcl.RSAKeyPair(&rng, 65537, nil, nil)

	// Message to encrypt
	MESSAGEstr := "Hello World\n"
	MESSAGE := []byte(MESSAGEstr)

	fmt.Printf("MESSAGE: ")
	fmt.Printf("%s\n\n", MESSAGE[:])

	// Signing message
	rtn, C := amcl.PKCS15(amcl.HASH_TYPE_RSA, MESSAGE)
	if rtn != 1 {
		fmt.Println("PKCS15 Error:", rtn)
		return
	}

	fmt.Printf("Padded MESSAGE: ")
	fmt.Printf("%x\n\n", C[:])

	// create signature in S
	S := amcl.RSA_DECRYPT(&RSA_PrivKey, C[:])

	fmt.Printf("Signed MESSAGE: ")
	fmt.Printf("%x\n\n", S[:])

	Cgot := amcl.RSA_ENCRYPT(&RSA_PubKey, S[:])

	fmt.Printf("Verify signature MESSAGE: ")
	fmt.Printf("%x\n\n", Cgot[:])

	// destroy private key
	fmt.Printf("Destroy private key\n\n")
	amcl.RSA_PRIVATE_KEY_KILL(&RSA_PrivKey)

	// NOTE - use the following only for testing
	rtn = int(amcl.OctetComp(C,Cgot))
	if rtn != 1 {
		fmt.Println("Error - Message doesn't correspond to the expected one:", rtn)
		return
	}
}