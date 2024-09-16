package response

import "fmt"


func WriteHttpResponse(content string) string{
  response := fmt.Sprintf("HTTP/1.1 200 OK\r\n"+
                       "Content-Type: text/plain\r\n"+
                       "Content-Length: %d\r\n\r\n%s", len(content), content)

  return response
}
