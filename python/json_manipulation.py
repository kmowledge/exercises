"""
Napisz program, który wczyta plik json ze zwrotkami z backoffice,
podzieli na pojedyncze zwrotki, sprawdzi, czy nie występują zwrotki ze zduplikowanym zapytaniem,
jeśli takie znajdzie, powinien je wydrukować w całości, jeśli nie, dać odpowiedni komunikat.
"""

import json
import re

with open("./ins/blik.json", "r") as read_file:
      data = json.load(read_file)

print(type(data))
str_data = str(data)
print(type(str_data))

zwrotki = re.findall(r"{\'localTraceId'\:.+?\'query\':.+?\'variables\':.+?\'response\':.+?}},", str_data)
print(zwrotki[0])
dict_data = dict(enumerate(str_data))
print(dict_data) # out >> {0: '{', 1: "'", 2: 'r', 3: 'e', 4: 'q', 5: 'u', 6: 'e', 7: 's', 8: 't',
# for chunk in "},".split()