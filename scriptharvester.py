import os
import sys

# Récupère le nom de domaine saisi par l'utilisateur en tant qu'argument du script
domain = sys.argv[1]
results_dir = sys.argv[2]

harvester_dir = os.path.abspath("theHarvester")

# Change directory to theHarvester
os.chdir(harvester_dir)

# Bing search command
file_name = os.path.join(results_dir, "harvester-bing.txt")
os.system(f"python.exe theHarvester.py -d {domain} -l 500 -b bing -f {file_name}")

# Yahoo search command
file_name = os.path.join(results_dir, "harvester-yahoo.txt")
os.system(f"python.exe theHarvester.py -d {domain} -l 500 -b yahoo -f {file_name}")

# DnsDumpster search command
file_name = os.path.join(results_dir, "harvester-dnsdumpster.txt")
os.system(f"python.exe theHarvester.py -d {domain} -l 500 -b dnsdumpster -f {file_name}")

# Shodan search command / Invalid source
#file_name = os.path.join(results_dir, "shodan-{}.txt".format(domain))
#os.system(f"python.exe theHarvester.py -d {domain} -l 500 -b shodan -f {file_name}")
