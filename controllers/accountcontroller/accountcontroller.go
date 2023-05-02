package accountcontroller

func Index(responde http.ResponseWriter, request *http.Request) {

	tmp, _ := template.ParseFiles("views/accountcontroller/index.html")
	tmp.Execute(responde, nil)

}