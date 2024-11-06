# ZADANIE 1. GENERATOR KODÓW POCZTOWYCH
# przyjmuje 2 stringi: "79-900" i "80-155" i zwraca listę kodów pomiędzy nimi

from random import randint
import logging

logging.basicConfig(filename='log.txt', level=logging.INFO, format='%(asctime)s - %(message)s')

def is_postcode_format(ispc: str) -> bool:
    try:
        assert type(ispc) == str
    except AssertionError:
        logging.info("Błędny typ inputu.") #Prompt z inputu zawsze będzie stringiem, więc to sprawdzenie jest naddatkowe dla prostego programiku, ale zostawiam na perspektywę dalszego developmentu.
        print(f'Błędny typ inputu.')
        return False
    try:
        pc_init, pc_ending = ispc.split('-', 1)
    except:
        logging.info("Błędny format kodu pocztowego. Nie udało sie podzielić liczb względem '-'.")
        print(f'Nie udało sie podzielić liczb względem "-".')
        return False
    try:
        assert len(pc_init) == 2 and len(pc_ending) == 3
    except AssertionError:
        logging.info("Błędny format kodu pocztowego. Nieprawidłowa długość cyfr.")
        print(f'Nieprawidłowa długość cyfr.\n')
        return False
    finally:
        return True

def enforce_format(startc: str, endc: str) -> None:
    postcodes = list()
    postcodes.extend([startc, endc])
    for i, c in enumerate(postcodes):
        while not is_postcode_format(c):
            c = input(f"Wartość '{c}' jest niepoprawna. Podaj ją ponownie.")
            if i == 0: startc = c
            elif i == 1: endc = c
      
def generate_postcodes_between(startc: str, endc: str) -> list:
    enforce_format(startc, endc)

    startc = [int(part) for part in startc.split('-', 1)]
    endc = [int(part) for part in endc.split('-', 1)]
    pc_list = list()
    for pc_init in range(startc[0], endc[0] + 1):

        a, b = 0, 1000 #default indexes
        if pc_init == startc[0]: a = startc[1] # custom starting index
        if pc_init == endc[0]: b = endc[1] + 1 # custom ending index
        for pc_ending in range(a, b):
            postcode = f'{pc_init:02}' + '-' + f'{pc_ending:03}'
            pc_list.append(postcode)

    return print(pc_list)



def launch_script():
    # startc = input(f"Podaj kod pocztowy, od którego mam generować kolejne.\n")
    # endc = input(f"Podaj kod pocztowy, do którego mam skończyć generowanie.\n")
    # testing mocks:
    print('generate_postcodes_between("79-900", "79-950")')
    pc_list = generate_postcodes_between("79-900", "79-950")
    print('/n/n/n')
    print('generate_postcodes_between("79-900", "81-155")')
    pc_list = generate_postcodes_between("79-900", "81-155")
    # Opakuj w main i popraw funkcję sprawdzającą format. Napraw logowanie (nie zapisuje logów). Dodaj log z KeyboardInterrupt i każdego odpalenia programu.

launch_script()
