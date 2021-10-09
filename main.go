package main
import (
    "encoding/json"
    "context"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

var collection1 *mongo.Collection
var collection2 *mongo.Collection
var ctx = context.TODO()

//MODELS
// User - Our struct for all user
type User struct {
	Id   string   `bson:"Id,omitempty"`
	Name    string   `bson:"Name,omitempty"`
	Email    string   `bson:"Email,omitempty"`
	Password    string   `bson:"Password,omitempty"`
}
// Post - Our struct for all post
type Post struct {
   Id   string   `bson:"Id,omitempty"`
   Caption    string   `bson:"Caption,omitempty"`
   ImageUrl    string   `bson:"ImageUrl,omitempty"`
   PostedTimestamp    string   `bson:"PostedTimestamp,omitempty"`
}



func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!, this task is done by Sriesh Agrawal(19BIT0407)")
    fmt.Println("Endpoint Hit: homePage")
}





//TASK ROUTE -1
func createNewUser(w http.ResponseWriter, r *http.Request) {  
    collection1 := db().Database("instagram_appointy").Collection("users")
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
     newUser := []interface{}{
        bson.D{{"Id" , user.Id},{"Name" , user.Name},{"Email" , user.Email},{"Password" , user.Password}},
    }
    res, insertErr := collection1.InsertMany(ctx, newUser)
    if insertErr != nil {
        log.Fatal(insertErr)
    }
    fmt.Println(res);
}   

//TASK ROUTE -2
func returnSingleUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
    collection1 := db().Database("instagram_appointy").Collection("users")
    cursor, err := collection1.Find(ctx, bson.M{"Id":key})
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var user bson.M
        if err = cursor.Decode(&user); err != nil {
            log.Fatal(err)
        }
        json.NewEncoder(w).Encode(user)
    }
}

//TASK ROUTE -3
func createNewPost(w http.ResponseWriter, r *http.Request) {  
    collection2 := db().Database("instagram_appointy").Collection("posts")
    var post Post
    err := json.NewDecoder(r.Body).Decode(&post)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
     newPost := []interface{}{
        bson.D{{"Id" , post.Id},{"Caption" , post.Caption},{"ImageUrl" , post.ImageUrl},{"PostedTimestamp" , post.PostedTimestamp}},
    }
    res, insertErr := collection2.InsertMany(ctx, newPost)
    if insertErr != nil {
        log.Fatal(insertErr)
    }
    fmt.Println(res);
}    

//TASK ROUTE -4
func returnSinglePost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
    collection2 := db().Database("instagram_appointy").Collection("posts")
    cursor, err := collection2.Find(ctx, bson.M{"Id":key})
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var post bson.M
        if err = cursor.Decode(&post); err != nil {
            log.Fatal(err)
        }
        json.NewEncoder(w).Encode(post)
    }
}

//TASK ROUTE -5

func returnPostOfUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
    collection2 := db().Database("instagram_appointy").Collection("posts")
    cursor, err := collection2.Find(ctx, bson.M{"Id":key})
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var post bson.M
        if err = cursor.Decode(&post); err != nil {
            log.Fatal(err)
        }
        json.NewEncoder(w).Encode(post)
    }
}


func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/users", createNewUser).Methods("POST")
    myRouter.HandleFunc("/posts", createNewPost).Methods("POST")
    myRouter.HandleFunc("/users/{id}", returnSingleUser)
    myRouter.HandleFunc("/posts/{id}", returnSinglePost)
    myRouter.HandleFunc("/posts/users/{id}", returnPostOfUser)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    handleRequests()
}



