#! /usr/bin/python3

import webbrowser

top_sites = ['imgur.com', 'msn.com', 'wikipedia.org']
for site in top_sites:
      url_pattern = ['https://', site, '/']
      webbrowser.open(''.join(url_pattern))