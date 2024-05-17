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

str_zwrotki = re.findall(r"{\'localTraceId'\:.+?\'query\':.+?\'variables\':.+?\'response\':.+?}},", str_data)
# print(str_zwrotki[0])


# matches items (params of request/response)
zwrotki = re.findall(r"('\w+'\:\s(?:({)|('))[\w\W]*?(?(2)}|(?(3)')),?\s*)", str_data)
[print(z) for z in zwrotki]
print('\n\n\n\n')
print(zwrotki)
print('\n\n\n\n')
print(data) # <class 'dict'>
# [print(k, v) for k, v in data.items()]
# for k, v in data.items(): 
#     print(f"{k}: {v}") #this fails to print dict items in separate lines, dunno why  

''' Problems:
1. nested dicts are cut.
'variables': {'
'response': {'

2. one nested item from 'variables' is listed on query level.
'type': 'blik'
'''

# how to set option: dot matches all? in case a need to apply pattern also to JSON (with \t \n).