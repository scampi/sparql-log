// Package extract parses SPARQL log files and retrives SPARQL queries.
// The connected components of the query's structure are extracted and
// written to files depending on a connected component complexity.
package extract

import (
    "strconv"
    "github.com/scampi/sparql-log/qparser"
    "hash/fnv"
    "hash"
    "net/url"
    "fmt"
    "regexp"
    "path"
    "github.com/golang/glog"
    "bufio"
    "strings"
    "os"
    "io/ioutil"
    "compress/bzip2"
    "compress/gzip"
)

// The format of the log files
type LogFormat uint

const (
    // Apache Combined Log Format http://httpd.apache.org/docs/current/logs.html
    TOMCAT LogFormat = iota
)

var logFormats = []string {
    "TOMCAT",
}

func (lf LogFormat) String() string {
    return logFormats[lf]
}

// Set method needed for the flag package
func (lf LogFormat) Set(s string) error {
    s = strings.ToUpper(s)
    switch s {
    case "TOMCAT":
        lf = TOMCAT
    }
    return fmt.Errorf("Unknown log format: [%v]", s)
}

// Extract process the log files in input with the given format, and dumps the
// connected components into output's subfolders by the component's complexity.
// Input log files may be Bzip2 or Gzip compressed.
func Extract(logFormat LogFormat, input, output string) {
    files, err := ioutil.ReadDir(input)
    if err != nil {
        glog.Fatal(err)
    }
    err = os.MkdirAll(output, os.ModePerm)
    if err != nil {
        glog.Fatal(err)
    }

    queries := make(map[string]*gzip.Writer)
    sg := &qparser.SparqlGraph{}
    h := fnv.New64a()
    uniq := make(map[uint64]bool)
    for _, file := range files {
        glog.Infof("Processing [%v]", file.Name())
        // Read logs
        fi, err := os.Open(path.Join(input, file.Name()))
        if err != nil {
            glog.Fatal(err)
        }
        defer fi.Close()
        var s *bufio.Scanner
        if strings.HasSuffix(fi.Name(), ".gz") {
            r, err := gzip.NewReader(fi)
            if err != nil {
                glog.Fatal(err)
            }
            defer r.Close()
            s = bufio.NewScanner(r)
        } else if strings.HasSuffix(fi.Name(), ".bz2") {
            s = bufio.NewScanner(bzip2.NewReader(fi))
        } else {
            s = bufio.NewScanner(fi)
        }
        switch logFormat {
        case TOMCAT:
            s.Split(tomcat)
        default:
            glog.Fatalf("Unknown format: [%v]", logFormat)
        }

        for s.Scan() {
            qparser.Reset(sg, s.Text())
            if err := sg.Parse(); err != nil {
                glog.Warningf("Failed to parse query\n%v\n%v", err, s.Text())
            }
            sg.Execute()
            for _, cc := range sg.ConnectedComponents() {
                if len(cc.Complexity) != 1 || cc.Complexity[0] != 1 {
                    query := "select * {\n" + cc.Body + "}\n"
                    qid := getQueryId(h, query)
                    if _, ok := uniq[qid]; !ok {
                        glog.Infof("%v%v", s.Text(), cc)
                        uniq[qid] = true
                        qc := ""
                        for i := range cc.Complexity {
                            qc += strconv.Itoa(cc.Complexity[i])
                            if i + 1 != len(cc.Complexity) {
                                qc += "-"
                            }
                        }
                        w := queries[qc]
                        if w == nil {
                            fo, err := os.OpenFile(path.Join(output, "query_" + qc + ".gz"), os.O_WRONLY | os.O_TRUNC | os.O_CREATE, os.ModePerm)
                            if err != nil {
                                glog.Fatal(err)
                            }
                            defer fo.Close()
                            defer fo.Sync()
                            w = gzip.NewWriter(fo)
                            defer w.Close()
                            queries[qc] = w
                        }
                        w.Write([]byte(query))
                        w.Write([]byte("###\n"))
                    }
                }
            }
        }
        if s.Err() != nil {
            glog.Fatal(s.Err())
        }
    }
}

var tomcatReg *regexp.Regexp = regexp.MustCompile("query=([^ ]+)")

// Tomcat reads the log file line by line and returns the decoded SPARQL query.
func tomcat(data []byte, atEOF bool) (advance int, token []byte, err error) {
    for {
        advance, token, err = bufio.ScanLines(data, atEOF)
        if err != nil {
            glog.Fatal(err)
        }
        if advance == 0 {
            return 0, nil, nil
        }
        m := tomcatReg.FindSubmatch(token)
        if m == nil {
            data = data[advance:]
            continue
        }
        dec, err := url.QueryUnescape(string(m[1]))
        if err != nil {
            glog.Fatalf("%v\n%v", m[1], err)
        }
        if ind := strings.LastIndex(dec, "}"); ind != -1 {
            // discard anything after the where clause
            dec = dec[:ind+1]
        } else {
            data = data[advance:]
            continue
        }
        token = []byte(dec)
        break
    }
    return
}

// getQueryId returns the query identifier for given query
func getQueryId(h hash.Hash64, query string) uint64 {
    h.Reset()
    h.Write([]byte(query))
    return h.Sum64()
}

