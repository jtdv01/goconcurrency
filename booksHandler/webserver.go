package booksHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type bookResource struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Link string `json:"link"`
}

type requestPayload struct {
	Title string `json:"title"`
	Link string `json:"link"`
}

type response struct {
	StatusCode int
	Books []bookResource
}

type Action struct {
	Id string
	Type string
	Payload requestPayload
	RetChan chan<- response
}

func GetBooks() map[string]bookResource {
	books := map[string]bookResource{}
	for i := 1; i < 6; i++ {
		id := fmt.Sprintf("%d", i)
		books[id] = bookResource {
			Id: id,
			Title: fmt.Sprintf("Book-%s", id),
			Link: fmt.Sprintf("http://link-to-book%s.com", id),
		}
	}

	return books
}


// writeResponse uses the pattern similar to MakeHandler.
func writeResponse(w http.ResponseWriter, resp response) {
    var err error
    var serializedPayload []byte

    if len(resp.Books) == 1 {
        serializedPayload, err = json.Marshal(resp.Books[0])
    } else {
        serializedPayload, err = json.Marshal(resp.Books)
    }

    if err != nil {
        writeError(w, http.StatusInternalServerError)
        fmt.Println("Error while serializing payload: ", err)
    } else {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(resp.StatusCode)
        w.Write(serializedPayload)
    }
}

// writeError allows us to return error message in JSON format.
func writeError(w http.ResponseWriter, statusCode int) {
    jsonMsg := struct {
        Msg  string `json:"msg"`
        Code int    `json:"code"`
    }{
        Code: statusCode,
        Msg:  http.StatusText(statusCode),
    }

    if serializedPayload, err := json.Marshal(jsonMsg); err != nil {
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        fmt.Println("Error while serializing payload: ", err)
    } else {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(statusCode)
        w.Write(serializedPayload)
    }
}

// MakeHandler shows a common pattern used reduce duplicated code.
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string, string, chan<- Action),
    endpoint string, actionCh chan<- Action) http.HandlerFunc {

    return func(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path
        method := r.Method

        msg := fmt.Sprintf("Received request [%s] for path: [%s]", method, path)
        log.Println(msg)

        id := path[len(endpoint):]
        log.Println("ID is ", id)
        fn(w, r, id, method, actionCh)
    }
}



func writeReponse(w http.ResponseWriter, resp response) {
	var err error
	var serializedPayload []byte

	if len(resp.Books) == 1 {
		serializedPayload, err = json.Marshal(resp.Books[0])
	} else {
		serializedPayload, err = json.Marshal(resp.Books)
	}

	if err != nil {
		fmt.Println("Error while serializing payload: ", err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		w.Write(serializedPayload)
	}

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Received request [%s] for path: [%s]", r.Method, r.URL.Path)
	log.Println(msg)

	response := fmt.Sprintf("Hello world! at Path %s", r.URL.Path)
	fmt.Fprintf(w, response)
}

func Main() {
	const PORT int = 8080
	// http.HandleFunc("/", helloWorldHandler)
	books := GetBooks()
	log.Println(fmt.Sprintf("%+v", books))

	actionCh := make(chan Action)
	go StartBooksManager(books, actionCh)
	http.HandleFunc("/api/books/", MakeHandler(BookHandler, "/api/books/", actionCh))
	log.Println("Starting")
	log.Println(fmt.Sprintf("Starting webserver at port %d", PORT))
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}