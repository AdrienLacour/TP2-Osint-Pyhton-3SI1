# TP2-Osint-Pyhton-3SI1
Adrien LACOUR et Mariam CISSE 

Installer le requirements.txt avec la commande: pip install -r requirements.txt

Afin de faire fonctionner le script pour Google dorks, urlscan et shodan, il faut:
Pour urlscan: 
	Insérer la clé API dans apikey.txt créée sur urlscan.io

Pour shodan: 
	Insérer la clé API dans apikey.txt créée sur le site officiel shodan.io

Pour Google dorks:
	Pour obtenir les clés nécessaires à l'utilisation de l'API de recherche Google, vous devez suivre ces étapes :
	-Créez un compte Google si vous n'en avez pas déjà un.
    	-Accédez à la page Google Developers Console.
    	-Cliquez sur "Select a Project" (Sélectionner un projet) en haut de la page, puis cliquez sur "New Project" (Nouveau projet).
    	-Donnez un nom à votre projet et cliquez sur "Create" (Créer).
    	-Dans le menu de gauche, cliquez sur "Credentials" (Identifiants).
    	-Cliquez sur le bouton "Create credentials" (Créer des identifiants), puis sélectionnez "API Key" (Clé d'API).
    	-Copiez la clé API générée.

	Pour obtenir l'ID de moteur de recherche personnalisé :
    	-Accédez à la page Google Custom Search Engine.
    	-Cliquez sur le bouton "Get started" (Commencer).
    	-Suivez les instructions pour configurer votre moteur de recherche personnalisé.
    	-Copiez l'ID de moteur de recherche personnalisé généré.

	Il faut faire ça pour que la clé fonctionne:
    	Pour vérifier si votre compte Google dispose de l'API de recherche Google activée, vous pouvez suivre les étapes suivantes:
	-Accédez au site https://console.cloud.google.com/.
    	-Connectez-vous à votre compte Google.
    	-Dans le menu de navigation à gauche, sélectionnez "API et services" -> "Bibliothèque".
    	-Dans la barre de recherche, tapez "Custom Search API" et appuyez sur Entrée.
    	-Si l'API est activée, vous devriez voir "Custom Search API" dans la liste des APIs avec un état "Activé".
    	-Si l'API n'est pas activée, cliquez sur le bouton "Activer".
