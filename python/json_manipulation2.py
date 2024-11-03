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