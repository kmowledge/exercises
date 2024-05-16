"""
Napisz program, który wczyta plik json ze zwrotkami z backoffice,
podzieli na pojedyncze zwrotki, sprawdzi, czy nie występują zwrotki ze zduplikowanym zapytaniem,
jeśli takie znajdzie, powinien je wydrukować w całości, jeśli nie, dać odpowiedni komunikat.
"""

import json
import re

with open("./ins/blik.json", "r") as read_file:
      data = json.load(read_file)

re.findall("\{"localTraceId" }")
{\s[ \t]+\"localTraceId..+,
"{\'localTraceId'\:.+?\'query\':.+?\'variables\':.+?\'response\':.+?}},"g
# for chunk in "},".split()