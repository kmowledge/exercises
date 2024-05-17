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

# str_zwrotki = re.findall(r"{\'localTraceId'\:.+?\'query\':.+?\'variables\':.+?\'response\':.+?}},", str_data)
# print(str_zwrotki[0])


# matches items (params of request/response)
items = re.findall(r"('\w+'\:\s(?:({)|('))[\w\W]*?(?(2)}|(?(3)')),?\s*)", str_data)
[print(z) for z in items]
print('\n\n\n\n')
print(data) # <class 'dict'>
# [print(k, v) for k, v in data.items()]
# for k, v in data.items(): 
#     print(f"{k}: {v}") #this fails to print dict items in separate lines, dunno why  

#matches request-response pairs
# rr_pairs = re.findall(r"('\w+'\:\s(?:({)|('))[\w\W]*?(?(2)}|(?(3)')),?\s*){4}", str_data)
pattern = r"('\w+'\:\s(?:({)|('))[\w\W]*?(?(2)}|(?(3)')),?\s*)"
rr_pairs = re.findall(pattern + r".*?" + pattern + r".*?" + pattern + r".*?" + pattern, str_data)
[print(rr, '\n', type(rr), '\n') for rr in rr_pairs] #list of tuples
print('\n\n\n\n')
print(rr_pairs)

''' 
items - Problems:
1. nested dicts are cut. [SOLVED]
'variables': {'
'response': {'
2. one nested item from 'variables' is listed on query level.[SOLVED]
'type': 'blik'

rr_pairs - Problems:
1. only "response" items are included. [SOLVED]
2. pairs are cut in the wrong place. [unsolved]
'''

# how to set option: dot matches all? in case a need to apply pattern also to JSON (with \t \n).