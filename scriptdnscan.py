import os
import sys

# Récupère le nom de domaine saisi par l'utilisateur en tant qu'argument du script
domain = sys.argv[1]
results_dir = sys.argv[2]

# executer la commande
file_name = os.path.join(results_dir, "dnscan.txt")
os.system(f'python.exe dnscan/dnscan.py -d {domain} -w dnscan/subdomains-100.txt -o {file_name}')
