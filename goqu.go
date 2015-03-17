package goqu

import(
  "fmt"
//  "os"
  "net/http"
  "html/template"
  "github.com/gorilla/mux"
  "github.com/gorilla/schema"
  "database/sql"
)

// var database_url := os.Getenv("DATABASE_URL")
var DatabaseUrl = "postgres://wouter:@127.0.0.1/asq?sslmode=disable"
var decoder = schema.NewDecoder()
var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/new.html"))

type Query struct {
  id int
  Name string `schema:name`
  Query string `schema:query`
  Active bool `schema:"-"`
}

func renderTemplate(w http.ResponseWriter, tmpl string, q *Query) {
  err := templates.ExecuteTemplate(w, tmpl + ".html", q)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, World!"))
}

func QueriesIndexHandler(w http.ResponseWriter, r *http.Request) {
  q := &Query{ Name: "", Query: "" }
  renderTemplate(w, "new", q)
}

func QueriesCreateHandler(w http.ResponseWriter, r *http.Request) {
  form_err := r.ParseForm()

  if form_err != nil {
    // Handle error
  }

  query := new(Query)
  decode_err := decoder.Decode(query, r.PostForm)

  if decode_err != nil {
    // Handle error
  }

  db, db_err := sql.Open("postgres", DatabaseUrl)
  if db_err != nil {
    // Handle error
  }

  var id int
  err := db.QueryRow(`INSERT INTO "queries" (name, query, active)
                    VALUES($1, $2, $3) RETURNING id`, query.Name, query.Query, true).Scan(&id)

  if err != nil {
    w.Write([]byte(fmt.Sprintf("Error: %s ", err)))
  } else {
    http.Redirect(w, r, fmt.Sprintf("/queries/%d", id), http.StatusFound)
  }
}

func QueriesHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  query_id := vars["id"]

  w.Write([]byte(fmt.Sprintf("Query id is %s ", query_id)))

  db, db_err := sql.Open("postgres", DatabaseUrl)
  if db_err != nil {
    // Handle error
  }

  var query string
  err := db.QueryRow(`SELECT query FROM "queries" where "queries"."id"=$1 LIMIT 1`, query_id).Scan(&query)

  if err != nil {
    w.Write([]byte(fmt.Sprintf("Error: %s ", err)))
  }

  w.Write([]byte(fmt.Sprintf("Result %s ", query)))
}

func init() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  r.HandleFunc("/queries", QueriesIndexHandler).Methods("GET")
  r.HandleFunc("/queries", QueriesCreateHandler).Methods("POST")
  r.HandleFunc("/queries/{id}", QueriesHandler)
  http.Handle("/", r)
}
