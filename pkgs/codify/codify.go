package codify

import (
	"crypto/sha256"
	"encoding/hex"
)

type Salting_Struct struct {

	Username string
	Salt string
}

// Runs a SHA 2 function on the string
func SHA(str string) string {

	bytes := []byte(str)

	// Converts string to sha2
    	h := sha256.New()                   	// new sha256 object
	h.Write(bytes)                  	// data is now converted to hex
	code := h.Sum(nil)      		// code is now the hex sum
	codestr := hex.EncodeToString(code) 	// converts hex to string
	
	return codestr
}

// Users username and password, outputs password + salt
func Password(user string, pass string) string {

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
	
    	// Search for the salted username, place in result the salt
    	c.Find(bson.M{"username": SHA(username)}).One(&result)
    	
        // close mongoDB session to free resources
    	session.Close()
	
	// Converts users input password + generated salt to bytes
	password = []byte(pass + result.salt)
			
	return SHA(password)
}
