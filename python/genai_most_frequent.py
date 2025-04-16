# Poniższy plik, o nazwie “data.zip” (około 100MB) to archiwum w formacie ZIP, które zawiera jeden plik tekstowy.
# W tym pliku w każdej linii znajduje się jedna liczba rzeczywista w zakresie od -180 do 180. Aby poznać rozwiązanie utwórz histogram ilości wystąpień każdej z liczb.
# Rozwiązaniem jest ciąg znaków zawierający dwie najczęściej występujące liczby w danych oddzielone przecinkiem.
import zipfile
import os
import matplotlib.pyplot as plt
import pandas as pd
import numpy as np

# with zipfile.ZipFile('data.zip', 'r') as zip_ref:
#     zip_ref.extractall('ins')
zip_contents = os.listdir('ins')

with open(f'ins/{zip_contents[0]}', 'r') as data:
    data = data.readlines()

data = np.array(data)
# plt.hist(data, bins=np.unique(data).size); plt.show()
fig, ax = plt.subplots()
ax.hist(data, bins=np.unique(data).size)
plt.show()