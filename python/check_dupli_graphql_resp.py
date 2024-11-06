"""
Napisz program, który wczyta plik json ze zwrotkami z backoffice,
podzieli na pojedyncze zwrotki, sprawdzi, czy nie występują zwrotki ze zduplikowanym zapytaniem,
jeśli takie znajdzie, powinien je wydrukować w całości, jeśli nie, dać odpowiedni komunikat.
"""
import json

def find_duplicate_exchanges(exchanges):
      seen_queries = set()
      seen_variables = set()
      duplicates = []

      for exchange in exchanges:
            query = exchange.get('query')
            variables = tuple(sorted(exchange.get('variables', {}).items()))

            if query in seen_queries and variables in seen_variables:
                  duplicates.append(exchange)
            else:
                  seen_queries.add(query)
                  seen_variables.add(variables)
      
      if len(duplicates) == 0:
            return 'No duplicate requests found in this file.'
      else:
            return duplicates

with open('./.ins/blik.json') as file:
      data = json.load(file)

exchanges = data['requestResponses'] 
duplicates = find_duplicate_exchanges(exchanges)
if type(duplicates) != str:
      for d in duplicates:
            print(d, '\n')
else:
      print(duplicates)

# ↓↓↓↓↓↓↓↓↓↓↓↓ How I tried to do this before knowing the simple way above ↓↓↓↓↓↓↓↓↓↓↓↓↓

str_data = str(data)

# matches items (params of request/response)
items = re.findall(r"('\w+'\:\s?(?:({)|('))[\w\W]*?(?(2)}|(?(3)')),?\s?)", str_data)
[print(z) for z in items] # items is a list of tuples
print('\n\n\n\n')

pattern = r"('\w+'\:\s(?:({)|('))[\w\W]*?(?(2)}|(?(3)')),?\s*)"
rr_pairs = re.findall(pattern + r".*?" + pattern + r".*?" + pattern + r".*?" + pattern, str_data)
[print(rr, '\n') for rr in rr_pairs]