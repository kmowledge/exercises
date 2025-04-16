import zipfile
import os
import openai
import numpy as np
from dotenv import load_dotenv
from collections import Counter
load_dotenv()
openai.api_key = os.getenv("OPENAI_API_KEY")
with zipfile.ZipFile('data.zip', 'r') as zip_ref:
    zip_ref.extractall('ins')
zip_contents = os.listdir('ins')

with open(f'ins/{zip_contents[0]}', 'r') as f:
    data = f.readlines()
data = np.array(data, dtype=float)
def split_data(data, chunk_size=10000):
    for i in range(0, len(data), chunk_size):
        yield data[i:i + chunk_size]

client = openai.OpenAI() 

def get_most_frequent_numbers(data_chunk):
    
    prompt = f'''
    Identify two most frequent numbers in this dataset: {data_chunk}. Print them.
    '''
    
    
    
    response = openai.Completion.create(
        engine='gpt-3.5-turbo',  
        messages=[
            {'role': 'system', 'content': 'You are a data analyzer. respond only with two numbers separated by comma'}
        ],
        max_tokens=200,
        temperature=0.0
    )

    
    return response.choices[0].text.strip()
frequent_numbers = []

for chunk in split_data(data):
    result = get_most_frequent_numbers(chunk)
    frequent_numbers.append(result)
all_results = []
for result in frequent_numbers:
    
    nums = result.split(',')
    all_results.extend([float(num.strip()) for num in nums])
counter = Counter(all_results)
most_common = counter.most_common(2)
print("\nFrequent numbers from each chunk:")
for result in frequent_numbers:
    print(result)
print(f"Global result: Two most frequent numbers are: {most_common[0][0]} and {most_common[1][0]}")
