package main;import("fmt";"strings";"regexp");func main(){ x:=(
`/* ---------------------------------------------------------`+
`-*/No:=""; P:=func(x interface{}) {o += fmt.Sprint(x)};P(Qpa`+
`ckageQ+N  Q main;import("fmt";"strings";"regexp");func main(`+
`){ x:=(Q)N  for i,ch:=range x {; if i%60==0 { P("\nQ") }; P(`+
`string(ch))N if i==len(x)-1 {P("Q)")}else if i%60==59 {P("Q+`+
`")} }; P("\n")Na:=strings.Replace(x,string(78),string(10),-1`+
`) /* newline \n */Nb:=strings.Replace(a,string(84),string(9)`+
`,-1)  /* tab     \t */Nc:=strings.Replace(b,string(81),strin`+
`g(96),-1) /* quote   \Q */NP(c); r:=regexp.MustCompile(Q(\x6`+
`0.*?\x60|\x22[^\n"]*\x22|/\*Q+NQ.*?\*/|\bpackage|import|func`+
`|interface|if|else|for|range|swiQ+NQtch|case|return|default\`+
`b)Q); Replace :=r.ReplaceAllStringFuncNfmt.Print (Replace (o`+
`, func (x string) string {  switch x[0]  {Ncase 0x22: return`+
` "\x1b[36m"+x+"\x1b[0m"; /* "double-quoted" */Ncase 0x60: re`+
`turn "\x1b[36m"+x+"\x1b[0m"; /*  Qback-quotedQ  */Ncase 0x2f`+
`: return "\x1b[90m"+x+"\x1b[0m"; /* general comment */Ndefau`+
`lt: return "\x1b[0;1m"+x+"\x1b[0m"} })); /* keyword    */}N`)
/* ----------------------------------------------------------*/
o:=""; P:=func(x interface{}) {o += fmt.Sprint(x)};P(`package`+
  ` main;import("fmt";"strings";"regexp");func main(){ x:=(`)
  for i,ch:=range x {; if i%60==0 { P("\n`") }; P(string(ch))
 if i==len(x)-1 {P("`)")}else if i%60==59 {P("`+")} }; P("\n")
a:=strings.Replace(x,string(78),string(10),-1) /* newline \n */
b:=strings.Replace(a,string(84),string(9),-1)  /* tab     \t */
c:=strings.Replace(b,string(81),string(96),-1) /* quote   \` */
P(c); r:=regexp.MustCompile(`(\x60.*?\x60|\x22[^\n"]*\x22|/\*`+
`.*?\*/|\bpackage|import|func|interface|if|else|for|range|swi`+
`tch|case|return|default\b)`); Replace :=r.ReplaceAllStringFunc
fmt.Print (Replace (o, func (x string) string {  switch x[0]  {
case 0x22: return "\x1b[36m"+x+"\x1b[0m"; /* "double-quoted" */
case 0x60: return "\x1b[36m"+x+"\x1b[0m"; /*  `back-quoted`  */
case 0x2f: return "\x1b[90m"+x+"\x1b[0m"; /* general comment */
default: return "\x1b[0;1m"+x+"\x1b[0m"} })); /* keyword    */}
