package main

import (
	"fmt"
 	"github.com/gorilla/mux"
	"log"
	"net/http"
	"html/template"
	"strings"
	_ "strings"
)

const linkHeader = "static/header.html"
const linkFooter = "static/footer.html"

type header struct {
	Link int
}

func webServer()  {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/", pageMain).Methods("GET")
	rtr.HandleFunc("/contact/", pageContact).Methods("GET")
	rtr.HandleFunc("/climate/", pageClimate).Methods("GET")
	rtr.HandleFunc("/verification", verification).Methods("GET")//.Queries("login")
	rtr.HandleFunc("/verification/", verification).Methods("GET").
					Queries("login","{login}").
					Queries("pass","{pass}")
	rtr.HandleFunc("/config/", pageConfig).Methods("GET")
	rtr.HandleFunc("/json/", pageJson).Methods("GET")
	rtr.HandleFunc("/blacklight/", pageBlacklight).Methods("GET")
	rtr.HandleFunc("/add/", pageAdd).Methods("GET").
					Queries("hum", "{hum}").
					Queries("temp", "{temp}").
					Queries("light", "{light}")

	http.Handle("/",rtr)

	fmt.Println("Server listen")
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("ListenAndServe = ",err)
	}
}

func pageMain(w http.ResponseWriter, r *http.Request) {
	tmpl,err := template.ParseFiles("static/index.html", linkHeader, linkFooter)
	if err != nil {
		fmt.Fprintf(w,err.Error())
	}
	tmpl.ExecuteTemplate(w,"header",  header{1})
	tmpl.ExecuteTemplate(w,"index", nil)
}

func pageClimate(w http.ResponseWriter, r *http.Request) {
	tmpl,err := template.ParseFiles("static/climate.html", linkHeader, linkFooter)
	if err != nil {
		fmt.Fprintf(w,err.Error())
	}

	tmpl.ExecuteTemplate(w,"header", header{2})
	tmpl.ExecuteTemplate(w,"climate",  writeData(readSql()) )
}

func pageContact(w http.ResponseWriter, r *http.Request) {
	tmpl,err := template.ParseFiles("static/contact.html", linkHeader, linkFooter)
	if err != nil {
		fmt.Fprintf(w,err.Error())
	}
	tmpl.ExecuteTemplate(w,"header",  header{3})
	tmpl.ExecuteTemplate(w,"contact",nil)
}

func pageConfig(w http.ResponseWriter, r *http.Request) {
	tmpl,err := template.ParseFiles("static/config.html", linkHeader, linkFooter)
	if err != nil {
		fmt.Fprintf(w,err.Error())
	}
	tmpl.ExecuteTemplate(w,"header",  header{4})
	tmpl.ExecuteTemplate(w,"config",nil)
}

func pageJson(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, getJson(dataType) )
}

func pageBlacklight(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, blacklight() )
}

func verification(w http.ResponseWriter, r *http.Request) {
	tmpl,err := template.ParseFiles("static/verification.html", linkHeader, linkFooter)
	if err != nil {
		fmt.Fprintf(w,err.Error())
	}
	vars := mux.Vars(r)
	var login string = vars["login"]
	pass := vars["pass"]

	if strings.Compare(login,"admin")==0 && strings.Compare(pass,"123")==0	{
		http.Redirect(w,r,"/config/",http.StatusSeeOther)  //переадрисация
	} else {
		tmpl.ExecuteTemplate(w,"header",  header{4})
		tmpl.ExecuteTemplate(w,"verification","test")
	}
}

func checkUser()  {

}

func pageAdd(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
		hum := vars["hum"]
		temp := vars["temp"]
		light := vars["light"]

	// подготавливаем данные
	add(hum,temp,light)
	fmt.Printf("ADD new data: hum=%v temp=%v  light=%v \n",hum,temp,light)
	//переадрисация
	http.Redirect(w,r,"/",http.StatusSeeOther)

}
//?hum=44&temp=22.5&light=55
