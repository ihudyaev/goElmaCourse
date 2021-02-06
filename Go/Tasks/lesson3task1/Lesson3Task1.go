//HTTP сервер - возвращает в ответе данные о запросе и тело запроса
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handlerEcho(w http.ResponseWriter, r *http.Request) {
	//Выводим параметры запрса при обращении на любой url
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	//Для POST запроса выводм тело запроса
	if r.Method == "POST" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		//Выводим тело также в консоль
		fmt.Printf("%s\n", reqBody)
		w.Write([]byte("Received a POST request\nBody:\n"))
		w.Write([]byte(reqBody))
		w.Write([]byte("\n"))
	}

}

func main() {
	//Добавляем обработчик - слушаем на порту 8000
	http.HandleFunc("/", handlerEcho)
	http.ListenAndServe(":8000", nil)
}
