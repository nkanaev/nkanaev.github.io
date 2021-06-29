lines = open('shq.go').readlines()
start = next(i for i, line in enumerate(lines) if line.startswith('/* --'))
content = ''.join(lines[start:])
content = content.replace('\n', 'N')
content = content.replace('\t', 'T')
content = content.replace('`', 'Q')
width = 60
chunks = [content[i:i+width] for i in range(0, len(content), width)]
for i, chunk in enumerate(chunks):
    print("`", chunk, "`+" if i != len(chunks)-1 else "`)", sep='')
