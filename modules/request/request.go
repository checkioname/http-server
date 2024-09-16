package request

import "strings"

type HttpRequest struct {
  Method         string
  RequestTarget  string
  HttpVersion    string
  Host           string 
  UserAgent      string
  Accept         string
  Body           string
}

func ParseRequest(request string) HttpRequest {
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
