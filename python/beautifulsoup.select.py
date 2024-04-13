# This might be universal script for scraping series of articles, like tutorials, but I would need
# to find more universal identifiers for text. Probably even change method for more appropriate.

import bs4, requests
"""
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
"""

# Read html into soup object.
# file = open('Manual,test-driven-go.html', 'rb') We have to open for reading, not writing as before.
with open('Manual,test-driven-go.html') as file:
      soup = bs4.BeautifulSoup(file, features="html.parser")

# (manually) Inspect the site for text elements. What's the mutual value?
# <p class="max-w-3xl w-full mx-auto decoration-primary/6 page-api-block:ml-0">
# h1,h2,h3 etc. and p elements

# Select html elements with the deserved text content. Write text of these elements into destination file.

elems = soup.select('p')
destinFile = open('tutorial.txt', 'wt')
for elem in elems: # type(elem) <class 'bs4.element.Tag'>
      selem = str(elem.string) # type(elem.string) <class 'bs4.element.NavigableString'>
      destinFile.write(selem)

# print(destinFile)
destinFile.close()

# (manually) Check the result. Open the file with scraped text of 1st chapter.
# Paragraphs are not maintained, one long string, and missing. Tried to prettify() the html before str().


# (manually) Inspect the 1. and 2. chapter's site for the button to move to the next one.
# Select html element that indicates it's the next chapter button. Click on it. Handle exception.
# OR derive the link url from button element and use it (change site without clicking the button).
# Contain the actions (scraping, clicking button) in the loop, until there's no button identifier found.
# (manually) Check the result.
# Polish the code, maybe it needs one blank line between chapters for readibilty.
print(type(elem.string))