package main

import (
	"fmt"
	 "net"
	 "os"
   "strings" 
)


type HttpRequest struct {
  Method string
  RequestTarget string
  HttpVersion string
  Host string 
  UserAgent string
  Accept string
  Body string
}

func parseRequest(request string) HttpRequest {
    lines := strings.Split(request, "\r\n")

    method := lines[0]
    requestTarget := strings.Fields(method)[1]
    host := ""
    accept := ""
    userAgent := ""

    for _, line := range lines[1:]{
      if strings.HasPrefix(line, "Host: ") {
        host = strings.TrimPrefix(line, "Host: ")
      } else if strings.HasPrefix(line, "Accept: ") {
        accept = strings.TrimPrefix(line, "Accept: ")
      } else if strings.HasPrefix(line, "User-Agent: ") {
        userAgent = strings.TrimPrefix(line, "User-Agent: ")
      }
    }

    return HttpRequest{
      Method: method,
      RequestTarget: requestTarget,
      Host: host,
      UserAgent: userAgent,
      Accept: accept,
    }

  }

func isValidPath(method string) bool{
    return method == "GET / HTTP/1.1" || strings.Contains(method,"GET /echo/") || strings.Contains(method,"GET /user-agent") || strings.Contains(method,"GET /files")
}

func getPathSizeAndContent(req HttpRequest) (string,int){
  content := strings.Split(req.RequestTarget,"/")
  lastElement := content[len(content)-1]
  var contentLength = len(lastElement)

  return lastElement, contentLength
}
  
func writeHttpResponse(content string) string{
  response := fmt.Sprintf("HTTP/1.1 200 OK\r\n"+
                       "Content-Type: text/plain\r\n"+
                       "Content-Length: %d\r\n\r\n%s", len(content), content)

  return response
}



func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	 listener, err := net.Listen("tcp", "0.0.0.0:4221") //Creates a tcp listener
	 if err != nil {
	 	fmt.Println("Failed to bind to port 4221")
	 	os.Exit(1)
	 }
   defer listener.Close()
  
   fmt.Println("Listening on port :4221...")
   for {

      conn, err := listener.Accept() // Waits for and return the next connection to the listener
  	  if err != nil {
    	 	fmt.Println("Error accepting connection: ", err.Error())
	     	os.Exit(1)
        continue
	    }
      go handleConnection(conn) // Manipula a conexão em uma goroutine
   }
 }

func handleConnection(conn net.Conn){
  defer conn.Close()
  fmt.Println("New connection from", conn.RemoteAddr())

  //Ler a request do cliente
  requestBytes := make([]byte, 1024)
  conn.Read(requestBytes)


  fmt.Println("--------------------")
  fmt.Println(string(requestBytes))
  fmt.Println("--------------------")

  req := parseRequest(string(requestBytes))  
  fmt.Printf("Parsed Request: \nMethod: %s \n", req.Method)
  //fmt.Println(strings.Fields(req.UserAgent)[1])
  if isValidPath(req.Method) == false{
    conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		fmt.Println("Path is not home.")
  }
  

  fmt.Println("--------------------")

  if strings.Contains(req.Method,"GET /echo") {
    content, _ := getPathSizeAndContent(req)
    response := writeHttpResponse(content)
    conn.Write([]byte(response))
  }

  if req.RequestTarget == "/user-agent"{
    userAgent := strings.Fields(req.UserAgent)
    fmt.Println("TAMANHO DO ARRAY USER AGENT",len(userAgent), req.UserAgent)
    //userAgentLast = userAgent[len(userAgent)-1]

    response := writeHttpResponse(req.UserAgent)

    fmt.Println("REPOSTA DA REQUISIÇÃO \n",response)                    

    fmt.Println("--------------------")
    conn.Write([]byte(response));
  }
  if strings.Contains(req.Method, "GET /files/"){
    fmt.Println("CAMINHO DO REQUEST",req.RequestTarget)
    
    fileDir := os.Args[2]
    fileName := strings.TrimPrefix(req.RequestTarget,"/files/")

    data, err := os.ReadFile(fileDir + fileName)
//    fmt.Println("erros na leitura do arquivo ", err)
		if err != nil {
      response := "HTTP/1.1 404 Not Found\r\n\r\n"
      conn.Write([]byte(response))
		} else {
      response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n%s", len(data), data)
      conn.Write([]byte(response));
		}
  } else{
    conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
  }
}


 
