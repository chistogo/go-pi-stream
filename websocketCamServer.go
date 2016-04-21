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
    http.ListenAndServe(":8080", nil)
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
    unixLisener, err := net.Listen("unix", "./socket")

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







//
// cespare | chistogo: play.golang.org                                                                                                                                                │ _cn
// 21:24:44          <-- | turtil (~turtil@unaffiliated/kylescottmcgill) has quit (Ping timeout: 252 seconds)                                                                                       │ _jesse_
// 21:25:54          --> | clouddig (~clouddig@cce02cs4040-fa12-z.ams.hpecore.net) has joined #go-nuts                                                                                              │ _rsc`
// 21:25:57      cespare | chistogo: if you want to read everything until the connection is closed, you can just use ioutil.ReadAll                                                                 │ a00001
// 21:26:43     chistogo | I did't know that. I will use that                                                                                                                                       │ aarwine
// 21:26:52      cespare | chistogo: You're overwriting your buffer from the beginning in every loop iteration                                                                                      │ abra0
// 21:27:08          --> | velovix (~velovix@149.169.218.242) has joined #go-nuts                                                                                                                   │ acmehendel
// 21:27:10          <-- | nathanleclaire (~nathanlec@c-73-189-234-64.hsd1.ca.comcast.net) has quit (Quit: nathanleclaire)                                                                          │ acrocity
// 21:27:10 +skelterjohn | also, yeah that                                                                                                                                                          │ adam^
// 21:27:25     chistogo | cespare Is that bad?                                                                                                                                                     │ adamcm
// 21:27:31      cespare | chistogo: it's wrong                                                                                                                                                     │ adeschamps
// 21:27:32 +skelterjohn | if you were to write .ReadFile, you'd need to move the buffer                                                                                                            │ aduermael_
// 21:27:36          <-- | charles (~quassel@pdpc/supporter/active/charles) has quit (Remote host closed the connection)                                                                            │ AG_Clinton
// 21:27:38 +skelterjohn | chistogo: it simply cannot work as written                                                                                                                               │ agatoxd
// 21:27:51      cespare | chistogo: with the small file, it so happens you slurp up the whole thing in one Read                                                                                    │ agundy
// 21:27:56          <-- | frostyfrog (~frostyfro@unaffiliated/frostyfrog) has quit (Read error: Connection reset by peer)                                                                          │ ahuman
// 21:27:58          --> | tonist (~tonistiig@c-24-23-191-227.hsd1.ca.comcast.net) has joined #go-nuts                                                                                              │ aidan-
// 21:28:18      cespare | chistogo: but with the large one, it takes multiple turns around the loop, and the way you've written it is broken because each read writes the same data at the         │ aissen
//                       | beginning of the buffer                                                                                                                                          

//Message types : websocket.BinaryMessage or websocket.TextMessage (UTF-8)