import traceback
import json
import sys
import requests
import urlscan
import os

# Get command-line arguments
domain = sys.argv[1]
results_dir = sys.argv[2]
URLSCAN_API_KEY = sys.argv[3]

try:
    headers = {'Content-Type': 'application/json', 'API-Key': URLSCAN_API_KEY}
    data = {"url": f"https://{domain}/", "visibility": "public"}
    response = requests.post('https://urlscan.io/api/v1/scan/',headers=headers, data=json.dumps(data))
    response_json = response.json()
    if response.status_code == 200:
        print("--------------------URLSCAN EN PROGRES--------------------")
        task_url = response_json["result"]
        #GOTTA FIND A WAY TO PARSE RESULT URL IN THE TERMINAL
        filename = os.path.join(results_dir, "urlscan.txt")
        with open(filename, "w", encoding='utf-8') as f:
            f.write(f"Les infos de l'urlscan ici: {task_url}")
    else:
	    print(f"L'URLscan a raté. Code de status: {response.status_code}")
except Exception as e:
		print("Impossible de faire la recherche.")
		traceback.print_exc()