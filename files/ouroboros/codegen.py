quine = r'''#include <stdio.h>
int main(){
char*c="?",o[1000],g[1000];
sprintf(g,"package main;import%cfmt%c;func main(){fmt.Print(%c%cs%c)}",34,34,96,37,96);
sprintf(o,c,!);
fprintf(stdout,g,o);
return 0;
}
'''

quinestr = ''
args = []
for i, ch in enumerate(quine):
    if ch in '\n"%':
        quinestr += '%c'
        args.append(ord(ch))
    elif ch == '?':
        quinestr += '%s'
        args.append('c')
    else:
        quinestr += ch

args = ','.join(map(str, args))

print(
    quine\
        .replace('?',
            quinestr\
                .replace('?', '%s')\
                .replace('!', args))\
        .replace('!', args),
    end='')
