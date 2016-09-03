package main

import (
    "fmt" //Printing
    "net/http"//Handling Http Requests
    "net" // Handles Unix Socket from Python Script
    "os"//Remove Unix Socket File if it exists
    "log" //Logging Errors
    "time"
    "io/ioutil" //Read in files
    "github.com/gorilla/websocket" //Create websockets easy
    //This is imported when you import io  "io/ioutil" // Write to file Easy
)

// func handler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

//Global Varrible to read the picture from
var raspPic []byte



func main() {
	fmt.Println("Welcome to Go Server!")

    //Handles Unix Socket Stuff
    go unixSocketCreate()

    

    http.HandleFunc("/", handler)
    http.ListenAndServe(":9696", nil)
}	

//Creates a unix socket that receives data from a python script to update the 
//Picture being showned via Websocket. Basicly this get pictures from camrea to
//send to websocket.
func unixSocketCreate() {
    //Remove unix socket file if it exisits 
    err := os.Remove("socket")

    //There will be an error if it doesnt so just ignore
    if err != nil {
        //fmt.Println(err)
    }

    //Create a unix Socket
    unixLisener, err := net.Listen("tcp", ":1776")

    //Creates a infinite loop to Accept Connections and starts a go thread to 
    //handle it
    for {
        //Accept Connection
        conn, err := unixLisener.Accept()
        //Pro Error handling here!!!
        if err != nil {
            fmt.Println("Error in excepting Unix Socket")
        }
        //Start a go thread to handle connection
        go unixSocketHandle(conn)
    }
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func unixSocketHandle(conn net.Conn) {

    raspPic, _ = ioutil.ReadAll(conn)
    // if(err != nil){
    //     log.Println(err)
    //     return
    // }
    
    fmt.Println("Picture Received from Unix Socket")
    err := ioutil.WriteFile("newPic.jpg", raspPic, 0777)
    if(err != nil){
        log.Println(err)
        return
    }
    //Close the connection
    conn.Close()
}



func handler(w http.ResponseWriter, r *http.Request) {
    
	// bytes, err := ioutil.ReadFile("pic.png")
 //    if err != nil {
 //        log.Println(err)
 //        return
 //    }


    //A server application uses the Upgrade function from an Upgrader object
    //with a HTTP request handler to get a pointer to a Conn:
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }

    //WriteMessage and ReadMessage methods to send and receive messages as a
    //slice of bytes

    for {
	    //Loop will wait untill a message is received
	    messageType, p, err := conn.ReadMessage()
	    if err != nil {
	    	fmt.Println("ERROR!!!")
	    	
	        return
	    }
	    messageType = messageType
	    p = p

	    for{ 
	    	time.Sleep(100 * time.Millisecond)
            err = conn.WriteMessage(websocket.BinaryMessage, raspPic); 
            if(err != nil ){
	           fmt.Println("ERROR!!!")
	           log.Println(err)
	           return
           }
	    }
	    fmt.Println("Loopin 4 Dayz")

	}
	fmt.Println("OUT OF THE LOOP")
}




