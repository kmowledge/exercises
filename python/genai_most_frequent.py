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
    prompt = f"Analyze this chunk of numbers and identify the two most frequent numbers: {data_chunk}"
    
    response = client.chat.completions.create(
        model="gpt-3.5-turbo",
        messages=[
            {"role": "system", "content": "You are a data analyzer. For each chunk, identify the two most frequent numbers and format your response as: 'For this chunk, the most frequent numbers are X and Y'"},
            {"role": "user", "content": prompt}
        ],
        max_tokens=200,
        temperature=0.0
    )

    return response.choices[0].message.content.strip()
current_dir = os.path.dirname(os.path.abspath(__file__))
output_file = os.path.join(current_dir, '20mln_float_results.txt')

frequent_numbers = []
with open(output_file, 'w', encoding='utf-8') as f:  
    f.write("Analysis Results:\n\n")
    print(f"Created output file at: {output_file}")  

for chunk in split_data(data):
    result = get_most_frequent_numbers(chunk)
    print(result)  
    
    
    with open(output_file, 'a', encoding='utf-8') as f:
        f.write(f"{result}\n")
    
    
    numbers = [float(num) for num in result.split() if num.replace('.','').isdigit()]
    frequent_numbers.extend(numbers)
counter = Counter(frequent_numbers)
most_common = counter.most_common(2)
final_result = f"\nGlobal result: The two most frequent numbers across all chunks are {most_common[0][0]} and {most_common[1][0]}"
print(final_result)
with open(output_file, 'a', encoding='utf-8') as f:
    f.write(final_result)
    print(f"All results saved to: {output_file}")
