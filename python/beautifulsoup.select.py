# This might be universal script for scraping series of articles, like tutorials, but I would need
# to find more universal html,css identifiers for text. Probably even search for completely another method.

import bs4, requests
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

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


# Read html into soup object.
# file = open('Manual,test-driven-go.html', 'rb') We have to open for reading, not writing as before.
with open('Manual,test-driven-go.html') as file:
      soup = bs4.BeautifulSoup(file, features="html.parser")

# (manually) Inspect the site for text elements. What's the mutual value?
# -> class value of p element: <p class="max-w-3xl w-full mx-auto decoration-primary/6 page-api-block:ml-0">
# -> h1,h2,h3 etc. + p elements

# Select html elements with the deserved text content. Write text of these elements into destination file.

elems = soup.select('p')
destinFile = open('tutorial.txt', 'wt')
for elem in elems: # type(elem) <class 'bs4.element.Tag'>
      selem = str(elem.string) # type(elem.string) <class 'bs4.element.NavigableString'>
      destinFile.write(selem)

# print(destinFile)
destinFile.close()

# (manually) Check the result. Open the file with scraped text of 1st chapter.
#-> Paragraphs are not maintained, one long string, and cull. Tried to prettify() the html before str().

# (manually) Inspect the 1. and 2. chapter's site for the button to move to the next one.
#-> last link on the page containing path: learn-go-with-tests

# Select html element that indicates it's the next chapter button. Click on it. Handle exception.
# OR derive the link url from button element and use it (change site without clicking the button).
"""
driver = webdriver.Chrome()
driver.get('https://quii.gitbook.io/learn-go-with-tests#read-the-book')

# <div class="flex flex-col md:flex-row mt-6 gap-2 max-w-3xl mx-auto page-api-block:ml-0"><a class="group text-sm p-2.5 flex gap-4 flex-1 flex-row items-center pr-4 border border-dark/3 rounded straight-corners:rounded-none hover:border-primary/6 dark:border-light/2 text-pretty dark:hover:border-primary-300/4 md:p-4 md:text-base" href="/learn-go-with-tests/go-fundamentals/install-go"><span class="flex flex-col flex-1"><span class="text-xs">Next</span><span class="text-dark dark:text-light/6 group-hover:text-primary line-clamp-2">Install Go</span></span><svg fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" shape-rendering="geometricPrecision" viewBox="0 0 24 24" class="hidden w-5 h-5 stroke-dark/5 group-hover:stroke-primary dark:stroke-light/4 md:block" height="24" width="24" style="color:currentColor"><path d="M9 18l6-6-6-6"></path></svg></a></div>

# elems2 = ''
try:
      elem2 = driver.find_element(By.CSS_SELECTOR,'body > div > div > div > div > main > div.flex.flex-col.md\:flex-row.mt-6.gap-2.max-w-3xl.mx-auto.page-api-block\:ml-0 > a')
      elem2.click()
      # elems2 = WebDriverWait(driver, 5).until(EC.presence_of_all_elements_located((By.PARTIAL_LINK_TEXT,"learn-go-with-tests")))
      # elems2[-1].click()
except Exception as e:
      print("Couldn't find element:", e)
finally:
      print(driver.current_url)
      driver.quit()
# Contain the actions (scraping, clicking button) in the loop, until there's no button identifier found.
# (manually) Check the result.
# Polish the code, maybe it needs one blank line between chapters for readibilty.
