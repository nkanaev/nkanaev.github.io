#!/usr/bin/env python3
import io
from collections import namedtuple

from parsimonious.grammar import Grammar
from parsimonious.nodes import NodeVisitor


WTFNode = namedtuple('WTFNode', ['name', 'attr', 'body'])

WTFGrammar = Grammar(r"""
root = (node / _)*
node = word _ "{" _ (attr / node / text / _)* _ "}"
attr = word _ ":" _ ~"[^;]*" _ ";"
text = ~r"\"[^\"]*\""
word = ~"[\w-]+"
_    = ~"\s*"
""")


class WTFVisitor(NodeVisitor):
    def visit_root(self, node, children):
        return [
            node
            for nodes in children
            for node in nodes
            if isinstance(node, WTFNode)]

    def visit_node(self, node, children):
        name = children[0].text
        attr = {}
        body = []
        for group in children[4]:
            for node in group:
                if isinstance(node, dict):
                    attr.update(node)
                elif isinstance(node, (str, WTFNode)):
                    body.append(node)
        return WTFNode(name, attr, body)

    def visit_attr(self, node, children):
        key, val = children[0], children[4]
        return {key.text: val.text}

    def visit_text(self, node, children):
        return node.text[1:-1]

    def generic_visit(self, node, children):
        return children or node


def _tohtml(node, out, indent=0, compact=False):
    ws = '' if compact else ' '
    nl = '' if compact else '\n'
    pp = lambda *args: print(*args, sep='', end='', file=out)

    # text
    if isinstance(node, str):
        pp(ws * indent, node, nl)
        return

    # opening tag
    pp(ws * indent, '<', node.name)
    for k, v in node.attr.items():
        pp(' ', '%s="%s"' % (k, v))
    pp('>', nl)

    # tag children
    for cn in node.body:
        _tohtml(cn, out, indent+4, compact)

    # closing tag
    pp(ws * indent, '</', node.name, '>')
    pp(nl)


def wtf2html(text, compact=False):
    tree = WTFGrammar.parse(text)
    nodes = WTFVisitor().visit(tree)

    out = io.StringIO()
    for node in nodes:
        _tohtml(node, out, compact=compact)

    return out.getvalue().rstrip()


def main():
    import sys
    if len(sys.argv) == 1:
        print('usage: %s file.wtfml' % sys.argv[0])
        return
    with open(sys.argv[1]) as f:
        print(wtf2html(f.read()))


if __name__ == '__main__':
    main()
