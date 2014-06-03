package codify

import (
	"crypto/sha256"
	"encoding/hex"
)

type Salting_Struct struct {

	Username string
	Salt string
}

// Users username and password, outputs password + salt
func SHA(user string, pass string) string {

	var password []byte
	var username []byte
	
	// Converts username to bytes
	username = []byte("user:" + user + pass)
	
	// Dial up a mongoDB session
	session, err := mgo.Dial("127.0.0.1:27017/")
    	if err != nil {
	     fmt.Println(err)
	     return ""		// Forcing a failure
    	}
    	
    	// Opens the "passwords" databases, "salts" collection
    	c := session.DB("passwords").C("salts")
    	
    	// Result with store username + password
    	result := Salting_Struct{}
    	
    	// Converts username to sha2
    	u := sha256.New()                   // new sha256 object
	u.Write(username)                   // data is now converted to hex
	code := u.Sum(nil)                  // code is now the hex sum
	codeusr := hex.EncodeToString(code) // converts hex to string
	
    	// Search for the salted username, place in result the salt
    	c.Find(bson.M{"username": codeusr}).One(&result)

	// Converts users input password + generated salt to bytes
	password = []byte(pass + result.salt)

	// Converts password + salt to sha2
	p := sha256.New()                   // new sha256 object
	p.Write(password)                   // data is now converted to hex
	code = p.Sum(nil)                   // code is now the hex sum
	codepass := hex.EncodeToString(code) // converts hex to string
	
        // close session to free resources
    	session.Close()
			
	return codestr
}
