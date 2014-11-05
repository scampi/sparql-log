package qparser

import (
    "strconv"
    "strings"
    "regexp"
    "sort"
)

type ConnectedComponent struct {
    Body string
    Complexity []int
}

type ConnectedComponents []ConnectedComponent

func (ccs ConnectedComponents) Len() int {
    return len(ccs)
}

func (ccs ConnectedComponents) Less(i, j int) bool {
    return ccs[i].Body < ccs[j].Body
}

func (ccs ConnectedComponents) Swap(i, j int) {
    tmp := ccs[i]
    ccs[i] = ccs[j]
    ccs[j] = tmp
}

type Schema struct {
    sts map[string]map[string][]string
    cnt int
    vars map[string]string
}

func NewSchema() *Schema {
    s := Schema{}
    s.sts = make(map[string]map[string][]string)
    s.vars = make(map[string]string)
    return &s
}

// Reset initialises the SparqlGraph
func Reset(sg *SparqlGraph, query string) {
    sg.Schema = NewSchema()
    sg.Buffer = query
    sg.Init()
}

func (s *Schema) getVar(str string) string {
    if v,ok := s.vars[str]; ok {
        return v
    }
    strVar := "?v" + strconv.Itoa(s.cnt)
    s.vars[str] = strVar
    s.cnt++
    return strVar
}

func (schema *Schema) addStatement(s, p, o string) {
    if p[0] == '?' {
        return
    }
    s = schema.getVar(s)
    if p != "<http://www.w3.org/1999/02/22-rdf-syntax-ns#type>" {
        o = schema.getVar(o)
    }
    if _, ok := schema.sts[s]; ok {
        if _, ok := schema.sts[s][p]; ok {
            for _, o2 := range schema.sts[s][p] {
                if o == o2 {
                    return
                }
            }
            schema.sts[s][p] = append(schema.sts[s][p], o)
        } else {
            schema.sts[s][p] = []string{ o }
        }
    } else {
        schema.sts[s] = map[string][]string {
            p : []string{ o },
        }
    }
}

func getKey(s string, ccs map[string][]string) (string, bool) {
    for k := range ccs {
        if strings.Contains(k, s) {
            return k, true
        }
    }
    return s, false
}

func (schema *Schema) ConnectedComponents() (ar ConnectedComponents) {
    // map of connected components
    // the key is the set of variables part of a component
    ccs := make(map[string][]string)
    for s, pos := range schema.sts {
        var cc []string
        key, _ := getKey(s + "-", ccs)
        for p, os := range pos {
            for _, o := range os {
                if o[0] == '?' {
                    newkey, ok := getKey(o + "-", ccs)
                    if newkey != key {
                        cc = append(cc, ccs[key]...)
                        delete(ccs, key)
                        if !ok {
                            key += newkey
                        } else {
                            key = newkey
                        }
                    }
                }
                cc = append(cc, "    " + s + " " + p + " " + o + " .")
            }
        }
        ccs[key] = append(ccs[key], cc...)
    }
    if len(ccs) == 0 {
        return
    }
    reg := regexp.MustCompile("    \\?[^ ]+")
    for _, v := range ccs {
        sort.Strings(v)
        body := strings.Join(v, "\n") + "\n"
        cc := ConnectedComponent{ Body : body }

        m := reg.FindAllStringIndex(body, -1)
        prev, cnt := "", 0
        for _, ind := range m {
            if cnt != 0 && prev != body[ind[0]:ind[1]] {
                cc.Complexity = append(cc.Complexity, cnt)
                cnt = 0
            }
            cnt++
            prev = body[ind[0]:ind[1]]
        }
        cc.Complexity = append(cc.Complexity, cnt)
        sort.Ints(cc.Complexity)
        ar = append(ar, cc)
    }
    sort.Sort(ar)
    return
}

