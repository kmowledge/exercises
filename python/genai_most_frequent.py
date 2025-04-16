import zipfile
import os
import openai
import numpy as np
import time
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
    prompt = f"Find exactly two most frequent numbers in this data: {data_chunk}. Return only these two numbers separated by comma, nothing else."
    
    max_retries = 3
    for attempt in range(max_retries):
        try:
            response = client.chat.completions.create(
                model="gpt-3.5-turbo",
                messages=[
                    {"role": "system", "content": "You are a data analyzer. Return ONLY two numbers separated by comma, no other text."},
                    {"role": "user", "content": prompt}
                ],
                max_tokens=50,
                temperature=0.0
            )
            result = response.choices[0].message.content.strip()
            return result
        except Exception as e:
            if attempt == max_retries - 1:
                raise
            print(f"Attempt {attempt + 1} failed, retrying... Error: {str(e)}")
            time.sleep(2)
current_dir = os.getcwd()
output_file = os.path.join(current_dir, '20mln_float_results.txt')
chunk_counter = 0

with open(output_file, 'w', encoding='utf-8') as f:
    f.write("chunk,number1,number2\n")  

frequent_numbers = []

for chunk in split_data(data):
    chunk_counter += 1
    try:
        result = get_most_frequent_numbers(chunk)
        print(f"Processing chunk {chunk_counter}")
        
        
        numbers = [num.strip() for num in result.split(',')]
        if len(numbers) == 2:
            with open(output_file, 'a', encoding='utf-8') as f:
                f.write(f"{chunk_counter},{numbers[0]},{numbers[1]}\n")
            frequent_numbers.extend(float(num) for num in numbers)
        else:
            print(f"Warning: Unexpected format in chunk {chunk_counter}: {result}")
    except Exception as e:
        print(f"Error processing chunk {chunk_counter}: {str(e)}")
        continue
counter = Counter(frequent_numbers)
most_common = counter.most_common(2)

with open(output_file, 'a', encoding='utf-8') as f:
    f.write(f"\nFINAL,{most_common[0][0]},{most_common[1][0]}\n")

print(f"Analysis complete. Results saved to {output_file}")
