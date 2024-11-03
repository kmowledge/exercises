import requests, json, os, time

isbn_to_find = '8871351460'

url = "https://data.bn.org.pl/api/institutions/bibs.json"

response = requests.get(url)
response.encoding = 'utf-8'

if response.status_code != 200:
    print('status_code: ', response.status_code)
    os.exit()

data = response.json()

with open('./.outs/book-nonspecific-query.json', 'wt', encoding='utf-8') as f:
    json.dump(data, f, indent=4, ensure_ascii=False)



def name_next_result(dir: str, name: str) -> tuple[str, int]:
    i = 1
    for filename in os.listdir(dir):
        if not os.listdir(dir):
            break
        i = int(''.join([char for char in filename if char.isdigit()])) + 1
    result_name = name
    return result_name, i

record_found = None

def find_record(field: str, value: object):
    global record_found #deklarujemy, że będziemy używać zmiennej spoza funkcji
    counter = 0
    result_name, i = '', 0 #inicjalizuję, żeby były dostępne na poziomie całej funkcji, a nie tylko w obrębie pętli
    for bib in data.get('bibs', []):
        if bib.get(field) == value:
            record_found = bib
            counter += 1
            result_name, i = name_next_result('./book-search-results', 'result')
    with open(f'./.outs/book-search-results/{result_name}{i}.json', 'wt', encoding='utf-8') as search_result:
        json.dump(bib, search_result, indent=4, ensure_ascii=False)
    
    print(record_found)
    print(f'\nLiczba pozycji w zbiorze z polem "{field}" o wartości "{value}" wynosi {counter}.')
    # return record_found - bez, bo odwołuję się do globalnej, find_record jest tu procedurą.

def find_nested(): #zamiast tej funkcji należałoby rozwinąć mechanizm find_record, ale już pal licho.
    marc["fields"][0]["001"]

# find_record('language', 'polski') działa
# find_record('marc["fields"]["001"]', 'b000000323260') źle, trzeba odwołać się do elementu listy.
# find_record('marc["fields"][0]["001"]', "b000000323260") tego Python nie zinterpretuje.
find_record(find_nested(), "b000000323260")
