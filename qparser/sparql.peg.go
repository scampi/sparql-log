package qparser

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

const end_symbol rune = 4

/* The rule types inferred from the grammar are below. */
type pegRule uint8

const (
	ruleUnknown pegRule = iota
	rulequeryContainer
	ruleprolog
	ruleprefixDecl
	rulebaseDecl
	rulequery
	ruleselectQuery
	ruleselect
	rulesubSelect
	ruleconstructQuery
	ruleconstruct
	ruledescribeQuery
	ruledescribe
	ruleaskQuery
	ruleprojectionElem
	ruledatasetClause
	rulewhereClause
	rulegroupGraphPattern
	rulegraphPattern
	rulegraphPatternNotTriples
	ruleserviceGraphPattern
	ruleoptionalGraphPattern
	rulegroupOrUnionGraphPattern
	rulegraphGraphPattern
	ruleminusGraphPattern
	rulebasicGraphPattern
	rulefilterOrBind
	ruleconstraint
	ruletriplesBlock
	ruletriplesSameSubjectPath
	rulevarOrTerm
	rulegraphTerm
	ruletriplesNodePath
	rulecollectionPath
	ruleblankNodePropertyListPath
	rulepropertyListPath
	ruleverbPath
	rulepath
	rulepathAlternative
	rulepathSequence
	rulepathElt
	rulepathPrimary
	rulepathNegatedPropertySet
	rulepathOneInPropertySet
	rulepathMod
	ruleobjectListPath
	ruleobjectPath
	rulegraphNodePath
	rulesolutionModifier
	rulegroupCondition
	ruleorderCondition
	rulelimitOffsetClauses
	rulelimit
	ruleoffset
	ruleexpression
	ruleconditionalOrExpression
	ruleconditionalAndExpression
	rulevalueLogical
	rulenumericExpression
	rulemultiplicativeExpression
	ruleunaryExpression
	ruleprimaryExpression
	rulebrackettedExpression
	rulefunctionCall
	rulein
	rulenotin
	ruleargList
	ruleaggregate
	rulecount
	rulegroupConcat
	rulebuiltinCall
	rulevar
	ruleiriref
	ruleiri
	ruleprefixedName
	ruleliteral
	rulestring
	rulestringLiteralA
	rulestringLiteralB
	rulestringLiteralLongA
	rulestringLiteralLongB
	ruleechar
	rulenumericLiteral
	rulesignedNumericLiteral
	rulebooleanLiteral
	ruleblankNode
	ruleblankNodeLabel
	ruleanon
	rulenil
	ruleVARNAME
	rulepnPrefix
	rulepnLocal
	rulepnChars
	rulepnCharsU
	rulepnCharsBase
	ruleplx
	rulepercent
	rulehex
	rulepnLocalEsc
	rulePREFIX
	ruleTRUE
	ruleFALSE
	ruleBASE
	ruleSELECT
	ruleREDUCED
	ruleDISTINCT
	ruleFROM
	ruleNAMED
	ruleWHERE
	ruleLBRACE
	ruleRBRACE
	ruleLBRACK
	ruleRBRACK
	ruleSEMICOLON
	ruleCOMMA
	ruleDOT
	ruleCOLON
	rulePIPE
	ruleSLASH
	ruleINVERSE
	ruleLPAREN
	ruleRPAREN
	ruleISA
	ruleNOT
	ruleSTAR
	ruleQUESTION
	rulePLUS
	ruleMINUS
	ruleOPTIONAL
	ruleUNION
	ruleLIMIT
	ruleOFFSET
	ruleINTEGER
	ruleCONSTRUCT
	ruleDESCRIBE
	ruleASK
	ruleOR
	ruleAND
	ruleEQ
	ruleNE
	ruleGT
	ruleLT
	ruleLE
	ruleGE
	ruleIN
	ruleNOTIN
	ruleAS
	ruleSTR
	ruleLANG
	ruleDATATYPE
	ruleIRI
	ruleURI
	ruleABS
	ruleCEIL
	ruleROUND
	ruleFLOOR
	ruleSTRLEN
	ruleUCASE
	ruleLCASE
	ruleENCODEFORURI
	ruleYEAR
	ruleMONTH
	ruleDAY
	ruleHOURS
	ruleMINUTES
	ruleSECONDS
	ruleTIMEZONE
	ruleTZ
	ruleMD5
	ruleSHA1
	ruleSHA256
	ruleSHA384
	ruleSHA512
	ruleISIRI
	ruleISURI
	ruleISBLANK
	ruleISLITERAL
	ruleISNUMERIC
	ruleLANGMATCHES
	ruleCONTAINS
	ruleSTRSTARTS
	ruleSTRENDS
	ruleSTRBEFORE
	ruleSTRAFTER
	ruleSTRLANG
	ruleSTRDT
	ruleSAMETERM
	ruleBOUND
	ruleBNODE
	ruleRAND
	ruleNOW
	ruleUUID
	ruleSTRUUID
	ruleCONCAT
	ruleSUBSTR
	ruleREPLACE
	ruleREGEX
	ruleIF
	ruleEXISTS
	ruleNOTEXIST
	ruleCOALESCE
	ruleFILTER
	ruleBIND
	ruleSUM
	ruleMIN
	ruleMAX
	ruleAVG
	ruleSAMPLE
	ruleCOUNT
	ruleGROUPCONCAT
	ruleSEPARATOR
	ruleASC
	ruleDESC
	ruleORDER
	ruleGROUP
	ruleBY
	ruleHAVING
	ruleGRAPH
	ruleMINUSSETOPER
	ruleSERVICE
	ruleSILENT
	ruleskip
	rulews
	rulecomment
	ruleendOfLine
	rulePegText
	ruleAction0
	ruleAction1
	ruleAction2
	ruleAction3
	ruleAction4
	ruleAction5
	ruleAction6

	rulePre_
	rule_In_
	rule_Suf
)

var rul3s = [...]string{
	"Unknown",
	"queryContainer",
	"prolog",
	"prefixDecl",
	"baseDecl",
	"query",
	"selectQuery",
	"select",
	"subSelect",
	"constructQuery",
	"construct",
	"describeQuery",
	"describe",
	"askQuery",
	"projectionElem",
	"datasetClause",
	"whereClause",
	"groupGraphPattern",
	"graphPattern",
	"graphPatternNotTriples",
	"serviceGraphPattern",
	"optionalGraphPattern",
	"groupOrUnionGraphPattern",
	"graphGraphPattern",
	"minusGraphPattern",
	"basicGraphPattern",
	"filterOrBind",
	"constraint",
	"triplesBlock",
	"triplesSameSubjectPath",
	"varOrTerm",
	"graphTerm",
	"triplesNodePath",
	"collectionPath",
	"blankNodePropertyListPath",
	"propertyListPath",
	"verbPath",
	"path",
	"pathAlternative",
	"pathSequence",
	"pathElt",
	"pathPrimary",
	"pathNegatedPropertySet",
	"pathOneInPropertySet",
	"pathMod",
	"objectListPath",
	"objectPath",
	"graphNodePath",
	"solutionModifier",
	"groupCondition",
	"orderCondition",
	"limitOffsetClauses",
	"limit",
	"offset",
	"expression",
	"conditionalOrExpression",
	"conditionalAndExpression",
	"valueLogical",
	"numericExpression",
	"multiplicativeExpression",
	"unaryExpression",
	"primaryExpression",
	"brackettedExpression",
	"functionCall",
	"in",
	"notin",
	"argList",
	"aggregate",
	"count",
	"groupConcat",
	"builtinCall",
	"var",
	"iriref",
	"iri",
	"prefixedName",
	"literal",
	"string",
	"stringLiteralA",
	"stringLiteralB",
	"stringLiteralLongA",
	"stringLiteralLongB",
	"echar",
	"numericLiteral",
	"signedNumericLiteral",
	"booleanLiteral",
	"blankNode",
	"blankNodeLabel",
	"anon",
	"nil",
	"VARNAME",
	"pnPrefix",
	"pnLocal",
	"pnChars",
	"pnCharsU",
	"pnCharsBase",
	"plx",
	"percent",
	"hex",
	"pnLocalEsc",
	"PREFIX",
	"TRUE",
	"FALSE",
	"BASE",
	"SELECT",
	"REDUCED",
	"DISTINCT",
	"FROM",
	"NAMED",
	"WHERE",
	"LBRACE",
	"RBRACE",
	"LBRACK",
	"RBRACK",
	"SEMICOLON",
	"COMMA",
	"DOT",
	"COLON",
	"PIPE",
	"SLASH",
	"INVERSE",
	"LPAREN",
	"RPAREN",
	"ISA",
	"NOT",
	"STAR",
	"QUESTION",
	"PLUS",
	"MINUS",
	"OPTIONAL",
	"UNION",
	"LIMIT",
	"OFFSET",
	"INTEGER",
	"CONSTRUCT",
	"DESCRIBE",
	"ASK",
	"OR",
	"AND",
	"EQ",
	"NE",
	"GT",
	"LT",
	"LE",
	"GE",
	"IN",
	"NOTIN",
	"AS",
	"STR",
	"LANG",
	"DATATYPE",
	"IRI",
	"URI",
	"ABS",
	"CEIL",
	"ROUND",
	"FLOOR",
	"STRLEN",
	"UCASE",
	"LCASE",
	"ENCODEFORURI",
	"YEAR",
	"MONTH",
	"DAY",
	"HOURS",
	"MINUTES",
	"SECONDS",
	"TIMEZONE",
	"TZ",
	"MD5",
	"SHA1",
	"SHA256",
	"SHA384",
	"SHA512",
	"ISIRI",
	"ISURI",
	"ISBLANK",
	"ISLITERAL",
	"ISNUMERIC",
	"LANGMATCHES",
	"CONTAINS",
	"STRSTARTS",
	"STRENDS",
	"STRBEFORE",
	"STRAFTER",
	"STRLANG",
	"STRDT",
	"SAMETERM",
	"BOUND",
	"BNODE",
	"RAND",
	"NOW",
	"UUID",
	"STRUUID",
	"CONCAT",
	"SUBSTR",
	"REPLACE",
	"REGEX",
	"IF",
	"EXISTS",
	"NOTEXIST",
	"COALESCE",
	"FILTER",
	"BIND",
	"SUM",
	"MIN",
	"MAX",
	"AVG",
	"SAMPLE",
	"COUNT",
	"GROUPCONCAT",
	"SEPARATOR",
	"ASC",
	"DESC",
	"ORDER",
	"GROUP",
	"BY",
	"HAVING",
	"GRAPH",
	"MINUSSETOPER",
	"SERVICE",
	"SILENT",
	"skip",
	"ws",
	"comment",
	"endOfLine",
	"PegText",
	"Action0",
	"Action1",
	"Action2",
	"Action3",
	"Action4",
	"Action5",
	"Action6",

	"Pre_",
	"_In_",
	"_Suf",
}

type tokenTree interface {
	Print()
	PrintSyntax()
	PrintSyntaxTree(buffer string)
	Add(rule pegRule, begin, end, next, depth int)
	Expand(index int) tokenTree
	Tokens() <-chan token32
	AST() *node32
	Error() []token32
	trim(length int)
}

type node32 struct {
	token32
	up, next *node32
}

func (node *node32) print(depth int, buffer string) {
	for node != nil {
		for c := 0; c < depth; c++ {
			fmt.Printf(" ")
		}
		fmt.Printf("\x1B[34m%v\x1B[m %v\n", rul3s[node.pegRule], strconv.Quote(buffer[node.begin:node.end]))
		if node.up != nil {
			node.up.print(depth+1, buffer)
		}
		node = node.next
	}
}

func (ast *node32) Print(buffer string) {
	ast.print(0, buffer)
}

type element struct {
	node *node32
	down *element
}

/* ${@} bit structure for abstract syntax tree */
type token16 struct {
	pegRule
	begin, end, next int16
}

func (t *token16) isZero() bool {
	return t.pegRule == ruleUnknown && t.begin == 0 && t.end == 0 && t.next == 0
}

func (t *token16) isParentOf(u token16) bool {
	return t.begin <= u.begin && t.end >= u.end && t.next > u.next
}

func (t *token16) getToken32() token32 {
	return token32{pegRule: t.pegRule, begin: int32(t.begin), end: int32(t.end), next: int32(t.next)}
}

func (t *token16) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v %v", rul3s[t.pegRule], t.begin, t.end, t.next)
}

type tokens16 struct {
	tree    []token16
	ordered [][]token16
}

func (t *tokens16) trim(length int) {
	t.tree = t.tree[0:length]
}

func (t *tokens16) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens16) Order() [][]token16 {
	if t.ordered != nil {
		return t.ordered
	}

	depths := make([]int16, 1, math.MaxInt16)
	for i, token := range t.tree {
		if token.pegRule == ruleUnknown {
			t.tree = t.tree[:i]
			break
		}
		depth := int(token.next)
		if length := len(depths); depth >= length {
			depths = depths[:depth+1]
		}
		depths[depth]++
	}
	depths = append(depths, 0)

	ordered, pool := make([][]token16, len(depths)), make([]token16, len(t.tree)+len(depths))
	for i, depth := range depths {
		depth++
		ordered[i], pool, depths[i] = pool[:depth], pool[depth:], 0
	}

	for i, token := range t.tree {
		depth := token.next
		token.next = int16(i)
		ordered[depth][depths[depth]] = token
		depths[depth]++
	}
	t.ordered = ordered
	return ordered
}

type state16 struct {
	token16
	depths []int16
	leaf   bool
}

func (t *tokens16) AST() *node32 {
	tokens := t.Tokens()
	stack := &element{node: &node32{token32: <-tokens}}
	for token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	return stack.node
}

func (t *tokens16) PreOrder() (<-chan state16, [][]token16) {
	s, ordered := make(chan state16, 6), t.Order()
	go func() {
		var states [8]state16
		for i, _ := range states {
			states[i].depths = make([]int16, len(ordered))
		}
		depths, state, depth := make([]int16, len(ordered)), 0, 1
		write := func(t token16, leaf bool) {
			S := states[state]
			state, S.pegRule, S.begin, S.end, S.next, S.leaf = (state+1)%8, t.pegRule, t.begin, t.end, int16(depth), leaf
			copy(S.depths, depths)
			s <- S
		}

		states[state].token16 = ordered[0][0]
		depths[0]++
		state++
		a, b := ordered[depth-1][depths[depth-1]-1], ordered[depth][depths[depth]]
	depthFirstSearch:
		for {
			for {
				if i := depths[depth]; i > 0 {
					if c, j := ordered[depth][i-1], depths[depth-1]; a.isParentOf(c) &&
						(j < 2 || !ordered[depth-1][j-2].isParentOf(c)) {
						if c.end != b.begin {
							write(token16{pegRule: rule_In_, begin: c.end, end: b.begin}, true)
						}
						break
					}
				}

				if a.begin < b.begin {
					write(token16{pegRule: rulePre_, begin: a.begin, end: b.begin}, true)
				}
				break
			}

			next := depth + 1
			if c := ordered[next][depths[next]]; c.pegRule != ruleUnknown && b.isParentOf(c) {
				write(b, false)
				depths[depth]++
				depth, a, b = next, b, c
				continue
			}

			write(b, true)
			depths[depth]++
			c, parent := ordered[depth][depths[depth]], true
			for {
				if c.pegRule != ruleUnknown && a.isParentOf(c) {
					b = c
					continue depthFirstSearch
				} else if parent && b.end != a.end {
					write(token16{pegRule: rule_Suf, begin: b.end, end: a.end}, true)
				}

				depth--
				if depth > 0 {
					a, b, c = ordered[depth-1][depths[depth-1]-1], a, ordered[depth][depths[depth]]
					parent = a.isParentOf(b)
					continue
				}

				break depthFirstSearch
			}
		}

		close(s)
	}()
	return s, ordered
}

func (t *tokens16) PrintSyntax() {
	tokens, ordered := t.PreOrder()
	max := -1
	for token := range tokens {
		if !token.leaf {
			fmt.Printf("%v", token.begin)
			for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
				fmt.Printf(" \x1B[36m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
			}
			fmt.Printf(" \x1B[36m%v\x1B[m\n", rul3s[token.pegRule])
		} else if token.begin == token.end {
			fmt.Printf("%v", token.begin)
			for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
				fmt.Printf(" \x1B[31m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
			}
			fmt.Printf(" \x1B[31m%v\x1B[m\n", rul3s[token.pegRule])
		} else {
			for c, end := token.begin, token.end; c < end; c++ {
				if i := int(c); max+1 < i {
					for j := max; j < i; j++ {
						fmt.Printf("skip %v %v\n", j, token.String())
					}
					max = i
				} else if i := int(c); i <= max {
					for j := i; j <= max; j++ {
						fmt.Printf("dupe %v %v\n", j, token.String())
					}
				} else {
					max = int(c)
				}
				fmt.Printf("%v", c)
				for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
					fmt.Printf(" \x1B[34m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
				}
				fmt.Printf(" \x1B[34m%v\x1B[m\n", rul3s[token.pegRule])
			}
			fmt.Printf("\n")
		}
	}
}

func (t *tokens16) PrintSyntaxTree(buffer string) {
	tokens, _ := t.PreOrder()
	for token := range tokens {
		for c := 0; c < int(token.next); c++ {
			fmt.Printf(" ")
		}
		fmt.Printf("\x1B[34m%v\x1B[m %v\n", rul3s[token.pegRule], strconv.Quote(buffer[token.begin:token.end]))
	}
}

func (t *tokens16) Add(rule pegRule, begin, end, depth, index int) {
	t.tree[index] = token16{pegRule: rule, begin: int16(begin), end: int16(end), next: int16(depth)}
}

func (t *tokens16) Tokens() <-chan token32 {
	s := make(chan token32, 16)
	go func() {
		for _, v := range t.tree {
			s <- v.getToken32()
		}
		close(s)
	}()
	return s
}

func (t *tokens16) Error() []token32 {
	ordered := t.Order()
	length := len(ordered)
	tokens, length := make([]token32, length), length-1
	for i, _ := range tokens {
		o := ordered[length-i]
		if len(o) > 1 {
			tokens[i] = o[len(o)-2].getToken32()
		}
	}
	return tokens
}

/* ${@} bit structure for abstract syntax tree */
type token32 struct {
	pegRule
	begin, end, next int32
}

func (t *token32) isZero() bool {
	return t.pegRule == ruleUnknown && t.begin == 0 && t.end == 0 && t.next == 0
}

func (t *token32) isParentOf(u token32) bool {
	return t.begin <= u.begin && t.end >= u.end && t.next > u.next
}

func (t *token32) getToken32() token32 {
	return token32{pegRule: t.pegRule, begin: int32(t.begin), end: int32(t.end), next: int32(t.next)}
}

func (t *token32) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v %v", rul3s[t.pegRule], t.begin, t.end, t.next)
}

type tokens32 struct {
	tree    []token32
	ordered [][]token32
}

func (t *tokens32) trim(length int) {
	t.tree = t.tree[0:length]
}

func (t *tokens32) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens32) Order() [][]token32 {
	if t.ordered != nil {
		return t.ordered
	}

	depths := make([]int32, 1, math.MaxInt16)
	for i, token := range t.tree {
		if token.pegRule == ruleUnknown {
			t.tree = t.tree[:i]
			break
		}
		depth := int(token.next)
		if length := len(depths); depth >= length {
			depths = depths[:depth+1]
		}
		depths[depth]++
	}
	depths = append(depths, 0)

	ordered, pool := make([][]token32, len(depths)), make([]token32, len(t.tree)+len(depths))
	for i, depth := range depths {
		depth++
		ordered[i], pool, depths[i] = pool[:depth], pool[depth:], 0
	}

	for i, token := range t.tree {
		depth := token.next
		token.next = int32(i)
		ordered[depth][depths[depth]] = token
		depths[depth]++
	}
	t.ordered = ordered
	return ordered
}

type state32 struct {
	token32
	depths []int32
	leaf   bool
}

func (t *tokens32) AST() *node32 {
	tokens := t.Tokens()
	stack := &element{node: &node32{token32: <-tokens}}
	for token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	return stack.node
}

func (t *tokens32) PreOrder() (<-chan state32, [][]token32) {
	s, ordered := make(chan state32, 6), t.Order()
	go func() {
		var states [8]state32
		for i, _ := range states {
			states[i].depths = make([]int32, len(ordered))
		}
		depths, state, depth := make([]int32, len(ordered)), 0, 1
		write := func(t token32, leaf bool) {
			S := states[state]
			state, S.pegRule, S.begin, S.end, S.next, S.leaf = (state+1)%8, t.pegRule, t.begin, t.end, int32(depth), leaf
			copy(S.depths, depths)
			s <- S
		}

		states[state].token32 = ordered[0][0]
		depths[0]++
		state++
		a, b := ordered[depth-1][depths[depth-1]-1], ordered[depth][depths[depth]]
	depthFirstSearch:
		for {
			for {
				if i := depths[depth]; i > 0 {
					if c, j := ordered[depth][i-1], depths[depth-1]; a.isParentOf(c) &&
						(j < 2 || !ordered[depth-1][j-2].isParentOf(c)) {
						if c.end != b.begin {
							write(token32{pegRule: rule_In_, begin: c.end, end: b.begin}, true)
						}
						break
					}
				}

				if a.begin < b.begin {
					write(token32{pegRule: rulePre_, begin: a.begin, end: b.begin}, true)
				}
				break
			}

			next := depth + 1
			if c := ordered[next][depths[next]]; c.pegRule != ruleUnknown && b.isParentOf(c) {
				write(b, false)
				depths[depth]++
				depth, a, b = next, b, c
				continue
			}

			write(b, true)
			depths[depth]++
			c, parent := ordered[depth][depths[depth]], true
			for {
				if c.pegRule != ruleUnknown && a.isParentOf(c) {
					b = c
					continue depthFirstSearch
				} else if parent && b.end != a.end {
					write(token32{pegRule: rule_Suf, begin: b.end, end: a.end}, true)
				}

				depth--
				if depth > 0 {
					a, b, c = ordered[depth-1][depths[depth-1]-1], a, ordered[depth][depths[depth]]
					parent = a.isParentOf(b)
					continue
				}

				break depthFirstSearch
			}
		}

		close(s)
	}()
	return s, ordered
}

func (t *tokens32) PrintSyntax() {
	tokens, ordered := t.PreOrder()
	max := -1
	for token := range tokens {
		if !token.leaf {
			fmt.Printf("%v", token.begin)
			for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
				fmt.Printf(" \x1B[36m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
			}
			fmt.Printf(" \x1B[36m%v\x1B[m\n", rul3s[token.pegRule])
		} else if token.begin == token.end {
			fmt.Printf("%v", token.begin)
			for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
				fmt.Printf(" \x1B[31m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
			}
			fmt.Printf(" \x1B[31m%v\x1B[m\n", rul3s[token.pegRule])
		} else {
			for c, end := token.begin, token.end; c < end; c++ {
				if i := int(c); max+1 < i {
					for j := max; j < i; j++ {
						fmt.Printf("skip %v %v\n", j, token.String())
					}
					max = i
				} else if i := int(c); i <= max {
					for j := i; j <= max; j++ {
						fmt.Printf("dupe %v %v\n", j, token.String())
					}
				} else {
					max = int(c)
				}
				fmt.Printf("%v", c)
				for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
					fmt.Printf(" \x1B[34m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
				}
				fmt.Printf(" \x1B[34m%v\x1B[m\n", rul3s[token.pegRule])
			}
			fmt.Printf("\n")
		}
	}
}

func (t *tokens32) PrintSyntaxTree(buffer string) {
	tokens, _ := t.PreOrder()
	for token := range tokens {
		for c := 0; c < int(token.next); c++ {
			fmt.Printf(" ")
		}
		fmt.Printf("\x1B[34m%v\x1B[m %v\n", rul3s[token.pegRule], strconv.Quote(buffer[token.begin:token.end]))
	}
}

func (t *tokens32) Add(rule pegRule, begin, end, depth, index int) {
	t.tree[index] = token32{pegRule: rule, begin: int32(begin), end: int32(end), next: int32(depth)}
}

func (t *tokens32) Tokens() <-chan token32 {
	s := make(chan token32, 16)
	go func() {
		for _, v := range t.tree {
			s <- v.getToken32()
		}
		close(s)
	}()
	return s
}

func (t *tokens32) Error() []token32 {
	ordered := t.Order()
	length := len(ordered)
	tokens, length := make([]token32, length), length-1
	for i, _ := range tokens {
		o := ordered[length-i]
		if len(o) > 1 {
			tokens[i] = o[len(o)-2].getToken32()
		}
	}
	return tokens
}

func (t *tokens16) Expand(index int) tokenTree {
	tree := t.tree
	if index >= len(tree) {
		expanded := make([]token32, 2*len(tree))
		for i, v := range tree {
			expanded[i] = v.getToken32()
		}
		return &tokens32{tree: expanded}
	}
	return nil
}

func (t *tokens32) Expand(index int) tokenTree {
	tree := t.tree
	if index >= len(tree) {
		expanded := make([]token32, 2*len(tree))
		copy(expanded, tree)
		t.tree = expanded
	}
	return nil
}

type SparqlGraph struct {
	*Schema
	label, s, p, o string

	Buffer string
	buffer []rune
	rules  [233]func() bool
	Parse  func(rule ...int) error
	Reset  func()
	tokenTree
}

type textPosition struct {
	line, symbol int
}

type textPositionMap map[int]textPosition

func translatePositions(buffer string, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

search:
	for i, c := range buffer[0:] {
		if c == '\n' {
			line, symbol = line+1, 0
		} else {
			symbol++
		}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {
				if i != positions[j] {
					continue search
				}
			}
			break search
		}
	}

	return translations
}

type parseError struct {
	p *SparqlGraph
}

func (e *parseError) Error() string {
	tokens, error := e.p.tokenTree.Error(), "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
	}
	translations := translatePositions(e.p.Buffer, positions)
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		error += fmt.Sprintf("parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n",
			rul3s[token.pegRule],
			translations[begin].line, translations[begin].symbol,
			translations[end].line, translations[end].symbol,
			/*strconv.Quote(*/ e.p.Buffer[begin:end] /*)*/)
	}

	return error
}

func (p *SparqlGraph) PrintSyntaxTree() {
	p.tokenTree.PrintSyntaxTree(p.Buffer)
}

func (p *SparqlGraph) Highlighter() {
	p.tokenTree.PrintSyntax()
}

func (p *SparqlGraph) Execute() {
	buffer, begin, end := p.Buffer, 0, 0
	for token := range p.tokenTree.Tokens() {
		switch token.pegRule {
		case rulePegText:
			begin, end = int(token.begin), int(token.end)
		case ruleAction0:
			p.s = p.label
		case ruleAction1:
			p.p = p.label
		case ruleAction2:
			p.o = p.label
			p.addStatement(p.s, p.p, p.o)
		case ruleAction3:
			p.label = buffer[begin:end]
		case ruleAction4:
			p.label = buffer[begin:end]
		case ruleAction5:
			p.label = buffer[begin:end]
		case ruleAction6:
			p.label = "<http://www.w3.org/1999/02/22-rdf-syntax-ns#type>"

		}
	}
}

func (p *SparqlGraph) Init() {
	p.buffer = []rune(p.Buffer)
	if len(p.buffer) == 0 || p.buffer[len(p.buffer)-1] != end_symbol {
		p.buffer = append(p.buffer, end_symbol)
	}

	var tree tokenTree = &tokens16{tree: make([]token16, math.MaxInt16)}
	position, depth, tokenIndex, buffer, rules := 0, 0, 0, p.buffer, p.rules

	p.Parse = func(rule ...int) error {
		r := 1
		if len(rule) > 0 {
			r = rule[0]
		}
		matches := p.rules[r]()
		p.tokenTree = tree
		if matches {
			p.tokenTree.trim(tokenIndex)
			return nil
		}
		return &parseError{p}
	}

	p.Reset = func() {
		position, tokenIndex, depth = 0, 0, 0
	}

	add := func(rule pegRule, begin int) {
		if t := tree.Expand(tokenIndex); t != nil {
			tree = t
		}
		tree.Add(rule, begin, position, depth, tokenIndex)
		tokenIndex++
	}

	matchDot := func() bool {
		if buffer[position] != end_symbol {
			position++
			return true
		}
		return false
	}

	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/

	/*matchRange := func(lower byte, upper byte) bool {
		if c := buffer[position]; c >= lower && c <= upper {
			position++
			return true
		}
		return false
	}*/

	rules = [...]func() bool{
		nil,
		/* 0 queryContainer <- <(skip prolog query !.)> */
		func() bool {
			position0, tokenIndex0, depth0 := position, tokenIndex, depth
			{
				position1 := position
				depth++
				if !rules[ruleskip]() {
					goto l0
				}
				{
					position2 := position
					depth++
				l3:
					{
						position4, tokenIndex4, depth4 := position, tokenIndex, depth
						{
							position5, tokenIndex5, depth5 := position, tokenIndex, depth
							{
								position7 := position
								depth++
								{
									position8 := position
									depth++
									{
										position9, tokenIndex9, depth9 := position, tokenIndex, depth
										if buffer[position] != rune('p') {
											goto l10
										}
										position++
										goto l9
									l10:
										position, tokenIndex, depth = position9, tokenIndex9, depth9
										if buffer[position] != rune('P') {
											goto l6
										}
										position++
									}
								l9:
									{
										position11, tokenIndex11, depth11 := position, tokenIndex, depth
										if buffer[position] != rune('r') {
											goto l12
										}
										position++
										goto l11
									l12:
										position, tokenIndex, depth = position11, tokenIndex11, depth11
										if buffer[position] != rune('R') {
											goto l6
										}
										position++
									}
								l11:
									{
										position13, tokenIndex13, depth13 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l14
										}
										position++
										goto l13
									l14:
										position, tokenIndex, depth = position13, tokenIndex13, depth13
										if buffer[position] != rune('E') {
											goto l6
										}
										position++
									}
								l13:
									{
										position15, tokenIndex15, depth15 := position, tokenIndex, depth
										if buffer[position] != rune('f') {
											goto l16
										}
										position++
										goto l15
									l16:
										position, tokenIndex, depth = position15, tokenIndex15, depth15
										if buffer[position] != rune('F') {
											goto l6
										}
										position++
									}
								l15:
									{
										position17, tokenIndex17, depth17 := position, tokenIndex, depth
										if buffer[position] != rune('i') {
											goto l18
										}
										position++
										goto l17
									l18:
										position, tokenIndex, depth = position17, tokenIndex17, depth17
										if buffer[position] != rune('I') {
											goto l6
										}
										position++
									}
								l17:
									{
										position19, tokenIndex19, depth19 := position, tokenIndex, depth
										if buffer[position] != rune('x') {
											goto l20
										}
										position++
										goto l19
									l20:
										position, tokenIndex, depth = position19, tokenIndex19, depth19
										if buffer[position] != rune('X') {
											goto l6
										}
										position++
									}
								l19:
									if !rules[ruleskip]() {
										goto l6
									}
									depth--
									add(rulePREFIX, position8)
								}
								{
									position21, tokenIndex21, depth21 := position, tokenIndex, depth
									if !rules[rulepnPrefix]() {
										goto l21
									}
									goto l22
								l21:
									position, tokenIndex, depth = position21, tokenIndex21, depth21
								}
							l22:
								{
									position23 := position
									depth++
									if buffer[position] != rune(':') {
										goto l6
									}
									position++
									if !rules[ruleskip]() {
										goto l6
									}
									depth--
									add(ruleCOLON, position23)
								}
								if !rules[ruleiri]() {
									goto l6
								}
								depth--
								add(ruleprefixDecl, position7)
							}
							goto l5
						l6:
							position, tokenIndex, depth = position5, tokenIndex5, depth5
							{
								position24 := position
								depth++
								{
									position25 := position
									depth++
									{
										position26, tokenIndex26, depth26 := position, tokenIndex, depth
										if buffer[position] != rune('b') {
											goto l27
										}
										position++
										goto l26
									l27:
										position, tokenIndex, depth = position26, tokenIndex26, depth26
										if buffer[position] != rune('B') {
											goto l4
										}
										position++
									}
								l26:
									{
										position28, tokenIndex28, depth28 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l29
										}
										position++
										goto l28
									l29:
										position, tokenIndex, depth = position28, tokenIndex28, depth28
										if buffer[position] != rune('A') {
											goto l4
										}
										position++
									}
								l28:
									{
										position30, tokenIndex30, depth30 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l31
										}
										position++
										goto l30
									l31:
										position, tokenIndex, depth = position30, tokenIndex30, depth30
										if buffer[position] != rune('S') {
											goto l4
										}
										position++
									}
								l30:
									{
										position32, tokenIndex32, depth32 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l33
										}
										position++
										goto l32
									l33:
										position, tokenIndex, depth = position32, tokenIndex32, depth32
										if buffer[position] != rune('E') {
											goto l4
										}
										position++
									}
								l32:
									if !rules[ruleskip]() {
										goto l4
									}
									depth--
									add(ruleBASE, position25)
								}
								if !rules[ruleiri]() {
									goto l4
								}
								depth--
								add(rulebaseDecl, position24)
							}
						}
					l5:
						goto l3
					l4:
						position, tokenIndex, depth = position4, tokenIndex4, depth4
					}
					depth--
					add(ruleprolog, position2)
				}
				{
					position34 := position
					depth++
					{
						switch buffer[position] {
						case 'A', 'a':
							{
								position36 := position
								depth++
								{
									position37 := position
									depth++
									{
										position38, tokenIndex38, depth38 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l39
										}
										position++
										goto l38
									l39:
										position, tokenIndex, depth = position38, tokenIndex38, depth38
										if buffer[position] != rune('A') {
											goto l0
										}
										position++
									}
								l38:
									{
										position40, tokenIndex40, depth40 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l41
										}
										position++
										goto l40
									l41:
										position, tokenIndex, depth = position40, tokenIndex40, depth40
										if buffer[position] != rune('S') {
											goto l0
										}
										position++
									}
								l40:
									{
										position42, tokenIndex42, depth42 := position, tokenIndex, depth
										if buffer[position] != rune('k') {
											goto l43
										}
										position++
										goto l42
									l43:
										position, tokenIndex, depth = position42, tokenIndex42, depth42
										if buffer[position] != rune('K') {
											goto l0
										}
										position++
									}
								l42:
									if !rules[ruleskip]() {
										goto l0
									}
									depth--
									add(ruleASK, position37)
								}
							l44:
								{
									position45, tokenIndex45, depth45 := position, tokenIndex, depth
									if !rules[ruledatasetClause]() {
										goto l45
									}
									goto l44
								l45:
									position, tokenIndex, depth = position45, tokenIndex45, depth45
								}
								if !rules[rulewhereClause]() {
									goto l0
								}
								depth--
								add(ruleaskQuery, position36)
							}
							break
						case 'D', 'd':
							{
								position46 := position
								depth++
								{
									position47 := position
									depth++
									{
										position48 := position
										depth++
										{
											position49, tokenIndex49, depth49 := position, tokenIndex, depth
											if buffer[position] != rune('d') {
												goto l50
											}
											position++
											goto l49
										l50:
											position, tokenIndex, depth = position49, tokenIndex49, depth49
											if buffer[position] != rune('D') {
												goto l0
											}
											position++
										}
									l49:
										{
											position51, tokenIndex51, depth51 := position, tokenIndex, depth
											if buffer[position] != rune('e') {
												goto l52
											}
											position++
											goto l51
										l52:
											position, tokenIndex, depth = position51, tokenIndex51, depth51
											if buffer[position] != rune('E') {
												goto l0
											}
											position++
										}
									l51:
										{
											position53, tokenIndex53, depth53 := position, tokenIndex, depth
											if buffer[position] != rune('s') {
												goto l54
											}
											position++
											goto l53
										l54:
											position, tokenIndex, depth = position53, tokenIndex53, depth53
											if buffer[position] != rune('S') {
												goto l0
											}
											position++
										}
									l53:
										{
											position55, tokenIndex55, depth55 := position, tokenIndex, depth
											if buffer[position] != rune('c') {
												goto l56
											}
											position++
											goto l55
										l56:
											position, tokenIndex, depth = position55, tokenIndex55, depth55
											if buffer[position] != rune('C') {
												goto l0
											}
											position++
										}
									l55:
										{
											position57, tokenIndex57, depth57 := position, tokenIndex, depth
											if buffer[position] != rune('r') {
												goto l58
											}
											position++
											goto l57
										l58:
											position, tokenIndex, depth = position57, tokenIndex57, depth57
											if buffer[position] != rune('R') {
												goto l0
											}
											position++
										}
									l57:
										{
											position59, tokenIndex59, depth59 := position, tokenIndex, depth
											if buffer[position] != rune('i') {
												goto l60
											}
											position++
											goto l59
										l60:
											position, tokenIndex, depth = position59, tokenIndex59, depth59
											if buffer[position] != rune('I') {
												goto l0
											}
											position++
										}
									l59:
										{
											position61, tokenIndex61, depth61 := position, tokenIndex, depth
											if buffer[position] != rune('b') {
												goto l62
											}
											position++
											goto l61
										l62:
											position, tokenIndex, depth = position61, tokenIndex61, depth61
											if buffer[position] != rune('B') {
												goto l0
											}
											position++
										}
									l61:
										{
											position63, tokenIndex63, depth63 := position, tokenIndex, depth
											if buffer[position] != rune('e') {
												goto l64
											}
											position++
											goto l63
										l64:
											position, tokenIndex, depth = position63, tokenIndex63, depth63
											if buffer[position] != rune('E') {
												goto l0
											}
											position++
										}
									l63:
										if !rules[ruleskip]() {
											goto l0
										}
										depth--
										add(ruleDESCRIBE, position48)
									}
									{
										switch buffer[position] {
										case '*':
											if !rules[ruleSTAR]() {
												goto l0
											}
											break
										case '$', '?':
											if !rules[rulevar]() {
												goto l0
											}
											break
										default:
											if !rules[ruleiriref]() {
												goto l0
											}
											break
										}
									}

									depth--
									add(ruledescribe, position47)
								}
							l66:
								{
									position67, tokenIndex67, depth67 := position, tokenIndex, depth
									if !rules[ruledatasetClause]() {
										goto l67
									}
									goto l66
								l67:
									position, tokenIndex, depth = position67, tokenIndex67, depth67
								}
								{
									position68, tokenIndex68, depth68 := position, tokenIndex, depth
									if !rules[rulewhereClause]() {
										goto l68
									}
									goto l69
								l68:
									position, tokenIndex, depth = position68, tokenIndex68, depth68
								}
							l69:
								if !rules[rulesolutionModifier]() {
									goto l0
								}
								depth--
								add(ruledescribeQuery, position46)
							}
							break
						case 'C', 'c':
							{
								position70 := position
								depth++
								{
									position71 := position
									depth++
									{
										position72 := position
										depth++
										{
											position73, tokenIndex73, depth73 := position, tokenIndex, depth
											if buffer[position] != rune('c') {
												goto l74
											}
											position++
											goto l73
										l74:
											position, tokenIndex, depth = position73, tokenIndex73, depth73
											if buffer[position] != rune('C') {
												goto l0
											}
											position++
										}
									l73:
										{
											position75, tokenIndex75, depth75 := position, tokenIndex, depth
											if buffer[position] != rune('o') {
												goto l76
											}
											position++
											goto l75
										l76:
											position, tokenIndex, depth = position75, tokenIndex75, depth75
											if buffer[position] != rune('O') {
												goto l0
											}
											position++
										}
									l75:
										{
											position77, tokenIndex77, depth77 := position, tokenIndex, depth
											if buffer[position] != rune('n') {
												goto l78
											}
											position++
											goto l77
										l78:
											position, tokenIndex, depth = position77, tokenIndex77, depth77
											if buffer[position] != rune('N') {
												goto l0
											}
											position++
										}
									l77:
										{
											position79, tokenIndex79, depth79 := position, tokenIndex, depth
											if buffer[position] != rune('s') {
												goto l80
											}
											position++
											goto l79
										l80:
											position, tokenIndex, depth = position79, tokenIndex79, depth79
											if buffer[position] != rune('S') {
												goto l0
											}
											position++
										}
									l79:
										{
											position81, tokenIndex81, depth81 := position, tokenIndex, depth
											if buffer[position] != rune('t') {
												goto l82
											}
											position++
											goto l81
										l82:
											position, tokenIndex, depth = position81, tokenIndex81, depth81
											if buffer[position] != rune('T') {
												goto l0
											}
											position++
										}
									l81:
										{
											position83, tokenIndex83, depth83 := position, tokenIndex, depth
											if buffer[position] != rune('r') {
												goto l84
											}
											position++
											goto l83
										l84:
											position, tokenIndex, depth = position83, tokenIndex83, depth83
											if buffer[position] != rune('R') {
												goto l0
											}
											position++
										}
									l83:
										{
											position85, tokenIndex85, depth85 := position, tokenIndex, depth
											if buffer[position] != rune('u') {
												goto l86
											}
											position++
											goto l85
										l86:
											position, tokenIndex, depth = position85, tokenIndex85, depth85
											if buffer[position] != rune('U') {
												goto l0
											}
											position++
										}
									l85:
										{
											position87, tokenIndex87, depth87 := position, tokenIndex, depth
											if buffer[position] != rune('c') {
												goto l88
											}
											position++
											goto l87
										l88:
											position, tokenIndex, depth = position87, tokenIndex87, depth87
											if buffer[position] != rune('C') {
												goto l0
											}
											position++
										}
									l87:
										{
											position89, tokenIndex89, depth89 := position, tokenIndex, depth
											if buffer[position] != rune('t') {
												goto l90
											}
											position++
											goto l89
										l90:
											position, tokenIndex, depth = position89, tokenIndex89, depth89
											if buffer[position] != rune('T') {
												goto l0
											}
											position++
										}
									l89:
										if !rules[ruleskip]() {
											goto l0
										}
										depth--
										add(ruleCONSTRUCT, position72)
									}
									if !rules[ruleLBRACE]() {
										goto l0
									}
									{
										position91, tokenIndex91, depth91 := position, tokenIndex, depth
										if !rules[ruletriplesBlock]() {
											goto l91
										}
										goto l92
									l91:
										position, tokenIndex, depth = position91, tokenIndex91, depth91
									}
								l92:
									if !rules[ruleRBRACE]() {
										goto l0
									}
									depth--
									add(ruleconstruct, position71)
								}
							l93:
								{
									position94, tokenIndex94, depth94 := position, tokenIndex, depth
									if !rules[ruledatasetClause]() {
										goto l94
									}
									goto l93
								l94:
									position, tokenIndex, depth = position94, tokenIndex94, depth94
								}
								if !rules[rulewhereClause]() {
									goto l0
								}
								if !rules[rulesolutionModifier]() {
									goto l0
								}
								depth--
								add(ruleconstructQuery, position70)
							}
							break
						default:
							{
								position95 := position
								depth++
								if !rules[ruleselect]() {
									goto l0
								}
							l96:
								{
									position97, tokenIndex97, depth97 := position, tokenIndex, depth
									if !rules[ruledatasetClause]() {
										goto l97
									}
									goto l96
								l97:
									position, tokenIndex, depth = position97, tokenIndex97, depth97
								}
								if !rules[rulewhereClause]() {
									goto l0
								}
								if !rules[rulesolutionModifier]() {
									goto l0
								}
								depth--
								add(ruleselectQuery, position95)
							}
							break
						}
					}

					depth--
					add(rulequery, position34)
				}
				{
					position98, tokenIndex98, depth98 := position, tokenIndex, depth
					if !matchDot() {
						goto l98
					}
					goto l0
				l98:
					position, tokenIndex, depth = position98, tokenIndex98, depth98
				}
				depth--
				add(rulequeryContainer, position1)
			}
			return true
		l0:
			position, tokenIndex, depth = position0, tokenIndex0, depth0
			return false
		},
		/* 1 prolog <- <(prefixDecl / baseDecl)*> */
		nil,
		/* 2 prefixDecl <- <(PREFIX pnPrefix? COLON iri)> */
		nil,
		/* 3 baseDecl <- <(BASE iri)> */
		nil,
		/* 4 query <- <((&('A' | 'a') askQuery) | (&('D' | 'd') describeQuery) | (&('C' | 'c') constructQuery) | (&('S' | 's') selectQuery))> */
		nil,
		/* 5 selectQuery <- <(select datasetClause* whereClause solutionModifier)> */
		nil,
		/* 6 select <- <(SELECT (DISTINCT / REDUCED)? (STAR / projectionElem+))> */
		func() bool {
			position104, tokenIndex104, depth104 := position, tokenIndex, depth
			{
				position105 := position
				depth++
				{
					position106 := position
					depth++
					{
						position107, tokenIndex107, depth107 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l108
						}
						position++
						goto l107
					l108:
						position, tokenIndex, depth = position107, tokenIndex107, depth107
						if buffer[position] != rune('S') {
							goto l104
						}
						position++
					}
				l107:
					{
						position109, tokenIndex109, depth109 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l110
						}
						position++
						goto l109
					l110:
						position, tokenIndex, depth = position109, tokenIndex109, depth109
						if buffer[position] != rune('E') {
							goto l104
						}
						position++
					}
				l109:
					{
						position111, tokenIndex111, depth111 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l112
						}
						position++
						goto l111
					l112:
						position, tokenIndex, depth = position111, tokenIndex111, depth111
						if buffer[position] != rune('L') {
							goto l104
						}
						position++
					}
				l111:
					{
						position113, tokenIndex113, depth113 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l114
						}
						position++
						goto l113
					l114:
						position, tokenIndex, depth = position113, tokenIndex113, depth113
						if buffer[position] != rune('E') {
							goto l104
						}
						position++
					}
				l113:
					{
						position115, tokenIndex115, depth115 := position, tokenIndex, depth
						if buffer[position] != rune('c') {
							goto l116
						}
						position++
						goto l115
					l116:
						position, tokenIndex, depth = position115, tokenIndex115, depth115
						if buffer[position] != rune('C') {
							goto l104
						}
						position++
					}
				l115:
					{
						position117, tokenIndex117, depth117 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l118
						}
						position++
						goto l117
					l118:
						position, tokenIndex, depth = position117, tokenIndex117, depth117
						if buffer[position] != rune('T') {
							goto l104
						}
						position++
					}
				l117:
					if !rules[ruleskip]() {
						goto l104
					}
					depth--
					add(ruleSELECT, position106)
				}
				{
					position119, tokenIndex119, depth119 := position, tokenIndex, depth
					{
						position121, tokenIndex121, depth121 := position, tokenIndex, depth
						if !rules[ruleDISTINCT]() {
							goto l122
						}
						goto l121
					l122:
						position, tokenIndex, depth = position121, tokenIndex121, depth121
						{
							position123 := position
							depth++
							{
								position124, tokenIndex124, depth124 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l125
								}
								position++
								goto l124
							l125:
								position, tokenIndex, depth = position124, tokenIndex124, depth124
								if buffer[position] != rune('R') {
									goto l119
								}
								position++
							}
						l124:
							{
								position126, tokenIndex126, depth126 := position, tokenIndex, depth
								if buffer[position] != rune('e') {
									goto l127
								}
								position++
								goto l126
							l127:
								position, tokenIndex, depth = position126, tokenIndex126, depth126
								if buffer[position] != rune('E') {
									goto l119
								}
								position++
							}
						l126:
							{
								position128, tokenIndex128, depth128 := position, tokenIndex, depth
								if buffer[position] != rune('d') {
									goto l129
								}
								position++
								goto l128
							l129:
								position, tokenIndex, depth = position128, tokenIndex128, depth128
								if buffer[position] != rune('D') {
									goto l119
								}
								position++
							}
						l128:
							{
								position130, tokenIndex130, depth130 := position, tokenIndex, depth
								if buffer[position] != rune('u') {
									goto l131
								}
								position++
								goto l130
							l131:
								position, tokenIndex, depth = position130, tokenIndex130, depth130
								if buffer[position] != rune('U') {
									goto l119
								}
								position++
							}
						l130:
							{
								position132, tokenIndex132, depth132 := position, tokenIndex, depth
								if buffer[position] != rune('c') {
									goto l133
								}
								position++
								goto l132
							l133:
								position, tokenIndex, depth = position132, tokenIndex132, depth132
								if buffer[position] != rune('C') {
									goto l119
								}
								position++
							}
						l132:
							{
								position134, tokenIndex134, depth134 := position, tokenIndex, depth
								if buffer[position] != rune('e') {
									goto l135
								}
								position++
								goto l134
							l135:
								position, tokenIndex, depth = position134, tokenIndex134, depth134
								if buffer[position] != rune('E') {
									goto l119
								}
								position++
							}
						l134:
							{
								position136, tokenIndex136, depth136 := position, tokenIndex, depth
								if buffer[position] != rune('d') {
									goto l137
								}
								position++
								goto l136
							l137:
								position, tokenIndex, depth = position136, tokenIndex136, depth136
								if buffer[position] != rune('D') {
									goto l119
								}
								position++
							}
						l136:
							if !rules[ruleskip]() {
								goto l119
							}
							depth--
							add(ruleREDUCED, position123)
						}
					}
				l121:
					goto l120
				l119:
					position, tokenIndex, depth = position119, tokenIndex119, depth119
				}
			l120:
				{
					position138, tokenIndex138, depth138 := position, tokenIndex, depth
					if !rules[ruleSTAR]() {
						goto l139
					}
					goto l138
				l139:
					position, tokenIndex, depth = position138, tokenIndex138, depth138
					{
						position142 := position
						depth++
						{
							position143, tokenIndex143, depth143 := position, tokenIndex, depth
							if !rules[rulevar]() {
								goto l144
							}
							goto l143
						l144:
							position, tokenIndex, depth = position143, tokenIndex143, depth143
							if !rules[ruleLPAREN]() {
								goto l104
							}
							if !rules[ruleexpression]() {
								goto l104
							}
							if !rules[ruleAS]() {
								goto l104
							}
							if !rules[rulevar]() {
								goto l104
							}
							if !rules[ruleRPAREN]() {
								goto l104
							}
						}
					l143:
						depth--
						add(ruleprojectionElem, position142)
					}
				l140:
					{
						position141, tokenIndex141, depth141 := position, tokenIndex, depth
						{
							position145 := position
							depth++
							{
								position146, tokenIndex146, depth146 := position, tokenIndex, depth
								if !rules[rulevar]() {
									goto l147
								}
								goto l146
							l147:
								position, tokenIndex, depth = position146, tokenIndex146, depth146
								if !rules[ruleLPAREN]() {
									goto l141
								}
								if !rules[ruleexpression]() {
									goto l141
								}
								if !rules[ruleAS]() {
									goto l141
								}
								if !rules[rulevar]() {
									goto l141
								}
								if !rules[ruleRPAREN]() {
									goto l141
								}
							}
						l146:
							depth--
							add(ruleprojectionElem, position145)
						}
						goto l140
					l141:
						position, tokenIndex, depth = position141, tokenIndex141, depth141
					}
				}
			l138:
				depth--
				add(ruleselect, position105)
			}
			return true
		l104:
			position, tokenIndex, depth = position104, tokenIndex104, depth104
			return false
		},
		/* 7 subSelect <- <(select whereClause solutionModifier)> */
		func() bool {
			position148, tokenIndex148, depth148 := position, tokenIndex, depth
			{
				position149 := position
				depth++
				if !rules[ruleselect]() {
					goto l148
				}
				if !rules[rulewhereClause]() {
					goto l148
				}
				if !rules[rulesolutionModifier]() {
					goto l148
				}
				depth--
				add(rulesubSelect, position149)
			}
			return true
		l148:
			position, tokenIndex, depth = position148, tokenIndex148, depth148
			return false
		},
		/* 8 constructQuery <- <(construct datasetClause* whereClause solutionModifier)> */
		nil,
		/* 9 construct <- <(CONSTRUCT LBRACE triplesBlock? RBRACE)> */
		nil,
		/* 10 describeQuery <- <(describe datasetClause* whereClause? solutionModifier)> */
		nil,
		/* 11 describe <- <(DESCRIBE ((&('*') STAR) | (&('$' | '?') var) | (&(':' | '<' | 'A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' | 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z' | 'a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'g' | 'h' | 'i' | 'j' | 'k' | 'l' | 'm' | 'n' | 'o' | 'p' | 'q' | 'r' | 's' | 't' | 'u' | 'v' | 'w' | 'x' | 'y' | 'z' | 'Ã' | 'Ä' | 'Å' | 'Æ' | 'Ç' | 'È' | 'É' | 'Ê' | 'Ë' | 'Í' | 'Î' | 'Ï' | 'Ð' | 'Ñ' | 'Ò' | 'Ó' | 'Ô' | 'Õ' | 'Ö' | '×' | 'Ø' | 'Ù' | 'Ú' | 'Û' | 'Ü' | 'Ý' | 'Þ' | 'ß' | 'à' | 'á' | 'â' | 'ã' | 'ä' | 'å' | 'æ' | 'ç' | 'è' | 'é' | 'ê' | 'ë' | 'ì' | 'í' | 'ï' | 'ð' | 'ñ' | 'ò' | 'ó') iriref)))> */
		nil,
		/* 12 askQuery <- <(ASK datasetClause* whereClause)> */
		nil,
		/* 13 projectionElem <- <(var / (LPAREN expression AS var RPAREN))> */
		nil,
		/* 14 datasetClause <- <(FROM NAMED? iriref)> */
		func() bool {
			position156, tokenIndex156, depth156 := position, tokenIndex, depth
			{
				position157 := position
				depth++
				{
					position158 := position
					depth++
					{
						position159, tokenIndex159, depth159 := position, tokenIndex, depth
						if buffer[position] != rune('f') {
							goto l160
						}
						position++
						goto l159
					l160:
						position, tokenIndex, depth = position159, tokenIndex159, depth159
						if buffer[position] != rune('F') {
							goto l156
						}
						position++
					}
				l159:
					{
						position161, tokenIndex161, depth161 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l162
						}
						position++
						goto l161
					l162:
						position, tokenIndex, depth = position161, tokenIndex161, depth161
						if buffer[position] != rune('R') {
							goto l156
						}
						position++
					}
				l161:
					{
						position163, tokenIndex163, depth163 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l164
						}
						position++
						goto l163
					l164:
						position, tokenIndex, depth = position163, tokenIndex163, depth163
						if buffer[position] != rune('O') {
							goto l156
						}
						position++
					}
				l163:
					{
						position165, tokenIndex165, depth165 := position, tokenIndex, depth
						if buffer[position] != rune('m') {
							goto l166
						}
						position++
						goto l165
					l166:
						position, tokenIndex, depth = position165, tokenIndex165, depth165
						if buffer[position] != rune('M') {
							goto l156
						}
						position++
					}
				l165:
					if !rules[ruleskip]() {
						goto l156
					}
					depth--
					add(ruleFROM, position158)
				}
				{
					position167, tokenIndex167, depth167 := position, tokenIndex, depth
					{
						position169 := position
						depth++
						{
							position170, tokenIndex170, depth170 := position, tokenIndex, depth
							if buffer[position] != rune('n') {
								goto l171
							}
							position++
							goto l170
						l171:
							position, tokenIndex, depth = position170, tokenIndex170, depth170
							if buffer[position] != rune('N') {
								goto l167
							}
							position++
						}
					l170:
						{
							position172, tokenIndex172, depth172 := position, tokenIndex, depth
							if buffer[position] != rune('a') {
								goto l173
							}
							position++
							goto l172
						l173:
							position, tokenIndex, depth = position172, tokenIndex172, depth172
							if buffer[position] != rune('A') {
								goto l167
							}
							position++
						}
					l172:
						{
							position174, tokenIndex174, depth174 := position, tokenIndex, depth
							if buffer[position] != rune('m') {
								goto l175
							}
							position++
							goto l174
						l175:
							position, tokenIndex, depth = position174, tokenIndex174, depth174
							if buffer[position] != rune('M') {
								goto l167
							}
							position++
						}
					l174:
						{
							position176, tokenIndex176, depth176 := position, tokenIndex, depth
							if buffer[position] != rune('e') {
								goto l177
							}
							position++
							goto l176
						l177:
							position, tokenIndex, depth = position176, tokenIndex176, depth176
							if buffer[position] != rune('E') {
								goto l167
							}
							position++
						}
					l176:
						{
							position178, tokenIndex178, depth178 := position, tokenIndex, depth
							if buffer[position] != rune('d') {
								goto l179
							}
							position++
							goto l178
						l179:
							position, tokenIndex, depth = position178, tokenIndex178, depth178
							if buffer[position] != rune('D') {
								goto l167
							}
							position++
						}
					l178:
						if !rules[ruleskip]() {
							goto l167
						}
						depth--
						add(ruleNAMED, position169)
					}
					goto l168
				l167:
					position, tokenIndex, depth = position167, tokenIndex167, depth167
				}
			l168:
				if !rules[ruleiriref]() {
					goto l156
				}
				depth--
				add(ruledatasetClause, position157)
			}
			return true
		l156:
			position, tokenIndex, depth = position156, tokenIndex156, depth156
			return false
		},
		/* 15 whereClause <- <(WHERE? groupGraphPattern)> */
		func() bool {
			position180, tokenIndex180, depth180 := position, tokenIndex, depth
			{
				position181 := position
				depth++
				{
					position182, tokenIndex182, depth182 := position, tokenIndex, depth
					{
						position184 := position
						depth++
						{
							position185, tokenIndex185, depth185 := position, tokenIndex, depth
							if buffer[position] != rune('w') {
								goto l186
							}
							position++
							goto l185
						l186:
							position, tokenIndex, depth = position185, tokenIndex185, depth185
							if buffer[position] != rune('W') {
								goto l182
							}
							position++
						}
					l185:
						{
							position187, tokenIndex187, depth187 := position, tokenIndex, depth
							if buffer[position] != rune('h') {
								goto l188
							}
							position++
							goto l187
						l188:
							position, tokenIndex, depth = position187, tokenIndex187, depth187
							if buffer[position] != rune('H') {
								goto l182
							}
							position++
						}
					l187:
						{
							position189, tokenIndex189, depth189 := position, tokenIndex, depth
							if buffer[position] != rune('e') {
								goto l190
							}
							position++
							goto l189
						l190:
							position, tokenIndex, depth = position189, tokenIndex189, depth189
							if buffer[position] != rune('E') {
								goto l182
							}
							position++
						}
					l189:
						{
							position191, tokenIndex191, depth191 := position, tokenIndex, depth
							if buffer[position] != rune('r') {
								goto l192
							}
							position++
							goto l191
						l192:
							position, tokenIndex, depth = position191, tokenIndex191, depth191
							if buffer[position] != rune('R') {
								goto l182
							}
							position++
						}
					l191:
						{
							position193, tokenIndex193, depth193 := position, tokenIndex, depth
							if buffer[position] != rune('e') {
								goto l194
							}
							position++
							goto l193
						l194:
							position, tokenIndex, depth = position193, tokenIndex193, depth193
							if buffer[position] != rune('E') {
								goto l182
							}
							position++
						}
					l193:
						if !rules[ruleskip]() {
							goto l182
						}
						depth--
						add(ruleWHERE, position184)
					}
					goto l183
				l182:
					position, tokenIndex, depth = position182, tokenIndex182, depth182
				}
			l183:
				if !rules[rulegroupGraphPattern]() {
					goto l180
				}
				depth--
				add(rulewhereClause, position181)
			}
			return true
		l180:
			position, tokenIndex, depth = position180, tokenIndex180, depth180
			return false
		},
		/* 16 groupGraphPattern <- <(LBRACE (subSelect / graphPattern) RBRACE)> */
		func() bool {
			position195, tokenIndex195, depth195 := position, tokenIndex, depth
			{
				position196 := position
				depth++
				if !rules[ruleLBRACE]() {
					goto l195
				}
				{
					position197, tokenIndex197, depth197 := position, tokenIndex, depth
					if !rules[rulesubSelect]() {
						goto l198
					}
					goto l197
				l198:
					position, tokenIndex, depth = position197, tokenIndex197, depth197
					if !rules[rulegraphPattern]() {
						goto l195
					}
				}
			l197:
				if !rules[ruleRBRACE]() {
					goto l195
				}
				depth--
				add(rulegroupGraphPattern, position196)
			}
			return true
		l195:
			position, tokenIndex, depth = position195, tokenIndex195, depth195
			return false
		},
		/* 17 graphPattern <- <(basicGraphPattern? (graphPatternNotTriples DOT? graphPattern)?)> */
		func() bool {
			{
				position200 := position
				depth++
				{
					position201, tokenIndex201, depth201 := position, tokenIndex, depth
					{
						position203 := position
						depth++
						{
							position204, tokenIndex204, depth204 := position, tokenIndex, depth
							if !rules[ruletriplesBlock]() {
								goto l205
							}
						l206:
							{
								position207, tokenIndex207, depth207 := position, tokenIndex, depth
								if !rules[rulefilterOrBind]() {
									goto l207
								}
								{
									position208, tokenIndex208, depth208 := position, tokenIndex, depth
									if !rules[ruleDOT]() {
										goto l208
									}
									goto l209
								l208:
									position, tokenIndex, depth = position208, tokenIndex208, depth208
								}
							l209:
								{
									position210, tokenIndex210, depth210 := position, tokenIndex, depth
									if !rules[ruletriplesBlock]() {
										goto l210
									}
									goto l211
								l210:
									position, tokenIndex, depth = position210, tokenIndex210, depth210
								}
							l211:
								goto l206
							l207:
								position, tokenIndex, depth = position207, tokenIndex207, depth207
							}
							goto l204
						l205:
							position, tokenIndex, depth = position204, tokenIndex204, depth204
							if !rules[rulefilterOrBind]() {
								goto l201
							}
							{
								position214, tokenIndex214, depth214 := position, tokenIndex, depth
								if !rules[ruleDOT]() {
									goto l214
								}
								goto l215
							l214:
								position, tokenIndex, depth = position214, tokenIndex214, depth214
							}
						l215:
							{
								position216, tokenIndex216, depth216 := position, tokenIndex, depth
								if !rules[ruletriplesBlock]() {
									goto l216
								}
								goto l217
							l216:
								position, tokenIndex, depth = position216, tokenIndex216, depth216
							}
						l217:
						l212:
							{
								position213, tokenIndex213, depth213 := position, tokenIndex, depth
								if !rules[rulefilterOrBind]() {
									goto l213
								}
								{
									position218, tokenIndex218, depth218 := position, tokenIndex, depth
									if !rules[ruleDOT]() {
										goto l218
									}
									goto l219
								l218:
									position, tokenIndex, depth = position218, tokenIndex218, depth218
								}
							l219:
								{
									position220, tokenIndex220, depth220 := position, tokenIndex, depth
									if !rules[ruletriplesBlock]() {
										goto l220
									}
									goto l221
								l220:
									position, tokenIndex, depth = position220, tokenIndex220, depth220
								}
							l221:
								goto l212
							l213:
								position, tokenIndex, depth = position213, tokenIndex213, depth213
							}
						}
					l204:
						depth--
						add(rulebasicGraphPattern, position203)
					}
					goto l202
				l201:
					position, tokenIndex, depth = position201, tokenIndex201, depth201
				}
			l202:
				{
					position222, tokenIndex222, depth222 := position, tokenIndex, depth
					{
						position224 := position
						depth++
						{
							switch buffer[position] {
							case 'S', 's':
								{
									position226 := position
									depth++
									{
										position227 := position
										depth++
										{
											position228, tokenIndex228, depth228 := position, tokenIndex, depth
											if buffer[position] != rune('s') {
												goto l229
											}
											position++
											goto l228
										l229:
											position, tokenIndex, depth = position228, tokenIndex228, depth228
											if buffer[position] != rune('S') {
												goto l222
											}
											position++
										}
									l228:
										{
											position230, tokenIndex230, depth230 := position, tokenIndex, depth
											if buffer[position] != rune('e') {
												goto l231
											}
											position++
											goto l230
										l231:
											position, tokenIndex, depth = position230, tokenIndex230, depth230
											if buffer[position] != rune('E') {
												goto l222
											}
											position++
										}
									l230:
										{
											position232, tokenIndex232, depth232 := position, tokenIndex, depth
											if buffer[position] != rune('r') {
												goto l233
											}
											position++
											goto l232
										l233:
											position, tokenIndex, depth = position232, tokenIndex232, depth232
											if buffer[position] != rune('R') {
												goto l222
											}
											position++
										}
									l232:
										{
											position234, tokenIndex234, depth234 := position, tokenIndex, depth
											if buffer[position] != rune('v') {
												goto l235
											}
											position++
											goto l234
										l235:
											position, tokenIndex, depth = position234, tokenIndex234, depth234
											if buffer[position] != rune('V') {
												goto l222
											}
											position++
										}
									l234:
										{
											position236, tokenIndex236, depth236 := position, tokenIndex, depth
											if buffer[position] != rune('i') {
												goto l237
											}
											position++
											goto l236
										l237:
											position, tokenIndex, depth = position236, tokenIndex236, depth236
											if buffer[position] != rune('I') {
												goto l222
											}
											position++
										}
									l236:
										{
											position238, tokenIndex238, depth238 := position, tokenIndex, depth
											if buffer[position] != rune('c') {
												goto l239
											}
											position++
											goto l238
										l239:
											position, tokenIndex, depth = position238, tokenIndex238, depth238
											if buffer[position] != rune('C') {
												goto l222
											}
											position++
										}
									l238:
										{
											position240, tokenIndex240, depth240 := position, tokenIndex, depth
											if buffer[position] != rune('e') {
												goto l241
											}
											position++
											goto l240
										l241:
											position, tokenIndex, depth = position240, tokenIndex240, depth240
											if buffer[position] != rune('E') {
												goto l222
											}
											position++
										}
									l240:
										if !rules[ruleskip]() {
											goto l222
										}
										depth--
										add(ruleSERVICE, position227)
									}
									{
										position242, tokenIndex242, depth242 := position, tokenIndex, depth
										{
											position244 := position
											depth++
											{
												position245, tokenIndex245, depth245 := position, tokenIndex, depth
												if buffer[position] != rune('s') {
													goto l246
												}
												position++
												goto l245
											l246:
												position, tokenIndex, depth = position245, tokenIndex245, depth245
												if buffer[position] != rune('S') {
													goto l242
												}
												position++
											}
										l245:
											{
												position247, tokenIndex247, depth247 := position, tokenIndex, depth
												if buffer[position] != rune('i') {
													goto l248
												}
												position++
												goto l247
											l248:
												position, tokenIndex, depth = position247, tokenIndex247, depth247
												if buffer[position] != rune('I') {
													goto l242
												}
												position++
											}
										l247:
											{
												position249, tokenIndex249, depth249 := position, tokenIndex, depth
												if buffer[position] != rune('l') {
													goto l250
												}
												position++
												goto l249
											l250:
												position, tokenIndex, depth = position249, tokenIndex249, depth249
												if buffer[position] != rune('L') {
													goto l242
												}
												position++
											}
										l249:
											{
												position251, tokenIndex251, depth251 := position, tokenIndex, depth
												if buffer[position] != rune('e') {
													goto l252
												}
												position++
												goto l251
											l252:
												position, tokenIndex, depth = position251, tokenIndex251, depth251
												if buffer[position] != rune('E') {
													goto l242
												}
												position++
											}
										l251:
											{
												position253, tokenIndex253, depth253 := position, tokenIndex, depth
												if buffer[position] != rune('n') {
													goto l254
												}
												position++
												goto l253
											l254:
												position, tokenIndex, depth = position253, tokenIndex253, depth253
												if buffer[position] != rune('N') {
													goto l242
												}
												position++
											}
										l253:
											{
												position255, tokenIndex255, depth255 := position, tokenIndex, depth
												if buffer[position] != rune('t') {
													goto l256
												}
												position++
												goto l255
											l256:
												position, tokenIndex, depth = position255, tokenIndex255, depth255
												if buffer[position] != rune('T') {
													goto l242
												}
												position++
											}
										l255:
											if !rules[ruleskip]() {
												goto l242
											}
											depth--
											add(ruleSILENT, position244)
										}
										goto l243
									l242:
										position, tokenIndex, depth = position242, tokenIndex242, depth242
									}
								l243:
									{
										position257, tokenIndex257, depth257 := position, tokenIndex, depth
										if !rules[rulevar]() {
											goto l258
										}
										goto l257
									l258:
										position, tokenIndex, depth = position257, tokenIndex257, depth257
										if !rules[ruleiriref]() {
											goto l222
										}
									}
								l257:
									if !rules[rulegroupGraphPattern]() {
										goto l222
									}
									depth--
									add(ruleserviceGraphPattern, position226)
								}
								break
							case 'M', 'm':
								{
									position259 := position
									depth++
									{
										position260 := position
										depth++
										{
											position261, tokenIndex261, depth261 := position, tokenIndex, depth
											if buffer[position] != rune('m') {
												goto l262
											}
											position++
											goto l261
										l262:
											position, tokenIndex, depth = position261, tokenIndex261, depth261
											if buffer[position] != rune('M') {
												goto l222
											}
											position++
										}
									l261:
										{
											position263, tokenIndex263, depth263 := position, tokenIndex, depth
											if buffer[position] != rune('i') {
												goto l264
											}
											position++
											goto l263
										l264:
											position, tokenIndex, depth = position263, tokenIndex263, depth263
											if buffer[position] != rune('I') {
												goto l222
											}
											position++
										}
									l263:
										{
											position265, tokenIndex265, depth265 := position, tokenIndex, depth
											if buffer[position] != rune('n') {
												goto l266
											}
											position++
											goto l265
										l266:
											position, tokenIndex, depth = position265, tokenIndex265, depth265
											if buffer[position] != rune('N') {
												goto l222
											}
											position++
										}
									l265:
										{
											position267, tokenIndex267, depth267 := position, tokenIndex, depth
											if buffer[position] != rune('u') {
												goto l268
											}
											position++
											goto l267
										l268:
											position, tokenIndex, depth = position267, tokenIndex267, depth267
											if buffer[position] != rune('U') {
												goto l222
											}
											position++
										}
									l267:
										{
											position269, tokenIndex269, depth269 := position, tokenIndex, depth
											if buffer[position] != rune('s') {
												goto l270
											}
											position++
											goto l269
										l270:
											position, tokenIndex, depth = position269, tokenIndex269, depth269
											if buffer[position] != rune('S') {
												goto l222
											}
											position++
										}
									l269:
										if !rules[ruleskip]() {
											goto l222
										}
										depth--
										add(ruleMINUSSETOPER, position260)
									}
									if !rules[rulegroupGraphPattern]() {
										goto l222
									}
									depth--
									add(ruleminusGraphPattern, position259)
								}
								break
							case 'G', 'g':
								{
									position271 := position
									depth++
									{
										position272 := position
										depth++
										{
											position273, tokenIndex273, depth273 := position, tokenIndex, depth
											if buffer[position] != rune('g') {
												goto l274
											}
											position++
											goto l273
										l274:
											position, tokenIndex, depth = position273, tokenIndex273, depth273
											if buffer[position] != rune('G') {
												goto l222
											}
											position++
										}
									l273:
										{
											position275, tokenIndex275, depth275 := position, tokenIndex, depth
											if buffer[position] != rune('r') {
												goto l276
											}
											position++
											goto l275
										l276:
											position, tokenIndex, depth = position275, tokenIndex275, depth275
											if buffer[position] != rune('R') {
												goto l222
											}
											position++
										}
									l275:
										{
											position277, tokenIndex277, depth277 := position, tokenIndex, depth
											if buffer[position] != rune('a') {
												goto l278
											}
											position++
											goto l277
										l278:
											position, tokenIndex, depth = position277, tokenIndex277, depth277
											if buffer[position] != rune('A') {
												goto l222
											}
											position++
										}
									l277:
										{
											position279, tokenIndex279, depth279 := position, tokenIndex, depth
											if buffer[position] != rune('p') {
												goto l280
											}
											position++
											goto l279
										l280:
											position, tokenIndex, depth = position279, tokenIndex279, depth279
											if buffer[position] != rune('P') {
												goto l222
											}
											position++
										}
									l279:
										{
											position281, tokenIndex281, depth281 := position, tokenIndex, depth
											if buffer[position] != rune('h') {
												goto l282
											}
											position++
											goto l281
										l282:
											position, tokenIndex, depth = position281, tokenIndex281, depth281
											if buffer[position] != rune('H') {
												goto l222
											}
											position++
										}
									l281:
										if !rules[ruleskip]() {
											goto l222
										}
										depth--
										add(ruleGRAPH, position272)
									}
									{
										position283, tokenIndex283, depth283 := position, tokenIndex, depth
										if !rules[rulevar]() {
											goto l284
										}
										goto l283
									l284:
										position, tokenIndex, depth = position283, tokenIndex283, depth283
										if !rules[ruleiriref]() {
											goto l222
										}
									}
								l283:
									if !rules[rulegroupGraphPattern]() {
										goto l222
									}
									depth--
									add(rulegraphGraphPattern, position271)
								}
								break
							case '{':
								if !rules[rulegroupOrUnionGraphPattern]() {
									goto l222
								}
								break
							default:
								{
									position285 := position
									depth++
									{
										position286 := position
										depth++
										{
											position287, tokenIndex287, depth287 := position, tokenIndex, depth
											if buffer[position] != rune('o') {
												goto l288
											}
											position++
											goto l287
										l288:
											position, tokenIndex, depth = position287, tokenIndex287, depth287
											if buffer[position] != rune('O') {
												goto l222
											}
											position++
										}
									l287:
										{
											position289, tokenIndex289, depth289 := position, tokenIndex, depth
											if buffer[position] != rune('p') {
												goto l290
											}
											position++
											goto l289
										l290:
											position, tokenIndex, depth = position289, tokenIndex289, depth289
											if buffer[position] != rune('P') {
												goto l222
											}
											position++
										}
									l289:
										{
											position291, tokenIndex291, depth291 := position, tokenIndex, depth
											if buffer[position] != rune('t') {
												goto l292
											}
											position++
											goto l291
										l292:
											position, tokenIndex, depth = position291, tokenIndex291, depth291
											if buffer[position] != rune('T') {
												goto l222
											}
											position++
										}
									l291:
										{
											position293, tokenIndex293, depth293 := position, tokenIndex, depth
											if buffer[position] != rune('i') {
												goto l294
											}
											position++
											goto l293
										l294:
											position, tokenIndex, depth = position293, tokenIndex293, depth293
											if buffer[position] != rune('I') {
												goto l222
											}
											position++
										}
									l293:
										{
											position295, tokenIndex295, depth295 := position, tokenIndex, depth
											if buffer[position] != rune('o') {
												goto l296
											}
											position++
											goto l295
										l296:
											position, tokenIndex, depth = position295, tokenIndex295, depth295
											if buffer[position] != rune('O') {
												goto l222
											}
											position++
										}
									l295:
										{
											position297, tokenIndex297, depth297 := position, tokenIndex, depth
											if buffer[position] != rune('n') {
												goto l298
											}
											position++
											goto l297
										l298:
											position, tokenIndex, depth = position297, tokenIndex297, depth297
											if buffer[position] != rune('N') {
												goto l222
											}
											position++
										}
									l297:
										{
											position299, tokenIndex299, depth299 := position, tokenIndex, depth
											if buffer[position] != rune('a') {
												goto l300
											}
											position++
											goto l299
										l300:
											position, tokenIndex, depth = position299, tokenIndex299, depth299
											if buffer[position] != rune('A') {
												goto l222
											}
											position++
										}
									l299:
										{
											position301, tokenIndex301, depth301 := position, tokenIndex, depth
											if buffer[position] != rune('l') {
												goto l302
											}
											position++
											goto l301
										l302:
											position, tokenIndex, depth = position301, tokenIndex301, depth301
											if buffer[position] != rune('L') {
												goto l222
											}
											position++
										}
									l301:
										if !rules[ruleskip]() {
											goto l222
										}
										depth--
										add(ruleOPTIONAL, position286)
									}
									if !rules[ruleLBRACE]() {
										goto l222
									}
									{
										position303, tokenIndex303, depth303 := position, tokenIndex, depth
										if !rules[rulesubSelect]() {
											goto l304
										}
										goto l303
									l304:
										position, tokenIndex, depth = position303, tokenIndex303, depth303
										if !rules[rulegraphPattern]() {
											goto l222
										}
									}
								l303:
									if !rules[ruleRBRACE]() {
										goto l222
									}
									depth--
									add(ruleoptionalGraphPattern, position285)
								}
								break
							}
						}

						depth--
						add(rulegraphPatternNotTriples, position224)
					}
					{
						position305, tokenIndex305, depth305 := position, tokenIndex, depth
						if !rules[ruleDOT]() {
							goto l305
						}
						goto l306
					l305:
						position, tokenIndex, depth = position305, tokenIndex305, depth305
					}
				l306:
					if !rules[rulegraphPattern]() {
						goto l222
					}
					goto l223
				l222:
					position, tokenIndex, depth = position222, tokenIndex222, depth222
				}
			l223:
				depth--
				add(rulegraphPattern, position200)
			}
			return true
		},
		/* 18 graphPatternNotTriples <- <((&('S' | 's') serviceGraphPattern) | (&('M' | 'm') minusGraphPattern) | (&('G' | 'g') graphGraphPattern) | (&('{') groupOrUnionGraphPattern) | (&('O' | 'o') optionalGraphPattern))> */
		nil,
		/* 19 serviceGraphPattern <- <(SERVICE SILENT? (var / iriref) groupGraphPattern)> */
		nil,
		/* 20 optionalGraphPattern <- <(OPTIONAL LBRACE (subSelect / graphPattern) RBRACE)> */
		nil,
		/* 21 groupOrUnionGraphPattern <- <(groupGraphPattern (UNION groupOrUnionGraphPattern)?)> */
		func() bool {
			position310, tokenIndex310, depth310 := position, tokenIndex, depth
			{
				position311 := position
				depth++
				if !rules[rulegroupGraphPattern]() {
					goto l310
				}
				{
					position312, tokenIndex312, depth312 := position, tokenIndex, depth
					{
						position314 := position
						depth++
						{
							position315, tokenIndex315, depth315 := position, tokenIndex, depth
							if buffer[position] != rune('u') {
								goto l316
							}
							position++
							goto l315
						l316:
							position, tokenIndex, depth = position315, tokenIndex315, depth315
							if buffer[position] != rune('U') {
								goto l312
							}
							position++
						}
					l315:
						{
							position317, tokenIndex317, depth317 := position, tokenIndex, depth
							if buffer[position] != rune('n') {
								goto l318
							}
							position++
							goto l317
						l318:
							position, tokenIndex, depth = position317, tokenIndex317, depth317
							if buffer[position] != rune('N') {
								goto l312
							}
							position++
						}
					l317:
						{
							position319, tokenIndex319, depth319 := position, tokenIndex, depth
							if buffer[position] != rune('i') {
								goto l320
							}
							position++
							goto l319
						l320:
							position, tokenIndex, depth = position319, tokenIndex319, depth319
							if buffer[position] != rune('I') {
								goto l312
							}
							position++
						}
					l319:
						{
							position321, tokenIndex321, depth321 := position, tokenIndex, depth
							if buffer[position] != rune('o') {
								goto l322
							}
							position++
							goto l321
						l322:
							position, tokenIndex, depth = position321, tokenIndex321, depth321
							if buffer[position] != rune('O') {
								goto l312
							}
							position++
						}
					l321:
						{
							position323, tokenIndex323, depth323 := position, tokenIndex, depth
							if buffer[position] != rune('n') {
								goto l324
							}
							position++
							goto l323
						l324:
							position, tokenIndex, depth = position323, tokenIndex323, depth323
							if buffer[position] != rune('N') {
								goto l312
							}
							position++
						}
					l323:
						if !rules[ruleskip]() {
							goto l312
						}
						depth--
						add(ruleUNION, position314)
					}
					if !rules[rulegroupOrUnionGraphPattern]() {
						goto l312
					}
					goto l313
				l312:
					position, tokenIndex, depth = position312, tokenIndex312, depth312
				}
			l313:
				depth--
				add(rulegroupOrUnionGraphPattern, position311)
			}
			return true
		l310:
			position, tokenIndex, depth = position310, tokenIndex310, depth310
			return false
		},
		/* 22 graphGraphPattern <- <(GRAPH (var / iriref) groupGraphPattern)> */
		nil,
		/* 23 minusGraphPattern <- <(MINUSSETOPER groupGraphPattern)> */
		nil,
		/* 24 basicGraphPattern <- <((triplesBlock (filterOrBind DOT? triplesBlock?)*) / (filterOrBind DOT? triplesBlock?)+)> */
		nil,
		/* 25 filterOrBind <- <((FILTER constraint) / (BIND LPAREN expression AS var RPAREN))> */
		func() bool {
			position328, tokenIndex328, depth328 := position, tokenIndex, depth
			{
				position329 := position
				depth++
				{
					position330, tokenIndex330, depth330 := position, tokenIndex, depth
					{
						position332 := position
						depth++
						{
							position333, tokenIndex333, depth333 := position, tokenIndex, depth
							if buffer[position] != rune('f') {
								goto l334
							}
							position++
							goto l333
						l334:
							position, tokenIndex, depth = position333, tokenIndex333, depth333
							if buffer[position] != rune('F') {
								goto l331
							}
							position++
						}
					l333:
						{
							position335, tokenIndex335, depth335 := position, tokenIndex, depth
							if buffer[position] != rune('i') {
								goto l336
							}
							position++
							goto l335
						l336:
							position, tokenIndex, depth = position335, tokenIndex335, depth335
							if buffer[position] != rune('I') {
								goto l331
							}
							position++
						}
					l335:
						{
							position337, tokenIndex337, depth337 := position, tokenIndex, depth
							if buffer[position] != rune('l') {
								goto l338
							}
							position++
							goto l337
						l338:
							position, tokenIndex, depth = position337, tokenIndex337, depth337
							if buffer[position] != rune('L') {
								goto l331
							}
							position++
						}
					l337:
						{
							position339, tokenIndex339, depth339 := position, tokenIndex, depth
							if buffer[position] != rune('t') {
								goto l340
							}
							position++
							goto l339
						l340:
							position, tokenIndex, depth = position339, tokenIndex339, depth339
							if buffer[position] != rune('T') {
								goto l331
							}
							position++
						}
					l339:
						{
							position341, tokenIndex341, depth341 := position, tokenIndex, depth
							if buffer[position] != rune('e') {
								goto l342
							}
							position++
							goto l341
						l342:
							position, tokenIndex, depth = position341, tokenIndex341, depth341
							if buffer[position] != rune('E') {
								goto l331
							}
							position++
						}
					l341:
						{
							position343, tokenIndex343, depth343 := position, tokenIndex, depth
							if buffer[position] != rune('r') {
								goto l344
							}
							position++
							goto l343
						l344:
							position, tokenIndex, depth = position343, tokenIndex343, depth343
							if buffer[position] != rune('R') {
								goto l331
							}
							position++
						}
					l343:
						if !rules[ruleskip]() {
							goto l331
						}
						depth--
						add(ruleFILTER, position332)
					}
					if !rules[ruleconstraint]() {
						goto l331
					}
					goto l330
				l331:
					position, tokenIndex, depth = position330, tokenIndex330, depth330
					{
						position345 := position
						depth++
						{
							position346, tokenIndex346, depth346 := position, tokenIndex, depth
							if buffer[position] != rune('b') {
								goto l347
							}
							position++
							goto l346
						l347:
							position, tokenIndex, depth = position346, tokenIndex346, depth346
							if buffer[position] != rune('B') {
								goto l328
							}
							position++
						}
					l346:
						{
							position348, tokenIndex348, depth348 := position, tokenIndex, depth
							if buffer[position] != rune('i') {
								goto l349
							}
							position++
							goto l348
						l349:
							position, tokenIndex, depth = position348, tokenIndex348, depth348
							if buffer[position] != rune('I') {
								goto l328
							}
							position++
						}
					l348:
						{
							position350, tokenIndex350, depth350 := position, tokenIndex, depth
							if buffer[position] != rune('n') {
								goto l351
							}
							position++
							goto l350
						l351:
							position, tokenIndex, depth = position350, tokenIndex350, depth350
							if buffer[position] != rune('N') {
								goto l328
							}
							position++
						}
					l350:
						{
							position352, tokenIndex352, depth352 := position, tokenIndex, depth
							if buffer[position] != rune('d') {
								goto l353
							}
							position++
							goto l352
						l353:
							position, tokenIndex, depth = position352, tokenIndex352, depth352
							if buffer[position] != rune('D') {
								goto l328
							}
							position++
						}
					l352:
						if !rules[ruleskip]() {
							goto l328
						}
						depth--
						add(ruleBIND, position345)
					}
					if !rules[ruleLPAREN]() {
						goto l328
					}
					if !rules[ruleexpression]() {
						goto l328
					}
					if !rules[ruleAS]() {
						goto l328
					}
					if !rules[rulevar]() {
						goto l328
					}
					if !rules[ruleRPAREN]() {
						goto l328
					}
				}
			l330:
				depth--
				add(rulefilterOrBind, position329)
			}
			return true
		l328:
			position, tokenIndex, depth = position328, tokenIndex328, depth328
			return false
		},
		/* 26 constraint <- <(brackettedExpression / builtinCall / functionCall)> */
		func() bool {
			position354, tokenIndex354, depth354 := position, tokenIndex, depth
			{
				position355 := position
				depth++
				{
					position356, tokenIndex356, depth356 := position, tokenIndex, depth
					if !rules[rulebrackettedExpression]() {
						goto l357
					}
					goto l356
				l357:
					position, tokenIndex, depth = position356, tokenIndex356, depth356
					if !rules[rulebuiltinCall]() {
						goto l358
					}
					goto l356
				l358:
					position, tokenIndex, depth = position356, tokenIndex356, depth356
					if !rules[rulefunctionCall]() {
						goto l354
					}
				}
			l356:
				depth--
				add(ruleconstraint, position355)
			}
			return true
		l354:
			position, tokenIndex, depth = position354, tokenIndex354, depth354
			return false
		},
		/* 27 triplesBlock <- <(triplesSameSubjectPath (DOT triplesSameSubjectPath)* DOT?)> */
		func() bool {
			position359, tokenIndex359, depth359 := position, tokenIndex, depth
			{
				position360 := position
				depth++
				if !rules[ruletriplesSameSubjectPath]() {
					goto l359
				}
			l361:
				{
					position362, tokenIndex362, depth362 := position, tokenIndex, depth
					if !rules[ruleDOT]() {
						goto l362
					}
					if !rules[ruletriplesSameSubjectPath]() {
						goto l362
					}
					goto l361
				l362:
					position, tokenIndex, depth = position362, tokenIndex362, depth362
				}
				{
					position363, tokenIndex363, depth363 := position, tokenIndex, depth
					if !rules[ruleDOT]() {
						goto l363
					}
					goto l364
				l363:
					position, tokenIndex, depth = position363, tokenIndex363, depth363
				}
			l364:
				depth--
				add(ruletriplesBlock, position360)
			}
			return true
		l359:
			position, tokenIndex, depth = position359, tokenIndex359, depth359
			return false
		},
		/* 28 triplesSameSubjectPath <- <((<varOrTerm> Action0 propertyListPath) / (triplesNodePath propertyListPath?))> */
		func() bool {
			position365, tokenIndex365, depth365 := position, tokenIndex, depth
			{
				position366 := position
				depth++
				{
					position367, tokenIndex367, depth367 := position, tokenIndex, depth
					{
						position369 := position
						depth++
						if !rules[rulevarOrTerm]() {
							goto l368
						}
						depth--
						add(rulePegText, position369)
					}
					{
						add(ruleAction0, position)
					}
					if !rules[rulepropertyListPath]() {
						goto l368
					}
					goto l367
				l368:
					position, tokenIndex, depth = position367, tokenIndex367, depth367
					if !rules[ruletriplesNodePath]() {
						goto l365
					}
					{
						position371, tokenIndex371, depth371 := position, tokenIndex, depth
						if !rules[rulepropertyListPath]() {
							goto l371
						}
						goto l372
					l371:
						position, tokenIndex, depth = position371, tokenIndex371, depth371
					}
				l372:
				}
			l367:
				depth--
				add(ruletriplesSameSubjectPath, position366)
			}
			return true
		l365:
			position, tokenIndex, depth = position365, tokenIndex365, depth365
			return false
		},
		/* 29 varOrTerm <- <(var / graphTerm)> */
		func() bool {
			position373, tokenIndex373, depth373 := position, tokenIndex, depth
			{
				position374 := position
				depth++
				{
					position375, tokenIndex375, depth375 := position, tokenIndex, depth
					if !rules[rulevar]() {
						goto l376
					}
					goto l375
				l376:
					position, tokenIndex, depth = position375, tokenIndex375, depth375
					{
						position377 := position
						depth++
						{
							position378, tokenIndex378, depth378 := position, tokenIndex, depth
							if !rules[ruleiriref]() {
								goto l379
							}
							goto l378
						l379:
							position, tokenIndex, depth = position378, tokenIndex378, depth378
							{
								switch buffer[position] {
								case '(':
									if !rules[rulenil]() {
										goto l373
									}
									break
								case '[', '_':
									{
										position381 := position
										depth++
										{
											position382, tokenIndex382, depth382 := position, tokenIndex, depth
											{
												position384 := position
												depth++
												if buffer[position] != rune('_') {
													goto l383
												}
												position++
												if buffer[position] != rune(':') {
													goto l383
												}
												position++
												{
													position385, tokenIndex385, depth385 := position, tokenIndex, depth
													if !rules[rulepnCharsU]() {
														goto l386
													}
													goto l385
												l386:
													position, tokenIndex, depth = position385, tokenIndex385, depth385
													if c := buffer[position]; c < rune('0') || c > rune('9') {
														goto l383
													}
													position++
												}
											l385:
												{
													position387, tokenIndex387, depth387 := position, tokenIndex, depth
													{
														position389, tokenIndex389, depth389 := position, tokenIndex, depth
													l391:
														{
															position392, tokenIndex392, depth392 := position, tokenIndex, depth
															{
																position393, tokenIndex393, depth393 := position, tokenIndex, depth
																if !rules[rulepnCharsU]() {
																	goto l394
																}
																goto l393
															l394:
																position, tokenIndex, depth = position393, tokenIndex393, depth393
																{
																	switch buffer[position] {
																	case '.':
																		if buffer[position] != rune('.') {
																			goto l392
																		}
																		position++
																		break
																	case '-':
																		if buffer[position] != rune('-') {
																			goto l392
																		}
																		position++
																		break
																	default:
																		if c := buffer[position]; c < rune('0') || c > rune('9') {
																			goto l392
																		}
																		position++
																		break
																	}
																}

															}
														l393:
															goto l391
														l392:
															position, tokenIndex, depth = position392, tokenIndex392, depth392
														}
														if !rules[rulepnCharsU]() {
															goto l390
														}
														goto l389
													l390:
														position, tokenIndex, depth = position389, tokenIndex389, depth389
														{
															position396, tokenIndex396, depth396 := position, tokenIndex, depth
															if c := buffer[position]; c < rune('0') || c > rune('9') {
																goto l397
															}
															position++
															goto l396
														l397:
															position, tokenIndex, depth = position396, tokenIndex396, depth396
															if buffer[position] != rune('-') {
																goto l387
															}
															position++
														}
													l396:
													}
												l389:
													goto l388
												l387:
													position, tokenIndex, depth = position387, tokenIndex387, depth387
												}
											l388:
												if !rules[ruleskip]() {
													goto l383
												}
												depth--
												add(ruleblankNodeLabel, position384)
											}
											goto l382
										l383:
											position, tokenIndex, depth = position382, tokenIndex382, depth382
											{
												position398 := position
												depth++
												if buffer[position] != rune('[') {
													goto l373
												}
												position++
											l399:
												{
													position400, tokenIndex400, depth400 := position, tokenIndex, depth
													if !rules[rulews]() {
														goto l400
													}
													goto l399
												l400:
													position, tokenIndex, depth = position400, tokenIndex400, depth400
												}
												if buffer[position] != rune(']') {
													goto l373
												}
												position++
												if !rules[ruleskip]() {
													goto l373
												}
												depth--
												add(ruleanon, position398)
											}
										}
									l382:
										depth--
										add(ruleblankNode, position381)
									}
									break
								case 'F', 'T', 'f', 't':
									if !rules[rulebooleanLiteral]() {
										goto l373
									}
									break
								case '"', '\'':
									if !rules[ruleliteral]() {
										goto l373
									}
									break
								default:
									if !rules[rulenumericLiteral]() {
										goto l373
									}
									break
								}
							}

						}
					l378:
						depth--
						add(rulegraphTerm, position377)
					}
				}
			l375:
				depth--
				add(rulevarOrTerm, position374)
			}
			return true
		l373:
			position, tokenIndex, depth = position373, tokenIndex373, depth373
			return false
		},
		/* 30 graphTerm <- <(iriref / ((&('(') nil) | (&('[' | '_') blankNode) | (&('F' | 'T' | 'f' | 't') booleanLiteral) | (&('"' | '\'') literal) | (&('+' | '-' | '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') numericLiteral)))> */
		nil,
		/* 31 triplesNodePath <- <(collectionPath / blankNodePropertyListPath)> */
		func() bool {
			position402, tokenIndex402, depth402 := position, tokenIndex, depth
			{
				position403 := position
				depth++
				{
					position404, tokenIndex404, depth404 := position, tokenIndex, depth
					{
						position406 := position
						depth++
						if !rules[ruleLPAREN]() {
							goto l405
						}
						if !rules[rulegraphNodePath]() {
							goto l405
						}
					l407:
						{
							position408, tokenIndex408, depth408 := position, tokenIndex, depth
							if !rules[rulegraphNodePath]() {
								goto l408
							}
							goto l407
						l408:
							position, tokenIndex, depth = position408, tokenIndex408, depth408
						}
						if !rules[ruleRPAREN]() {
							goto l405
						}
						depth--
						add(rulecollectionPath, position406)
					}
					goto l404
				l405:
					position, tokenIndex, depth = position404, tokenIndex404, depth404
					{
						position409 := position
						depth++
						{
							position410 := position
							depth++
							if buffer[position] != rune('[') {
								goto l402
							}
							position++
							if !rules[ruleskip]() {
								goto l402
							}
							depth--
							add(ruleLBRACK, position410)
						}
						if !rules[rulepropertyListPath]() {
							goto l402
						}
						{
							position411 := position
							depth++
							if buffer[position] != rune(']') {
								goto l402
							}
							position++
							if !rules[ruleskip]() {
								goto l402
							}
							depth--
							add(ruleRBRACK, position411)
						}
						depth--
						add(ruleblankNodePropertyListPath, position409)
					}
				}
			l404:
				depth--
				add(ruletriplesNodePath, position403)
			}
			return true
		l402:
			position, tokenIndex, depth = position402, tokenIndex402, depth402
			return false
		},
		/* 32 collectionPath <- <(LPAREN graphNodePath+ RPAREN)> */
		nil,
		/* 33 blankNodePropertyListPath <- <(LBRACK propertyListPath RBRACK)> */
		nil,
		/* 34 propertyListPath <- <(<(var / verbPath)> Action1 objectListPath (SEMICOLON propertyListPath?)?)> */
		func() bool {
			position414, tokenIndex414, depth414 := position, tokenIndex, depth
			{
				position415 := position
				depth++
				{
					position416 := position
					depth++
					{
						position417, tokenIndex417, depth417 := position, tokenIndex, depth
						if !rules[rulevar]() {
							goto l418
						}
						goto l417
					l418:
						position, tokenIndex, depth = position417, tokenIndex417, depth417
						{
							position419 := position
							depth++
							if !rules[rulepath]() {
								goto l414
							}
							depth--
							add(ruleverbPath, position419)
						}
					}
				l417:
					depth--
					add(rulePegText, position416)
				}
				{
					add(ruleAction1, position)
				}
				{
					position421 := position
					depth++
					if !rules[ruleobjectPath]() {
						goto l414
					}
				l422:
					{
						position423, tokenIndex423, depth423 := position, tokenIndex, depth
						if !rules[ruleCOMMA]() {
							goto l423
						}
						if !rules[ruleobjectPath]() {
							goto l423
						}
						goto l422
					l423:
						position, tokenIndex, depth = position423, tokenIndex423, depth423
					}
					depth--
					add(ruleobjectListPath, position421)
				}
				{
					position424, tokenIndex424, depth424 := position, tokenIndex, depth
					if !rules[ruleSEMICOLON]() {
						goto l424
					}
					{
						position426, tokenIndex426, depth426 := position, tokenIndex, depth
						if !rules[rulepropertyListPath]() {
							goto l426
						}
						goto l427
					l426:
						position, tokenIndex, depth = position426, tokenIndex426, depth426
					}
				l427:
					goto l425
				l424:
					position, tokenIndex, depth = position424, tokenIndex424, depth424
				}
			l425:
				depth--
				add(rulepropertyListPath, position415)
			}
			return true
		l414:
			position, tokenIndex, depth = position414, tokenIndex414, depth414
			return false
		},
		/* 35 verbPath <- <path> */
		nil,
		/* 36 path <- <pathAlternative> */
		func() bool {
			position429, tokenIndex429, depth429 := position, tokenIndex, depth
			{
				position430 := position
				depth++
				{
					position431 := position
					depth++
					if !rules[rulepathSequence]() {
						goto l429
					}
				l432:
					{
						position433, tokenIndex433, depth433 := position, tokenIndex, depth
						if !rules[rulePIPE]() {
							goto l433
						}
						if !rules[rulepathSequence]() {
							goto l433
						}
						goto l432
					l433:
						position, tokenIndex, depth = position433, tokenIndex433, depth433
					}
					depth--
					add(rulepathAlternative, position431)
				}
				depth--
				add(rulepath, position430)
			}
			return true
		l429:
			position, tokenIndex, depth = position429, tokenIndex429, depth429
			return false
		},
		/* 37 pathAlternative <- <(pathSequence (PIPE pathSequence)*)> */
		nil,
		/* 38 pathSequence <- <(pathElt (SLASH pathElt)*)> */
		func() bool {
			position435, tokenIndex435, depth435 := position, tokenIndex, depth
			{
				position436 := position
				depth++
				if !rules[rulepathElt]() {
					goto l435
				}
			l437:
				{
					position438, tokenIndex438, depth438 := position, tokenIndex, depth
					if !rules[ruleSLASH]() {
						goto l438
					}
					if !rules[rulepathElt]() {
						goto l438
					}
					goto l437
				l438:
					position, tokenIndex, depth = position438, tokenIndex438, depth438
				}
				depth--
				add(rulepathSequence, position436)
			}
			return true
		l435:
			position, tokenIndex, depth = position435, tokenIndex435, depth435
			return false
		},
		/* 39 pathElt <- <(INVERSE? pathPrimary pathMod?)> */
		func() bool {
			position439, tokenIndex439, depth439 := position, tokenIndex, depth
			{
				position440 := position
				depth++
				{
					position441, tokenIndex441, depth441 := position, tokenIndex, depth
					if !rules[ruleINVERSE]() {
						goto l441
					}
					goto l442
				l441:
					position, tokenIndex, depth = position441, tokenIndex441, depth441
				}
			l442:
				{
					position443 := position
					depth++
					{
						position444, tokenIndex444, depth444 := position, tokenIndex, depth
						if !rules[ruleiriref]() {
							goto l445
						}
						goto l444
					l445:
						position, tokenIndex, depth = position444, tokenIndex444, depth444
						{
							switch buffer[position] {
							case '(':
								if !rules[ruleLPAREN]() {
									goto l439
								}
								if !rules[rulepath]() {
									goto l439
								}
								if !rules[ruleRPAREN]() {
									goto l439
								}
								break
							case '!':
								if !rules[ruleNOT]() {
									goto l439
								}
								{
									position447 := position
									depth++
									{
										position448, tokenIndex448, depth448 := position, tokenIndex, depth
										if !rules[rulepathOneInPropertySet]() {
											goto l449
										}
										goto l448
									l449:
										position, tokenIndex, depth = position448, tokenIndex448, depth448
										if !rules[ruleLPAREN]() {
											goto l439
										}
										{
											position450, tokenIndex450, depth450 := position, tokenIndex, depth
											if !rules[rulepathOneInPropertySet]() {
												goto l450
											}
										l452:
											{
												position453, tokenIndex453, depth453 := position, tokenIndex, depth
												if !rules[rulePIPE]() {
													goto l453
												}
												if !rules[rulepathOneInPropertySet]() {
													goto l453
												}
												goto l452
											l453:
												position, tokenIndex, depth = position453, tokenIndex453, depth453
											}
											goto l451
										l450:
											position, tokenIndex, depth = position450, tokenIndex450, depth450
										}
									l451:
										if !rules[ruleRPAREN]() {
											goto l439
										}
									}
								l448:
									depth--
									add(rulepathNegatedPropertySet, position447)
								}
								break
							default:
								if !rules[ruleISA]() {
									goto l439
								}
								break
							}
						}

					}
				l444:
					depth--
					add(rulepathPrimary, position443)
				}
				{
					position454, tokenIndex454, depth454 := position, tokenIndex, depth
					{
						position456 := position
						depth++
						{
							switch buffer[position] {
							case '+':
								if !rules[rulePLUS]() {
									goto l454
								}
								break
							case '?':
								{
									position458 := position
									depth++
									if buffer[position] != rune('?') {
										goto l454
									}
									position++
									if !rules[ruleskip]() {
										goto l454
									}
									depth--
									add(ruleQUESTION, position458)
								}
								break
							default:
								if !rules[ruleSTAR]() {
									goto l454
								}
								break
							}
						}

						{
							position459, tokenIndex459, depth459 := position, tokenIndex, depth
							if !matchDot() {
								goto l459
							}
							goto l454
						l459:
							position, tokenIndex, depth = position459, tokenIndex459, depth459
						}
						depth--
						add(rulepathMod, position456)
					}
					goto l455
				l454:
					position, tokenIndex, depth = position454, tokenIndex454, depth454
				}
			l455:
				depth--
				add(rulepathElt, position440)
			}
			return true
		l439:
			position, tokenIndex, depth = position439, tokenIndex439, depth439
			return false
		},
		/* 40 pathPrimary <- <(iriref / ((&('(') (LPAREN path RPAREN)) | (&('!') (NOT pathNegatedPropertySet)) | (&('a') ISA)))> */
		nil,
		/* 41 pathNegatedPropertySet <- <(pathOneInPropertySet / (LPAREN (pathOneInPropertySet (PIPE pathOneInPropertySet)*)? RPAREN))> */
		nil,
		/* 42 pathOneInPropertySet <- <(iriref / ISA / (INVERSE (iriref / ISA)))> */
		func() bool {
			position462, tokenIndex462, depth462 := position, tokenIndex, depth
			{
				position463 := position
				depth++
				{
					position464, tokenIndex464, depth464 := position, tokenIndex, depth
					if !rules[ruleiriref]() {
						goto l465
					}
					goto l464
				l465:
					position, tokenIndex, depth = position464, tokenIndex464, depth464
					if !rules[ruleISA]() {
						goto l466
					}
					goto l464
				l466:
					position, tokenIndex, depth = position464, tokenIndex464, depth464
					if !rules[ruleINVERSE]() {
						goto l462
					}
					{
						position467, tokenIndex467, depth467 := position, tokenIndex, depth
						if !rules[ruleiriref]() {
							goto l468
						}
						goto l467
					l468:
						position, tokenIndex, depth = position467, tokenIndex467, depth467
						if !rules[ruleISA]() {
							goto l462
						}
					}
				l467:
				}
			l464:
				depth--
				add(rulepathOneInPropertySet, position463)
			}
			return true
		l462:
			position, tokenIndex, depth = position462, tokenIndex462, depth462
			return false
		},
		/* 43 pathMod <- <(((&('+') PLUS) | (&('?') QUESTION) | (&('*') STAR)) !.)> */
		nil,
		/* 44 objectListPath <- <(objectPath (COMMA objectPath)*)> */
		nil,
		/* 45 objectPath <- <(<graphNodePath> Action2)> */
		func() bool {
			position471, tokenIndex471, depth471 := position, tokenIndex, depth
			{
				position472 := position
				depth++
				{
					position473 := position
					depth++
					if !rules[rulegraphNodePath]() {
						goto l471
					}
					depth--
					add(rulePegText, position473)
				}
				{
					add(ruleAction2, position)
				}
				depth--
				add(ruleobjectPath, position472)
			}
			return true
		l471:
			position, tokenIndex, depth = position471, tokenIndex471, depth471
			return false
		},
		/* 46 graphNodePath <- <(varOrTerm / triplesNodePath)> */
		func() bool {
			position475, tokenIndex475, depth475 := position, tokenIndex, depth
			{
				position476 := position
				depth++
				{
					position477, tokenIndex477, depth477 := position, tokenIndex, depth
					if !rules[rulevarOrTerm]() {
						goto l478
					}
					goto l477
				l478:
					position, tokenIndex, depth = position477, tokenIndex477, depth477
					if !rules[ruletriplesNodePath]() {
						goto l475
					}
				}
			l477:
				depth--
				add(rulegraphNodePath, position476)
			}
			return true
		l475:
			position, tokenIndex, depth = position475, tokenIndex475, depth475
			return false
		},
		/* 47 solutionModifier <- <((ORDER BY orderCondition+) / ((&('H' | 'h') (HAVING constraint)) | (&('G' | 'g') (GROUP BY groupCondition+)) | (&('L' | 'O' | 'l' | 'o') limitOffsetClauses)))?> */
		func() bool {
			{
				position480 := position
				depth++
				{
					position481, tokenIndex481, depth481 := position, tokenIndex, depth
					{
						position483, tokenIndex483, depth483 := position, tokenIndex, depth
						{
							position485 := position
							depth++
							{
								position486, tokenIndex486, depth486 := position, tokenIndex, depth
								if buffer[position] != rune('o') {
									goto l487
								}
								position++
								goto l486
							l487:
								position, tokenIndex, depth = position486, tokenIndex486, depth486
								if buffer[position] != rune('O') {
									goto l484
								}
								position++
							}
						l486:
							{
								position488, tokenIndex488, depth488 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l489
								}
								position++
								goto l488
							l489:
								position, tokenIndex, depth = position488, tokenIndex488, depth488
								if buffer[position] != rune('R') {
									goto l484
								}
								position++
							}
						l488:
							{
								position490, tokenIndex490, depth490 := position, tokenIndex, depth
								if buffer[position] != rune('d') {
									goto l491
								}
								position++
								goto l490
							l491:
								position, tokenIndex, depth = position490, tokenIndex490, depth490
								if buffer[position] != rune('D') {
									goto l484
								}
								position++
							}
						l490:
							{
								position492, tokenIndex492, depth492 := position, tokenIndex, depth
								if buffer[position] != rune('e') {
									goto l493
								}
								position++
								goto l492
							l493:
								position, tokenIndex, depth = position492, tokenIndex492, depth492
								if buffer[position] != rune('E') {
									goto l484
								}
								position++
							}
						l492:
							{
								position494, tokenIndex494, depth494 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l495
								}
								position++
								goto l494
							l495:
								position, tokenIndex, depth = position494, tokenIndex494, depth494
								if buffer[position] != rune('R') {
									goto l484
								}
								position++
							}
						l494:
							if !rules[ruleskip]() {
								goto l484
							}
							depth--
							add(ruleORDER, position485)
						}
						if !rules[ruleBY]() {
							goto l484
						}
						{
							position498 := position
							depth++
							{
								position499, tokenIndex499, depth499 := position, tokenIndex, depth
								{
									position501, tokenIndex501, depth501 := position, tokenIndex, depth
									{
										position503, tokenIndex503, depth503 := position, tokenIndex, depth
										{
											position505 := position
											depth++
											{
												position506, tokenIndex506, depth506 := position, tokenIndex, depth
												if buffer[position] != rune('a') {
													goto l507
												}
												position++
												goto l506
											l507:
												position, tokenIndex, depth = position506, tokenIndex506, depth506
												if buffer[position] != rune('A') {
													goto l504
												}
												position++
											}
										l506:
											{
												position508, tokenIndex508, depth508 := position, tokenIndex, depth
												if buffer[position] != rune('s') {
													goto l509
												}
												position++
												goto l508
											l509:
												position, tokenIndex, depth = position508, tokenIndex508, depth508
												if buffer[position] != rune('S') {
													goto l504
												}
												position++
											}
										l508:
											{
												position510, tokenIndex510, depth510 := position, tokenIndex, depth
												if buffer[position] != rune('c') {
													goto l511
												}
												position++
												goto l510
											l511:
												position, tokenIndex, depth = position510, tokenIndex510, depth510
												if buffer[position] != rune('C') {
													goto l504
												}
												position++
											}
										l510:
											if !rules[ruleskip]() {
												goto l504
											}
											depth--
											add(ruleASC, position505)
										}
										goto l503
									l504:
										position, tokenIndex, depth = position503, tokenIndex503, depth503
										{
											position512 := position
											depth++
											{
												position513, tokenIndex513, depth513 := position, tokenIndex, depth
												if buffer[position] != rune('d') {
													goto l514
												}
												position++
												goto l513
											l514:
												position, tokenIndex, depth = position513, tokenIndex513, depth513
												if buffer[position] != rune('D') {
													goto l501
												}
												position++
											}
										l513:
											{
												position515, tokenIndex515, depth515 := position, tokenIndex, depth
												if buffer[position] != rune('e') {
													goto l516
												}
												position++
												goto l515
											l516:
												position, tokenIndex, depth = position515, tokenIndex515, depth515
												if buffer[position] != rune('E') {
													goto l501
												}
												position++
											}
										l515:
											{
												position517, tokenIndex517, depth517 := position, tokenIndex, depth
												if buffer[position] != rune('s') {
													goto l518
												}
												position++
												goto l517
											l518:
												position, tokenIndex, depth = position517, tokenIndex517, depth517
												if buffer[position] != rune('S') {
													goto l501
												}
												position++
											}
										l517:
											{
												position519, tokenIndex519, depth519 := position, tokenIndex, depth
												if buffer[position] != rune('c') {
													goto l520
												}
												position++
												goto l519
											l520:
												position, tokenIndex, depth = position519, tokenIndex519, depth519
												if buffer[position] != rune('C') {
													goto l501
												}
												position++
											}
										l519:
											if !rules[ruleskip]() {
												goto l501
											}
											depth--
											add(ruleDESC, position512)
										}
									}
								l503:
									goto l502
								l501:
									position, tokenIndex, depth = position501, tokenIndex501, depth501
								}
							l502:
								if !rules[rulebrackettedExpression]() {
									goto l500
								}
								goto l499
							l500:
								position, tokenIndex, depth = position499, tokenIndex499, depth499
								if !rules[rulefunctionCall]() {
									goto l521
								}
								goto l499
							l521:
								position, tokenIndex, depth = position499, tokenIndex499, depth499
								if !rules[rulebuiltinCall]() {
									goto l522
								}
								goto l499
							l522:
								position, tokenIndex, depth = position499, tokenIndex499, depth499
								if !rules[rulevar]() {
									goto l484
								}
							}
						l499:
							depth--
							add(ruleorderCondition, position498)
						}
					l496:
						{
							position497, tokenIndex497, depth497 := position, tokenIndex, depth
							{
								position523 := position
								depth++
								{
									position524, tokenIndex524, depth524 := position, tokenIndex, depth
									{
										position526, tokenIndex526, depth526 := position, tokenIndex, depth
										{
											position528, tokenIndex528, depth528 := position, tokenIndex, depth
											{
												position530 := position
												depth++
												{
													position531, tokenIndex531, depth531 := position, tokenIndex, depth
													if buffer[position] != rune('a') {
														goto l532
													}
													position++
													goto l531
												l532:
													position, tokenIndex, depth = position531, tokenIndex531, depth531
													if buffer[position] != rune('A') {
														goto l529
													}
													position++
												}
											l531:
												{
													position533, tokenIndex533, depth533 := position, tokenIndex, depth
													if buffer[position] != rune('s') {
														goto l534
													}
													position++
													goto l533
												l534:
													position, tokenIndex, depth = position533, tokenIndex533, depth533
													if buffer[position] != rune('S') {
														goto l529
													}
													position++
												}
											l533:
												{
													position535, tokenIndex535, depth535 := position, tokenIndex, depth
													if buffer[position] != rune('c') {
														goto l536
													}
													position++
													goto l535
												l536:
													position, tokenIndex, depth = position535, tokenIndex535, depth535
													if buffer[position] != rune('C') {
														goto l529
													}
													position++
												}
											l535:
												if !rules[ruleskip]() {
													goto l529
												}
												depth--
												add(ruleASC, position530)
											}
											goto l528
										l529:
											position, tokenIndex, depth = position528, tokenIndex528, depth528
											{
												position537 := position
												depth++
												{
													position538, tokenIndex538, depth538 := position, tokenIndex, depth
													if buffer[position] != rune('d') {
														goto l539
													}
													position++
													goto l538
												l539:
													position, tokenIndex, depth = position538, tokenIndex538, depth538
													if buffer[position] != rune('D') {
														goto l526
													}
													position++
												}
											l538:
												{
													position540, tokenIndex540, depth540 := position, tokenIndex, depth
													if buffer[position] != rune('e') {
														goto l541
													}
													position++
													goto l540
												l541:
													position, tokenIndex, depth = position540, tokenIndex540, depth540
													if buffer[position] != rune('E') {
														goto l526
													}
													position++
												}
											l540:
												{
													position542, tokenIndex542, depth542 := position, tokenIndex, depth
													if buffer[position] != rune('s') {
														goto l543
													}
													position++
													goto l542
												l543:
													position, tokenIndex, depth = position542, tokenIndex542, depth542
													if buffer[position] != rune('S') {
														goto l526
													}
													position++
												}
											l542:
												{
													position544, tokenIndex544, depth544 := position, tokenIndex, depth
													if buffer[position] != rune('c') {
														goto l545
													}
													position++
													goto l544
												l545:
													position, tokenIndex, depth = position544, tokenIndex544, depth544
													if buffer[position] != rune('C') {
														goto l526
													}
													position++
												}
											l544:
												if !rules[ruleskip]() {
													goto l526
												}
												depth--
												add(ruleDESC, position537)
											}
										}
									l528:
										goto l527
									l526:
										position, tokenIndex, depth = position526, tokenIndex526, depth526
									}
								l527:
									if !rules[rulebrackettedExpression]() {
										goto l525
									}
									goto l524
								l525:
									position, tokenIndex, depth = position524, tokenIndex524, depth524
									if !rules[rulefunctionCall]() {
										goto l546
									}
									goto l524
								l546:
									position, tokenIndex, depth = position524, tokenIndex524, depth524
									if !rules[rulebuiltinCall]() {
										goto l547
									}
									goto l524
								l547:
									position, tokenIndex, depth = position524, tokenIndex524, depth524
									if !rules[rulevar]() {
										goto l497
									}
								}
							l524:
								depth--
								add(ruleorderCondition, position523)
							}
							goto l496
						l497:
							position, tokenIndex, depth = position497, tokenIndex497, depth497
						}
						goto l483
					l484:
						position, tokenIndex, depth = position483, tokenIndex483, depth483
						{
							switch buffer[position] {
							case 'H', 'h':
								{
									position549 := position
									depth++
									{
										position550, tokenIndex550, depth550 := position, tokenIndex, depth
										if buffer[position] != rune('h') {
											goto l551
										}
										position++
										goto l550
									l551:
										position, tokenIndex, depth = position550, tokenIndex550, depth550
										if buffer[position] != rune('H') {
											goto l481
										}
										position++
									}
								l550:
									{
										position552, tokenIndex552, depth552 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l553
										}
										position++
										goto l552
									l553:
										position, tokenIndex, depth = position552, tokenIndex552, depth552
										if buffer[position] != rune('A') {
											goto l481
										}
										position++
									}
								l552:
									{
										position554, tokenIndex554, depth554 := position, tokenIndex, depth
										if buffer[position] != rune('v') {
											goto l555
										}
										position++
										goto l554
									l555:
										position, tokenIndex, depth = position554, tokenIndex554, depth554
										if buffer[position] != rune('V') {
											goto l481
										}
										position++
									}
								l554:
									{
										position556, tokenIndex556, depth556 := position, tokenIndex, depth
										if buffer[position] != rune('i') {
											goto l557
										}
										position++
										goto l556
									l557:
										position, tokenIndex, depth = position556, tokenIndex556, depth556
										if buffer[position] != rune('I') {
											goto l481
										}
										position++
									}
								l556:
									{
										position558, tokenIndex558, depth558 := position, tokenIndex, depth
										if buffer[position] != rune('n') {
											goto l559
										}
										position++
										goto l558
									l559:
										position, tokenIndex, depth = position558, tokenIndex558, depth558
										if buffer[position] != rune('N') {
											goto l481
										}
										position++
									}
								l558:
									{
										position560, tokenIndex560, depth560 := position, tokenIndex, depth
										if buffer[position] != rune('g') {
											goto l561
										}
										position++
										goto l560
									l561:
										position, tokenIndex, depth = position560, tokenIndex560, depth560
										if buffer[position] != rune('G') {
											goto l481
										}
										position++
									}
								l560:
									if !rules[ruleskip]() {
										goto l481
									}
									depth--
									add(ruleHAVING, position549)
								}
								if !rules[ruleconstraint]() {
									goto l481
								}
								break
							case 'G', 'g':
								{
									position562 := position
									depth++
									{
										position563, tokenIndex563, depth563 := position, tokenIndex, depth
										if buffer[position] != rune('g') {
											goto l564
										}
										position++
										goto l563
									l564:
										position, tokenIndex, depth = position563, tokenIndex563, depth563
										if buffer[position] != rune('G') {
											goto l481
										}
										position++
									}
								l563:
									{
										position565, tokenIndex565, depth565 := position, tokenIndex, depth
										if buffer[position] != rune('r') {
											goto l566
										}
										position++
										goto l565
									l566:
										position, tokenIndex, depth = position565, tokenIndex565, depth565
										if buffer[position] != rune('R') {
											goto l481
										}
										position++
									}
								l565:
									{
										position567, tokenIndex567, depth567 := position, tokenIndex, depth
										if buffer[position] != rune('o') {
											goto l568
										}
										position++
										goto l567
									l568:
										position, tokenIndex, depth = position567, tokenIndex567, depth567
										if buffer[position] != rune('O') {
											goto l481
										}
										position++
									}
								l567:
									{
										position569, tokenIndex569, depth569 := position, tokenIndex, depth
										if buffer[position] != rune('u') {
											goto l570
										}
										position++
										goto l569
									l570:
										position, tokenIndex, depth = position569, tokenIndex569, depth569
										if buffer[position] != rune('U') {
											goto l481
										}
										position++
									}
								l569:
									{
										position571, tokenIndex571, depth571 := position, tokenIndex, depth
										if buffer[position] != rune('p') {
											goto l572
										}
										position++
										goto l571
									l572:
										position, tokenIndex, depth = position571, tokenIndex571, depth571
										if buffer[position] != rune('P') {
											goto l481
										}
										position++
									}
								l571:
									if !rules[ruleskip]() {
										goto l481
									}
									depth--
									add(ruleGROUP, position562)
								}
								if !rules[ruleBY]() {
									goto l481
								}
								{
									position575 := position
									depth++
									{
										position576, tokenIndex576, depth576 := position, tokenIndex, depth
										if !rules[rulefunctionCall]() {
											goto l577
										}
										goto l576
									l577:
										position, tokenIndex, depth = position576, tokenIndex576, depth576
										{
											switch buffer[position] {
											case '$', '?':
												if !rules[rulevar]() {
													goto l481
												}
												break
											case '(':
												if !rules[ruleLPAREN]() {
													goto l481
												}
												if !rules[ruleexpression]() {
													goto l481
												}
												{
													position579, tokenIndex579, depth579 := position, tokenIndex, depth
													if !rules[ruleAS]() {
														goto l579
													}
													if !rules[rulevar]() {
														goto l579
													}
													goto l580
												l579:
													position, tokenIndex, depth = position579, tokenIndex579, depth579
												}
											l580:
												if !rules[ruleRPAREN]() {
													goto l481
												}
												break
											default:
												if !rules[rulebuiltinCall]() {
													goto l481
												}
												break
											}
										}

									}
								l576:
									depth--
									add(rulegroupCondition, position575)
								}
							l573:
								{
									position574, tokenIndex574, depth574 := position, tokenIndex, depth
									{
										position581 := position
										depth++
										{
											position582, tokenIndex582, depth582 := position, tokenIndex, depth
											if !rules[rulefunctionCall]() {
												goto l583
											}
											goto l582
										l583:
											position, tokenIndex, depth = position582, tokenIndex582, depth582
											{
												switch buffer[position] {
												case '$', '?':
													if !rules[rulevar]() {
														goto l574
													}
													break
												case '(':
													if !rules[ruleLPAREN]() {
														goto l574
													}
													if !rules[ruleexpression]() {
														goto l574
													}
													{
														position585, tokenIndex585, depth585 := position, tokenIndex, depth
														if !rules[ruleAS]() {
															goto l585
														}
														if !rules[rulevar]() {
															goto l585
														}
														goto l586
													l585:
														position, tokenIndex, depth = position585, tokenIndex585, depth585
													}
												l586:
													if !rules[ruleRPAREN]() {
														goto l574
													}
													break
												default:
													if !rules[rulebuiltinCall]() {
														goto l574
													}
													break
												}
											}

										}
									l582:
										depth--
										add(rulegroupCondition, position581)
									}
									goto l573
								l574:
									position, tokenIndex, depth = position574, tokenIndex574, depth574
								}
								break
							default:
								{
									position587 := position
									depth++
									{
										position588, tokenIndex588, depth588 := position, tokenIndex, depth
										if !rules[rulelimit]() {
											goto l589
										}
										{
											position590, tokenIndex590, depth590 := position, tokenIndex, depth
											if !rules[ruleoffset]() {
												goto l590
											}
											goto l591
										l590:
											position, tokenIndex, depth = position590, tokenIndex590, depth590
										}
									l591:
										goto l588
									l589:
										position, tokenIndex, depth = position588, tokenIndex588, depth588
										if !rules[ruleoffset]() {
											goto l481
										}
										{
											position592, tokenIndex592, depth592 := position, tokenIndex, depth
											if !rules[rulelimit]() {
												goto l592
											}
											goto l593
										l592:
											position, tokenIndex, depth = position592, tokenIndex592, depth592
										}
									l593:
									}
								l588:
									depth--
									add(rulelimitOffsetClauses, position587)
								}
								break
							}
						}

					}
				l483:
					goto l482
				l481:
					position, tokenIndex, depth = position481, tokenIndex481, depth481
				}
			l482:
				depth--
				add(rulesolutionModifier, position480)
			}
			return true
		},
		/* 48 groupCondition <- <(functionCall / ((&('$' | '?') var) | (&('(') (LPAREN expression (AS var)? RPAREN)) | (&('A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'H' | 'I' | 'L' | 'M' | 'N' | 'R' | 'S' | 'T' | 'U' | 'Y' | 'a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'h' | 'i' | 'l' | 'm' | 'n' | 'r' | 's' | 't' | 'u' | 'y') builtinCall)))> */
		nil,
		/* 49 orderCondition <- <(((ASC / DESC)? brackettedExpression) / functionCall / builtinCall / var)> */
		nil,
		/* 50 limitOffsetClauses <- <((limit offset?) / (offset limit?))> */
		nil,
		/* 51 limit <- <(LIMIT INTEGER)> */
		func() bool {
			position597, tokenIndex597, depth597 := position, tokenIndex, depth
			{
				position598 := position
				depth++
				{
					position599 := position
					depth++
					{
						position600, tokenIndex600, depth600 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l601
						}
						position++
						goto l600
					l601:
						position, tokenIndex, depth = position600, tokenIndex600, depth600
						if buffer[position] != rune('L') {
							goto l597
						}
						position++
					}
				l600:
					{
						position602, tokenIndex602, depth602 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l603
						}
						position++
						goto l602
					l603:
						position, tokenIndex, depth = position602, tokenIndex602, depth602
						if buffer[position] != rune('I') {
							goto l597
						}
						position++
					}
				l602:
					{
						position604, tokenIndex604, depth604 := position, tokenIndex, depth
						if buffer[position] != rune('m') {
							goto l605
						}
						position++
						goto l604
					l605:
						position, tokenIndex, depth = position604, tokenIndex604, depth604
						if buffer[position] != rune('M') {
							goto l597
						}
						position++
					}
				l604:
					{
						position606, tokenIndex606, depth606 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l607
						}
						position++
						goto l606
					l607:
						position, tokenIndex, depth = position606, tokenIndex606, depth606
						if buffer[position] != rune('I') {
							goto l597
						}
						position++
					}
				l606:
					{
						position608, tokenIndex608, depth608 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l609
						}
						position++
						goto l608
					l609:
						position, tokenIndex, depth = position608, tokenIndex608, depth608
						if buffer[position] != rune('T') {
							goto l597
						}
						position++
					}
				l608:
					if !rules[ruleskip]() {
						goto l597
					}
					depth--
					add(ruleLIMIT, position599)
				}
				if !rules[ruleINTEGER]() {
					goto l597
				}
				depth--
				add(rulelimit, position598)
			}
			return true
		l597:
			position, tokenIndex, depth = position597, tokenIndex597, depth597
			return false
		},
		/* 52 offset <- <(OFFSET INTEGER)> */
		func() bool {
			position610, tokenIndex610, depth610 := position, tokenIndex, depth
			{
				position611 := position
				depth++
				{
					position612 := position
					depth++
					{
						position613, tokenIndex613, depth613 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l614
						}
						position++
						goto l613
					l614:
						position, tokenIndex, depth = position613, tokenIndex613, depth613
						if buffer[position] != rune('O') {
							goto l610
						}
						position++
					}
				l613:
					{
						position615, tokenIndex615, depth615 := position, tokenIndex, depth
						if buffer[position] != rune('f') {
							goto l616
						}
						position++
						goto l615
					l616:
						position, tokenIndex, depth = position615, tokenIndex615, depth615
						if buffer[position] != rune('F') {
							goto l610
						}
						position++
					}
				l615:
					{
						position617, tokenIndex617, depth617 := position, tokenIndex, depth
						if buffer[position] != rune('f') {
							goto l618
						}
						position++
						goto l617
					l618:
						position, tokenIndex, depth = position617, tokenIndex617, depth617
						if buffer[position] != rune('F') {
							goto l610
						}
						position++
					}
				l617:
					{
						position619, tokenIndex619, depth619 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l620
						}
						position++
						goto l619
					l620:
						position, tokenIndex, depth = position619, tokenIndex619, depth619
						if buffer[position] != rune('S') {
							goto l610
						}
						position++
					}
				l619:
					{
						position621, tokenIndex621, depth621 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l622
						}
						position++
						goto l621
					l622:
						position, tokenIndex, depth = position621, tokenIndex621, depth621
						if buffer[position] != rune('E') {
							goto l610
						}
						position++
					}
				l621:
					{
						position623, tokenIndex623, depth623 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l624
						}
						position++
						goto l623
					l624:
						position, tokenIndex, depth = position623, tokenIndex623, depth623
						if buffer[position] != rune('T') {
							goto l610
						}
						position++
					}
				l623:
					if !rules[ruleskip]() {
						goto l610
					}
					depth--
					add(ruleOFFSET, position612)
				}
				if !rules[ruleINTEGER]() {
					goto l610
				}
				depth--
				add(ruleoffset, position611)
			}
			return true
		l610:
			position, tokenIndex, depth = position610, tokenIndex610, depth610
			return false
		},
		/* 53 expression <- <conditionalOrExpression> */
		func() bool {
			position625, tokenIndex625, depth625 := position, tokenIndex, depth
			{
				position626 := position
				depth++
				if !rules[ruleconditionalOrExpression]() {
					goto l625
				}
				depth--
				add(ruleexpression, position626)
			}
			return true
		l625:
			position, tokenIndex, depth = position625, tokenIndex625, depth625
			return false
		},
		/* 54 conditionalOrExpression <- <(conditionalAndExpression (OR conditionalOrExpression)?)> */
		func() bool {
			position627, tokenIndex627, depth627 := position, tokenIndex, depth
			{
				position628 := position
				depth++
				if !rules[ruleconditionalAndExpression]() {
					goto l627
				}
				{
					position629, tokenIndex629, depth629 := position, tokenIndex, depth
					{
						position631 := position
						depth++
						if buffer[position] != rune('|') {
							goto l629
						}
						position++
						if buffer[position] != rune('|') {
							goto l629
						}
						position++
						if !rules[ruleskip]() {
							goto l629
						}
						depth--
						add(ruleOR, position631)
					}
					if !rules[ruleconditionalOrExpression]() {
						goto l629
					}
					goto l630
				l629:
					position, tokenIndex, depth = position629, tokenIndex629, depth629
				}
			l630:
				depth--
				add(ruleconditionalOrExpression, position628)
			}
			return true
		l627:
			position, tokenIndex, depth = position627, tokenIndex627, depth627
			return false
		},
		/* 55 conditionalAndExpression <- <(valueLogical (AND conditionalAndExpression)?)> */
		func() bool {
			position632, tokenIndex632, depth632 := position, tokenIndex, depth
			{
				position633 := position
				depth++
				{
					position634 := position
					depth++
					if !rules[rulenumericExpression]() {
						goto l632
					}
					{
						position635, tokenIndex635, depth635 := position, tokenIndex, depth
						{
							switch buffer[position] {
							case 'N', 'n':
								{
									position638 := position
									depth++
									{
										position639 := position
										depth++
										{
											position640, tokenIndex640, depth640 := position, tokenIndex, depth
											if buffer[position] != rune('n') {
												goto l641
											}
											position++
											goto l640
										l641:
											position, tokenIndex, depth = position640, tokenIndex640, depth640
											if buffer[position] != rune('N') {
												goto l635
											}
											position++
										}
									l640:
										{
											position642, tokenIndex642, depth642 := position, tokenIndex, depth
											if buffer[position] != rune('o') {
												goto l643
											}
											position++
											goto l642
										l643:
											position, tokenIndex, depth = position642, tokenIndex642, depth642
											if buffer[position] != rune('O') {
												goto l635
											}
											position++
										}
									l642:
										{
											position644, tokenIndex644, depth644 := position, tokenIndex, depth
											if buffer[position] != rune('t') {
												goto l645
											}
											position++
											goto l644
										l645:
											position, tokenIndex, depth = position644, tokenIndex644, depth644
											if buffer[position] != rune('T') {
												goto l635
											}
											position++
										}
									l644:
										if buffer[position] != rune(' ') {
											goto l635
										}
										position++
										{
											position646, tokenIndex646, depth646 := position, tokenIndex, depth
											if buffer[position] != rune('i') {
												goto l647
											}
											position++
											goto l646
										l647:
											position, tokenIndex, depth = position646, tokenIndex646, depth646
											if buffer[position] != rune('I') {
												goto l635
											}
											position++
										}
									l646:
										{
											position648, tokenIndex648, depth648 := position, tokenIndex, depth
											if buffer[position] != rune('n') {
												goto l649
											}
											position++
											goto l648
										l649:
											position, tokenIndex, depth = position648, tokenIndex648, depth648
											if buffer[position] != rune('N') {
												goto l635
											}
											position++
										}
									l648:
										if !rules[ruleskip]() {
											goto l635
										}
										depth--
										add(ruleNOTIN, position639)
									}
									if !rules[ruleargList]() {
										goto l635
									}
									depth--
									add(rulenotin, position638)
								}
								break
							case 'I', 'i':
								{
									position650 := position
									depth++
									{
										position651 := position
										depth++
										{
											position652, tokenIndex652, depth652 := position, tokenIndex, depth
											if buffer[position] != rune('i') {
												goto l653
											}
											position++
											goto l652
										l653:
											position, tokenIndex, depth = position652, tokenIndex652, depth652
											if buffer[position] != rune('I') {
												goto l635
											}
											position++
										}
									l652:
										{
											position654, tokenIndex654, depth654 := position, tokenIndex, depth
											if buffer[position] != rune('n') {
												goto l655
											}
											position++
											goto l654
										l655:
											position, tokenIndex, depth = position654, tokenIndex654, depth654
											if buffer[position] != rune('N') {
												goto l635
											}
											position++
										}
									l654:
										if !rules[ruleskip]() {
											goto l635
										}
										depth--
										add(ruleIN, position651)
									}
									if !rules[ruleargList]() {
										goto l635
									}
									depth--
									add(rulein, position650)
								}
								break
							default:
								{
									position656, tokenIndex656, depth656 := position, tokenIndex, depth
									{
										position658 := position
										depth++
										if buffer[position] != rune('<') {
											goto l657
										}
										position++
										if !rules[ruleskip]() {
											goto l657
										}
										depth--
										add(ruleLT, position658)
									}
									goto l656
								l657:
									position, tokenIndex, depth = position656, tokenIndex656, depth656
									{
										position660 := position
										depth++
										if buffer[position] != rune('>') {
											goto l659
										}
										position++
										if buffer[position] != rune('=') {
											goto l659
										}
										position++
										if !rules[ruleskip]() {
											goto l659
										}
										depth--
										add(ruleGE, position660)
									}
									goto l656
								l659:
									position, tokenIndex, depth = position656, tokenIndex656, depth656
									{
										switch buffer[position] {
										case '>':
											{
												position662 := position
												depth++
												if buffer[position] != rune('>') {
													goto l635
												}
												position++
												if !rules[ruleskip]() {
													goto l635
												}
												depth--
												add(ruleGT, position662)
											}
											break
										case '<':
											{
												position663 := position
												depth++
												if buffer[position] != rune('<') {
													goto l635
												}
												position++
												if buffer[position] != rune('=') {
													goto l635
												}
												position++
												if !rules[ruleskip]() {
													goto l635
												}
												depth--
												add(ruleLE, position663)
											}
											break
										case '!':
											{
												position664 := position
												depth++
												if buffer[position] != rune('!') {
													goto l635
												}
												position++
												if buffer[position] != rune('=') {
													goto l635
												}
												position++
												if !rules[ruleskip]() {
													goto l635
												}
												depth--
												add(ruleNE, position664)
											}
											break
										default:
											if !rules[ruleEQ]() {
												goto l635
											}
											break
										}
									}

								}
							l656:
								if !rules[rulenumericExpression]() {
									goto l635
								}
								break
							}
						}

						goto l636
					l635:
						position, tokenIndex, depth = position635, tokenIndex635, depth635
					}
				l636:
					depth--
					add(rulevalueLogical, position634)
				}
				{
					position665, tokenIndex665, depth665 := position, tokenIndex, depth
					{
						position667 := position
						depth++
						if buffer[position] != rune('&') {
							goto l665
						}
						position++
						if buffer[position] != rune('&') {
							goto l665
						}
						position++
						if !rules[ruleskip]() {
							goto l665
						}
						depth--
						add(ruleAND, position667)
					}
					if !rules[ruleconditionalAndExpression]() {
						goto l665
					}
					goto l666
				l665:
					position, tokenIndex, depth = position665, tokenIndex665, depth665
				}
			l666:
				depth--
				add(ruleconditionalAndExpression, position633)
			}
			return true
		l632:
			position, tokenIndex, depth = position632, tokenIndex632, depth632
			return false
		},
		/* 56 valueLogical <- <(numericExpression ((&('N' | 'n') notin) | (&('I' | 'i') in) | (&('!' | '<' | '=' | '>') ((LT / GE / ((&('>') GT) | (&('<') LE) | (&('!') NE) | (&('=') EQ))) numericExpression)))?)> */
		nil,
		/* 57 numericExpression <- <(multiplicativeExpression (((PLUS / MINUS) multiplicativeExpression) / signedNumericLiteral)*)> */
		func() bool {
			position669, tokenIndex669, depth669 := position, tokenIndex, depth
			{
				position670 := position
				depth++
				if !rules[rulemultiplicativeExpression]() {
					goto l669
				}
			l671:
				{
					position672, tokenIndex672, depth672 := position, tokenIndex, depth
					{
						position673, tokenIndex673, depth673 := position, tokenIndex, depth
						{
							position675, tokenIndex675, depth675 := position, tokenIndex, depth
							if !rules[rulePLUS]() {
								goto l676
							}
							goto l675
						l676:
							position, tokenIndex, depth = position675, tokenIndex675, depth675
							if !rules[ruleMINUS]() {
								goto l674
							}
						}
					l675:
						if !rules[rulemultiplicativeExpression]() {
							goto l674
						}
						goto l673
					l674:
						position, tokenIndex, depth = position673, tokenIndex673, depth673
						{
							position677 := position
							depth++
							{
								position678, tokenIndex678, depth678 := position, tokenIndex, depth
								if buffer[position] != rune('+') {
									goto l679
								}
								position++
								goto l678
							l679:
								position, tokenIndex, depth = position678, tokenIndex678, depth678
								if buffer[position] != rune('-') {
									goto l672
								}
								position++
							}
						l678:
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l672
							}
							position++
						l680:
							{
								position681, tokenIndex681, depth681 := position, tokenIndex, depth
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l681
								}
								position++
								goto l680
							l681:
								position, tokenIndex, depth = position681, tokenIndex681, depth681
							}
							{
								position682, tokenIndex682, depth682 := position, tokenIndex, depth
								if buffer[position] != rune('.') {
									goto l682
								}
								position++
							l684:
								{
									position685, tokenIndex685, depth685 := position, tokenIndex, depth
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l685
									}
									position++
									goto l684
								l685:
									position, tokenIndex, depth = position685, tokenIndex685, depth685
								}
								goto l683
							l682:
								position, tokenIndex, depth = position682, tokenIndex682, depth682
							}
						l683:
							if !rules[ruleskip]() {
								goto l672
							}
							depth--
							add(rulesignedNumericLiteral, position677)
						}
					}
				l673:
					goto l671
				l672:
					position, tokenIndex, depth = position672, tokenIndex672, depth672
				}
				depth--
				add(rulenumericExpression, position670)
			}
			return true
		l669:
			position, tokenIndex, depth = position669, tokenIndex669, depth669
			return false
		},
		/* 58 multiplicativeExpression <- <(unaryExpression ((STAR / SLASH) unaryExpression)*)> */
		func() bool {
			position686, tokenIndex686, depth686 := position, tokenIndex, depth
			{
				position687 := position
				depth++
				if !rules[ruleunaryExpression]() {
					goto l686
				}
			l688:
				{
					position689, tokenIndex689, depth689 := position, tokenIndex, depth
					{
						position690, tokenIndex690, depth690 := position, tokenIndex, depth
						if !rules[ruleSTAR]() {
							goto l691
						}
						goto l690
					l691:
						position, tokenIndex, depth = position690, tokenIndex690, depth690
						if !rules[ruleSLASH]() {
							goto l689
						}
					}
				l690:
					if !rules[ruleunaryExpression]() {
						goto l689
					}
					goto l688
				l689:
					position, tokenIndex, depth = position689, tokenIndex689, depth689
				}
				depth--
				add(rulemultiplicativeExpression, position687)
			}
			return true
		l686:
			position, tokenIndex, depth = position686, tokenIndex686, depth686
			return false
		},
		/* 59 unaryExpression <- <(((&('+') PLUS) | (&('-') MINUS) | (&('!') NOT))? primaryExpression)> */
		func() bool {
			position692, tokenIndex692, depth692 := position, tokenIndex, depth
			{
				position693 := position
				depth++
				{
					position694, tokenIndex694, depth694 := position, tokenIndex, depth
					{
						switch buffer[position] {
						case '+':
							if !rules[rulePLUS]() {
								goto l694
							}
							break
						case '-':
							if !rules[ruleMINUS]() {
								goto l694
							}
							break
						default:
							if !rules[ruleNOT]() {
								goto l694
							}
							break
						}
					}

					goto l695
				l694:
					position, tokenIndex, depth = position694, tokenIndex694, depth694
				}
			l695:
				{
					position697 := position
					depth++
					{
						position698, tokenIndex698, depth698 := position, tokenIndex, depth
						if !rules[rulebuiltinCall]() {
							goto l699
						}
						goto l698
					l699:
						position, tokenIndex, depth = position698, tokenIndex698, depth698
						if !rules[rulefunctionCall]() {
							goto l700
						}
						goto l698
					l700:
						position, tokenIndex, depth = position698, tokenIndex698, depth698
						if !rules[ruleiriref]() {
							goto l701
						}
						goto l698
					l701:
						position, tokenIndex, depth = position698, tokenIndex698, depth698
						{
							switch buffer[position] {
							case 'A', 'C', 'G', 'M', 'S', 'a', 'c', 'g', 'm', 's':
								{
									position703 := position
									depth++
									{
										switch buffer[position] {
										case 'G', 'g':
											{
												position705 := position
												depth++
												{
													position706 := position
													depth++
													{
														position707, tokenIndex707, depth707 := position, tokenIndex, depth
														if buffer[position] != rune('g') {
															goto l708
														}
														position++
														goto l707
													l708:
														position, tokenIndex, depth = position707, tokenIndex707, depth707
														if buffer[position] != rune('G') {
															goto l692
														}
														position++
													}
												l707:
													{
														position709, tokenIndex709, depth709 := position, tokenIndex, depth
														if buffer[position] != rune('r') {
															goto l710
														}
														position++
														goto l709
													l710:
														position, tokenIndex, depth = position709, tokenIndex709, depth709
														if buffer[position] != rune('R') {
															goto l692
														}
														position++
													}
												l709:
													{
														position711, tokenIndex711, depth711 := position, tokenIndex, depth
														if buffer[position] != rune('o') {
															goto l712
														}
														position++
														goto l711
													l712:
														position, tokenIndex, depth = position711, tokenIndex711, depth711
														if buffer[position] != rune('O') {
															goto l692
														}
														position++
													}
												l711:
													{
														position713, tokenIndex713, depth713 := position, tokenIndex, depth
														if buffer[position] != rune('u') {
															goto l714
														}
														position++
														goto l713
													l714:
														position, tokenIndex, depth = position713, tokenIndex713, depth713
														if buffer[position] != rune('U') {
															goto l692
														}
														position++
													}
												l713:
													{
														position715, tokenIndex715, depth715 := position, tokenIndex, depth
														if buffer[position] != rune('p') {
															goto l716
														}
														position++
														goto l715
													l716:
														position, tokenIndex, depth = position715, tokenIndex715, depth715
														if buffer[position] != rune('P') {
															goto l692
														}
														position++
													}
												l715:
													if buffer[position] != rune('_') {
														goto l692
													}
													position++
													{
														position717, tokenIndex717, depth717 := position, tokenIndex, depth
														if buffer[position] != rune('c') {
															goto l718
														}
														position++
														goto l717
													l718:
														position, tokenIndex, depth = position717, tokenIndex717, depth717
														if buffer[position] != rune('C') {
															goto l692
														}
														position++
													}
												l717:
													{
														position719, tokenIndex719, depth719 := position, tokenIndex, depth
														if buffer[position] != rune('o') {
															goto l720
														}
														position++
														goto l719
													l720:
														position, tokenIndex, depth = position719, tokenIndex719, depth719
														if buffer[position] != rune('O') {
															goto l692
														}
														position++
													}
												l719:
													{
														position721, tokenIndex721, depth721 := position, tokenIndex, depth
														if buffer[position] != rune('n') {
															goto l722
														}
														position++
														goto l721
													l722:
														position, tokenIndex, depth = position721, tokenIndex721, depth721
														if buffer[position] != rune('N') {
															goto l692
														}
														position++
													}
												l721:
													{
														position723, tokenIndex723, depth723 := position, tokenIndex, depth
														if buffer[position] != rune('c') {
															goto l724
														}
														position++
														goto l723
													l724:
														position, tokenIndex, depth = position723, tokenIndex723, depth723
														if buffer[position] != rune('C') {
															goto l692
														}
														position++
													}
												l723:
													{
														position725, tokenIndex725, depth725 := position, tokenIndex, depth
														if buffer[position] != rune('a') {
															goto l726
														}
														position++
														goto l725
													l726:
														position, tokenIndex, depth = position725, tokenIndex725, depth725
														if buffer[position] != rune('A') {
															goto l692
														}
														position++
													}
												l725:
													{
														position727, tokenIndex727, depth727 := position, tokenIndex, depth
														if buffer[position] != rune('t') {
															goto l728
														}
														position++
														goto l727
													l728:
														position, tokenIndex, depth = position727, tokenIndex727, depth727
														if buffer[position] != rune('T') {
															goto l692
														}
														position++
													}
												l727:
													if !rules[ruleskip]() {
														goto l692
													}
													depth--
													add(ruleGROUPCONCAT, position706)
												}
												if !rules[ruleLPAREN]() {
													goto l692
												}
												{
													position729, tokenIndex729, depth729 := position, tokenIndex, depth
													if !rules[ruleDISTINCT]() {
														goto l729
													}
													goto l730
												l729:
													position, tokenIndex, depth = position729, tokenIndex729, depth729
												}
											l730:
												if !rules[ruleexpression]() {
													goto l692
												}
												{
													position731, tokenIndex731, depth731 := position, tokenIndex, depth
													if !rules[ruleSEMICOLON]() {
														goto l731
													}
													{
														position733 := position
														depth++
														{
															position734, tokenIndex734, depth734 := position, tokenIndex, depth
															if buffer[position] != rune('s') {
																goto l735
															}
															position++
															goto l734
														l735:
															position, tokenIndex, depth = position734, tokenIndex734, depth734
															if buffer[position] != rune('S') {
																goto l731
															}
															position++
														}
													l734:
														{
															position736, tokenIndex736, depth736 := position, tokenIndex, depth
															if buffer[position] != rune('e') {
																goto l737
															}
															position++
															goto l736
														l737:
															position, tokenIndex, depth = position736, tokenIndex736, depth736
															if buffer[position] != rune('E') {
																goto l731
															}
															position++
														}
													l736:
														{
															position738, tokenIndex738, depth738 := position, tokenIndex, depth
															if buffer[position] != rune('p') {
																goto l739
															}
															position++
															goto l738
														l739:
															position, tokenIndex, depth = position738, tokenIndex738, depth738
															if buffer[position] != rune('P') {
																goto l731
															}
															position++
														}
													l738:
														{
															position740, tokenIndex740, depth740 := position, tokenIndex, depth
															if buffer[position] != rune('a') {
																goto l741
															}
															position++
															goto l740
														l741:
															position, tokenIndex, depth = position740, tokenIndex740, depth740
															if buffer[position] != rune('A') {
																goto l731
															}
															position++
														}
													l740:
														{
															position742, tokenIndex742, depth742 := position, tokenIndex, depth
															if buffer[position] != rune('r') {
																goto l743
															}
															position++
															goto l742
														l743:
															position, tokenIndex, depth = position742, tokenIndex742, depth742
															if buffer[position] != rune('R') {
																goto l731
															}
															position++
														}
													l742:
														{
															position744, tokenIndex744, depth744 := position, tokenIndex, depth
															if buffer[position] != rune('a') {
																goto l745
															}
															position++
															goto l744
														l745:
															position, tokenIndex, depth = position744, tokenIndex744, depth744
															if buffer[position] != rune('A') {
																goto l731
															}
															position++
														}
													l744:
														{
															position746, tokenIndex746, depth746 := position, tokenIndex, depth
															if buffer[position] != rune('t') {
																goto l747
															}
															position++
															goto l746
														l747:
															position, tokenIndex, depth = position746, tokenIndex746, depth746
															if buffer[position] != rune('T') {
																goto l731
															}
															position++
														}
													l746:
														{
															position748, tokenIndex748, depth748 := position, tokenIndex, depth
															if buffer[position] != rune('o') {
																goto l749
															}
															position++
															goto l748
														l749:
															position, tokenIndex, depth = position748, tokenIndex748, depth748
															if buffer[position] != rune('O') {
																goto l731
															}
															position++
														}
													l748:
														{
															position750, tokenIndex750, depth750 := position, tokenIndex, depth
															if buffer[position] != rune('r') {
																goto l751
															}
															position++
															goto l750
														l751:
															position, tokenIndex, depth = position750, tokenIndex750, depth750
															if buffer[position] != rune('R') {
																goto l731
															}
															position++
														}
													l750:
														if !rules[ruleskip]() {
															goto l731
														}
														depth--
														add(ruleSEPARATOR, position733)
													}
													if !rules[ruleEQ]() {
														goto l731
													}
													if !rules[rulestring]() {
														goto l731
													}
													goto l732
												l731:
													position, tokenIndex, depth = position731, tokenIndex731, depth731
												}
											l732:
												if !rules[ruleRPAREN]() {
													goto l692
												}
												depth--
												add(rulegroupConcat, position705)
											}
											break
										case 'C', 'c':
											{
												position752 := position
												depth++
												{
													position753 := position
													depth++
													{
														position754, tokenIndex754, depth754 := position, tokenIndex, depth
														if buffer[position] != rune('c') {
															goto l755
														}
														position++
														goto l754
													l755:
														position, tokenIndex, depth = position754, tokenIndex754, depth754
														if buffer[position] != rune('C') {
															goto l692
														}
														position++
													}
												l754:
													{
														position756, tokenIndex756, depth756 := position, tokenIndex, depth
														if buffer[position] != rune('o') {
															goto l757
														}
														position++
														goto l756
													l757:
														position, tokenIndex, depth = position756, tokenIndex756, depth756
														if buffer[position] != rune('O') {
															goto l692
														}
														position++
													}
												l756:
													{
														position758, tokenIndex758, depth758 := position, tokenIndex, depth
														if buffer[position] != rune('u') {
															goto l759
														}
														position++
														goto l758
													l759:
														position, tokenIndex, depth = position758, tokenIndex758, depth758
														if buffer[position] != rune('U') {
															goto l692
														}
														position++
													}
												l758:
													{
														position760, tokenIndex760, depth760 := position, tokenIndex, depth
														if buffer[position] != rune('n') {
															goto l761
														}
														position++
														goto l760
													l761:
														position, tokenIndex, depth = position760, tokenIndex760, depth760
														if buffer[position] != rune('N') {
															goto l692
														}
														position++
													}
												l760:
													{
														position762, tokenIndex762, depth762 := position, tokenIndex, depth
														if buffer[position] != rune('t') {
															goto l763
														}
														position++
														goto l762
													l763:
														position, tokenIndex, depth = position762, tokenIndex762, depth762
														if buffer[position] != rune('T') {
															goto l692
														}
														position++
													}
												l762:
													if !rules[ruleskip]() {
														goto l692
													}
													depth--
													add(ruleCOUNT, position753)
												}
												if !rules[ruleLPAREN]() {
													goto l692
												}
												{
													position764, tokenIndex764, depth764 := position, tokenIndex, depth
													if !rules[ruleDISTINCT]() {
														goto l764
													}
													goto l765
												l764:
													position, tokenIndex, depth = position764, tokenIndex764, depth764
												}
											l765:
												{
													position766, tokenIndex766, depth766 := position, tokenIndex, depth
													if !rules[ruleSTAR]() {
														goto l767
													}
													goto l766
												l767:
													position, tokenIndex, depth = position766, tokenIndex766, depth766
													if !rules[ruleexpression]() {
														goto l692
													}
												}
											l766:
												if !rules[ruleRPAREN]() {
													goto l692
												}
												depth--
												add(rulecount, position752)
											}
											break
										default:
											{
												position768, tokenIndex768, depth768 := position, tokenIndex, depth
												{
													position770 := position
													depth++
													{
														position771, tokenIndex771, depth771 := position, tokenIndex, depth
														if buffer[position] != rune('s') {
															goto l772
														}
														position++
														goto l771
													l772:
														position, tokenIndex, depth = position771, tokenIndex771, depth771
														if buffer[position] != rune('S') {
															goto l769
														}
														position++
													}
												l771:
													{
														position773, tokenIndex773, depth773 := position, tokenIndex, depth
														if buffer[position] != rune('u') {
															goto l774
														}
														position++
														goto l773
													l774:
														position, tokenIndex, depth = position773, tokenIndex773, depth773
														if buffer[position] != rune('U') {
															goto l769
														}
														position++
													}
												l773:
													{
														position775, tokenIndex775, depth775 := position, tokenIndex, depth
														if buffer[position] != rune('m') {
															goto l776
														}
														position++
														goto l775
													l776:
														position, tokenIndex, depth = position775, tokenIndex775, depth775
														if buffer[position] != rune('M') {
															goto l769
														}
														position++
													}
												l775:
													if !rules[ruleskip]() {
														goto l769
													}
													depth--
													add(ruleSUM, position770)
												}
												goto l768
											l769:
												position, tokenIndex, depth = position768, tokenIndex768, depth768
												{
													position778 := position
													depth++
													{
														position779, tokenIndex779, depth779 := position, tokenIndex, depth
														if buffer[position] != rune('m') {
															goto l780
														}
														position++
														goto l779
													l780:
														position, tokenIndex, depth = position779, tokenIndex779, depth779
														if buffer[position] != rune('M') {
															goto l777
														}
														position++
													}
												l779:
													{
														position781, tokenIndex781, depth781 := position, tokenIndex, depth
														if buffer[position] != rune('i') {
															goto l782
														}
														position++
														goto l781
													l782:
														position, tokenIndex, depth = position781, tokenIndex781, depth781
														if buffer[position] != rune('I') {
															goto l777
														}
														position++
													}
												l781:
													{
														position783, tokenIndex783, depth783 := position, tokenIndex, depth
														if buffer[position] != rune('n') {
															goto l784
														}
														position++
														goto l783
													l784:
														position, tokenIndex, depth = position783, tokenIndex783, depth783
														if buffer[position] != rune('N') {
															goto l777
														}
														position++
													}
												l783:
													if !rules[ruleskip]() {
														goto l777
													}
													depth--
													add(ruleMIN, position778)
												}
												goto l768
											l777:
												position, tokenIndex, depth = position768, tokenIndex768, depth768
												{
													switch buffer[position] {
													case 'S', 's':
														{
															position786 := position
															depth++
															{
																position787, tokenIndex787, depth787 := position, tokenIndex, depth
																if buffer[position] != rune('s') {
																	goto l788
																}
																position++
																goto l787
															l788:
																position, tokenIndex, depth = position787, tokenIndex787, depth787
																if buffer[position] != rune('S') {
																	goto l692
																}
																position++
															}
														l787:
															{
																position789, tokenIndex789, depth789 := position, tokenIndex, depth
																if buffer[position] != rune('a') {
																	goto l790
																}
																position++
																goto l789
															l790:
																position, tokenIndex, depth = position789, tokenIndex789, depth789
																if buffer[position] != rune('A') {
																	goto l692
																}
																position++
															}
														l789:
															{
																position791, tokenIndex791, depth791 := position, tokenIndex, depth
																if buffer[position] != rune('m') {
																	goto l792
																}
																position++
																goto l791
															l792:
																position, tokenIndex, depth = position791, tokenIndex791, depth791
																if buffer[position] != rune('M') {
																	goto l692
																}
																position++
															}
														l791:
															{
																position793, tokenIndex793, depth793 := position, tokenIndex, depth
																if buffer[position] != rune('p') {
																	goto l794
																}
																position++
																goto l793
															l794:
																position, tokenIndex, depth = position793, tokenIndex793, depth793
																if buffer[position] != rune('P') {
																	goto l692
																}
																position++
															}
														l793:
															{
																position795, tokenIndex795, depth795 := position, tokenIndex, depth
																if buffer[position] != rune('l') {
																	goto l796
																}
																position++
																goto l795
															l796:
																position, tokenIndex, depth = position795, tokenIndex795, depth795
																if buffer[position] != rune('L') {
																	goto l692
																}
																position++
															}
														l795:
															{
																position797, tokenIndex797, depth797 := position, tokenIndex, depth
																if buffer[position] != rune('e') {
																	goto l798
																}
																position++
																goto l797
															l798:
																position, tokenIndex, depth = position797, tokenIndex797, depth797
																if buffer[position] != rune('E') {
																	goto l692
																}
																position++
															}
														l797:
															if !rules[ruleskip]() {
																goto l692
															}
															depth--
															add(ruleSAMPLE, position786)
														}
														break
													case 'A', 'a':
														{
															position799 := position
															depth++
															{
																position800, tokenIndex800, depth800 := position, tokenIndex, depth
																if buffer[position] != rune('a') {
																	goto l801
																}
																position++
																goto l800
															l801:
																position, tokenIndex, depth = position800, tokenIndex800, depth800
																if buffer[position] != rune('A') {
																	goto l692
																}
																position++
															}
														l800:
															{
																position802, tokenIndex802, depth802 := position, tokenIndex, depth
																if buffer[position] != rune('v') {
																	goto l803
																}
																position++
																goto l802
															l803:
																position, tokenIndex, depth = position802, tokenIndex802, depth802
																if buffer[position] != rune('V') {
																	goto l692
																}
																position++
															}
														l802:
															{
																position804, tokenIndex804, depth804 := position, tokenIndex, depth
																if buffer[position] != rune('g') {
																	goto l805
																}
																position++
																goto l804
															l805:
																position, tokenIndex, depth = position804, tokenIndex804, depth804
																if buffer[position] != rune('G') {
																	goto l692
																}
																position++
															}
														l804:
															if !rules[ruleskip]() {
																goto l692
															}
															depth--
															add(ruleAVG, position799)
														}
														break
													default:
														{
															position806 := position
															depth++
															{
																position807, tokenIndex807, depth807 := position, tokenIndex, depth
																if buffer[position] != rune('m') {
																	goto l808
																}
																position++
																goto l807
															l808:
																position, tokenIndex, depth = position807, tokenIndex807, depth807
																if buffer[position] != rune('M') {
																	goto l692
																}
																position++
															}
														l807:
															{
																position809, tokenIndex809, depth809 := position, tokenIndex, depth
																if buffer[position] != rune('a') {
																	goto l810
																}
																position++
																goto l809
															l810:
																position, tokenIndex, depth = position809, tokenIndex809, depth809
																if buffer[position] != rune('A') {
																	goto l692
																}
																position++
															}
														l809:
															{
																position811, tokenIndex811, depth811 := position, tokenIndex, depth
																if buffer[position] != rune('x') {
																	goto l812
																}
																position++
																goto l811
															l812:
																position, tokenIndex, depth = position811, tokenIndex811, depth811
																if buffer[position] != rune('X') {
																	goto l692
																}
																position++
															}
														l811:
															if !rules[ruleskip]() {
																goto l692
															}
															depth--
															add(ruleMAX, position806)
														}
														break
													}
												}

											}
										l768:
											if !rules[ruleLPAREN]() {
												goto l692
											}
											{
												position813, tokenIndex813, depth813 := position, tokenIndex, depth
												if !rules[ruleDISTINCT]() {
													goto l813
												}
												goto l814
											l813:
												position, tokenIndex, depth = position813, tokenIndex813, depth813
											}
										l814:
											if !rules[ruleexpression]() {
												goto l692
											}
											if !rules[ruleRPAREN]() {
												goto l692
											}
											break
										}
									}

									depth--
									add(ruleaggregate, position703)
								}
								break
							case '$', '?':
								if !rules[rulevar]() {
									goto l692
								}
								break
							case 'F', 'T', 'f', 't':
								if !rules[rulebooleanLiteral]() {
									goto l692
								}
								break
							case '(':
								if !rules[rulebrackettedExpression]() {
									goto l692
								}
								break
							case '"', '\'':
								if !rules[ruleliteral]() {
									goto l692
								}
								break
							default:
								if !rules[rulenumericLiteral]() {
									goto l692
								}
								break
							}
						}

					}
				l698:
					depth--
					add(ruleprimaryExpression, position697)
				}
				depth--
				add(ruleunaryExpression, position693)
			}
			return true
		l692:
			position, tokenIndex, depth = position692, tokenIndex692, depth692
			return false
		},
		/* 60 primaryExpression <- <(builtinCall / functionCall / iriref / ((&('A' | 'C' | 'G' | 'M' | 'S' | 'a' | 'c' | 'g' | 'm' | 's') aggregate) | (&('$' | '?') var) | (&('F' | 'T' | 'f' | 't') booleanLiteral) | (&('(') brackettedExpression) | (&('"' | '\'') literal) | (&('+' | '-' | '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') numericLiteral)))> */
		nil,
		/* 61 brackettedExpression <- <(LPAREN expression RPAREN)> */
		func() bool {
			position816, tokenIndex816, depth816 := position, tokenIndex, depth
			{
				position817 := position
				depth++
				if !rules[ruleLPAREN]() {
					goto l816
				}
				if !rules[ruleexpression]() {
					goto l816
				}
				if !rules[ruleRPAREN]() {
					goto l816
				}
				depth--
				add(rulebrackettedExpression, position817)
			}
			return true
		l816:
			position, tokenIndex, depth = position816, tokenIndex816, depth816
			return false
		},
		/* 62 functionCall <- <(iriref argList)> */
		func() bool {
			position818, tokenIndex818, depth818 := position, tokenIndex, depth
			{
				position819 := position
				depth++
				if !rules[ruleiriref]() {
					goto l818
				}
				if !rules[ruleargList]() {
					goto l818
				}
				depth--
				add(rulefunctionCall, position819)
			}
			return true
		l818:
			position, tokenIndex, depth = position818, tokenIndex818, depth818
			return false
		},
		/* 63 in <- <(IN argList)> */
		nil,
		/* 64 notin <- <(NOTIN argList)> */
		nil,
		/* 65 argList <- <(nil / (LPAREN expression (COMMA expression)* RPAREN))> */
		func() bool {
			position822, tokenIndex822, depth822 := position, tokenIndex, depth
			{
				position823 := position
				depth++
				{
					position824, tokenIndex824, depth824 := position, tokenIndex, depth
					if !rules[rulenil]() {
						goto l825
					}
					goto l824
				l825:
					position, tokenIndex, depth = position824, tokenIndex824, depth824
					if !rules[ruleLPAREN]() {
						goto l822
					}
					if !rules[ruleexpression]() {
						goto l822
					}
				l826:
					{
						position827, tokenIndex827, depth827 := position, tokenIndex, depth
						if !rules[ruleCOMMA]() {
							goto l827
						}
						if !rules[ruleexpression]() {
							goto l827
						}
						goto l826
					l827:
						position, tokenIndex, depth = position827, tokenIndex827, depth827
					}
					if !rules[ruleRPAREN]() {
						goto l822
					}
				}
			l824:
				depth--
				add(ruleargList, position823)
			}
			return true
		l822:
			position, tokenIndex, depth = position822, tokenIndex822, depth822
			return false
		},
		/* 66 aggregate <- <((&('G' | 'g') groupConcat) | (&('C' | 'c') count) | (&('A' | 'M' | 'S' | 'a' | 'm' | 's') ((SUM / MIN / ((&('S' | 's') SAMPLE) | (&('A' | 'a') AVG) | (&('M' | 'm') MAX))) LPAREN DISTINCT? expression RPAREN)))> */
		nil,
		/* 67 count <- <(COUNT LPAREN DISTINCT? (STAR / expression) RPAREN)> */
		nil,
		/* 68 groupConcat <- <(GROUPCONCAT LPAREN DISTINCT? expression (SEMICOLON SEPARATOR EQ string)? RPAREN)> */
		nil,
		/* 69 builtinCall <- <(((STR / LANG / DATATYPE / IRI / URI / STRLEN / MONTH / MINUTES / SECONDS / TIMEZONE / SHA1 / SHA256 / SHA384 / ISIRI / ISURI / ISBLANK / ISLITERAL / ((&('I' | 'i') ISNUMERIC) | (&('S' | 's') SHA512) | (&('M' | 'm') MD5) | (&('T' | 't') TZ) | (&('H' | 'h') HOURS) | (&('D' | 'd') DAY) | (&('Y' | 'y') YEAR) | (&('E' | 'e') ENCODEFORURI) | (&('L' | 'l') LCASE) | (&('U' | 'u') UCASE) | (&('F' | 'f') FLOOR) | (&('R' | 'r') ROUND) | (&('C' | 'c') CEIL) | (&('A' | 'a') ABS))) LPAREN expression RPAREN) / ((STRSTARTS / STRENDS / STRBEFORE / STRAFTER / STRLANG / STRDT / ((&('S' | 's') SAMETERM) | (&('C' | 'c') CONTAINS) | (&('L' | 'l') LANGMATCHES))) LPAREN expression COMMA expression RPAREN) / (BOUND LPAREN var RPAREN) / (((&('S' | 's') STRUUID) | (&('U' | 'u') UUID) | (&('N' | 'n') NOW) | (&('R' | 'r') RAND)) nil) / ((&('E' | 'N' | 'e' | 'n') ((EXISTS / NOTEXIST) groupGraphPattern)) | (&('I' | 'i') (IF LPAREN expression COMMA expression COMMA expression RPAREN)) | (&('C' | 'c') ((CONCAT / COALESCE) argList)) | (&('B' | 'b') (BNODE ((LPAREN expression RPAREN) / nil))) | (&('R' | 'S' | 'r' | 's') ((SUBSTR / REPLACE / REGEX) LPAREN expression COMMA expression (COMMA expression)? RPAREN))))> */
		func() bool {
			position831, tokenIndex831, depth831 := position, tokenIndex, depth
			{
				position832 := position
				depth++
				{
					position833, tokenIndex833, depth833 := position, tokenIndex, depth
					{
						position835, tokenIndex835, depth835 := position, tokenIndex, depth
						{
							position837 := position
							depth++
							{
								position838, tokenIndex838, depth838 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l839
								}
								position++
								goto l838
							l839:
								position, tokenIndex, depth = position838, tokenIndex838, depth838
								if buffer[position] != rune('S') {
									goto l836
								}
								position++
							}
						l838:
							{
								position840, tokenIndex840, depth840 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l841
								}
								position++
								goto l840
							l841:
								position, tokenIndex, depth = position840, tokenIndex840, depth840
								if buffer[position] != rune('T') {
									goto l836
								}
								position++
							}
						l840:
							{
								position842, tokenIndex842, depth842 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l843
								}
								position++
								goto l842
							l843:
								position, tokenIndex, depth = position842, tokenIndex842, depth842
								if buffer[position] != rune('R') {
									goto l836
								}
								position++
							}
						l842:
							if !rules[ruleskip]() {
								goto l836
							}
							depth--
							add(ruleSTR, position837)
						}
						goto l835
					l836:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position845 := position
							depth++
							{
								position846, tokenIndex846, depth846 := position, tokenIndex, depth
								if buffer[position] != rune('l') {
									goto l847
								}
								position++
								goto l846
							l847:
								position, tokenIndex, depth = position846, tokenIndex846, depth846
								if buffer[position] != rune('L') {
									goto l844
								}
								position++
							}
						l846:
							{
								position848, tokenIndex848, depth848 := position, tokenIndex, depth
								if buffer[position] != rune('a') {
									goto l849
								}
								position++
								goto l848
							l849:
								position, tokenIndex, depth = position848, tokenIndex848, depth848
								if buffer[position] != rune('A') {
									goto l844
								}
								position++
							}
						l848:
							{
								position850, tokenIndex850, depth850 := position, tokenIndex, depth
								if buffer[position] != rune('n') {
									goto l851
								}
								position++
								goto l850
							l851:
								position, tokenIndex, depth = position850, tokenIndex850, depth850
								if buffer[position] != rune('N') {
									goto l844
								}
								position++
							}
						l850:
							{
								position852, tokenIndex852, depth852 := position, tokenIndex, depth
								if buffer[position] != rune('g') {
									goto l853
								}
								position++
								goto l852
							l853:
								position, tokenIndex, depth = position852, tokenIndex852, depth852
								if buffer[position] != rune('G') {
									goto l844
								}
								position++
							}
						l852:
							if !rules[ruleskip]() {
								goto l844
							}
							depth--
							add(ruleLANG, position845)
						}
						goto l835
					l844:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position855 := position
							depth++
							{
								position856, tokenIndex856, depth856 := position, tokenIndex, depth
								if buffer[position] != rune('d') {
									goto l857
								}
								position++
								goto l856
							l857:
								position, tokenIndex, depth = position856, tokenIndex856, depth856
								if buffer[position] != rune('D') {
									goto l854
								}
								position++
							}
						l856:
							{
								position858, tokenIndex858, depth858 := position, tokenIndex, depth
								if buffer[position] != rune('a') {
									goto l859
								}
								position++
								goto l858
							l859:
								position, tokenIndex, depth = position858, tokenIndex858, depth858
								if buffer[position] != rune('A') {
									goto l854
								}
								position++
							}
						l858:
							{
								position860, tokenIndex860, depth860 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l861
								}
								position++
								goto l860
							l861:
								position, tokenIndex, depth = position860, tokenIndex860, depth860
								if buffer[position] != rune('T') {
									goto l854
								}
								position++
							}
						l860:
							{
								position862, tokenIndex862, depth862 := position, tokenIndex, depth
								if buffer[position] != rune('a') {
									goto l863
								}
								position++
								goto l862
							l863:
								position, tokenIndex, depth = position862, tokenIndex862, depth862
								if buffer[position] != rune('A') {
									goto l854
								}
								position++
							}
						l862:
							{
								position864, tokenIndex864, depth864 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l865
								}
								position++
								goto l864
							l865:
								position, tokenIndex, depth = position864, tokenIndex864, depth864
								if buffer[position] != rune('T') {
									goto l854
								}
								position++
							}
						l864:
							{
								position866, tokenIndex866, depth866 := position, tokenIndex, depth
								if buffer[position] != rune('y') {
									goto l867
								}
								position++
								goto l866
							l867:
								position, tokenIndex, depth = position866, tokenIndex866, depth866
								if buffer[position] != rune('Y') {
									goto l854
								}
								position++
							}
						l866:
							{
								position868, tokenIndex868, depth868 := position, tokenIndex, depth
								if buffer[position] != rune('p') {
									goto l869
								}
								position++
								goto l868
							l869:
								position, tokenIndex, depth = position868, tokenIndex868, depth868
								if buffer[position] != rune('P') {
									goto l854
								}
								position++
							}
						l868:
							{
								position870, tokenIndex870, depth870 := position, tokenIndex, depth
								if buffer[position] != rune('e') {
									goto l871
								}
								position++
								goto l870
							l871:
								position, tokenIndex, depth = position870, tokenIndex870, depth870
								if buffer[position] != rune('E') {
									goto l854
								}
								position++
							}
						l870:
							if !rules[ruleskip]() {
								goto l854
							}
							depth--
							add(ruleDATATYPE, position855)
						}
						goto l835
					l854:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position873 := position
							depth++
							{
								position874, tokenIndex874, depth874 := position, tokenIndex, depth
								if buffer[position] != rune('i') {
									goto l875
								}
								position++
								goto l874
							l875:
								position, tokenIndex, depth = position874, tokenIndex874, depth874
								if buffer[position] != rune('I') {
									goto l872
								}
								position++
							}
						l874:
							{
								position876, tokenIndex876, depth876 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l877
								}
								position++
								goto l876
							l877:
								position, tokenIndex, depth = position876, tokenIndex876, depth876
								if buffer[position] != rune('R') {
									goto l872
								}
								position++
							}
						l876:
							{
								position878, tokenIndex878, depth878 := position, tokenIndex, depth
								if buffer[position] != rune('i') {
									goto l879
								}
								position++
								goto l878
							l879:
								position, tokenIndex, depth = position878, tokenIndex878, depth878
								if buffer[position] != rune('I') {
									goto l872
								}
								position++
							}
						l878:
							if !rules[ruleskip]() {
								goto l872
							}
							depth--
							add(ruleIRI, position873)
						}
						goto l835
					l872:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position881 := position
							depth++
							{
								position882, tokenIndex882, depth882 := position, tokenIndex, depth
								if buffer[position] != rune('u') {
									goto l883
								}
								position++
								goto l882
							l883:
								position, tokenIndex, depth = position882, tokenIndex882, depth882
								if buffer[position] != rune('U') {
									goto l880
								}
								position++
							}
						l882:
							{
								position884, tokenIndex884, depth884 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l885
								}
								position++
								goto l884
							l885:
								position, tokenIndex, depth = position884, tokenIndex884, depth884
								if buffer[position] != rune('R') {
									goto l880
								}
								position++
							}
						l884:
							{
								position886, tokenIndex886, depth886 := position, tokenIndex, depth
								if buffer[position] != rune('i') {
									goto l887
								}
								position++
								goto l886
							l887:
								position, tokenIndex, depth = position886, tokenIndex886, depth886
								if buffer[position] != rune('I') {
									goto l880
								}
								position++
							}
						l886:
							if !rules[ruleskip]() {
								goto l880
							}
							depth--
							add(ruleURI, position881)
						}
						goto l835
					l880:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position889 := position
							depth++
							{
								position890, tokenIndex890, depth890 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l891
								}
								position++
								goto l890
							l891:
								position, tokenIndex, depth = position890, tokenIndex890, depth890
								if buffer[position] != rune('S') {
									goto l888
								}
								position++
							}
						l890:
							{
								position892, tokenIndex892, depth892 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l893
								}
								position++
								goto l892
							l893:
								position, tokenIndex, depth = position892, tokenIndex892, depth892
								if buffer[position] != rune('T') {
									goto l888
								}
								position++
							}
						l892:
							{
								position894, tokenIndex894, depth894 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l895
								}
								position++
								goto l894
							l895:
								position, tokenIndex, depth = position894, tokenIndex894, depth894
								if buffer[position] != rune('R') {
									goto l888
								}
								position++
							}
						l894:
							{
								position896, tokenIndex896, depth896 := position, tokenIndex, depth
								if buffer[position] != rune('l') {
									goto l897
								}
								position++
								goto l896
							l897:
								position, tokenIndex, depth = position896, tokenIndex896, depth896
								if buffer[position] != rune('L') {
									goto l888
								}
								position++
							}
						l896:
							{
								position898, tokenIndex898, depth898 := position, tokenIndex, depth
								if buffer[position] != rune('e') {
									goto l899
								}
								position++
								goto l898
							l899:
								position, tokenIndex, depth = position898, tokenIndex898, depth898
								if buffer[position] != rune('E') {
									goto l888
								}
								position++
							}
						l898:
							{
								position900, tokenIndex900, depth900 := position, tokenIndex, depth
								if buffer[position] != rune('n') {
									goto l901
								}
								position++
								goto l900
							l901:
								position, tokenIndex, depth = position900, tokenIndex900, depth900
								if buffer[position] != rune('N') {
									goto l888
								}
								position++
							}
						l900:
							if !rules[ruleskip]() {
								goto l888
							}
							depth--
							add(ruleSTRLEN, position889)
						}
						goto l835
					l888:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position903 := position
							depth++
							{
								position904, tokenIndex904, depth904 := position, tokenIndex, depth
								if buffer[position] != rune('m') {
									goto l905
								}
								position++
								goto l904
							l905:
								position, tokenIndex, depth = position904, tokenIndex904, depth904
								if buffer[position] != rune('M') {
									goto l902
								}
								position++
							}
						l904:
							{
								position906, tokenIndex906, depth906 := position, tokenIndex, depth
								if buffer[position] != rune('o') {
									goto l907
								}
								position++
								goto l906
							l907:
								position, tokenIndex, depth = position906, tokenIndex906, depth906
								if buffer[position] != rune('O') {
									goto l902
								}
								position++
							}
						l906:
							{
								position908, tokenIndex908, depth908 := position, tokenIndex, depth
								if buffer[position] != rune('n') {
									goto l909
								}
								position++
								goto l908
							l909:
								position, tokenIndex, depth = position908, tokenIndex908, depth908
								if buffer[position] != rune('N') {
									goto l902
								}
								position++
							}
						l908:
							{
								position910, tokenIndex910, depth910 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l911
								}
								position++
								goto l910
							l911:
								position, tokenIndex, depth = position910, tokenIndex910, depth910
								if buffer[position] != rune('T') {
									goto l902
								}
								position++
							}
						l910:
							{
								position912, tokenIndex912, depth912 := position, tokenIndex, depth
								if buffer[position] != rune('h') {
									goto l913
								}
								position++
								goto l912
							l913:
								position, tokenIndex, depth = position912, tokenIndex912, depth912
								if buffer[position] != rune('H') {
									goto l902
								}
								position++
							}
						l912:
							if !rules[ruleskip]() {
								goto l902
							}
							depth--
							add(ruleMONTH, position903)
						}
						goto l835
					l902:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position915 := position
							depth++
							{
								position916, tokenIndex916, depth916 := position, tokenIndex, depth
								if buffer[position] != rune('m') {
									goto l917
								}
								position++
								goto l916
							l917:
								position, tokenIndex, depth = position916, tokenIndex916, depth916
								if buffer[position] != rune('M') {
									goto l914
								}
								position++
							}
						l916:
							{
								position918, tokenIndex918, depth918 := position, tokenIndex, depth
								if buffer[position] != rune('i') {
									goto l919
								}
								position++
								goto l918
							l919:
								position, tokenIndex, depth = position918, tokenIndex918, depth918
								if buffer[position] != rune('I') {
									goto l914
								}
								position++
							}
						l918:
							{
								position920, tokenIndex920, depth920 := position, tokenIndex, depth
								if buffer[position] != rune('n') {
									goto l921
								}
								position++
								goto l920
							l921:
								position, tokenIndex, depth = position920, tokenIndex920, depth920
								if buffer[position] != rune('N') {
									goto l914
								}
								position++
							}
						l920:
							{
								position922, tokenIndex922, depth922 := position, tokenIndex, depth
								if buffer[position] != rune('u') {
									goto l923
								}
								position++
								goto l922
							l923:
								position, tokenIndex, depth = position922, tokenIndex922, depth922
								if buffer[position] != rune('U') {
									goto l914
								}
								position++
							}
						l922:
							{
								position924, tokenIndex924, depth924 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l925
								}
								position++
								goto l924
							l925:
								position, tokenIndex, depth = position924, tokenIndex924, depth924
								if buffer[position] != rune('T') {
									goto l914
								}
								position++
							}
						l924:
							{
								position926, tokenIndex926, depth926 := position, tokenIndex, depth
								if buffer[position] != rune('e') {
									goto l927
								}
								position++
								goto l926
							l927:
								position, tokenIndex, depth = position926, tokenIndex926, depth926
								if buffer[position] != rune('E') {
									goto l914
								}
								position++
							}
						l926:
							{
								position928, tokenIndex928, depth928 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l929
								}
								position++
								goto l928
							l929:
								position, tokenIndex, depth = position928, tokenIndex928, depth928
								if buffer[position] != rune('S') {
									goto l914
								}
								position++
							}
						l928:
							if !rules[ruleskip]() {
								goto l914
							}
							depth--
							add(ruleMINUTES, position915)
						}
						goto l835
					l914:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position931 := position
							depth++
							{
								position932, tokenIndex932, depth932 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l933
								}
								position++
								goto l932
							l933:
								position, tokenIndex, depth = position932, tokenIndex932, depth932
								if buffer[position] != rune('S') {
									goto l930
								}
								position++
							}
						l932:
							{
								position934, tokenIndex934, depth934 := position, tokenIndex, depth
								if buffer[position] != rune('e') {
									goto l935
								}
								position++
								goto l934
							l935:
								position, tokenIndex, depth = position934, tokenIndex934, depth934
								if buffer[position] != rune('E') {
									goto l930
								}
								position++
							}
						l934:
							{
								position936, tokenIndex936, depth936 := position, tokenIndex, depth
								if buffer[position] != rune('c') {
									goto l937
								}
								position++
								goto l936
							l937:
								position, tokenIndex, depth = position936, tokenIndex936, depth936
								if buffer[position] != rune('C') {
									goto l930
								}
								position++
							}
						l936:
							{
								position938, tokenIndex938, depth938 := position, tokenIndex, depth
								if buffer[position] != rune('o') {
									goto l939
								}
								position++
								goto l938
							l939:
								position, tokenIndex, depth = position938, tokenIndex938, depth938
								if buffer[position] != rune('O') {
									goto l930
								}
								position++
							}
						l938:
							{
								position940, tokenIndex940, depth940 := position, tokenIndex, depth
								if buffer[position] != rune('n') {
									goto l941
								}
								position++
								goto l940
							l941:
								position, tokenIndex, depth = position940, tokenIndex940, depth940
								if buffer[position] != rune('N') {
									goto l930
								}
								position++
							}
						l940:
							{
								position942, tokenIndex942, depth942 := position, tokenIndex, depth
								if buffer[position] != rune('d') {
									goto l943
								}
								position++
								goto l942
							l943:
								position, tokenIndex, depth = position942, tokenIndex942, depth942
								if buffer[position] != rune('D') {
									goto l930
								}
								position++
							}
						l942:
							{
								position944, tokenIndex944, depth944 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l945
								}
								position++
								goto l944
							l945:
								position, tokenIndex, depth = position944, tokenIndex944, depth944
								if buffer[position] != rune('S') {
									goto l930
								}
								position++
							}
						l944:
							if !rules[ruleskip]() {
								goto l930
							}
							depth--
							add(ruleSECONDS, position931)
						}
						goto l835
					l930:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position947 := position
							depth++
							{
								position948, tokenIndex948, depth948 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l949
								}
								position++
								goto l948
							l949:
								position, tokenIndex, depth = position948, tokenIndex948, depth948
								if buffer[position] != rune('T') {
									goto l946
								}
								position++
							}
						l948:
							{
								position950, tokenIndex950, depth950 := position, tokenIndex, depth
								if buffer[position] != rune('i') {
									goto l951
								}
								position++
								goto l950
							l951:
								position, tokenIndex, depth = position950, tokenIndex950, depth950
								if buffer[position] != rune('I') {
									goto l946
								}
								position++
							}
						l950:
							{
								position952, tokenIndex952, depth952 := position, tokenIndex, depth
								if buffer[position] != rune('m') {
									goto l953
								}
								position++
								goto l952
							l953:
								position, tokenIndex, depth = position952, tokenIndex952, depth952
								if buffer[position] != rune('M') {
									goto l946
								}
								position++
							}
						l952:
							{
								position954, tokenIndex954, depth954 := position, tokenIndex, depth
								if buffer[position] != rune('e') {
									goto l955
								}
								position++
								goto l954
							l955:
								position, tokenIndex, depth = position954, tokenIndex954, depth954
								if buffer[position] != rune('E') {
									goto l946
								}
								position++
							}
						l954:
							{
								position956, tokenIndex956, depth956 := position, tokenIndex, depth
								if buffer[position] != rune('z') {
									goto l957
								}
								position++
								goto l956
							l957:
								position, tokenIndex, depth = position956, tokenIndex956, depth956
								if buffer[position] != rune('Z') {
									goto l946
								}
								position++
							}
						l956:
							{
								position958, tokenIndex958, depth958 := position, tokenIndex, depth
								if buffer[position] != rune('o') {
									goto l959
								}
								position++
								goto l958
							l959:
								position, tokenIndex, depth = position958, tokenIndex958, depth958
								if buffer[position] != rune('O') {
									goto l946
								}
								position++
							}
						l958:
							{
								position960, tokenIndex960, depth960 := position, tokenIndex, depth
								if buffer[position] != rune('n') {
									goto l961
								}
								position++
								goto l960
							l961:
								position, tokenIndex, depth = position960, tokenIndex960, depth960
								if buffer[position] != rune('N') {
									goto l946
								}
								position++
							}
						l960:
							{
								position962, tokenIndex962, depth962 := position, tokenIndex, depth
								if buffer[position] != rune('e') {
									goto l963
								}
								position++
								goto l962
							l963:
								position, tokenIndex, depth = position962, tokenIndex962, depth962
								if buffer[position] != rune('E') {
									goto l946
								}
								position++
							}
						l962:
							if !rules[ruleskip]() {
								goto l946
							}
							depth--
							add(ruleTIMEZONE, position947)
						}
						goto l835
					l946:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position965 := position
							depth++
							{
								position966, tokenIndex966, depth966 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l967
								}
								position++
								goto l966
							l967:
								position, tokenIndex, depth = position966, tokenIndex966, depth966
								if buffer[position] != rune('S') {
									goto l964
								}
								position++
							}
						l966:
							{
								position968, tokenIndex968, depth968 := position, tokenIndex, depth
								if buffer[position] != rune('h') {
									goto l969
								}
								position++
								goto l968
							l969:
								position, tokenIndex, depth = position968, tokenIndex968, depth968
								if buffer[position] != rune('H') {
									goto l964
								}
								position++
							}
						l968:
							{
								position970, tokenIndex970, depth970 := position, tokenIndex, depth
								if buffer[position] != rune('a') {
									goto l971
								}
								position++
								goto l970
							l971:
								position, tokenIndex, depth = position970, tokenIndex970, depth970
								if buffer[position] != rune('A') {
									goto l964
								}
								position++
							}
						l970:
							if buffer[position] != rune('1') {
								goto l964
							}
							position++
							if !rules[ruleskip]() {
								goto l964
							}
							depth--
							add(ruleSHA1, position965)
						}
						goto l835
					l964:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position973 := position
							depth++
							{
								position974, tokenIndex974, depth974 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l975
								}
								position++
								goto l974
							l975:
								position, tokenIndex, depth = position974, tokenIndex974, depth974
								if buffer[position] != rune('S') {
									goto l972
								}
								position++
							}
						l974:
							{
								position976, tokenIndex976, depth976 := position, tokenIndex, depth
								if buffer[position] != rune('h') {
									goto l977
								}
								position++
								goto l976
							l977:
								position, tokenIndex, depth = position976, tokenIndex976, depth976
								if buffer[position] != rune('H') {
									goto l972
								}
								position++
							}
						l976:
							{
								position978, tokenIndex978, depth978 := position, tokenIndex, depth
								if buffer[position] != rune('a') {
									goto l979
								}
								position++
								goto l978
							l979:
								position, tokenIndex, depth = position978, tokenIndex978, depth978
								if buffer[position] != rune('A') {
									goto l972
								}
								position++
							}
						l978:
							if buffer[position] != rune('2') {
								goto l972
							}
							position++
							if buffer[position] != rune('5') {
								goto l972
							}
							position++
							if buffer[position] != rune('6') {
								goto l972
							}
							position++
							if !rules[ruleskip]() {
								goto l972
							}
							depth--
							add(ruleSHA256, position973)
						}
						goto l835
					l972:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position981 := position
							depth++
							{
								position982, tokenIndex982, depth982 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l983
								}
								position++
								goto l982
							l983:
								position, tokenIndex, depth = position982, tokenIndex982, depth982
								if buffer[position] != rune('S') {
									goto l980
								}
								position++
							}
						l982:
							{
								position984, tokenIndex984, depth984 := position, tokenIndex, depth
								if buffer[position] != rune('h') {
									goto l985
								}
								position++
								goto l984
							l985:
								position, tokenIndex, depth = position984, tokenIndex984, depth984
								if buffer[position] != rune('H') {
									goto l980
								}
								position++
							}
						l984:
							{
								position986, tokenIndex986, depth986 := position, tokenIndex, depth
								if buffer[position] != rune('a') {
									goto l987
								}
								position++
								goto l986
							l987:
								position, tokenIndex, depth = position986, tokenIndex986, depth986
								if buffer[position] != rune('A') {
									goto l980
								}
								position++
							}
						l986:
							if buffer[position] != rune('3') {
								goto l980
							}
							position++
							if buffer[position] != rune('8') {
								goto l980
							}
							position++
							if buffer[position] != rune('4') {
								goto l980
							}
							position++
							if !rules[ruleskip]() {
								goto l980
							}
							depth--
							add(ruleSHA384, position981)
						}
						goto l835
					l980:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position989 := position
							depth++
							{
								position990, tokenIndex990, depth990 := position, tokenIndex, depth
								if buffer[position] != rune('i') {
									goto l991
								}
								position++
								goto l990
							l991:
								position, tokenIndex, depth = position990, tokenIndex990, depth990
								if buffer[position] != rune('I') {
									goto l988
								}
								position++
							}
						l990:
							{
								position992, tokenIndex992, depth992 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l993
								}
								position++
								goto l992
							l993:
								position, tokenIndex, depth = position992, tokenIndex992, depth992
								if buffer[position] != rune('S') {
									goto l988
								}
								position++
							}
						l992:
							{
								position994, tokenIndex994, depth994 := position, tokenIndex, depth
								if buffer[position] != rune('i') {
									goto l995
								}
								position++
								goto l994
							l995:
								position, tokenIndex, depth = position994, tokenIndex994, depth994
								if buffer[position] != rune('I') {
									goto l988
								}
								position++
							}
						l994:
							{
								position996, tokenIndex996, depth996 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l997
								}
								position++
								goto l996
							l997:
								position, tokenIndex, depth = position996, tokenIndex996, depth996
								if buffer[position] != rune('R') {
									goto l988
								}
								position++
							}
						l996:
							{
								position998, tokenIndex998, depth998 := position, tokenIndex, depth
								if buffer[position] != rune('i') {
									goto l999
								}
								position++
								goto l998
							l999:
								position, tokenIndex, depth = position998, tokenIndex998, depth998
								if buffer[position] != rune('I') {
									goto l988
								}
								position++
							}
						l998:
							if !rules[ruleskip]() {
								goto l988
							}
							depth--
							add(ruleISIRI, position989)
						}
						goto l835
					l988:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position1001 := position
							depth++
							{
								position1002, tokenIndex1002, depth1002 := position, tokenIndex, depth
								if buffer[position] != rune('i') {
									goto l1003
								}
								position++
								goto l1002
							l1003:
								position, tokenIndex, depth = position1002, tokenIndex1002, depth1002
								if buffer[position] != rune('I') {
									goto l1000
								}
								position++
							}
						l1002:
							{
								position1004, tokenIndex1004, depth1004 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l1005
								}
								position++
								goto l1004
							l1005:
								position, tokenIndex, depth = position1004, tokenIndex1004, depth1004
								if buffer[position] != rune('S') {
									goto l1000
								}
								position++
							}
						l1004:
							{
								position1006, tokenIndex1006, depth1006 := position, tokenIndex, depth
								if buffer[position] != rune('u') {
									goto l1007
								}
								position++
								goto l1006
							l1007:
								position, tokenIndex, depth = position1006, tokenIndex1006, depth1006
								if buffer[position] != rune('U') {
									goto l1000
								}
								position++
							}
						l1006:
							{
								position1008, tokenIndex1008, depth1008 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l1009
								}
								position++
								goto l1008
							l1009:
								position, tokenIndex, depth = position1008, tokenIndex1008, depth1008
								if buffer[position] != rune('R') {
									goto l1000
								}
								position++
							}
						l1008:
							{
								position1010, tokenIndex1010, depth1010 := position, tokenIndex, depth
								if buffer[position] != rune('i') {
									goto l1011
								}
								position++
								goto l1010
							l1011:
								position, tokenIndex, depth = position1010, tokenIndex1010, depth1010
								if buffer[position] != rune('I') {
									goto l1000
								}
								position++
							}
						l1010:
							if !rules[ruleskip]() {
								goto l1000
							}
							depth--
							add(ruleISURI, position1001)
						}
						goto l835
					l1000:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position1013 := position
							depth++
							{
								position1014, tokenIndex1014, depth1014 := position, tokenIndex, depth
								if buffer[position] != rune('i') {
									goto l1015
								}
								position++
								goto l1014
							l1015:
								position, tokenIndex, depth = position1014, tokenIndex1014, depth1014
								if buffer[position] != rune('I') {
									goto l1012
								}
								position++
							}
						l1014:
							{
								position1016, tokenIndex1016, depth1016 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l1017
								}
								position++
								goto l1016
							l1017:
								position, tokenIndex, depth = position1016, tokenIndex1016, depth1016
								if buffer[position] != rune('S') {
									goto l1012
								}
								position++
							}
						l1016:
							{
								position1018, tokenIndex1018, depth1018 := position, tokenIndex, depth
								if buffer[position] != rune('b') {
									goto l1019
								}
								position++
								goto l1018
							l1019:
								position, tokenIndex, depth = position1018, tokenIndex1018, depth1018
								if buffer[position] != rune('B') {
									goto l1012
								}
								position++
							}
						l1018:
							{
								position1020, tokenIndex1020, depth1020 := position, tokenIndex, depth
								if buffer[position] != rune('l') {
									goto l1021
								}
								position++
								goto l1020
							l1021:
								position, tokenIndex, depth = position1020, tokenIndex1020, depth1020
								if buffer[position] != rune('L') {
									goto l1012
								}
								position++
							}
						l1020:
							{
								position1022, tokenIndex1022, depth1022 := position, tokenIndex, depth
								if buffer[position] != rune('a') {
									goto l1023
								}
								position++
								goto l1022
							l1023:
								position, tokenIndex, depth = position1022, tokenIndex1022, depth1022
								if buffer[position] != rune('A') {
									goto l1012
								}
								position++
							}
						l1022:
							{
								position1024, tokenIndex1024, depth1024 := position, tokenIndex, depth
								if buffer[position] != rune('n') {
									goto l1025
								}
								position++
								goto l1024
							l1025:
								position, tokenIndex, depth = position1024, tokenIndex1024, depth1024
								if buffer[position] != rune('N') {
									goto l1012
								}
								position++
							}
						l1024:
							{
								position1026, tokenIndex1026, depth1026 := position, tokenIndex, depth
								if buffer[position] != rune('k') {
									goto l1027
								}
								position++
								goto l1026
							l1027:
								position, tokenIndex, depth = position1026, tokenIndex1026, depth1026
								if buffer[position] != rune('K') {
									goto l1012
								}
								position++
							}
						l1026:
							if !rules[ruleskip]() {
								goto l1012
							}
							depth--
							add(ruleISBLANK, position1013)
						}
						goto l835
					l1012:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							position1029 := position
							depth++
							{
								position1030, tokenIndex1030, depth1030 := position, tokenIndex, depth
								if buffer[position] != rune('i') {
									goto l1031
								}
								position++
								goto l1030
							l1031:
								position, tokenIndex, depth = position1030, tokenIndex1030, depth1030
								if buffer[position] != rune('I') {
									goto l1028
								}
								position++
							}
						l1030:
							{
								position1032, tokenIndex1032, depth1032 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l1033
								}
								position++
								goto l1032
							l1033:
								position, tokenIndex, depth = position1032, tokenIndex1032, depth1032
								if buffer[position] != rune('S') {
									goto l1028
								}
								position++
							}
						l1032:
							{
								position1034, tokenIndex1034, depth1034 := position, tokenIndex, depth
								if buffer[position] != rune('l') {
									goto l1035
								}
								position++
								goto l1034
							l1035:
								position, tokenIndex, depth = position1034, tokenIndex1034, depth1034
								if buffer[position] != rune('L') {
									goto l1028
								}
								position++
							}
						l1034:
							{
								position1036, tokenIndex1036, depth1036 := position, tokenIndex, depth
								if buffer[position] != rune('i') {
									goto l1037
								}
								position++
								goto l1036
							l1037:
								position, tokenIndex, depth = position1036, tokenIndex1036, depth1036
								if buffer[position] != rune('I') {
									goto l1028
								}
								position++
							}
						l1036:
							{
								position1038, tokenIndex1038, depth1038 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l1039
								}
								position++
								goto l1038
							l1039:
								position, tokenIndex, depth = position1038, tokenIndex1038, depth1038
								if buffer[position] != rune('T') {
									goto l1028
								}
								position++
							}
						l1038:
							{
								position1040, tokenIndex1040, depth1040 := position, tokenIndex, depth
								if buffer[position] != rune('e') {
									goto l1041
								}
								position++
								goto l1040
							l1041:
								position, tokenIndex, depth = position1040, tokenIndex1040, depth1040
								if buffer[position] != rune('E') {
									goto l1028
								}
								position++
							}
						l1040:
							{
								position1042, tokenIndex1042, depth1042 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l1043
								}
								position++
								goto l1042
							l1043:
								position, tokenIndex, depth = position1042, tokenIndex1042, depth1042
								if buffer[position] != rune('R') {
									goto l1028
								}
								position++
							}
						l1042:
							{
								position1044, tokenIndex1044, depth1044 := position, tokenIndex, depth
								if buffer[position] != rune('a') {
									goto l1045
								}
								position++
								goto l1044
							l1045:
								position, tokenIndex, depth = position1044, tokenIndex1044, depth1044
								if buffer[position] != rune('A') {
									goto l1028
								}
								position++
							}
						l1044:
							{
								position1046, tokenIndex1046, depth1046 := position, tokenIndex, depth
								if buffer[position] != rune('l') {
									goto l1047
								}
								position++
								goto l1046
							l1047:
								position, tokenIndex, depth = position1046, tokenIndex1046, depth1046
								if buffer[position] != rune('L') {
									goto l1028
								}
								position++
							}
						l1046:
							if !rules[ruleskip]() {
								goto l1028
							}
							depth--
							add(ruleISLITERAL, position1029)
						}
						goto l835
					l1028:
						position, tokenIndex, depth = position835, tokenIndex835, depth835
						{
							switch buffer[position] {
							case 'I', 'i':
								{
									position1049 := position
									depth++
									{
										position1050, tokenIndex1050, depth1050 := position, tokenIndex, depth
										if buffer[position] != rune('i') {
											goto l1051
										}
										position++
										goto l1050
									l1051:
										position, tokenIndex, depth = position1050, tokenIndex1050, depth1050
										if buffer[position] != rune('I') {
											goto l834
										}
										position++
									}
								l1050:
									{
										position1052, tokenIndex1052, depth1052 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1053
										}
										position++
										goto l1052
									l1053:
										position, tokenIndex, depth = position1052, tokenIndex1052, depth1052
										if buffer[position] != rune('S') {
											goto l834
										}
										position++
									}
								l1052:
									{
										position1054, tokenIndex1054, depth1054 := position, tokenIndex, depth
										if buffer[position] != rune('n') {
											goto l1055
										}
										position++
										goto l1054
									l1055:
										position, tokenIndex, depth = position1054, tokenIndex1054, depth1054
										if buffer[position] != rune('N') {
											goto l834
										}
										position++
									}
								l1054:
									{
										position1056, tokenIndex1056, depth1056 := position, tokenIndex, depth
										if buffer[position] != rune('u') {
											goto l1057
										}
										position++
										goto l1056
									l1057:
										position, tokenIndex, depth = position1056, tokenIndex1056, depth1056
										if buffer[position] != rune('U') {
											goto l834
										}
										position++
									}
								l1056:
									{
										position1058, tokenIndex1058, depth1058 := position, tokenIndex, depth
										if buffer[position] != rune('m') {
											goto l1059
										}
										position++
										goto l1058
									l1059:
										position, tokenIndex, depth = position1058, tokenIndex1058, depth1058
										if buffer[position] != rune('M') {
											goto l834
										}
										position++
									}
								l1058:
									{
										position1060, tokenIndex1060, depth1060 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1061
										}
										position++
										goto l1060
									l1061:
										position, tokenIndex, depth = position1060, tokenIndex1060, depth1060
										if buffer[position] != rune('E') {
											goto l834
										}
										position++
									}
								l1060:
									{
										position1062, tokenIndex1062, depth1062 := position, tokenIndex, depth
										if buffer[position] != rune('r') {
											goto l1063
										}
										position++
										goto l1062
									l1063:
										position, tokenIndex, depth = position1062, tokenIndex1062, depth1062
										if buffer[position] != rune('R') {
											goto l834
										}
										position++
									}
								l1062:
									{
										position1064, tokenIndex1064, depth1064 := position, tokenIndex, depth
										if buffer[position] != rune('i') {
											goto l1065
										}
										position++
										goto l1064
									l1065:
										position, tokenIndex, depth = position1064, tokenIndex1064, depth1064
										if buffer[position] != rune('I') {
											goto l834
										}
										position++
									}
								l1064:
									{
										position1066, tokenIndex1066, depth1066 := position, tokenIndex, depth
										if buffer[position] != rune('c') {
											goto l1067
										}
										position++
										goto l1066
									l1067:
										position, tokenIndex, depth = position1066, tokenIndex1066, depth1066
										if buffer[position] != rune('C') {
											goto l834
										}
										position++
									}
								l1066:
									if !rules[ruleskip]() {
										goto l834
									}
									depth--
									add(ruleISNUMERIC, position1049)
								}
								break
							case 'S', 's':
								{
									position1068 := position
									depth++
									{
										position1069, tokenIndex1069, depth1069 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1070
										}
										position++
										goto l1069
									l1070:
										position, tokenIndex, depth = position1069, tokenIndex1069, depth1069
										if buffer[position] != rune('S') {
											goto l834
										}
										position++
									}
								l1069:
									{
										position1071, tokenIndex1071, depth1071 := position, tokenIndex, depth
										if buffer[position] != rune('h') {
											goto l1072
										}
										position++
										goto l1071
									l1072:
										position, tokenIndex, depth = position1071, tokenIndex1071, depth1071
										if buffer[position] != rune('H') {
											goto l834
										}
										position++
									}
								l1071:
									{
										position1073, tokenIndex1073, depth1073 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l1074
										}
										position++
										goto l1073
									l1074:
										position, tokenIndex, depth = position1073, tokenIndex1073, depth1073
										if buffer[position] != rune('A') {
											goto l834
										}
										position++
									}
								l1073:
									if buffer[position] != rune('5') {
										goto l834
									}
									position++
									if buffer[position] != rune('1') {
										goto l834
									}
									position++
									if buffer[position] != rune('2') {
										goto l834
									}
									position++
									if !rules[ruleskip]() {
										goto l834
									}
									depth--
									add(ruleSHA512, position1068)
								}
								break
							case 'M', 'm':
								{
									position1075 := position
									depth++
									{
										position1076, tokenIndex1076, depth1076 := position, tokenIndex, depth
										if buffer[position] != rune('m') {
											goto l1077
										}
										position++
										goto l1076
									l1077:
										position, tokenIndex, depth = position1076, tokenIndex1076, depth1076
										if buffer[position] != rune('M') {
											goto l834
										}
										position++
									}
								l1076:
									{
										position1078, tokenIndex1078, depth1078 := position, tokenIndex, depth
										if buffer[position] != rune('d') {
											goto l1079
										}
										position++
										goto l1078
									l1079:
										position, tokenIndex, depth = position1078, tokenIndex1078, depth1078
										if buffer[position] != rune('D') {
											goto l834
										}
										position++
									}
								l1078:
									if !rules[ruleskip]() {
										goto l834
									}
									depth--
									add(ruleMD5, position1075)
								}
								break
							case 'T', 't':
								{
									position1080 := position
									depth++
									{
										position1081, tokenIndex1081, depth1081 := position, tokenIndex, depth
										if buffer[position] != rune('t') {
											goto l1082
										}
										position++
										goto l1081
									l1082:
										position, tokenIndex, depth = position1081, tokenIndex1081, depth1081
										if buffer[position] != rune('T') {
											goto l834
										}
										position++
									}
								l1081:
									{
										position1083, tokenIndex1083, depth1083 := position, tokenIndex, depth
										if buffer[position] != rune('z') {
											goto l1084
										}
										position++
										goto l1083
									l1084:
										position, tokenIndex, depth = position1083, tokenIndex1083, depth1083
										if buffer[position] != rune('Z') {
											goto l834
										}
										position++
									}
								l1083:
									if !rules[ruleskip]() {
										goto l834
									}
									depth--
									add(ruleTZ, position1080)
								}
								break
							case 'H', 'h':
								{
									position1085 := position
									depth++
									{
										position1086, tokenIndex1086, depth1086 := position, tokenIndex, depth
										if buffer[position] != rune('h') {
											goto l1087
										}
										position++
										goto l1086
									l1087:
										position, tokenIndex, depth = position1086, tokenIndex1086, depth1086
										if buffer[position] != rune('H') {
											goto l834
										}
										position++
									}
								l1086:
									{
										position1088, tokenIndex1088, depth1088 := position, tokenIndex, depth
										if buffer[position] != rune('o') {
											goto l1089
										}
										position++
										goto l1088
									l1089:
										position, tokenIndex, depth = position1088, tokenIndex1088, depth1088
										if buffer[position] != rune('O') {
											goto l834
										}
										position++
									}
								l1088:
									{
										position1090, tokenIndex1090, depth1090 := position, tokenIndex, depth
										if buffer[position] != rune('u') {
											goto l1091
										}
										position++
										goto l1090
									l1091:
										position, tokenIndex, depth = position1090, tokenIndex1090, depth1090
										if buffer[position] != rune('U') {
											goto l834
										}
										position++
									}
								l1090:
									{
										position1092, tokenIndex1092, depth1092 := position, tokenIndex, depth
										if buffer[position] != rune('r') {
											goto l1093
										}
										position++
										goto l1092
									l1093:
										position, tokenIndex, depth = position1092, tokenIndex1092, depth1092
										if buffer[position] != rune('R') {
											goto l834
										}
										position++
									}
								l1092:
									{
										position1094, tokenIndex1094, depth1094 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1095
										}
										position++
										goto l1094
									l1095:
										position, tokenIndex, depth = position1094, tokenIndex1094, depth1094
										if buffer[position] != rune('S') {
											goto l834
										}
										position++
									}
								l1094:
									if !rules[ruleskip]() {
										goto l834
									}
									depth--
									add(ruleHOURS, position1085)
								}
								break
							case 'D', 'd':
								{
									position1096 := position
									depth++
									{
										position1097, tokenIndex1097, depth1097 := position, tokenIndex, depth
										if buffer[position] != rune('d') {
											goto l1098
										}
										position++
										goto l1097
									l1098:
										position, tokenIndex, depth = position1097, tokenIndex1097, depth1097
										if buffer[position] != rune('D') {
											goto l834
										}
										position++
									}
								l1097:
									{
										position1099, tokenIndex1099, depth1099 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l1100
										}
										position++
										goto l1099
									l1100:
										position, tokenIndex, depth = position1099, tokenIndex1099, depth1099
										if buffer[position] != rune('A') {
											goto l834
										}
										position++
									}
								l1099:
									{
										position1101, tokenIndex1101, depth1101 := position, tokenIndex, depth
										if buffer[position] != rune('y') {
											goto l1102
										}
										position++
										goto l1101
									l1102:
										position, tokenIndex, depth = position1101, tokenIndex1101, depth1101
										if buffer[position] != rune('Y') {
											goto l834
										}
										position++
									}
								l1101:
									if !rules[ruleskip]() {
										goto l834
									}
									depth--
									add(ruleDAY, position1096)
								}
								break
							case 'Y', 'y':
								{
									position1103 := position
									depth++
									{
										position1104, tokenIndex1104, depth1104 := position, tokenIndex, depth
										if buffer[position] != rune('y') {
											goto l1105
										}
										position++
										goto l1104
									l1105:
										position, tokenIndex, depth = position1104, tokenIndex1104, depth1104
										if buffer[position] != rune('Y') {
											goto l834
										}
										position++
									}
								l1104:
									{
										position1106, tokenIndex1106, depth1106 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1107
										}
										position++
										goto l1106
									l1107:
										position, tokenIndex, depth = position1106, tokenIndex1106, depth1106
										if buffer[position] != rune('E') {
											goto l834
										}
										position++
									}
								l1106:
									{
										position1108, tokenIndex1108, depth1108 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l1109
										}
										position++
										goto l1108
									l1109:
										position, tokenIndex, depth = position1108, tokenIndex1108, depth1108
										if buffer[position] != rune('A') {
											goto l834
										}
										position++
									}
								l1108:
									{
										position1110, tokenIndex1110, depth1110 := position, tokenIndex, depth
										if buffer[position] != rune('r') {
											goto l1111
										}
										position++
										goto l1110
									l1111:
										position, tokenIndex, depth = position1110, tokenIndex1110, depth1110
										if buffer[position] != rune('R') {
											goto l834
										}
										position++
									}
								l1110:
									if !rules[ruleskip]() {
										goto l834
									}
									depth--
									add(ruleYEAR, position1103)
								}
								break
							case 'E', 'e':
								{
									position1112 := position
									depth++
									{
										position1113, tokenIndex1113, depth1113 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1114
										}
										position++
										goto l1113
									l1114:
										position, tokenIndex, depth = position1113, tokenIndex1113, depth1113
										if buffer[position] != rune('E') {
											goto l834
										}
										position++
									}
								l1113:
									{
										position1115, tokenIndex1115, depth1115 := position, tokenIndex, depth
										if buffer[position] != rune('n') {
											goto l1116
										}
										position++
										goto l1115
									l1116:
										position, tokenIndex, depth = position1115, tokenIndex1115, depth1115
										if buffer[position] != rune('N') {
											goto l834
										}
										position++
									}
								l1115:
									{
										position1117, tokenIndex1117, depth1117 := position, tokenIndex, depth
										if buffer[position] != rune('c') {
											goto l1118
										}
										position++
										goto l1117
									l1118:
										position, tokenIndex, depth = position1117, tokenIndex1117, depth1117
										if buffer[position] != rune('C') {
											goto l834
										}
										position++
									}
								l1117:
									{
										position1119, tokenIndex1119, depth1119 := position, tokenIndex, depth
										if buffer[position] != rune('o') {
											goto l1120
										}
										position++
										goto l1119
									l1120:
										position, tokenIndex, depth = position1119, tokenIndex1119, depth1119
										if buffer[position] != rune('O') {
											goto l834
										}
										position++
									}
								l1119:
									{
										position1121, tokenIndex1121, depth1121 := position, tokenIndex, depth
										if buffer[position] != rune('d') {
											goto l1122
										}
										position++
										goto l1121
									l1122:
										position, tokenIndex, depth = position1121, tokenIndex1121, depth1121
										if buffer[position] != rune('D') {
											goto l834
										}
										position++
									}
								l1121:
									{
										position1123, tokenIndex1123, depth1123 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1124
										}
										position++
										goto l1123
									l1124:
										position, tokenIndex, depth = position1123, tokenIndex1123, depth1123
										if buffer[position] != rune('E') {
											goto l834
										}
										position++
									}
								l1123:
									if buffer[position] != rune('_') {
										goto l834
									}
									position++
									{
										position1125, tokenIndex1125, depth1125 := position, tokenIndex, depth
										if buffer[position] != rune('f') {
											goto l1126
										}
										position++
										goto l1125
									l1126:
										position, tokenIndex, depth = position1125, tokenIndex1125, depth1125
										if buffer[position] != rune('F') {
											goto l834
										}
										position++
									}
								l1125:
									{
										position1127, tokenIndex1127, depth1127 := position, tokenIndex, depth
										if buffer[position] != rune('o') {
											goto l1128
										}
										position++
										goto l1127
									l1128:
										position, tokenIndex, depth = position1127, tokenIndex1127, depth1127
										if buffer[position] != rune('O') {
											goto l834
										}
										position++
									}
								l1127:
									{
										position1129, tokenIndex1129, depth1129 := position, tokenIndex, depth
										if buffer[position] != rune('r') {
											goto l1130
										}
										position++
										goto l1129
									l1130:
										position, tokenIndex, depth = position1129, tokenIndex1129, depth1129
										if buffer[position] != rune('R') {
											goto l834
										}
										position++
									}
								l1129:
									if buffer[position] != rune('_') {
										goto l834
									}
									position++
									{
										position1131, tokenIndex1131, depth1131 := position, tokenIndex, depth
										if buffer[position] != rune('u') {
											goto l1132
										}
										position++
										goto l1131
									l1132:
										position, tokenIndex, depth = position1131, tokenIndex1131, depth1131
										if buffer[position] != rune('U') {
											goto l834
										}
										position++
									}
								l1131:
									{
										position1133, tokenIndex1133, depth1133 := position, tokenIndex, depth
										if buffer[position] != rune('r') {
											goto l1134
										}
										position++
										goto l1133
									l1134:
										position, tokenIndex, depth = position1133, tokenIndex1133, depth1133
										if buffer[position] != rune('R') {
											goto l834
										}
										position++
									}
								l1133:
									{
										position1135, tokenIndex1135, depth1135 := position, tokenIndex, depth
										if buffer[position] != rune('i') {
											goto l1136
										}
										position++
										goto l1135
									l1136:
										position, tokenIndex, depth = position1135, tokenIndex1135, depth1135
										if buffer[position] != rune('I') {
											goto l834
										}
										position++
									}
								l1135:
									if !rules[ruleskip]() {
										goto l834
									}
									depth--
									add(ruleENCODEFORURI, position1112)
								}
								break
							case 'L', 'l':
								{
									position1137 := position
									depth++
									{
										position1138, tokenIndex1138, depth1138 := position, tokenIndex, depth
										if buffer[position] != rune('l') {
											goto l1139
										}
										position++
										goto l1138
									l1139:
										position, tokenIndex, depth = position1138, tokenIndex1138, depth1138
										if buffer[position] != rune('L') {
											goto l834
										}
										position++
									}
								l1138:
									{
										position1140, tokenIndex1140, depth1140 := position, tokenIndex, depth
										if buffer[position] != rune('c') {
											goto l1141
										}
										position++
										goto l1140
									l1141:
										position, tokenIndex, depth = position1140, tokenIndex1140, depth1140
										if buffer[position] != rune('C') {
											goto l834
										}
										position++
									}
								l1140:
									{
										position1142, tokenIndex1142, depth1142 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l1143
										}
										position++
										goto l1142
									l1143:
										position, tokenIndex, depth = position1142, tokenIndex1142, depth1142
										if buffer[position] != rune('A') {
											goto l834
										}
										position++
									}
								l1142:
									{
										position1144, tokenIndex1144, depth1144 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1145
										}
										position++
										goto l1144
									l1145:
										position, tokenIndex, depth = position1144, tokenIndex1144, depth1144
										if buffer[position] != rune('S') {
											goto l834
										}
										position++
									}
								l1144:
									{
										position1146, tokenIndex1146, depth1146 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1147
										}
										position++
										goto l1146
									l1147:
										position, tokenIndex, depth = position1146, tokenIndex1146, depth1146
										if buffer[position] != rune('E') {
											goto l834
										}
										position++
									}
								l1146:
									if !rules[ruleskip]() {
										goto l834
									}
									depth--
									add(ruleLCASE, position1137)
								}
								break
							case 'U', 'u':
								{
									position1148 := position
									depth++
									{
										position1149, tokenIndex1149, depth1149 := position, tokenIndex, depth
										if buffer[position] != rune('u') {
											goto l1150
										}
										position++
										goto l1149
									l1150:
										position, tokenIndex, depth = position1149, tokenIndex1149, depth1149
										if buffer[position] != rune('U') {
											goto l834
										}
										position++
									}
								l1149:
									{
										position1151, tokenIndex1151, depth1151 := position, tokenIndex, depth
										if buffer[position] != rune('c') {
											goto l1152
										}
										position++
										goto l1151
									l1152:
										position, tokenIndex, depth = position1151, tokenIndex1151, depth1151
										if buffer[position] != rune('C') {
											goto l834
										}
										position++
									}
								l1151:
									{
										position1153, tokenIndex1153, depth1153 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l1154
										}
										position++
										goto l1153
									l1154:
										position, tokenIndex, depth = position1153, tokenIndex1153, depth1153
										if buffer[position] != rune('A') {
											goto l834
										}
										position++
									}
								l1153:
									{
										position1155, tokenIndex1155, depth1155 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1156
										}
										position++
										goto l1155
									l1156:
										position, tokenIndex, depth = position1155, tokenIndex1155, depth1155
										if buffer[position] != rune('S') {
											goto l834
										}
										position++
									}
								l1155:
									{
										position1157, tokenIndex1157, depth1157 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1158
										}
										position++
										goto l1157
									l1158:
										position, tokenIndex, depth = position1157, tokenIndex1157, depth1157
										if buffer[position] != rune('E') {
											goto l834
										}
										position++
									}
								l1157:
									if !rules[ruleskip]() {
										goto l834
									}
									depth--
									add(ruleUCASE, position1148)
								}
								break
							case 'F', 'f':
								{
									position1159 := position
									depth++
									{
										position1160, tokenIndex1160, depth1160 := position, tokenIndex, depth
										if buffer[position] != rune('f') {
											goto l1161
										}
										position++
										goto l1160
									l1161:
										position, tokenIndex, depth = position1160, tokenIndex1160, depth1160
										if buffer[position] != rune('F') {
											goto l834
										}
										position++
									}
								l1160:
									{
										position1162, tokenIndex1162, depth1162 := position, tokenIndex, depth
										if buffer[position] != rune('l') {
											goto l1163
										}
										position++
										goto l1162
									l1163:
										position, tokenIndex, depth = position1162, tokenIndex1162, depth1162
										if buffer[position] != rune('L') {
											goto l834
										}
										position++
									}
								l1162:
									{
										position1164, tokenIndex1164, depth1164 := position, tokenIndex, depth
										if buffer[position] != rune('o') {
											goto l1165
										}
										position++
										goto l1164
									l1165:
										position, tokenIndex, depth = position1164, tokenIndex1164, depth1164
										if buffer[position] != rune('O') {
											goto l834
										}
										position++
									}
								l1164:
									{
										position1166, tokenIndex1166, depth1166 := position, tokenIndex, depth
										if buffer[position] != rune('o') {
											goto l1167
										}
										position++
										goto l1166
									l1167:
										position, tokenIndex, depth = position1166, tokenIndex1166, depth1166
										if buffer[position] != rune('O') {
											goto l834
										}
										position++
									}
								l1166:
									{
										position1168, tokenIndex1168, depth1168 := position, tokenIndex, depth
										if buffer[position] != rune('r') {
											goto l1169
										}
										position++
										goto l1168
									l1169:
										position, tokenIndex, depth = position1168, tokenIndex1168, depth1168
										if buffer[position] != rune('R') {
											goto l834
										}
										position++
									}
								l1168:
									if !rules[ruleskip]() {
										goto l834
									}
									depth--
									add(ruleFLOOR, position1159)
								}
								break
							case 'R', 'r':
								{
									position1170 := position
									depth++
									{
										position1171, tokenIndex1171, depth1171 := position, tokenIndex, depth
										if buffer[position] != rune('r') {
											goto l1172
										}
										position++
										goto l1171
									l1172:
										position, tokenIndex, depth = position1171, tokenIndex1171, depth1171
										if buffer[position] != rune('R') {
											goto l834
										}
										position++
									}
								l1171:
									{
										position1173, tokenIndex1173, depth1173 := position, tokenIndex, depth
										if buffer[position] != rune('o') {
											goto l1174
										}
										position++
										goto l1173
									l1174:
										position, tokenIndex, depth = position1173, tokenIndex1173, depth1173
										if buffer[position] != rune('O') {
											goto l834
										}
										position++
									}
								l1173:
									{
										position1175, tokenIndex1175, depth1175 := position, tokenIndex, depth
										if buffer[position] != rune('u') {
											goto l1176
										}
										position++
										goto l1175
									l1176:
										position, tokenIndex, depth = position1175, tokenIndex1175, depth1175
										if buffer[position] != rune('U') {
											goto l834
										}
										position++
									}
								l1175:
									{
										position1177, tokenIndex1177, depth1177 := position, tokenIndex, depth
										if buffer[position] != rune('n') {
											goto l1178
										}
										position++
										goto l1177
									l1178:
										position, tokenIndex, depth = position1177, tokenIndex1177, depth1177
										if buffer[position] != rune('N') {
											goto l834
										}
										position++
									}
								l1177:
									{
										position1179, tokenIndex1179, depth1179 := position, tokenIndex, depth
										if buffer[position] != rune('d') {
											goto l1180
										}
										position++
										goto l1179
									l1180:
										position, tokenIndex, depth = position1179, tokenIndex1179, depth1179
										if buffer[position] != rune('D') {
											goto l834
										}
										position++
									}
								l1179:
									if !rules[ruleskip]() {
										goto l834
									}
									depth--
									add(ruleROUND, position1170)
								}
								break
							case 'C', 'c':
								{
									position1181 := position
									depth++
									{
										position1182, tokenIndex1182, depth1182 := position, tokenIndex, depth
										if buffer[position] != rune('c') {
											goto l1183
										}
										position++
										goto l1182
									l1183:
										position, tokenIndex, depth = position1182, tokenIndex1182, depth1182
										if buffer[position] != rune('C') {
											goto l834
										}
										position++
									}
								l1182:
									{
										position1184, tokenIndex1184, depth1184 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1185
										}
										position++
										goto l1184
									l1185:
										position, tokenIndex, depth = position1184, tokenIndex1184, depth1184
										if buffer[position] != rune('E') {
											goto l834
										}
										position++
									}
								l1184:
									{
										position1186, tokenIndex1186, depth1186 := position, tokenIndex, depth
										if buffer[position] != rune('i') {
											goto l1187
										}
										position++
										goto l1186
									l1187:
										position, tokenIndex, depth = position1186, tokenIndex1186, depth1186
										if buffer[position] != rune('I') {
											goto l834
										}
										position++
									}
								l1186:
									{
										position1188, tokenIndex1188, depth1188 := position, tokenIndex, depth
										if buffer[position] != rune('l') {
											goto l1189
										}
										position++
										goto l1188
									l1189:
										position, tokenIndex, depth = position1188, tokenIndex1188, depth1188
										if buffer[position] != rune('L') {
											goto l834
										}
										position++
									}
								l1188:
									if !rules[ruleskip]() {
										goto l834
									}
									depth--
									add(ruleCEIL, position1181)
								}
								break
							default:
								{
									position1190 := position
									depth++
									{
										position1191, tokenIndex1191, depth1191 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l1192
										}
										position++
										goto l1191
									l1192:
										position, tokenIndex, depth = position1191, tokenIndex1191, depth1191
										if buffer[position] != rune('A') {
											goto l834
										}
										position++
									}
								l1191:
									{
										position1193, tokenIndex1193, depth1193 := position, tokenIndex, depth
										if buffer[position] != rune('b') {
											goto l1194
										}
										position++
										goto l1193
									l1194:
										position, tokenIndex, depth = position1193, tokenIndex1193, depth1193
										if buffer[position] != rune('B') {
											goto l834
										}
										position++
									}
								l1193:
									{
										position1195, tokenIndex1195, depth1195 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1196
										}
										position++
										goto l1195
									l1196:
										position, tokenIndex, depth = position1195, tokenIndex1195, depth1195
										if buffer[position] != rune('S') {
											goto l834
										}
										position++
									}
								l1195:
									if !rules[ruleskip]() {
										goto l834
									}
									depth--
									add(ruleABS, position1190)
								}
								break
							}
						}

					}
				l835:
					if !rules[ruleLPAREN]() {
						goto l834
					}
					if !rules[ruleexpression]() {
						goto l834
					}
					if !rules[ruleRPAREN]() {
						goto l834
					}
					goto l833
				l834:
					position, tokenIndex, depth = position833, tokenIndex833, depth833
					{
						position1198, tokenIndex1198, depth1198 := position, tokenIndex, depth
						{
							position1200 := position
							depth++
							{
								position1201, tokenIndex1201, depth1201 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l1202
								}
								position++
								goto l1201
							l1202:
								position, tokenIndex, depth = position1201, tokenIndex1201, depth1201
								if buffer[position] != rune('S') {
									goto l1199
								}
								position++
							}
						l1201:
							{
								position1203, tokenIndex1203, depth1203 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l1204
								}
								position++
								goto l1203
							l1204:
								position, tokenIndex, depth = position1203, tokenIndex1203, depth1203
								if buffer[position] != rune('T') {
									goto l1199
								}
								position++
							}
						l1203:
							{
								position1205, tokenIndex1205, depth1205 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l1206
								}
								position++
								goto l1205
							l1206:
								position, tokenIndex, depth = position1205, tokenIndex1205, depth1205
								if buffer[position] != rune('R') {
									goto l1199
								}
								position++
							}
						l1205:
							{
								position1207, tokenIndex1207, depth1207 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l1208
								}
								position++
								goto l1207
							l1208:
								position, tokenIndex, depth = position1207, tokenIndex1207, depth1207
								if buffer[position] != rune('S') {
									goto l1199
								}
								position++
							}
						l1207:
							{
								position1209, tokenIndex1209, depth1209 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l1210
								}
								position++
								goto l1209
							l1210:
								position, tokenIndex, depth = position1209, tokenIndex1209, depth1209
								if buffer[position] != rune('T') {
									goto l1199
								}
								position++
							}
						l1209:
							{
								position1211, tokenIndex1211, depth1211 := position, tokenIndex, depth
								if buffer[position] != rune('a') {
									goto l1212
								}
								position++
								goto l1211
							l1212:
								position, tokenIndex, depth = position1211, tokenIndex1211, depth1211
								if buffer[position] != rune('A') {
									goto l1199
								}
								position++
							}
						l1211:
							{
								position1213, tokenIndex1213, depth1213 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l1214
								}
								position++
								goto l1213
							l1214:
								position, tokenIndex, depth = position1213, tokenIndex1213, depth1213
								if buffer[position] != rune('R') {
									goto l1199
								}
								position++
							}
						l1213:
							{
								position1215, tokenIndex1215, depth1215 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l1216
								}
								position++
								goto l1215
							l1216:
								position, tokenIndex, depth = position1215, tokenIndex1215, depth1215
								if buffer[position] != rune('T') {
									goto l1199
								}
								position++
							}
						l1215:
							{
								position1217, tokenIndex1217, depth1217 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l1218
								}
								position++
								goto l1217
							l1218:
								position, tokenIndex, depth = position1217, tokenIndex1217, depth1217
								if buffer[position] != rune('S') {
									goto l1199
								}
								position++
							}
						l1217:
							if !rules[ruleskip]() {
								goto l1199
							}
							depth--
							add(ruleSTRSTARTS, position1200)
						}
						goto l1198
					l1199:
						position, tokenIndex, depth = position1198, tokenIndex1198, depth1198
						{
							position1220 := position
							depth++
							{
								position1221, tokenIndex1221, depth1221 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l1222
								}
								position++
								goto l1221
							l1222:
								position, tokenIndex, depth = position1221, tokenIndex1221, depth1221
								if buffer[position] != rune('S') {
									goto l1219
								}
								position++
							}
						l1221:
							{
								position1223, tokenIndex1223, depth1223 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l1224
								}
								position++
								goto l1223
							l1224:
								position, tokenIndex, depth = position1223, tokenIndex1223, depth1223
								if buffer[position] != rune('T') {
									goto l1219
								}
								position++
							}
						l1223:
							{
								position1225, tokenIndex1225, depth1225 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l1226
								}
								position++
								goto l1225
							l1226:
								position, tokenIndex, depth = position1225, tokenIndex1225, depth1225
								if buffer[position] != rune('R') {
									goto l1219
								}
								position++
							}
						l1225:
							{
								position1227, tokenIndex1227, depth1227 := position, tokenIndex, depth
								if buffer[position] != rune('e') {
									goto l1228
								}
								position++
								goto l1227
							l1228:
								position, tokenIndex, depth = position1227, tokenIndex1227, depth1227
								if buffer[position] != rune('E') {
									goto l1219
								}
								position++
							}
						l1227:
							{
								position1229, tokenIndex1229, depth1229 := position, tokenIndex, depth
								if buffer[position] != rune('n') {
									goto l1230
								}
								position++
								goto l1229
							l1230:
								position, tokenIndex, depth = position1229, tokenIndex1229, depth1229
								if buffer[position] != rune('N') {
									goto l1219
								}
								position++
							}
						l1229:
							{
								position1231, tokenIndex1231, depth1231 := position, tokenIndex, depth
								if buffer[position] != rune('d') {
									goto l1232
								}
								position++
								goto l1231
							l1232:
								position, tokenIndex, depth = position1231, tokenIndex1231, depth1231
								if buffer[position] != rune('D') {
									goto l1219
								}
								position++
							}
						l1231:
							{
								position1233, tokenIndex1233, depth1233 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l1234
								}
								position++
								goto l1233
							l1234:
								position, tokenIndex, depth = position1233, tokenIndex1233, depth1233
								if buffer[position] != rune('S') {
									goto l1219
								}
								position++
							}
						l1233:
							if !rules[ruleskip]() {
								goto l1219
							}
							depth--
							add(ruleSTRENDS, position1220)
						}
						goto l1198
					l1219:
						position, tokenIndex, depth = position1198, tokenIndex1198, depth1198
						{
							position1236 := position
							depth++
							{
								position1237, tokenIndex1237, depth1237 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l1238
								}
								position++
								goto l1237
							l1238:
								position, tokenIndex, depth = position1237, tokenIndex1237, depth1237
								if buffer[position] != rune('S') {
									goto l1235
								}
								position++
							}
						l1237:
							{
								position1239, tokenIndex1239, depth1239 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l1240
								}
								position++
								goto l1239
							l1240:
								position, tokenIndex, depth = position1239, tokenIndex1239, depth1239
								if buffer[position] != rune('T') {
									goto l1235
								}
								position++
							}
						l1239:
							{
								position1241, tokenIndex1241, depth1241 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l1242
								}
								position++
								goto l1241
							l1242:
								position, tokenIndex, depth = position1241, tokenIndex1241, depth1241
								if buffer[position] != rune('R') {
									goto l1235
								}
								position++
							}
						l1241:
							{
								position1243, tokenIndex1243, depth1243 := position, tokenIndex, depth
								if buffer[position] != rune('b') {
									goto l1244
								}
								position++
								goto l1243
							l1244:
								position, tokenIndex, depth = position1243, tokenIndex1243, depth1243
								if buffer[position] != rune('B') {
									goto l1235
								}
								position++
							}
						l1243:
							{
								position1245, tokenIndex1245, depth1245 := position, tokenIndex, depth
								if buffer[position] != rune('e') {
									goto l1246
								}
								position++
								goto l1245
							l1246:
								position, tokenIndex, depth = position1245, tokenIndex1245, depth1245
								if buffer[position] != rune('E') {
									goto l1235
								}
								position++
							}
						l1245:
							{
								position1247, tokenIndex1247, depth1247 := position, tokenIndex, depth
								if buffer[position] != rune('f') {
									goto l1248
								}
								position++
								goto l1247
							l1248:
								position, tokenIndex, depth = position1247, tokenIndex1247, depth1247
								if buffer[position] != rune('F') {
									goto l1235
								}
								position++
							}
						l1247:
							{
								position1249, tokenIndex1249, depth1249 := position, tokenIndex, depth
								if buffer[position] != rune('o') {
									goto l1250
								}
								position++
								goto l1249
							l1250:
								position, tokenIndex, depth = position1249, tokenIndex1249, depth1249
								if buffer[position] != rune('O') {
									goto l1235
								}
								position++
							}
						l1249:
							{
								position1251, tokenIndex1251, depth1251 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l1252
								}
								position++
								goto l1251
							l1252:
								position, tokenIndex, depth = position1251, tokenIndex1251, depth1251
								if buffer[position] != rune('R') {
									goto l1235
								}
								position++
							}
						l1251:
							{
								position1253, tokenIndex1253, depth1253 := position, tokenIndex, depth
								if buffer[position] != rune('e') {
									goto l1254
								}
								position++
								goto l1253
							l1254:
								position, tokenIndex, depth = position1253, tokenIndex1253, depth1253
								if buffer[position] != rune('E') {
									goto l1235
								}
								position++
							}
						l1253:
							if !rules[ruleskip]() {
								goto l1235
							}
							depth--
							add(ruleSTRBEFORE, position1236)
						}
						goto l1198
					l1235:
						position, tokenIndex, depth = position1198, tokenIndex1198, depth1198
						{
							position1256 := position
							depth++
							{
								position1257, tokenIndex1257, depth1257 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l1258
								}
								position++
								goto l1257
							l1258:
								position, tokenIndex, depth = position1257, tokenIndex1257, depth1257
								if buffer[position] != rune('S') {
									goto l1255
								}
								position++
							}
						l1257:
							{
								position1259, tokenIndex1259, depth1259 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l1260
								}
								position++
								goto l1259
							l1260:
								position, tokenIndex, depth = position1259, tokenIndex1259, depth1259
								if buffer[position] != rune('T') {
									goto l1255
								}
								position++
							}
						l1259:
							{
								position1261, tokenIndex1261, depth1261 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l1262
								}
								position++
								goto l1261
							l1262:
								position, tokenIndex, depth = position1261, tokenIndex1261, depth1261
								if buffer[position] != rune('R') {
									goto l1255
								}
								position++
							}
						l1261:
							{
								position1263, tokenIndex1263, depth1263 := position, tokenIndex, depth
								if buffer[position] != rune('a') {
									goto l1264
								}
								position++
								goto l1263
							l1264:
								position, tokenIndex, depth = position1263, tokenIndex1263, depth1263
								if buffer[position] != rune('A') {
									goto l1255
								}
								position++
							}
						l1263:
							{
								position1265, tokenIndex1265, depth1265 := position, tokenIndex, depth
								if buffer[position] != rune('f') {
									goto l1266
								}
								position++
								goto l1265
							l1266:
								position, tokenIndex, depth = position1265, tokenIndex1265, depth1265
								if buffer[position] != rune('F') {
									goto l1255
								}
								position++
							}
						l1265:
							{
								position1267, tokenIndex1267, depth1267 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l1268
								}
								position++
								goto l1267
							l1268:
								position, tokenIndex, depth = position1267, tokenIndex1267, depth1267
								if buffer[position] != rune('T') {
									goto l1255
								}
								position++
							}
						l1267:
							{
								position1269, tokenIndex1269, depth1269 := position, tokenIndex, depth
								if buffer[position] != rune('e') {
									goto l1270
								}
								position++
								goto l1269
							l1270:
								position, tokenIndex, depth = position1269, tokenIndex1269, depth1269
								if buffer[position] != rune('E') {
									goto l1255
								}
								position++
							}
						l1269:
							{
								position1271, tokenIndex1271, depth1271 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l1272
								}
								position++
								goto l1271
							l1272:
								position, tokenIndex, depth = position1271, tokenIndex1271, depth1271
								if buffer[position] != rune('R') {
									goto l1255
								}
								position++
							}
						l1271:
							if !rules[ruleskip]() {
								goto l1255
							}
							depth--
							add(ruleSTRAFTER, position1256)
						}
						goto l1198
					l1255:
						position, tokenIndex, depth = position1198, tokenIndex1198, depth1198
						{
							position1274 := position
							depth++
							{
								position1275, tokenIndex1275, depth1275 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l1276
								}
								position++
								goto l1275
							l1276:
								position, tokenIndex, depth = position1275, tokenIndex1275, depth1275
								if buffer[position] != rune('S') {
									goto l1273
								}
								position++
							}
						l1275:
							{
								position1277, tokenIndex1277, depth1277 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l1278
								}
								position++
								goto l1277
							l1278:
								position, tokenIndex, depth = position1277, tokenIndex1277, depth1277
								if buffer[position] != rune('T') {
									goto l1273
								}
								position++
							}
						l1277:
							{
								position1279, tokenIndex1279, depth1279 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l1280
								}
								position++
								goto l1279
							l1280:
								position, tokenIndex, depth = position1279, tokenIndex1279, depth1279
								if buffer[position] != rune('R') {
									goto l1273
								}
								position++
							}
						l1279:
							{
								position1281, tokenIndex1281, depth1281 := position, tokenIndex, depth
								if buffer[position] != rune('l') {
									goto l1282
								}
								position++
								goto l1281
							l1282:
								position, tokenIndex, depth = position1281, tokenIndex1281, depth1281
								if buffer[position] != rune('L') {
									goto l1273
								}
								position++
							}
						l1281:
							{
								position1283, tokenIndex1283, depth1283 := position, tokenIndex, depth
								if buffer[position] != rune('a') {
									goto l1284
								}
								position++
								goto l1283
							l1284:
								position, tokenIndex, depth = position1283, tokenIndex1283, depth1283
								if buffer[position] != rune('A') {
									goto l1273
								}
								position++
							}
						l1283:
							{
								position1285, tokenIndex1285, depth1285 := position, tokenIndex, depth
								if buffer[position] != rune('n') {
									goto l1286
								}
								position++
								goto l1285
							l1286:
								position, tokenIndex, depth = position1285, tokenIndex1285, depth1285
								if buffer[position] != rune('N') {
									goto l1273
								}
								position++
							}
						l1285:
							{
								position1287, tokenIndex1287, depth1287 := position, tokenIndex, depth
								if buffer[position] != rune('g') {
									goto l1288
								}
								position++
								goto l1287
							l1288:
								position, tokenIndex, depth = position1287, tokenIndex1287, depth1287
								if buffer[position] != rune('G') {
									goto l1273
								}
								position++
							}
						l1287:
							if !rules[ruleskip]() {
								goto l1273
							}
							depth--
							add(ruleSTRLANG, position1274)
						}
						goto l1198
					l1273:
						position, tokenIndex, depth = position1198, tokenIndex1198, depth1198
						{
							position1290 := position
							depth++
							{
								position1291, tokenIndex1291, depth1291 := position, tokenIndex, depth
								if buffer[position] != rune('s') {
									goto l1292
								}
								position++
								goto l1291
							l1292:
								position, tokenIndex, depth = position1291, tokenIndex1291, depth1291
								if buffer[position] != rune('S') {
									goto l1289
								}
								position++
							}
						l1291:
							{
								position1293, tokenIndex1293, depth1293 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l1294
								}
								position++
								goto l1293
							l1294:
								position, tokenIndex, depth = position1293, tokenIndex1293, depth1293
								if buffer[position] != rune('T') {
									goto l1289
								}
								position++
							}
						l1293:
							{
								position1295, tokenIndex1295, depth1295 := position, tokenIndex, depth
								if buffer[position] != rune('r') {
									goto l1296
								}
								position++
								goto l1295
							l1296:
								position, tokenIndex, depth = position1295, tokenIndex1295, depth1295
								if buffer[position] != rune('R') {
									goto l1289
								}
								position++
							}
						l1295:
							{
								position1297, tokenIndex1297, depth1297 := position, tokenIndex, depth
								if buffer[position] != rune('d') {
									goto l1298
								}
								position++
								goto l1297
							l1298:
								position, tokenIndex, depth = position1297, tokenIndex1297, depth1297
								if buffer[position] != rune('D') {
									goto l1289
								}
								position++
							}
						l1297:
							{
								position1299, tokenIndex1299, depth1299 := position, tokenIndex, depth
								if buffer[position] != rune('t') {
									goto l1300
								}
								position++
								goto l1299
							l1300:
								position, tokenIndex, depth = position1299, tokenIndex1299, depth1299
								if buffer[position] != rune('T') {
									goto l1289
								}
								position++
							}
						l1299:
							if !rules[ruleskip]() {
								goto l1289
							}
							depth--
							add(ruleSTRDT, position1290)
						}
						goto l1198
					l1289:
						position, tokenIndex, depth = position1198, tokenIndex1198, depth1198
						{
							switch buffer[position] {
							case 'S', 's':
								{
									position1302 := position
									depth++
									{
										position1303, tokenIndex1303, depth1303 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1304
										}
										position++
										goto l1303
									l1304:
										position, tokenIndex, depth = position1303, tokenIndex1303, depth1303
										if buffer[position] != rune('S') {
											goto l1197
										}
										position++
									}
								l1303:
									{
										position1305, tokenIndex1305, depth1305 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l1306
										}
										position++
										goto l1305
									l1306:
										position, tokenIndex, depth = position1305, tokenIndex1305, depth1305
										if buffer[position] != rune('A') {
											goto l1197
										}
										position++
									}
								l1305:
									{
										position1307, tokenIndex1307, depth1307 := position, tokenIndex, depth
										if buffer[position] != rune('m') {
											goto l1308
										}
										position++
										goto l1307
									l1308:
										position, tokenIndex, depth = position1307, tokenIndex1307, depth1307
										if buffer[position] != rune('M') {
											goto l1197
										}
										position++
									}
								l1307:
									{
										position1309, tokenIndex1309, depth1309 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1310
										}
										position++
										goto l1309
									l1310:
										position, tokenIndex, depth = position1309, tokenIndex1309, depth1309
										if buffer[position] != rune('E') {
											goto l1197
										}
										position++
									}
								l1309:
									{
										position1311, tokenIndex1311, depth1311 := position, tokenIndex, depth
										if buffer[position] != rune('t') {
											goto l1312
										}
										position++
										goto l1311
									l1312:
										position, tokenIndex, depth = position1311, tokenIndex1311, depth1311
										if buffer[position] != rune('T') {
											goto l1197
										}
										position++
									}
								l1311:
									{
										position1313, tokenIndex1313, depth1313 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1314
										}
										position++
										goto l1313
									l1314:
										position, tokenIndex, depth = position1313, tokenIndex1313, depth1313
										if buffer[position] != rune('E') {
											goto l1197
										}
										position++
									}
								l1313:
									{
										position1315, tokenIndex1315, depth1315 := position, tokenIndex, depth
										if buffer[position] != rune('r') {
											goto l1316
										}
										position++
										goto l1315
									l1316:
										position, tokenIndex, depth = position1315, tokenIndex1315, depth1315
										if buffer[position] != rune('R') {
											goto l1197
										}
										position++
									}
								l1315:
									{
										position1317, tokenIndex1317, depth1317 := position, tokenIndex, depth
										if buffer[position] != rune('m') {
											goto l1318
										}
										position++
										goto l1317
									l1318:
										position, tokenIndex, depth = position1317, tokenIndex1317, depth1317
										if buffer[position] != rune('M') {
											goto l1197
										}
										position++
									}
								l1317:
									if !rules[ruleskip]() {
										goto l1197
									}
									depth--
									add(ruleSAMETERM, position1302)
								}
								break
							case 'C', 'c':
								{
									position1319 := position
									depth++
									{
										position1320, tokenIndex1320, depth1320 := position, tokenIndex, depth
										if buffer[position] != rune('c') {
											goto l1321
										}
										position++
										goto l1320
									l1321:
										position, tokenIndex, depth = position1320, tokenIndex1320, depth1320
										if buffer[position] != rune('C') {
											goto l1197
										}
										position++
									}
								l1320:
									{
										position1322, tokenIndex1322, depth1322 := position, tokenIndex, depth
										if buffer[position] != rune('o') {
											goto l1323
										}
										position++
										goto l1322
									l1323:
										position, tokenIndex, depth = position1322, tokenIndex1322, depth1322
										if buffer[position] != rune('O') {
											goto l1197
										}
										position++
									}
								l1322:
									{
										position1324, tokenIndex1324, depth1324 := position, tokenIndex, depth
										if buffer[position] != rune('n') {
											goto l1325
										}
										position++
										goto l1324
									l1325:
										position, tokenIndex, depth = position1324, tokenIndex1324, depth1324
										if buffer[position] != rune('N') {
											goto l1197
										}
										position++
									}
								l1324:
									{
										position1326, tokenIndex1326, depth1326 := position, tokenIndex, depth
										if buffer[position] != rune('t') {
											goto l1327
										}
										position++
										goto l1326
									l1327:
										position, tokenIndex, depth = position1326, tokenIndex1326, depth1326
										if buffer[position] != rune('T') {
											goto l1197
										}
										position++
									}
								l1326:
									{
										position1328, tokenIndex1328, depth1328 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l1329
										}
										position++
										goto l1328
									l1329:
										position, tokenIndex, depth = position1328, tokenIndex1328, depth1328
										if buffer[position] != rune('A') {
											goto l1197
										}
										position++
									}
								l1328:
									{
										position1330, tokenIndex1330, depth1330 := position, tokenIndex, depth
										if buffer[position] != rune('i') {
											goto l1331
										}
										position++
										goto l1330
									l1331:
										position, tokenIndex, depth = position1330, tokenIndex1330, depth1330
										if buffer[position] != rune('I') {
											goto l1197
										}
										position++
									}
								l1330:
									{
										position1332, tokenIndex1332, depth1332 := position, tokenIndex, depth
										if buffer[position] != rune('n') {
											goto l1333
										}
										position++
										goto l1332
									l1333:
										position, tokenIndex, depth = position1332, tokenIndex1332, depth1332
										if buffer[position] != rune('N') {
											goto l1197
										}
										position++
									}
								l1332:
									{
										position1334, tokenIndex1334, depth1334 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1335
										}
										position++
										goto l1334
									l1335:
										position, tokenIndex, depth = position1334, tokenIndex1334, depth1334
										if buffer[position] != rune('S') {
											goto l1197
										}
										position++
									}
								l1334:
									if !rules[ruleskip]() {
										goto l1197
									}
									depth--
									add(ruleCONTAINS, position1319)
								}
								break
							default:
								{
									position1336 := position
									depth++
									{
										position1337, tokenIndex1337, depth1337 := position, tokenIndex, depth
										if buffer[position] != rune('l') {
											goto l1338
										}
										position++
										goto l1337
									l1338:
										position, tokenIndex, depth = position1337, tokenIndex1337, depth1337
										if buffer[position] != rune('L') {
											goto l1197
										}
										position++
									}
								l1337:
									{
										position1339, tokenIndex1339, depth1339 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l1340
										}
										position++
										goto l1339
									l1340:
										position, tokenIndex, depth = position1339, tokenIndex1339, depth1339
										if buffer[position] != rune('A') {
											goto l1197
										}
										position++
									}
								l1339:
									{
										position1341, tokenIndex1341, depth1341 := position, tokenIndex, depth
										if buffer[position] != rune('n') {
											goto l1342
										}
										position++
										goto l1341
									l1342:
										position, tokenIndex, depth = position1341, tokenIndex1341, depth1341
										if buffer[position] != rune('N') {
											goto l1197
										}
										position++
									}
								l1341:
									{
										position1343, tokenIndex1343, depth1343 := position, tokenIndex, depth
										if buffer[position] != rune('g') {
											goto l1344
										}
										position++
										goto l1343
									l1344:
										position, tokenIndex, depth = position1343, tokenIndex1343, depth1343
										if buffer[position] != rune('G') {
											goto l1197
										}
										position++
									}
								l1343:
									{
										position1345, tokenIndex1345, depth1345 := position, tokenIndex, depth
										if buffer[position] != rune('m') {
											goto l1346
										}
										position++
										goto l1345
									l1346:
										position, tokenIndex, depth = position1345, tokenIndex1345, depth1345
										if buffer[position] != rune('M') {
											goto l1197
										}
										position++
									}
								l1345:
									{
										position1347, tokenIndex1347, depth1347 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l1348
										}
										position++
										goto l1347
									l1348:
										position, tokenIndex, depth = position1347, tokenIndex1347, depth1347
										if buffer[position] != rune('A') {
											goto l1197
										}
										position++
									}
								l1347:
									{
										position1349, tokenIndex1349, depth1349 := position, tokenIndex, depth
										if buffer[position] != rune('t') {
											goto l1350
										}
										position++
										goto l1349
									l1350:
										position, tokenIndex, depth = position1349, tokenIndex1349, depth1349
										if buffer[position] != rune('T') {
											goto l1197
										}
										position++
									}
								l1349:
									{
										position1351, tokenIndex1351, depth1351 := position, tokenIndex, depth
										if buffer[position] != rune('c') {
											goto l1352
										}
										position++
										goto l1351
									l1352:
										position, tokenIndex, depth = position1351, tokenIndex1351, depth1351
										if buffer[position] != rune('C') {
											goto l1197
										}
										position++
									}
								l1351:
									{
										position1353, tokenIndex1353, depth1353 := position, tokenIndex, depth
										if buffer[position] != rune('h') {
											goto l1354
										}
										position++
										goto l1353
									l1354:
										position, tokenIndex, depth = position1353, tokenIndex1353, depth1353
										if buffer[position] != rune('H') {
											goto l1197
										}
										position++
									}
								l1353:
									{
										position1355, tokenIndex1355, depth1355 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1356
										}
										position++
										goto l1355
									l1356:
										position, tokenIndex, depth = position1355, tokenIndex1355, depth1355
										if buffer[position] != rune('E') {
											goto l1197
										}
										position++
									}
								l1355:
									{
										position1357, tokenIndex1357, depth1357 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1358
										}
										position++
										goto l1357
									l1358:
										position, tokenIndex, depth = position1357, tokenIndex1357, depth1357
										if buffer[position] != rune('S') {
											goto l1197
										}
										position++
									}
								l1357:
									if !rules[ruleskip]() {
										goto l1197
									}
									depth--
									add(ruleLANGMATCHES, position1336)
								}
								break
							}
						}

					}
				l1198:
					if !rules[ruleLPAREN]() {
						goto l1197
					}
					if !rules[ruleexpression]() {
						goto l1197
					}
					if !rules[ruleCOMMA]() {
						goto l1197
					}
					if !rules[ruleexpression]() {
						goto l1197
					}
					if !rules[ruleRPAREN]() {
						goto l1197
					}
					goto l833
				l1197:
					position, tokenIndex, depth = position833, tokenIndex833, depth833
					{
						position1360 := position
						depth++
						{
							position1361, tokenIndex1361, depth1361 := position, tokenIndex, depth
							if buffer[position] != rune('b') {
								goto l1362
							}
							position++
							goto l1361
						l1362:
							position, tokenIndex, depth = position1361, tokenIndex1361, depth1361
							if buffer[position] != rune('B') {
								goto l1359
							}
							position++
						}
					l1361:
						{
							position1363, tokenIndex1363, depth1363 := position, tokenIndex, depth
							if buffer[position] != rune('o') {
								goto l1364
							}
							position++
							goto l1363
						l1364:
							position, tokenIndex, depth = position1363, tokenIndex1363, depth1363
							if buffer[position] != rune('O') {
								goto l1359
							}
							position++
						}
					l1363:
						{
							position1365, tokenIndex1365, depth1365 := position, tokenIndex, depth
							if buffer[position] != rune('u') {
								goto l1366
							}
							position++
							goto l1365
						l1366:
							position, tokenIndex, depth = position1365, tokenIndex1365, depth1365
							if buffer[position] != rune('U') {
								goto l1359
							}
							position++
						}
					l1365:
						{
							position1367, tokenIndex1367, depth1367 := position, tokenIndex, depth
							if buffer[position] != rune('n') {
								goto l1368
							}
							position++
							goto l1367
						l1368:
							position, tokenIndex, depth = position1367, tokenIndex1367, depth1367
							if buffer[position] != rune('N') {
								goto l1359
							}
							position++
						}
					l1367:
						{
							position1369, tokenIndex1369, depth1369 := position, tokenIndex, depth
							if buffer[position] != rune('d') {
								goto l1370
							}
							position++
							goto l1369
						l1370:
							position, tokenIndex, depth = position1369, tokenIndex1369, depth1369
							if buffer[position] != rune('D') {
								goto l1359
							}
							position++
						}
					l1369:
						if !rules[ruleskip]() {
							goto l1359
						}
						depth--
						add(ruleBOUND, position1360)
					}
					if !rules[ruleLPAREN]() {
						goto l1359
					}
					if !rules[rulevar]() {
						goto l1359
					}
					if !rules[ruleRPAREN]() {
						goto l1359
					}
					goto l833
				l1359:
					position, tokenIndex, depth = position833, tokenIndex833, depth833
					{
						switch buffer[position] {
						case 'S', 's':
							{
								position1373 := position
								depth++
								{
									position1374, tokenIndex1374, depth1374 := position, tokenIndex, depth
									if buffer[position] != rune('s') {
										goto l1375
									}
									position++
									goto l1374
								l1375:
									position, tokenIndex, depth = position1374, tokenIndex1374, depth1374
									if buffer[position] != rune('S') {
										goto l1371
									}
									position++
								}
							l1374:
								{
									position1376, tokenIndex1376, depth1376 := position, tokenIndex, depth
									if buffer[position] != rune('t') {
										goto l1377
									}
									position++
									goto l1376
								l1377:
									position, tokenIndex, depth = position1376, tokenIndex1376, depth1376
									if buffer[position] != rune('T') {
										goto l1371
									}
									position++
								}
							l1376:
								{
									position1378, tokenIndex1378, depth1378 := position, tokenIndex, depth
									if buffer[position] != rune('r') {
										goto l1379
									}
									position++
									goto l1378
								l1379:
									position, tokenIndex, depth = position1378, tokenIndex1378, depth1378
									if buffer[position] != rune('R') {
										goto l1371
									}
									position++
								}
							l1378:
								{
									position1380, tokenIndex1380, depth1380 := position, tokenIndex, depth
									if buffer[position] != rune('u') {
										goto l1381
									}
									position++
									goto l1380
								l1381:
									position, tokenIndex, depth = position1380, tokenIndex1380, depth1380
									if buffer[position] != rune('U') {
										goto l1371
									}
									position++
								}
							l1380:
								{
									position1382, tokenIndex1382, depth1382 := position, tokenIndex, depth
									if buffer[position] != rune('u') {
										goto l1383
									}
									position++
									goto l1382
								l1383:
									position, tokenIndex, depth = position1382, tokenIndex1382, depth1382
									if buffer[position] != rune('U') {
										goto l1371
									}
									position++
								}
							l1382:
								{
									position1384, tokenIndex1384, depth1384 := position, tokenIndex, depth
									if buffer[position] != rune('i') {
										goto l1385
									}
									position++
									goto l1384
								l1385:
									position, tokenIndex, depth = position1384, tokenIndex1384, depth1384
									if buffer[position] != rune('I') {
										goto l1371
									}
									position++
								}
							l1384:
								{
									position1386, tokenIndex1386, depth1386 := position, tokenIndex, depth
									if buffer[position] != rune('d') {
										goto l1387
									}
									position++
									goto l1386
								l1387:
									position, tokenIndex, depth = position1386, tokenIndex1386, depth1386
									if buffer[position] != rune('D') {
										goto l1371
									}
									position++
								}
							l1386:
								if !rules[ruleskip]() {
									goto l1371
								}
								depth--
								add(ruleSTRUUID, position1373)
							}
							break
						case 'U', 'u':
							{
								position1388 := position
								depth++
								{
									position1389, tokenIndex1389, depth1389 := position, tokenIndex, depth
									if buffer[position] != rune('u') {
										goto l1390
									}
									position++
									goto l1389
								l1390:
									position, tokenIndex, depth = position1389, tokenIndex1389, depth1389
									if buffer[position] != rune('U') {
										goto l1371
									}
									position++
								}
							l1389:
								{
									position1391, tokenIndex1391, depth1391 := position, tokenIndex, depth
									if buffer[position] != rune('u') {
										goto l1392
									}
									position++
									goto l1391
								l1392:
									position, tokenIndex, depth = position1391, tokenIndex1391, depth1391
									if buffer[position] != rune('U') {
										goto l1371
									}
									position++
								}
							l1391:
								{
									position1393, tokenIndex1393, depth1393 := position, tokenIndex, depth
									if buffer[position] != rune('i') {
										goto l1394
									}
									position++
									goto l1393
								l1394:
									position, tokenIndex, depth = position1393, tokenIndex1393, depth1393
									if buffer[position] != rune('I') {
										goto l1371
									}
									position++
								}
							l1393:
								{
									position1395, tokenIndex1395, depth1395 := position, tokenIndex, depth
									if buffer[position] != rune('d') {
										goto l1396
									}
									position++
									goto l1395
								l1396:
									position, tokenIndex, depth = position1395, tokenIndex1395, depth1395
									if buffer[position] != rune('D') {
										goto l1371
									}
									position++
								}
							l1395:
								if !rules[ruleskip]() {
									goto l1371
								}
								depth--
								add(ruleUUID, position1388)
							}
							break
						case 'N', 'n':
							{
								position1397 := position
								depth++
								{
									position1398, tokenIndex1398, depth1398 := position, tokenIndex, depth
									if buffer[position] != rune('n') {
										goto l1399
									}
									position++
									goto l1398
								l1399:
									position, tokenIndex, depth = position1398, tokenIndex1398, depth1398
									if buffer[position] != rune('N') {
										goto l1371
									}
									position++
								}
							l1398:
								{
									position1400, tokenIndex1400, depth1400 := position, tokenIndex, depth
									if buffer[position] != rune('o') {
										goto l1401
									}
									position++
									goto l1400
								l1401:
									position, tokenIndex, depth = position1400, tokenIndex1400, depth1400
									if buffer[position] != rune('O') {
										goto l1371
									}
									position++
								}
							l1400:
								{
									position1402, tokenIndex1402, depth1402 := position, tokenIndex, depth
									if buffer[position] != rune('w') {
										goto l1403
									}
									position++
									goto l1402
								l1403:
									position, tokenIndex, depth = position1402, tokenIndex1402, depth1402
									if buffer[position] != rune('W') {
										goto l1371
									}
									position++
								}
							l1402:
								if !rules[ruleskip]() {
									goto l1371
								}
								depth--
								add(ruleNOW, position1397)
							}
							break
						default:
							{
								position1404 := position
								depth++
								{
									position1405, tokenIndex1405, depth1405 := position, tokenIndex, depth
									if buffer[position] != rune('r') {
										goto l1406
									}
									position++
									goto l1405
								l1406:
									position, tokenIndex, depth = position1405, tokenIndex1405, depth1405
									if buffer[position] != rune('R') {
										goto l1371
									}
									position++
								}
							l1405:
								{
									position1407, tokenIndex1407, depth1407 := position, tokenIndex, depth
									if buffer[position] != rune('a') {
										goto l1408
									}
									position++
									goto l1407
								l1408:
									position, tokenIndex, depth = position1407, tokenIndex1407, depth1407
									if buffer[position] != rune('A') {
										goto l1371
									}
									position++
								}
							l1407:
								{
									position1409, tokenIndex1409, depth1409 := position, tokenIndex, depth
									if buffer[position] != rune('n') {
										goto l1410
									}
									position++
									goto l1409
								l1410:
									position, tokenIndex, depth = position1409, tokenIndex1409, depth1409
									if buffer[position] != rune('N') {
										goto l1371
									}
									position++
								}
							l1409:
								{
									position1411, tokenIndex1411, depth1411 := position, tokenIndex, depth
									if buffer[position] != rune('d') {
										goto l1412
									}
									position++
									goto l1411
								l1412:
									position, tokenIndex, depth = position1411, tokenIndex1411, depth1411
									if buffer[position] != rune('D') {
										goto l1371
									}
									position++
								}
							l1411:
								if !rules[ruleskip]() {
									goto l1371
								}
								depth--
								add(ruleRAND, position1404)
							}
							break
						}
					}

					if !rules[rulenil]() {
						goto l1371
					}
					goto l833
				l1371:
					position, tokenIndex, depth = position833, tokenIndex833, depth833
					{
						switch buffer[position] {
						case 'E', 'N', 'e', 'n':
							{
								position1414, tokenIndex1414, depth1414 := position, tokenIndex, depth
								{
									position1416 := position
									depth++
									{
										position1417, tokenIndex1417, depth1417 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1418
										}
										position++
										goto l1417
									l1418:
										position, tokenIndex, depth = position1417, tokenIndex1417, depth1417
										if buffer[position] != rune('E') {
											goto l1415
										}
										position++
									}
								l1417:
									{
										position1419, tokenIndex1419, depth1419 := position, tokenIndex, depth
										if buffer[position] != rune('x') {
											goto l1420
										}
										position++
										goto l1419
									l1420:
										position, tokenIndex, depth = position1419, tokenIndex1419, depth1419
										if buffer[position] != rune('X') {
											goto l1415
										}
										position++
									}
								l1419:
									{
										position1421, tokenIndex1421, depth1421 := position, tokenIndex, depth
										if buffer[position] != rune('i') {
											goto l1422
										}
										position++
										goto l1421
									l1422:
										position, tokenIndex, depth = position1421, tokenIndex1421, depth1421
										if buffer[position] != rune('I') {
											goto l1415
										}
										position++
									}
								l1421:
									{
										position1423, tokenIndex1423, depth1423 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1424
										}
										position++
										goto l1423
									l1424:
										position, tokenIndex, depth = position1423, tokenIndex1423, depth1423
										if buffer[position] != rune('S') {
											goto l1415
										}
										position++
									}
								l1423:
									{
										position1425, tokenIndex1425, depth1425 := position, tokenIndex, depth
										if buffer[position] != rune('t') {
											goto l1426
										}
										position++
										goto l1425
									l1426:
										position, tokenIndex, depth = position1425, tokenIndex1425, depth1425
										if buffer[position] != rune('T') {
											goto l1415
										}
										position++
									}
								l1425:
									{
										position1427, tokenIndex1427, depth1427 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1428
										}
										position++
										goto l1427
									l1428:
										position, tokenIndex, depth = position1427, tokenIndex1427, depth1427
										if buffer[position] != rune('S') {
											goto l1415
										}
										position++
									}
								l1427:
									if !rules[ruleskip]() {
										goto l1415
									}
									depth--
									add(ruleEXISTS, position1416)
								}
								goto l1414
							l1415:
								position, tokenIndex, depth = position1414, tokenIndex1414, depth1414
								{
									position1429 := position
									depth++
									{
										position1430, tokenIndex1430, depth1430 := position, tokenIndex, depth
										if buffer[position] != rune('n') {
											goto l1431
										}
										position++
										goto l1430
									l1431:
										position, tokenIndex, depth = position1430, tokenIndex1430, depth1430
										if buffer[position] != rune('N') {
											goto l831
										}
										position++
									}
								l1430:
									{
										position1432, tokenIndex1432, depth1432 := position, tokenIndex, depth
										if buffer[position] != rune('o') {
											goto l1433
										}
										position++
										goto l1432
									l1433:
										position, tokenIndex, depth = position1432, tokenIndex1432, depth1432
										if buffer[position] != rune('O') {
											goto l831
										}
										position++
									}
								l1432:
									{
										position1434, tokenIndex1434, depth1434 := position, tokenIndex, depth
										if buffer[position] != rune('t') {
											goto l1435
										}
										position++
										goto l1434
									l1435:
										position, tokenIndex, depth = position1434, tokenIndex1434, depth1434
										if buffer[position] != rune('T') {
											goto l831
										}
										position++
									}
								l1434:
									if buffer[position] != rune(' ') {
										goto l831
									}
									position++
									{
										position1436, tokenIndex1436, depth1436 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1437
										}
										position++
										goto l1436
									l1437:
										position, tokenIndex, depth = position1436, tokenIndex1436, depth1436
										if buffer[position] != rune('E') {
											goto l831
										}
										position++
									}
								l1436:
									{
										position1438, tokenIndex1438, depth1438 := position, tokenIndex, depth
										if buffer[position] != rune('x') {
											goto l1439
										}
										position++
										goto l1438
									l1439:
										position, tokenIndex, depth = position1438, tokenIndex1438, depth1438
										if buffer[position] != rune('X') {
											goto l831
										}
										position++
									}
								l1438:
									{
										position1440, tokenIndex1440, depth1440 := position, tokenIndex, depth
										if buffer[position] != rune('i') {
											goto l1441
										}
										position++
										goto l1440
									l1441:
										position, tokenIndex, depth = position1440, tokenIndex1440, depth1440
										if buffer[position] != rune('I') {
											goto l831
										}
										position++
									}
								l1440:
									{
										position1442, tokenIndex1442, depth1442 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1443
										}
										position++
										goto l1442
									l1443:
										position, tokenIndex, depth = position1442, tokenIndex1442, depth1442
										if buffer[position] != rune('S') {
											goto l831
										}
										position++
									}
								l1442:
									{
										position1444, tokenIndex1444, depth1444 := position, tokenIndex, depth
										if buffer[position] != rune('t') {
											goto l1445
										}
										position++
										goto l1444
									l1445:
										position, tokenIndex, depth = position1444, tokenIndex1444, depth1444
										if buffer[position] != rune('T') {
											goto l831
										}
										position++
									}
								l1444:
									{
										position1446, tokenIndex1446, depth1446 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1447
										}
										position++
										goto l1446
									l1447:
										position, tokenIndex, depth = position1446, tokenIndex1446, depth1446
										if buffer[position] != rune('S') {
											goto l831
										}
										position++
									}
								l1446:
									if !rules[ruleskip]() {
										goto l831
									}
									depth--
									add(ruleNOTEXIST, position1429)
								}
							}
						l1414:
							if !rules[rulegroupGraphPattern]() {
								goto l831
							}
							break
						case 'I', 'i':
							{
								position1448 := position
								depth++
								{
									position1449, tokenIndex1449, depth1449 := position, tokenIndex, depth
									if buffer[position] != rune('i') {
										goto l1450
									}
									position++
									goto l1449
								l1450:
									position, tokenIndex, depth = position1449, tokenIndex1449, depth1449
									if buffer[position] != rune('I') {
										goto l831
									}
									position++
								}
							l1449:
								{
									position1451, tokenIndex1451, depth1451 := position, tokenIndex, depth
									if buffer[position] != rune('f') {
										goto l1452
									}
									position++
									goto l1451
								l1452:
									position, tokenIndex, depth = position1451, tokenIndex1451, depth1451
									if buffer[position] != rune('F') {
										goto l831
									}
									position++
								}
							l1451:
								if !rules[ruleskip]() {
									goto l831
								}
								depth--
								add(ruleIF, position1448)
							}
							if !rules[ruleLPAREN]() {
								goto l831
							}
							if !rules[ruleexpression]() {
								goto l831
							}
							if !rules[ruleCOMMA]() {
								goto l831
							}
							if !rules[ruleexpression]() {
								goto l831
							}
							if !rules[ruleCOMMA]() {
								goto l831
							}
							if !rules[ruleexpression]() {
								goto l831
							}
							if !rules[ruleRPAREN]() {
								goto l831
							}
							break
						case 'C', 'c':
							{
								position1453, tokenIndex1453, depth1453 := position, tokenIndex, depth
								{
									position1455 := position
									depth++
									{
										position1456, tokenIndex1456, depth1456 := position, tokenIndex, depth
										if buffer[position] != rune('c') {
											goto l1457
										}
										position++
										goto l1456
									l1457:
										position, tokenIndex, depth = position1456, tokenIndex1456, depth1456
										if buffer[position] != rune('C') {
											goto l1454
										}
										position++
									}
								l1456:
									{
										position1458, tokenIndex1458, depth1458 := position, tokenIndex, depth
										if buffer[position] != rune('o') {
											goto l1459
										}
										position++
										goto l1458
									l1459:
										position, tokenIndex, depth = position1458, tokenIndex1458, depth1458
										if buffer[position] != rune('O') {
											goto l1454
										}
										position++
									}
								l1458:
									{
										position1460, tokenIndex1460, depth1460 := position, tokenIndex, depth
										if buffer[position] != rune('n') {
											goto l1461
										}
										position++
										goto l1460
									l1461:
										position, tokenIndex, depth = position1460, tokenIndex1460, depth1460
										if buffer[position] != rune('N') {
											goto l1454
										}
										position++
									}
								l1460:
									{
										position1462, tokenIndex1462, depth1462 := position, tokenIndex, depth
										if buffer[position] != rune('c') {
											goto l1463
										}
										position++
										goto l1462
									l1463:
										position, tokenIndex, depth = position1462, tokenIndex1462, depth1462
										if buffer[position] != rune('C') {
											goto l1454
										}
										position++
									}
								l1462:
									{
										position1464, tokenIndex1464, depth1464 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l1465
										}
										position++
										goto l1464
									l1465:
										position, tokenIndex, depth = position1464, tokenIndex1464, depth1464
										if buffer[position] != rune('A') {
											goto l1454
										}
										position++
									}
								l1464:
									{
										position1466, tokenIndex1466, depth1466 := position, tokenIndex, depth
										if buffer[position] != rune('t') {
											goto l1467
										}
										position++
										goto l1466
									l1467:
										position, tokenIndex, depth = position1466, tokenIndex1466, depth1466
										if buffer[position] != rune('T') {
											goto l1454
										}
										position++
									}
								l1466:
									if !rules[ruleskip]() {
										goto l1454
									}
									depth--
									add(ruleCONCAT, position1455)
								}
								goto l1453
							l1454:
								position, tokenIndex, depth = position1453, tokenIndex1453, depth1453
								{
									position1468 := position
									depth++
									{
										position1469, tokenIndex1469, depth1469 := position, tokenIndex, depth
										if buffer[position] != rune('c') {
											goto l1470
										}
										position++
										goto l1469
									l1470:
										position, tokenIndex, depth = position1469, tokenIndex1469, depth1469
										if buffer[position] != rune('C') {
											goto l831
										}
										position++
									}
								l1469:
									{
										position1471, tokenIndex1471, depth1471 := position, tokenIndex, depth
										if buffer[position] != rune('o') {
											goto l1472
										}
										position++
										goto l1471
									l1472:
										position, tokenIndex, depth = position1471, tokenIndex1471, depth1471
										if buffer[position] != rune('O') {
											goto l831
										}
										position++
									}
								l1471:
									{
										position1473, tokenIndex1473, depth1473 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l1474
										}
										position++
										goto l1473
									l1474:
										position, tokenIndex, depth = position1473, tokenIndex1473, depth1473
										if buffer[position] != rune('A') {
											goto l831
										}
										position++
									}
								l1473:
									{
										position1475, tokenIndex1475, depth1475 := position, tokenIndex, depth
										if buffer[position] != rune('l') {
											goto l1476
										}
										position++
										goto l1475
									l1476:
										position, tokenIndex, depth = position1475, tokenIndex1475, depth1475
										if buffer[position] != rune('L') {
											goto l831
										}
										position++
									}
								l1475:
									{
										position1477, tokenIndex1477, depth1477 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1478
										}
										position++
										goto l1477
									l1478:
										position, tokenIndex, depth = position1477, tokenIndex1477, depth1477
										if buffer[position] != rune('E') {
											goto l831
										}
										position++
									}
								l1477:
									{
										position1479, tokenIndex1479, depth1479 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1480
										}
										position++
										goto l1479
									l1480:
										position, tokenIndex, depth = position1479, tokenIndex1479, depth1479
										if buffer[position] != rune('S') {
											goto l831
										}
										position++
									}
								l1479:
									{
										position1481, tokenIndex1481, depth1481 := position, tokenIndex, depth
										if buffer[position] != rune('c') {
											goto l1482
										}
										position++
										goto l1481
									l1482:
										position, tokenIndex, depth = position1481, tokenIndex1481, depth1481
										if buffer[position] != rune('C') {
											goto l831
										}
										position++
									}
								l1481:
									{
										position1483, tokenIndex1483, depth1483 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1484
										}
										position++
										goto l1483
									l1484:
										position, tokenIndex, depth = position1483, tokenIndex1483, depth1483
										if buffer[position] != rune('E') {
											goto l831
										}
										position++
									}
								l1483:
									if !rules[ruleskip]() {
										goto l831
									}
									depth--
									add(ruleCOALESCE, position1468)
								}
							}
						l1453:
							if !rules[ruleargList]() {
								goto l831
							}
							break
						case 'B', 'b':
							{
								position1485 := position
								depth++
								{
									position1486, tokenIndex1486, depth1486 := position, tokenIndex, depth
									if buffer[position] != rune('b') {
										goto l1487
									}
									position++
									goto l1486
								l1487:
									position, tokenIndex, depth = position1486, tokenIndex1486, depth1486
									if buffer[position] != rune('B') {
										goto l831
									}
									position++
								}
							l1486:
								{
									position1488, tokenIndex1488, depth1488 := position, tokenIndex, depth
									if buffer[position] != rune('n') {
										goto l1489
									}
									position++
									goto l1488
								l1489:
									position, tokenIndex, depth = position1488, tokenIndex1488, depth1488
									if buffer[position] != rune('N') {
										goto l831
									}
									position++
								}
							l1488:
								{
									position1490, tokenIndex1490, depth1490 := position, tokenIndex, depth
									if buffer[position] != rune('o') {
										goto l1491
									}
									position++
									goto l1490
								l1491:
									position, tokenIndex, depth = position1490, tokenIndex1490, depth1490
									if buffer[position] != rune('O') {
										goto l831
									}
									position++
								}
							l1490:
								{
									position1492, tokenIndex1492, depth1492 := position, tokenIndex, depth
									if buffer[position] != rune('d') {
										goto l1493
									}
									position++
									goto l1492
								l1493:
									position, tokenIndex, depth = position1492, tokenIndex1492, depth1492
									if buffer[position] != rune('D') {
										goto l831
									}
									position++
								}
							l1492:
								{
									position1494, tokenIndex1494, depth1494 := position, tokenIndex, depth
									if buffer[position] != rune('e') {
										goto l1495
									}
									position++
									goto l1494
								l1495:
									position, tokenIndex, depth = position1494, tokenIndex1494, depth1494
									if buffer[position] != rune('E') {
										goto l831
									}
									position++
								}
							l1494:
								if !rules[ruleskip]() {
									goto l831
								}
								depth--
								add(ruleBNODE, position1485)
							}
							{
								position1496, tokenIndex1496, depth1496 := position, tokenIndex, depth
								if !rules[ruleLPAREN]() {
									goto l1497
								}
								if !rules[ruleexpression]() {
									goto l1497
								}
								if !rules[ruleRPAREN]() {
									goto l1497
								}
								goto l1496
							l1497:
								position, tokenIndex, depth = position1496, tokenIndex1496, depth1496
								if !rules[rulenil]() {
									goto l831
								}
							}
						l1496:
							break
						default:
							{
								position1498, tokenIndex1498, depth1498 := position, tokenIndex, depth
								{
									position1500 := position
									depth++
									{
										position1501, tokenIndex1501, depth1501 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1502
										}
										position++
										goto l1501
									l1502:
										position, tokenIndex, depth = position1501, tokenIndex1501, depth1501
										if buffer[position] != rune('S') {
											goto l1499
										}
										position++
									}
								l1501:
									{
										position1503, tokenIndex1503, depth1503 := position, tokenIndex, depth
										if buffer[position] != rune('u') {
											goto l1504
										}
										position++
										goto l1503
									l1504:
										position, tokenIndex, depth = position1503, tokenIndex1503, depth1503
										if buffer[position] != rune('U') {
											goto l1499
										}
										position++
									}
								l1503:
									{
										position1505, tokenIndex1505, depth1505 := position, tokenIndex, depth
										if buffer[position] != rune('b') {
											goto l1506
										}
										position++
										goto l1505
									l1506:
										position, tokenIndex, depth = position1505, tokenIndex1505, depth1505
										if buffer[position] != rune('B') {
											goto l1499
										}
										position++
									}
								l1505:
									{
										position1507, tokenIndex1507, depth1507 := position, tokenIndex, depth
										if buffer[position] != rune('s') {
											goto l1508
										}
										position++
										goto l1507
									l1508:
										position, tokenIndex, depth = position1507, tokenIndex1507, depth1507
										if buffer[position] != rune('S') {
											goto l1499
										}
										position++
									}
								l1507:
									{
										position1509, tokenIndex1509, depth1509 := position, tokenIndex, depth
										if buffer[position] != rune('t') {
											goto l1510
										}
										position++
										goto l1509
									l1510:
										position, tokenIndex, depth = position1509, tokenIndex1509, depth1509
										if buffer[position] != rune('T') {
											goto l1499
										}
										position++
									}
								l1509:
									{
										position1511, tokenIndex1511, depth1511 := position, tokenIndex, depth
										if buffer[position] != rune('r') {
											goto l1512
										}
										position++
										goto l1511
									l1512:
										position, tokenIndex, depth = position1511, tokenIndex1511, depth1511
										if buffer[position] != rune('R') {
											goto l1499
										}
										position++
									}
								l1511:
									if !rules[ruleskip]() {
										goto l1499
									}
									depth--
									add(ruleSUBSTR, position1500)
								}
								goto l1498
							l1499:
								position, tokenIndex, depth = position1498, tokenIndex1498, depth1498
								{
									position1514 := position
									depth++
									{
										position1515, tokenIndex1515, depth1515 := position, tokenIndex, depth
										if buffer[position] != rune('r') {
											goto l1516
										}
										position++
										goto l1515
									l1516:
										position, tokenIndex, depth = position1515, tokenIndex1515, depth1515
										if buffer[position] != rune('R') {
											goto l1513
										}
										position++
									}
								l1515:
									{
										position1517, tokenIndex1517, depth1517 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1518
										}
										position++
										goto l1517
									l1518:
										position, tokenIndex, depth = position1517, tokenIndex1517, depth1517
										if buffer[position] != rune('E') {
											goto l1513
										}
										position++
									}
								l1517:
									{
										position1519, tokenIndex1519, depth1519 := position, tokenIndex, depth
										if buffer[position] != rune('p') {
											goto l1520
										}
										position++
										goto l1519
									l1520:
										position, tokenIndex, depth = position1519, tokenIndex1519, depth1519
										if buffer[position] != rune('P') {
											goto l1513
										}
										position++
									}
								l1519:
									{
										position1521, tokenIndex1521, depth1521 := position, tokenIndex, depth
										if buffer[position] != rune('l') {
											goto l1522
										}
										position++
										goto l1521
									l1522:
										position, tokenIndex, depth = position1521, tokenIndex1521, depth1521
										if buffer[position] != rune('L') {
											goto l1513
										}
										position++
									}
								l1521:
									{
										position1523, tokenIndex1523, depth1523 := position, tokenIndex, depth
										if buffer[position] != rune('a') {
											goto l1524
										}
										position++
										goto l1523
									l1524:
										position, tokenIndex, depth = position1523, tokenIndex1523, depth1523
										if buffer[position] != rune('A') {
											goto l1513
										}
										position++
									}
								l1523:
									{
										position1525, tokenIndex1525, depth1525 := position, tokenIndex, depth
										if buffer[position] != rune('c') {
											goto l1526
										}
										position++
										goto l1525
									l1526:
										position, tokenIndex, depth = position1525, tokenIndex1525, depth1525
										if buffer[position] != rune('C') {
											goto l1513
										}
										position++
									}
								l1525:
									{
										position1527, tokenIndex1527, depth1527 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1528
										}
										position++
										goto l1527
									l1528:
										position, tokenIndex, depth = position1527, tokenIndex1527, depth1527
										if buffer[position] != rune('E') {
											goto l1513
										}
										position++
									}
								l1527:
									if !rules[ruleskip]() {
										goto l1513
									}
									depth--
									add(ruleREPLACE, position1514)
								}
								goto l1498
							l1513:
								position, tokenIndex, depth = position1498, tokenIndex1498, depth1498
								{
									position1529 := position
									depth++
									{
										position1530, tokenIndex1530, depth1530 := position, tokenIndex, depth
										if buffer[position] != rune('r') {
											goto l1531
										}
										position++
										goto l1530
									l1531:
										position, tokenIndex, depth = position1530, tokenIndex1530, depth1530
										if buffer[position] != rune('R') {
											goto l831
										}
										position++
									}
								l1530:
									{
										position1532, tokenIndex1532, depth1532 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1533
										}
										position++
										goto l1532
									l1533:
										position, tokenIndex, depth = position1532, tokenIndex1532, depth1532
										if buffer[position] != rune('E') {
											goto l831
										}
										position++
									}
								l1532:
									{
										position1534, tokenIndex1534, depth1534 := position, tokenIndex, depth
										if buffer[position] != rune('g') {
											goto l1535
										}
										position++
										goto l1534
									l1535:
										position, tokenIndex, depth = position1534, tokenIndex1534, depth1534
										if buffer[position] != rune('G') {
											goto l831
										}
										position++
									}
								l1534:
									{
										position1536, tokenIndex1536, depth1536 := position, tokenIndex, depth
										if buffer[position] != rune('e') {
											goto l1537
										}
										position++
										goto l1536
									l1537:
										position, tokenIndex, depth = position1536, tokenIndex1536, depth1536
										if buffer[position] != rune('E') {
											goto l831
										}
										position++
									}
								l1536:
									{
										position1538, tokenIndex1538, depth1538 := position, tokenIndex, depth
										if buffer[position] != rune('x') {
											goto l1539
										}
										position++
										goto l1538
									l1539:
										position, tokenIndex, depth = position1538, tokenIndex1538, depth1538
										if buffer[position] != rune('X') {
											goto l831
										}
										position++
									}
								l1538:
									if !rules[ruleskip]() {
										goto l831
									}
									depth--
									add(ruleREGEX, position1529)
								}
							}
						l1498:
							if !rules[ruleLPAREN]() {
								goto l831
							}
							if !rules[ruleexpression]() {
								goto l831
							}
							if !rules[ruleCOMMA]() {
								goto l831
							}
							if !rules[ruleexpression]() {
								goto l831
							}
							{
								position1540, tokenIndex1540, depth1540 := position, tokenIndex, depth
								if !rules[ruleCOMMA]() {
									goto l1540
								}
								if !rules[ruleexpression]() {
									goto l1540
								}
								goto l1541
							l1540:
								position, tokenIndex, depth = position1540, tokenIndex1540, depth1540
							}
						l1541:
							if !rules[ruleRPAREN]() {
								goto l831
							}
							break
						}
					}

				}
			l833:
				depth--
				add(rulebuiltinCall, position832)
			}
			return true
		l831:
			position, tokenIndex, depth = position831, tokenIndex831, depth831
			return false
		},
		/* 70 var <- <(<(('?' / '$') VARNAME)> Action3 skip)> */
		func() bool {
			position1542, tokenIndex1542, depth1542 := position, tokenIndex, depth
			{
				position1543 := position
				depth++
				{
					position1544 := position
					depth++
					{
						position1545, tokenIndex1545, depth1545 := position, tokenIndex, depth
						if buffer[position] != rune('?') {
							goto l1546
						}
						position++
						goto l1545
					l1546:
						position, tokenIndex, depth = position1545, tokenIndex1545, depth1545
						if buffer[position] != rune('$') {
							goto l1542
						}
						position++
					}
				l1545:
					{
						position1547 := position
						depth++
						{
							position1548, tokenIndex1548, depth1548 := position, tokenIndex, depth
							if !rules[rulepnCharsU]() {
								goto l1549
							}
							goto l1548
						l1549:
							position, tokenIndex, depth = position1548, tokenIndex1548, depth1548
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1542
							}
							position++
						}
					l1548:
					l1550:
						{
							position1551, tokenIndex1551, depth1551 := position, tokenIndex, depth
							{
								position1552, tokenIndex1552, depth1552 := position, tokenIndex, depth
								if !rules[rulepnCharsU]() {
									goto l1553
								}
								goto l1552
							l1553:
								position, tokenIndex, depth = position1552, tokenIndex1552, depth1552
								{
									switch buffer[position] {
									case 'â':
										if c := buffer[position]; c < rune('‿') || c > rune('⁀') {
											goto l1551
										}
										position++
										break
									case 'Ì', 'Í':
										if c := buffer[position]; c < rune('̀') || c > rune('ͯ') {
											goto l1551
										}
										position++
										break
									case 'Â':
										if buffer[position] != rune('·') {
											goto l1551
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l1551
										}
										position++
										break
									}
								}

							}
						l1552:
							goto l1550
						l1551:
							position, tokenIndex, depth = position1551, tokenIndex1551, depth1551
						}
						depth--
						add(ruleVARNAME, position1547)
					}
					depth--
					add(rulePegText, position1544)
				}
				{
					add(ruleAction3, position)
				}
				if !rules[ruleskip]() {
					goto l1542
				}
				depth--
				add(rulevar, position1543)
			}
			return true
		l1542:
			position, tokenIndex, depth = position1542, tokenIndex1542, depth1542
			return false
		},
		/* 71 iriref <- <(iri / prefixedName)> */
		func() bool {
			position1556, tokenIndex1556, depth1556 := position, tokenIndex, depth
			{
				position1557 := position
				depth++
				{
					position1558, tokenIndex1558, depth1558 := position, tokenIndex, depth
					if !rules[ruleiri]() {
						goto l1559
					}
					goto l1558
				l1559:
					position, tokenIndex, depth = position1558, tokenIndex1558, depth1558
					{
						position1560 := position
						depth++
						{
							position1561, tokenIndex1561, depth1561 := position, tokenIndex, depth
							if !rules[rulepnPrefix]() {
								goto l1561
							}
							goto l1562
						l1561:
							position, tokenIndex, depth = position1561, tokenIndex1561, depth1561
						}
					l1562:
						if buffer[position] != rune(':') {
							goto l1556
						}
						position++
						{
							position1563 := position
							depth++
						l1564:
							{
								position1565, tokenIndex1565, depth1565 := position, tokenIndex, depth
								{
									switch buffer[position] {
									case '%', '\\':
										{
											position1567 := position
											depth++
											{
												position1568, tokenIndex1568, depth1568 := position, tokenIndex, depth
												{
													position1570 := position
													depth++
													if buffer[position] != rune('%') {
														goto l1569
													}
													position++
													if !rules[rulehex]() {
														goto l1569
													}
													if !rules[rulehex]() {
														goto l1569
													}
													depth--
													add(rulepercent, position1570)
												}
												goto l1568
											l1569:
												position, tokenIndex, depth = position1568, tokenIndex1568, depth1568
												{
													position1571 := position
													depth++
													if buffer[position] != rune('\\') {
														goto l1565
													}
													position++
													{
														switch buffer[position] {
														case '%':
															if buffer[position] != rune('%') {
																goto l1565
															}
															position++
															break
														case '@':
															if buffer[position] != rune('@') {
																goto l1565
															}
															position++
															break
														case '#':
															if buffer[position] != rune('#') {
																goto l1565
															}
															position++
															break
														case '?':
															if buffer[position] != rune('?') {
																goto l1565
															}
															position++
															break
														case '/':
															if buffer[position] != rune('/') {
																goto l1565
															}
															position++
															break
														case '=':
															if buffer[position] != rune('=') {
																goto l1565
															}
															position++
															break
														case ';':
															if buffer[position] != rune(';') {
																goto l1565
															}
															position++
															break
														case ',':
															if buffer[position] != rune(',') {
																goto l1565
															}
															position++
															break
														case '+':
															if buffer[position] != rune('+') {
																goto l1565
															}
															position++
															break
														case '*':
															if buffer[position] != rune('*') {
																goto l1565
															}
															position++
															break
														case ')':
															if buffer[position] != rune(')') {
																goto l1565
															}
															position++
															break
														case '(':
															if buffer[position] != rune('(') {
																goto l1565
															}
															position++
															break
														case '\'':
															if buffer[position] != rune('\'') {
																goto l1565
															}
															position++
															break
														case '&':
															if buffer[position] != rune('&') {
																goto l1565
															}
															position++
															break
														case '$':
															if buffer[position] != rune('$') {
																goto l1565
															}
															position++
															break
														case '!':
															if buffer[position] != rune('!') {
																goto l1565
															}
															position++
															break
														case '-':
															if buffer[position] != rune('-') {
																goto l1565
															}
															position++
															break
														case '.':
															if buffer[position] != rune('.') {
																goto l1565
															}
															position++
															break
														case '~':
															if buffer[position] != rune('~') {
																goto l1565
															}
															position++
															break
														default:
															if buffer[position] != rune('_') {
																goto l1565
															}
															position++
															break
														}
													}

													depth--
													add(rulepnLocalEsc, position1571)
												}
											}
										l1568:
											depth--
											add(ruleplx, position1567)
										}
										break
									case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l1565
										}
										position++
										break
									case ':':
										if buffer[position] != rune(':') {
											goto l1565
										}
										position++
										break
									default:
										if !rules[rulepnCharsU]() {
											goto l1565
										}
										break
									}
								}

								goto l1564
							l1565:
								position, tokenIndex, depth = position1565, tokenIndex1565, depth1565
							}
							depth--
							add(rulepnLocal, position1563)
						}
						if !rules[ruleskip]() {
							goto l1556
						}
						depth--
						add(ruleprefixedName, position1560)
					}
				}
			l1558:
				depth--
				add(ruleiriref, position1557)
			}
			return true
		l1556:
			position, tokenIndex, depth = position1556, tokenIndex1556, depth1556
			return false
		},
		/* 72 iri <- <(<('<' (!'>' .)* '>')> Action4 skip)> */
		func() bool {
			position1573, tokenIndex1573, depth1573 := position, tokenIndex, depth
			{
				position1574 := position
				depth++
				{
					position1575 := position
					depth++
					if buffer[position] != rune('<') {
						goto l1573
					}
					position++
				l1576:
					{
						position1577, tokenIndex1577, depth1577 := position, tokenIndex, depth
						{
							position1578, tokenIndex1578, depth1578 := position, tokenIndex, depth
							if buffer[position] != rune('>') {
								goto l1578
							}
							position++
							goto l1577
						l1578:
							position, tokenIndex, depth = position1578, tokenIndex1578, depth1578
						}
						if !matchDot() {
							goto l1577
						}
						goto l1576
					l1577:
						position, tokenIndex, depth = position1577, tokenIndex1577, depth1577
					}
					if buffer[position] != rune('>') {
						goto l1573
					}
					position++
					depth--
					add(rulePegText, position1575)
				}
				{
					add(ruleAction4, position)
				}
				if !rules[ruleskip]() {
					goto l1573
				}
				depth--
				add(ruleiri, position1574)
			}
			return true
		l1573:
			position, tokenIndex, depth = position1573, tokenIndex1573, depth1573
			return false
		},
		/* 73 prefixedName <- <(pnPrefix? ':' pnLocal skip)> */
		nil,
		/* 74 literal <- <(<(string (('@' ([a-z] / [A-Z])+ ('-' ((&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]) | (&('A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' | 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z') [A-Z]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'g' | 'h' | 'i' | 'j' | 'k' | 'l' | 'm' | 'n' | 'o' | 'p' | 'q' | 'r' | 's' | 't' | 'u' | 'v' | 'w' | 'x' | 'y' | 'z') [a-z]))+)*) / ('^' '^' iriref))?)> Action5 skip)> */
		func() bool {
			position1581, tokenIndex1581, depth1581 := position, tokenIndex, depth
			{
				position1582 := position
				depth++
				{
					position1583 := position
					depth++
					if !rules[rulestring]() {
						goto l1581
					}
					{
						position1584, tokenIndex1584, depth1584 := position, tokenIndex, depth
						{
							position1586, tokenIndex1586, depth1586 := position, tokenIndex, depth
							if buffer[position] != rune('@') {
								goto l1587
							}
							position++
							{
								position1590, tokenIndex1590, depth1590 := position, tokenIndex, depth
								if c := buffer[position]; c < rune('a') || c > rune('z') {
									goto l1591
								}
								position++
								goto l1590
							l1591:
								position, tokenIndex, depth = position1590, tokenIndex1590, depth1590
								if c := buffer[position]; c < rune('A') || c > rune('Z') {
									goto l1587
								}
								position++
							}
						l1590:
						l1588:
							{
								position1589, tokenIndex1589, depth1589 := position, tokenIndex, depth
								{
									position1592, tokenIndex1592, depth1592 := position, tokenIndex, depth
									if c := buffer[position]; c < rune('a') || c > rune('z') {
										goto l1593
									}
									position++
									goto l1592
								l1593:
									position, tokenIndex, depth = position1592, tokenIndex1592, depth1592
									if c := buffer[position]; c < rune('A') || c > rune('Z') {
										goto l1589
									}
									position++
								}
							l1592:
								goto l1588
							l1589:
								position, tokenIndex, depth = position1589, tokenIndex1589, depth1589
							}
						l1594:
							{
								position1595, tokenIndex1595, depth1595 := position, tokenIndex, depth
								if buffer[position] != rune('-') {
									goto l1595
								}
								position++
								{
									switch buffer[position] {
									case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l1595
										}
										position++
										break
									case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
										if c := buffer[position]; c < rune('A') || c > rune('Z') {
											goto l1595
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('a') || c > rune('z') {
											goto l1595
										}
										position++
										break
									}
								}

							l1596:
								{
									position1597, tokenIndex1597, depth1597 := position, tokenIndex, depth
									{
										switch buffer[position] {
										case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
											if c := buffer[position]; c < rune('0') || c > rune('9') {
												goto l1597
											}
											position++
											break
										case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
											if c := buffer[position]; c < rune('A') || c > rune('Z') {
												goto l1597
											}
											position++
											break
										default:
											if c := buffer[position]; c < rune('a') || c > rune('z') {
												goto l1597
											}
											position++
											break
										}
									}

									goto l1596
								l1597:
									position, tokenIndex, depth = position1597, tokenIndex1597, depth1597
								}
								goto l1594
							l1595:
								position, tokenIndex, depth = position1595, tokenIndex1595, depth1595
							}
							goto l1586
						l1587:
							position, tokenIndex, depth = position1586, tokenIndex1586, depth1586
							if buffer[position] != rune('^') {
								goto l1584
							}
							position++
							if buffer[position] != rune('^') {
								goto l1584
							}
							position++
							if !rules[ruleiriref]() {
								goto l1584
							}
						}
					l1586:
						goto l1585
					l1584:
						position, tokenIndex, depth = position1584, tokenIndex1584, depth1584
					}
				l1585:
					depth--
					add(rulePegText, position1583)
				}
				{
					add(ruleAction5, position)
				}
				if !rules[ruleskip]() {
					goto l1581
				}
				depth--
				add(ruleliteral, position1582)
			}
			return true
		l1581:
			position, tokenIndex, depth = position1581, tokenIndex1581, depth1581
			return false
		},
		/* 75 string <- <(stringLiteralA / stringLiteralB / stringLiteralLongA / stringLiteralLongB)> */
		func() bool {
			position1601, tokenIndex1601, depth1601 := position, tokenIndex, depth
			{
				position1602 := position
				depth++
				{
					position1603, tokenIndex1603, depth1603 := position, tokenIndex, depth
					{
						position1605 := position
						depth++
						if buffer[position] != rune('\'') {
							goto l1604
						}
						position++
					l1606:
						{
							position1607, tokenIndex1607, depth1607 := position, tokenIndex, depth
							{
								position1608, tokenIndex1608, depth1608 := position, tokenIndex, depth
								{
									position1610, tokenIndex1610, depth1610 := position, tokenIndex, depth
									{
										switch buffer[position] {
										case '\r':
											if buffer[position] != rune('\r') {
												goto l1610
											}
											position++
											break
										case '\n':
											if buffer[position] != rune('\n') {
												goto l1610
											}
											position++
											break
										case '\\':
											if buffer[position] != rune('\\') {
												goto l1610
											}
											position++
											break
										default:
											if buffer[position] != rune('\'') {
												goto l1610
											}
											position++
											break
										}
									}

									goto l1609
								l1610:
									position, tokenIndex, depth = position1610, tokenIndex1610, depth1610
								}
								if !matchDot() {
									goto l1609
								}
								goto l1608
							l1609:
								position, tokenIndex, depth = position1608, tokenIndex1608, depth1608
								if !rules[ruleechar]() {
									goto l1607
								}
							}
						l1608:
							goto l1606
						l1607:
							position, tokenIndex, depth = position1607, tokenIndex1607, depth1607
						}
						if buffer[position] != rune('\'') {
							goto l1604
						}
						position++
						depth--
						add(rulestringLiteralA, position1605)
					}
					goto l1603
				l1604:
					position, tokenIndex, depth = position1603, tokenIndex1603, depth1603
					{
						position1613 := position
						depth++
						if buffer[position] != rune('"') {
							goto l1612
						}
						position++
					l1614:
						{
							position1615, tokenIndex1615, depth1615 := position, tokenIndex, depth
							{
								position1616, tokenIndex1616, depth1616 := position, tokenIndex, depth
								{
									position1618, tokenIndex1618, depth1618 := position, tokenIndex, depth
									{
										switch buffer[position] {
										case '\r':
											if buffer[position] != rune('\r') {
												goto l1618
											}
											position++
											break
										case '\n':
											if buffer[position] != rune('\n') {
												goto l1618
											}
											position++
											break
										case '\\':
											if buffer[position] != rune('\\') {
												goto l1618
											}
											position++
											break
										default:
											if buffer[position] != rune('"') {
												goto l1618
											}
											position++
											break
										}
									}

									goto l1617
								l1618:
									position, tokenIndex, depth = position1618, tokenIndex1618, depth1618
								}
								if !matchDot() {
									goto l1617
								}
								goto l1616
							l1617:
								position, tokenIndex, depth = position1616, tokenIndex1616, depth1616
								if !rules[ruleechar]() {
									goto l1615
								}
							}
						l1616:
							goto l1614
						l1615:
							position, tokenIndex, depth = position1615, tokenIndex1615, depth1615
						}
						if buffer[position] != rune('"') {
							goto l1612
						}
						position++
						depth--
						add(rulestringLiteralB, position1613)
					}
					goto l1603
				l1612:
					position, tokenIndex, depth = position1603, tokenIndex1603, depth1603
					{
						position1621 := position
						depth++
						if buffer[position] != rune('\'') {
							goto l1620
						}
						position++
						if buffer[position] != rune('\'') {
							goto l1620
						}
						position++
						if buffer[position] != rune('\'') {
							goto l1620
						}
						position++
					l1622:
						{
							position1623, tokenIndex1623, depth1623 := position, tokenIndex, depth
							{
								position1624, tokenIndex1624, depth1624 := position, tokenIndex, depth
								{
									position1626, tokenIndex1626, depth1626 := position, tokenIndex, depth
									if buffer[position] != rune('\'') {
										goto l1627
									}
									position++
									goto l1626
								l1627:
									position, tokenIndex, depth = position1626, tokenIndex1626, depth1626
									if buffer[position] != rune('\'') {
										goto l1624
									}
									position++
									if buffer[position] != rune('\'') {
										goto l1624
									}
									position++
								}
							l1626:
								goto l1625
							l1624:
								position, tokenIndex, depth = position1624, tokenIndex1624, depth1624
							}
						l1625:
							{
								position1628, tokenIndex1628, depth1628 := position, tokenIndex, depth
								{
									position1630, tokenIndex1630, depth1630 := position, tokenIndex, depth
									{
										position1631, tokenIndex1631, depth1631 := position, tokenIndex, depth
										if buffer[position] != rune('\'') {
											goto l1632
										}
										position++
										goto l1631
									l1632:
										position, tokenIndex, depth = position1631, tokenIndex1631, depth1631
										if buffer[position] != rune('\\') {
											goto l1630
										}
										position++
									}
								l1631:
									goto l1629
								l1630:
									position, tokenIndex, depth = position1630, tokenIndex1630, depth1630
								}
								if !matchDot() {
									goto l1629
								}
								goto l1628
							l1629:
								position, tokenIndex, depth = position1628, tokenIndex1628, depth1628
								if !rules[ruleechar]() {
									goto l1623
								}
							}
						l1628:
							goto l1622
						l1623:
							position, tokenIndex, depth = position1623, tokenIndex1623, depth1623
						}
						if buffer[position] != rune('\'') {
							goto l1620
						}
						position++
						if buffer[position] != rune('\'') {
							goto l1620
						}
						position++
						if buffer[position] != rune('\'') {
							goto l1620
						}
						position++
						depth--
						add(rulestringLiteralLongA, position1621)
					}
					goto l1603
				l1620:
					position, tokenIndex, depth = position1603, tokenIndex1603, depth1603
					{
						position1633 := position
						depth++
						if buffer[position] != rune('"') {
							goto l1601
						}
						position++
						if buffer[position] != rune('"') {
							goto l1601
						}
						position++
						if buffer[position] != rune('"') {
							goto l1601
						}
						position++
					l1634:
						{
							position1635, tokenIndex1635, depth1635 := position, tokenIndex, depth
							{
								position1636, tokenIndex1636, depth1636 := position, tokenIndex, depth
								{
									position1638, tokenIndex1638, depth1638 := position, tokenIndex, depth
									if buffer[position] != rune('"') {
										goto l1639
									}
									position++
									goto l1638
								l1639:
									position, tokenIndex, depth = position1638, tokenIndex1638, depth1638
									if buffer[position] != rune('"') {
										goto l1636
									}
									position++
									if buffer[position] != rune('"') {
										goto l1636
									}
									position++
								}
							l1638:
								goto l1637
							l1636:
								position, tokenIndex, depth = position1636, tokenIndex1636, depth1636
							}
						l1637:
							{
								position1640, tokenIndex1640, depth1640 := position, tokenIndex, depth
								{
									position1642, tokenIndex1642, depth1642 := position, tokenIndex, depth
									{
										position1643, tokenIndex1643, depth1643 := position, tokenIndex, depth
										if buffer[position] != rune('"') {
											goto l1644
										}
										position++
										goto l1643
									l1644:
										position, tokenIndex, depth = position1643, tokenIndex1643, depth1643
										if buffer[position] != rune('\\') {
											goto l1642
										}
										position++
									}
								l1643:
									goto l1641
								l1642:
									position, tokenIndex, depth = position1642, tokenIndex1642, depth1642
								}
								if !matchDot() {
									goto l1641
								}
								goto l1640
							l1641:
								position, tokenIndex, depth = position1640, tokenIndex1640, depth1640
								if !rules[ruleechar]() {
									goto l1635
								}
							}
						l1640:
							goto l1634
						l1635:
							position, tokenIndex, depth = position1635, tokenIndex1635, depth1635
						}
						if buffer[position] != rune('"') {
							goto l1601
						}
						position++
						if buffer[position] != rune('"') {
							goto l1601
						}
						position++
						if buffer[position] != rune('"') {
							goto l1601
						}
						position++
						depth--
						add(rulestringLiteralLongB, position1633)
					}
				}
			l1603:
				depth--
				add(rulestring, position1602)
			}
			return true
		l1601:
			position, tokenIndex, depth = position1601, tokenIndex1601, depth1601
			return false
		},
		/* 76 stringLiteralA <- <('\'' ((!((&('\r') '\r') | (&('\n') '\n') | (&('\\') '\\') | (&('\'') '\'')) .) / echar)* '\'')> */
		nil,
		/* 77 stringLiteralB <- <('"' ((!((&('\r') '\r') | (&('\n') '\n') | (&('\\') '\\') | (&('"') '"')) .) / echar)* '"')> */
		nil,
		/* 78 stringLiteralLongA <- <('\'' '\'' '\'' (('\'' / ('\'' '\''))? ((!('\'' / '\\') .) / echar))* ('\'' '\'' '\''))> */
		nil,
		/* 79 stringLiteralLongB <- <('"' '"' '"' (('"' / ('"' '"'))? ((!('"' / '\\') .) / echar))* ('"' '"' '"'))> */
		nil,
		/* 80 echar <- <('\\' ((&('\'') '\'') | (&('"') '"') | (&('\\') '\\') | (&('f') 'f') | (&('r') 'r') | (&('n') 'n') | (&('b') 'b') | (&('t') 't') | (&('u') 'u')))> */
		func() bool {
			position1649, tokenIndex1649, depth1649 := position, tokenIndex, depth
			{
				position1650 := position
				depth++
				if buffer[position] != rune('\\') {
					goto l1649
				}
				position++
				{
					switch buffer[position] {
					case '\'':
						if buffer[position] != rune('\'') {
							goto l1649
						}
						position++
						break
					case '"':
						if buffer[position] != rune('"') {
							goto l1649
						}
						position++
						break
					case '\\':
						if buffer[position] != rune('\\') {
							goto l1649
						}
						position++
						break
					case 'f':
						if buffer[position] != rune('f') {
							goto l1649
						}
						position++
						break
					case 'r':
						if buffer[position] != rune('r') {
							goto l1649
						}
						position++
						break
					case 'n':
						if buffer[position] != rune('n') {
							goto l1649
						}
						position++
						break
					case 'b':
						if buffer[position] != rune('b') {
							goto l1649
						}
						position++
						break
					case 't':
						if buffer[position] != rune('t') {
							goto l1649
						}
						position++
						break
					default:
						if buffer[position] != rune('u') {
							goto l1649
						}
						position++
						break
					}
				}

				depth--
				add(ruleechar, position1650)
			}
			return true
		l1649:
			position, tokenIndex, depth = position1649, tokenIndex1649, depth1649
			return false
		},
		/* 81 numericLiteral <- <(('+' / '-')? [0-9]+ ('.' [0-9]*)? skip)> */
		func() bool {
			position1652, tokenIndex1652, depth1652 := position, tokenIndex, depth
			{
				position1653 := position
				depth++
				{
					position1654, tokenIndex1654, depth1654 := position, tokenIndex, depth
					{
						position1656, tokenIndex1656, depth1656 := position, tokenIndex, depth
						if buffer[position] != rune('+') {
							goto l1657
						}
						position++
						goto l1656
					l1657:
						position, tokenIndex, depth = position1656, tokenIndex1656, depth1656
						if buffer[position] != rune('-') {
							goto l1654
						}
						position++
					}
				l1656:
					goto l1655
				l1654:
					position, tokenIndex, depth = position1654, tokenIndex1654, depth1654
				}
			l1655:
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l1652
				}
				position++
			l1658:
				{
					position1659, tokenIndex1659, depth1659 := position, tokenIndex, depth
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1659
					}
					position++
					goto l1658
				l1659:
					position, tokenIndex, depth = position1659, tokenIndex1659, depth1659
				}
				{
					position1660, tokenIndex1660, depth1660 := position, tokenIndex, depth
					if buffer[position] != rune('.') {
						goto l1660
					}
					position++
				l1662:
					{
						position1663, tokenIndex1663, depth1663 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1663
						}
						position++
						goto l1662
					l1663:
						position, tokenIndex, depth = position1663, tokenIndex1663, depth1663
					}
					goto l1661
				l1660:
					position, tokenIndex, depth = position1660, tokenIndex1660, depth1660
				}
			l1661:
				if !rules[ruleskip]() {
					goto l1652
				}
				depth--
				add(rulenumericLiteral, position1653)
			}
			return true
		l1652:
			position, tokenIndex, depth = position1652, tokenIndex1652, depth1652
			return false
		},
		/* 82 signedNumericLiteral <- <(('+' / '-') [0-9]+ ('.' [0-9]*)? skip)> */
		nil,
		/* 83 booleanLiteral <- <(TRUE / FALSE)> */
		func() bool {
			position1665, tokenIndex1665, depth1665 := position, tokenIndex, depth
			{
				position1666 := position
				depth++
				{
					position1667, tokenIndex1667, depth1667 := position, tokenIndex, depth
					{
						position1669 := position
						depth++
						{
							position1670, tokenIndex1670, depth1670 := position, tokenIndex, depth
							if buffer[position] != rune('t') {
								goto l1671
							}
							position++
							goto l1670
						l1671:
							position, tokenIndex, depth = position1670, tokenIndex1670, depth1670
							if buffer[position] != rune('T') {
								goto l1668
							}
							position++
						}
					l1670:
						{
							position1672, tokenIndex1672, depth1672 := position, tokenIndex, depth
							if buffer[position] != rune('r') {
								goto l1673
							}
							position++
							goto l1672
						l1673:
							position, tokenIndex, depth = position1672, tokenIndex1672, depth1672
							if buffer[position] != rune('R') {
								goto l1668
							}
							position++
						}
					l1672:
						{
							position1674, tokenIndex1674, depth1674 := position, tokenIndex, depth
							if buffer[position] != rune('u') {
								goto l1675
							}
							position++
							goto l1674
						l1675:
							position, tokenIndex, depth = position1674, tokenIndex1674, depth1674
							if buffer[position] != rune('U') {
								goto l1668
							}
							position++
						}
					l1674:
						{
							position1676, tokenIndex1676, depth1676 := position, tokenIndex, depth
							if buffer[position] != rune('e') {
								goto l1677
							}
							position++
							goto l1676
						l1677:
							position, tokenIndex, depth = position1676, tokenIndex1676, depth1676
							if buffer[position] != rune('E') {
								goto l1668
							}
							position++
						}
					l1676:
						if !rules[ruleskip]() {
							goto l1668
						}
						depth--
						add(ruleTRUE, position1669)
					}
					goto l1667
				l1668:
					position, tokenIndex, depth = position1667, tokenIndex1667, depth1667
					{
						position1678 := position
						depth++
						{
							position1679, tokenIndex1679, depth1679 := position, tokenIndex, depth
							if buffer[position] != rune('f') {
								goto l1680
							}
							position++
							goto l1679
						l1680:
							position, tokenIndex, depth = position1679, tokenIndex1679, depth1679
							if buffer[position] != rune('F') {
								goto l1665
							}
							position++
						}
					l1679:
						{
							position1681, tokenIndex1681, depth1681 := position, tokenIndex, depth
							if buffer[position] != rune('a') {
								goto l1682
							}
							position++
							goto l1681
						l1682:
							position, tokenIndex, depth = position1681, tokenIndex1681, depth1681
							if buffer[position] != rune('A') {
								goto l1665
							}
							position++
						}
					l1681:
						{
							position1683, tokenIndex1683, depth1683 := position, tokenIndex, depth
							if buffer[position] != rune('l') {
								goto l1684
							}
							position++
							goto l1683
						l1684:
							position, tokenIndex, depth = position1683, tokenIndex1683, depth1683
							if buffer[position] != rune('L') {
								goto l1665
							}
							position++
						}
					l1683:
						{
							position1685, tokenIndex1685, depth1685 := position, tokenIndex, depth
							if buffer[position] != rune('s') {
								goto l1686
							}
							position++
							goto l1685
						l1686:
							position, tokenIndex, depth = position1685, tokenIndex1685, depth1685
							if buffer[position] != rune('S') {
								goto l1665
							}
							position++
						}
					l1685:
						{
							position1687, tokenIndex1687, depth1687 := position, tokenIndex, depth
							if buffer[position] != rune('e') {
								goto l1688
							}
							position++
							goto l1687
						l1688:
							position, tokenIndex, depth = position1687, tokenIndex1687, depth1687
							if buffer[position] != rune('E') {
								goto l1665
							}
							position++
						}
					l1687:
						if !rules[ruleskip]() {
							goto l1665
						}
						depth--
						add(ruleFALSE, position1678)
					}
				}
			l1667:
				depth--
				add(rulebooleanLiteral, position1666)
			}
			return true
		l1665:
			position, tokenIndex, depth = position1665, tokenIndex1665, depth1665
			return false
		},
		/* 84 blankNode <- <(blankNodeLabel / anon)> */
		nil,
		/* 85 blankNodeLabel <- <('_' ':' (pnCharsU / [0-9]) (((pnCharsU / ((&('.') '.') | (&('-') '-') | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9])))* pnCharsU) / ([0-9] / '-'))? skip)> */
		nil,
		/* 86 anon <- <('[' ws* ']' skip)> */
		nil,
		/* 87 nil <- <('(' ws* ')' skip)> */
		func() bool {
			position1692, tokenIndex1692, depth1692 := position, tokenIndex, depth
			{
				position1693 := position
				depth++
				if buffer[position] != rune('(') {
					goto l1692
				}
				position++
			l1694:
				{
					position1695, tokenIndex1695, depth1695 := position, tokenIndex, depth
					if !rules[rulews]() {
						goto l1695
					}
					goto l1694
				l1695:
					position, tokenIndex, depth = position1695, tokenIndex1695, depth1695
				}
				if buffer[position] != rune(')') {
					goto l1692
				}
				position++
				if !rules[ruleskip]() {
					goto l1692
				}
				depth--
				add(rulenil, position1693)
			}
			return true
		l1692:
			position, tokenIndex, depth = position1692, tokenIndex1692, depth1692
			return false
		},
		/* 88 VARNAME <- <((pnCharsU / [0-9]) (pnCharsU / ((&('â') [‿-⁀]) | (&('Ì' | 'Í') [̀-ͯ]) | (&('Â') '·') | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9])))*)> */
		nil,
		/* 89 pnPrefix <- <(pnCharsBase pnChars*)> */
		func() bool {
			position1697, tokenIndex1697, depth1697 := position, tokenIndex, depth
			{
				position1698 := position
				depth++
				if !rules[rulepnCharsBase]() {
					goto l1697
				}
			l1699:
				{
					position1700, tokenIndex1700, depth1700 := position, tokenIndex, depth
					{
						position1701 := position
						depth++
						{
							switch buffer[position] {
							case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l1700
								}
								position++
								break
							case '-':
								if buffer[position] != rune('-') {
									goto l1700
								}
								position++
								break
							default:
								if !rules[rulepnCharsU]() {
									goto l1700
								}
								break
							}
						}

						depth--
						add(rulepnChars, position1701)
					}
					goto l1699
				l1700:
					position, tokenIndex, depth = position1700, tokenIndex1700, depth1700
				}
				depth--
				add(rulepnPrefix, position1698)
			}
			return true
		l1697:
			position, tokenIndex, depth = position1697, tokenIndex1697, depth1697
			return false
		},
		/* 90 pnLocal <- <((&('%' | '\\') plx) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]) | (&(':') ':') | (&('A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' | 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z' | '_' | 'a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'g' | 'h' | 'i' | 'j' | 'k' | 'l' | 'm' | 'n' | 'o' | 'p' | 'q' | 'r' | 's' | 't' | 'u' | 'v' | 'w' | 'x' | 'y' | 'z' | 'Ã' | 'Ä' | 'Å' | 'Æ' | 'Ç' | 'È' | 'É' | 'Ê' | 'Ë' | 'Í' | 'Î' | 'Ï' | 'Ð' | 'Ñ' | 'Ò' | 'Ó' | 'Ô' | 'Õ' | 'Ö' | '×' | 'Ø' | 'Ù' | 'Ú' | 'Û' | 'Ü' | 'Ý' | 'Þ' | 'ß' | 'à' | 'á' | 'â' | 'ã' | 'ä' | 'å' | 'æ' | 'ç' | 'è' | 'é' | 'ê' | 'ë' | 'ì' | 'í' | 'ï' | 'ð' | 'ñ' | 'ò' | 'ó') pnCharsU))*> */
		nil,
		/* 91 pnChars <- <((&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]) | (&('-') '-') | (&('A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' | 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z' | '_' | 'a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'g' | 'h' | 'i' | 'j' | 'k' | 'l' | 'm' | 'n' | 'o' | 'p' | 'q' | 'r' | 's' | 't' | 'u' | 'v' | 'w' | 'x' | 'y' | 'z' | 'Ã' | 'Ä' | 'Å' | 'Æ' | 'Ç' | 'È' | 'É' | 'Ê' | 'Ë' | 'Í' | 'Î' | 'Ï' | 'Ð' | 'Ñ' | 'Ò' | 'Ó' | 'Ô' | 'Õ' | 'Ö' | '×' | 'Ø' | 'Ù' | 'Ú' | 'Û' | 'Ü' | 'Ý' | 'Þ' | 'ß' | 'à' | 'á' | 'â' | 'ã' | 'ä' | 'å' | 'æ' | 'ç' | 'è' | 'é' | 'ê' | 'ë' | 'ì' | 'í' | 'ï' | 'ð' | 'ñ' | 'ò' | 'ó') pnCharsU))> */
		nil,
		/* 92 pnCharsU <- <(pnCharsBase / '_')> */
		func() bool {
			position1705, tokenIndex1705, depth1705 := position, tokenIndex, depth
			{
				position1706 := position
				depth++
				{
					position1707, tokenIndex1707, depth1707 := position, tokenIndex, depth
					if !rules[rulepnCharsBase]() {
						goto l1708
					}
					goto l1707
				l1708:
					position, tokenIndex, depth = position1707, tokenIndex1707, depth1707
					if buffer[position] != rune('_') {
						goto l1705
					}
					position++
				}
			l1707:
				depth--
				add(rulepnCharsU, position1706)
			}
			return true
		l1705:
			position, tokenIndex, depth = position1705, tokenIndex1705, depth1705
			return false
		},
		/* 93 pnCharsBase <- <([À-Ö] / [Ø-ö] / [Ͱ-ͽ] / [‌-‍] / [⁰-↏] / [豈-﷏] / ((&('ð' | 'ñ' | 'ò' | 'ó') [𐀀-󯿿]) | (&('ï') [ﷰ-�]) | (&('ã' | 'ä' | 'å' | 'æ' | 'ç' | 'è' | 'é' | 'ê' | 'ë' | 'ì' | 'í') [、-퟿]) | (&('â') [Ⰰ-⿯]) | (&('Í' | 'Î' | 'Ï' | 'Ð' | 'Ñ' | 'Ò' | 'Ó' | 'Ô' | 'Õ' | 'Ö' | '×' | 'Ø' | 'Ù' | 'Ú' | 'Û' | 'Ü' | 'Ý' | 'Þ' | 'ß' | 'à' | 'á') [Ϳ-῿]) | (&('Ã' | 'Ä' | 'Å' | 'Æ' | 'Ç' | 'È' | 'É' | 'Ê' | 'Ë') [ø-˿]) | (&('A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' | 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z') [A-Z]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'g' | 'h' | 'i' | 'j' | 'k' | 'l' | 'm' | 'n' | 'o' | 'p' | 'q' | 'r' | 's' | 't' | 'u' | 'v' | 'w' | 'x' | 'y' | 'z') [a-z])))> */
		func() bool {
			position1709, tokenIndex1709, depth1709 := position, tokenIndex, depth
			{
				position1710 := position
				depth++
				{
					position1711, tokenIndex1711, depth1711 := position, tokenIndex, depth
					if c := buffer[position]; c < rune('À') || c > rune('Ö') {
						goto l1712
					}
					position++
					goto l1711
				l1712:
					position, tokenIndex, depth = position1711, tokenIndex1711, depth1711
					if c := buffer[position]; c < rune('Ø') || c > rune('ö') {
						goto l1713
					}
					position++
					goto l1711
				l1713:
					position, tokenIndex, depth = position1711, tokenIndex1711, depth1711
					if c := buffer[position]; c < rune('Ͱ') || c > rune('ͽ') {
						goto l1714
					}
					position++
					goto l1711
				l1714:
					position, tokenIndex, depth = position1711, tokenIndex1711, depth1711
					if c := buffer[position]; c < rune('\u200c') || c > rune('\u200d') {
						goto l1715
					}
					position++
					goto l1711
				l1715:
					position, tokenIndex, depth = position1711, tokenIndex1711, depth1711
					if c := buffer[position]; c < rune('⁰') || c > rune('\u218f') {
						goto l1716
					}
					position++
					goto l1711
				l1716:
					position, tokenIndex, depth = position1711, tokenIndex1711, depth1711
					if c := buffer[position]; c < rune('豈') || c > rune('\ufdcf') {
						goto l1717
					}
					position++
					goto l1711
				l1717:
					position, tokenIndex, depth = position1711, tokenIndex1711, depth1711
					{
						switch buffer[position] {
						case 'ð', 'ñ', 'ò', 'ó':
							if c := buffer[position]; c < rune('𐀀') || c > rune('\U000effff') {
								goto l1709
							}
							position++
							break
						case 'ï':
							if c := buffer[position]; c < rune('ﷰ') || c > rune('�') {
								goto l1709
							}
							position++
							break
						case 'ã', 'ä', 'å', 'æ', 'ç', 'è', 'é', 'ê', 'ë', 'ì', 'í':
							if c := buffer[position]; c < rune('、') || c > rune('\ud7ff') {
								goto l1709
							}
							position++
							break
						case 'â':
							if c := buffer[position]; c < rune('Ⰰ') || c > rune('\u2fef') {
								goto l1709
							}
							position++
							break
						case 'Í', 'Î', 'Ï', 'Ð', 'Ñ', 'Ò', 'Ó', 'Ô', 'Õ', 'Ö', '×', 'Ø', 'Ù', 'Ú', 'Û', 'Ü', 'Ý', 'Þ', 'ß', 'à', 'á':
							if c := buffer[position]; c < rune('\u037f') || c > rune('\u1fff') {
								goto l1709
							}
							position++
							break
						case 'Ã', 'Ä', 'Å', 'Æ', 'Ç', 'È', 'É', 'Ê', 'Ë':
							if c := buffer[position]; c < rune('ø') || c > rune('˿') {
								goto l1709
							}
							position++
							break
						case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
							if c := buffer[position]; c < rune('A') || c > rune('Z') {
								goto l1709
							}
							position++
							break
						default:
							if c := buffer[position]; c < rune('a') || c > rune('z') {
								goto l1709
							}
							position++
							break
						}
					}

				}
			l1711:
				depth--
				add(rulepnCharsBase, position1710)
			}
			return true
		l1709:
			position, tokenIndex, depth = position1709, tokenIndex1709, depth1709
			return false
		},
		/* 94 plx <- <(percent / pnLocalEsc)> */
		nil,
		/* 95 percent <- <('%' hex hex)> */
		nil,
		/* 96 hex <- <((&('a' | 'b' | 'c' | 'd' | 'e' | 'f') [a-f]) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]) | (&('A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' | 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z') [A-Z]))> */
		func() bool {
			position1721, tokenIndex1721, depth1721 := position, tokenIndex, depth
			{
				position1722 := position
				depth++
				{
					switch buffer[position] {
					case 'a', 'b', 'c', 'd', 'e', 'f':
						if c := buffer[position]; c < rune('a') || c > rune('f') {
							goto l1721
						}
						position++
						break
					case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1721
						}
						position++
						break
					default:
						if c := buffer[position]; c < rune('A') || c > rune('Z') {
							goto l1721
						}
						position++
						break
					}
				}

				depth--
				add(rulehex, position1722)
			}
			return true
		l1721:
			position, tokenIndex, depth = position1721, tokenIndex1721, depth1721
			return false
		},
		/* 97 pnLocalEsc <- <('\\' ((&('%') '%') | (&('@') '@') | (&('#') '#') | (&('?') '?') | (&('/') '/') | (&('=') '=') | (&(';') ';') | (&(',') ',') | (&('+') '+') | (&('*') '*') | (&(')') ')') | (&('(') '(') | (&('\'') '\'') | (&('&') '&') | (&('$') '$') | (&('!') '!') | (&('-') '-') | (&('.') '.') | (&('~') '~') | (&('_') '_')))> */
		nil,
		/* 98 PREFIX <- <(('p' / 'P') ('r' / 'R') ('e' / 'E') ('f' / 'F') ('i' / 'I') ('x' / 'X') skip)> */
		nil,
		/* 99 TRUE <- <(('t' / 'T') ('r' / 'R') ('u' / 'U') ('e' / 'E') skip)> */
		nil,
		/* 100 FALSE <- <(('f' / 'F') ('a' / 'A') ('l' / 'L') ('s' / 'S') ('e' / 'E') skip)> */
		nil,
		/* 101 BASE <- <(('b' / 'B') ('a' / 'A') ('s' / 'S') ('e' / 'E') skip)> */
		nil,
		/* 102 SELECT <- <(('s' / 'S') ('e' / 'E') ('l' / 'L') ('e' / 'E') ('c' / 'C') ('t' / 'T') skip)> */
		nil,
		/* 103 REDUCED <- <(('r' / 'R') ('e' / 'E') ('d' / 'D') ('u' / 'U') ('c' / 'C') ('e' / 'E') ('d' / 'D') skip)> */
		nil,
		/* 104 DISTINCT <- <(('d' / 'D') ('i' / 'I') ('s' / 'S') ('t' / 'T') ('i' / 'I') ('n' / 'N') ('c' / 'C') ('t' / 'T') skip)> */
		func() bool {
			position1731, tokenIndex1731, depth1731 := position, tokenIndex, depth
			{
				position1732 := position
				depth++
				{
					position1733, tokenIndex1733, depth1733 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l1734
					}
					position++
					goto l1733
				l1734:
					position, tokenIndex, depth = position1733, tokenIndex1733, depth1733
					if buffer[position] != rune('D') {
						goto l1731
					}
					position++
				}
			l1733:
				{
					position1735, tokenIndex1735, depth1735 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l1736
					}
					position++
					goto l1735
				l1736:
					position, tokenIndex, depth = position1735, tokenIndex1735, depth1735
					if buffer[position] != rune('I') {
						goto l1731
					}
					position++
				}
			l1735:
				{
					position1737, tokenIndex1737, depth1737 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l1738
					}
					position++
					goto l1737
				l1738:
					position, tokenIndex, depth = position1737, tokenIndex1737, depth1737
					if buffer[position] != rune('S') {
						goto l1731
					}
					position++
				}
			l1737:
				{
					position1739, tokenIndex1739, depth1739 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l1740
					}
					position++
					goto l1739
				l1740:
					position, tokenIndex, depth = position1739, tokenIndex1739, depth1739
					if buffer[position] != rune('T') {
						goto l1731
					}
					position++
				}
			l1739:
				{
					position1741, tokenIndex1741, depth1741 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l1742
					}
					position++
					goto l1741
				l1742:
					position, tokenIndex, depth = position1741, tokenIndex1741, depth1741
					if buffer[position] != rune('I') {
						goto l1731
					}
					position++
				}
			l1741:
				{
					position1743, tokenIndex1743, depth1743 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l1744
					}
					position++
					goto l1743
				l1744:
					position, tokenIndex, depth = position1743, tokenIndex1743, depth1743
					if buffer[position] != rune('N') {
						goto l1731
					}
					position++
				}
			l1743:
				{
					position1745, tokenIndex1745, depth1745 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l1746
					}
					position++
					goto l1745
				l1746:
					position, tokenIndex, depth = position1745, tokenIndex1745, depth1745
					if buffer[position] != rune('C') {
						goto l1731
					}
					position++
				}
			l1745:
				{
					position1747, tokenIndex1747, depth1747 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l1748
					}
					position++
					goto l1747
				l1748:
					position, tokenIndex, depth = position1747, tokenIndex1747, depth1747
					if buffer[position] != rune('T') {
						goto l1731
					}
					position++
				}
			l1747:
				if !rules[ruleskip]() {
					goto l1731
				}
				depth--
				add(ruleDISTINCT, position1732)
			}
			return true
		l1731:
			position, tokenIndex, depth = position1731, tokenIndex1731, depth1731
			return false
		},
		/* 105 FROM <- <(('f' / 'F') ('r' / 'R') ('o' / 'O') ('m' / 'M') skip)> */
		nil,
		/* 106 NAMED <- <(('n' / 'N') ('a' / 'A') ('m' / 'M') ('e' / 'E') ('d' / 'D') skip)> */
		nil,
		/* 107 WHERE <- <(('w' / 'W') ('h' / 'H') ('e' / 'E') ('r' / 'R') ('e' / 'E') skip)> */
		nil,
		/* 108 LBRACE <- <('{' skip)> */
		func() bool {
			position1752, tokenIndex1752, depth1752 := position, tokenIndex, depth
			{
				position1753 := position
				depth++
				if buffer[position] != rune('{') {
					goto l1752
				}
				position++
				if !rules[ruleskip]() {
					goto l1752
				}
				depth--
				add(ruleLBRACE, position1753)
			}
			return true
		l1752:
			position, tokenIndex, depth = position1752, tokenIndex1752, depth1752
			return false
		},
		/* 109 RBRACE <- <('}' skip)> */
		func() bool {
			position1754, tokenIndex1754, depth1754 := position, tokenIndex, depth
			{
				position1755 := position
				depth++
				if buffer[position] != rune('}') {
					goto l1754
				}
				position++
				if !rules[ruleskip]() {
					goto l1754
				}
				depth--
				add(ruleRBRACE, position1755)
			}
			return true
		l1754:
			position, tokenIndex, depth = position1754, tokenIndex1754, depth1754
			return false
		},
		/* 110 LBRACK <- <('[' skip)> */
		nil,
		/* 111 RBRACK <- <(']' skip)> */
		nil,
		/* 112 SEMICOLON <- <(';' skip)> */
		func() bool {
			position1758, tokenIndex1758, depth1758 := position, tokenIndex, depth
			{
				position1759 := position
				depth++
				if buffer[position] != rune(';') {
					goto l1758
				}
				position++
				if !rules[ruleskip]() {
					goto l1758
				}
				depth--
				add(ruleSEMICOLON, position1759)
			}
			return true
		l1758:
			position, tokenIndex, depth = position1758, tokenIndex1758, depth1758
			return false
		},
		/* 113 COMMA <- <(',' skip)> */
		func() bool {
			position1760, tokenIndex1760, depth1760 := position, tokenIndex, depth
			{
				position1761 := position
				depth++
				if buffer[position] != rune(',') {
					goto l1760
				}
				position++
				if !rules[ruleskip]() {
					goto l1760
				}
				depth--
				add(ruleCOMMA, position1761)
			}
			return true
		l1760:
			position, tokenIndex, depth = position1760, tokenIndex1760, depth1760
			return false
		},
		/* 114 DOT <- <('.' skip)> */
		func() bool {
			position1762, tokenIndex1762, depth1762 := position, tokenIndex, depth
			{
				position1763 := position
				depth++
				if buffer[position] != rune('.') {
					goto l1762
				}
				position++
				if !rules[ruleskip]() {
					goto l1762
				}
				depth--
				add(ruleDOT, position1763)
			}
			return true
		l1762:
			position, tokenIndex, depth = position1762, tokenIndex1762, depth1762
			return false
		},
		/* 115 COLON <- <(':' skip)> */
		nil,
		/* 116 PIPE <- <('|' skip)> */
		func() bool {
			position1765, tokenIndex1765, depth1765 := position, tokenIndex, depth
			{
				position1766 := position
				depth++
				if buffer[position] != rune('|') {
					goto l1765
				}
				position++
				if !rules[ruleskip]() {
					goto l1765
				}
				depth--
				add(rulePIPE, position1766)
			}
			return true
		l1765:
			position, tokenIndex, depth = position1765, tokenIndex1765, depth1765
			return false
		},
		/* 117 SLASH <- <('/' skip)> */
		func() bool {
			position1767, tokenIndex1767, depth1767 := position, tokenIndex, depth
			{
				position1768 := position
				depth++
				if buffer[position] != rune('/') {
					goto l1767
				}
				position++
				if !rules[ruleskip]() {
					goto l1767
				}
				depth--
				add(ruleSLASH, position1768)
			}
			return true
		l1767:
			position, tokenIndex, depth = position1767, tokenIndex1767, depth1767
			return false
		},
		/* 118 INVERSE <- <('^' skip)> */
		func() bool {
			position1769, tokenIndex1769, depth1769 := position, tokenIndex, depth
			{
				position1770 := position
				depth++
				if buffer[position] != rune('^') {
					goto l1769
				}
				position++
				if !rules[ruleskip]() {
					goto l1769
				}
				depth--
				add(ruleINVERSE, position1770)
			}
			return true
		l1769:
			position, tokenIndex, depth = position1769, tokenIndex1769, depth1769
			return false
		},
		/* 119 LPAREN <- <('(' skip)> */
		func() bool {
			position1771, tokenIndex1771, depth1771 := position, tokenIndex, depth
			{
				position1772 := position
				depth++
				if buffer[position] != rune('(') {
					goto l1771
				}
				position++
				if !rules[ruleskip]() {
					goto l1771
				}
				depth--
				add(ruleLPAREN, position1772)
			}
			return true
		l1771:
			position, tokenIndex, depth = position1771, tokenIndex1771, depth1771
			return false
		},
		/* 120 RPAREN <- <(')' skip)> */
		func() bool {
			position1773, tokenIndex1773, depth1773 := position, tokenIndex, depth
			{
				position1774 := position
				depth++
				if buffer[position] != rune(')') {
					goto l1773
				}
				position++
				if !rules[ruleskip]() {
					goto l1773
				}
				depth--
				add(ruleRPAREN, position1774)
			}
			return true
		l1773:
			position, tokenIndex, depth = position1773, tokenIndex1773, depth1773
			return false
		},
		/* 121 ISA <- <('a' Action6 skip)> */
		func() bool {
			position1775, tokenIndex1775, depth1775 := position, tokenIndex, depth
			{
				position1776 := position
				depth++
				if buffer[position] != rune('a') {
					goto l1775
				}
				position++
				{
					add(ruleAction6, position)
				}
				if !rules[ruleskip]() {
					goto l1775
				}
				depth--
				add(ruleISA, position1776)
			}
			return true
		l1775:
			position, tokenIndex, depth = position1775, tokenIndex1775, depth1775
			return false
		},
		/* 122 NOT <- <('!' skip)> */
		func() bool {
			position1778, tokenIndex1778, depth1778 := position, tokenIndex, depth
			{
				position1779 := position
				depth++
				if buffer[position] != rune('!') {
					goto l1778
				}
				position++
				if !rules[ruleskip]() {
					goto l1778
				}
				depth--
				add(ruleNOT, position1779)
			}
			return true
		l1778:
			position, tokenIndex, depth = position1778, tokenIndex1778, depth1778
			return false
		},
		/* 123 STAR <- <('*' skip)> */
		func() bool {
			position1780, tokenIndex1780, depth1780 := position, tokenIndex, depth
			{
				position1781 := position
				depth++
				if buffer[position] != rune('*') {
					goto l1780
				}
				position++
				if !rules[ruleskip]() {
					goto l1780
				}
				depth--
				add(ruleSTAR, position1781)
			}
			return true
		l1780:
			position, tokenIndex, depth = position1780, tokenIndex1780, depth1780
			return false
		},
		/* 124 QUESTION <- <('?' skip)> */
		nil,
		/* 125 PLUS <- <('+' skip)> */
		func() bool {
			position1783, tokenIndex1783, depth1783 := position, tokenIndex, depth
			{
				position1784 := position
				depth++
				if buffer[position] != rune('+') {
					goto l1783
				}
				position++
				if !rules[ruleskip]() {
					goto l1783
				}
				depth--
				add(rulePLUS, position1784)
			}
			return true
		l1783:
			position, tokenIndex, depth = position1783, tokenIndex1783, depth1783
			return false
		},
		/* 126 MINUS <- <('-' skip)> */
		func() bool {
			position1785, tokenIndex1785, depth1785 := position, tokenIndex, depth
			{
				position1786 := position
				depth++
				if buffer[position] != rune('-') {
					goto l1785
				}
				position++
				if !rules[ruleskip]() {
					goto l1785
				}
				depth--
				add(ruleMINUS, position1786)
			}
			return true
		l1785:
			position, tokenIndex, depth = position1785, tokenIndex1785, depth1785
			return false
		},
		/* 127 OPTIONAL <- <(('o' / 'O') ('p' / 'P') ('t' / 'T') ('i' / 'I') ('o' / 'O') ('n' / 'N') ('a' / 'A') ('l' / 'L') skip)> */
		nil,
		/* 128 UNION <- <(('u' / 'U') ('n' / 'N') ('i' / 'I') ('o' / 'O') ('n' / 'N') skip)> */
		nil,
		/* 129 LIMIT <- <(('l' / 'L') ('i' / 'I') ('m' / 'M') ('i' / 'I') ('t' / 'T') skip)> */
		nil,
		/* 130 OFFSET <- <(('o' / 'O') ('f' / 'F') ('f' / 'F') ('s' / 'S') ('e' / 'E') ('t' / 'T') skip)> */
		nil,
		/* 131 INTEGER <- <([0-9]+ skip)> */
		func() bool {
			position1791, tokenIndex1791, depth1791 := position, tokenIndex, depth
			{
				position1792 := position
				depth++
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l1791
				}
				position++
			l1793:
				{
					position1794, tokenIndex1794, depth1794 := position, tokenIndex, depth
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1794
					}
					position++
					goto l1793
				l1794:
					position, tokenIndex, depth = position1794, tokenIndex1794, depth1794
				}
				if !rules[ruleskip]() {
					goto l1791
				}
				depth--
				add(ruleINTEGER, position1792)
			}
			return true
		l1791:
			position, tokenIndex, depth = position1791, tokenIndex1791, depth1791
			return false
		},
		/* 132 CONSTRUCT <- <(('c' / 'C') ('o' / 'O') ('n' / 'N') ('s' / 'S') ('t' / 'T') ('r' / 'R') ('u' / 'U') ('c' / 'C') ('t' / 'T') skip)> */
		nil,
		/* 133 DESCRIBE <- <(('d' / 'D') ('e' / 'E') ('s' / 'S') ('c' / 'C') ('r' / 'R') ('i' / 'I') ('b' / 'B') ('e' / 'E') skip)> */
		nil,
		/* 134 ASK <- <(('a' / 'A') ('s' / 'S') ('k' / 'K') skip)> */
		nil,
		/* 135 OR <- <('|' '|' skip)> */
		nil,
		/* 136 AND <- <('&' '&' skip)> */
		nil,
		/* 137 EQ <- <('=' skip)> */
		func() bool {
			position1800, tokenIndex1800, depth1800 := position, tokenIndex, depth
			{
				position1801 := position
				depth++
				if buffer[position] != rune('=') {
					goto l1800
				}
				position++
				if !rules[ruleskip]() {
					goto l1800
				}
				depth--
				add(ruleEQ, position1801)
			}
			return true
		l1800:
			position, tokenIndex, depth = position1800, tokenIndex1800, depth1800
			return false
		},
		/* 138 NE <- <('!' '=' skip)> */
		nil,
		/* 139 GT <- <('>' skip)> */
		nil,
		/* 140 LT <- <('<' skip)> */
		nil,
		/* 141 LE <- <('<' '=' skip)> */
		nil,
		/* 142 GE <- <('>' '=' skip)> */
		nil,
		/* 143 IN <- <(('i' / 'I') ('n' / 'N') skip)> */
		nil,
		/* 144 NOTIN <- <(('n' / 'N') ('o' / 'O') ('t' / 'T') ' ' ('i' / 'I') ('n' / 'N') skip)> */
		nil,
		/* 145 AS <- <(('a' / 'A') ('s' / 'S') skip)> */
		func() bool {
			position1809, tokenIndex1809, depth1809 := position, tokenIndex, depth
			{
				position1810 := position
				depth++
				{
					position1811, tokenIndex1811, depth1811 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l1812
					}
					position++
					goto l1811
				l1812:
					position, tokenIndex, depth = position1811, tokenIndex1811, depth1811
					if buffer[position] != rune('A') {
						goto l1809
					}
					position++
				}
			l1811:
				{
					position1813, tokenIndex1813, depth1813 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l1814
					}
					position++
					goto l1813
				l1814:
					position, tokenIndex, depth = position1813, tokenIndex1813, depth1813
					if buffer[position] != rune('S') {
						goto l1809
					}
					position++
				}
			l1813:
				if !rules[ruleskip]() {
					goto l1809
				}
				depth--
				add(ruleAS, position1810)
			}
			return true
		l1809:
			position, tokenIndex, depth = position1809, tokenIndex1809, depth1809
			return false
		},
		/* 146 STR <- <(('s' / 'S') ('t' / 'T') ('r' / 'R') skip)> */
		nil,
		/* 147 LANG <- <(('l' / 'L') ('a' / 'A') ('n' / 'N') ('g' / 'G') skip)> */
		nil,
		/* 148 DATATYPE <- <(('d' / 'D') ('a' / 'A') ('t' / 'T') ('a' / 'A') ('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E') skip)> */
		nil,
		/* 149 IRI <- <(('i' / 'I') ('r' / 'R') ('i' / 'I') skip)> */
		nil,
		/* 150 URI <- <(('u' / 'U') ('r' / 'R') ('i' / 'I') skip)> */
		nil,
		/* 151 ABS <- <(('a' / 'A') ('b' / 'B') ('s' / 'S') skip)> */
		nil,
		/* 152 CEIL <- <(('c' / 'C') ('e' / 'E') ('i' / 'I') ('l' / 'L') skip)> */
		nil,
		/* 153 ROUND <- <(('r' / 'R') ('o' / 'O') ('u' / 'U') ('n' / 'N') ('d' / 'D') skip)> */
		nil,
		/* 154 FLOOR <- <(('f' / 'F') ('l' / 'L') ('o' / 'O') ('o' / 'O') ('r' / 'R') skip)> */
		nil,
		/* 155 STRLEN <- <(('s' / 'S') ('t' / 'T') ('r' / 'R') ('l' / 'L') ('e' / 'E') ('n' / 'N') skip)> */
		nil,
		/* 156 UCASE <- <(('u' / 'U') ('c' / 'C') ('a' / 'A') ('s' / 'S') ('e' / 'E') skip)> */
		nil,
		/* 157 LCASE <- <(('l' / 'L') ('c' / 'C') ('a' / 'A') ('s' / 'S') ('e' / 'E') skip)> */
		nil,
		/* 158 ENCODEFORURI <- <(('e' / 'E') ('n' / 'N') ('c' / 'C') ('o' / 'O') ('d' / 'D') ('e' / 'E') '_' ('f' / 'F') ('o' / 'O') ('r' / 'R') '_' ('u' / 'U') ('r' / 'R') ('i' / 'I') skip)> */
		nil,
		/* 159 YEAR <- <(('y' / 'Y') ('e' / 'E') ('a' / 'A') ('r' / 'R') skip)> */
		nil,
		/* 160 MONTH <- <(('m' / 'M') ('o' / 'O') ('n' / 'N') ('t' / 'T') ('h' / 'H') skip)> */
		nil,
		/* 161 DAY <- <(('d' / 'D') ('a' / 'A') ('y' / 'Y') skip)> */
		nil,
		/* 162 HOURS <- <(('h' / 'H') ('o' / 'O') ('u' / 'U') ('r' / 'R') ('s' / 'S') skip)> */
		nil,
		/* 163 MINUTES <- <(('m' / 'M') ('i' / 'I') ('n' / 'N') ('u' / 'U') ('t' / 'T') ('e' / 'E') ('s' / 'S') skip)> */
		nil,
		/* 164 SECONDS <- <(('s' / 'S') ('e' / 'E') ('c' / 'C') ('o' / 'O') ('n' / 'N') ('d' / 'D') ('s' / 'S') skip)> */
		nil,
		/* 165 TIMEZONE <- <(('t' / 'T') ('i' / 'I') ('m' / 'M') ('e' / 'E') ('z' / 'Z') ('o' / 'O') ('n' / 'N') ('e' / 'E') skip)> */
		nil,
		/* 166 TZ <- <(('t' / 'T') ('z' / 'Z') skip)> */
		nil,
		/* 167 MD5 <- <(('m' / 'M') ('d' / 'D') skip)> */
		nil,
		/* 168 SHA1 <- <(('s' / 'S') ('h' / 'H') ('a' / 'A') '1' skip)> */
		nil,
		/* 169 SHA256 <- <(('s' / 'S') ('h' / 'H') ('a' / 'A') '2' '5' '6' skip)> */
		nil,
		/* 170 SHA384 <- <(('s' / 'S') ('h' / 'H') ('a' / 'A') '3' '8' '4' skip)> */
		nil,
		/* 171 SHA512 <- <(('s' / 'S') ('h' / 'H') ('a' / 'A') '5' '1' '2' skip)> */
		nil,
		/* 172 ISIRI <- <(('i' / 'I') ('s' / 'S') ('i' / 'I') ('r' / 'R') ('i' / 'I') skip)> */
		nil,
		/* 173 ISURI <- <(('i' / 'I') ('s' / 'S') ('u' / 'U') ('r' / 'R') ('i' / 'I') skip)> */
		nil,
		/* 174 ISBLANK <- <(('i' / 'I') ('s' / 'S') ('b' / 'B') ('l' / 'L') ('a' / 'A') ('n' / 'N') ('k' / 'K') skip)> */
		nil,
		/* 175 ISLITERAL <- <(('i' / 'I') ('s' / 'S') ('l' / 'L') ('i' / 'I') ('t' / 'T') ('e' / 'E') ('r' / 'R') ('a' / 'A') ('l' / 'L') skip)> */
		nil,
		/* 176 ISNUMERIC <- <(('i' / 'I') ('s' / 'S') ('n' / 'N') ('u' / 'U') ('m' / 'M') ('e' / 'E') ('r' / 'R') ('i' / 'I') ('c' / 'C') skip)> */
		nil,
		/* 177 LANGMATCHES <- <(('l' / 'L') ('a' / 'A') ('n' / 'N') ('g' / 'G') ('m' / 'M') ('a' / 'A') ('t' / 'T') ('c' / 'C') ('h' / 'H') ('e' / 'E') ('s' / 'S') skip)> */
		nil,
		/* 178 CONTAINS <- <(('c' / 'C') ('o' / 'O') ('n' / 'N') ('t' / 'T') ('a' / 'A') ('i' / 'I') ('n' / 'N') ('s' / 'S') skip)> */
		nil,
		/* 179 STRSTARTS <- <(('s' / 'S') ('t' / 'T') ('r' / 'R') ('s' / 'S') ('t' / 'T') ('a' / 'A') ('r' / 'R') ('t' / 'T') ('s' / 'S') skip)> */
		nil,
		/* 180 STRENDS <- <(('s' / 'S') ('t' / 'T') ('r' / 'R') ('e' / 'E') ('n' / 'N') ('d' / 'D') ('s' / 'S') skip)> */
		nil,
		/* 181 STRBEFORE <- <(('s' / 'S') ('t' / 'T') ('r' / 'R') ('b' / 'B') ('e' / 'E') ('f' / 'F') ('o' / 'O') ('r' / 'R') ('e' / 'E') skip)> */
		nil,
		/* 182 STRAFTER <- <(('s' / 'S') ('t' / 'T') ('r' / 'R') ('a' / 'A') ('f' / 'F') ('t' / 'T') ('e' / 'E') ('r' / 'R') skip)> */
		nil,
		/* 183 STRLANG <- <(('s' / 'S') ('t' / 'T') ('r' / 'R') ('l' / 'L') ('a' / 'A') ('n' / 'N') ('g' / 'G') skip)> */
		nil,
		/* 184 STRDT <- <(('s' / 'S') ('t' / 'T') ('r' / 'R') ('d' / 'D') ('t' / 'T') skip)> */
		nil,
		/* 185 SAMETERM <- <(('s' / 'S') ('a' / 'A') ('m' / 'M') ('e' / 'E') ('t' / 'T') ('e' / 'E') ('r' / 'R') ('m' / 'M') skip)> */
		nil,
		/* 186 BOUND <- <(('b' / 'B') ('o' / 'O') ('u' / 'U') ('n' / 'N') ('d' / 'D') skip)> */
		nil,
		/* 187 BNODE <- <(('b' / 'B') ('n' / 'N') ('o' / 'O') ('d' / 'D') ('e' / 'E') skip)> */
		nil,
		/* 188 RAND <- <(('r' / 'R') ('a' / 'A') ('n' / 'N') ('d' / 'D') skip)> */
		nil,
		/* 189 NOW <- <(('n' / 'N') ('o' / 'O') ('w' / 'W') skip)> */
		nil,
		/* 190 UUID <- <(('u' / 'U') ('u' / 'U') ('i' / 'I') ('d' / 'D') skip)> */
		nil,
		/* 191 STRUUID <- <(('s' / 'S') ('t' / 'T') ('r' / 'R') ('u' / 'U') ('u' / 'U') ('i' / 'I') ('d' / 'D') skip)> */
		nil,
		/* 192 CONCAT <- <(('c' / 'C') ('o' / 'O') ('n' / 'N') ('c' / 'C') ('a' / 'A') ('t' / 'T') skip)> */
		nil,
		/* 193 SUBSTR <- <(('s' / 'S') ('u' / 'U') ('b' / 'B') ('s' / 'S') ('t' / 'T') ('r' / 'R') skip)> */
		nil,
		/* 194 REPLACE <- <(('r' / 'R') ('e' / 'E') ('p' / 'P') ('l' / 'L') ('a' / 'A') ('c' / 'C') ('e' / 'E') skip)> */
		nil,
		/* 195 REGEX <- <(('r' / 'R') ('e' / 'E') ('g' / 'G') ('e' / 'E') ('x' / 'X') skip)> */
		nil,
		/* 196 IF <- <(('i' / 'I') ('f' / 'F') skip)> */
		nil,
		/* 197 EXISTS <- <(('e' / 'E') ('x' / 'X') ('i' / 'I') ('s' / 'S') ('t' / 'T') ('s' / 'S') skip)> */
		nil,
		/* 198 NOTEXIST <- <(('n' / 'N') ('o' / 'O') ('t' / 'T') ' ' ('e' / 'E') ('x' / 'X') ('i' / 'I') ('s' / 'S') ('t' / 'T') ('s' / 'S') skip)> */
		nil,
		/* 199 COALESCE <- <(('c' / 'C') ('o' / 'O') ('a' / 'A') ('l' / 'L') ('e' / 'E') ('s' / 'S') ('c' / 'C') ('e' / 'E') skip)> */
		nil,
		/* 200 FILTER <- <(('f' / 'F') ('i' / 'I') ('l' / 'L') ('t' / 'T') ('e' / 'E') ('r' / 'R') skip)> */
		nil,
		/* 201 BIND <- <(('b' / 'B') ('i' / 'I') ('n' / 'N') ('d' / 'D') skip)> */
		nil,
		/* 202 SUM <- <(('s' / 'S') ('u' / 'U') ('m' / 'M') skip)> */
		nil,
		/* 203 MIN <- <(('m' / 'M') ('i' / 'I') ('n' / 'N') skip)> */
		nil,
		/* 204 MAX <- <(('m' / 'M') ('a' / 'A') ('x' / 'X') skip)> */
		nil,
		/* 205 AVG <- <(('a' / 'A') ('v' / 'V') ('g' / 'G') skip)> */
		nil,
		/* 206 SAMPLE <- <(('s' / 'S') ('a' / 'A') ('m' / 'M') ('p' / 'P') ('l' / 'L') ('e' / 'E') skip)> */
		nil,
		/* 207 COUNT <- <(('c' / 'C') ('o' / 'O') ('u' / 'U') ('n' / 'N') ('t' / 'T') skip)> */
		nil,
		/* 208 GROUPCONCAT <- <(('g' / 'G') ('r' / 'R') ('o' / 'O') ('u' / 'U') ('p' / 'P') '_' ('c' / 'C') ('o' / 'O') ('n' / 'N') ('c' / 'C') ('a' / 'A') ('t' / 'T') skip)> */
		nil,
		/* 209 SEPARATOR <- <(('s' / 'S') ('e' / 'E') ('p' / 'P') ('a' / 'A') ('r' / 'R') ('a' / 'A') ('t' / 'T') ('o' / 'O') ('r' / 'R') skip)> */
		nil,
		/* 210 ASC <- <(('a' / 'A') ('s' / 'S') ('c' / 'C') skip)> */
		nil,
		/* 211 DESC <- <(('d' / 'D') ('e' / 'E') ('s' / 'S') ('c' / 'C') skip)> */
		nil,
		/* 212 ORDER <- <(('o' / 'O') ('r' / 'R') ('d' / 'D') ('e' / 'E') ('r' / 'R') skip)> */
		nil,
		/* 213 GROUP <- <(('g' / 'G') ('r' / 'R') ('o' / 'O') ('u' / 'U') ('p' / 'P') skip)> */
		nil,
		/* 214 BY <- <(('b' / 'B') ('y' / 'Y') skip)> */
		func() bool {
			position1883, tokenIndex1883, depth1883 := position, tokenIndex, depth
			{
				position1884 := position
				depth++
				{
					position1885, tokenIndex1885, depth1885 := position, tokenIndex, depth
					if buffer[position] != rune('b') {
						goto l1886
					}
					position++
					goto l1885
				l1886:
					position, tokenIndex, depth = position1885, tokenIndex1885, depth1885
					if buffer[position] != rune('B') {
						goto l1883
					}
					position++
				}
			l1885:
				{
					position1887, tokenIndex1887, depth1887 := position, tokenIndex, depth
					if buffer[position] != rune('y') {
						goto l1888
					}
					position++
					goto l1887
				l1888:
					position, tokenIndex, depth = position1887, tokenIndex1887, depth1887
					if buffer[position] != rune('Y') {
						goto l1883
					}
					position++
				}
			l1887:
				if !rules[ruleskip]() {
					goto l1883
				}
				depth--
				add(ruleBY, position1884)
			}
			return true
		l1883:
			position, tokenIndex, depth = position1883, tokenIndex1883, depth1883
			return false
		},
		/* 215 HAVING <- <(('h' / 'H') ('a' / 'A') ('v' / 'V') ('i' / 'I') ('n' / 'N') ('g' / 'G') skip)> */
		nil,
		/* 216 GRAPH <- <(('g' / 'G') ('r' / 'R') ('a' / 'A') ('p' / 'P') ('h' / 'H') skip)> */
		nil,
		/* 217 MINUSSETOPER <- <(('m' / 'M') ('i' / 'I') ('n' / 'N') ('u' / 'U') ('s' / 'S') skip)> */
		nil,
		/* 218 SERVICE <- <(('s' / 'S') ('e' / 'E') ('r' / 'R') ('v' / 'V') ('i' / 'I') ('c' / 'C') ('e' / 'E') skip)> */
		nil,
		/* 219 SILENT <- <(('s' / 'S') ('i' / 'I') ('l' / 'L') ('e' / 'E') ('n' / 'N') ('t' / 'T') skip)> */
		nil,
		/* 220 skip <- <(ws / comment)*> */
		func() bool {
			{
				position1895 := position
				depth++
			l1896:
				{
					position1897, tokenIndex1897, depth1897 := position, tokenIndex, depth
					{
						position1898, tokenIndex1898, depth1898 := position, tokenIndex, depth
						if !rules[rulews]() {
							goto l1899
						}
						goto l1898
					l1899:
						position, tokenIndex, depth = position1898, tokenIndex1898, depth1898
						{
							position1900 := position
							depth++
							if buffer[position] != rune('#') {
								goto l1897
							}
							position++
						l1901:
							{
								position1902, tokenIndex1902, depth1902 := position, tokenIndex, depth
								{
									position1903, tokenIndex1903, depth1903 := position, tokenIndex, depth
									if !rules[ruleendOfLine]() {
										goto l1903
									}
									goto l1902
								l1903:
									position, tokenIndex, depth = position1903, tokenIndex1903, depth1903
								}
								if !matchDot() {
									goto l1902
								}
								goto l1901
							l1902:
								position, tokenIndex, depth = position1902, tokenIndex1902, depth1902
							}
							if !rules[ruleendOfLine]() {
								goto l1897
							}
							depth--
							add(rulecomment, position1900)
						}
					}
				l1898:
					goto l1896
				l1897:
					position, tokenIndex, depth = position1897, tokenIndex1897, depth1897
				}
				depth--
				add(ruleskip, position1895)
			}
			return true
		},
		/* 221 ws <- <((&('\v') '\v') | (&('\f') '\f') | (&('\t') '\t') | (&(' ') ' ') | (&('\n' | '\r') endOfLine))> */
		func() bool {
			position1904, tokenIndex1904, depth1904 := position, tokenIndex, depth
			{
				position1905 := position
				depth++
				{
					switch buffer[position] {
					case '\v':
						if buffer[position] != rune('\v') {
							goto l1904
						}
						position++
						break
					case '\f':
						if buffer[position] != rune('\f') {
							goto l1904
						}
						position++
						break
					case '\t':
						if buffer[position] != rune('\t') {
							goto l1904
						}
						position++
						break
					case ' ':
						if buffer[position] != rune(' ') {
							goto l1904
						}
						position++
						break
					default:
						if !rules[ruleendOfLine]() {
							goto l1904
						}
						break
					}
				}

				depth--
				add(rulews, position1905)
			}
			return true
		l1904:
			position, tokenIndex, depth = position1904, tokenIndex1904, depth1904
			return false
		},
		/* 222 comment <- <('#' (!endOfLine .)* endOfLine)> */
		nil,
		/* 223 endOfLine <- <(('\r' '\n') / '\n' / '\r')> */
		func() bool {
			position1908, tokenIndex1908, depth1908 := position, tokenIndex, depth
			{
				position1909 := position
				depth++
				{
					position1910, tokenIndex1910, depth1910 := position, tokenIndex, depth
					if buffer[position] != rune('\r') {
						goto l1911
					}
					position++
					if buffer[position] != rune('\n') {
						goto l1911
					}
					position++
					goto l1910
				l1911:
					position, tokenIndex, depth = position1910, tokenIndex1910, depth1910
					if buffer[position] != rune('\n') {
						goto l1912
					}
					position++
					goto l1910
				l1912:
					position, tokenIndex, depth = position1910, tokenIndex1910, depth1910
					if buffer[position] != rune('\r') {
						goto l1908
					}
					position++
				}
			l1910:
				depth--
				add(ruleendOfLine, position1909)
			}
			return true
		l1908:
			position, tokenIndex, depth = position1908, tokenIndex1908, depth1908
			return false
		},
		nil,
		/* 226 Action0 <- <{ p.s = p.label }> */
		nil,
		/* 227 Action1 <- <{ p.p = p.label }> */
		nil,
		/* 228 Action2 <- <{ p.o = p.label; p.addStatement(p.s, p.p, p.o) }> */
		nil,
		/* 229 Action3 <- <{ p.label = buffer[begin:end] }> */
		nil,
		/* 230 Action4 <- <{ p.label = buffer[begin:end] }> */
		nil,
		/* 231 Action5 <- <{ p.label = buffer[begin:end] }> */
		nil,
		/* 232 Action6 <- <{ p.label = "<http://www.w3.org/1999/02/22-rdf-syntax-ns#type>" }> */
		nil,
	}
	p.rules = rules
}
