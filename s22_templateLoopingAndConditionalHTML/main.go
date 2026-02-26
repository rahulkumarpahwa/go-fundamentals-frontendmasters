package main

import (
	"mod22/data"
	"net/http"
	"text/template"
)

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./template/template.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
		return
	}
	html.Execute(w, data.GetAll()) // instead of sending the first one exhibition we will send the whole slice of structure.
}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/template", handleTemplate)
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	http.ListenAndServe(":3333", server)

}

// we will send all the exhibition structs from the slice at once and in the template to handle them we will use the range and .
/**
  {{ range . }}
     <article>
       <h2>{{ .Title }}</h2>
       <p>
         {{ .Description }}
       </p>
       <img src="/gallery/{{ .Image }}" fetchpriority="high" decoding="sync" />
     </article>
      {{ end }}
*/

// Looping over the range :
// here range is used instead of FOR loop and . represents the each struct of the slice. So, internally we don't need to change anything.

// write add the two braces in the start and end of the article to show its start and end.

// make sure to add the fs (file server), this is used to serve the static files such as css and images and remove the script tag in the template.

// Adding the CSS with template variables :
// we can also change the css in the similar way. like lets add the new property in the struct Exhibition in the exhibitions.go and

// then we will try to apply that in the template as well using the same {{ }} and the dot, but this time we will pass this as a property value to the border-color in style attribute as :

//  <article style="border: 5px solid {{ .Color }}">

// In, Normal html file this may give the error but in the template file, this works fine.

// if-else conditions using the template:
// Code here :
//  <article style="border: 5px solid {{ .Color }}" class="{{if .CurrentlyOpened}} opened{{else}}closed {{end}}">
// we will add the css classes and put the attributes on the opened and closed based upon the need.
// now, when the CurrentlyOpened is true for the struct passed then it will have the class 'opened' in it and when false then closed class will be added in the tag.
// note : The if-else logic was actually working, but the extra spaces in your template were breaking the CSS class matching. make sure to have not the proper spacing.

// {
// 	Title:           "Life in Ancient Greek",
// 	Description:     "Uncover the world of ancient Greece through the sculptures, tools, and jewelry found in ruins from over 2000 years ago that have been unearthed through modern science and technology.",
// 	Image:           "ancient-greece.png",
// 	Color:           "red",
// 	CurrentlyOpened: true, // here is propety added
// },

// Also, Note: when we change only the CSS files, or the template (template is accessed again by the server everytime it needs it) just then we don't need to restart again as the GO Code remains unchanged.
// Note : If you don't see the change, you may have to clear your browser cache.

// Also, note that:
// So sometimes you wanna make that you don't want everythingto be like within the expressions but I don't want that space in the final output ie. around the classess like opened and closed around them in html tag when seen in the source code in the browser. Well, to solve that,we have a special syntax is a trimming syntax over the double curly brace. It's just a hyphen that you can add on each side of the declaration. Like this, in this case, it's saying that it should remove spaces after or before that declaration.

// Eg :
//  <article style="border: 5px solid red" class=" opened "> from browser page source.
// so solution is :

// <article style="border: 5px solid {{ .Color }}" class="{{- if .CurrentlyOpened -}}opened{{- else -}} closed {{- end -}}">
// make sure to have the space between the - and the text.

// Just that. So then I can keep a lot of spaces here.Even I can send enters as well, like new line. So if I refresh now,I don't have any space in the class in source code but I still have them in my template.
