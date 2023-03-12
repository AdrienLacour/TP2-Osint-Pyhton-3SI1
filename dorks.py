import requests
from bs4 import BeautifulSoup
import sys
import os
import urllib.parse

# Set up Google API key and search engine ID
domain = sys.argv[1]
results_dir = sys.argv[2]
GOOGLE_API_KEY = sys.argv[3]
SEARCH_ENGINE_ID = sys.argv[4]

# Define function to perform Google search
def google_search(query):
    # URL encode query string
    query = urllib.parse.quote_plus(query)
    # Construct Google API search URL
    url = f'https://www.googleapis.com/customsearch/v1?key={GOOGLE_API_KEY}&cx={SEARCH_ENGINE_ID}&q={query}'
    # Send GET request to Google API search URL
    response = requests.get(url)
    # Parse JSON response
    results = response.json()
    # Return list of search result URLs
    return [item['link'] for item in results.get('items', [])]

# Define main function to perform Google Dorks
def main():
    # Define Google Dork queries
    queries = [
        f'site:{domain} filetype:pdf',
        f'intitle:"confidential" site:{domain}',
        f'inurl:admin site:{domain}',
       # f'inurl:"phpmyadmin" site:{domain}',
       # f'intext:"utilisateurs enregistrés" site:{domain}',
       # f'intext:"mot de passe" site:{domain}',
       # f'intext:"nom d\'utilisateur" site:{domain}',
       # f'intext:"identifiant" site:{domain}'
    ]
    # Perform Google searches for each query
    for query in queries:
        print(f'Searching for: {query}')
        results = google_search(query)
        # Print search results
        filename = os.path.join(results_dir, "dorks.txt")
        with open(filename, "w", encoding='utf-8') as f:
            for url in results:
                f.write(f"Google dork: {url}\n")
       
if __name__ == '__main__':
    main()
