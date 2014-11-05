package main

import (
    "github.com/golang/glog"
    "github.com/scampi/sparql-log/extract"
    "os"
    "flag"
    "fmt"
)

var logFormat extract.LogFormat
var input = flag.String("input", "", "The path to the log folder")
var output = flag.String("output", "", "The path to the output folder")

func init() {
    flag.Var(&logFormat, "log-format", "The format of the logs")
}

func missingOption(option string) {
    fmt.Println("Missing option -" + option)
    flag.Usage()
    os.Exit(1)
}

func main() {
    flag.Parse()
    defer glog.Flush()

    if *input == "" { missingOption("input") }
    if *output == "" { missingOption("output") }
    extract.Extract(logFormat, *input, *output)
}

