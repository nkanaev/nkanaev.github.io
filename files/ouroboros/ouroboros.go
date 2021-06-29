package main;import"fmt";func main(){fmt.Print(`#include <stdio.h>
int main(){
char*c="#include <stdio.h>%cint main(){%cchar*c=%c%s%c,o[1000],g[1000];%csprintf(g,%cpackage main;import%ccfmt%cc;func main(){fmt.Print(%cc%ccs%cc)}%c,34,34,96,37,96);%csprintf(o,c,10,10,34,c,34,10,34,37,37,37,37,37,34,10,10,10,10,10);%cfprintf(stdout,g,o);%creturn 0;%c}%c",o[1000],g[1000];
sprintf(g,"package main;import%cfmt%c;func main(){fmt.Print(%c%cs%c)}",34,34,96,37,96);
sprintf(o,c,10,10,34,c,34,10,34,37,37,37,37,37,34,10,10,10,10,10);
fprintf(stdout,g,o);
return 0;
}
`)}