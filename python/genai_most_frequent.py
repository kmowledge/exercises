# Poniższy plik, o nazwie “data.zip” (około 100MB) to archiwum w formacie ZIP, które zawiera jeden plik tekstowy.
# W tym pliku w każdej linii znajduje się jedna liczba rzeczywista w zakresie od -180 do 180. Aby poznać rozwiązanie utwórz histogram ilości wystąpień każdej z liczb.
# Rozwiązaniem jest ciąg znaków zawierający dwie najczęściej występujące liczby w danych oddzielone przecinkiem.
import zipfile
import os
import matplotlib.pyplot as plt
import pandas as pd
import numpy as np
# from openai import OpenAI
import openai
# from dotenv import load_dotenv

# with zipfile.ZipFile('data.zip', 'r') as zip_ref:
#     zip_ref.extractall('ins')


zip_contents = os.listdir('ins')

with open(f'ins/{zip_contents[0]}', 'r') as data:
    data = data.readlines()
data = np.array(data)
# data.shape (20000000,)

# fig, ax = plt.subplots()
# ax.hist({data}, bins=np.unique({data}).size)
# plt.show()

# load_dotenv()
# client = OpenAI(api_key = os.getenv('OPENAI_API_KEY'))
openai.api_key = os.getenv("OPENAI_API_KEY")

prompt = f'''
Identify two most frequent numbers in this dataset: {data}. Print them.
'''

# response = client.chat.completions.create(
response = openai.Completion.create(
    engine='gpt-4',
    prompt=prompt,
    max_tokens=200,
    temperature=0.0
    # messages=[
    #     {'role': 'user', 'content': prompt},
    #     {'role': 'user', 'content': ': '}
    # ]
)

# if I don't split data into 10k chunks the calculation of 2mln elements array will yield bill of $6000 compared to $12 if I split them.