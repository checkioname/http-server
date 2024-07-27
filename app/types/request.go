package request

type struct Request{
  Method string
  Host string 
  UserAgent string
  Accept string


  func parseRequest(request string) Request {
    lines := strings.Split(request, "\r\n")

    method := lines[0]
    host := lines[1]
    userAgent := lines[2]
    accept := lines[3]

    return Request{
      Method: method,
      Host: host,
      UserAgent: userAgent,
      Accept: accept
    }
  }
}
  



