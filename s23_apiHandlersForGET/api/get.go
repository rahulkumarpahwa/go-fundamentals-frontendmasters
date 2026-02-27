package api

import (
	"encoding/json"
	"fmt"
	"mod23/data"
	"net/http"
	"strconv"
)

// here we will take the query and if we recive the valid query, which contains the index then we will return that exhibition of the slice and else we will return all the exhibitions from data/exhibitions.go

// same put the Responsewriter and then http.Request, that's the signature of every handler.

// here we need is to just the send the output JSON as we have known the marshal and unmarshal but here we will use something else.
// In the json package, we have NewEncoder(), which will take the argument writer w, we have.

//So in this case, instead of receiving a string and using our Marshall and Marshal to work with the strings, we can just send directly the stream of JSON to the output.

//This is even faster. I don't even need to create a temporary string in memory. So basically in this case,while the encoder is taking out my object and creating the string forthe JSON, we are sending the device to the network directly.

// We go to a string.And we asked that encoder to encode one particular object, which object?The one that we have in data, GetAll, very simple.

// Do we need something else?No.The only thing that if you want to be a good citizen for API's, the only thingthat you need to add is a header to specify that you're sending JSON.Because by default, HTTP understands that you're sending text.

//Some APIs will never care about that header, butmaybe it's a good idea to add a header.And to add a header, we also talk to the writer.And we say, header, it's a function that returns an object value.So with some functions for example set we can set a header.

// if you are experienced with HTTP headers butthe header is content type and the value that we need to send isthe type that for JSON is typically application JSON.

// That's a handler, and to test the handler,I need to register the handler in my server.
// So we need to go to the main and register this handler.

// func Get(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(data.GetAll())
// }

// now we will take the query from the URL say /api/exhibitions?id=1 then we will return the first one struct of the exhibition.

// Well if I have 'id' then what I need first is to convert that into an integer because this is something the user can change. So I cannot just treat it as a number.I need to convert that into a number, and see if it's a real number.

// remember I mentioned before that there is a package 'strconv' forstring conversions.
// Here we have a lot of functions actually to convert from and to strings.And there is one with the name Atoi.
// Atoi. It's actually similar to Parseint that is prepared.
// So it's parsing this as an integer usingautomatically the decimal base, so base 10.
// So in this case I can parse that ID and every converter tool will return me the,let's say the finalid and also possible error because maybe andyou know what maybe this is not working.

// we will aslo check that the ID should not be greater than the length of the slice.
// Make sense?Because if that happens, I'm trying to get an exhibition that doesn't exist.

// So in those cases I'm going to return an error.And we already return a 500, setting the header,but I will show you another way to return an error, a faster way.
// You can talk to http and say http.Error.
// It's a function receiving the writer,A message like Invalid Exhibition, andthen the status that you wanna send back, status bad request.This is just a utility function to reduce the steps.

// first start the server and go to browser and type http://localhost:3334/api/exhibitions?id=1 we will get the first structure data.

// What happens if I parse, Hello or maybe Panic,as an id I'm getting invalid exhibition.

// Note : By the way here I said get but am I verifying thatthat request was actually made with a Get or not?Actually, no, I can actually make a POST to that URL, or a DELETE to that URL.I'm also returning a GET.In this case, it's harmless, because nothing happens.But with a POST or a DELETE, you need to be sure that the verb that you'rereceiving is the one that you're expecting.

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// /api/exhibitons?id=1
	// id := r.URL.Query()["id"];  // returns the slice of values eg. ?id=34,3,4, etc and we have to take out the first one.
	id := r.URL.Query().Get("id") // directtly returns the first one of the slice
	if id != "" {                 // when needed only when exhibition.
		final_id, err := strconv.Atoi(id)
		if err == nil && final_id < len(data.GetAll()) {
			newData := data.GetAll()[final_id-1]
			json.NewEncoder(w).Encode(newData) // here we will send that particular id data we want to get.
		} else {
			fmt.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
	} else { // when we want all the exhibitions.
		err := json.NewEncoder(w).Encode(data.GetAll())
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

}
