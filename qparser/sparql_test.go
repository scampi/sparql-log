package qparser

import (
    "testing"
    "strings"
    "sort"
    "reflect"
)

func sortPatterns(data ConnectedComponents) {
    for i := range data {
        ar := strings.Split(data[i].Body, "\n")
        sort.Strings(ar)
        data[i].Body = strings.Join(ar, "\n") + "\n"
    }
    sort.Sort(data)
}

func assert(t *testing.T, query string, expected ConnectedComponents) {
    sg := &SparqlGraph{}
    Reset(sg, query)
    if err := sg.Parse(); err != nil {
        t.Errorf("Failed to parse query\n%v", err)
    }
    sg.Execute()
    actual := sg.ConnectedComponents()
    if len(actual) != len(expected) {
        t.Errorf("Expected %v, but got %v", len(expected), len(actual))
    }
    sortPatterns(expected)
    if !reflect.DeepEqual(expected, actual) {
        t.Errorf("Expected %v, but got %v", expected, actual)
    }
}

func TestEmptyPnLocal(t *testing.T) {
    q := `
    Select ?name ?population ?lat ?long from lgd: 
    { 
        ?s ?p ?o
    }
    `
    assert(t, q, nil)
}

func TestService(t *testing.T) {
    q := `
    SELECT * WHERE {
        SERVICE <http://kegg.bio2rdf.org/sparql> {
            ?enzyme <http://bio2rdf.org/ns/kegg#xSubstrate> ?cpd.
            ?enzyme a <http://bio2rdf.org/ns/kegg#Enzyme>.
            ?reaction <http://bio2rdf.org/ns/kegg#xEnzyme> ?enzyme.
            ?reaction <http://bio2rdf.org/ns/kegg#equation> ?equation.
        }
    }
    `
    expected := ConnectedComponents{
        ConnectedComponent{
`    ?v0 <http://bio2rdf.org/ns/kegg#xSubstrate> ?v1 .
    ?v0 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://bio2rdf.org/ns/kegg#Enzyme> .
    ?v2 <http://bio2rdf.org/ns/kegg#xEnzyme> ?v0 .
    ?v2 <http://bio2rdf.org/ns/kegg#equation> ?v3 .
`,
            []int{ 2, 2 },
        },
    }
    assert(t, q, expected)
}

func TestLiteral1(t *testing.T) {
    q := `
    SELECT ?value WHERE {
        ?value <http://xmlns.com/foaf/0.1/name> "\u00DCberSoldier"@en .
    }
    LIMIT 1000
    `
    expected := ConnectedComponents{
        ConnectedComponent{
            "    ?v0 <http://xmlns.com/foaf/0.1/name> ?v1 .\n",
            []int{ 1 },
        },
    }
    assert(t, q, expected)
}

func TestLiteral2(t *testing.T) {
    q := `
    SELECT ?value WHERE {
        ?value <http://www.w3.org/2000/01/rdf-schema#comment> "\u041A"@ru .
    }
    LIMIT 10
    `
    expected := ConnectedComponents{
        ConnectedComponent{
            "    ?v0 <http://www.w3.org/2000/01/rdf-schema#comment> ?v1 .\n",
            []int{ 1 },
        },
    }
    assert(t, q, expected)
}

func TestOnePattern(t *testing.T) {
    q := `
    ASK
    WHERE
      { ?s  ?p  ?o . }
    `
    assert(t, q, nil)
}

func TestStar1(t *testing.T) {
    q := `
    select * {
        ?s a <:Person>; <name> "toto"
    }
    `
    expected := ConnectedComponents{
        ConnectedComponent{
            "    ?v0 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <:Person> .\n" +
            "    ?v0 <name> ?v1 .\n",
            []int{ 2 },
        },
    }
    assert(t, q, expected)
}

func TestStar2(t *testing.T) {
    q := `
    select * {
        ?s a <:Person>; <name> "toto" .
        ?o a <:Person>; <age> "42"
    }
    `
    expected := ConnectedComponents{
        ConnectedComponent{
            "    ?v0 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <:Person> .\n" +
            "    ?v0 <name> ?v1 .\n",
            []int{ 2 },
        }, ConnectedComponent{
            "    ?v2 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <:Person> .\n" +
            "    ?v2 <age> ?v3 .\n",
            []int{ 2 },
        },
    }
    assert(t, q, expected)
}

func TestStar3(t *testing.T) {
    q := `
    select * {
        ?s <name> "toto" .
        ?o <age> "42" .
        ?s a <:Person> .
        ?o a <:Person> .
    }
    `
    expected := ConnectedComponents{
        ConnectedComponent{
            "    ?v0 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <:Person> .\n" +
            "    ?v0 <name> ?v1 .\n",
            []int{ 2 },
        }, ConnectedComponent{
            "    ?v2 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <:Person> .\n" +
            "    ?v2 <age> ?v3 .\n",
            []int{ 2 },
        },
    }
    assert(t, q, expected)
}

func TestFilter(t *testing.T) {
    q := `
    select * {
        ?s <name> "toto" .
        ?s a <:Person> .
        filter isuri(?s)
    }
    `
    expected := ConnectedComponents{
        ConnectedComponent{
            "    ?v0 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <:Person> .\n" +
            "    ?v0 <name> ?v1 .\n",
            []int{ 2 },
        },
    }
    assert(t, q, expected)
}

func TestPath1(t *testing.T) {
    q := `
    select * {
        ?s a <:Person>; <name> "toto"; <knows> ?o .
        ?o a <:Person>; <age> "42"
    }
    `
    expected := ConnectedComponents{
        ConnectedComponent{
            "    ?v0 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <:Person> .\n" +
            "    ?v0 <name> ?v1 .\n" +
            "    ?v0 <knows> ?v2 .\n" +
            "    ?v2 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <:Person> .\n" +
            "    ?v2 <age> ?v3 .\n",
            []int{ 2, 3 },
        },
    }
    assert(t, q, expected)
}

func TestPath2(t *testing.T) {
    q := `
    select * {
        ?s a <:Person>; <name> "toto" .
        ?o a <:Person>; <age> "42"; <knows> ?s .
    }
    `
    expected := ConnectedComponents{
        ConnectedComponent{
            "    ?v0 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <:Person> .\n" +
            "    ?v0 <name> ?v1 .\n" +
            "    ?v2 <knows> ?v0 .\n" +
            "    ?v2 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <:Person> .\n" +
            "    ?v2 <age> ?v3 .\n",
            []int{ 2, 3 },
        },
    }
    assert(t, q, expected)
}

