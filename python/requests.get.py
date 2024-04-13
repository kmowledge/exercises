import requests

#res = requests.get("https://www.univ-orleans.fr/iut-orleans/informatique/intra/tuto/django/Tuto-Django.pdf")
res = requests.get("https://wolnelektury.pl/media/book/txt/pan-tadeusz.txt")

try:
      res.raise_for_status()
except Exception as exc:
      print('Operation thrown the following error: %s' % (exc))

file = open('pan-tadeusz.txt', 'wb')

for chunk in res.iter_content(100000):
      file.write(chunk)

file.close()