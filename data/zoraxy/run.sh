#!/bin/bash
set -e

echo "Initialisation de Zoraxy pour Home Assistant..."

# Aller dans le dossier où se trouvent les fichiers de Zoraxy
cd /usr/src/app

# Vérifier si le binaire Zoraxy est présent ou le compiler
if [[ ! -f /usr/local/bin/zoraxy ]]; then
  echo "Compilation de Zoraxy..."
  go mod tidy
  go build -o /usr/local/bin/zoraxy
fi

# Démarrage de Zoraxy
echo "Lancement de Zoraxy..."
/usr/local/bin/zoraxy
