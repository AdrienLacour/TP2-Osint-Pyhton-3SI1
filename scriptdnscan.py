import os
import sys

# Récupère le nom de domaine saisi par l'utilisateur en tant qu'argument du script
domain = sys.argv[1]
results_dir = sys.argv[2]

# exécuter la commande en utilisant un résolveur DNS personnalisé
file_name = os.path.join(results_dir, "dnscan.txt")
os.system(f'python.exe dnscan/dnscan.py -d {domain} -w dnscan/subdomains-100.txt -o {file_name} -R 8.8.8.8')


