"""
Napisz program, który wczyta plik json ze zwrotkami z backoffice,
podzieli na pojedyncze zwrotki, sprawdzi, czy nie występują zwrotki ze zduplikowanym zapytaniem,
jeśli takie znajdzie, powinien je wydrukować w całości, jeśli nie, dać odpowiedni komunikat.
"""

import json
import re

with open("./.ins/blik.json", "r") as read_file:
      data = json.load(read_file)

str_data = str(data)

# matches items (params of request/response)
items = re.findall(r"('\w+'\:\s?(?:({)|('))[\w\W]*?(?(2)}|(?(3)')),?\s?)", str_data)
[print(z) for z in items] # items is a list of tuples
print('\n\n\n\n')

pattern = r"('\w+'\:\s(?:({)|('))[\w\W]*?(?(2)}|(?(3)')),?\s*)"
rr_pairs = re.findall(pattern + r".*?" + pattern + r".*?" + pattern + r".*?" + pattern, str_data)
[print(rr, '\n') for rr in rr_pairs]

''' 
items - Problems:
1. nested dicts are cut. [SOLVED]
'variables': {'
'response': {'
2. one nested item from 'variables' is listed on query level.[SOLVED]
'type': 'blik'
3. there are minor false-positives, probably easy to eradicate [unsolved]
("'localTraceId': '', ", '', "'")
, ", '{', '')


rr_pairs - Problems:
1. only "response" items are included. [SOLVED]
2. pairs are cut in the wrong place. [unsolved]
'''