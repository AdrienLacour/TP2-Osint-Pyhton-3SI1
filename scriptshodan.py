import sys
import shodan
import sys
import os

# Récupère le nom de domaine saisi par l'utilisateur en tant qu'argument du script
domain = sys.argv[1]
results_dir = sys.argv[2]
SHODAN_API_KEY = sys.argv[3]

try:
    # Creation d'une instance
    api = shodan.Shodan(SHODAN_API_KEY)

    # Demande de la recherche a l'userà
    query = domain

    # Lance la recherche
    result = api.search(query)
    print("Results number : ", len(result["matches"]))

    # Ouverture d'un fichier pour enregistrer les resultats
    filename = os.path.join(results_dir, "shodan.txt")
    with open(filename, "w", encoding='utf-8') as f:
     for match in result["matches"]:
       # Extraction des informations pertinentes
       f.write("IP: {}".format(match['ip_str']) + "\n")

       if 'org' in match:
            f.write("Organisation: {}".format(match.get("org")) + "\n")

       f.write("Emplacement: {}, {}".format(match.get("location", {}).get("city", "") or "", match.get("location", {}).get("country_name", "") or "") + "\n")

       if 'domains' in match:
            f.write("Domains: {}".format(",".join(match.get("domains"))) + "\n")

       if 'server' in match:
            f.write("Server: {}".format(match.get("server")) + "\n")

       f.write("\n\n")

     print("Les resultats de la recherche ont ete enregistres dans "  + filename)

except shodan.APIError as e:
    print("Erreur de l'API: {}".format(e))
