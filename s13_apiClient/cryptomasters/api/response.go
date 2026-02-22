package api

// here we will create the struct for the api response and we will use this to get the Rate struct value from it.
// we don't need to take all the properties in the response we will take only we need here
// https://mholt.github.io/json-to-go/

type APIResponse struct {
	Result   string `json:"result"`
	Provider string `json:"provider"`
	BaseCode string `json:"base_code"`
	Rates    map[string]float64  `json:"rates"`  // slice of map of string to float64 value. 
}

// Properties in structures can contain metadata that is just a string ie. `json:"result"` is the meta data about the result. It's kind of a comment that you can add on properties.
// It's just a string.Will go do something with that?No, it's like a comment.However, back caches, libraries, frameworks,they can read using reflection APIs, those comments, labels.The json library can read metadata when it's marshaling or unmarshaling objects.
// Reflection APIs lets you inspect types on the fly.So what it means is that when you use json column,it means I'm going to give you an instruction for the JSON library.Then I could add more things here for other libraries, okay?It can be an object to database mapping library or things like that,that can create the SQL code for you, things like that.
// The thing is that JSON needs the the instruction that you're going to send tothe library in double quotes.So that's why I'm not using double quotes there because then they have a quoteproblem.So that's why typically we use backtick.That is another way to the final string in Go.So then we can use double quotes inside.So roughly, what you need is here,you're going to specify how is that property named in the JSON.
//So then you can map, instead of here,if you say, for example, well, first you can use capital letters now.And if you have one property with a hyphen, with a space, or something likethat, then you have a different thing on your side, and now it would work.So this is the mapping definition between the JSON and your Go structure. 
