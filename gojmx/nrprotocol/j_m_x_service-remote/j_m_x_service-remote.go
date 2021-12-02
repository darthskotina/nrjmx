// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/newrelic/nrjmx/gojmx/nrprotocol"
)

var _ = nrprotocol.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  void connect(JMXConfig config, i64 timeoutMs)")
  fmt.Fprintln(os.Stderr, "  void disconnect()")
  fmt.Fprintln(os.Stderr, "   queryMbean(string beanName, i64 timeoutMs)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := nrprotocol.NewJMXServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "connect":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "Connect requires 2 args")
      flag.Usage()
    }
    arg9 := flag.Arg(1)
    mbTrans10 := thrift.NewTMemoryBufferLen(len(arg9))
    defer mbTrans10.Close()
    _, err11 := mbTrans10.WriteString(arg9)
    if err11 != nil {
      Usage()
      return
    }
    factory12 := thrift.NewTJSONProtocolFactory()
    jsProt13 := factory12.GetProtocol(mbTrans10)
    argvalue0 := nrprotocol.NewJMXConfig()
    err14 := argvalue0.Read(jsProt13)
    if err14 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1, err15 := (strconv.ParseInt(flag.Arg(2), 10, 64))
    if err15 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.Connect(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "disconnect":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "Disconnect requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.Disconnect(context.Background()))
    fmt.Print("\n")
    break
  case "queryMbean":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "QueryMbean requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1, err17 := (strconv.ParseInt(flag.Arg(2), 10, 64))
    if err17 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.QueryMbean(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
