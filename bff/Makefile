# Récupère le nom du module Go depuis le fichier go.mod
MODULE_NAME := $(shell grep '^module ' go.mod | awk '{print $$2}')

# Tag Docker
TAG := latest

# Cible pour construire l'image Docker
docker-build:
	docker build -t $(MODULE_NAME):$(TAG) .

# Pour vérifier le nom du module extrait
print-module:
	@echo "Nom du module Go: $(MODULE_NAME)"	

# Docker push
docker-push:
	docker push $(MODULE_NAME):$(TAG)