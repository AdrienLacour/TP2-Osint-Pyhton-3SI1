import subprocess
import os
import tkinter as tk

# Définition des variables globales pour domain et results_dir
domain = ""
results_dir = ""

def center_window(root, w, h):
    # Récupération des dimensions de l'écran
    screen_width = mainapp.winfo_screenwidth()
    screen_height = mainapp.winfo_screenheight()

    # Calcul des coordonnées pour centrer la fenêtre
    x = (screen_width/2) - (w/2)
    y = (screen_height/2) - (h/2)

    # Configuration de la fenêtre avec ces coordonnées
    mainapp.geometry("%dx%d+%d+%d" % (w, h, x, y))
    
def exit_app():
    mainapp.quit()
    mainapp.destroy()

def open_apikey_file():
    os.system("notepad.exe apikey.txt")

# Lecture des clés API depuis le fichier apikey.txt
shodanapikey = ''
urlscanapikey = ''

with open('apikey.txt') as f:
    for line in f:
        key, value = line.strip().split(':')
        if key == 'shodanapikey':
            shodanapikey = value
        elif key == 'urlscanapikey':
            urlscanapikey = value
            
def get_domain():
    global domain, results_dir
    domain = domain_entry.get()
    if not domain:
        return
    domain_label.config(text="Vous avez choisi comme nom de domaine : " + domain)

    # Créer le chemin du dossier
    results_dir = os.path.join(os.getcwd(), domain)

    # Vérifier si le dossier existe déjà
    if not os.path.exists(results_dir):
        os.mkdir(results_dir)
        print("Le dossier " + domain + " a été créé avec succès.".format(results_dir))
        result_label.config(text="Le dossier " + domain + " a été créé avec succès.")
    else:
        print("Le dossier " + domain + " existe déjà.".format(results_dir))
        result_label.config(text="Le dossier " + domain + " existe déjà.")       

# Affichage des options
def show_genpwd_button():
    launch_button.pack_forget()
    launch_button2.pack_forget()
    domain_label3.pack()
    domain_entry.pack(pady=20)
    domain_button.pack(pady=20)
    domain_label2.pack()
    script1_check.pack(pady=5)
    script2_check.pack(pady=5)
    script3_check.pack(pady=5)
    script4_check.pack(pady=5)
    script5_check.pack(pady=5)
    execute_button.pack()
    exit_button.pack(pady=40)
    
# Tkinter App
mainapp = tk.Tk()
mainapp.title("OSINT")
center_window(mainapp, 1000,600)

# Boutons de base
launch_button = tk.Button(mainapp, text="Lancer une recherche OSINT", command=show_genpwd_button)
launch_button.pack(pady=100)
launch_button2 = tk.Button(mainapp, text="Renseigner mes clés APIs", command=open_apikey_file)
launch_button2.pack(pady=20)

# Texte
domain_label2 = tk.Label(mainapp, text="Choisissez les outils que vous souhaitez:")
domain_label2.pack_forget()
domain_label3 = tk.Label(mainapp, text="Saisissez le nom de domaine:")
domain_label3.pack_forget()

# Zone de saisie
domain_entry = tk.Entry(mainapp, validate="key")
domain_entry.bind("<Return>", lambda x: get_domain())

# Bouton de validation
domain_button = tk.Button(mainapp, text="Valider", command=get_domain)

# Label de résultat
result_label = tk.Label(mainapp, text="")
result_label.pack()

# Label de confirmation création dossier
domain_label = tk.Label(mainapp, text="")
domain_label.pack()

# Checkbox pour exécuter les scripts
script1_var = tk.BooleanVar(value=True)
script1_check = tk.Checkbutton(mainapp, text="Shodan", variable=script1_var)
script1_check.pack_forget()

script2_var = tk.BooleanVar(value=True)
script2_check = tk.Checkbutton(mainapp, text="Dnscan", variable=script2_var)
script2_check.pack_forget()

script3_var = tk.BooleanVar(value=True)
script3_check = tk.Checkbutton(mainapp, text="Urlscan", variable=script3_var)
script3_check.pack_forget()

script4_var = tk.BooleanVar(value=True)
script4_check = tk.Checkbutton(mainapp, text="The Harvester", variable=script4_var)
script4_check.pack_forget()

script5_var = tk.BooleanVar(value=True)
script5_check = tk.Checkbutton(mainapp, text="Google Dorks", variable=script5_var)
script5_check.pack_forget()

def execute_scripts():
    global domain, results_dir
    
    # Vérifier si le domaine a été saisi
    if not domain:
        result_label.config(text="Aucun domaine saisi.")
        return
    # Vérifier si les clés API ont été chargées avec succès
    if not shodanapikey:
        print("Erreur lors du chargement des clés API depuis le fichier apikey.txt")
        exit(1)
    # Effacer le texte dans la zone de saisie
    domain_entry.delete("0", "end")
    
    # Exécuter les scripts en fonction de la valeur des checkbox
    if script1_var.get():
        subprocess.run(['python.exe', 'scriptshodan.py', domain, results_dir, shodanapikey])

    if script2_var.get():
        subprocess.run(['python.exe', 'scriptdnscan.py', domain, results_dir])

    if script3_var.get():
        subprocess.run(['python.exe', 'scripturlscan.py', domain, results_dir, urlscanapikey])

    if script4_var.get():
        subprocess.run(['python.exe', 'scriptharvester.py', domain, results_dir])

    if script5_var.get():
        subprocess.run(['python.exe', 'dorks.py', domain, results_dir])

    print("Les resultats ont ete enregistres dans le dossier " + domain)

# Bouton pour exécuter les scripts
execute_button = tk.Button(mainapp, text="Exécuter les scripts", command=execute_scripts)
execute_button.pack_forget()

#Bouton Exit
exit_button = tk.Button(mainapp, text="Exit", command=exit_app, width=20)
exit_button.pack_forget()

# Boucle Tkinter
mainapp.mainloop()
