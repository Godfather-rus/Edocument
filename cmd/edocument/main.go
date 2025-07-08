package main

import (
	"log"
)

func main() {
	ap, err := app.NewApp()
	if err != nil {
		log.Fatal("failed to create app", "error", err)
		return
	}

	if err = ap.Run(); err != nil {
		log.Fatal("failed to run app", "error", err)
	}
}

//func InitServer() {
//
//	log.Println("Server running")
//
//	r := chi.NewRouter()
//
//	r.Post(
//		"/api/docs", UploadDoc)
//
//	r.Get("/api/docs", GetEdocsList)
//	r.Get("/api/docs/{id}", GetEdoc)
//
//	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
//		result := "hello"
//		resultJson, err := json.Marshal(result)
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte(err.Error()))
//
//			return
//		}
//		w.Write(resultJson)
//	})
//
//	srv := &http.Server{Addr: ":8080", Handler: r}
//
//	go func() {
//		err := srv.ListenAndServe()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}()
//
//	// Wait for an interrupt
//	c := make(chan os.Signal, 1)
//	signal.Notify(c, os.Interrupt)
//	<-c
//	log.Println("Shutting down server...")
//
//	// Attempt a graceful shutdown
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	if err := srv.Shutdown(ctx); err != nil {
//		log.Fatalf("Server forced to shutdown: %v", err)
//	}
//	log.Println("Server exiting")
//}
//
//func GetEdoc(w http.ResponseWriter, r *http.Request) {
//	edocDAO := ConnectDB()
//
//	id := chi.URLParam(r, "id")
//	res, err := edocDAO.FindOne(id)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		w.Write([]byte(err.Error()))
//
//		return
//	}
//	resultJson, err := json.Marshal(*res)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		w.Write([]byte(err.Error()))
//
//		return
//	}
//	w.Write(resultJson)
//}
//
//func GetEdocsList(w http.ResponseWriter, r *http.Request) {
//	edocDAO := ConnectDB()
//	//var data Data
//
//	res, err := edocDAO.FindAll(context.Background(), bson.D{})
//
//	//for i, v := range *res {
//	//
//	//}
//	//result := *res
//	resultJson, err := json.Marshal(*res)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		w.Write([]byte(err.Error()))
//
//		return
//	}
//	w.Write(resultJson)
//}
//
//func UploadDoc(w http.ResponseWriter, r *http.Request) {
//	edocDAO := ConnectDB()
//
//	err := r.ParseMultipartForm(32 << 20)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	var data Data
//
//	//Access form fields
//	//id := r.FormValue("id")
//	//fmt.Fprintf(w, "id: %s\n", id)
//
//	jsonData := r.FormValue("json")
//	//fmt.Fprintf(w, "json: %s\n", jsonfield)
//
//	//err = json.Unmarshal([]byte(jsonData), &data.Data.Json)
//	//if err != nil {
//	//	//log.Fatal(err)
//	//}
//
//	//ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
//	//defer cancel()
//	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
//	//if err != nil {
//	//	http.Error(w, err.Error(), http.StatusInternalServerError)
//	//	return
//	//}
//	//collection := client.Database("baz").Collection("qux")
//	//
//	//res, err := collection.InsertOne(context.Background(), bson.M{"hello": "world"})
//	//if err != nil {
//	//	http.Error(w, err.Error(), http.StatusInternalServerError)
//	//	return
//	//}
//	//id := res.InsertedID
//
//	var doc map[string]interface{}
//	err = json.Unmarshal([]byte(jsonData), &doc)
//
//	res, err := edocDAO.Insert(context.Background(), doc)
//	if err != nil {
//		log.Fatal("DB error: ", err)
//	}
//	id := res.InsertedID
//	fmt.Fprintf(w, "id: %s\n", id)
//
//	// Access uploaded files
//	file, header, err := r.FormFile("file")
//	if err != nil {
//		fmt.Fprintf(w, "No file uploaded or error: %v\n", err)
//		return
//	}
//	defer file.Close()
//
//	res, err = edocDAO.Insert(context.Background(), bson.M{"hello": "world"})
//	if err != nil {
//		log.Fatal("DB error: ", err)
//	}
//
//	fmt.Fprintf(w, "File name: %s\n", header.Filename)
//	// You can then read from 'file' (an io.Reader) to process the uploaded content
//
//	//err := json.NewDecoder(r.Body).Decode(&item)
//	//if err != nil {
//	//	w.WriteHeader(http.StatusBadRequest)
//	//	w.Write([]byte(err.Error()))
//	//	return
//	//}
//	//fmt.Println(item)
//
//	jsonItem, err := json.Marshal(data)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		w.Write([]byte(err.Error()))
//		return
//	}
//	w.Write(jsonItem)
//}
//
//func ConnectDB() *EdocDAO {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
//	if err != nil {
//		log.Fatal("failed to init db connection")
//	}
//
//	//defer func() {
//	//	if err = client.Disconnect(ctx); err != nil {
//	//		panic(err)
//	//	}
//	//}()
//
//	edocDAO, err := NewEdocDAO(ctx, client)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return edocDAO
//
//	//repo := repository.NewRepository(client)
//}
