package api

import (
	"encoding/json"
	"mod24/data"
	"net/http"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newExhibition data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&newExhibition)
		if err != nil {
			http.Error(w, "UnSupported Content", http.StatusNotAcceptable)
			return
		}
		data.Add(newExhibition)
		w.WriteHeader(http.StatusCreated) // we can send this status only as well.
		w.Write([]byte("New Exhibition has been added"))
	} else {
		http.Error(w, "UnSupported Method", http.StatusMethodNotAllowed)
	}

}

// In this case, I want to make sure that it's really a post.
// So for that their request has a method.
// So it's simple and we can check if it's post, get, delete or the one that we want.

// For the methods, we can use a string or we can check with in the constants in http.MethodPost, for example.
// If it's not a post, we can just right now throw an error over HTTP.We are passing the writer The message like Unsupported Method.
// The HTP status code that I want is probably one that isknown as MethodNotAllowed. So in this URL, we are not allowing that method.

// now, we will store the exhibition data came from the request in my exhibition collection we have in api package, we have Add() method there.
// which internally uses the append method, which creates the new list and not add in current one.

// Now we can go back to post and the next question is how to read the exhibitionthat is coming from the browser or from the client ?
// Well there are many ways to get data from the client but if the data is coming in the body as JSON andthat's the most common scenario then we can Then try to decode the body.
// So we've seen how to create a new encoder for the output.We also have a way to create a new decoder for the input. So we can receive the body as JSON and recreate an object from our side.

//I actually need to create the variable and send that variable to,in this case, the decoder and say hey, decode me that JSON into this variable.
// now we will create the vraible as newExhibition and then pass it as &newExhibition so that any chnage by the decoder will reflect back to it.
// also, we will handle the error which may be send by the decoder as well.

// now, we will call the method data.Add(newExhibition) and then pass the exhibition in it.
// finally, we will either send the status 'created' or the write over the http as well.

// make the new route in the server in main.go with using the HandleFunc as:
// server.HandleFunc("/api/make-exhibition", api.HandlePost)

// we will use the 'Postman' here to create the newExhibition and send the data to localhost:3335/api/make-exhibition
// send the data as below:
/* {
    "Title": "Life at Frontend Masters",
    "Description": "Uncover the world of ancient INDIA through the sculptures, tools, and jewelry found in ruins from over 2000 years ago that have been unearthed through modern science and technology.",
    "Image": "fem.jpg",
    "Color": "ORANGE",
    "CurrentlyOpened": true
} */

// and we will see the data added by checking the data added at the localhost:3335/api/get-exhibitions

// also, we can go the template route and then we will see the new added post as well. for that we will go to the localhost:3335/template.
// also the status code in postman on adding will be 201.
