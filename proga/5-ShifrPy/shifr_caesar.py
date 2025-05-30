with open('input.txt', 'r', encoding='utf-8') as file:
    text = file.read()

shift = 3
alp = 'abcdefghijklmnopqrstuvwxyz'
shifted_alp = alp[shift:] + alp[:shift]

text = text.lower()
result = ''

for char in text:
    if char in alp:
        ind = alp.index(char)
        result += shifted_alp[ind]
    else:
        result += char

with open('output.txt', 'w', encoding='utf-8') as file:
    file.write(result)