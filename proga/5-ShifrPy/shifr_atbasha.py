with open('input.txt', 'r', encoding='utf-8') as file:
    text = file.read()

alp = 'abcdefghijklmnopqrstuvwxyz'
r_alp = alp[::-1]

text = text.lower()
r_text = ''

for char in text:
    if char in alp:
        ind = alp.index(char)
        r_text += r_alp[ind]
    else:
        r_text += char

with open('output.txt', 'w', encoding='utf-8') as file:
    file.write(r_text)