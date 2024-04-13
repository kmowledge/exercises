# This might be universal script for scraping series of articles, like tutorials, but I would need
# to find more universal identifiers for text. Probably even change method for more appropriate.

import bs4, requests

res = requests.get('https://quii.gitbook.io/learn-go-with-tests#read-the-book')

try:
      res.raise_for_status()
except Exception as exc:
      print("Request brough an exception: %s" %exc)

print(res.text[:250])
file = open('Manual,test-driven-go.html', 'wb')

for chunk in res.iter_content(100):
      print(chunk)

# <note for myself> Argument of iter_content is size of a chunk, in bytes.
# It prevents loading all at once, allows for concurrent task. It's simply bufferring. chunk=buffer
for chunk in res.iter_content(100000):
      file.write(chunk)

file.close()

# To-do next steps:
# Read html into soup object.
# (manually) Inspect the site for the text elements. What's the mutual value?
# Select html elements with the deserved text content. Write text of these elements into destination file.
# (manually) Check the result. Open the file with scraped text of 1st chapter.
# (manually) Inspect the 1. and 2. chapter's site for the button to move to the next one.
# Select html element that indicates it's the next chapter button. Click on it. Handle exception.
# Contain the actions (scraping, clicking button) in the loop, until there's no button identifier found.
# (manually) Check the result.
# Polish the code, maybe it needs one blank line between chapters for readibilty.